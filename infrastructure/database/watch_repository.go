//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../../mock/$GOPACKAGE/$GOFILE

package database

import (
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
)

// WatchRepository 整備済み品apple watchの情報を読み書きするクライアント
type WatchRepository interface {
	FindWatch(param *model.WatchRequestParam) (model.Watches, error)
	FindWatchAll() (model.Watches, error)
	FindByURL(url string) (*model.Watch, error)
	IsExist(watch *model.Watch) (bool, uint, time.Time, error)
	AddWatch(watch *model.Watch) error
	UpdateWatch(watch *model.Watch) error
	UpdateAllSoldTemporary() error
	RemoveWatch(id int64) error
}
