package usecase

import (
	"fmt"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/service"
)

// WatchInteractor WatchUseCaseの実装
type WatchInteractor struct {
	ms service.WatchService
}

var _ WatchUseCase = &WatchInteractor{}

// NewWatchInteractor WatchInteractorを初期化して返す
func NewWatchInteractor(ms service.WatchService) (*WatchInteractor, error) {
	if ms == nil {
		return nil, fmt.Errorf("watch service must not be nil")
	}
	return &WatchInteractor{
		ms: ms,
	}, nil
}

// GetWatches 検索条件に当てはまるmacエンティティのリストを返す
func (mi *WatchInteractor) GetWatches(mrp model.WatchRequestParam) (model.Watches, error) {
	macs, err := mi.ms.Find(&mrp)
	return macs, err
}
