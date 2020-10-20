package parser

import (
	"strings"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/utils"
	"github.com/stretchr/testify/assert"
)

const detailMacHTML = `
<div class="as-productinfosection-panel Overview-panel row">
    <div class="as-productinfosection-sidepanel column large-3 small-12">
        <h3 data-autom="sectionTitle">概要</h3>
    </div>
    <div class="as-productinfosection-mainpanel column large-9 small-12">

            <div class="para-list">
            <p>
                2019年11月発売
            </p>
        </div>
        <div class="para-list">
            <p>
                Touch IDセンサーが組み込まれたTouch Bar
            </p>
        </div>
        <div class="para-list">
            <p>
                IPSテクノロジー搭載16インチ（対角）LEDバックライトディスプレイ、3,072 x 1,920ピクセル標準解像度、226ppi、数百万色以上対応
            </p>
        </div>
        <div class="para-list">
            <p>
                16GB 2,666MHz DDR4オンボードメモリ
            </p>
        </div>
        <div class="para-list">
            <p>
                512GB SSD<sup>1</sup>
            </p>
        </div>
        <div class="para-list">
            <p>
                720p FaceTime HDカメラ
            </p>
        </div>
        <div class="para-list as-pdp-lastparalist">
            <p>
                AMD Radeon Pro 5300M（4GB GDDR6メ‍モ‍リ搭載）
            </p>
        </div>
    </div>
</div>
`

func TestLoadMacInformationFromTitle(t *testing.T) {
	assert := assert.New(t)
	mac := &model.Mac{}
	{
		// 16インチMacBook Proの場合
		pageParser := Parser{
			Title:     "16インチMacBook Pro 2.4GHz 8コアIntel Core i9 Retinaディスプレイモデル - スペースグレイ [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://apple.com",
		}
		pageParser.LoadMacInformationFromTitle(mac)
		assert.Equal(mac.Inch, float32(16))
		assert.Equal(mac.CPU, "2.4GHz 8コアIntel Core i9")
		assert.Equal(mac.Color, "スペースグレイ")
		assert.Equal(mac.Amount, 30000)
		assert.Equal(mac.Name, "MacBook Pro")
		assert.Equal(mac.URL, "https://apple.com")
	}
	{
		// 15.4インチMacBook Proの場合
		// 16インチMacBook Proの場合
		pageParser := Parser{
			Title:     "15.4インチMacBook Pro 2.9GHz 6コアIntel Core i9 Retinaディスプレイモデル - シルバー [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://apple.com",
		}
		pageParser.LoadMacInformationFromTitle(mac)
		assert.Equal(mac.Inch, float32(15.4))
		assert.Equal(mac.CPU, "2.9GHz 6コアIntel Core i9")
		assert.Equal(mac.Color, "シルバー")
		assert.Equal(mac.Amount, 30000)
		assert.Equal(mac.Name, "MacBook Pro")
		assert.Equal(mac.URL, "https://apple.com")
	}
	{
		// 13.3インチMacBook Proの場合
		pageParser := Parser{
			Title:     "13.3インチMacBook Pro 1.4GHzクアッドコアIntel Core i5 Retinaディスプレイモデル - スペースグレイ [整備済製品]",
			AmountStr: "30,000円（税別）",
			DetailURL: "https://apple.com",
		}
		pageParser.LoadMacInformationFromTitle(mac)
		assert.Equal(mac.Inch, float32(13.3))
		assert.Equal(mac.CPU, "1.4GHzクアッドコアIntel Core i5")
		assert.Equal(mac.Color, "スペースグレイ")
		assert.Equal(mac.Amount, 30000)
		assert.Equal(mac.Name, "MacBook Pro")
		assert.Equal(mac.URL, "https://apple.com")
	}
	// TODO: MacBook Airに関するテストを増やす
	// TODO: MacBook以外のMacに関するテストを増やす
	// {
	// 	// Mac Proの場合
	// 	pageParser := Parser{
	// 		Title:     "Mac Pro 3.2GHz 16コア Intel Xeon W、Radeon Pro 580X [整備済製品]",
	// 		AmountStr: "30,000円（税別）",
	// 		DetailURL: "https://apple.com",
	// 	}
	// 	pageParser.LoadMacInformationFromTitle(mac)
	// 	assert.Equal(mac.Inch, float32(0.0))
	// 	assert.Equal(mac.CPU, "3.2GHz 16コア Intel Xeon W、Radeon Pro 580X")
	// 	assert.Equal(mac.Color, "")
	// 	assert.Equal(mac.Amount, 30000)
	// 	assert.Equal(mac.Name, "Mac Pro")
	// 	assert.Equal(mac.URL, "https://apple.com")
	// }

}

func TestLoadMacInformationFromDetailHTML(t *testing.T) {
	assert := assert.New(t)
	pageParser := Parser{Title: "title", AmountStr: "30000円", DetailURL: "https://apple.com"}
	mac := &model.Mac{}
	doc, _ := goquery.NewDocumentFromReader(strings.NewReader(detailMacHTML))
	{
		pageParser.LoadMacInformationFromDetailHTML(mac, doc)
		assert.Equal(utils.GetReleaseYearAndMonth(2019, 11), mac.ReleaseDate)
		assert.Equal(true, mac.TouchBar)
		assert.Equal("512GB", mac.Strage)
	}
}
