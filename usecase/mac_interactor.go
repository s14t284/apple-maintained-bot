package usecase

import (
	"fmt"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/service"
)

// MacInteractor MacUseCaseの実装
type MacInteractor struct {
	ms service.MacService
}

var _ MacUseCase = &MacInteractor{}

// NewMacInteractor MacInteractorを初期化して返す
func NewMacInteractor(ms service.MacService) (*MacInteractor, error) {
	if ms == nil {
		return nil, fmt.Errorf("mac service must not be nil")
	}
	return &MacInteractor{
		ms: ms,
	}, nil
}

// GetMacs 検索条件に当てはまるmacエンティティのリストを返す
func (mi *MacInteractor) GetMacs(mrp model.MacRequestParam) (model.Macs, error) {
	macs, err := mi.ms.Find(&mrp)
	return macs, err
}
