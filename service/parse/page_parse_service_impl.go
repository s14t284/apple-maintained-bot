package parse

import (
	"fmt"

	"github.com/s14t284/apple-maitained-bot/domain"
	"github.com/s14t284/apple-maitained-bot/infrastructure/web"
)

// PageParseServiceImpl 商品ページのパーサーの実装
type PageParseServiceImpl struct {
	ppr web.PageParseRepository
}

var _ PageParseService = &PageParseServiceImpl{}

// NewPageParseServiceImpl PageParseServiceImplを初期化
func NewPageParseServiceImpl(ppr web.PageParseRepository) (*PageParseServiceImpl, error) {
	if ppr == nil {
		return nil, fmt.Errorf("page parse repository must not be nil")
	}
	return &PageParseServiceImpl{
		ppr: ppr,
	}, nil
}

// ParsePage 商品ページのパース
func (ppi *PageParseServiceImpl) ParsePage(target string, page domain.Page) (interface{}, error) {
	switch target {
	case "mac":
		return ppi.ppr.ParseMacPage(page)
	case "ipad":
		return ppi.ppr.ParseIPadPage(page)
	case "watch":
		return ppi.ppr.ParseWatchPage(page)
	default:
		return nil, fmt.Errorf("target must be `mac`, `ipad`, or `watch`")
	}
}
