package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/s14t284/apple-maitained-bot/utils"
)

const rootURL = "https://www.apple.com"
const shopListEndPoint = "/jp/shop/refurbished/"

var products = []string{"mac", "ipad", "watch"}

func main() {
	for _, product := range products {
		doc, err := utils.GetGoQueryObject(rootURL + shopListEndPoint + product)
		if err != nil {
			panic(err)
		}
		doc.Find("div .refurbished-category-grid-no-js > ul > li").Each(func(i int, s *goquery.Selection) {
			title := s.Find("h3 > a").Text()
			price := s.Find("div,.as-currentprice,.producttile-currentprice").Text()
			fmt.Printf("%s %s\n", title, price)
			attr, _ := s.Find("a").Attr("href")
			fmt.Println(attr)
		})
	}
}
