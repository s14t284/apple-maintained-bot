// TODO: ロジックの検証しかしていないので、integration testで疎通確認する
package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	mockUseCase "github.com/s14t284/apple-maitained-bot/mock/usecase"
	"github.com/s14t284/apple-maitained-bot/usecase"
)

func TestGetController_GetMacHandler(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	mockMu := mockUseCase.NewMockMacUseCase(ctrl)
	mockIu := mockUseCase.NewMockIPadUseCase(ctrl)
	mockWu := mockUseCase.NewMockWatchUseCase(ctrl)

	mrp := model.MacRequestParam{
		Name:       "MacBook Pro",
		Color:      "ゴールド",
		IsSold:     "false",
		MaxInch:    20,
		MinInch:    10,
		MaxMemory:  64,
		MinMemory:  8,
		MaxStorage: 1000,
		MinStorage: 256,
		MaxAmount:  100000,
		MinAmount:  50000,
	}
	input := url.Values{}
	input.Set("name", "pro")
	input.Set("color", "ゴールド")
	input.Set("is_sold", "false")
	input.Set("max_amount", "100000")
	input.Set("min_amount", "50000")
	input.Set("max_inch", "20")
	input.Set("min_inch", "10")
	input.Set("max_storage", "1000")
	input.Set("min_storage", "256")
	input.Set("max_memory", "64")
	input.Set("min_memory", "8")
	header := http.Header{}
	header.Set("Content-Type", "application/json")

	expectedEntity := model.Macs{model.Mac{Name: "MacBook Pro"}}
	expected, _ := json.Marshal(expectedEntity)
	mockMu.EXPECT().GetMacs(mrp).Return(expectedEntity, nil).Times(1)
	{
		// 正常系
		gc, err := NewGetController(mockMu, mockIu, mockWu)
		if err != nil {
			t.FailNow()
		}
		req := &http.Request{
			Method: http.MethodGet,
			Form:   input,
			Header: header,
		}
		got := httptest.NewRecorder()
		gc.GetMacHandler(got, req)

		a.Equal(http.StatusOK, got.Code)
		a.Equal(string(expected), got.Body.String())
	}
}

func TestGetController_GetIPadHandler(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	mockMu := mockUseCase.NewMockMacUseCase(ctrl)
	mockIu := mockUseCase.NewMockIPadUseCase(ctrl)
	mockWu := mockUseCase.NewMockWatchUseCase(ctrl)

	irp := model.IPadRequestParam{
		Name:       "IPad Pro",
		Color:      "ゴールド",
		IsSold:     "false",
		MaxInch:    20,
		MinInch:    10,
		MaxStorage: 64,
		MinStorage: 8,
		MaxAmount:  100000,
		MinAmount:  50000,
	}
	input := url.Values{}
	input.Set("name", "pro")
	input.Set("color", "ゴールド")
	input.Set("is_sold", "false")
	input.Set("max_amount", "100000")
	input.Set("min_amount", "50000")
	input.Set("max_inch", "20")
	input.Set("min_inch", "10")
	input.Set("max_storage", "64")
	input.Set("min_storage", "8")
	header := http.Header{}
	header.Set("Content-Type", "application/json")

	expectedEntity := model.IPads{model.IPad{Name: "IPad Pro"}}
	expected, _ := json.Marshal(expectedEntity)
	mockIu.EXPECT().GetIPads(irp).Return(expectedEntity, nil).Times(1)
	{
		// 正常系
		gc, err := NewGetController(mockMu, mockIu, mockWu)
		if err != nil {
			t.FailNow()
		}
		req := &http.Request{
			Method: http.MethodGet,
			Form:   input,
			Header: header,
		}
		got := httptest.NewRecorder()
		gc.GetIPadHandler(got, req)

		a.Equal(http.StatusOK, got.Code)
		a.Equal(string(expected), got.Body.String())
	}
}

func TestGetController_GetWatchHandler(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	mockMu := mockUseCase.NewMockMacUseCase(ctrl)
	mockIu := mockUseCase.NewMockIPadUseCase(ctrl)
	mockWu := mockUseCase.NewMockWatchUseCase(ctrl)

	wrp := model.WatchRequestParam{
		Name:       "NAME",
		Color:      "ゴールド",
		IsSold:     "false",
		MaxInch:    20,
		MinInch:    10,
		MaxStorage: 64,
		MinStorage: 8,
		MaxAmount:  100000,
		MinAmount:  50000,
	}
	input := url.Values{}
	input.Set("name", "name")
	input.Set("color", "ゴールド")
	input.Set("is_sold", "false")
	input.Set("max_amount", "100000")
	input.Set("min_amount", "50000")
	input.Set("max_inch", "20")
	input.Set("min_inch", "10")
	input.Set("max_storage", "64")
	input.Set("min_storage", "8")
	header := http.Header{}
	header.Set("Content-Type", "application/json")

	expectedEntity := model.Watches{model.Watch{Name: "AppleWatchSE"}}
	expected, _ := json.Marshal(expectedEntity)
	mockWu.EXPECT().GetWatches(wrp).Return(expectedEntity, nil).Times(1)
	{
		// 正常系
		gc, err := NewGetController(mockMu, mockIu, mockWu)
		if err != nil {
			t.FailNow()
		}
		req := &http.Request{
			Method: http.MethodGet,
			Form:   input,
			Header: header,
		}
		got := httptest.NewRecorder()
		gc.GetWatchHandler(got, req)

		a.Equal(http.StatusOK, got.Code)
		a.Equal(string(expected), got.Body.String())
	}
}

func TestGetController_HealthCheck(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	mockMu := mockUseCase.NewMockMacUseCase(ctrl)
	mockIu := mockUseCase.NewMockIPadUseCase(ctrl)
	mockWu := mockUseCase.NewMockWatchUseCase(ctrl)
	{
		// 正常系
		gc, err := NewGetController(mockMu, mockIu, mockWu)
		if err != nil {
			t.FailNow()
		}
		req := &http.Request{
			Method: http.MethodGet,
		}
		got := httptest.NewRecorder()
		gc.HealthCheck(got, req)

		a.Equal(http.StatusOK, got.Code)
		a.Equal("{\"message\": \"ok\"}", got.Body.String())
	}
}

func TestNewGetController(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMu := mockUseCase.NewMockMacUseCase(ctrl)
	mockIu := mockUseCase.NewMockIPadUseCase(ctrl)
	mockWu := mockUseCase.NewMockWatchUseCase(ctrl)
	type args struct {
		mu usecase.MacUseCase
		iu usecase.IPadUseCase
		wu usecase.WatchUseCase
	}
	tests := []struct {
		name    string
		args    args
		want    *GetController
		wantErr bool
	}{
		{"正常系", args{mockMu, mockIu, mockWu}, &GetController{mockMu, mockIu, mockWu}, false},
		{"異常系_MacUseCaseがnil", args{nil, mockIu, mockWu}, nil, true},
		{"異常系_IPadUseCaseがnil", args{mockMu, nil, mockWu}, nil, true},
		{"異常系_WatchUseCaseがnil", args{mockMu, mockIu, nil}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGetController(tt.args.mu, tt.args.iu, tt.args.wu)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGetController() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGetController() got = %v, want %v", got, tt.want)
			}
		})
	}
}
