package usecase

// CrawlerUseCase クローラーのinterface
type CrawlerUseCase interface {
	CrawlMacPage() error
	CrawlIPadPage() error
	CrawlWatchPage() error
}
