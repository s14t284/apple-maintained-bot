package repository

import "github.com/s14t284/apple-maitained-bot/domain/model"

// MacRepository 整備済み品macの情報を読み書きするクライアント
type MacRepository interface {
	FindMacAll() (model.Macs, error)
	FindByURL(url string) (*model.Mac, error)
	AddMac(mac *model.Mac) error
	UpdateMac(mac *model.Mac) error
	UpdateAllSoldTemporary() error
	RemoveMac(id int64) error
}
