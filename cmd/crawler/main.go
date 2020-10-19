package main

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"github.com/s14t284/apple-maitained-bot/config"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
	"github.com/s14t284/apple-maitained-bot/usecase"
	"github.com/s14t284/apple-maitained-bot/usecase/repository"
	"github.com/s14t284/apple-maitained-bot/utils/crawler"
)

const rootURL = "https://www.apple.com"
const shopListEndPoint = "/jp/shop/refurbished/"

var products = []string{"mac", "ipad", "watch"}

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}
	psqlClient, err := infrastructure.PsqlNewClientImpl(config.PsqlConfig)
	if err != nil {
		log.Errorf(err.Error())
		panic(err)
	}

	mpr := database.MacRepositoryImpl{SQLClient: psqlClient}
	ipr := database.IPadRepositoryImpl{SQLClient: psqlClient}
	wpr := database.WatchRepositoryImpl{SQLClient: psqlClient}
	var macInteractor repository.MacRepository = usecase.NewMacInteractor(mpr)
	var ipadInteractor repository.IPadRepository = usecase.NewIPadInteractor(ipr)
	var watchInteractor repository.WatchRepository = usecase.NewWatchInteractor(wpr)

	for _, product := range products {
		switch product {
		case products[0]:
			crawler.CrawlMacPage(rootURL, shopListEndPoint, macInteractor)
		case products[1]:
			crawler.CrawlIPadPage(rootURL, shopListEndPoint, ipadInteractor)
		case products[2]:
			crawler.CrawlWatchPage(rootURL, shopListEndPoint, watchInteractor)
		default:
			panic(fmt.Errorf("invalid path parameter: %s", product))
		}
	}
}
