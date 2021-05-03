package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/s14t284/apple-maitained-bot/config"
	"github.com/s14t284/apple-maitained-bot/handler"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
	"github.com/s14t284/apple-maitained-bot/infrastructure/web"
	"github.com/s14t284/apple-maitained-bot/service"
	"github.com/s14t284/apple-maitained-bot/usecase"

	"github.com/labstack/gommon/log"
	"github.com/robfig/cron/v3"
)

func getCronConfig(crawler usecase.CrawlerUseCase) (*cron.Cron, error) {
	c := cron.New()

	// 整備済み品収集
	_, err := c.AddFunc("CRON_TZ=Asia/Tokyo 0 8-22 * * *", func() {
		log.Info("start crawling maintained products")
		crawl := func(f func() error) {
			err := f()
			if err != nil {
				log.Errorf("crawling error [error][%w]", err)
			}
		}
		go crawl(crawler.CrawlMacPage)
		go crawl(crawler.CrawlIPadPage)
		go crawl(crawler.CrawlWatchPage)
	})

	return c, err
}

func main() {
	// 設定読み込み
	conf, err := config.LoadConfig()
	if err != nil {
		log.Errorf("cannot load config: %s", err.Error())
		panic(err)
	}
	// scraper
	scraper, err := web.NewScraperImpl()
	if err != nil {
		err = fmt.Errorf("failed to initialize scraper [error][%w]", err)
		log.Error(err)
		panic(err)
	}
	// parser
	parser, err := web.NewPageParserImpl()
	if err != nil {
		err = fmt.Errorf("failed to initialize parser [error][%w]", err)
		log.Error(err)
		panic(err)
	}
	// slack notifier
	notifier, err := infrastructure.NewSlackNotifyRepositoryImpl(conf.SlackNotifyConfig)
	if err != nil {
		err = fmt.Errorf("failed to initialize slack notifier [error][%w]", err)
		log.Error(err)
		panic(err)
	}
	// DB接続
	psqlClient, err := database.PsqlNewClientImpl(conf.PsqlConfig)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	macInteractor := service.NewMacService(database.MacRepositoryImpl{SQLClient: psqlClient})
	ipadInteractor := service.NewIPadService(database.IPadRepositoryImpl{SQLClient: psqlClient})
	watchInteractor := service.NewWatchService(database.WatchRepositoryImpl{SQLClient: psqlClient})
	// crawler
	crawler, err := usecase.NewCrawlerControllerImpl(macInteractor, ipadInteractor, watchInteractor, parser, scraper, notifier)
	if err != nil {
		log.Error(err)
	}

	// cronを設定
	c, err := getCronConfig(crawler)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	c.Start()

	// 仮のエンドポイント
	// TODO: 修正する
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		_, err := w.Write([]byte("{\"message\": \"ok\"}"))
		if err != nil {
			log.Error(err)
		}
	})

	http.HandleFunc("/mac", handler.GetMacHandler(macInteractor))

	http.HandleFunc("/ipad", handler.GetIPadHandler(ipadInteractor))

	http.HandleFunc("/watch", handler.GetWatchHandler(watchInteractor))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // set default port
	}
	go func() {
		log.Info("Run Server...")
		err = http.ListenAndServe(":"+port, nil)
		if err != nil {
			log.Fatal("Error ListenAndServe: ", err)
		}
	}()
	// 一度クローリングを実行
	err = crawler.CrawlMacPage()
	if err != nil {
		log.Error("crawling macbook page error: " + err.Error())
	}
	err = crawler.CrawlIPadPage()
	if err != nil {
		log.Error("crawling ipda page error: " + err.Error())
	}
	err = crawler.CrawlWatchPage()
	if err != nil {
		log.Error("crawling apple watch page error: " + err.Error())
	}
}
