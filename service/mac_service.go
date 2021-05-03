//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package service

import (
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
)

type MacService interface {
	Find(param *model.MacRequestParam) (model.Macs, error)
	FindAll() (model.Macs, error)
	FindByURL(url string) (*model.Mac, error)
	IsExist(mac *model.Mac) (bool, uint, time.Time, error)
	Add(mac *model.Mac) error
	Update(mac *model.Mac) error
	UpdateAllSoldTemporary() error
	Remove(id int64) error
}
