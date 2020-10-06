package parser

import "github.com/s14t284/apple-maitained-bot/domain/model"

// IParser 整備済み品ページのパーサーのインターフェース
type IParser interface {
	ParseMacPage() (*model.Mac, error)
	ParseIPadPage() (*model.IPad, error)
	ParseWatchPage() (*model.Watch, error)
}

// Parser 整備済み品ページのパーサー
type Parser struct {
	Title     string
	AmountStr string
	DetailURL string
}
