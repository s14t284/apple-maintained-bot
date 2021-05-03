package web

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/s14t284/apple-maitained-bot/domain"
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/utils"
)

const (
	SALES = "月発売"
	INCH  = "インチ"
)

// PageParserImpl 商品ページのパーサーの実装
type PageParserImpl struct {
}

// NewPageParserImpl PageParserImplを初期化
func NewPageParserImpl() (*PageParserImpl, error) {
	return &PageParserImpl{}, nil
}

// ParsePage 商品ページのパース
func (ppi *PageParserImpl) ParsePage(target string, page domain.Page) (interface{}, error) {
	switch target {
	case "mac":
		return ppi.parseMacPage(page)
	case "ipad":
		return ppi.parseIPadPage(page)
	case "watch":
		return ppi.parseWatchPage(page)
	default:
		return nil, fmt.Errorf("target must be `mac`, `ipad`, or `watch`")
	}
}

// parseMacPage macに関するページをパースして、macに関する情報のオブジェクトを返却
// TODO: macbook以外にも対応
func (ppi *PageParserImpl) parseMacPage(page domain.Page) (*model.Mac, error) {
	var mac model.Mac
	err := ppi.loadMacInformationFromTitle(&mac, page)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mac page [error][%w]", err)
	}
	err = ppi.loadMacInformationFromDetailHTML(&mac, page.Document)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mac page [error][%w]", err)
	}
	return &mac, nil
}

// parseMacPage ipadに関するページをパースして、ipadに関する情報のオブジェクトを返却
func (ppi *PageParserImpl) parseIPadPage(page domain.Page) (*model.IPad, error) {
	var ipad model.IPad
	err := ppi.loadIPadInformationFromTitle(&ipad, page)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mac page [error][%w]", err)
	}
	err = ppi.loadIPadInformationFromDetailHTML(&ipad, page.Document)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mac page [error][%w]", err)
	}
	return &ipad, nil
}

// parseWatchPage watchに関するページをパースして、watchに関する情報のオブジェクトを返却
func (ppi *PageParserImpl) parseWatchPage(page domain.Page) (*model.Watch, error) {
	var watch model.Watch
	err := ppi.loadWatchInformationFromTitle(&watch, page)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mac page [error][%w]", err)
	}
	err = ppi.loadWatchInformationFromDetailHTML(&watch, page.Document)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mac page [error][%w]", err)
	}
	return &watch, nil
}

// loadMacInformationFromDetailHTML 詳細ページのHTMLから情報を取得する
func (ppi *PageParserImpl) loadMacInformationFromDetailHTML(mac *model.Mac, doc *goquery.Document) error {
	detail := doc.Find(".as-productinfosection-mainpanel").First()
	detailRegExp, err := regexp.Compile(`(\n|\s)`)
	if err != nil {
		return err
	}
	storageRegExp, err := regexp.Compile(`(T|G)B`)
	if err != nil {
		return err
	}

	var localErr error
	detail.Find("div .para-list > p").Each(func(_ int, s *goquery.Selection) {
		text := detailRegExp.ReplaceAllLiteralString(s.Text(), "")
		if strings.Contains(text, SALES) {
			// 発売年月
			year, localErr := strconv.Atoi(text[:4])
			if localErr != nil {
				err = fmt.Errorf("failed to parse year [error][%w]", localErr)
			}
			month, localErr := strconv.Atoi(text[strings.Index(text, "年")+3 : strings.Index(text, "月")])
			if err != nil {
				err = fmt.Errorf("failed to parse month [error][%w]", localErr)
			}
			mac.ReleaseDate = utils.GetReleaseYearAndMonth(year, month)
		} else if strings.Contains(text, "TouchBar") {
			// タッチバーがあるかないか
			mac.TouchBar = true
		} else if strings.Contains(text, "SSD") {
			// ストレージ容量
			mac.Storage, localErr = parseStorage(text[:storageRegExp.FindStringIndex(text)[0]+2])
			if localErr != nil {
				err = fmt.Errorf("failed to parse storage [error][%w]", localErr)
			}
		} else if strings.Contains(text, "GB") && mac.Memory == 0 {
			// メモリ
			mac.Memory, localErr = strconv.Atoi(text[:strings.Index(text, "GB")])
			if localErr != nil {
				err = fmt.Errorf("failed to parse memory [error][%w]", localErr)
			}
		}
	})
	return err
}

