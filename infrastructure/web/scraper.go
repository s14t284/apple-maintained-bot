//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package web

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/s14t284/apple-maitained-bot/domain"
)

// Scraper スクレイピングを行うオブジェクトのinterface
type Scraper interface {
	Scrape(targetPath string) (*goquery.Document, error)
	ScrapeMaintainedPage(doc *goquery.Document) ([]domain.Page, error)
}
