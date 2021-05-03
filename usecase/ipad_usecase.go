package usecase

import (
	"github.com/s14t284/apple-maitained-bot/domain/model"
)

type IPadUseCase interface {
	Find(param *model.IPadRequestParam) (model.IPads, error)
	FindAll() (model.IPads, error)
	FindByURL(url string) (model.IPad, error)
	IsExist(mac *model.IPad) (bool, error)
	Add(mac *model.IPad) error
	Update(mac *model.IPad) error
	UpdateAllSoldTemporary() error
	Remove(id int64) error
}
