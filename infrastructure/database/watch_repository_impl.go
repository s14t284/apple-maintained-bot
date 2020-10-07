package database

import (
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
)

// WatchRepositoryImpl apple watchに関する情報を操作するための実装
type WatchRepositoryImpl struct {
	SQLClient *infrastructure.SQLClient
}

// FindWatchAll 整備済みapple watchの全ての情報を返す
func (watchRepository *WatchRepositoryImpl) FindWatchAll() (model.Watches, error) {
	var watches model.Watches
	result := watchRepository.SQLClient.Client.Where("is_sold is false").Order("id DESC").Find(&watches)
	if result.Error != nil {
		return nil, result.Error
	}
	return watches, nil
}

// FindByURL 指定したURLに一致するapple watchを取得
func (watchRepository *WatchRepositoryImpl) FindByURL(url string) (*model.Watch, error) {
	var watch model.Watch
	result := watchRepository.SQLClient.Client.Where("url = ?", url).Find(&watch)
	if watch.URL != url {
		return nil, result.Error
	}
	return &watch, result.Error
}

// AddWatch 整備済み品apple watchの情報を保存する
func (watchRepository *WatchRepositoryImpl) AddWatch(watch *model.Watch) error {
	result := watchRepository.SQLClient.Client.Create(watch)
	return result.Error
}

// UpdateWatch  整備済み品apple watchの情報を更新する
func (watchRepository *WatchRepositoryImpl) UpdateWatch(watch *model.Watch) (err error) {
	result := watchRepository.SQLClient.Client.Save(watch)
	return result.Error
}

// UpdateAllSoldTemporary 一旦全てを売り切れ判定にする
func (watchRepository *WatchRepositoryImpl) UpdateAllSoldTemporary() error {
	result := watchRepository.SQLClient.Client.Exec("UPDATE watches SET is_sold = true")
	return result.Error
}

// RemoveWatch 整備済み品apple watch情報を削除する
func (watchRepository *WatchRepositoryImpl) RemoveWatch(id int64) error {
	result := watchRepository.SQLClient.Client.Delete(&model.Watch{}, id)
	return result.Error
}
