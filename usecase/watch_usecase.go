package usecase

import (
	"github.com/s14t284/apple-maitained-bot/domain/model"
)

type WatchUseCase interface {
	Find(param *model.WatchRequestParam) (model.Watches, error)
	FindAll() (model.Watches, error)
	FindByURL(url string) (model.Watch, error)
	IsExist(mac *model.Watch) (bool, error)
	Add(mac *model.Watch) error
	Update(mac *model.Watch) error
	UpdateAllSoldTemporary() error
	Remove(id int64) error
}
