package domain

import "github.com/PuerkitoBio/goquery"

// Page 整備済み品ページの情報
type Page struct {
	Title     string            // 商品タイトル
	AmountStr string            // 商品の金額の文字列
	DetailURL string            // 詳細ページへのURL
	Document  *goquery.Document // Webページの構造
}
