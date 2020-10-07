package database

import (
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
)

// IPadRepositoryImpl ipadに関する情報を操作するための実装
type IPadRepositoryImpl struct {
	SQLClient *infrastructure.SQLClient
}

// FindIPadAll 整備済みipadの全ての情報を返す
func (ipadRepository *IPadRepositoryImpl) FindIPadAll() (model.IPads, error) {
	var ipads model.IPads
	result := ipadRepository.SQLClient.Client.Where("is_sold is false").Order("id DESC").Find(&ipads) // idの大きい順
	if result.Error != nil {
		return nil, result.Error
	}
	return ipads, nil
}

// FindByURL 指定したURLに一致するipadを取得
func (ipadRepository *IPadRepositoryImpl) FindByURL(url string) (*model.IPad, error) {
	var ipad model.IPad
	result := ipadRepository.SQLClient.Client.Where("url = ?", url).Find(&ipad)
	if ipad.URL != url {
		return nil, result.Error
	}
	return &ipad, result.Error
}

// AddIPad 整備済み品ipadの情報を保存する
func (ipadRepository *IPadRepositoryImpl) AddIPad(ipad *model.IPad) error {
	result := ipadRepository.SQLClient.Client.Create(ipad)
	return result.Error
}

// UpdateIPad 整備済み品ipad情報を更新する
func (ipadRepository *IPadRepositoryImpl) UpdateIPad(ipad *model.IPad) (err error) {
	result := ipadRepository.SQLClient.Client.Save(ipad)
	return result.Error
}

// UpdateAllSoldTemporary 一旦全てを売り切れ判定にする
func (ipadRepository *IPadRepositoryImpl) UpdateAllSoldTemporary() error {
	result := ipadRepository.SQLClient.Client.Exec("UPDATE ipads SET is_sold = true")
	return result.Error
}

// RemoveIPad 整備済み品ipad情報を削除する
func (ipadRepository *IPadRepositoryImpl) RemoveIPad(id int64) error {
	result := ipadRepository.SQLClient.Client.Delete(&model.IPad{}, id)
	return result.Error
}
