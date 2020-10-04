package utils

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/labstack/gommon/log"
)

const userAgent = "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1"

// GetGoQueryObject 引数で指定したURLにアクセスしてそのURLのHTML情報を取得
func GetGoQueryObject(requestURL string) (*goquery.Document, error) {
	requestBody := url.Values{}
	req, err := http.NewRequest("GET", requestURL, strings.NewReader(requestBody.Encode()))
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	req.Header.Add("User-Agent", userAgent)

	client := &http.Client{}
	resp, err := client.Do(req)
	if resp.StatusCode != 200 {
		log.Errorf("status code error: %d %s", resp.StatusCode, resp.StatusCode)
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Errorf(err.Error())
		return nil, err
	}
	return doc, nil
}
