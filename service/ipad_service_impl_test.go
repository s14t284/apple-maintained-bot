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

func TestNewIPadServiceImpl(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	{
		// success
		ipi := NewIPadServiceImpl(mockIpr)
		a.NotNil(ipi)
	}
	{
		// failed because database is nil
		ipi := NewIPadServiceImpl(nil)
		a.Nil(ipi)
	}
}

func TestIPadServiceImpl_FindIPad(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	expected := make(model.IPads, 1)
	expected[0] = model.IPad{Name: "IPad Pro"}
	{
		// success
		ipi := NewIPadServiceImpl(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindIPad(&model.IPadRequestParam{}).Return(expected, nil)
		actual, err := ipi.Find(&model.IPadRequestParam{})
		a.NotNil(actual)
		a.NoError(err)
		a.Equal(expected, actual)
	}
	{
		// failed
		ipi := NewIPadServiceImpl(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindIPad(&model.IPadRequestParam{}).Return(nil, fmt.Errorf("error"))
		actual, err := ipi.Find(&model.IPadRequestParam{})
		a.Nil(actual)
		a.Error(err)
	}
}

func TestIPadServiceImpl_FindIPadAll(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	var expected model.IPads = make(model.IPads, 1)
	expected[0] = model.IPad{}
	{
		// success
		ipi := NewIPadServiceImpl(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindIPadAll().Return(expected, nil)
		actual, err := ipi.FindAll()
		a.NotNil(actual)
		a.NoError(err)
		a.Equal(expected, actual)
	}
	{
		// failed
		ipi := NewIPadServiceImpl(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindIPadAll().Return(nil, fmt.Errorf("error"))
		actual, err := ipi.FindAll()
		a.Nil(actual)
		a.Error(err)
	}
}

func TestIPadServiceImpl_FindByURL(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	expected := &model.IPad{}
	url := "https://apple.com"
	{
		// success
		ipi := NewIPadServiceImpl(mockIpr)
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
		ipi := NewIPadServiceImpl(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindByURL(url).Return(nil, fmt.Errorf("error"))
		actual, err := ipi.FindByURL(url)
		a.Nil(actual)
		a.Error(err)
	}
}

func TestIPadServiceImpl_IsExist(t *testing.T) {
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
		ipi := NewIPadServiceImpl(mockIpr)
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
		ipi := NewIPadServiceImpl(mockIpr)
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

func TestIPadServiceImpl_AddIPad(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	input := &model.IPad{}
	{
		// success
		ipi := NewIPadServiceImpl(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().AddIPad(input).Return(nil)
		err := ipi.Add(input)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewIPadServiceImpl(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().AddIPad(input).Return(fmt.Errorf("error"))
		err := ipi.Add(input)
		a.Error(err)
	}
}

func TestIPadServiceImpl_UpdateIPad(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	input := &model.IPad{ID: 1}
	failedInput := &model.IPad{ID: 0}
	{
		// success
		ipi := NewIPadServiceImpl(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().UpdateIPad(input).Return(nil)
		err := ipi.Update(input)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewIPadServiceImpl(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().UpdateIPad(failedInput).Times(0)
		err := ipi.Update(failedInput)
		a.Error(err, fmt.Errorf("cannot update ipad because invalid ipad id: %d", 0))
	}
}

func TestIPadServiceImpl_UpdateAllSoldTemporary(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	{
		// success
		ipi := NewIPadServiceImpl(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().UpdateAllSoldTemporary().Return(nil)
		err := ipi.UpdateAllSoldTemporary()
		a.NoError(err)
	}
	{
		// failed
		ipi := NewIPadServiceImpl(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().UpdateAllSoldTemporary().Return(fmt.Errorf("error"))
		err := ipi.UpdateAllSoldTemporary()
		a.Error(err)
	}
}

func TestIPadServiceImpl_RemoveIPad(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	id := int64(1)
	{
		// success
		ipi := NewIPadServiceImpl(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().RemoveIPad(id).Return(nil)
		err := ipi.Remove(id)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewIPadServiceImpl(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().RemoveIPad(id).Return(fmt.Errorf("error"))
		err := ipi.Remove(id)
		a.Error(err)
	}
}
