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
	"github.com/s14t284/apple-maitained-bot/controller"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
	"github.com/s14t284/apple-maitained-bot/infrastructure/web"
	"github.com/s14t284/apple-maitained-bot/service"
	"github.com/s14t284/apple-maitained-bot/service/parse"
	"github.com/s14t284/apple-maitained-bot/usecase"
)

var activeConWg sync.WaitGroup
var numberOfActive = 0
var listener *net.TCPListener

type clients struct {
	srv         *http.Server
	crawlerCron *cron.Cron
	crawler     usecase.CrawlerUseCase
}

func execCrawl(f func() error) {
	err := f()
	if err != nil {
		log.Errorf("crawling error [error][%w]", err)
	}
}

// TODO: 将来的に cron ではなく、API呼び出しでcrawlingするように修正
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

func getAPIServer(gc *controller.GetController) (*http.Server, error) {
	mux := http.NewServeMux()
	// GetControllerを登録
	mux.HandleFunc("/", gc.HealthCheck)
	mux.HandleFunc("/mac", gc.GetMacHandler)
	mux.HandleFunc("/ipad", gc.GetIPadHandler)
	mux.HandleFunc("/watch", gc.GetWatchHandler)

	// http srv 用の設定とserve
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // set default port
	}
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	laddr, err := net.ResolveTCPAddr("tcp", ":"+port)
	if err != nil {
		return nil, err
	}
	listener, err = net.ListenTCP("tcp", laddr)
	if err != nil {
		return nil, err
	}

	// ctrl + c を受け取った時に終了できるように設定
	go func() {
		sig := <-ch
		log.Info(sig)
		err := listener.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}()

	srv := &http.Server{
		Addr:    port,
		Handler: mux,
		ConnState: func(conn net.Conn, state http.ConnState) {
			if state == http.StateActive {
				activeConWg.Add(1)
				numberOfActive++
			} else if state == http.StateIdle || state == http.StateHijacked {
				activeConWg.Done()
				numberOfActive--
			}
			log.Info("Number of active connection: " + strconv.Itoa(numberOfActive) + "\n")
		}}
	return srv, nil
}

func initializeClients(conf *config.Config) (*clients, error) {
	// pps
	ppr, err := web.NewPageParseRepositoryImpl()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize ppr [error][%w]", err)
	}
	pps, err := parse.NewPageParseServiceImpl(ppr)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize pps [error][%w]", err)
	}
	// scraper
	scraper, err := service.NewScrapeServiceImpl()
	if err != nil {
		return nil, fmt.Errorf("failed to initialize scraper [error][%w]", err)
	}
	// slack notifier
	notifier, err := infrastructure.NewSlackNotifyRepositoryImpl(conf.SlackNotifyConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize slack notifier [error][%w]", err)
	}
	// DB接続
	psqlClient, err := database.PsqlNewClientImpl(conf.PsqlConfig)
	if err != nil {
		return nil, err
	}
	// service, usecase
	macService := service.NewMacServiceImpl(database.MacRepositoryImpl{SQLClient: psqlClient})
	macUseCase, err := usecase.NewMacInteractor(macService)
	if err != nil {
		return nil, err
	}
	ipadService := service.NewIPadServiceImpl(database.IPadRepositoryImpl{SQLClient: psqlClient})
	ipadUseCase, err := usecase.NewIPadInteractor(ipadService)
	if err != nil {
		return nil, err
	}
	watchService := service.NewWatchServiceImpl(database.WatchRepositoryImpl{SQLClient: psqlClient})
	watchUseCase, err := usecase.NewWatchInteractor(watchService)
	if err != nil {
		return nil, err
	}
	getController, err := controller.NewGetController(macUseCase, ipadUseCase, watchUseCase)
	if err != nil {
		return nil, err
	}
	server, err := getAPIServer(getController)
	if err != nil {
		return nil, err
	}

	// crawler
	crawler, err := usecase.NewCrawlerInteractor(macService, ipadService, watchService, pps, scraper, notifier)
	if err != nil {
		return nil, err
	}
	// cronを設定
	crawlerCron, err := getCronConfig(crawler)
	if err != nil {
		return nil, err
	}

	return &clients{
		srv:         server,
		crawlerCron: crawlerCron,
		crawler:     crawler,
	}, nil
}

func main() {
	// 設定読み込み
	conf, err := config.LoadConfig()
	if err != nil {
		log.Errorf("cannot load config: %s", err.Error())
		panic(err)
	}
	clients, err := initializeClients(conf)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	log.Info("Start Crawling")
	clients.crawlerCron.Start()
	// 一度クローリングを実行
	go execCrawl(clients.crawler.CrawlMacPage)
	go execCrawl(clients.crawler.CrawlIPadPage)
	go execCrawl(clients.crawler.CrawlWatchPage)

	log.Info("Run Server...")
	err = clients.srv.Serve(listener)
	if err != nil {
		log.Fatal("serve error: " + err.Error())
	}
	defer activeConWg.Wait()
}
