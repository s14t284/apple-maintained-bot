//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package repository

import (
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
)

// IPadRepository 整備済み品ipadの情報を読み書きするクライアント
type IPadRepository interface {
	FindIPad(*model.IPad) (model.IPads, error)
	FindIPadAll() (model.IPads, error)
	FindByURL(url string) (*model.IPad, error)
	IsExist(ipad *model.IPad) (bool, uint, time.Time, error)
	AddIPad(ipad *model.IPad) error
	UpdateIPad(ipad *model.IPad) error
	UpdateAllSoldTemporary() error
	RemoveIPad(id int64) error
}
