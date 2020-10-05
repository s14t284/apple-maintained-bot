package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
	"github.com/s14t284/apple-maitained-bot/config"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
	"github.com/s14t284/apple-maitained-bot/parser"
	"github.com/s14t284/apple-maitained-bot/usecase"
	"github.com/s14t284/apple-maitained-bot/usecase/repository"
	"github.com/s14t284/apple-maitained-bot/utils"
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
	var macInteractor repository.MacRepository
	macInteractor = &usecase.MacInteractor{MacRepository: database.MacRepositoryImpl{SQLClient: psqlClient}}
	fmt.Println(macInteractor.FindMacAll())

	for _, product := range products {
		doc, err := utils.GetGoQueryObject(rootURL + shopListEndPoint + product)
		if err != nil {
			panic(err)
		}
		doc.Find("div .refurbished-category-grid-no-js > ul > li").Each(func(_ int, s *goquery.Selection) {
			title := s.Find("h3 > a").Text()
			amount := s.Find("div,.as-currentprice,.producttile-currentprice").Text()
			href, _ := s.Find("a").Attr("href")
			if product == products[0] {
				var macParser parser.IMacParser
				macParser = &parser.MacParser{Title: title, AmountStr: amount, DetailURL: rootURL + href}
				mac, _ := macParser.ParseMacPage()
				fmt.Println(mac)
				err := macInteractor.AddMac(mac)
				if err != nil {
					log.Errorf(err.Error())
				}
			} else if product == products[1] {
				var ipadParser parser.IIpadParser
				ipadParser = &parser.IPadParser{Title: title, AmountStr: amount, DetailURL: rootURL + href}
				ipad, _ := ipadParser.ParseIPadPage()
				fmt.Println(ipad)
			} else if product == products[2] {
				var watchParser parser.IWatchParser
				watchParser = &parser.WatchParser{Title: title, AmountStr: amount, DetailURL: rootURL + href}
				watchParser.ParseWatchPage()
				watch, _ := watchParser.ParseWatchPage()
				fmt.Println(watch)
			}
		})
	}
}
