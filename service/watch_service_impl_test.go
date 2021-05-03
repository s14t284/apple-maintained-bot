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

func TestNewWatchServiceImpl(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	{
		// success
		mpi := NewWatchServiceImpl(mockWr)
		a.NotNil(mpi)
	}
	{
		// failed because database is nil
		mpi := NewWatchServiceImpl(nil)
		a.Nil(mpi)
	}
}

func TestWatchServiceImpl_FindWatch(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWpr := database.NewMockWatchRepository(ctrl)
	expected := make(model.Watches, 1)
	expected[0] = model.Watch{Name: "Apple Watch Series 4"}
	{
		// success
		ipi := NewWatchServiceImpl(mockWpr)
		if ipi == nil {
			t.FailNow()
		}
		mockWpr.EXPECT().FindWatch(&model.WatchRequestParam{}).Return(expected, nil)
		actual, err := ipi.Find(&model.WatchRequestParam{})
		a.NotNil(actual)
		a.NoError(err)
		a.Equal(expected, actual)
	}
	{
		// failed
		ipi := NewWatchServiceImpl(mockWpr)
		if ipi == nil {
			t.FailNow()
		}
		mockWpr.EXPECT().FindWatch(&model.WatchRequestParam{}).Return(nil, fmt.Errorf("error"))
		actual, err := ipi.Find(&model.WatchRequestParam{})
		a.Nil(actual)
		a.Error(err)
	}
}

func TestWatchServiceImpl_FindWatchAll(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	var expected model.Watches = make(model.Watches, 1)
	expected[0] = model.Watch{}
	{
		// success
		mpi := NewWatchServiceImpl(mockWr)
		if mpi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().FindWatchAll().Return(expected, nil)
		actual, err := mpi.FindAll()
		a.Equal(expected, actual)
		a.NoError(err)
	}
	{
		// failed
		mpi := NewWatchServiceImpl(mockWr)
		if mpi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().FindWatchAll().Return(nil, fmt.Errorf("error"))
		actual, err := mpi.FindAll()
		a.Nil(actual)
		a.Error(err)
	}
}

func TestWatchServiceImpl_FindByURL(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	expected := &model.Watch{}
	url := "https://apple.com"
	{
		// success
		mpi := NewWatchServiceImpl(mockWr)
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
		mpi := NewWatchServiceImpl(mockWr)
		if mpi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().FindByURL(url).Return(nil, fmt.Errorf("error"))
		actual, err := mpi.FindByURL(url)
		a.Nil(actual)
		a.Error(err)
	}
}

func TestWatchServiceImpl_IsExist(t *testing.T) {
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
		mpi := NewWatchServiceImpl(mockWr)
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
		mpi := NewWatchServiceImpl(mockWr)
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

func TestWatchServiceImpl_AddWatch(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	input := &model.Watch{}
	{
		// success
		ipi := NewWatchServiceImpl(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().AddWatch(input).Return(nil)
		err := ipi.Add(input)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewWatchServiceImpl(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().AddWatch(input).Return(fmt.Errorf("error"))
		err := ipi.Add(input)
		a.Error(err)
	}
}

func TestWatchServiceImpl_UpdateWatch(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	input := &model.Watch{ID: 1}
	failedInput := &model.Watch{ID: 0}
	{
		// success
		ipi := NewWatchServiceImpl(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().UpdateWatch(input).Return(nil)
		err := ipi.Update(input)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewWatchServiceImpl(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().UpdateWatch(failedInput).Times(0)
		err := ipi.Update(failedInput)
		a.EqualError(err, fmt.Sprintf("cannot update watch because invalid watch id: %d", 0))
	}
}

func TestWatchServiceImpl_UpdateAllSoldTemporary(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	{
		// success
		ipi := NewWatchServiceImpl(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().UpdateAllSoldTemporary().Return(nil)
		err := ipi.UpdateAllSoldTemporary()
		a.NoError(err)
	}
	{
		// failed
		ipi := NewWatchServiceImpl(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().UpdateAllSoldTemporary().Return(fmt.Errorf("error"))
		err := ipi.UpdateAllSoldTemporary()
		a.Error(err)
	}
}

func TestWatchServiceImpl_RemoveWatch(t *testing.T) {
	a := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	id := int64(1)
	{
		// success
		ipi := NewWatchServiceImpl(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().RemoveWatch(id).Return(nil)
		err := ipi.Remove(id)
		a.NoError(err)
	}
	{
		// failed
		ipi := NewWatchServiceImpl(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().RemoveWatch(id).Return(fmt.Errorf("error"))
		err := ipi.Remove(id)
		a.Error(err)
	}
}
