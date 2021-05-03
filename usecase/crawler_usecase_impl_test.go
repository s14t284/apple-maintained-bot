package usecase

import (
	"path"
	"testing"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/s14t284/apple-maitained-bot/domain"
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/mock/infrastructure"
	"github.com/s14t284/apple-maitained-bot/mock/parse"
	"github.com/s14t284/apple-maitained-bot/mock/service"
)

const endPoint = "/jp/shop/refurbished/"

func TestNewCrawlerControllerImpl(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ms := service.NewMockMacService(ctrl)
	is := service.NewMockIPadService(ctrl)
	ws := service.NewMockWatchService(ctrl)
	pps := parse.NewMockPageParseService(ctrl)
	scraper := service.NewMockScrapeService(ctrl)
	notifier := infrastructure.NewMockSlackNotifyRepository(ctrl)
	{
		// 正常系
		cci, err := NewCrawlerUseCaseImpl(ms, is, ws, pps, scraper, notifier)
		a.NotNil(cci)
		a.NoError(err)
	}
	{
		// 異常系
		// mac databaseがnil
		cci, err := NewCrawlerUseCaseImpl(nil, is, ws, pps, scraper, notifier)
		a.Nil(cci)
		a.Error(err)
	}
	{
		// 異常系
		// ipad databaseがnil
		cci, err := NewCrawlerUseCaseImpl(ms, nil, ws, pps, scraper, notifier)
		a.Nil(cci)
		a.Error(err)
	}
	{
		// 異常系
		// watch databaseがnil
		cci, err := NewCrawlerUseCaseImpl(ms, is, nil, pps, scraper, notifier)
		a.Nil(cci)
		a.Error(err)
	}
	{
		// 異常系
		// parserがnil
		cci, err := NewCrawlerUseCaseImpl(ms, is, ws, nil, scraper, notifier)
		a.Nil(cci)
		a.Error(err)
	}
	{
		// 異常系
		// scraperがnil
		cci, err := NewCrawlerUseCaseImpl(ms, is, ws, pps, nil, notifier)
		a.Nil(cci)
		a.Error(err)
	}
	{
		// 異常系
		// slack notifier がnil
		cci, err := NewCrawlerUseCaseImpl(ms, is, ws, pps, scraper, nil)
		a.Nil(cci)
		a.Error(err)
	}
}

func TestCrawlerControllerImpl_CrawlMacPage(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ms := service.NewMockMacService(ctrl)
	is := service.NewMockIPadService(ctrl)
	ws := service.NewMockWatchService(ctrl)
	pps := parse.NewMockPageParseService(ctrl)
	scraper := service.NewMockScrapeService(ctrl)
	notifier := infrastructure.NewMockSlackNotifyRepository(ctrl)
	doc := &goquery.Document{}
	pages := []domain.Page{
		{Title: "MacBook PRO 15.4インチ", AmountStr: "30000円", DetailURL: "/detail", Document: doc},
		{Title: "MacBook Ais", AmountStr: "15000円", DetailURL: "/detail", Document: doc},
	}
	notifierPage := []domain.Page{
		{Title: "MacBook Ais", DetailURL: "/detail"},
	}
	macs := []*model.Mac{
		{Name: "MacBook PRO 15.4インチ", Amount: 30000},
		{Name: "MacBook Ais", Amount: 15000},
	}
	{
		// 正常系
		cci, err := NewCrawlerUseCaseImpl(ms, is, ws, pps, scraper, notifier)
		if err != nil {
			t.FailNow()
		}
		gomock.InOrder(
			scraper.EXPECT().Scrape(path.Join(endPoint+"mac")).Return(doc, nil),
			scraper.EXPECT().ScrapeMaintainedPage(doc).Return(pages, nil),
			ms.EXPECT().UpdateAllSoldTemporary().Return(nil),
			pps.EXPECT().ParsePage("mac", pages[0]).Return(macs[0], nil),
			ms.EXPECT().IsExist(macs[0]).Return(true, uint(0), time.Now(), nil),
			ms.EXPECT().Update(macs[0]).Return(nil),
			pps.EXPECT().ParsePage("mac", pages[1]).Return(macs[1], nil),
			ms.EXPECT().IsExist(macs[1]).Return(false, uint(1), time.Now(), nil),
			ms.EXPECT().Add(macs[1]).Return(nil),
			notifier.EXPECT().HookToSlack(notifierPage, "mac").Return(nil),
		)
		err = cci.CrawlMacPage()
		a.NoError(err)
	}
}

