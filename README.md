# URL Harvester

A web crawling tool that collects URLs from a given domain and stores them in Parquet format.

## Features
- Crawl inbound and outbound URLs from a given domain.
- Store crawled data in Parquet format.
- Support for proxies (e.g., Tor, SOCKS5).
- Configurable maximum crawl depth and page count.
- Saves output to a specified directory, with default paths for macOS and Linux.

## Download the Binary

You can download the latest binary from the [releases page](https://github.com/mihneamanolache/url-harvester/releases) and add it to your `PATH` for easy execution.

### Example:
1. Download the binary for your platform.
2. Add the binary to your `PATH` by running:
```bash
   mv url-harvester /usr/local/bin
```

## Build from Source
To build from source, clone the repository and run the following commands:

```bash
# Step 1: Clone the repository
git clone https://github.com/mihneamanolache/url-harvester
# Step 2: Change the working directory
cd url-harvester
# Step 3: Build the binary
go build -o url-harvester cmd/main.go
# Step 4: Move the binary to your PATH
mv url-harvester /usr/local/bin
```
This will allow you to run url-harvester from anywhere on your system.

## Usage
The URL Harvester tool supports the following flags:

```bash
Usage of url-harvester:
  -depth int
        Maximum crawl depth (default 1)
  -domain string
        Domain to crawl
  -output string
        Output directory (default "/tmp")
  -pages int
        Maximum number of pages to crawl (default 100)
  -proxy string
        Proxy address (e.g., socks5://
```

### Usage example:
```bash
url-harvester -query <domain> -depth <max-depth> -maxpages <max-pages> -proxy <proxy-server>
```

### Output Schema:
```
{
  "URL": "string",          // The base URL of the crawled domain (e.g., "https://www.mihnea.dev")
  "Links": [
    {
      "Protocol": "string",  // Protocol of the link (e.g., "https")
      "Host": "string",      // Host of the link (e.g., "mihnea.dev")
      "Path": "string",      // Path of the link (e.g., "/about")
      "Link": "string",      // Full URL of the link (e.g., "https://mihnea.dev/about")
      "Title": "string",     // The title of the link (e.g., "About Us")
      "IsInbound": "bool",   // Whether the link is inbound (true) or outbound (false)
      "Depth": "int"         // The depth at which the link was found
    }
  ]
}
```

NOTE: By default, the crawled data will be saved as a Parquet file on your Desktop (macOS) or /root/ (Linux). You can change the output directory with the -output flag.

## Contributing
We welcome contributions! Feel free to open issues or submit pull requests to improve the tool.

## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
# url-harvester
