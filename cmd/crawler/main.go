package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
	"github.com/s14t284/apple-maitained-bot/config"
	"github.com/s14t284/apple-maitained-bot/domain/model"
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

func searchMacFromExists(newMac *model.Mac, macInteractor repository.MacRepository) (flag bool, err error) {
	mac, _ := macInteractor.FindByURL(newMac.URL)
	if mac != nil {
		mac.IsSold = false
		log.Infof("Unsold: %s", mac.URL)
		err = macInteractor.UpdateMac(mac)
		flag = true
		return
	}
	return
}

func searchIPadFromExists(newIPad *model.IPad, ipadInteractor repository.IPadRepository) (flag bool, err error) {
	ipad, _ := ipadInteractor.FindByURL(newIPad.URL)
	if ipad != nil {
		ipad.IsSold = false
		log.Infof("Unsold: %s", ipad.URL)
		err = ipadInteractor.UpdateIPad(ipad)
		flag = true
		return
	}
	return
}

func searchWatchFromExists(newWatch *model.Watch, watchInteractor repository.WatchRepository) (flag bool, err error) {
	watch, _ := watchInteractor.FindByURL(newWatch.URL)
	if watch != nil {
		watch.IsSold = false
		log.Infof("Unsold: %s", watch.URL)
		err = watchInteractor.UpdateWatch(watch)
		flag = true
		return
	}
	return
}

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
	var ipadInteractor repository.IPadRepository
	ipadInteractor = &usecase.IPadInteractor{IPadRepository: database.IPadRepositoryImpl{SQLClient: psqlClient}}
	var watchInteractor repository.WatchRepository
	watchInteractor = &usecase.WatchInteractor{WatchRepository: database.WatchRepositoryImpl{SQLClient: psqlClient}}

	for _, product := range products {
		doc, err := utils.GetGoQueryObject(rootURL + shopListEndPoint + product)
		if err != nil {
			panic(err)
		}

		// 一旦、全て売れていることにする
		// スクレイピングの際に売れ残っている判定を実施する
		switch product {
		case products[0]:
			macInteractor.UpdateAllSoldTemporary()
		case products[1]:
			ipadInteractor.UpdateAllSoldTemporary()
		case products[2]:
			watchInteractor.UpdateAllSoldTemporary()
		default:
			panic(fmt.Errorf("invalid path parameter: %s", product))
		}
		doc.Find("div .refurbished-category-grid-no-js > ul > li").Each(func(_ int, s *goquery.Selection) {
			title := s.Find("h3 > a").Text()
			amount := s.Find("div,.as-currentprice,.producttile-currentprice").Text()
			href, _ := s.Find("a").Attr("href")
			var pageParser parser.IParser
			pageParser = &parser.Parser{Title: title, AmountStr: amount, DetailURL: rootURL + href}
			switch product {
			case products[0]:
				mac, _ := pageParser.ParseMacPage()
				flag, err := searchMacFromExists(mac, macInteractor)
				if err != nil {
					log.Errorf(err.Error())
				}
				if !flag {
					err := macInteractor.AddMac(mac)
					if err != nil {
						log.Errorf(err.Error())
					}
				}
			case products[1]:
				ipad, _ := pageParser.ParseIPadPage()
				flag, err := searchIPadFromExists(ipad, ipadInteractor)
				if err != nil {
					log.Errorf(err.Error())
				}
				if !flag {
					err := ipadInteractor.AddIPad(ipad)
					if err != nil {
						log.Errorf(err.Error())
					}
				}
			case products[2]:
				watch, _ := pageParser.ParseWatchPage()
				flag, err := searchWatchFromExists(watch, watchInteractor)
				if err != nil {
					log.Errorf(err.Error())
				}
				if !flag {
					err := watchInteractor.AddWatch(watch)
					if err != nil {
						log.Errorf(err.Error())
					}
				}
			default:
				err = fmt.Errorf("invalid product parameter: %s", product)
				panic(err)
			}
		})
	}
}
