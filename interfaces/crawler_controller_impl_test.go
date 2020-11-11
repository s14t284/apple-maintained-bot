package interfaces

import (
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/mock/gomock"
	"github.com/s14t284/apple-maitained-bot/domain"
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/mock/repository"
	"github.com/stretchr/testify/assert"
)

func TestNewCrawlerControllerImpl(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := repository.NewMockMacRepository(ctrl)
	ir := repository.NewMockIPadRepository(ctrl)
	wr := repository.NewMockWatchRepository(ctrl)
	parser := repository.NewMockPageParser(ctrl)
	scraper := repository.NewMockScraper(ctrl)
	notifier := repository.NewMockSlackNotifyRepository(ctrl)
	{
		// 正常系
		cci, err := NewCrawlerControllerImpl(mr, ir, wr, parser, scraper, notifier)
		assert.NotNil(cci)
		assert.NoError(err)
	}
	{
		// 異常系
		// mac repositoryがnil
		cci, err := NewCrawlerControllerImpl(nil, ir, wr, parser, scraper, notifier)
		assert.Nil(cci)
		assert.Error(err)
	}
	{
		// 異常系
		// ipad repositoryがnil
		cci, err := NewCrawlerControllerImpl(mr, nil, wr, parser, scraper, notifier)
		assert.Nil(cci)
		assert.Error(err)
	}
	{
		// 異常系
		// watch repositoryがnil
		cci, err := NewCrawlerControllerImpl(mr, ir, nil, parser, scraper, notifier)
		assert.Nil(cci)
		assert.Error(err)
	}
	{
		// 異常系
		// parserがnil
		cci, err := NewCrawlerControllerImpl(mr, ir, wr, nil, scraper, notifier)
		assert.Nil(cci)
		assert.Error(err)
	}
	{
		// 異常系
		// scraperがnil
		cci, err := NewCrawlerControllerImpl(mr, ir, wr, parser, nil, notifier)
		assert.Nil(cci)
		assert.Error(err)
	}
	{
		// 異常系
		// slack notifier がnil
		cci, err := NewCrawlerControllerImpl(mr, ir, wr, parser, scraper, nil)
		assert.Nil(cci)
		assert.Error(err)
	}
}

func TestCrawlerControllerImpl_CrawlMacPage(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := repository.NewMockMacRepository(ctrl)
	ir := repository.NewMockIPadRepository(ctrl)
	wr := repository.NewMockWatchRepository(ctrl)
	parser := repository.NewMockPageParser(ctrl)
	scraper := repository.NewMockScraper(ctrl)
	notifier := repository.NewMockSlackNotifyRepository(ctrl)
	rootURL := "https://apple.com"
	endPoint := "/products/"
	doc := &goquery.Document{}
	pages := []domain.Page{
		{Title: "MacBook PRO 15.4インチ", AmountStr: "30000円", DetailURL: "/detail", Document: doc},
		{Title: "MacBook Air", AmountStr: "15000円", DetailURL: "/detail", Document: doc},
	}
	notifierPage := []domain.Page{
		{Title: "MacBook Air", DetailURL: "/detail"},
	}
	macs := []*model.Mac{
		{Name: "MacBook PRO 15.4インチ", Amount: 30000},
		{Name: "MacBook Air", Amount: 15000},
	}
	{
		// 正常系
		cci, err := NewCrawlerControllerImpl(mr, ir, wr, parser, scraper, notifier)
		if err != nil {
			t.FailNow()
		}
		gomock.InOrder(
			scraper.EXPECT().Scrape(rootURL+endPoint+"mac").Return(doc, nil),
			scraper.EXPECT().ScrapeMaintainedPage(doc).Return(pages, nil),
			mr.EXPECT().UpdateAllSoldTemporary().Return(nil),
			parser.EXPECT().ParsePage("mac", pages[0]).Return(macs[0], nil),
			mr.EXPECT().IsExist(macs[0]).Return(true, uint(0), time.Now(), nil),
			mr.EXPECT().UpdateMac(macs[0]).Return(nil),
			parser.EXPECT().ParsePage("mac", pages[1]).Return(macs[1], nil),
			mr.EXPECT().IsExist(macs[1]).Return(false, uint(1), time.Now(), nil),
			mr.EXPECT().AddMac(macs[1]).Return(nil),
			notifier.EXPECT().HookToSlack(notifierPage, "mac").Return(nil),
		)
		err = cci.CrawlMacPage(rootURL, endPoint)
		assert.NoError(err)
	}
}