// loadMacInformationFromTitle タイトルから情報を取得する
func (ppi *PageParserImpl) loadMacInformationFromTitle(mac *model.Mac, page domain.Page) error {
	nameRegExp, err := regexp.Compile(`Retinaディスプレイ.*\s\-`)
	if err != nil {
		return err
	}
	name := nameRegExp.ReplaceAllLiteralString(strings.Replace(page.Title, " [整備済製品]", "", 1), "")
	if strings.Contains(page.Title, "MacBook") {
		// インチ数
		strs := strings.Split(name, INCH)
		inch, err := strconv.ParseFloat(strs[0], 32)
		if err != nil {
			return fmt.Errorf("failed to parse inch [error][%w]", err)
		}
		mac.Inch = float32(inch)
		name = strs[1]
		// CPU
		cpuRegExp, err := regexp.Compile(`\d+\.\dGHz.+i\d`)
		if err != nil {
			return err
		}
		mac.CPU = cpuRegExp.FindString(name)
		name = strings.Replace(name, mac.CPU+" ", "", 1)
		// 色
		strs = strings.Split(name, "  ")
		name = strs[0]
		mac.Color = strs[1]

	}
	if strings.Contains(page.Title, "Mac mini") {
		// CPU
		cpuRegExp, err := regexp.Compile(`\d+\.\dGHz.+i\d`)
		if err != nil {
			return err
		}
		mac.CPU = cpuRegExp.FindString(name)
		name = strings.Replace(name, mac.CPU+" ", "", 1)
		// 色
		strs := strings.Split(name, " - ")
		name = strs[0]
		mac.Color = strs[1]

	} else {
		// TODO: iMacやMacBookAirにも対応させる
		mac.Name = page.Title
	}
	// 金額
	amountRegExp, err := regexp.Compile(`(,|円（税別）|\s)`)
	if err != nil {
		return err
	}
	mac.Amount, err = strconv.Atoi(amountRegExp.ReplaceAllLiteralString(page.AmountStr, ""))
	if err != nil {
		return fmt.Errorf("failed to parse amount [error][%w]", err)
	}
	// 名前
	mac.Name = name
	// URL
	mac.URL = page.DetailURL
	if !strings.HasPrefix(mac.URL, rootURL) {
		mac.URL = rootURL + mac.URL
	}
	return nil
}

// loadIPadInformationFromDetailHTML 詳細ページのHTMLから情報を取得する
func (ppi *PageParserImpl) loadIPadInformationFromDetailHTML(ipad *model.IPad, doc *goquery.Document) error {
	detail := doc.Find(".as-productinfosection-mainpanel").First()
	detailRegExp, err := regexp.Compile(`(\n|\s)`)
	if err != nil {
		return err
	}

	detail.Find("div .para-list > p").Each(func(_ int, s *goquery.Selection) {
		text := detailRegExp.ReplaceAllLiteralString(s.Text(), "")
		if strings.Contains(text, SALES) {
			// 発売年月
			year, localErr := strconv.Atoi(text[:4])
			if localErr != nil {
				err = fmt.Errorf("failed to parse year [error][%w]", localErr)
			}
			month, localErr := strconv.Atoi(text[strings.Index(text, "年")+3 : strings.Index(text, "月")])
			if localErr != nil {
				err = fmt.Errorf("failed to parse year [error][%w]", localErr)
			}
			ipad.ReleaseDate = utils.GetReleaseYearAndMonth(year, month)
		} else if strings.Contains(text, "メガピクセル") {
			ipad.Camera = text
		} else if strings.Contains(text, INCH) {
			strs := strings.Split(text, INCH)
			inch, localErr := strconv.ParseFloat(strs[0], 32)
			if localErr != nil {
				err = fmt.Errorf("failed to parse year [error][%w]", localErr)
			}
			ipad.Inch = float32(inch)
		}
	})
	return err
}

