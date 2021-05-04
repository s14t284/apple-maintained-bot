package service

import (
	"fmt"
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
)

// IPadServiceImpl ipadの情報をやりとりするためのgateway
type IPadServiceImpl struct {
	ipr database.IPadRepository
}

var _ IPadService = &IPadServiceImpl{}

// NewIPadServiceImpl IPadServiceImplを生成
func NewIPadServiceImpl(ipr database.IPadRepository) *IPadServiceImpl {
	if ipr == nil {
		return nil
	}
	return &IPadServiceImpl{ipr: ipr}
}

// Find 引数に指定したipadの情報に合致するipadの一覧を取得
func (interactor *IPadServiceImpl) Find(param *model.IPadRequestParam) (model.IPads, error) {
	ipads, err := interactor.ipr.FindIPad(param)
	return ipads, err
}

// FindAll ipadの情報を取得
func (interactor *IPadServiceImpl) FindAll() (model.IPads, error) {
	ipads, err := interactor.ipr.FindIPadAll()
	return ipads, err
}

// FindByURL 指定したURLを持つipadを取得
func (interactor *IPadServiceImpl) FindByURL(url string) (*model.IPad, error) {
	return interactor.ipr.FindByURL(url)
}

// IsExist 指定したオブジェクトが存在するかを取得
func (interactor *IPadServiceImpl) IsExist(ipad *model.IPad) (bool, uint, time.Time, error) {
	return interactor.ipr.IsExist(ipad)
}

// Add ipadの情報を追加
func (interactor *IPadServiceImpl) Add(ipad *model.IPad) (err error) {
	err = interactor.ipr.AddIPad(ipad)
	return
}

// Update ipadの情報の更新
func (interactor *IPadServiceImpl) Update(ipad *model.IPad) (err error) {
	if ipad.ID <= 0 {
		return fmt.Errorf("cannot update ipad because invalid ipad id: %d", ipad.ID)
	}
	err = interactor.ipr.UpdateIPad(ipad)
	return
}

// UpdateAllSoldTemporary 一旦全てを売り切れ判定にする
func (interactor *IPadServiceImpl) UpdateAllSoldTemporary() (err error) {
	err = interactor.ipr.UpdateAllSoldTemporary()
	return
}

// Remove ipadの情報を削除
func (interactor *IPadServiceImpl) Remove(id int64) (err error) {
	err = interactor.ipr.RemoveIPad(id)
	return
}
