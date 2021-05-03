package usecase

import (
	"github.com/s14t284/apple-maitained-bot/domain/model"
)

type MacUseCase interface {
	Find(param *model.MacRequestParam) (model.Macs, error)
	FindAll() (model.Macs, error)
	FindByURL(url string) (model.Mac, error)
	IsExist(mac *model.Mac) (bool, error)
	Add(mac *model.Mac) error
	Update(mac *model.Mac) error
	UpdateAllSoldTemporary() error
	Remove(id int64) error
}