// loadIPadInformationFromTitle タイトルから情報を取得する
func (ppi *PageParserImpl) loadIPadInformationFromTitle(ipad *model.IPad, page domain.Page) error {
	// 不要な部分を削除
	nameRegExp, err := regexp.Compile(`\s*(\(|\[|（).+(\])`)
	if err != nil {
		return err
	}
	name := nameRegExp.ReplaceAllLiteralString(page.Title, "")
	strs := strings.Split(name, INCH)
	name = strs[len(strs)-1]
	// 色
	strs = strings.Split(name, " - ")
	name = strs[0]
	ipad.Color = strs[1]
	// 名前・ストレージ
	if strings.Contains(name, "Wi-Fi + Cellular") {
		strs = strings.Split(name, " Wi-Fi + Cellular ")
	} else if strings.Contains(name, "Wi-Fiモデル") {
		strs = strings.Split(name, " Wi-Fiモデル ")
	} else {
		strs = strings.Split(name, " Wi-Fi ")
	}
	ipad.Name = strs[0]
	ipad.Storage, err = parseStorage(strs[1])
	if err != nil {
		return fmt.Errorf("failed to parse storage [error][%w]", err)
	}
	// 金額
	amountRegExp, err := regexp.Compile(`(,|円（税別）|\s)`)
	if err != nil {
		return err
	}
	ipad.Amount, err = strconv.Atoi(amountRegExp.ReplaceAllLiteralString(page.AmountStr, ""))
	if err != nil {
		return fmt.Errorf("failed to parse amount [error][%w]", err)
	}
	// URL
	ipad.URL = page.DetailURL
	if !strings.HasPrefix(ipad.URL, rootURL) {
		ipad.URL = rootURL + ipad.URL
	}
	return nil
}

// loadWatchInformationFromDetailHTML 詳細ページのHTMLから情報を取得する
func (ppi *PageParserImpl) loadWatchInformationFromDetailHTML(watch *model.Watch, doc *goquery.Document) error {
	detail := doc.Find(".as-productinfosection-mainpanel").First()
	detailRegExp, err := regexp.Compile(`(\n|\s)`)
	if err != nil {
		return err
	}

	detail.Find("div .para-list > p").Each(func(_ int, s *goquery.Selection) {
		text := detailRegExp.ReplaceAllLiteralString(s.Text(), "")
		if strings.Contains(text, SALES) {
			// 発売年月
			year, localErr := strconv.Atoi(text[:4])
			if localErr != nil {
				err = fmt.Errorf("failed to parse year [error][%w]", localErr)
			}
			month, localErr := strconv.Atoi(text[strings.Index(text, "年")+3 : strings.Index(text, "月")])
			if localErr != nil {
				err = fmt.Errorf("failed to parse year [error][%w]", localErr)
			}
			watch.ReleaseDate = utils.GetReleaseYearAndMonth(year, month)
		} else if strings.Contains(text, "GB") {
			// ストレージ
			r, localErr := regexp.Compile(`[0-9]+GB`)
			if localErr != nil {
				err = localErr
			}
			watch.Storage, localErr = parseStorage(r.FindString(text))
			if localErr != nil {
				err = fmt.Errorf("failed to parse storage [error][%w]", localErr)
			}
		}
	})
	return err
}

// loadWatchInformationFromTitle タイトルから情報を取得する
func (ppi *PageParserImpl) loadWatchInformationFromTitle(watch *model.Watch, page domain.Page) error {
	// 不要な部分を削除
	nameRegExp, err := regexp.Compile(`\s*(（.+）|\[.+\])`)
	if err != nil {
		return err
	}
	name := nameRegExp.ReplaceAllLiteralString(page.Title, "")
	// Cellularモデルかどうか
	if strings.Contains(page.Title, "Cellular") {
		watch.IsCellular = true
	}
	// 名前・色
	colorRegExp, err := regexp.Compile(`\d+mm`)
	if err != nil {
		return err
	}
	strs := strings.Split(name, "- ")
	watch.Name = strs[0]
	watch.Color = colorRegExp.ReplaceAllLiteralString(strs[1], "")
	// 金額
	amountRegExp, err := regexp.Compile(`(,|円（税別）|\s)`)
	if err != nil {
		return err
	}
	watch.Amount, err = strconv.Atoi(amountRegExp.ReplaceAllLiteralString(page.AmountStr, ""))
	if err != nil {
		return fmt.Errorf("failed to parse amount [error][%w]", err)
	}
	// URL
	watch.URL = page.DetailURL
	if !strings.HasPrefix(watch.URL, rootURL) {
		watch.URL = rootURL + watch.URL
	}
	return nil
}

func parseStorage(storageStr string) (int, error) {
	var str, suffix string
	var coef int
	if strings.HasSuffix(storageStr, "GB") {
		coef = 1
		suffix = "GB"
	} else if strings.HasSuffix(storageStr, "TB") {
		coef = 1000
		suffix = "TB"
	} else {
		return -1, fmt.Errorf("suffix of storage string must have 'GB' or 'TB'")
	}
	str = strings.Replace(storageStr, suffix, "", 1)

	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return -1, fmt.Errorf("failed to parse storage stirng to int [error][%w]", err)
	}
	return int(val) * coef, nil
}
