package interfaces

type CrawlerController interface {
	CrawlMacPage() error
	CrawlIPadPage() error
	CrawlWatchPage() error
}
