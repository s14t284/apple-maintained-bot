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

func TestNewWatchService(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	{
		// success
		mpi := NewWatchService(mockWr)
		a.NotNil(mpi)
	}
	{
		// failed because database is nil
		mpi := NewWatchService(nil)
		a.Nil(mpi)
	}
}

func TestWatchService_FindWatch(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWpr := database.NewMockWatchRepository(ctrl)
	expected := make(model.Watches, 1)
	expected[0] = model.Watch{Name: "Apple Watch Series 4"}
	{
		// success
		ipi := NewWatchService(mockWpr)
		if ipi == nil {
			t.FailNow()
		}
		mockWpr.EXPECT().FindWatch(&model.WatchRequestParam{}).Return(expected, nil)
		actual, err := ipi.FindWatch(&model.WatchRequestParam{})
		a.NotNil(actual)
		a.NoError(err)
		a.Equal(expected, actual)
	}
	{
		// failed
		ipi := NewWatchService(mockWpr)
		if ipi == nil {
			t.FailNow()
		}
		mockWpr.EXPECT().FindWatch(&model.WatchRequestParam{}).Return(nil, fmt.Errorf("error"))
		actual, err := ipi.FindWatch(&model.WatchRequestParam{})
		a.Nil(actual)
		a.Error(err)
	}
}

func TestWatchService_FindWatchAll(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	var expected model.Watches = make(model.Watches, 1)
	expected[0] = model.Watch{}
	{
		// success
		mpi := NewWatchService(mockWr)
		if mpi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().FindWatchAll().Return(expected, nil)
		actual, err := mpi.FindWatchAll()
		a.Equal(expected, actual)
		a.NoError(err)
	}
	{
		// failed
		mpi := NewWatchService(mockWr)
		if mpi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().FindWatchAll().Return(nil, fmt.Errorf("error"))
		actual, err := mpi.FindWatchAll()
		a.Nil(actual)
		a.Error(err)
	}
}

func TestWatchService_FindByURL(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	expected := &model.Watch{}
	url := "https://apple.com"
	{
		// success
		mpi := NewWatchService(mockWr)
		if mpi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().FindByURL(url).Return(expected, nil)
		actual, err := mpi.FindByURL(url)
		a.Equal(expected, actual)
		a.NoError(err)
	}
	{
		// failed
		mpi := NewWatchService(mockWr)
		if mpi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().FindByURL(url).Return(nil, fmt.Errorf("error"))
		actual, err := mpi.FindByURL(url)
		a.Nil(actual)
		a.Error(err)
	}
}

func TestWatchService_IsExist(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	input := &model.Watch{}

	// mock output
	eIsExist := true
	eID := uint(1)
	eT := time.Now()
	{
		// success
		mpi := NewWatchService(mockWr)
		if mpi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().IsExist(input).Return(eIsExist, eID, eT, nil)
		aIsExist, aID, aT, err := mpi.IsExist(input)
		a.NoError(err)
		a.Equal(eIsExist, aIsExist)
		a.Equal(eID, aID)
		a.Equal(eT, aT)
	}
	{
		// failed
		mpi := NewWatchService(mockWr)
		if mpi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().IsExist(input).Return(false, uint(0), time.Time{}, fmt.Errorf("error"))
		aIsExist, aID, aT, err := mpi.IsExist(input)
		a.Error(err)
		a.False(aIsExist)
		a.Equal(uint(0), aID)
		a.Equal(time.Time{}, aT)
	}
}

func TestWatchService_AddWatch(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	input := &model.Watch{}
	{
		// success
		ipi := NewWatchService(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().AddWatch(input).Return(nil)
		err := ipi.AddWatch(input)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewWatchService(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().AddWatch(input).Return(fmt.Errorf("error"))
		err := ipi.AddWatch(input)
		a.Error(err)
	}
}

func TestWatchService_UpdateWatch(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	input := &model.Watch{ID: 1}
	failedInput := &model.Watch{ID: 0}
	{
		// success
		ipi := NewWatchService(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().UpdateWatch(input).Return(nil)
		err := ipi.UpdateWatch(input)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewWatchService(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().UpdateWatch(failedInput).Times(0)
		err := ipi.UpdateWatch(failedInput)
		a.EqualError(err, fmt.Sprintf("cannot update watch because invalid watch id: %d", 0))
	}
}

func TestWatchService_UpdateAllSoldTemporary(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	{
		// success
		ipi := NewWatchService(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().UpdateAllSoldTemporary().Return(nil)
		err := ipi.UpdateAllSoldTemporary()
		a.NoError(err)
	}
	{
		// failed
		ipi := NewWatchService(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().UpdateAllSoldTemporary().Return(fmt.Errorf("error"))
		err := ipi.UpdateAllSoldTemporary()
		a.Error(err)
	}
}

func TestWatchService_RemoveWatch(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	id := int64(1)
	{
		// success
		ipi := NewWatchService(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().RemoveWatch(id).Return(nil)
		err := ipi.RemoveWatch(id)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewWatchService(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().RemoveWatch(id).Return(fmt.Errorf("error"))
		err := ipi.RemoveWatch(id)
		a.Error(err)
	}
}
