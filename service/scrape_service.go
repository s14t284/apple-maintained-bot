//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE

package service

import (
	"github.com/s14t284/apple-maitained-bot/domain"

	"github.com/PuerkitoBio/goquery"
)

// ScrapeService スクレイピングを行うオブジェクトのinterface
type ScrapeService interface {
	Scrape(targetPath string) (*goquery.Document, error)
	ScrapeMaintainedPage(doc *goquery.Document) ([]domain.Page, error)
}
