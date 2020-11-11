//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package interfaces

type CrawlerController interface {
	CrawlMacPage(rootURL, endPoint string) error
	CrawlIPadPage(rootURL, endPoint string) error
	CrawlWatchPage(rootURL, endPoint string) error
}
