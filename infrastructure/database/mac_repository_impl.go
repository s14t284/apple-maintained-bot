package database

import (
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
)

// MacRepositoryImpl macbookに関する情報を操作するための実装
type MacRepositoryImpl struct {
	SQLClient *infrastructure.SQLClient
}

// FindMacAll 整備済みmacの全ての情報を返す
func (macRepository *MacRepositoryImpl) FindMacAll() (model.Macs, error) {
	var macs model.Macs
	result := macRepository.SQLClient.Client.Find(&macs)
	if result.Error != nil {
		return nil, result.Error
	}
	return macs, nil
}

// AddMac 整備済み品macの情報を保存する
func (macRepository *MacRepositoryImpl) AddMac(mac *model.Mac) error {
	result := macRepository.SQLClient.Client.Create(mac)
	return result.Error
}

// UpdateMac  整備済み品mac情報を更新する
func (macRepository *MacRepositoryImpl) UpdateMac(mac *model.Mac) (err error) {
	result := macRepository.SQLClient.Client.Save(mac)
	return result.Error
}

// RemoveMac 整備済み品mac情報を削除する
func (macRepository *MacRepositoryImpl) RemoveMac(id int64) error {
	result := macRepository.SQLClient.Client.Delete(&model.Mac{}, id)
	return result.Error
}
