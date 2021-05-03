package controller

// CrawlerController クローラーのinterface
type CrawlerController interface {
	CrawlMacPage() error
	CrawlIPadPage() error
	CrawlWatchPage() error
}
