// Code generated by MockGen. DO NOT EDIT.
// Source: ipad_usecase.go

// Package usecase is a generated GoMock package.
package usecase

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	model "github.com/s14t284/apple-maitained-bot/domain/model"
)

// MockIPadUseCase is a mock of IPadUseCase interface.
type MockIPadUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockIPadUseCaseMockRecorder
}

// MockIPadUseCaseMockRecorder is the mock recorder for MockIPadUseCase.
type MockIPadUseCaseMockRecorder struct {
	mock *MockIPadUseCase
}

// NewMockIPadUseCase creates a new mock instance.
func NewMockIPadUseCase(ctrl *gomock.Controller) *MockIPadUseCase {
	mock := &MockIPadUseCase{ctrl: ctrl}
	mock.recorder = &MockIPadUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIPadUseCase) EXPECT() *MockIPadUseCaseMockRecorder {
	return m.recorder
}

// GetIPads mocks base method.
func (m *MockIPadUseCase) GetIPads(irp model.IPadRequestParam) (model.IPads, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIPads", irp)
	ret0, _ := ret[0].(model.IPads)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIPads indicates an expected call of GetIPads.
func (mr *MockIPadUseCaseMockRecorder) GetIPads(irp interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIPads", reflect.TypeOf((*MockIPadUseCase)(nil).GetIPads), irp)
}
