package usecase

import (
	"fmt"
	"path"

	"github.com/s14t284/apple-maitained-bot/domain"
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
	"github.com/s14t284/apple-maitained-bot/service"
	"github.com/s14t284/apple-maitained-bot/service/parse"

	"github.com/labstack/gommon/log"
)

const shopListEndPoint = "/jp/shop/refurbished/"

// CrawlerInteractor 整備済み品のクローラー
type CrawlerInteractor struct {
	ms            service.MacService
	is            service.IPadService
	ws            service.WatchService
	parser        parse.PageParseService
	scraper       service.ScrapeService
	slackNotifier infrastructure.SlackNotifyRepository
}

var _ CrawlerUseCase = &CrawlerInteractor{}

// NewCrawlerInteractor CrawlerControllerImplを初期化
func NewCrawlerInteractor(
	ms service.MacService,
	is service.IPadService,
	ws service.WatchService,
	parser parse.PageParseService,
	scraper service.ScrapeService,
	slackNotifier infrastructure.SlackNotifyRepository,
) (*CrawlerInteractor, error) {
	if ms == nil {
		return nil, fmt.Errorf("mac parse is nil")
	}
	if is == nil {
		return nil, fmt.Errorf("ipad parse is nil")
	}
	if ws == nil {
		return nil, fmt.Errorf("watch parse is nilj")
	}
	if parser == nil {
		return nil, fmt.Errorf("page parse is nil")
	}
	if scraper == nil {
		return nil, fmt.Errorf("scraper is nil")
	}
	if slackNotifier == nil {
		return nil, fmt.Errorf("slack notifier is nil")
	}
	return &CrawlerInteractor{
		ms:            ms,
		is:            is,
		ws:            ws,
		parser:        parser,
		scraper:       scraper,
		slackNotifier: slackNotifier,
	}, nil
}

// CrawlMacPage macに関する整備済み品ページをクローリング
func (ci *CrawlerInteractor) CrawlMacPage() error {
	mu := path.Join(shopListEndPoint, "mac")
	doc, err := ci.scraper.Scrape(mu)
	if err != nil {
		log.Warnf("cannot crawl whole page. Maybe apple store is maintenance now.")
		return err
	}

	pages, err := ci.scraper.ScrapeMaintainedPage(doc)
	if err != nil {
		return fmt.Errorf("failed to crawl mac page because failed scraping [error][%w]", err)
	}

	// 一旦、全て売れていることにする
	// クローリングの際に売れ残っている判定を実施する
	err = ci.ms.UpdateAllSoldTemporary()
	if err != nil {
		return fmt.Errorf("failed to update all products to sold tempolary [error][%w]", err)
	}

	var productPage []domain.Page
	for _, page := range pages {
		// タイトルなどから情報をパース
		iF, err := ci.parser.ParsePage("mac", page)
		if err != nil {
			log.Errorf(err.Error())
		}
		mac := iF.(*model.Mac)
		// すでにDBに格納されているか確認
		isExist, id, createdAt, err := ci.ms.IsExist(mac)
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
			err = ci.ms.Update(mac)
		} else {
			err = ci.ms.Add(mac)
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
	err = ci.slackNotifier.HookToSlack(productPage, "mac")
	if err != nil {
		log.Errorf(err.Error())
	}
	return err
}

// CrawlIPadPage ipadに関する整備済み品ページをクローリング
func (ci *CrawlerInteractor) CrawlIPadPage() error {
	iu := path.Join(shopListEndPoint, "ipad")
	doc, err := ci.scraper.Scrape(iu)
	if err != nil {
		log.Warnf("cannot crawl whole page. Maybe apple store is maintenance now.")
		return err
	}

	pages, err := ci.scraper.ScrapeMaintainedPage(doc)
	if err != nil {
		return fmt.Errorf("failed to crawl ipad page because failed scraping [error][%w]", err)
	}

	// 一旦、全て売れていることにする
	// クローリングの際に売れ残っている判定を実施する
	err = ci.is.UpdateAllSoldTemporary()
	if err != nil {
		return fmt.Errorf("failed to update all products to sold tempolary [error][%w]", err)
	}

	var productPage []domain.Page
	for _, page := range pages {
		iF, err := ci.parser.ParsePage("ipad", page)
		if err != nil {
			log.Errorf(err.Error())
		}
		ipad := iF.(*model.IPad)
		// すでにDBに格納されているか確認
		isExist, id, createdAt, err := ci.is.IsExist(ipad)
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
			err = ci.is.Update(ipad)
		} else {
			err = ci.is.Add(ipad)
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
	err = ci.slackNotifier.HookToSlack(productPage, "ipad")
	if err != nil {
		log.Errorf(err.Error())
	}
	return err
}

// CrawlWatchPage watchに関する整備済み品ページをクローリング
func (ci *CrawlerInteractor) CrawlWatchPage() error {
	wu := path.Join(shopListEndPoint, "watch")
	doc, err := ci.scraper.Scrape(wu)
	if err != nil {
		log.Warnf("cannot crawl whole page. Maybe apple store is maintenance now.")
		return err
	}

	pages, err := ci.scraper.ScrapeMaintainedPage(doc)
	if err != nil {
		return fmt.Errorf("failed to crawl ipad page because failed scraping [error][%w]", err)
	}

	// 一旦、全て売れていることにする
	// クローリングの際に売れ残っている判定を実施する
	err = ci.ws.UpdateAllSoldTemporary()
	if err != nil {
		return fmt.Errorf("failed to update all products to sold tempolary [error][%w]", err)
	}

	var productPage []domain.Page
	for _, page := range pages {
		iF, err := ci.parser.ParsePage("watch", page)
		if err != nil {
			log.Errorf(err.Error())
		}
		watch := iF.(*model.Watch)
		// すでにDBに格納されているか確認
		isExist, id, createdAt, err := ci.ws.IsExist(watch)
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
			err = ci.ws.Update(watch)
		} else {
			err = ci.ws.Add(watch)
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
	err = ci.slackNotifier.HookToSlack(productPage, "apple watch")
	if err != nil {
		log.Errorf(err.Error())
	}
	return err
}
