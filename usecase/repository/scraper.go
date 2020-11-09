package repository

import "github.com/PuerkitoBio/goquery"

type Scraper interface {
	Scrape(url string) (*goquery.Document, error)
	ScrapeMaintainedPage(doc *goquery.Document) (titles, amounts, hrefs []string)
}