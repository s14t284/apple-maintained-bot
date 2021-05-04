//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package service

import (
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
)

type IPadService interface {
	Find(param *model.IPadRequestParam) (model.IPads, error)
	FindAll() (model.IPads, error)
	FindByURL(url string) (*model.IPad, error)
	IsExist(ipad *model.IPad) (bool, uint, time.Time, error)
	Add(ipad *model.IPad) error
	Update(ipad *model.IPad) error
	UpdateAllSoldTemporary() error
	Remove(id int64) error
}
