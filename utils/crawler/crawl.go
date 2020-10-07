package crawler

import (
	"github.com/labstack/gommon/log"
	"github.com/s14t284/apple-maitained-bot/parser"
	"github.com/s14t284/apple-maitained-bot/usecase/repository"
	"github.com/s14t284/apple-maitained-bot/utils/scraper"
)

// CrawlMacPage macに関する整備済み品ページをクローリング
func CrawlMacPage(rootURL string, endPoint string, mr repository.MacRepository) {
	doc, err := scraper.GetGoQueryObject(rootURL + endPoint + "mac")
	if err != nil {
		panic(err)
	}

	// 一旦、全て売れていることにする
	// クローリングの際に売れ残っている判定を実施する
	mr.UpdateAllSoldTemporary()

	titles, amounts, hrefs := scraper.ScrapeMaintainedPage(doc)
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
			macInDB.IsSold = false
			log.Infof("Unsold: %s", macInDB.URL)
			err = mr.UpdateMac(macInDB)
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
	doc, err := scraper.GetGoQueryObject(rootURL + endPoint + "ipad")
	if err != nil {
		panic(err)
	}

	// 一旦、全て売れていることにする
	// クローリングの際に売れ残っている判定を実施する
	ir.UpdateAllSoldTemporary()

	titles, amounts, hrefs := scraper.ScrapeMaintainedPage(doc)
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
			ipadInDB.IsSold = false
			log.Infof("Unsold: %s", ipadInDB.URL)
			err = ir.UpdateIPad(ipadInDB)
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
	doc, err := scraper.GetGoQueryObject(rootURL + endPoint + "watch")
	if err != nil {
		panic(err)
	}

	// 一旦、全て売れていることにする
	// クローリングの際に売れ残っている判定を実施する
	wr.UpdateAllSoldTemporary()

	titles, amounts, hrefs := scraper.ScrapeMaintainedPage(doc)
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
			watchInDB.IsSold = false
			log.Infof("Unsold: %s", watchInDB.URL)
			err = wr.UpdateWatch(watchInDB)
		} else {
			err = wr.AddWatch(watch)
		}
		if err != nil {
			log.Errorf(err.Error())
		}
	}
}
