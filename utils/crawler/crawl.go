package crawler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/labstack/gommon/log"
	"github.com/s14t284/apple-maitained-bot/domain"
	"github.com/s14t284/apple-maitained-bot/parser"
	"github.com/s14t284/apple-maitained-bot/usecase/repository"
	"github.com/s14t284/apple-maitained-bot/utils/scraper"
)

func hookToSlack(titles []string, urls []string, productKind string) (err error) {
	// slackへ送るメッセージの整形
	attachments := domain.Attachments{}
	// TODO: channel名なども環境変数から取得
	payload := domain.Payload{
		Channel:   "#random",
		UserName:  "AppleMaintainedBot",
		IconEmoji: ":apple:",
	}
	if len(titles) == 0 {
		payload.Text = productKind + "の整備済み品の追加はありませんでした"
	} else {
		for i := 0; i < len(titles); i++ {
			attachment := domain.Attachment{Title: titles[i], TitleLink: urls[i], Color: "good", AuthorName: "apple"}
			attachments = append(attachments, attachment)
		}
		payload.Text = productKind + "の整備済み品が追加されました"
		payload.Attachments = attachments
	}
	rp, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// webhookへリクエスト
	values := url.Values{}
	values.Set("payload", string(rp))
	req, err := http.NewRequest("POST", os.Getenv("SLACK_WEBHOOK_URL"), strings.NewReader(values.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		fmt.Println(string(rp))
		log.Errorf("status code error: %d %s", resp.StatusCode, resp.StatusCode)
		return
	}
	defer resp.Body.Close()
	return
}

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
	var newProducts, newUrls []string
	for i := 0; i < len(titles); i++ {
		// タイトルなどから情報をパース
		var pageParser parser.IParser = &parser.Parser{Title: titles[i], AmountStr: amounts[i], DetailURL: rootURL + hrefs[i]}
		mac, _ := pageParser.ParseMacPage()
		// すでにDBに格納されているか確認
		isExist, id, err := mr.IsExist(mac)
		if err != nil {
			log.Errorf(err.Error())
		}
		// 格納されている場合、まだ売れていないように戻し、URLを更新
		// 格納されていない場合、情報を追加
		if isExist {
			mac.ID = id
			mac.IsSold = false
			log.Infof("Unsold: %s", mac.URL)
			err = mr.UpdateMac(mac)
		} else {
			err = mr.AddMac(mac)
			newProducts = append(newProducts, titles[i])
			newUrls = append(newUrls, rootURL+hrefs[i])
		}
		if err != nil {
			log.Errorf(err.Error())
		}
	}
	err = hookToSlack(newProducts, newUrls, "mac")
	if err != nil {
		log.Errorf(err.Error())
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
	var newProducts, newUrls []string
	for i := 0; i < len(titles); i++ {
		// タイトルなどから情報をパース
		var pageParser parser.IParser = &parser.Parser{Title: titles[i], AmountStr: amounts[i], DetailURL: rootURL + hrefs[i]}
		ipad, _ := pageParser.ParseIPadPage()
		// すでにDBに格納されているか確認
		isExist, id, err := ir.IsExist(ipad)
		if err != nil {
			log.Errorf(err.Error())
		}
		// 格納されている場合、まだ売れていないように戻し、URLを更新
		// 格納されていない場合、情報を追加
		if isExist {
			ipad.ID = id
			ipad.IsSold = false
			log.Infof("Unsold: %s", ipad.URL)
			err = ir.UpdateIPad(ipad)
		} else {
			err = ir.AddIPad(ipad)
			newProducts = append(newProducts, titles[i])
			newUrls = append(newUrls, rootURL+hrefs[i])
		}
		if err != nil {
			log.Errorf(err.Error())
		}
	}
	err = hookToSlack(newProducts, newUrls, "ipad")
	if err != nil {
		log.Errorf(err.Error())
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
	var newProducts, newUrls []string
	for i := 0; i < len(titles); i++ {
		// タイトルなどから情報をパース
		var pageParser parser.IParser = &parser.Parser{Title: titles[i], AmountStr: amounts[i], DetailURL: rootURL + hrefs[i]}
		watch, _ := pageParser.ParseWatchPage()
		// すでにDBに格納されているか確認
		isExist, id, err := wr.IsExist(watch)
		if err != nil {
			log.Errorf(err.Error())
		}
		// 格納されている場合、まだ売れていないように戻し、URLを更新
		// 格納されていない場合、情報を追加
		if isExist {
			watch.ID = id
			watch.IsSold = false
			log.Infof("Unsold: %s", watch.URL)
			err = wr.UpdateWatch(watch)
		} else {
			err = wr.AddWatch(watch)
			newProducts = append(newProducts, titles[i])
			newUrls = append(newUrls, rootURL+hrefs[i])
		}
		if err != nil {
			log.Errorf(err.Error())
		}
	}
	err = hookToSlack(newProducts, newUrls, "apple watch")
	if err != nil {
		log.Errorf(err.Error())
	}
}
