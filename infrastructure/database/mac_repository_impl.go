package database

import (
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/s14t284/apple-maitained-bot/domain/model"
)

// MacRepositoryImpl macbookに関する情報を操作するための実装
type MacRepositoryImpl struct {
	SQLClient *SQLClient
}

var _ MacRepository = &MacRepositoryImpl{}

// FindMac 整備済み品macの情報を検索して返す
func (mr MacRepositoryImpl) FindMac(param *model.MacRequestParam) (model.Macs, error) {
	var macs model.Macs
	var def model.MacRequestParam
	m := make(map[string]interface{})
	if param.Name != def.Name {
		m["name LIKE ?"] = param.Name
	}
	if param.Color != def.Color {
		m["color = ?"] = param.Color
	}
	if param.MaxInch != def.MaxInch {
		m["inch < ?"] = param.MaxInch
	}
	if param.MinInch != def.MinInch {
		m["inch > ?"] = param.MinInch
	}
	if param.MaxStorage != def.MaxStorage {
		m["storage < ?"] = param.MaxStorage
	}
	if param.MinStorage != def.MinStorage {
		m["storage > ?"] = param.MinStorage
	}
	if param.MaxAmount != def.MaxAmount {
		m["amount < ?"] = param.MaxAmount
	}
	if param.MinAmount != def.MinAmount {
		m["amount > ?"] = param.MinAmount
	}
	if param.MaxMemory != def.MaxMemory {
		m["memory < ?"] = param.MaxMemory
	}
	if param.MinMemory != def.MinMemory {
		m["memory > ?"] = param.MinMemory
	}
	if param.IsSold != def.IsSold {
		switch param.IsSold {
		case "true":
			m["is_sold = ?"] = true
		case "false":
			m["is_sold = ?"] = false
		}
		m["is_sold = ?"] = param.IsSold
	}
	if param.TouchBar != def.TouchBar {
		switch param.TouchBar {
		case "true":
			m["touch_bar = ?"] = true
		case "false":
			m["touch_bar = ?"] = false
		}
	}
	q := mr.SQLClient.Client.Where("id > ?", 0)
	for k, v := range m {
		q = q.Where(k, v)
	}
	result := q.Order("id DESC").Find(&macs)
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
