package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/s14t284/apple-maitained-bot/config"

	"github.com/labstack/gommon/log"

	"github.com/s14t284/apple-maitained-bot/domain"
)

// SlackNotifyRepositoryImpl slackへ通知する実装
type SlackNotifyRepositoryImpl struct {
	client     http.Client
	channel    string
	userName   string
	iconEmoji  string
	webHookURL string
}

// NewSlackNotifyRepositoryImpl slackへ通知する実装の初期化
func NewSlackNotifyRepositoryImpl(c config.SlackNotifyConfig) (*SlackNotifyRepositoryImpl, error) {
	return &SlackNotifyRepositoryImpl{
		client:     http.Client{},
		channel:    c.Channel,
		userName:   c.UserName,
		iconEmoji:  c.IconEmoji,
		webHookURL: c.WebHookURL,
	}, nil
}

// HookToSlack slackへ通知
func (s *SlackNotifyRepositoryImpl) HookToSlack(pages []domain.Page, kind string) error {
	// slackへ送るメッセージの整形
	attachments := domain.Attachments{}
	if len(pages) == 0 {
		return nil
	}

	for _, page := range pages {
		attachment := domain.Attachment{
			AuthorName: "apple",
			Color:      "good",
			Title:      page.Title,
			TitleLink:  page.DetailURL,
		}
		attachments = append(attachments, attachment)
	}
	payload := domain.Payload{
		Channel:     s.channel,
		UserName:    s.userName,
		IconEmoji:   s.iconEmoji,
		Attachments: attachments,
		Text:        kind + "の整備済み品が追加されました",
	}
	rp, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	// webhookへリクエスト
	values := url.Values{}
	values.Set("payload", string(rp))
	req, err := http.NewRequest("POST", s.webHookURL, strings.NewReader(values.Encode()))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		err = fmt.Errorf("status code error: [status_code][%d] [status][%s]", resp.StatusCode, resp.Status)
		log.Error(err)
		return err
	}
	return resp.Body.Close()
}
