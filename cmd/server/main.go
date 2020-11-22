package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/s14t284/apple-maitained-bot/handler"

	"github.com/s14t284/apple-maitained-bot/interfaces"

	"github.com/labstack/gommon/log"
	"github.com/robfig/cron/v3"
	"github.com/s14t284/apple-maitained-bot/config"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
	"github.com/s14t284/apple-maitained-bot/usecase"
)

func getCronConfig(crawler interfaces.CrawlerController) *cron.Cron {
	c := cron.New()

	// Macの整備済み品収集
	c.AddFunc("CRON_TZ=Asia/Tokyo 0 8-22 * * *", func() {
		log.Info("start crawling maintained products of mac")
		crawler.CrawlMacPage()
	})
	// IPadの整備済み品収集
	c.AddFunc("CRON_TZ=Asia/Tokyo 0 8-22 * * *", func() {
		log.Info("start crawling maintained products of ipad")
		crawler.CrawlIPadPage()
	})
	// apple watchの整備済み品収集
	c.AddFunc("CRON_TZ=Asia/Tokyo 0 8-22 * * *", func() {
		log.Info("start crawling maintained products of apple watch")
		crawler.CrawlWatchPage()
	})

	return c
}

func main() {
	// 設定読み込み
	config, err := config.LoadConfig()
	if err != nil {
		log.Errorf("cannot load config: %s", err.Error())
		panic(err)
	}
	// scraper
	scraper, err := infrastructure.NewScraperImpl()
	if err != nil {
		err = fmt.Errorf("failed to initialize scraper [error][%w]", err)
		log.Error(err)
		panic(err)
	}
	// parser
	parser, err := infrastructure.NewPageParserImpl()
	if err != nil {
		err = fmt.Errorf("failed to initialize parser [error][%w]", err)
		log.Error(err)
		panic(err)
	}
	// slack notifier
	notifier, err := infrastructure.NewSlackNotifyRepositoryImpl(config.SlackNotifyConfig)
	if err != nil {
		err = fmt.Errorf("failed to initialize slack notifier [error][%w]", err)
		log.Error(err)
		panic(err)
	}
	// DB接続
	psqlClient, err := infrastructure.PsqlNewClientImpl(config.PsqlConfig)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	macInteractor := usecase.NewMacInteractor(database.MacRepositoryImpl{SQLClient: psqlClient})
	ipadInteractor := usecase.NewIPadInteractor(database.IPadRepositoryImpl{SQLClient: psqlClient})
	watchInteractor := usecase.NewWatchInteractor(database.WatchRepositoryImpl{SQLClient: psqlClient})
	// crawler
	crawler, err := interfaces.NewCrawlerControllerImpl(macInteractor, ipadInteractor, watchInteractor, parser, scraper, notifier)
	if err != nil {
		log.Error(err)
	}

	// cronを設定
	c := getCronConfig(crawler)
	c.Start()

	// 仮のエンドポイント
	// TODO: 修正する
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("{\"message\": \"ok\"}"))
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
	crawler.CrawlMacPage()
	crawler.CrawlIPadPage()
	crawler.CrawlWatchPage()
}
