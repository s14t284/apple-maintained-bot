package database

import (
	"errors"
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure"
	"gorm.io/gorm"
)

// MacRepositoryImpl macbookに関する情報を操作するための実装
type MacRepositoryImpl struct {
	SQLClient *infrastructure.SQLClient
}

// FindMac 整備済み品macの情報を検索して返す
func (mr MacRepositoryImpl) FindMac(mac *model.Mac) (model.Macs, error) {
	var macs model.Macs
	result := mr.SQLClient.Client.Where(mac).Order("id DESC").Find(&macs)
	if result.Error != nil {
		return nil, result.Error
	}
	return macs, nil
}

// FindMacAll 整備済みmacの全ての情報を返す
func (mr MacRepositoryImpl) FindMacAll() (model.Macs, error) {
	var macs model.Macs
	result := mr.SQLClient.Client.Where("is_sold is false").Order("id DESC").Find(&macs)
	if result.Error != nil {
		return nil, result.Error
	}
	return macs, nil
}

// FindByURL 指定したURLに一致するmacを取得
func (mr MacRepositoryImpl) FindByURL(url string) (*model.Mac, error) {
	var mac model.Mac
	result := mr.SQLClient.Client.Where("url = ?", url).Find(&mac)
	if mac.URL != url {
		return nil, result.Error
	}
	return &mac, result.Error
}

// IsExist オブジェクトがDB内に存在しているかどうか
func (mr MacRepositoryImpl) IsExist(mac *model.Mac) (bool, uint, time.Time, error) {
	tmp := &model.Mac{}
	err := mr.SQLClient.Client.Where(
		&model.Mac{
			Name:        mac.Name,
			Inch:        mac.Inch,
			CPU:         mac.CPU,
			Memory:      mac.Memory,
			Storage:     mac.Storage,
			TouchBar:    mac.TouchBar,
			Color:       mac.Color,
			Amount:      mac.Amount,
			ReleaseDate: mac.ReleaseDate}).First(tmp).Error
	if err == nil {
		return true, tmp.ID, tmp.CreatedAt, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, 0, time.Time{}, nil
	}
	return false, 0, time.Time{}, err
}

// UpdateAllSoldTemporary 一旦全てを売り切れ判定にする
func (mr MacRepositoryImpl) UpdateAllSoldTemporary() error {
	result := mr.SQLClient.Client.Exec("UPDATE macs SET is_sold = true")
	return result.Error
}

// AddMac 整備済み品macの情報を保存する
func (mr MacRepositoryImpl) AddMac(mac *model.Mac) error {
	result := mr.SQLClient.Client.Create(mac)
	return result.Error
}

// UpdateMac  整備済み品mac情報を更新する
func (mr MacRepositoryImpl) UpdateMac(mac *model.Mac) (err error) {
	result := mr.SQLClient.Client.Save(mac)
	return result.Error
}

// RemoveMac 整備済み品mac情報を削除する
func (mr MacRepositoryImpl) RemoveMac(id int64) error {
	result := mr.SQLClient.Client.Delete(&model.Mac{}, id)
	return result.Error
}