func TestCrawlerControllerImpl_CrawlIPadPage(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := repository.NewMockMacRepository(ctrl)
	ir := repository.NewMockIPadRepository(ctrl)
	wr := repository.NewMockWatchRepository(ctrl)
	parser := repository.NewMockPageParser(ctrl)
	scraper := repository.NewMockScraper(ctrl)
	notifier := repository.NewMockSlackNotifyRepository(ctrl)
	rootURL := "https://apple.com"
	endPoint := "/products/"
	doc := &goquery.Document{}
	pages := []domain.Page{
		{Title: "IPad Pro", AmountStr: "30000円", DetailURL: "/detail", Document: doc},
		{Title: "IPad Air", AmountStr: "15000円", DetailURL: "/detail", Document: doc},
	}
	notifierPage := []domain.Page{
		{Title: "IPad Air", DetailURL: "/detail"},
	}
	ipads := []*model.IPad{
		{Name: "IPad PRO", Amount: 30000},
		{Name: "IPad Air", Amount: 15000},
	}
	{
		// 正常系
		cci, err := NewCrawlerControllerImpl(mr, ir, wr, parser, scraper, notifier)
		if err != nil {
			t.FailNow()
		}
		gomock.InOrder(
			scraper.EXPECT().Scrape(rootURL+endPoint+"ipad").Return(doc, nil),
			scraper.EXPECT().ScrapeMaintainedPage(doc).Return(pages, nil),
			ir.EXPECT().UpdateAllSoldTemporary().Return(nil),
			parser.EXPECT().ParsePage("ipad", pages[0]).Return(ipads[0], nil),
			ir.EXPECT().IsExist(ipads[0]).Return(true, uint(0), time.Now(), nil),
			ir.EXPECT().UpdateIPad(ipads[0]).Return(nil),
			parser.EXPECT().ParsePage("ipad", pages[1]).Return(ipads[1], nil),
			ir.EXPECT().IsExist(ipads[1]).Return(false, uint(1), time.Now(), nil),
			ir.EXPECT().AddIPad(ipads[1]).Return(nil),
			notifier.EXPECT().HookToSlack(notifierPage, "ipad").Return(nil),
		)
		err = cci.CrawlIPadPage(rootURL, endPoint)
		assert.NoError(err)
	}
}
func TestCrawlerControllerImpl_CrawlWatchPage(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := repository.NewMockMacRepository(ctrl)
	ir := repository.NewMockIPadRepository(ctrl)
	wr := repository.NewMockWatchRepository(ctrl)
	parser := repository.NewMockPageParser(ctrl)
	scraper := repository.NewMockScraper(ctrl)
	notifier := repository.NewMockSlackNotifyRepository(ctrl)
	rootURL := "https://apple.com"
	endPoint := "/products/"
	doc := &goquery.Document{}
	pages := []domain.Page{
		{Title: "apple watch 4", AmountStr: "30000円", DetailURL: "/detail", Document: doc},
		{Title: "apple watch with Nike", AmountStr: "15000円", DetailURL: "/detail", Document: doc},
	}
	notifierPage := []domain.Page{
		{Title: "apple watch with Nike", DetailURL: "/detail"},
	}
	watches := []*model.Watch{
		{Name: "apple watch 4", Amount: 30000},
		{Name: "apple watch with Nike", Amount: 15000},
	}
	{
		// 正常系
		cci, err := NewCrawlerControllerImpl(mr, ir, wr, parser, scraper, notifier)
		if err != nil {
			t.FailNow()
		}
		gomock.InOrder(
			scraper.EXPECT().Scrape(rootURL+endPoint+"watch").Return(doc, nil),
			scraper.EXPECT().ScrapeMaintainedPage(doc).Return(pages, nil),
			wr.EXPECT().UpdateAllSoldTemporary().Return(nil),
			parser.EXPECT().ParsePage("watch", pages[0]).Return(watches[0], nil),
			wr.EXPECT().IsExist(watches[0]).Return(true, uint(0), time.Now(), nil),
			wr.EXPECT().UpdateWatch(watches[0]).Return(nil),
			parser.EXPECT().ParsePage("watch", pages[1]).Return(watches[1], nil),
			wr.EXPECT().IsExist(watches[1]).Return(false, uint(1), time.Now(), nil),
			wr.EXPECT().AddWatch(watches[1]).Return(nil),
			notifier.EXPECT().HookToSlack(notifierPage, "apple watch").Return(nil),
		)
		err = cci.CrawlWatchPage(rootURL, endPoint)
		assert.NoError(err)
	}
}
