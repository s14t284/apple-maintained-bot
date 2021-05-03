package parse

import (
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"

	"github.com/s14t284/apple-maitained-bot/domain"
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/infrastructure/web"
	mock "github.com/s14t284/apple-maitained-bot/mock/web"
)

func TestNewPageParseServiceImpl(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPpr := mock.NewMockPageParseRepository(ctrl)
	ppsi, err := NewPageParseServiceImpl(mockPpr)
	if err != nil {
		t.FailNow()
	}

	// when
	type args struct {
		ppr web.PageParseRepository
	}
	var tests = []struct {
		name    string
		args    args
		want    *PageParseServiceImpl
		wantErr bool
	}{
		{"正常系", args{mockPpr}, ppsi, false},
		{"異常系_repositoryがnil", args{nil}, nil, true},
	}
	// then
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewPageParseServiceImpl(tt.args.ppr)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewPageParseServiceImpl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPageParseServiceImpl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPageParseServiceImpl_ParsePage(t *testing.T) {
	// given
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockPpr := mock.NewMockPageParseRepository(ctrl)
	page := domain.Page{}
	mac := model.Mac{}
	ipad := model.IPad{}
	watch := model.Watch{}
	mockPpr.EXPECT().ParseMacPage(page).Return(&mac, nil).Times(1)
	mockPpr.EXPECT().ParseIPadPage(page).Return(&ipad, nil).Times(1)
	mockPpr.EXPECT().ParseWatchPage(page).Return(&watch, nil).Times(1)

	// when
	type fields struct {
		ppr web.PageParseRepository
	}
	type args struct {
		target string
		page   domain.Page
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		{"正常系_mac", fields{mockPpr}, args{"mac", page}, &mac, false},
		{"正常系_ipad", fields{mockPpr}, args{"ipad", page}, &ipad, false},
		{"正常系_watch", fields{mockPpr}, args{"watch", page}, &watch, false},
		{"異常系", fields{mockPpr}, args{"chromebook", page}, nil, true},
	}
	// then
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ppi := &PageParseServiceImpl{
				ppr: tt.fields.ppr,
			}
			got, err := ppi.ParsePage(tt.args.target, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParsePage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParsePage() got = %v, want %v", got, tt.want)
			}
		})
	}
}
