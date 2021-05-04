package usecase

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	ms "github.com/s14t284/apple-maitained-bot/mock/service"
	"github.com/s14t284/apple-maitained-bot/service"
)

func TestWatchInteractor_GetWatches(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMs := ms.NewMockWatchService(ctrl)
	mrp1 := model.WatchRequestParam{}
	mrp2 := model.WatchRequestParam{Color: "グリーン"}
	mrp3 := model.WatchRequestParam{Name: "surface"}
	rt1 := model.Watches{model.Watch{}}
	rt2 := model.Watches{}
	mockMs.EXPECT().Find(&mrp1).Return(rt1, nil).Times(1)
	mockMs.EXPECT().Find(&mrp2).Return(rt2, nil).Times(1)
	mockMs.EXPECT().Find(&mrp3).Return(nil, fmt.Errorf("dummy")).Times(1)

	type fields struct {
		ms service.WatchService
	}
	type args struct {
		mrp model.WatchRequestParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Watches
		wantErr bool
	}{
		{"正常系", fields{mockMs}, args{mrp1}, rt1, false},
		{"正常系_結果が空", fields{mockMs}, args{mrp2}, rt2, false},
		{"異常系_検索できなかった", fields{mockMs}, args{mrp3}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mi := &WatchInteractor{
				ms: tt.fields.ms,
			}
			got, err := mi.GetWatches(tt.args.mrp)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWatches() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWatches() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewWatchInteractor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMs := ms.NewMockWatchService(ctrl)

	type args struct {
		ms service.WatchService
	}
	tests := []struct {
		name    string
		args    args
		want    *WatchInteractor
		wantErr bool
	}{
		{"正常系", args{mockMs}, &WatchInteractor{mockMs}, false},
		{"異常系_WatchServiceがnil", args{mockMs}, &WatchInteractor{mockMs}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewWatchInteractor(tt.args.ms)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewWatchInteractor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWatchInteractor() got = %v, want %v", got, tt.want)
			}
		})
	}
}