func TestCrawlerControllerImpl_CrawlIPadPage(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ms := service.NewMockMacService(ctrl)
	is := service.NewMockIPadService(ctrl)
	ws := service.NewMockWatchService(ctrl)
	pps := parse.NewMockPageParseService(ctrl)
	scraper := service.NewMockScrapeService(ctrl)
	notifier := infrastructure.NewMockSlackNotifyRepository(ctrl)
	doc := &goquery.Document{}
	pages := []domain.Page{
		{Title: "IPad Pro", AmountStr: "30000円", DetailURL: "/detail", Document: doc},
		{Title: "IPad Ais", AmountStr: "15000円", DetailURL: "/detail", Document: doc},
	}
	notifierPage := []domain.Page{
		{Title: "IPad Ais", DetailURL: "/detail"},
	}
	ipads := []*model.IPad{
		{Name: "IPad PRO", Amount: 30000},
		{Name: "IPad Ais", Amount: 15000},
	}
	{
		// 正常系
		cci, err := NewCrawlerUseCaseImpl(ms, is, ws, pps, scraper, notifier)
		if err != nil {
			t.FailNow()
		}
		gomock.InOrder(
			scraper.EXPECT().Scrape(path.Join(endPoint, "ipad")).Return(doc, nil),
			scraper.EXPECT().ScrapeMaintainedPage(doc).Return(pages, nil),
			is.EXPECT().UpdateAllSoldTemporary().Return(nil),
			pps.EXPECT().ParsePage("ipad", pages[0]).Return(ipads[0], nil),
			is.EXPECT().IsExist(ipads[0]).Return(true, uint(0), time.Now(), nil),
			is.EXPECT().Update(ipads[0]).Return(nil),
			pps.EXPECT().ParsePage("ipad", pages[1]).Return(ipads[1], nil),
			is.EXPECT().IsExist(ipads[1]).Return(false, uint(1), time.Now(), nil),
			is.EXPECT().Add(ipads[1]).Return(nil),
			notifier.EXPECT().HookToSlack(notifierPage, "ipad").Return(nil),
		)
		err = cci.CrawlIPadPage()
		a.NoError(err)
	}
}
func TestCrawlerControllerImpl_CrawlWatchPage(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	ms := service.NewMockMacService(ctrl)
	is := service.NewMockIPadService(ctrl)
	ws := service.NewMockWatchService(ctrl)
	pps := parse.NewMockPageParseService(ctrl)
	scraper := service.NewMockScrapeService(ctrl)
	notifier := infrastructure.NewMockSlackNotifyRepository(ctrl)
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
		cci, err := NewCrawlerUseCaseImpl(ms, is, ws, pps, scraper, notifier)
		if err != nil {
			t.FailNow()
		}
		gomock.InOrder(
			scraper.EXPECT().Scrape(path.Join(endPoint+"watch")).Return(doc, nil),
			scraper.EXPECT().ScrapeMaintainedPage(doc).Return(pages, nil),
			ws.EXPECT().UpdateAllSoldTemporary().Return(nil),
			pps.EXPECT().ParsePage("watch", pages[0]).Return(watches[0], nil),
			ws.EXPECT().IsExist(watches[0]).Return(true, uint(0), time.Now(), nil),
			ws.EXPECT().Update(watches[0]).Return(nil),
			pps.EXPECT().ParsePage("watch", pages[1]).Return(watches[1], nil),
			ws.EXPECT().IsExist(watches[1]).Return(false, uint(1), time.Now(), nil),
			ws.EXPECT().Add(watches[1]).Return(nil),
			notifier.EXPECT().HookToSlack(notifierPage, "apple watch").Return(nil),
		)
		err = cci.CrawlWatchPage()
		a.NoError(err)
	}
}
