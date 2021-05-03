package usecase

import (
	"fmt"
	"path"

	"github.com/s14t284/apple-maitained-bot/domain"
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
	"github.com/s14t284/apple-maitained-bot/infrastructure/web"
	parser2 "github.com/s14t284/apple-maitained-bot/service/parser"

	"github.com/labstack/gommon/log"
)

const shopListEndPoint = "/jp/shop/refurbished/"

// CrawlerUseCaseImpl 整備済み品のクローラー
type CrawlerUseCaseImpl struct {
	mr            database.MacRepository
	ir            database.IPadRepository
	wr            database.WatchRepository
	parser        parser2.PageParseService
	scraper       web.Scraper
	slackNotifier infrastructure.SlackNotifyRepository
}

// NewCrawlerControllerImpl CrawlerControllerImplを初期化
func NewCrawlerControllerImpl(
	mr database.MacRepository,
	ir database.IPadRepository,
	wr database.WatchRepository,
	parser parser2.PageParseService,
	scraper web.Scraper,
	slackNotifier infrastructure.SlackNotifyRepository,
) (*CrawlerUseCaseImpl, error) {
	if mr == nil {
		return nil, fmt.Errorf("mac parser is nil")
	}
	if ir == nil {
		return nil, fmt.Errorf("ipad parser is nil")
	}
	if wr == nil {
		return nil, fmt.Errorf("watch parser is nilj")
	}
	if parser == nil {
		return nil, fmt.Errorf("page parser is nil")
	}
	if scraper == nil {
		return nil, fmt.Errorf("scraper is nil")
	}
	if slackNotifier == nil {
		return nil, fmt.Errorf("slack notifier is nil")
	}
	return &CrawlerUseCaseImpl{
		mr:            mr,
		ir:            ir,
		wr:            wr,
		parser:        parser,
		scraper:       scraper,
		slackNotifier: slackNotifier,
	}, nil
}

// CrawlMacPage macに関する整備済み品ページをクローリング
func (c *CrawlerUseCaseImpl) CrawlMacPage() error {
	mu := path.Join(shopListEndPoint, "mac")
	doc, err := c.scraper.Scrape(mu)
	if err != nil {
		log.Warnf("cannot crawl whole page. Maybe apple store is maintenance now.")
		return err
	}

	pages, err := c.scraper.ScrapeMaintainedPage(doc)
	if err != nil {
		return fmt.Errorf("failed to crawl mac page because failed scraping [error][%w]", err)
	}

	// 一旦、全て売れていることにする
	// クローリングの際に売れ残っている判定を実施する
	err = c.mr.UpdateAllSoldTemporary()
	if err != nil {
		return fmt.Errorf("failed to update all products to sold tempolary [error][%w]", err)
	}

	var productPage []domain.Page
	for _, page := range pages {
		// タイトルなどから情報をパース
		iF, err := c.parser.ParsePage("mac", page)
		if err != nil {
			log.Errorf(err.Error())
		}
		mac := iF.(*model.Mac)
		// すでにDBに格納されているか確認
		isExist, id, createdAt, err := c.mr.IsExist(mac)
		if err != nil {
			log.Errorf(err.Error())
		}
		// 格納されている場合、まだ売れていないように戻し、URLを更新
		// 格納されていない場合、情報を追加
		if isExist {
			mac.ID = id
			mac.IsSold = false
			mac.CreatedAt = createdAt
			log.Infof("Unsold: %s", mac.URL)
			err = c.mr.UpdateMac(mac)
		} else {
			err = c.mr.AddMac(mac)
			if err == nil {
				productPage = append(productPage, domain.Page{
					Title:     page.Title,
					DetailURL: page.DetailURL,
				})
			}
		}
		if err != nil {
			log.Errorf(err.Error())
		}
	}
	err = c.slackNotifier.HookToSlack(productPage, "mac")
	if err != nil {
		log.Errorf(err.Error())
	}
	return err
}

