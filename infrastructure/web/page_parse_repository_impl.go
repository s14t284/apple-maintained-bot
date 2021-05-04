package web

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/s14t284/apple-maitained-bot/domain"
	"github.com/s14t284/apple-maitained-bot/domain/model"

	"github.com/PuerkitoBio/goquery"
)

// PageParseRepositoryImpl 商品ページごとのパーサーの詳細
type PageParseRepositoryImpl struct {
}

var _ PageParseRepository = &PageParseRepositoryImpl{}

const (
	SALES   = "月発売"
	INCH    = "インチ"
	rootURL = "https://www.apple.com"
)

// NewPageParseRepositoryImpl PageParseRepositoryImplを初期化
func NewPageParseRepositoryImpl() (*PageParseRepositoryImpl, error) {
	return &PageParseRepositoryImpl{}, nil
}

// ParseMacPage macに関するページをパースして、macに関する情報のオブジェクトを返却
// TODO: macbook以外にも対応
func (ppri *PageParseRepositoryImpl) ParseMacPage(page domain.Page) (*model.Mac, error) {
	var mac model.Mac
	err := ppri.loadMacInformationFromTitle(&mac, page)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mac page [error][%w]", err)
	}
	err = ppri.loadMacInformationFromDetailHTML(&mac, page.Document)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mac page [error][%w]", err)
	}
	return &mac, nil
}

// ParseMacPage ipadに関するページをパースして、ipadに関する情報のオブジェクトを返却
func (ppri *PageParseRepositoryImpl) ParseIPadPage(page domain.Page) (*model.IPad, error) {
	var ipad model.IPad
	err := ppri.loadIPadInformationFromTitle(&ipad, page)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mac page [error][%w]", err)
	}
	err = ppri.loadIPadInformationFromDetailHTML(&ipad, page.Document)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mac page [error][%w]", err)
	}
	return &ipad, nil
}

// ParseWatchPage watchに関するページをパースして、watchに関する情報のオブジェクトを返却
func (ppri *PageParseRepositoryImpl) ParseWatchPage(page domain.Page) (*model.Watch, error) {
	var watch model.Watch
	err := ppri.loadWatchInformationFromTitle(&watch, page)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mac page [error][%w]", err)
	}
	err = ppri.loadWatchInformationFromDetailHTML(&watch, page.Document)
	if err != nil {
		return nil, fmt.Errorf("failed to parse mac page [error][%w]", err)
	}
	return &watch, nil
}

