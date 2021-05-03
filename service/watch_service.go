//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package service

import (
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
)

type WatchService interface {
	Find(param *model.WatchRequestParam) (model.Watches, error)
	FindAll() (model.Watches, error)
	FindByURL(url string) (*model.Watch, error)
	IsExist(watch *model.Watch) (bool, uint, time.Time, error)
	Add(watch *model.Watch) error
	Update(watch *model.Watch) error
	UpdateAllSoldTemporary() error
	Remove(id int64) error
}
