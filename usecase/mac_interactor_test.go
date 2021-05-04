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

func TestMacInteractor_GetMacs(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMs := ms.NewMockMacService(ctrl)
	mrp1 := model.MacRequestParam{}
	mrp2 := model.MacRequestParam{Color: "グリーン"}
	mrp3 := model.MacRequestParam{Name: "surface"}
	rt1 := model.Macs{model.Mac{}}
	rt2 := model.Macs{}
	mockMs.EXPECT().Find(&mrp1).Return(rt1, nil).Times(1)
	mockMs.EXPECT().Find(&mrp2).Return(rt2, nil).Times(1)
	mockMs.EXPECT().Find(&mrp3).Return(nil, fmt.Errorf("dummy")).Times(1)

	type fields struct {
		ms service.MacService
	}
	type args struct {
		mrp model.MacRequestParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.Macs
		wantErr bool
	}{
		{"正常系", fields{mockMs}, args{mrp1}, rt1, false},
		{"正常系_結果が空", fields{mockMs}, args{mrp2}, rt2, false},
		{"異常系_検索できなかった", fields{mockMs}, args{mrp3}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mi := &MacInteractor{
				ms: tt.fields.ms,
			}
			got, err := mi.GetMacs(tt.args.mrp)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMacs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMacs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMacInteractor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMs := ms.NewMockMacService(ctrl)

	type args struct {
		ms service.MacService
	}
	tests := []struct {
		name    string
		args    args
		want    *MacInteractor
		wantErr bool
	}{
		{"正常系", args{mockMs}, &MacInteractor{mockMs}, false},
		{"異常系_MacServiceがnil", args{mockMs}, &MacInteractor{mockMs}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMacInteractor(tt.args.ms)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMacInteractor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMacInteractor() got = %v, want %v", got, tt.want)
			}
		})
	}
}
