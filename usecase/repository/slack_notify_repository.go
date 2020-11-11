//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package repository

import "github.com/s14t284/apple-maitained-bot/domain"

// SlackNotifyRepository slackへ通知するリポジトリ
type SlackNotifyRepository interface {
	HookToSlack(pages []domain.Page, kind string) error
}
