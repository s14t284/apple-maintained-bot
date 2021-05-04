//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

// CrawlerUseCase クローラーのinterface
type CrawlerUseCase interface {
	CrawlMacPage() error
	CrawlIPadPage() error
	CrawlWatchPage() error
}
