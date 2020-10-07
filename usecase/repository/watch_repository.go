package repository

import "github.com/s14t284/apple-maitained-bot/domain/model"

// WatchRepository 整備済み品apple watchの情報を読み書きするクライアント
type WatchRepository interface {
	FindWatchAll() (model.Watches, error)
	FindByURL(url string) (*model.Watch, error)
	AddWatch(watch *model.Watch) error
	UpdateWatch(watch *model.Watch) error
	UpdateAllSoldTemporary() error
	RemoveWatch(id int64) error
}
