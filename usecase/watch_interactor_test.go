package usecase

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/mock/database"
	"github.com/stretchr/testify/assert"
)

func TestNewWatchInteractor(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	{
		// success
		mpi := NewWatchInteractor(mockWr)
		assert.NotNil(mpi)
	}
	{
		// failed because database is nil
		mpi := NewWatchInteractor(nil)
		assert.Nil(mpi)
	}
}

func TestWatchInteractor_FindWatch(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWpr := database.NewMockWatchRepository(ctrl)
	expected := make(model.Watches, 1)
	expected[0] = model.Watch{Name: "Apple Watch Series 4"}
	{
		// success
		ipi := NewWatchInteractor(mockWpr)
		if ipi == nil {
			t.FailNow()
		}
		mockWpr.EXPECT().FindWatch(&model.WatchRequestParam{}).Return(expected, nil)
		actual, err := ipi.FindWatch(&model.WatchRequestParam{})
		assert.NotNil(actual)
		assert.NoError(err)
		assert.Equal(expected, actual)
	}
	{
		// failed
		ipi := NewWatchInteractor(mockWpr)
		if ipi == nil {
			t.FailNow()
		}
		mockWpr.EXPECT().FindWatch(&model.WatchRequestParam{}).Return(nil, fmt.Errorf("error"))
		actual, err := ipi.FindWatch(&model.WatchRequestParam{})
		assert.Nil(actual)
		assert.Error(err)
	}
}

func TestWatchInteractor_FindWatchAll(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	var expected model.Watches = make(model.Watches, 1)
	expected[0] = model.Watch{}
	{
		// success
		mpi := NewWatchInteractor(mockWr)
		if mpi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().FindWatchAll().Return(expected, nil)
		actual, err := mpi.FindWatchAll()
		assert.Equal(expected, actual)
		assert.NoError(err)
	}
	{
		// failed
		mpi := NewWatchInteractor(mockWr)
		if mpi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().FindWatchAll().Return(nil, fmt.Errorf("error"))
		actual, err := mpi.FindWatchAll()
		assert.Nil(actual)
		assert.Error(err)
	}
}

func TestWatchInteractor_FindByURL(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	expected := &model.Watch{}
	url := "https://apple.com"
	{
		// success
		mpi := NewWatchInteractor(mockWr)
		if mpi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().FindByURL(url).Return(expected, nil)
		actual, err := mpi.FindByURL(url)
		assert.Equal(expected, actual)
		assert.NoError(err)
	}
	{
		// failed
		mpi := NewWatchInteractor(mockWr)
		if mpi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().FindByURL(url).Return(nil, fmt.Errorf("error"))
		actual, err := mpi.FindByURL(url)
		assert.Nil(actual)
		assert.Error(err)
	}
}

func TestWatchInteractor_IsExist(t *testing.T) {
	assert := assert.New(t)
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
		mpi := NewWatchInteractor(mockWr)
		if mpi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().IsExist(input).Return(eIsExist, eID, eT, nil)
		aIsExist, aID, aT, err := mpi.IsExist(input)
		assert.NoError(err)
		assert.Equal(eIsExist, aIsExist)
		assert.Equal(eID, aID)
		assert.Equal(eT, aT)
	}
	{
		// failed
		mpi := NewWatchInteractor(mockWr)
		if mpi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().IsExist(input).Return(false, uint(0), time.Time{}, fmt.Errorf("error"))
		aIsExist, aID, aT, err := mpi.IsExist(input)
		assert.Error(err)
		assert.False(aIsExist)
		assert.Equal(uint(0), aID)
		assert.Equal(time.Time{}, aT)
	}
}

func TestWatchInteractor_AddWatch(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	input := &model.Watch{}
	{
		// success
		ipi := NewWatchInteractor(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().AddWatch(input).Return(nil)
		err := ipi.AddWatch(input)
		assert.NoError(err)
	}
	{
		// failed
		ipi := NewWatchInteractor(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().AddWatch(input).Return(fmt.Errorf("error"))
		err := ipi.AddWatch(input)
		assert.Error(err)
	}
}

func TestWatchInteractor_UpdateWatch(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	input := &model.Watch{ID: 1}
	failedInput := &model.Watch{ID: 0}
	{
		// success
		ipi := NewWatchInteractor(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().UpdateWatch(input).Return(nil)
		err := ipi.UpdateWatch(input)
		assert.NoError(err)
	}
	{
		// failed
		ipi := NewWatchInteractor(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().UpdateWatch(failedInput).Times(0)
		err := ipi.UpdateWatch(failedInput)
		assert.EqualError(err, fmt.Sprintf("cannot update watch because invalid watch id: %d", 0))
	}
}

func TestWatchInteractor_UpdateAllSoldTemporary(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	{
		// success
		ipi := NewWatchInteractor(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().UpdateAllSoldTemporary().Return(nil)
		err := ipi.UpdateAllSoldTemporary()
		assert.NoError(err)
	}
	{
		// failed
		ipi := NewWatchInteractor(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().UpdateAllSoldTemporary().Return(fmt.Errorf("error"))
		err := ipi.UpdateAllSoldTemporary()
		assert.Error(err)
	}
}

func TestWatchInteractor_RemoveWatch(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockWr := database.NewMockWatchRepository(ctrl)
	id := int64(1)
	{
		// success
		ipi := NewWatchInteractor(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().RemoveWatch(id).Return(nil)
		err := ipi.RemoveWatch(id)
		assert.NoError(err)
	}
	{
		// failed
		ipi := NewWatchInteractor(mockWr)
		if ipi == nil {
			t.FailNow()
		}
		mockWr.EXPECT().RemoveWatch(id).Return(fmt.Errorf("error"))
		err := ipi.RemoveWatch(id)
		assert.Error(err)
	}
}
