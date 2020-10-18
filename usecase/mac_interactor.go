package usecase

import (
	"fmt"
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
)

// MacInteractor macbookの情報をやりとりするためのgateway
type MacInteractor struct {
	mr database.MacRepositoryImpl
}

// NewMacInteractor MacInteractorを生成
func NewMacInteractor(dbClient *infrastructure.SQLClient) *MacInteractor {
	return &MacInteractor{mr: database.MacRepositoryImpl{SQLClient: dbClient}}
}

// FindMacAll macbookの情報を全て取得
func (interactor *MacInteractor) FindMacAll() (model.Macs, error) {
	macs, err := interactor.mr.FindMacAll()
	return macs, err
}

// FindByURL 指定したURLを持つmacを取得
func (interactor *MacInteractor) FindByURL(url string) (*model.Mac, error) {
	return interactor.mr.FindByURL(url)
}

// IsExist 指定したオブジェクトが存在するかを取得
func (interactor *MacInteractor) IsExist(mac *model.Mac) (bool, uint, time.Time, error) {
	return interactor.mr.IsExist(mac)
}

// AddMac macbookの情報を追加
func (interactor *MacInteractor) AddMac(mac *model.Mac) (err error) {
	err = interactor.mr.AddMac(mac)
	return
}

// UpdateMac macの情報の更新
func (interactor *MacInteractor) UpdateMac(mac *model.Mac) (err error) {
	if mac.ID <= 0 {
		return fmt.Errorf("cannot logical update mac because invalid mac id: %d", mac.ID)
	}
	err = interactor.mr.UpdateMac(mac)
	return
}

// UpdateAllSoldTemporary 一旦全てを売り切れ判定にする
func (interactor *MacInteractor) UpdateAllSoldTemporary() (err error) {
	err = interactor.mr.UpdateAllSoldTemporary()
	return
}

// RemoveMac macbookの情報を削除
func (interactor *MacInteractor) RemoveMac(id int64) (err error) {
	err = interactor.mr.RemoveMac(id)
	return
}
