//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repository

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/s14t284/apple-maitained-bot/domain"
)

type Scraper interface {
	Scrape(url string) (*goquery.Document, error)
	ScrapeMaintainedPage(doc *goquery.Document) ([]domain.Page, error)
}