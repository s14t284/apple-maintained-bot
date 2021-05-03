//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package web

import (
	"github.com/s14t284/apple-maitained-bot/domain"
)

// PageParser 商品ページのパーサー
type PageParser interface {
	ParsePage(target string, page domain.Page) (interface{}, error)
}
