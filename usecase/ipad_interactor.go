package usecase

import (
	"fmt"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
)

// IPadInteractor ipadの情報をやりとりするためのgateway
type IPadInteractor struct {
	ipr database.IPadRepository
}

// NewIPadInteractor IPadInteractorを生成
func NewIPadInteractor(ipr database.IPadRepository) *IPadInteractor {
	if ipr == nil {
		return nil
	}
	return &IPadInteractor{ipr: ipr}
}

// FindIPad 引数に指定したipadの情報に合致するipadの一覧を取得
func (interactor *IPadInteractor) FindIPad(param *model.IPadRequestParam) (model.IPads, error) {
	ipads, err := interactor.ipr.FindIPad(param)
	return ipads, err
}

// FindIPadAll ipadの情報を取得
func (interactor *IPadInteractor) FindIPadAll() (model.IPads, error) {
	ipads, err := interactor.ipr.FindIPadAll()
	return ipads, err
}

// FindByURL 指定したURLを持つipadを取得
func (interactor *IPadInteractor) FindByURL(url string) (*model.IPad, error) {
	return interactor.ipr.FindByURL(url)
}

// IsExist 指定したオブジェクトが存在するかを取得
func (interactor *IPadInteractor) IsExist(ipad *model.IPad) (bool, uint, time.Time, error) {
	return interactor.ipr.IsExist(ipad)
}

// AddIPad ipadの情報を追加
func (interactor *IPadInteractor) AddIPad(ipad *model.IPad) (err error) {
	err = interactor.ipr.AddIPad(ipad)
	return
}

// UpdateIPad ipadの情報の更新
func (interactor *IPadInteractor) UpdateIPad(ipad *model.IPad) (err error) {
	if ipad.ID <= 0 {
		return fmt.Errorf("cannot update ipad because invalid ipad id: %d", ipad.ID)
	}
	err = interactor.ipr.UpdateIPad(ipad)
	return
}

// UpdateAllSoldTemporary 一旦全てを売り切れ判定にする
func (interactor *IPadInteractor) UpdateAllSoldTemporary() (err error) {
	err = interactor.ipr.UpdateAllSoldTemporary()
	return
}

// RemoveIPad ipadの情報を削除
func (interactor *IPadInteractor) RemoveIPad(id int64) (err error) {
	err = interactor.ipr.RemoveIPad(id)
	return
}
