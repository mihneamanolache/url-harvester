package main

import (
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/mihneamanolache/url-harvester/internal/types"
	"github.com/playwright-community/playwright-go"
	"github.com/xitongsys/parquet-go-source/local"
	"github.com/xitongsys/parquet-go/writer"
)

var (
	maxDepth        *int
	inboundCount    int
	outboundCount   int
)

var mu sync.Mutex

func isCrawled(crawled map[string]bool, host, path string) bool {
	key := host + path
	return crawled[key]
}

func markAsCrawled(crawled map[string]bool, host, path string) {
	key := host + path
	crawled[key] = true
}

func crawlPage(baseURL string, depth int, crawled map[string]bool, parsedURL *url.URL, parquetWriter *writer.ParquetWriter, urlQueue chan *url.URL, proxy *playwright.Proxy) {
	if depth > *maxDepth {
		return
	}

	pw, err := playwright.Run()
	if err != nil {
		log.Printf("Could not start Playwright: %v", err)
		return
	}
	defer pw.Stop()

	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Proxy: proxy,
	})
	if err != nil {
		log.Printf("Could not launch browser: %v", err)
		return
	}
	defer browser.Close()

	page, err := browser.NewPage()
	if err != nil {
		log.Printf("Could not create page: %v", err)
		return
	}
	defer page.Close()

	fmt.Println("Navigating to", parsedURL.String())
	if _, err := page.Goto(parsedURL.String()); err != nil {
		log.Printf("Error navigating to %s: %v", parsedURL.String(), err)
		return
	}

	// Get the final URL after redirection
	finalURL := page.URL()
	parsedURL, err = url.Parse(finalURL)
	if err != nil {
		log.Printf("Error parsing final URL %v: %v", finalURL, err)
		return
	}

	// After navigation, collect links from the page
	entries, err := page.Locator("a").All()
	if err != nil {
		log.Printf("Could not get entries: %v", err)
		return
	}

	var linkInfos []types.LinkInfo
	for _, entry := range entries {
		href, err := entry.GetAttribute("href")
		if err != nil || href == "" {
			continue
		}

		if strings.HasPrefix(href, "mailto:") || strings.HasPrefix(href, "tel:") {
			continue
		}

		if strings.HasPrefix(href, "/") {
			href = parsedURL.Scheme + "://" + parsedURL.Host + href
		}

		relURL, err := url.Parse(href)
		if err != nil {
			log.Printf("Warning: Could not parse URL: %v", err)
			continue
		}

		isInbound := relURL.Host == parsedURL.Host

		mu.Lock()
		if isInbound {
			inboundCount++
		} else {
			outboundCount++
		}
		mu.Unlock()

		mu.Lock()
		if isCrawled(crawled, relURL.Host, relURL.Path) {
			mu.Unlock()
			continue
		}
		markAsCrawled(crawled, relURL.Host, relURL.Path)
		mu.Unlock()

		title, _ := entry.InnerText()

		linkInfo := types.LinkInfo{
			Protocol:  relURL.Scheme,
			Host:      relURL.Host,
			Path:      relURL.Path,
			Link:      relURL.String(),
			Title:     title,
			IsInbound: isInbound,
			Depth:     depth + 1,
		}

		linkInfos = append(linkInfos, linkInfo)
		if isInbound {
			urlQueue <- relURL
		}
	}

	crawledData := types.CrawledData{
		URL:   baseURL,
		Links: linkInfos,
	}

	if err := parquetWriter.Write(crawledData); err != nil {
		log.Printf("Failed to write to Parquet file: %v", err)
	}
}

func parseProxy(proxyStr string) (*playwright.Proxy, error) {
	proxyURL, err := url.Parse(proxyStr)
	if err != nil {
		return nil, fmt.Errorf("invalid proxy URL: %w", err)
	}

	if proxyURL.Scheme == "" {
		proxyURL.Scheme = "http"
	}

	var username, password string
	if proxyURL.User != nil {
		username = proxyURL.User.Username()
		password, _ = proxyURL.User.Password()
	}

	hostParts := strings.Split(proxyURL.Host, ":")
	if len(hostParts) != 2 {
		return nil, fmt.Errorf("invalid proxy host:port format")
	}

	return &playwright.Proxy{
		Server:   proxyURL.Scheme + "://" + hostParts[0] + ":" + hostParts[1],
		Username: &username,
		Password: &password,
	}, nil
}

func main() {
	query := flag.String("query", "", "Initial URL/domain to crawl")
	proxyServer := flag.String("proxy", "", "Proxy server URL (e.g., socks5://127.0.0.1:9050)")
	maxDepth = flag.Int("depth", 3, "Maximum depth for crawling")
	maxPages := flag.Int("maxpages", 500, "Maximum number of pages to crawl")
	outputPath := flag.String("output", "", "Output directory for Parquet file")

	flag.Parse()

	if *outputPath == "" {
		if strings.Contains(runtime.GOOS, "darwin") { 
			*outputPath = fmt.Sprintf("/Users/%s/Desktop", os.Getenv("USER"))
		} else { // Linux
			*outputPath = "/root"
		}
	}

	if _, err := os.Stat(*outputPath); os.IsNotExist(err) {
		log.Fatalf("Output directory does not exist: %v", err)
	}

	var proxy *playwright.Proxy
	if *proxyServer != "" {
		proxy, _ = parseProxy(*proxyServer)
	}

	baseURL := *query
	if !strings.HasPrefix(baseURL, "http") {
		baseURL = "https://" + baseURL
	}
	if !strings.HasSuffix(baseURL, "/") {
		baseURL = baseURL + "/"
	}

	parsedBaseURL, err := url.Parse(baseURL)
	if err != nil {
		log.Fatalf("Invalid URL: %v", err)
	}

	crawled := make(map[string]bool)
	markAsCrawled(crawled, parsedBaseURL.Host, parsedBaseURL.Path)

	parquetFileName := fmt.Sprintf("%s.parquet", strings.Replace(parsedBaseURL.Host, ".", "_", -1))
	parquetFile, err := local.NewLocalFileWriter(fmt.Sprintf("%s/%s", *outputPath, parquetFileName))
	if err != nil {
		log.Fatalf("Failed to open Parquet file: %v", err)
	}
	defer parquetFile.Close()

	parquetWriter, err := writer.NewParquetWriter(parquetFile, new(types.CrawledData), 4)
	if err != nil {
		log.Fatalf("Failed to create Parquet writer: %v", err)
	}
	defer parquetWriter.WriteStop()

	// Initialize the URL queue and start crawling
	urlQueue := make(chan *url.URL, *maxPages)
	urlQueue <- parsedBaseURL

	pageCount := 0
	for pageCount < *maxPages {
		select {
		case nextURL := <-urlQueue:
			pageCount++
			crawlPage(nextURL.String(), 0, crawled, nextURL, parquetWriter, urlQueue, proxy)
		default:
			fmt.Println("Done crawling.")
			fmt.Printf("Total Inbound URLs: %d\n", inboundCount)
			fmt.Printf("Total Outbound URLs: %d\n", outboundCount)
            fmt.Printf("File saved to: %s/%s\n", *outputPath, parquetFileName)
			return
		}
	}
}
