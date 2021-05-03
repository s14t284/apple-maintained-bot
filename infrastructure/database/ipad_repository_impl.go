package database

import (
	"errors"
	"time"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"gorm.io/gorm"
)

// IPadRepositoryImpl ipadに関する情報を操作するための実装
type IPadRepositoryImpl struct {
	SQLClient *SQLClient
}

// FindIPad 整備済み品ipadの情報を検索して返す
func (ipr IPadRepositoryImpl) FindIPad(param *model.IPadRequestParam) (model.IPads, error) {
	var ipads model.IPads
	var def model.IPadRequestParam
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
	if param.IsSold != def.IsSold {
		switch param.IsSold {
		case "true":
			m["is_sold = ?"] = true
		case "false":
			m["is_sold = ?"] = false
		}
		m["is_sold = ?"] = param.IsSold
	}
	q := ipr.SQLClient.Client.Where("id > ?", 0)
	for k, v := range m {
		q = q.Where(k, v)
	}
	result := q.Order("id DESC").Find(&ipads)
	if result.Error != nil {
		return nil, result.Error
	}
	return ipads, nil
}

// FindIPadAll 整備済みipadの全ての情報を返す
func (ipr IPadRepositoryImpl) FindIPadAll() (model.IPads, error) {
	var ipads model.IPads
	result := ipr.SQLClient.Client.Where("is_sold is false").Order("id DESC").Find(&ipads) // idの大きい順
	if result.Error != nil {
		return nil, result.Error
	}
	return ipads, nil
}

// FindByURL 指定したURLに一致するipadを取得
func (ipr IPadRepositoryImpl) FindByURL(url string) (*model.IPad, error) {
	var ipad model.IPad
	result := ipr.SQLClient.Client.Where("url = ?", url).Find(&ipad)
	if ipad.URL != url {
		return nil, result.Error
	}
	return &ipad, result.Error
}

// IsExist オブジェクトがDB内に存在しているかどうか
func (ipr IPadRepositoryImpl) IsExist(ipad *model.IPad) (bool, uint, time.Time, error) {
	tmp := &model.IPad{}
	err := ipr.SQLClient.Client.Where(&model.IPad{
		Name:        ipad.Name,
		Inch:        ipad.Inch,
		Storage:     ipad.Storage,
		Camera:      ipad.Camera,
		Color:       ipad.Color,
		Amount:      ipad.Amount,
		ReleaseDate: ipad.ReleaseDate}).First(tmp).Error
	if err == nil {
		return true, tmp.ID, tmp.CreatedAt, nil
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, 0, time.Time{}, nil
	}
	return false, 0, time.Time{}, err
}

// AddIPad 整備済み品ipadの情報を保存する
func (ipr IPadRepositoryImpl) AddIPad(ipad *model.IPad) error {
	result := ipr.SQLClient.Client.Create(ipad)
	return result.Error
}

// UpdateIPad 整備済み品ipad情報を更新する
func (ipr IPadRepositoryImpl) UpdateIPad(ipad *model.IPad) (err error) {
	result := ipr.SQLClient.Client.Save(ipad)
	return result.Error
}

// UpdateAllSoldTemporary 一旦全てを売り切れ判定にする
func (ipr IPadRepositoryImpl) UpdateAllSoldTemporary() error {
	result := ipr.SQLClient.Client.Exec("UPDATE ipads SET is_sold = true")
	return result.Error
}

// RemoveIPad 整備済み品ipad情報を削除する
func (ipr IPadRepositoryImpl) RemoveIPad(id int64) error {
	result := ipr.SQLClient.Client.Delete(&model.IPad{}, id)
	return result.Error
}
