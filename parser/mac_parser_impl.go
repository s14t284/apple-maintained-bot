package parser

import (
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/utils"
)

func loadMacInformationFromDetailURL(mac *model.Mac, doc *goquery.Document) {
	detail := doc.Find(".as-productinfosection-mainpanel").First()
	detailRegExp, _ := regexp.Compile(`(\n|\s)`)
	storageRegExp, _ := regexp.Compile(`(T|G)B`)
	detail.Find("div .para-list > p").Each(func(_ int, s *goquery.Selection) {
		text := detailRegExp.ReplaceAllLiteralString(s.Text(), "")
		if strings.Index(text, "月発売") > -1 {
			// 発売年月
			year, _ := strconv.Atoi(text[:4])
			month, _ := strconv.Atoi(text[strings.Index(text, "年")+4 : strings.Index(text, "月")])
			timeZone, err := time.LoadLocation("Asia/Tokyo")
			if err != nil {
				log.Errorf(err.Error())
			}
			mac.ReleaseDate = time.Date(year, time.Month(month), 1, 9, 0, 0, 0, timeZone)
		} else if strings.Index(text, "TouchBar") > -1 {
			// タッチバーがあるかないか
			mac.TouchBar = true
		} else if strings.Index(text, "SSD") > -1 {
			// ストレージ容量
			mac.Strage = text[:storageRegExp.FindStringIndex(text)[0]+2]
		} else if strings.Index(text, "GB") > -1 && mac.Memory == 0 {
			// メモリ
			mac.Memory, _ = strconv.Atoi(text[:strings.Index(text, "GB")])
		}
	})
}

// ParseMacPage macに関するページをパースして、macに関する情報のオブジェクトを返却
func (parser *Parser) ParseMacPage() (*model.Mac, error) {
	var mac model.Mac
	// 最初に詳細情報が取ってこれるかを確認
	doc, err := utils.GetGoQueryObject(parser.DetailURL)
	if err != nil {
		log.Errorf("cannot open detail product page. url: %s", parser.DetailURL)
		return nil, err
	}
	// ノートパソコン以外は飛ばす
	// TODO: デスクトップにも対応
	if strings.Index(parser.Title, "MacBook") == -1 {
		mac.Name = parser.Title
		return &mac, nil
	}
	// 不要な部分を削除
	nameRegExp, _ := regexp.Compile(`Retinaディスプレイ.*\s\-`)
	name := nameRegExp.ReplaceAllLiteralString(strings.Replace(parser.Title, " [整備済製品]", "", 1), "")
	// インチ数
	strs := strings.Split(name, "インチ")
	mac.Inch, _ = strconv.ParseFloat(strs[0], 64)
	name = strs[1]
	// CPU
	cpuRegExp, _ := regexp.Compile(`\d+\.\dGHz.+i\d`)
	mac.CPU = cpuRegExp.FindString(name)
	name = strings.Replace(name, mac.CPU+" ", "", 1)
	// 色
	strs = strings.Split(name, "  ")
	name = strs[0]
	mac.Color = strs[1]
	// 金額
	amountRegExp, _ := regexp.Compile(`(,|円（税別）|\s)`)
	mac.Amount, _ = strconv.Atoi(amountRegExp.ReplaceAllLiteralString(parser.AmountStr, ""))
	// 名前
	mac.Name = name
	// URL
	mac.URL = parser.DetailURL

	// その他の情報
	loadMacInformationFromDetailURL(&mac, doc)
	return &mac, nil
}
