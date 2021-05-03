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

func TestNewIPadInteractor(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	{
		// success
		ipi := NewIPadInteractor(mockIpr)
		assert.NotNil(ipi)
	}
	{
		// failed because database is nil
		ipi := NewIPadInteractor(nil)
		assert.Nil(ipi)
	}
}

func TestIPadInteractor_FindIPad(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	expected := make(model.IPads, 1)
	expected[0] = model.IPad{Name: "IPad Pro"}
	{
		// success
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindIPad(&model.IPadRequestParam{}).Return(expected, nil)
		actual, err := ipi.FindIPad(&model.IPadRequestParam{})
		assert.NotNil(actual)
		assert.NoError(err)
		assert.Equal(expected, actual)
	}
	{
		// failed
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindIPad(&model.IPadRequestParam{}).Return(nil, fmt.Errorf("error"))
		actual, err := ipi.FindIPad(&model.IPadRequestParam{})
		assert.Nil(actual)
		assert.Error(err)
	}
}

func TestIPadInteractor_FindIPadAll(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	var expected model.IPads = make(model.IPads, 1)
	expected[0] = model.IPad{}
	{
		// success
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindIPadAll().Return(expected, nil)
		actual, err := ipi.FindIPadAll()
		assert.NotNil(actual)
		assert.NoError(err)
		assert.Equal(expected, actual)
	}
	{
		// failed
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindIPadAll().Return(nil, fmt.Errorf("error"))
		actual, err := ipi.FindIPadAll()
		assert.Nil(actual)
		assert.Error(err)
	}
}

func TestIPadInteractor_FindByURL(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	expected := &model.IPad{}
	url := "https://apple.com"
	{
		// success
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindByURL(url).Return(expected, nil)
		actual, err := ipi.FindByURL(url)
		assert.NotNil(actual)
		assert.NoError(err)
		assert.Equal(expected, actual)
	}
	{
		// failed
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().FindByURL(url).Return(nil, fmt.Errorf("error"))
		actual, err := ipi.FindByURL(url)
		assert.Nil(actual)
		assert.Error(err)
	}
}

func TestIPadInteractor_IsExist(t *testing.T) {
	assert := assert.New(t)
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
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().IsExist(input).Return(eIsExist, eID, eT, nil)
		aIsExist, aID, aT, err := ipi.IsExist(input)
		assert.NoError(err)
		assert.Equal(eIsExist, aIsExist)
		assert.Equal(eID, aID)
		assert.Equal(eT, aT)
	}
	{
		// failed
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().IsExist(input).Return(false, uint(0), time.Time{}, fmt.Errorf("error"))
		aIsExist, aID, aT, err := ipi.IsExist(input)
		assert.Error(err)
		assert.False(aIsExist)
		assert.Equal(uint(0), aID)
		assert.Equal(time.Time{}, aT)
	}
}

func TestIPadInteractor_AddIPad(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	input := &model.IPad{}
	{
		// success
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().AddIPad(input).Return(nil)
		err := ipi.AddIPad(input)
		assert.NoError(err)
	}
	{
		// failed
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().AddIPad(input).Return(fmt.Errorf("error"))
		err := ipi.AddIPad(input)
		assert.Error(err)
	}
}

func TestIPadInteractor_UpdateIPad(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	input := &model.IPad{ID: 1}
	failedInput := &model.IPad{ID: 0}
	{
		// success
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().UpdateIPad(input).Return(nil)
		err := ipi.UpdateIPad(input)
		assert.NoError(err)
	}
	{
		// failed
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().UpdateIPad(failedInput).Times(0)
		err := ipi.UpdateIPad(failedInput)
		assert.Error(err, fmt.Errorf("cannot update ipad because invalid ipad id: %d", 0))
	}
}

func TestIPadInteractor_UpdateAllSoldTemporary(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	{
		// success
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().UpdateAllSoldTemporary().Return(nil)
		err := ipi.UpdateAllSoldTemporary()
		assert.NoError(err)
	}
	{
		// failed
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().UpdateAllSoldTemporary().Return(fmt.Errorf("error"))
		err := ipi.UpdateAllSoldTemporary()
		assert.Error(err)
	}
}

func TestIPadInteractor_RemoveIPad(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockIpr := database.NewMockIPadRepository(ctrl)
	id := int64(1)
	{
		// success
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().RemoveIPad(id).Return(nil)
		err := ipi.RemoveIPad(id)
		assert.NoError(err)
	}
	{
		// failed
		ipi := NewIPadInteractor(mockIpr)
		if ipi == nil {
			t.FailNow()
		}
		mockIpr.EXPECT().RemoveIPad(id).Return(fmt.Errorf("error"))
		err := ipi.RemoveIPad(id)
		assert.Error(err)
	}
}
