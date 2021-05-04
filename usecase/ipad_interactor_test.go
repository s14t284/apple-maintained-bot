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

func TestIPadInteractor_GetIPads(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMs := ms.NewMockIPadService(ctrl)
	mrp1 := model.IPadRequestParam{}
	mrp2 := model.IPadRequestParam{Color: "グリーン"}
	mrp3 := model.IPadRequestParam{Name: "surface"}
	rt1 := model.IPads{model.IPad{}}
	rt2 := model.IPads{}
	mockMs.EXPECT().Find(&mrp1).Return(rt1, nil).Times(1)
	mockMs.EXPECT().Find(&mrp2).Return(rt2, nil).Times(1)
	mockMs.EXPECT().Find(&mrp3).Return(nil, fmt.Errorf("dummy")).Times(1)

	type fields struct {
		ms service.IPadService
	}
	type args struct {
		mrp model.IPadRequestParam
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.IPads
		wantErr bool
	}{
		{"正常系", fields{mockMs}, args{mrp1}, rt1, false},
		{"正常系_結果が空", fields{mockMs}, args{mrp2}, rt2, false},
		{"異常系_検索できなかった", fields{mockMs}, args{mrp3}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mi := &IPadInteractor{
				ms: tt.fields.ms,
			}
			got, err := mi.GetIPads(tt.args.mrp)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetIPads() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetIPads() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewIPadInteractor(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockMs := ms.NewMockIPadService(ctrl)

	type args struct {
		ms service.IPadService
	}
	tests := []struct {
		name    string
		args    args
		want    *IPadInteractor
		wantErr bool
	}{
		{"正常系", args{mockMs}, &IPadInteractor{mockMs}, false},
		{"異常系_IPadServiceがnil", args{mockMs}, &IPadInteractor{mockMs}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewIPadInteractor(tt.args.ms)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewIPadInteractor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewIPadInteractor() got = %v, want %v", got, tt.want)
			}
		})
	}
}
