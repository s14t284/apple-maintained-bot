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
	result := macRepository.SQLClient.Client.Where("is_sold is false").Order("id DESC").Find(&macs)
	if result.Error != nil {
		return nil, result.Error
	}
	return macs, nil
}

// FindByURL 指定したURLに一致するmacを取得
func (macRepository *MacRepositoryImpl) FindByURL(url string) (*model.Mac, error) {
	var mac model.Mac
	result := macRepository.SQLClient.Client.Where("url = ?", url).Find(&mac)
	if mac.URL != url {
		return nil, result.Error
	}
	return &mac, result.Error
}

// UpdateAllSoldTemporary 一旦全てを売り切れ判定にする
func (macRepository *MacRepositoryImpl) UpdateAllSoldTemporary() error {
	result := macRepository.SQLClient.Client.Exec("UPDATE macs SET is_sold = true")
	return result.Error
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
