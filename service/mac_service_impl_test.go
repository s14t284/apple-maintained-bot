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

func TestNewMacServiceImpl(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := database.NewMockMacRepository(ctrl)
	{
		// success
		mpi := NewMacServiceImpl(mockMr)
		a.NotNil(mpi)
	}
	{
		// failed because database is nil
		mpi := NewMacServiceImpl(nil)
		a.Nil(mpi)
	}
}

func TestMacServiceImpl_FindMac(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMpr := database.NewMockMacRepository(ctrl)
	expected := make(model.Macs, 1)
	expected[0] = model.Mac{Name: "MacBook Pro"}
	{
		// success
		ipi := NewMacServiceImpl(mockMpr)
		if ipi == nil {
			t.FailNow()
		}
		mockMpr.EXPECT().FindMac(&model.MacRequestParam{}).Return(expected, nil)
		actual, err := ipi.Find(&model.MacRequestParam{})
		a.NotNil(actual)
		a.NoError(err)
		a.Equal(expected, actual)
	}
	{
		// failed
		ipi := NewMacServiceImpl(mockMpr)
		if ipi == nil {
			t.FailNow()
		}
		mockMpr.EXPECT().FindMac(&model.MacRequestParam{}).Return(nil, fmt.Errorf("error"))
		actual, err := ipi.Find(&model.MacRequestParam{})
		a.Nil(actual)
		a.Error(err)
	}
}

func TestMacServiceImpl_FindMacAll(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := database.NewMockMacRepository(ctrl)
	var expected model.Macs = make(model.Macs, 1)
	expected[0] = model.Mac{}
	{
		// success
		mpi := NewMacServiceImpl(mockMr)
		if mpi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().FindMacAll().Return(expected, nil)
		actual, err := mpi.FindAll()
		a.Equal(expected, actual)
		a.NoError(err)
	}
	{
		// failed
		mpi := NewMacServiceImpl(mockMr)
		if mpi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().FindMacAll().Return(nil, fmt.Errorf("error"))
		actual, err := mpi.FindAll()
		a.Nil(actual)
		a.Error(err)
	}
}

func TestMacServiceImpl_FindByURL(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := database.NewMockMacRepository(ctrl)
	expected := &model.Mac{}
	url := "https://apple.com"
	{
		// success
		mpi := NewMacServiceImpl(mockMr)
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
		mpi := NewMacServiceImpl(mockMr)
		if mpi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().FindByURL(url).Return(nil, fmt.Errorf("error"))
		actual, err := mpi.FindByURL(url)
		a.Nil(actual)
		a.Error(err)
	}
}

func TestMacServiceImpl_IsExist(t *testing.T) {
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
		mpi := NewMacServiceImpl(mockMr)
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
		mpi := NewMacServiceImpl(mockMr)
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

func TestMacServiceImpl_AddMac(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := database.NewMockMacRepository(ctrl)
	input := &model.Mac{}
	{
		// success
		ipi := NewMacServiceImpl(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().AddMac(input).Return(nil)
		err := ipi.Add(input)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewMacServiceImpl(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().AddMac(input).Return(fmt.Errorf("error"))
		err := ipi.Add(input)
		a.Error(err)
	}
}

func TestMacServiceImpl_UpdateMac(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := database.NewMockMacRepository(ctrl)
	input := &model.Mac{ID: 1}
	failedInput := &model.Mac{ID: 0}
	{
		// success
		ipi := NewMacServiceImpl(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().UpdateMac(input).Return(nil)
		err := ipi.Update(input)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewMacServiceImpl(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().UpdateMac(failedInput).Times(0)
		err := ipi.Update(failedInput)
		a.Error(err, fmt.Errorf("cannot update mac because invalid mac id: %d", 0))
	}
}

func TestMacServiceImpl_UpdateAllSoldTemporary(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := database.NewMockMacRepository(ctrl)
	{
		// success
		ipi := NewMacServiceImpl(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().UpdateAllSoldTemporary().Return(nil)
		err := ipi.UpdateAllSoldTemporary()
		a.NoError(err)
	}
	{
		// failed
		ipi := NewMacServiceImpl(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().UpdateAllSoldTemporary().Return(fmt.Errorf("error"))
		err := ipi.UpdateAllSoldTemporary()
		a.Error(err)
	}
}

func TestMacServiceImpl_RemoveMac(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := database.NewMockMacRepository(ctrl)
	id := int64(1)
	{
		// success
		ipi := NewMacServiceImpl(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().RemoveMac(id).Return(nil)
		err := ipi.Remove(id)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewMacServiceImpl(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().RemoveMac(id).Return(fmt.Errorf("error"))
		err := ipi.Remove(id)
		a.Error(err)
	}
}
