//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package repository

import (
	"github.com/s14t284/apple-maitained-bot/domain"
)

type PageParser interface {
	ParsePage(target string, page domain.Page) (interface{}, error)
}