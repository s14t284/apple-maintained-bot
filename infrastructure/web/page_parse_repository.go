//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package web

import (
	"github.com/s14t284/apple-maitained-bot/domain"
	"github.com/s14t284/apple-maitained-bot/domain/model"
)

type PageParseRepository interface {
	ParseMacPage(page domain.Page) (*model.Mac, error)
	ParseIPadPage(page domain.Page) (*model.IPad, error)
	ParseWatchPage(page domain.Page) (*model.Watch, error)
}
