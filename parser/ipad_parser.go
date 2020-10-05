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

// IIpadParser ipadに関するページのインタフェース
type IIpadParser interface {
	ParseIPadPage() (*model.IPad, error)
}

// IPadParser ipadに関するページのパーサー
type IPadParser struct {
	Title     string
	AmountStr string
	DetailURL string
}

func loadIPadInformationFromDetailURL(ipad *model.IPad, doc *goquery.Document) {
	detail := doc.Find(".as-productinfosection-mainpanel").First()
	detailRegExp, _ := regexp.Compile(`(\n|\s)`)
	detail.Find("div .para-list > p").Each(func(_ int, s *goquery.Selection) {
		text := detailRegExp.ReplaceAllLiteralString(s.Text(), "")
		if strings.Index(text, "月発売") > -1 {
			// 発売年月
			year, _ := strconv.Atoi(text[:4])
			month, _ := strconv.Atoi(text[strings.Index(text, "年"):strings.Index(text, "月")])
			ipad.ReleaseDate = time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
		} else if strings.Index(text, "メガピクセル") > -1 {
			ipad.Camera = text
		} else if strings.Index(text, "インチ") > -1 {
			strs := strings.Split(text, "インチ")
			inch, _ := strconv.ParseFloat(strs[0], 32)
			ipad.Inch = float32(inch)
		}
	})
}

// ParseIPadPage ipadに関するページをパースして、ipadに関する情報のオブジェクトを返却
func (parser *IPadParser) ParseIPadPage() (*model.IPad, error) {
	var ipad model.IPad
	// 最初に詳細情報が取ってこれるかを確認
	doc, err := utils.GetGoQueryObject(parser.DetailURL)
	if err != nil {
		log.Errorf("cannot open detail product page. url: %s", parser.DetailURL)
		return nil, err
	}
	// 不要な部分を削除
	nameRegExp, _ := regexp.Compile(`\s*(\(|\[|（).+(\])`)
	name := nameRegExp.ReplaceAllLiteralString(parser.Title, "")
	strs := strings.Split(name, "インチ")
	name = strs[len(strs)-1]
	// 色
	strs = strings.Split(name, " - ")
	name = strs[0]
	ipad.Color = strs[1]
	// 名前・ストレージ
	if strings.Index(name, "Wi-Fi + Cellular") > -1 {
		strs = strings.Split(name, " Wi-Fi + Cellular ")
	} else if strings.Index(name, "Wi-Fiモデル") > -1 {
		strs = strings.Split(name, " Wi-Fiモデル ")
	} else {
		strs = strings.Split(name, " Wi-Fi ")
	}
	ipad.Name = strs[0]
	ipad.Strage = strs[1]
	// 金額
	amountRegExp, _ := regexp.Compile(`(,|円（税別）|\s)`)
	ipad.Amount, _ = strconv.Atoi(amountRegExp.ReplaceAllLiteralString(parser.AmountStr, ""))

	// その他の情報
	loadIPadInformationFromDetailURL(&ipad, doc)
	return &ipad, nil
}
