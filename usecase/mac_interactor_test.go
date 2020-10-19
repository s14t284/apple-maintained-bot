package usecase

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/s14t284/apple-maitained-bot/domain/model"
	"github.com/s14t284/apple-maitained-bot/mock/repository"
	"github.com/stretchr/testify/assert"
)

func TestNewMacInteractor(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := repository.NewMockMacRepository(ctrl)
	{
		// success
		mpi := NewMacInteractor(mockMr)
		assert.NotNil(mpi)
	}
	{
		// failed because repository is nil
		mpi := NewMacInteractor(nil)
		assert.Nil(mpi)
	}
}

func TestMacInteractor_FindMacAll(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := repository.NewMockMacRepository(ctrl)
	var expected model.Macs = make(model.Macs, 1)
	expected[0] = model.Mac{}
	{
		// success
		mpi := NewMacInteractor(mockMr)
		if mpi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().FindMacAll().Return(expected, nil)
		actual, err := mpi.FindMacAll()
		assert.Equal(expected, actual)
		assert.NoError(err)
	}
	{
		// failed
		mpi := NewMacInteractor(mockMr)
		if mpi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().FindMacAll().Return(nil, fmt.Errorf("error"))
		actual, err := mpi.FindMacAll()
		assert.Nil(actual)
		assert.Error(err)
	}
}

func TestMacInteractor_FindByURL(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := repository.NewMockMacRepository(ctrl)
	expected := &model.Mac{}
	url := "https://apple.com"
	{
		// success
		mpi := NewMacInteractor(mockMr)
		if mpi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().FindByURL(url).Return(expected, nil)
		actual, err := mpi.FindByURL(url)
		assert.Equal(expected, actual)
		assert.NoError(err)
	}
	{
		// failed
		mpi := NewMacInteractor(mockMr)
		if mpi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().FindByURL(url).Return(nil, fmt.Errorf("error"))
		actual, err := mpi.FindByURL(url)
		assert.Nil(actual)
		assert.Error(err)
	}
}

func TestMacInteractor_IsExist(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := repository.NewMockMacRepository(ctrl)
	input := &model.Mac{}

	// mock output
	eIsExist := true
	eID := uint(1)
	eT := time.Now()
	{
		// success
		mpi := NewMacInteractor(mockMr)
		if mpi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().IsExist(input).Return(eIsExist, eID, eT, nil)
		aIsExist, aID, aT, err := mpi.IsExist(input)
		assert.NoError(err)
		assert.Equal(eIsExist, aIsExist)
		assert.Equal(eID, aID)
		assert.Equal(eT, aT)
	}
	{
		// failed
		mpi := NewMacInteractor(mockMr)
		if mpi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().IsExist(input).Return(false, uint(0), time.Time{}, fmt.Errorf("error"))
		aIsExist, aID, aT, err := mpi.IsExist(input)
		assert.Error(err)
		assert.False(aIsExist)
		assert.Equal(uint(0), aID)
		assert.Equal(time.Time{}, aT)
	}
}

func TestMacInteractor_AddMac(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := repository.NewMockMacRepository(ctrl)
	input := &model.Mac{}
	{
		// success
		ipi := NewMacInteractor(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().AddMac(input).Return(nil)
		err := ipi.AddMac(input)
		assert.NoError(err)
	}
	{
		// failed
		ipi := NewMacInteractor(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().AddMac(input).Return(fmt.Errorf("error"))
		err := ipi.AddMac(input)
		assert.Error(err)
	}
}

func TestMacInteractor_UpdateMac(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := repository.NewMockMacRepository(ctrl)
	input := &model.Mac{ID: 1}
	failedInput := &model.Mac{ID: 0}
	{
		// success
		ipi := NewMacInteractor(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().UpdateMac(input).Return(nil)
		err := ipi.UpdateMac(input)
		assert.NoError(err)
	}
	{
		// failed
		ipi := NewMacInteractor(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().UpdateMac(failedInput).Times(0)
		err := ipi.UpdateMac(failedInput)
		assert.Error(err, fmt.Errorf("cannot update mac because invalid mac id: %d", 0))
	}
}

func TestMacInteractor_UpdateAllSoldTemporary(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := repository.NewMockMacRepository(ctrl)
	{
		// success
		ipi := NewMacInteractor(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().UpdateAllSoldTemporary().Return(nil)
		err := ipi.UpdateAllSoldTemporary()
		assert.NoError(err)
	}
	{
		// failed
		ipi := NewMacInteractor(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().UpdateAllSoldTemporary().Return(fmt.Errorf("error"))
		err := ipi.UpdateAllSoldTemporary()
		assert.Error(err)
	}
}

func TestMacInteractor_RemoveMac(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockMr := repository.NewMockMacRepository(ctrl)
	id := int64(1)
	{
		// success
		ipi := NewMacInteractor(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().RemoveMac(id).Return(nil)
		err := ipi.RemoveMac(id)
		assert.NoError(err)
	}
	{
		// failed
		ipi := NewMacInteractor(mockMr)
		if ipi == nil {
			t.FailNow()
		}
		mockMr.EXPECT().RemoveMac(id).Return(fmt.Errorf("error"))
		err := ipi.RemoveMac(id)
		assert.Error(err)
	}
}
