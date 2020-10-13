package parser

import (
	"testing"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/stretchr/testify/assert"
)

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
