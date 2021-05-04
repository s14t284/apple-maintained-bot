// Code generated by MockGen. DO NOT EDIT.
// Source: watch_usecase.go

// Package usecase is a generated GoMock package.
package usecase

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/s14t284/apple-maitained-bot/domain/model"
)

// MockWatchUseCase is a mock of WatchUseCase interface.
type MockWatchUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockWatchUseCaseMockRecorder
}

// MockWatchUseCaseMockRecorder is the mock recorder for MockWatchUseCase.
type MockWatchUseCaseMockRecorder struct {
	mock *MockWatchUseCase
}

// NewMockWatchUseCase creates a new mock instance.
func NewMockWatchUseCase(ctrl *gomock.Controller) *MockWatchUseCase {
	mock := &MockWatchUseCase{ctrl: ctrl}
	mock.recorder = &MockWatchUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWatchUseCase) EXPECT() *MockWatchUseCaseMockRecorder {
	return m.recorder
}

// GetWatches mocks base method.
func (m *MockWatchUseCase) GetWatches(wrp model.WatchRequestParam) (model.Watches, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetWatches", wrp)
	ret0, _ := ret[0].(model.Watches)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetWatches indicates an expected call of GetWatches.
func (mr *MockWatchUseCaseMockRecorder) GetWatches(wrp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetWatches", reflect.TypeOf((*MockWatchUseCase)(nil).GetWatches), wrp)
}