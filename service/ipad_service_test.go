package service

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/mock/database"
)

func TestNewIPadService(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	{
		// success
		ipi := NewIPadService(mockIpr)
		a.NotNil(ipi)
	}
	{
		// failed because database is nil
		ipi := NewIPadService(nil)
		a.Nil(ipi)
	}
}

func TestIPadService_FindIPad(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	expected := make(model.IPads, 1)
	expected[0] = model.IPad{Name: "IPad Pro"}
	{
		// success
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindIPad(&model.IPadRequestParam{}).Return(expected, nil)
		actual, err := ipi.FindIPad(&model.IPadRequestParam{})
		a.NotNil(actual)
		a.NoError(err)
		a.Equal(expected, actual)
	}
	{
		// failed
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindIPad(&model.IPadRequestParam{}).Return(nil, fmt.Errorf("error"))
		actual, err := ipi.FindIPad(&model.IPadRequestParam{})
		a.Nil(actual)
		a.Error(err)
	}
}

func TestIPadService_FindIPadAll(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	var expected model.IPads = make(model.IPads, 1)
	expected[0] = model.IPad{}
	{
		// success
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindIPadAll().Return(expected, nil)
		actual, err := ipi.FindIPadAll()
		a.NotNil(actual)
		a.NoError(err)
		a.Equal(expected, actual)
	}
	{
		// failed
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindIPadAll().Return(nil, fmt.Errorf("error"))
		actual, err := ipi.FindIPadAll()
		a.Nil(actual)
		a.Error(err)
	}
}

func TestIPadService_FindByURL(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	expected := &model.IPad{}
	url := "https://apple.com"
	{
		// success
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindByURL(url).Return(expected, nil)
		actual, err := ipi.FindByURL(url)
		a.NotNil(actual)
		a.NoError(err)
		a.Equal(expected, actual)
	}
	{
		// failed
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindByURL(url).Return(nil, fmt.Errorf("error"))
		actual, err := ipi.FindByURL(url)
		a.Nil(actual)
		a.Error(err)
	}
}

func TestIPadService_IsExist(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	input := &model.IPad{}

	// mock output
	eIsExist := true
	eID := uint(1)
	eT := time.Now()
	{
		// success
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().IsExist(input).Return(eIsExist, eID, eT, nil)
		aIsExist, aID, aT, err := ipi.IsExist(input)
		a.NoError(err)
		a.Equal(eIsExist, aIsExist)
		a.Equal(eID, aID)
		a.Equal(eT, aT)
	}
	{
		// failed
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().IsExist(input).Return(false, uint(0), time.Time{}, fmt.Errorf("error"))
		aIsExist, aID, aT, err := ipi.IsExist(input)
		a.Error(err)
		a.False(aIsExist)
		a.Equal(uint(0), aID)
		a.Equal(time.Time{}, aT)
	}
}

func TestIPadService_AddIPad(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	input := &model.IPad{}
	{
		// success
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().AddIPad(input).Return(nil)
		err := ipi.AddIPad(input)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().AddIPad(input).Return(fmt.Errorf("error"))
		err := ipi.AddIPad(input)
		a.Error(err)
	}
}

func TestIPadService_UpdateIPad(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	input := &model.IPad{ID: 1}
	failedInput := &model.IPad{ID: 0}
	{
		// success
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().UpdateIPad(input).Return(nil)
		err := ipi.UpdateIPad(input)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().UpdateIPad(failedInput).Times(0)
		err := ipi.UpdateIPad(failedInput)
		a.Error(err, fmt.Errorf("cannot update ipad because invalid ipad id: %d", 0))
	}
}

func TestIPadService_UpdateAllSoldTemporary(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	{
		// success
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().UpdateAllSoldTemporary().Return(nil)
		err := ipi.UpdateAllSoldTemporary()
		a.NoError(err)
	}
	{
		// failed
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().UpdateAllSoldTemporary().Return(fmt.Errorf("error"))
		err := ipi.UpdateAllSoldTemporary()
		a.Error(err)
	}
}

func TestIPadService_RemoveIPad(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	id := int64(1)
	{
		// success
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().RemoveIPad(id).Return(nil)
		err := ipi.RemoveIPad(id)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewIPadService(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().RemoveIPad(id).Return(fmt.Errorf("error"))
		err := ipi.RemoveIPad(id)
		a.Error(err)
	}
}
