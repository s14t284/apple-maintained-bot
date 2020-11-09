//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE
package repository

import "github.com/PuerkitoBio/goquery"

type Scraper interface {
	Scrape(url string) (*goquery.Document, error)
	ScrapeMaintainedPage(doc *goquery.Document) (titles, amounts, hrefs []string)
}