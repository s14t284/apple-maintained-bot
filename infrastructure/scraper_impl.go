package infrastructure

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
	"github.com/s14t284/apple-maitained-bot/domain"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const userAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1"

type ScraperImpl struct {
	client *http.Client
}

// NewScraperImpl ScraperImplの初期化
func NewScraperImpl(client *http.Client) (*ScraperImpl, error) {
	if client == nil {
		return nil, fmt.Errorf("http client is nil")
	}
	return &ScraperImpl{client: client}, nil
}

// Scrape 指定したurlをgoqueryオブジェクトに変換するメソッド
func (si *ScraperImpl) Scrape(targetURL string) (doc *goquery.Document, err error) {
	requestBody := url.Values{}
	req, err := http.NewRequest("GET", targetURL, strings.NewReader(requestBody.Encode()))
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	req.Header.Add("User-Agent", userAgent)

	resp, _ := si.client.Do(req)
	if resp.StatusCode != 200 {
		log.Errorf("status code error: %d %s", resp.StatusCode, resp.Status)
		return nil, fmt.Errorf("cannot access error")
	}
	defer func() {
		closeErr := resp.Body.Close()
		if err == nil {
			err = closeErr
		}
	}()

	// doc, err := goquery.NewDocumentFromResponse(resp)
	doc, err = goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	return doc, nil
}

// ScrapeMaintainedPage 整備済み品ページから共通の情報を取得
func (si *ScraperImpl) ScrapeMaintainedPage(doc *goquery.Document) ([]domain.Page, error) {
	// 一件ずつスクレイピング
	pages := make([]domain.Page, 0)
	var err error
	doc.Find("div .refurbished-category-grid-no-js > ul > li").Each(func(_ int, s *goquery.Selection) {
		// タイトル、金額、詳細ページへのURL
		title := s.Find("h3 > a").Text()
		amount := s.Find("div,.as-currentprice,.producttile-currentprice").Text()
		href, _ := s.Find("a").Attr("href")
		detailDoc, localErr := si.Scrape(href)
		if localErr != nil {
			err = fmt.Errorf("failed to scrape detail page [url][%s][error][%w]", href, localErr)
			log.Errorf(err.Error())
		}
		// 格納
		page := domain.Page{
			Title:     title,
			AmountStr: amount,
			DetailURL: href,
			Document:  detailDoc,
		}
		pages = append(pages, page)
		time.Sleep(time.Second)
	})
	return pages, err
}
