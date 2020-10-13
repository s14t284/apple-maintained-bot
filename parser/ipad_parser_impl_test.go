package parser

import (
	"fmt"
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/utils"
	"github.com/stretchr/testify/assert"
)

const detailHTML = `
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

func TestLoadIPadInformationFromDetailHTML(t *testing.T) {
	assert := assert.New(t)
	pageParser := Parser{Title: "title", AmountStr: "30000円", DetailURL: "https://apple.com"}
	ipad := &model.IPad{}
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(detailHTML))
	{
		pageParser.LoadIPadInformationFromDetailHTML(ipad, doc)
		fmt.Println(ipad.ReleaseDate)
		fmt.Println(utils.GetReleaseYearAndMonth(2015, 9))
		assert.Equal(ipad.ReleaseDate, utils.GetReleaseYearAndMonth(2015, 9))
		assert.Equal(ipad.Camera, "8メガピクセルiSightカメラ")
		assert.Equal(ipad.Inch, float32(7.9))
	}

}
