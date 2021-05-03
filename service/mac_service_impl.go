package service

import (
	"fmt"
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
)

// MacServiceImpl macbookの情報をやりとりするためのgateway
type MacServiceImpl struct {
	mr database.MacRepository
}

var _ MacService = &MacServiceImpl{}

// NewMacServiceImpl MacServiceImplを生成
func NewMacServiceImpl(mr database.MacRepository) *MacServiceImpl {
	if mr == nil {
		return nil
	}
	return &MacServiceImpl{mr: mr}
}

// Find 引数に指定したmacの情報に合致するmacの一覧を取得
func (interactor *MacServiceImpl) Find(param *model.MacRequestParam) (model.Macs, error) {
	macs, err := interactor.mr.FindMac(param)
	return macs, err
}

// FindAll macの情報を全て取得
func (interactor *MacServiceImpl) FindAll() (model.Macs, error) {
	macs, err := interactor.mr.FindMacAll()
	return macs, err
}

// FindByURL 指定したURLを持つmacを取得
func (interactor *MacServiceImpl) FindByURL(url string) (*model.Mac, error) {
	return interactor.mr.FindByURL(url)
}

// IsExist 指定したオブジェクトが存在するかを取得
func (interactor *MacServiceImpl) IsExist(mac *model.Mac) (bool, uint, time.Time, error) {
	return interactor.mr.IsExist(mac)
}

// Add macの情報を追加
func (interactor *MacServiceImpl) Add(mac *model.Mac) (err error) {
	err = interactor.mr.AddMac(mac)
	return
}

// Update macの情報の更新
func (interactor *MacServiceImpl) Update(mac *model.Mac) (err error) {
	if mac.ID <= 0 {
		return fmt.Errorf("cannot update mac because invalid mac id: %d", mac.ID)
	}
	err = interactor.mr.UpdateMac(mac)
	return
}

// UpdateAllSoldTemporary 一旦全てを売り切れ判定にする
func (interactor *MacServiceImpl) UpdateAllSoldTemporary() (err error) {
	err = interactor.mr.UpdateAllSoldTemporary()
	return
}

// Remove macの情報を削除
func (interactor *MacServiceImpl) Remove(id int64) (err error) {
	err = interactor.mr.RemoveMac(id)
	return
}
