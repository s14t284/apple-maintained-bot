package usecase

import (
	"fmt"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
)

// IPadInteractor ipadの情報をやりとりするためのgateway
type IPadInteractor struct {
	IPadRepository database.IPadRepositoryImpl
}

// NewIPadInteractor IPadInteractorを生成
func NewIPadInteractor(dbClient *infrastructure.SQLClient) *IPadInteractor {
	return &IPadInteractor{IPadRepository: database.IPadRepositoryImpl{SQLClient: dbClient}}
}

// FindIPadAll ipadの情報を取得
func (interactor *IPadInteractor) FindIPadAll() (model.IPads, error) {
	ipads, err := interactor.IPadRepository.FindIPadAll()
	return ipads, err
}

// FindByURL 指定したURLを持つipadを取得
func (interactor *IPadInteractor) FindByURL(url string) (*model.IPad, error) {
	return interactor.IPadRepository.FindByURL(url)
}

// AddIPad ipadの情報を追加
func (interactor *IPadInteractor) AddIPad(ipad *model.IPad) (err error) {
	err = interactor.IPadRepository.AddIPad(ipad)
	return
}

// UpdateIPad ipadの情報の更新
func (interactor *IPadInteractor) UpdateIPad(ipad *model.IPad) (err error) {
	if ipad.ID <= 0 {
		return fmt.Errorf("cannot logical update ipad because invalid ipad id: %d", ipad.ID)
	}
	err = interactor.IPadRepository.UpdateIPad(ipad)
	return
}

// UpdateAllSoldTemporary 一旦全てを売り切れ判定にする
func (interactor *IPadInteractor) UpdateAllSoldTemporary() (err error) {
	err = interactor.IPadRepository.UpdateAllSoldTemporary()
	return
}

// RemoveIPad ipadの情報を削除
func (interactor *IPadInteractor) RemoveIPad(id int64) (err error) {
	err = interactor.IPadRepository.RemoveIPad(id)
	return
}
