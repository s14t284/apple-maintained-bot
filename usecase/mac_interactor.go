package usecase

import (
	"fmt"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure/database"
)

// MacInteractor macbookの情報をやりとりするためのgateway
type MacInteractor struct {
	MacRepository database.MacRepositoryImpl
}

// FindMacAll macbookの情報を全て取得
func (interactor *MacInteractor) FindMacAll() (model.Macs, error) {
	macs, err := interactor.MacRepository.FindMacAll()
	return macs, err
}

// AddMac macbookの情報を追加
func (interactor *MacInteractor) AddMac(mac *model.Mac) (err error) {
	err = interactor.MacRepository.AddMac(mac)
	return
}

// UpdateMac macの情報の更新
func (interactor *MacInteractor) UpdateMac(mac *model.Mac) (err error) {
	if mac.ID <= 0 {
		return fmt.Errorf("cannot logical update mac because invalid mac id: %d", mac.ID)
	}
	err = interactor.MacRepository.UpdateMac(mac)
	return
}

// RemoveMac macbookの情報を削除
func (interactor *MacInteractor) RemoveMac(id int64) (err error) {
	err = interactor.MacRepository.RemoveMac(id)
	return
}
