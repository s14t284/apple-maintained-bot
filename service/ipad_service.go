package service

import (
	"fmt"
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
)

// IPadService ipadの情報をやりとりするためのgateway
type IPadService struct {
	ipr database.IPadRepository
}

// NewIPadService IPadInteractorを生成
func NewIPadService(ipr database.IPadRepository) *IPadService {
	if ipr == nil {
		return nil
	}
	return &IPadService{ipr: ipr}
}

// FindIPad 引数に指定したipadの情報に合致するipadの一覧を取得
func (interactor *IPadService) FindIPad(param *model.IPadRequestParam) (model.IPads, error) {
	ipads, err := interactor.ipr.FindIPad(param)
	return ipads, err
}

// FindIPadAll ipadの情報を取得
func (interactor *IPadService) FindIPadAll() (model.IPads, error) {
	ipads, err := interactor.ipr.FindIPadAll()
	return ipads, err
}

// FindByURL 指定したURLを持つipadを取得
func (interactor *IPadService) FindByURL(url string) (*model.IPad, error) {
	return interactor.ipr.FindByURL(url)
}

// IsExist 指定したオブジェクトが存在するかを取得
func (interactor *IPadService) IsExist(ipad *model.IPad) (bool, uint, time.Time, error) {
	return interactor.ipr.IsExist(ipad)
}

// AddIPad ipadの情報を追加
func (interactor *IPadService) AddIPad(ipad *model.IPad) (err error) {
	err = interactor.ipr.AddIPad(ipad)
	return
}

// UpdateIPad ipadの情報の更新
func (interactor *IPadService) UpdateIPad(ipad *model.IPad) (err error) {
	if ipad.ID <= 0 {
		return fmt.Errorf("cannot update ipad because invalid ipad id: %d", ipad.ID)
	}
	err = interactor.ipr.UpdateIPad(ipad)
	return
}

// UpdateAllSoldTemporary 一旦全てを売り切れ判定にする
func (interactor *IPadService) UpdateAllSoldTemporary() (err error) {
	err = interactor.ipr.UpdateAllSoldTemporary()
	return
}

// RemoveIPad ipadの情報を削除
func (interactor *IPadService) RemoveIPad(id int64) (err error) {
	err = interactor.ipr.RemoveIPad(id)
	return
}
