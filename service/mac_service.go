package service

import (
	"fmt"
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
)

// MacService macbookの情報をやりとりするためのgateway
type MacService struct {
	mr database.MacRepository
}

// NewMacService MacInteractorを生成
func NewMacService(mr database.MacRepository) *MacService {
	if mr == nil {
		return nil
	}
	return &MacService{mr: mr}
}

// FindMac 引数に指定したmacの情報に合致するmacの一覧を取得
func (interactor *MacService) FindMac(param *model.MacRequestParam) (model.Macs, error) {
	macs, err := interactor.mr.FindMac(param)
	return macs, err
}

// FindMacAll macの情報を全て取得
func (interactor *MacService) FindMacAll() (model.Macs, error) {
	macs, err := interactor.mr.FindMacAll()
	return macs, err
}

// FindByURL 指定したURLを持つmacを取得
func (interactor *MacService) FindByURL(url string) (*model.Mac, error) {
	return interactor.mr.FindByURL(url)
}

// IsExist 指定したオブジェクトが存在するかを取得
func (interactor *MacService) IsExist(mac *model.Mac) (bool, uint, time.Time, error) {
	return interactor.mr.IsExist(mac)
}

// AddMac macの情報を追加
func (interactor *MacService) AddMac(mac *model.Mac) (err error) {
	err = interactor.mr.AddMac(mac)
	return
}

// UpdateMac macの情報の更新
func (interactor *MacService) UpdateMac(mac *model.Mac) (err error) {
	if mac.ID <= 0 {
		return fmt.Errorf("cannot update mac because invalid mac id: %d", mac.ID)
	}
	err = interactor.mr.UpdateMac(mac)
	return
}

// UpdateAllSoldTemporary 一旦全てを売り切れ判定にする
func (interactor *MacService) UpdateAllSoldTemporary() (err error) {
	err = interactor.mr.UpdateAllSoldTemporary()
	return
}

// RemoveMac macの情報を削除
func (interactor *MacService) RemoveMac(id int64) (err error) {
	err = interactor.mr.RemoveMac(id)
	return
}
