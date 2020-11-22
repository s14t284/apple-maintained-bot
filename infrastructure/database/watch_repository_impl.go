package database

import (
	"errors"
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
	"gorm.io/gorm"
)

// WatchRepositoryImpl apple watchに関する情報を操作するための実装
type WatchRepositoryImpl struct {
	SQLClient *infrastructure.SQLClient
}

// FindWatch 整備済み品apple watchの情報を検索して返す
func (wr WatchRepositoryImpl) FindWatch(watch *model.Watch) (model.Watches, error) {
	var watches model.Watches
	result := wr.SQLClient.Client.Where(watch).Order("id DESC").Find(&watches)
	if result.Error != nil {
		return nil, result.Error
	}
	return watches, nil
}

// FindWatchAll 整備済みapple watchの全ての情報を返す
func (wr WatchRepositoryImpl) FindWatchAll() (model.Watches, error) {
	var watches model.Watches
	result := wr.SQLClient.Client.Where("is_sold is false").Order("id DESC").Find(&watches)
	if result.Error != nil {
		return nil, result.Error
	}
	return watches, nil
}

// FindByURL 指定したURLに一致するapple watchを取得
func (wr WatchRepositoryImpl) FindByURL(url string) (*model.Watch, error) {
	var watch model.Watch
	result := wr.SQLClient.Client.Where("url = ?", url).Find(&watch)
	if watch.URL != url {
		return nil, result.Error
	}
	return &watch, result.Error
}

// IsExist オブジェクトがDB内に存在しているかどうか
func (wr WatchRepositoryImpl) IsExist(watch *model.Watch) (bool, uint, time.Time, error) {
	tmp := &model.Watch{}
	err := wr.SQLClient.Client.Where(&model.Watch{
		Name:        watch.Name,
		Storage:     watch.Storage,
		Color:       watch.Color,
		IsCellular:  watch.IsCellular,
		Amount:      watch.Amount,
		ReleaseDate: watch.ReleaseDate}).First(tmp).Error
	if err == nil {
		return true, tmp.ID, tmp.CreatedAt, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, 0, time.Time{}, nil
	}
	return false, 0, time.Time{}, err
}

// AddWatch 整備済み品apple watchの情報を保存する
func (wr WatchRepositoryImpl) AddWatch(watch *model.Watch) error {
	result := wr.SQLClient.Client.Create(watch)
	return result.Error
}

// UpdateWatch  整備済み品apple watchの情報を更新する
func (wr WatchRepositoryImpl) UpdateWatch(watch *model.Watch) (err error) {
	result := wr.SQLClient.Client.Save(watch)
	return result.Error
}

// UpdateAllSoldTemporary 一旦全てを売り切れ判定にする
func (wr WatchRepositoryImpl) UpdateAllSoldTemporary() error {
	result := wr.SQLClient.Client.Exec("UPDATE watches SET is_sold = true")
	return result.Error
}

// RemoveWatch 整備済み品apple watch情報を削除する
func (wr WatchRepositoryImpl) RemoveWatch(id int64) error {
	result := wr.SQLClient.Client.Delete(&model.Watch{}, id)
	return result.Error
}
