package parser

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/utils"
	"github.com/stretchr/testify/assert"
)

const detailIPadHTML = `
<div class="as-productinfosection-panel Overview-panel row">

    <div class="as-productinfosection-sidepanel column large-3 small-12">
        <h3 data-autom="sectionTitle">概要</h3>
    </div><div class="as-productinfosection-mainpanel column large-9 small-12">

            <div class="para-list">
            <p>
                2015年9月発売モデル
            </p>
        </div>
        <div class="para-list">
            <p>
                Wi-Fi（802.11a/b/g/n/ac）
            </p>
        </div>
        <div class="para-list">
            <p>
                Bluetooth 4.2テクノロジー
            </p>
        </div>
        <div class="para-list">
            <p>
                7.9インチのRetinaディスプレイ
            </p>
        </div>
        <div class="para-list">
            <p>
                8メガピクセルiSightカメラ
            </p>
        </div>
        <div class="para-list">
            <p>
                FaceTime HDカメラ
            </p>
        </div>
        <div class="para-list">
            <p>
                1080p HDビデオ撮影
            </p>
        </div>
        <div class="para-list">
            <p>
                64ビットアーキテクチャ搭載A8チップ
            </p>
        </div>
        <div class="para-list">
            <p>
                M8モーションコプロセッサ
            </p>
        </div>
        <div class="para-list">
            <p>
                10時間のバッテリー駆動時間
            </p>
        </div>
        <div class="para-list">
            <p>
                マルチタッチスクリーン
            </p>
        </div>
        <div class="para-list as-pdp-lastparalist">
            <p>
                304 g、6.1 mm
            </p>
        </div>
    </div>
</div>
`

func TestLoadIPadInformationFromTitle(t *testing.T) {
	assert := assert.New(t)
	ipad := &model.IPad{}
	{
		// IPad Proの場合
		// インチ数がタイトルに含まれているが、インチは詳細ページから取得する
		pageParser := Parser{
			Title:     "12.9インチiPad Pro Wi-Fi + Cellular 512GB - スペースグレイ（第2世代） [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://apple.com",
		}
		pageParser.LoadIPadInformationFromTitle(ipad)
		assert.Equal(ipad.Color, "スペースグレイ")
		assert.Equal(ipad.Name, "iPad Pro")
		assert.Equal(ipad.Strage, "512GB")
		assert.Equal(ipad.Amount, 30000)
		assert.Equal(ipad.URL, "https://apple.com")

	}
	{
		// IPad Airの場合
		pageParser := Parser{
			Title:     "iPad Air Wi-Fiモデル 64GB - ゴールド [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://apple.com",
		}
		pageParser.LoadIPadInformationFromTitle(ipad)
		assert.Equal(ipad.Color, "ゴールド")
		assert.Equal(ipad.Name, "iPad Air")
		assert.Equal(ipad.Strage, "64GB")
		assert.Equal(ipad.Amount, 30000)
		assert.Equal(ipad.URL, "https://apple.com")
	}
	{
		// IPad miniの場合
		pageParser := Parser{
			Title:     "iPad mini 4 Wi-Fi 128GB - スペースグレイ [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://apple.com",
		}
		pageParser.LoadIPadInformationFromTitle(ipad)
		assert.Equal(ipad.Color, "スペースグレイ")
		assert.Equal(ipad.Name, "iPad mini 4")
		assert.Equal(ipad.Strage, "128GB")
		assert.Equal(ipad.Amount, 30000)
		assert.Equal(ipad.URL, "https://apple.com")
	}
	{
		// 通常IPadの場合
		pageParser := Parser{
			Title:     "iPad Wi-Fi 128GB - シルバー（第7世代） [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://apple.com",
		}
		pageParser.LoadIPadInformationFromTitle(ipad)
		assert.Equal(ipad.Color, "シルバー")
		assert.Equal(ipad.Name, "iPad")
		assert.Equal(ipad.Strage, "128GB")
		assert.Equal(ipad.Amount, 30000)
		assert.Equal(ipad.URL, "https://apple.com")
	}
}

func TestLoadIPadInformationFromDetailHTML(t *testing.T) {
	assert := assert.New(t)
	pageParser := Parser{Title: "title", AmountStr: "30000円", DetailURL: "https://apple.com"}
	ipad := &model.IPad{}
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(detailIPadHTML))
	{
		pageParser.LoadIPadInformationFromDetailHTML(ipad, doc)
		assert.Equal(utils.GetReleaseYearAndMonth(2015, 9), ipad.ReleaseDate)
		assert.Equal("8メガピクセルiSightカメラ", ipad.Camera)
		assert.Equal(float32(7.9), ipad.Inch)
	}

}