// CrawlIPadPage ipadに関する整備済み品ページをクローリング
func (c *CrawlerUseCaseImpl) CrawlIPadPage() error {
	iu := path.Join(shopListEndPoint, "ipad")
	doc, err := c.scraper.Scrape(iu)
	if err != nil {
		log.Warnf("cannot crawl whole page. Maybe apple store is maintenance now.")
		return err
	}

	pages, err := c.scraper.ScrapeMaintainedPage(doc)
	if err != nil {
		return fmt.Errorf("failed to crawl ipad page because failed scraping [error][%w]", err)
	}

	// 一旦、全て売れていることにする
	// クローリングの際に売れ残っている判定を実施する
	err = c.ir.UpdateAllSoldTemporary()
	if err != nil {
		return fmt.Errorf("failed to update all products to sold tempolary [error][%w]", err)
	}

	var productPage []domain.Page
	for _, page := range pages {
		iF, err := c.parser.ParsePage("ipad", page)
		if err != nil {
			log.Errorf(err.Error())
		}
		ipad := iF.(*model.IPad)
		// すでにDBに格納されているか確認
		isExist, id, createdAt, err := c.ir.IsExist(ipad)
		if err != nil {
			log.Errorf(err.Error())
		}
		// 格納されている場合、まだ売れていないように戻し、URLを更新
		// 格納されていない場合、情報を追加
		if isExist {
			ipad.ID = id
			ipad.IsSold = false
			ipad.CreatedAt = createdAt
			log.Infof("Unsold: %s", ipad.URL)
			err = c.ir.UpdateIPad(ipad)
		} else {
			err = c.ir.AddIPad(ipad)
			if err == nil {
				productPage = append(productPage, domain.Page{
					Title:     page.Title,
					DetailURL: page.DetailURL,
				})
			}
		}
		if err != nil {
			log.Errorf(err.Error())
		}
	}
	err = c.slackNotifier.HookToSlack(productPage, "ipad")
	if err != nil {
		log.Errorf(err.Error())
	}
	return err
}

// CrawlWatchPage watchに関する整備済み品ページをクローリング
func (c *CrawlerUseCaseImpl) CrawlWatchPage() error {
	wu := path.Join(shopListEndPoint, "watch")
	doc, err := c.scraper.Scrape(wu)
	if err != nil {
		log.Warnf("cannot crawl whole page. Maybe apple store is maintenance now.")
		return err
	}

	pages, err := c.scraper.ScrapeMaintainedPage(doc)
	if err != nil {
		return fmt.Errorf("failed to crawl ipad page because failed scraping [error][%w]", err)
	}

	// 一旦、全て売れていることにする
	// クローリングの際に売れ残っている判定を実施する
	err = c.wr.UpdateAllSoldTemporary()
	if err != nil {
		return fmt.Errorf("failed to update all products to sold tempolary [error][%w]", err)
	}

	var productPage []domain.Page
	for _, page := range pages {
		iF, err := c.parser.ParsePage("watch", page)
		if err != nil {
			log.Errorf(err.Error())
		}
		watch := iF.(*model.Watch)
		// すでにDBに格納されているか確認
		isExist, id, createdAt, err := c.wr.IsExist(watch)
		if err != nil {
			log.Errorf(err.Error())
		}
		// 格納されている場合、まだ売れていないように戻し、URLを更新
		// 格納されていない場合、情報を追加
		if isExist {
			watch.ID = id
			watch.IsSold = false
			watch.CreatedAt = createdAt
			log.Infof("Unsold: %s", watch.URL)
			err = c.wr.UpdateWatch(watch)
		} else {
			err = c.wr.AddWatch(watch)
			if err == nil {
				productPage = append(productPage, domain.Page{
					Title:     page.Title,
					DetailURL: page.DetailURL,
				})
			}
		}
		if err != nil {
			log.Errorf(err.Error())
		}
	}
	err = c.slackNotifier.HookToSlack(productPage, "apple watch")
	if err != nil {
		log.Errorf(err.Error())
	}
	return err
}
