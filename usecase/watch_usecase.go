package usecase

import (
	"github.com/s14t284/apple-maitained-bot/domain/model"
)

// WatchUseCase AppleWatchに関する情報をやりとりするUseCase
type WatchUseCase interface {
	GetWatches(wrp model.WatchRequestParam) (model.Watches, error)
}
