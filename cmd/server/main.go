package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"sync"

	"github.com/labstack/gommon/log"
	"github.com/robfig/cron/v3"

	"github.com/s14t284/apple-maitained-bot/config"
	"github.com/s14t284/apple-maitained-bot/handler"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
	"github.com/s14t284/apple-maitained-bot/infrastructure/web"
	"github.com/s14t284/apple-maitained-bot/service"
	"github.com/s14t284/apple-maitained-bot/service/parse"
	"github.com/s14t284/apple-maitained-bot/usecase"
)

var activeConWg sync.WaitGroup
var numberOfActive = 0

func execCrawl(f func() error) {
	err := f()
	if err != nil {
		log.Errorf("crawling error [error][%w]", err)
	}
}

func getCronConfig(crawler usecase.CrawlerUseCase) (*cron.Cron, error) {
	c := cron.New()

	// 整備済み品収集
	_, err := c.AddFunc("CRON_TZ=Asia/Tokyo 0 8-22 * * *", func() {
		log.Info("start crawling maintained products")
		go execCrawl(crawler.CrawlMacPage)
		go execCrawl(crawler.CrawlIPadPage)
		go execCrawl(crawler.CrawlWatchPage)
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
	// pps
	ppr, err := web.NewPageParseRepositoryImpl()
	if err != nil {
		err = fmt.Errorf("failed to initialize ppr [error][%w]", err)
		log.Error(err)
		panic(err)
	}
	pps, err := parse.NewPageParseServiceImpl(ppr)
	if err != nil {
		err = fmt.Errorf("failed to initialize pps [error][%w]", err)
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
	crawler, err := usecase.NewCrawlerControllerImpl(macInteractor, ipadInteractor, watchInteractor, pps, scraper, notifier)
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
	// 一度クローリングを実行
	go execCrawl(crawler.CrawlMacPage)
	go execCrawl(crawler.CrawlIPadPage)
	go execCrawl(crawler.CrawlWatchPage)

	// http server 用の設定とserve
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // set default port
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	laddr, err := net.ResolveTCPAddr("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
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

	go func() {
		sig := <-ch
		log.Info(sig)
		err := listener.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}()

	log.Info("Run Server...")
	srv := &http.Server{Handler: nil, ConnState: func(conn net.Conn, state http.ConnState) {
		if state == http.StateActive {
			activeConWg.Add(1)
			numberOfActive++
		} else if state == http.StateIdle || state == http.StateHijacked {
			activeConWg.Done()
			numberOfActive--
		}
		log.Info("Number of active connection: " + strconv.Itoa(numberOfActive) + "\n")
	}}
	err = srv.Serve(listener)
	if err != nil {
		log.Fatal("serve error: " + err.Error())
		panic(err)
	}
	defer activeConWg.Wait()
}
