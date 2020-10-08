package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/labstack/gommon/log"
	"github.com/robfig/cron"
	"github.com/s14t284/apple-maitained-bot/config"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
	"github.com/s14t284/apple-maitained-bot/usecase"
	"github.com/s14t284/apple-maitained-bot/usecase/repository"
	"github.com/s14t284/apple-maitained-bot/utils/crawler"
)

const rootURL = "https://www.apple.com"
const shopListEndPoint = "/jp/shop/refurbished/"

func getCronConfig(mr repository.MacRepository, ir repository.IPadRepository, wr repository.WatchRepository) *cron.Cron {
	c := cron.New()

	// Macの整備済み品収集
	c.AddFunc("CRON_TZ=Asia/Tokyo 0 8-22 * * *", func() {
		log.Info("start crawling maintained products of mac")
		crawler.CrawlMacPage(rootURL, shopListEndPoint, mr)
	})
	// IPadの整備済み品収集
	c.AddFunc("CRON_TZ=Asia/Tokyo 20 8-22 * * *", func() {
		log.Info("start crawling maintained products of ipad")
		crawler.CrawlIPadPage(rootURL, shopListEndPoint, ir)
	})
	// apple watchの整備済み品収集
	c.AddFunc("CRON_TZ=Asia/Tokyo 40 8-22 * * *", func() {
		log.Info("start crawling maintained products of apple watch")
		crawler.CrawlWatchPage(rootURL, shopListEndPoint, wr)
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
	// DB接続
	psqlClient, err := infrastructure.PsqlNewClientImpl(config.PsqlConfig)
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	var macInteractor repository.MacRepository = usecase.NewMacInteractor(psqlClient)
	var ipadInteractor repository.IPadRepository = usecase.NewIPadInteractor(psqlClient)
	var watchInteractor repository.WatchRepository = usecase.NewWatchInteractor(psqlClient)

	c := getCronConfig(macInteractor, ipadInteractor, watchInteractor)
	c.Start()

	// 仮のエンドポイント
	// TODO: 修正する
	http.HandleFunc("/mac", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		macs, err := macInteractor.FindMacAll()
		if err != nil {
			log.Errorf(err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
		json, err := json.Marshal(macs)
		if err != nil {
			log.Errorf(err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(json)
	})

	http.HandleFunc("/ipad", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ipads, err := ipadInteractor.FindIPadAll()
		if err != nil {
			log.Errorf(err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
		json, err := json.Marshal(ipads)
		if err != nil {
			log.Errorf(err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(json)
	})

	http.HandleFunc("/watch", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "GET" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		watches, err := watchInteractor.FindWatchAll()
		if err != nil {
			log.Errorf(err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
		json, err := json.Marshal(watches)
		if err != nil {
			log.Errorf(err.Error())
			w.WriteHeader(http.StatusBadRequest)
		}
		w.Write(json)
	})

	port := os.Getenv("PORT")
	if err != nil {
		port = "8080" // set default port
	}
	log.Info("Run Server...")
	err = http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("Error ListenAndServe: ", err)
	}
}
