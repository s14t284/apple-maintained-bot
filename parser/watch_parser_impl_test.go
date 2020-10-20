package parser

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/utils"
	"github.com/stretchr/testify/assert"
)

const detailWatchHTML = `
<div class="as-productinfosection-panel Overview-panel row">
    <div class="as-productinfosection-sidepanel column large-3 small-12">
        <h3 data-autom="sectionTitle">概要</h3>
    </div>
    <div class="as-productinfosection-mainpanel column large-9 small-12">

            <div class="para-list">
            <p>
                2018年9月発売モデル
            </p>
        </div>
        <div class="para-list">
            <p>
                S4（64ビットデュアルコアプロセッサ搭載。S3プロセッサと比べて最大2倍高速）
            </p>
        </div>
        <div class="para-list">
            <p>
                50メートルの耐水性能<sup>1</sup>
            </p>
        </div>
        <div class="para-list">
            <p>
                感圧タッチ対応LTPO OLED Retinaディスプレイ（1,000ニトの輝度）
            </p>
        </div>
        <div class="para-list">
            <p><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">
                Wi-Fi（802.11b / g / n 2.4GHz）
            </font></font></p>
        </div>
        <div class="para-list">
            <p><font style="vertical-align: inherit;"><font style="vertical-align: inherit;">
                Bluetooth 5.0
            </font></font></p>
        </div>
        <div class="para-list">
            <p>
                光学式心拍センサー
            </p>
        </div>
        <div class="para-list">
            <p>
                進化した加速度センサー
            </p>
        </div>
        <div class="para-list">
            <p>
                進化したジャイロスコープ
            </p>
        </div>
        <div class="para-list">
            <p>
                環境光センサー
            </p>
        </div>
        <div class="para-list">
            <p>
                容量16GB<sup>2</sup>
            </p>
        </div>
        <div class="para-list as-pdp-lastparalist">
            <p>
                全面セラミックとサファイアクリスタルの裏蓋
            </p>
        </div>
    </div>
</div>
`

func TestLoadWatchInformationFromTitle(t *testing.T) {
	assert := assert.New(t)
	watch := &model.Watch{}
	{
		// Apple Watch Series 4（GPS + Cellularモデル）の場合
		pageParser := Parser{
			Title:     "Apple Watch Series 4（GPS + Cellularモデル）- 44mmシルバーアルミニウムケースとホワイトスポーツバンド [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://apple.com",
		}
		pageParser.LoadWatchInformationFromTitle(watch)
		assert.Equal(watch.Color, "シルバーアルミニウムケースとホワイトスポーツバンド")
		assert.Equal(watch.Name, "Apple Watch Series 4")
		assert.Equal(watch.Amount, 30000)
		assert.Equal(watch.URL, "https://apple.com")
	}
	{
		// Apple Watch Series 4（GPSモデル）の場合
		// GPSとGPS + Cellularは変更できるため、取得していない
		pageParser := Parser{
			Title:     "Apple Watch Series 4（GPSモデル）- 44mmゴールドアルミニウムケースとピンクサンドスポーツバンド [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://apple.com",
		}
		pageParser.LoadWatchInformationFromTitle(watch)
		assert.Equal(watch.Color, "ゴールドアルミニウムケースとピンクサンドスポーツバンド")
		assert.Equal(watch.Name, "Apple Watch Series 4")
		assert.Equal(watch.Amount, 30000)
		assert.Equal(watch.URL, "https://apple.com")
	}
	{
		// Apple Watch Series Nike+ Series 4の場合
		pageParser := Parser{
			Title:     "Apple Watch Nike+ Series 4（GPS + Cellularモデル）- 40mmスペースグレイアルミニウムケースとアンスラサイト/ブラックNikeスポーツバンド [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://apple.com",
		}
		pageParser.LoadWatchInformationFromTitle(watch)
		assert.Equal(watch.Color, "スペースグレイアルミニウムケースとアンスラサイト/ブラックNikeスポーツバンド")
		assert.Equal(watch.Name, "Apple Watch Nike+ Series 4")
		assert.Equal(watch.Amount, 30000)
		assert.Equal(watch.URL, "https://apple.com")
	}
}

func TestLoadWatchInformationFromDetailHTML(t *testing.T) {
	assert := assert.New(t)
	pageParser := Parser{Title: "title", AmountStr: "30000円", DetailURL: "https://apple.com"}
	watch := &model.Watch{}
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(detailWatchHTML))
	{
		pageParser.LoadWatchInformationFromDetailHTML(watch, doc)
		assert.Equal(utils.GetReleaseYearAndMonth(2018, 9), watch.ReleaseDate)
		assert.Equal("16GB", watch.Strage)
	}
}
