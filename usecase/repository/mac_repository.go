//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package repository

import (
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
)

// MacRepository 整備済み品macの情報を読み書きするクライアント
type MacRepository interface {
	FindMac(*model.Mac) (model.Macs, error)
	FindMacAll() (model.Macs, error)
	FindByURL(url string) (*model.Mac, error)
	IsExist(mac *model.Mac) (bool, uint, time.Time, error)
	AddMac(mac *model.Mac) error
	UpdateMac(mac *model.Mac) error
	UpdateAllSoldTemporary() error
	RemoveMac(id int64) error
}
