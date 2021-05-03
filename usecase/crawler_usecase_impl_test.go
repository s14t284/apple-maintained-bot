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
	"github.com/s14t284/apple-maitained-bot/mock/database"
	"github.com/s14t284/apple-maitained-bot/mock/infrastructure"
	"github.com/s14t284/apple-maitained-bot/mock/parse"
	"github.com/s14t284/apple-maitained-bot/mock/web"
)

const endPoint = "/jp/shop/refurbished/"

func TestNewCrawlerControllerImpl(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := database.NewMockMacRepository(ctrl)
	ir := database.NewMockIPadRepository(ctrl)
	wr := database.NewMockWatchRepository(ctrl)
	pps := parse.NewMockPageParseService(ctrl)
	scraper := web.NewMockScraper(ctrl)
	notifier := infrastructure.NewMockSlackNotifyRepository(ctrl)
	{
		// 正常系
		cci, err := NewCrawlerControllerImpl(mr, ir, wr, pps, scraper, notifier)
		a.NotNil(cci)
		a.NoError(err)
	}
	{
		// 異常系
		// mac databaseがnil
		cci, err := NewCrawlerControllerImpl(nil, ir, wr, pps, scraper, notifier)
		a.Nil(cci)
		a.Error(err)
	}
	{
		// 異常系
		// ipad databaseがnil
		cci, err := NewCrawlerControllerImpl(mr, nil, wr, pps, scraper, notifier)
		a.Nil(cci)
		a.Error(err)
	}
	{
		// 異常系
		// watch databaseがnil
		cci, err := NewCrawlerControllerImpl(mr, ir, nil, pps, scraper, notifier)
		a.Nil(cci)
		a.Error(err)
	}
	{
		// 異常系
		// parserがnil
		cci, err := NewCrawlerControllerImpl(mr, ir, wr, nil, scraper, notifier)
		a.Nil(cci)
		a.Error(err)
	}
	{
		// 異常系
		// scraperがnil
		cci, err := NewCrawlerControllerImpl(mr, ir, wr, pps, nil, notifier)
		a.Nil(cci)
		a.Error(err)
	}
	{
		// 異常系
		// slack notifier がnil
		cci, err := NewCrawlerControllerImpl(mr, ir, wr, pps, scraper, nil)
		a.Nil(cci)
		a.Error(err)
	}
}

func TestCrawlerControllerImpl_CrawlMacPage(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mr := database.NewMockMacRepository(ctrl)
	ir := database.NewMockIPadRepository(ctrl)
	wr := database.NewMockWatchRepository(ctrl)
	pps := parse.NewMockPageParseService(ctrl)
	scraper := web.NewMockScraper(ctrl)
	notifier := infrastructure.NewMockSlackNotifyRepository(ctrl)
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
		cci, err := NewCrawlerControllerImpl(mr, ir, wr, pps, scraper, notifier)
		if err != nil {
			t.FailNow()
		}
		gomock.InOrder(
			scraper.EXPECT().Scrape(path.Join(endPoint+"mac")).Return(doc, nil),
			scraper.EXPECT().ScrapeMaintainedPage(doc).Return(pages, nil),
			mr.EXPECT().UpdateAllSoldTemporary().Return(nil),
			pps.EXPECT().ParsePage("mac", pages[0]).Return(macs[0], nil),
			mr.EXPECT().IsExist(macs[0]).Return(true, uint(0), time.Now(), nil),
			mr.EXPECT().UpdateMac(macs[0]).Return(nil),
			pps.EXPECT().ParsePage("mac", pages[1]).Return(macs[1], nil),
			mr.EXPECT().IsExist(macs[1]).Return(false, uint(1), time.Now(), nil),
			mr.EXPECT().AddMac(macs[1]).Return(nil),
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
	mr := database.NewMockMacRepository(ctrl)
	ir := database.NewMockIPadRepository(ctrl)
	wr := database.NewMockWatchRepository(ctrl)
	pps := parse.NewMockPageParseService(ctrl)
	scraper := web.NewMockScraper(ctrl)
	notifier := infrastructure.NewMockSlackNotifyRepository(ctrl)
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
		cci, err := NewCrawlerControllerImpl(mr, ir, wr, pps, scraper, notifier)
		if err != nil {
			t.FailNow()
		}
		gomock.InOrder(
			scraper.EXPECT().Scrape(path.Join(endPoint, "ipad")).Return(doc, nil),
			scraper.EXPECT().ScrapeMaintainedPage(doc).Return(pages, nil),
			ir.EXPECT().UpdateAllSoldTemporary().Return(nil),
			pps.EXPECT().ParsePage("ipad", pages[0]).Return(ipads[0], nil),
			ir.EXPECT().IsExist(ipads[0]).Return(true, uint(0), time.Now(), nil),
			ir.EXPECT().UpdateIPad(ipads[0]).Return(nil),
			pps.EXPECT().ParsePage("ipad", pages[1]).Return(ipads[1], nil),
			ir.EXPECT().IsExist(ipads[1]).Return(false, uint(1), time.Now(), nil),
			ir.EXPECT().AddIPad(ipads[1]).Return(nil),
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
	mr := database.NewMockMacRepository(ctrl)
	ir := database.NewMockIPadRepository(ctrl)
	wr := database.NewMockWatchRepository(ctrl)
	pps := parse.NewMockPageParseService(ctrl)
	scraper := web.NewMockScraper(ctrl)
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
		cci, err := NewCrawlerControllerImpl(mr, ir, wr, pps, scraper, notifier)
		if err != nil {
			t.FailNow()
		}
		gomock.InOrder(
			scraper.EXPECT().Scrape(path.Join(endPoint+"watch")).Return(doc, nil),
			scraper.EXPECT().ScrapeMaintainedPage(doc).Return(pages, nil),
			wr.EXPECT().UpdateAllSoldTemporary().Return(nil),
			pps.EXPECT().ParsePage("watch", pages[0]).Return(watches[0], nil),
			wr.EXPECT().IsExist(watches[0]).Return(true, uint(0), time.Now(), nil),
			wr.EXPECT().UpdateWatch(watches[0]).Return(nil),
			pps.EXPECT().ParsePage("watch", pages[1]).Return(watches[1], nil),
			wr.EXPECT().IsExist(watches[1]).Return(false, uint(1), time.Now(), nil),
			wr.EXPECT().AddWatch(watches[1]).Return(nil),
			notifier.EXPECT().HookToSlack(notifierPage, "apple watch").Return(nil),
		)
		err = cci.CrawlWatchPage()
		a.NoError(err)
	}
}