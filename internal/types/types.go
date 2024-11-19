package types

type LinkInfo struct {
	Protocol  string `parquet:"name=protocol, type=BYTE_ARRAY, convertedtype=UTF8"`
	Host      string `parquet:"name=host, type=BYTE_ARRAY, convertedtype=UTF8"`
	Path      string `parquet:"name=path, type=BYTE_ARRAY, convertedtype=UTF8"`
	Link      string `parquet:"name=link, type=BYTE_ARRAY, convertedtype=UTF8"`
	Title     string `parquet:"name=title, type=BYTE_ARRAY, convertedtype=UTF8"`
	IsInbound bool   `parquet:"name=is_inbound, type=BOOLEAN"`
	Depth     int    `parquet:"name=depth, type=INT32"`
    PageType  string `parquet:"name=page_type, type=BYTE_ARRAY, convertedtype=UTF8"`
}

type CrawledData struct {
	URL   string        `parquet:"name=url, type=BYTE_ARRAY, convertedtype=UTF8"`
	Links []LinkInfo    `parquet:"name=links, type=LIST"`
    PageType string     `parquet:"name=page_type, type=BYTE_ARRAY, convertedtype=UTF8"`
}
