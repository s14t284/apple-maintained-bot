package service

import (
	"fmt"
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
)

// WatchServiceImpl apple watchの情報をやりとりするためのgateway
type WatchServiceImpl struct {
	wr database.WatchRepository
}

var _ WatchService = &WatchServiceImpl{}

// NewWatchServiceImpl WatchServiceImplを生成
func NewWatchServiceImpl(wr database.WatchRepository) *WatchServiceImpl {
	if wr == nil {
		return nil
	}
	return &WatchServiceImpl{wr: wr}
}

// Find 引数に指定したapple watchの情報に合致するapple watchの一覧を取得
func (interactor *WatchServiceImpl) Find(param *model.WatchRequestParam) (model.Watches, error) {
	watches, err := interactor.wr.FindWatch(param)
	return watches, err
}

// FindAll apple watchの情報を取得
func (interactor *WatchServiceImpl) FindAll() (model.Watches, error) {
	watches, err := interactor.wr.FindWatchAll()
	return watches, err
}

// FindByURL 指定したURLを持つapple watchを取得
func (interactor *WatchServiceImpl) FindByURL(url string) (*model.Watch, error) {
	return interactor.wr.FindByURL(url)
}

// IsExist 指定したオブジェクトが存在するかを取得
func (interactor *WatchServiceImpl) IsExist(watch *model.Watch) (bool, uint, time.Time, error) {
	return interactor.wr.IsExist(watch)
}

// Add apple watchの情報を追加
func (interactor *WatchServiceImpl) Add(watch *model.Watch) (err error) {
	err = interactor.wr.AddWatch(watch)
	return
}

// Update apple watchの情報の更新
func (interactor *WatchServiceImpl) Update(watch *model.Watch) (err error) {
	if watch.ID <= 0 {
		return fmt.Errorf("cannot update watch because invalid watch id: %d", watch.ID)
	}
	err = interactor.wr.UpdateWatch(watch)
	return
}

// UpdateAllSoldTemporary 一旦全てを売り切れ判定にする
func (interactor *WatchServiceImpl) UpdateAllSoldTemporary() (err error) {
	err = interactor.wr.UpdateAllSoldTemporary()
	return
}

// Remove apple watchの情報を削除
func (interactor *WatchServiceImpl) Remove(id int64) (err error) {
	err = interactor.wr.RemoveWatch(id)
	return
}
