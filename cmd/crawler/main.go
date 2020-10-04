package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
)

var rootURL = "https://www.apple.com/jp/shop/refurbished/"
var products = []string{"mac", "ipad", "watch"}
var userAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1"

func main() {
	for _, product := range products {
		requestBody := url.Values{}
		req, err := http.NewRequest("GET", rootURL+product, strings.NewReader(requestBody.Encode()))
		if err != nil {
			log.Errorf(err.Error())
			panic(err)
		}
		req.Header.Add("User-Agent", userAgent)

		client := &http.Client{}
		resp, err := client.Do(req)
		if resp.StatusCode != 200 {
			log.Errorf("status code error: %d %s", resp.StatusCode, resp.StatusCode)
			panic(err)
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromResponse(resp)
		if err != nil {
			log.Errorf(err.Error())
			panic(err)
		}

		doc.Find("div .refurbished-category-grid-no-js > ul > li").Each(func(i int, s *goquery.Selection) {
			title := s.Find("h3 > a").Text()
			price := s.Find("div,.as-currentprice,.producttile-currentprice").Text()
			fmt.Printf("%s %s\n", title, price)
		})
	}
}
