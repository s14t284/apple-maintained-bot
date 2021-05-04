//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"github.com/s14t284/apple-maitained-bot/domain/model"
)

// MacUseCase Macに関する情報をやりとりするUseCase
type MacUseCase interface {
	GetMacs(mrp model.MacRequestParam) (model.Macs, error)
}
