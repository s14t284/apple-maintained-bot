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

func TestNewMacService(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := database.NewMockMacRepository(ctrl)
	{
		// success
		mpi := NewMacService(mockMr)
		a.NotNil(mpi)
	}
	{
		// failed because database is nil
		mpi := NewMacService(nil)
		a.Nil(mpi)
	}
}

func TestMacService_FindMac(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMpr := database.NewMockMacRepository(ctrl)
	expected := make(model.Macs, 1)
	expected[0] = model.Mac{Name: "MacBook Pro"}
	{
		// success
		ipi := NewMacService(mockMpr)
		if ipi == nil {
			t.FailNow()
		}
		mockMpr.EXPECT().FindMac(&model.MacRequestParam{}).Return(expected, nil)
		actual, err := ipi.FindMac(&model.MacRequestParam{})
		a.NotNil(actual)
		a.NoError(err)
		a.Equal(expected, actual)
	}
	{
		// failed
		ipi := NewMacService(mockMpr)
		if ipi == nil {
			t.FailNow()
		}
		mockMpr.EXPECT().FindMac(&model.MacRequestParam{}).Return(nil, fmt.Errorf("error"))
		actual, err := ipi.FindMac(&model.MacRequestParam{})
		a.Nil(actual)
		a.Error(err)
	}
}

func TestMacService_FindMacAll(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := database.NewMockMacRepository(ctrl)
	var expected model.Macs = make(model.Macs, 1)
	expected[0] = model.Mac{}
	{
		// success
		mpi := NewMacService(mockMr)
		if mpi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().FindMacAll().Return(expected, nil)
		actual, err := mpi.FindMacAll()
		a.Equal(expected, actual)
		a.NoError(err)
	}
	{
		// failed
		mpi := NewMacService(mockMr)
		if mpi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().FindMacAll().Return(nil, fmt.Errorf("error"))
		actual, err := mpi.FindMacAll()
		a.Nil(actual)
		a.Error(err)
	}
}

func TestMacService_FindByURL(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := database.NewMockMacRepository(ctrl)
	expected := &model.Mac{}
	url := "https://apple.com"
	{
		// success
		mpi := NewMacService(mockMr)
		if mpi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().FindByURL(url).Return(expected, nil)
		actual, err := mpi.FindByURL(url)
		a.Equal(expected, actual)
		a.NoError(err)
	}
	{
		// failed
		mpi := NewMacService(mockMr)
		if mpi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().FindByURL(url).Return(nil, fmt.Errorf("error"))
		actual, err := mpi.FindByURL(url)
		a.Nil(actual)
		a.Error(err)
	}
}

func TestMacService_IsExist(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := database.NewMockMacRepository(ctrl)
	input := &model.Mac{}

	// mock output
	eIsExist := true
	eID := uint(1)
	eT := time.Now()
	{
		// success
		mpi := NewMacService(mockMr)
		if mpi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().IsExist(input).Return(eIsExist, eID, eT, nil)
		aIsExist, aID, aT, err := mpi.IsExist(input)
		a.NoError(err)
		a.Equal(eIsExist, aIsExist)
		a.Equal(eID, aID)
		a.Equal(eT, aT)
	}
	{
		// failed
		mpi := NewMacService(mockMr)
		if mpi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().IsExist(input).Return(false, uint(0), time.Time{}, fmt.Errorf("error"))
		aIsExist, aID, aT, err := mpi.IsExist(input)
		a.Error(err)
		a.False(aIsExist)
		a.Equal(uint(0), aID)
		a.Equal(time.Time{}, aT)
	}
}

func TestMacService_AddMac(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := database.NewMockMacRepository(ctrl)
	input := &model.Mac{}
	{
		// success
		ipi := NewMacService(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().AddMac(input).Return(nil)
		err := ipi.AddMac(input)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewMacService(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().AddMac(input).Return(fmt.Errorf("error"))
		err := ipi.AddMac(input)
		a.Error(err)
	}
}

func TestMacService_UpdateMac(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := database.NewMockMacRepository(ctrl)
	input := &model.Mac{ID: 1}
	failedInput := &model.Mac{ID: 0}
	{
		// success
		ipi := NewMacService(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().UpdateMac(input).Return(nil)
		err := ipi.UpdateMac(input)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewMacService(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().UpdateMac(failedInput).Times(0)
		err := ipi.UpdateMac(failedInput)
		a.Error(err, fmt.Errorf("cannot update mac because invalid mac id: %d", 0))
	}
}

func TestMacService_UpdateAllSoldTemporary(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := database.NewMockMacRepository(ctrl)
	{
		// success
		ipi := NewMacService(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().UpdateAllSoldTemporary().Return(nil)
		err := ipi.UpdateAllSoldTemporary()
		a.NoError(err)
	}
	{
		// failed
		ipi := NewMacService(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().UpdateAllSoldTemporary().Return(fmt.Errorf("error"))
		err := ipi.UpdateAllSoldTemporary()
		a.Error(err)
	}
}

func TestMacService_RemoveMac(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := database.NewMockMacRepository(ctrl)
	id := int64(1)
	{
		// success
		ipi := NewMacService(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().RemoveMac(id).Return(nil)
		err := ipi.RemoveMac(id)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewMacService(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().RemoveMac(id).Return(fmt.Errorf("error"))
		err := ipi.RemoveMac(id)
		a.Error(err)
	}
}
