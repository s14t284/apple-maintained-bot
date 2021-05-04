//go:generate mockgen -source=$GOFILE -package=$GOPACKAGE -destination=../mock/$GOPACKAGE/$GOFILE
package usecase

import (
	"github.com/s14t284/apple-maitained-bot/domain/model"
)

// IPadUseCase IPadに関する情報をやりとりするUseCase
type IPadUseCase interface {
	GetIPads(irp model.IPadRequestParam) (model.IPads, error)
}
