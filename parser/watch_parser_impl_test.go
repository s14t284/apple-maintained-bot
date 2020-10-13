package parser

import (
	"testing"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/stretchr/testify/assert"
)

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
