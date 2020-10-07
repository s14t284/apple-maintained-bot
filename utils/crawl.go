package utils

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
	"github.com/s14t284/apple-maitained-bot/parser"
	"github.com/s14t284/apple-maitained-bot/usecase/repository"
)

const userAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1"

// GetGoQueryObject 引数で指定したURLにアクセスしてそのURLのHTML情報を取得
func GetGoQueryObject(requestURL string) (*goquery.Document, error) {
	requestBody := url.Values{}
	req, err := http.NewRequest("GET", requestURL, strings.NewReader(requestBody.Encode()))
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	req.Header.Add("User-Agent", userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		log.Errorf("status code error: %d %s", resp.StatusCode, resp.StatusCode)
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	return doc, nil
}

// CrawlMacPage macに関する整備済み品ページをクローリング
func CrawlMacPage(rootURL string, endPoint string, mr repository.MacRepository) {
	doc, err := GetGoQueryObject(rootURL + endPoint + "mac")
	if err != nil {
		panic(err)
	}

	// 一旦、全て売れていることにする
	// クローリングの際に売れ残っている判定を実施する
	mr.UpdateAllSoldTemporary()

	titles, amounts, hrefs := scrapeMaintainedPage(doc)
	for i := 0; i < len(titles); i++ {
		// タイトルなどから情報をパース
		var pageParser parser.IParser = &parser.Parser{Title: titles[i], AmountStr: amounts[i], DetailURL: rootURL + hrefs[i]}
		mac, _ := pageParser.ParseMacPage()
		// すでにDBに格納されているか確認
		macInDB, err := mr.FindByURL(mac.URL)
		if err != nil {
			log.Errorf(err.Error())
		}
		// 格納されている場合、まだ売れていないように戻す
		// 格納されていない場合、情報を追加
		if macInDB != nil {
			mac.IsSold = false
			log.Infof("Unsold: %s", mac.URL)
			err = mr.UpdateMac(mac)
		} else {
			err = mr.AddMac(mac)
		}
		if err != nil {
			log.Errorf(err.Error())
		}
	}
}

// CrawlIPadPage ipadに関する整備済み品ページをクローリング
func CrawlIPadPage(rootURL string, endPoint string, ir repository.IPadRepository) {
	doc, err := GetGoQueryObject(rootURL + endPoint + "ipad")
	if err != nil {
		panic(err)
	}

	// 一旦、全て売れていることにする
	// クローリングの際に売れ残っている判定を実施する
	ir.UpdateAllSoldTemporary()

	titles, amounts, hrefs := scrapeMaintainedPage(doc)
	for i := 0; i < len(titles); i++ {
		// タイトルなどから情報をパース
		var pageParser parser.IParser = &parser.Parser{Title: titles[i], AmountStr: amounts[i], DetailURL: rootURL + hrefs[i]}
		ipad, _ := pageParser.ParseIPadPage()
		// すでにDBに格納されているか確認
		ipadInDB, err := ir.FindByURL(ipad.URL)
		if err != nil {
			log.Errorf(err.Error())
		}
		// 格納されている場合、まだ売れていないように戻す
		// 格納されていない場合、情報を追加
		if ipadInDB != nil {
			ipad.IsSold = false
			log.Infof("Unsold: %s", ipad.URL)
			err = ir.UpdateIPad(ipad)
		} else {
			err = ir.AddIPad(ipad)
		}
		if err != nil {
			log.Errorf(err.Error())
		}
	}
}

// CrawlWatchPage watchに関する整備済み品ページをクローリング
func CrawlWatchPage(rootURL string, endPoint string, wr repository.WatchRepository) {
	doc, err := GetGoQueryObject(rootURL + endPoint + "watch")
	if err != nil {
		panic(err)
	}

	// 一旦、全て売れていることにする
	// クローリングの際に売れ残っている判定を実施する
	wr.UpdateAllSoldTemporary()

	titles, amounts, hrefs := scrapeMaintainedPage(doc)
	for i := 0; i < len(titles); i++ {
		// タイトルなどから情報をパース
		var pageParser parser.IParser = &parser.Parser{Title: titles[i], AmountStr: amounts[i], DetailURL: rootURL + hrefs[i]}
		watch, _ := pageParser.ParseWatchPage()
		// すでにDBに格納されているか確認
		watchInDB, err := wr.FindByURL(watch.URL)
		if err != nil {
			log.Errorf(err.Error())
		}
		// 格納されている場合、まだ売れていないように戻す
		// 格納されていない場合、情報を追加
		if watchInDB != nil {
			watch.IsSold = false
			log.Infof("Unsold: %s", watch.URL)
			err = wr.UpdateWatch(watch)
		} else {
			err = wr.AddWatch(watch)
		}
		if err != nil {
			log.Errorf(err.Error())
		}
	}
}

// ScrapeMaintainedPage 整備済み品ページから共通の情報を取得
func scrapeMaintainedPage(doc *goquery.Document) (titles []string, amounts []string, hrefs []string) {
	// 一件ずつスクレイピング
	doc.Find("div .refurbished-category-grid-no-js > ul > li").Each(func(_ int, s *goquery.Selection) {
		// タイトル、金額、詳細ページへのURL
		title := s.Find("h3 > a").Text()
		amount := s.Find("div,.as-currentprice,.producttile-currentprice").Text()
		href, _ := s.Find("a").Attr("href")
		// 格納
		titles = append(titles, title)
		amounts = append(amounts, amount)
		hrefs = append(hrefs, href)
	})
	return
}
