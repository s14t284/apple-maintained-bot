//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package parse

import (
	"github.com/s14t284/apple-maitained-bot/domain"
)

// PageParseService 商品ページのパーサー
type PageParseService interface {
	ParsePage(target string, page domain.Page) (interface{}, error)
}
