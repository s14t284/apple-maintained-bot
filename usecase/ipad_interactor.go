package usecase

import (
	"fmt"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/service"
)

// IPadInteractor IPadUseCaseの実装
type IPadInteractor struct {
	ms service.IPadService
}

var _ IPadUseCase = &IPadInteractor{}

// NewIPadInteractor IPadInteractorを初期化して返す
func NewIPadInteractor(ms service.IPadService) (*IPadInteractor, error) {
	if ms == nil {
		return nil, fmt.Errorf("ipad service must not be nil")
	}
	return &IPadInteractor{
		ms: ms,
	}, nil
}

// GetIPads 検索条件に当てはまるmacエンティティのリストを返す
func (mi *IPadInteractor) GetIPads(mrp model.IPadRequestParam) (model.IPads, error) {
	macs, err := mi.ms.Find(&mrp)
	return macs, err
}