// loadMacInformationFromDetailHTML 詳細ページのHTMLから情報を取得する
func (ppri *PageParseRepositoryImpl) loadMacInformationFromDetailHTML(mac *model.Mac, doc *goquery.Document) error {
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
			mac.ReleaseDate = getReleaseYearAndMonth(year, month)
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
func (ppri *PageParseRepositoryImpl) loadMacInformationFromTitle(mac *model.Mac, page domain.Page) error {
	title := getReformattedTitle(page.Title)
	mac.Name = title

	// インチ数
	mac.Inch = getInch(title)
	// CPU
	cpuRegExp, err := regexp.Compile(`\d+\.\dGHz.+i\d`)
	if err != nil {
		return err
	}
	mac.CPU = cpuRegExp.FindString(title)
	// 色
	mac.Color = getColor(title)
	// 金額
	mac.Amount, err = getAmount(page.AmountStr)
	if err != nil {
		return err
	}
	// URL
	mac.URL = getURL(page.DetailURL)
	return nil
}

// loadIPadInformationFromDetailHTML 詳細ページのHTMLから情報を取得する
func (ppri *PageParseRepositoryImpl) loadIPadInformationFromDetailHTML(ipad *model.IPad, doc *goquery.Document) error {
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
			ipad.ReleaseDate = getReleaseYearAndMonth(year, month)
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
func (ppri *PageParseRepositoryImpl) loadIPadInformationFromTitle(ipad *model.IPad, page domain.Page) error {
	// 不要な部分を削除
	title := getReformattedTitle(page.Title)
	ipad.Name = title
	// 色
	ipad.Color = getColor(title)
	// ストレージ
	storage, err := parseStorage(title)
	if err != nil {
		return fmt.Errorf("failed to parse storage [error][%w]", err)
	}
	ipad.Storage = storage
	// 金額
	ipad.Amount, err = getAmount(page.AmountStr)
	if err != nil {
		return err
	}
	// URL
	ipad.URL = getURL(page.DetailURL)
	return nil
}

// loadWatchInformationFromDetailHTML 詳細ページのHTMLから情報を取得する
func (ppri *PageParseRepositoryImpl) loadWatchInformationFromDetailHTML(watch *model.Watch, doc *goquery.Document) error {
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
			watch.ReleaseDate = getReleaseYearAndMonth(year, month)
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
func (ppri *PageParseRepositoryImpl) loadWatchInformationFromTitle(watch *model.Watch, page domain.Page) error {
	// 不要な部分を削除
	title := getReformattedTitle(page.Title)
	watch.Name = title

	// Cellularモデルかどうか
	if strings.Contains(page.Title, "Cellular") {
		watch.IsCellular = true
	}
	// 色
	watch.Color = getColor(title)
	// 金額
	amount, err := getAmount(page.AmountStr)
	if err != nil {
		return err
	}
	watch.Amount = amount
	// URL
	watch.URL = getURL(page.DetailURL)
	return nil
}

func parseStorage(str string) (int, error) {
	var suffix string
	var coef int
	if strings.Contains(str, "GB") {
		coef = 1
		suffix = "GB"
	} else if strings.Contains(str, "TB") {
		coef = 1000
		suffix = "TB"
	} else {
		return -1, fmt.Errorf("suffix of storage string must have 'GB' or 'TB'")
	}

	re, err := regexp.Compile(`\d+` + suffix)
	if err != nil {
		return -1, err
	}
	val, err := strconv.Atoi(strings.Replace(re.FindString(str), suffix, "", 1))
	if err != nil {
		return -1, fmt.Errorf("failed to parse storage stirng to int [error][%w]", err)
	}
	return val * coef, nil
}

func getColor(str string) string {
	for _, color := range []string{"ゴールド", "シルバー", "スペースグレイ"} {
		if strings.Contains(str, color) {
			return color
		}
	}
	return ""
}

func getInch(str string) float32 {
	inchRegExp, err := regexp.Compile(`(\d|\.)+インチ`)
	if err != nil {
		return 0.0
	}
	inchStr := inchRegExp.FindString(str)
	if inchStr != "" {
		inch, err := strconv.ParseFloat(strings.Replace(inchStr, "インチ", "", 1), 32)
		if err != nil {
			return 0.0
		}
		return float32(inch)
	}
	return 0.0
}

func getReformattedTitle(str string) string {
	return strings.Replace(str, " [整備済製品]", "", 1)
}

func getAmount(str string) (int, error) {
	amountRegExp, err := regexp.Compile(`(\d|,|円（税別）|\s)+`)
	if err != nil {
		return 0, err
	}
	amountSuffixRegExp, err := regexp.Compile(`(,|円（税別）|\s)+`)
	if err != nil {
		return 0, err
	}
	amountStr := amountRegExp.FindString(str)
	amount, err := strconv.Atoi(amountSuffixRegExp.ReplaceAllLiteralString(amountStr, ""))
	if err != nil {
		return 0, err
	}
	return amount, nil
}

func getURL(str string) string {
	if !strings.HasPrefix(str, rootURL) {
		return rootURL + str
	}
	return str
}

// getReleaseYearAndMonth 商品が売られ始めた年月を表現するための時間を返却
func getReleaseYearAndMonth(year int, month int) time.Time {
	timeZone, _ := time.LoadLocation("Asia/Tokyo")
	return time.Date(year, time.Month(month), 1, 9, 0, 0, 0, timeZone)
}
