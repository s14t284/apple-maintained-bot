// Code generated by MockGen. DO NOT EDIT.
// Source: crawler_usecase.go

// Package usecase is a generated GoMock package.
package usecase

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCrawlerUseCase is a mock of CrawlerUseCase interface.
type MockCrawlerUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockCrawlerUseCaseMockRecorder
}

// MockCrawlerUseCaseMockRecorder is the mock recorder for MockCrawlerUseCase.
type MockCrawlerUseCaseMockRecorder struct {
	mock *MockCrawlerUseCase
}

// NewMockCrawlerUseCase creates a new mock instance.
func NewMockCrawlerUseCase(ctrl *gomock.Controller) *MockCrawlerUseCase {
	mock := &MockCrawlerUseCase{ctrl: ctrl}
	mock.recorder = &MockCrawlerUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCrawlerUseCase) EXPECT() *MockCrawlerUseCaseMockRecorder {
	return m.recorder
}

// CrawlIPadPage mocks base method.
func (m *MockCrawlerUseCase) CrawlIPadPage() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrawlIPadPage")
	ret0, _ := ret[0].(error)
	return ret0
}

// CrawlIPadPage indicates an expected call of CrawlIPadPage.
func (mr *MockCrawlerUseCaseMockRecorder) CrawlIPadPage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrawlIPadPage", reflect.TypeOf((*MockCrawlerUseCase)(nil).CrawlIPadPage))
}

// CrawlMacPage mocks base method.
func (m *MockCrawlerUseCase) CrawlMacPage() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrawlMacPage")
	ret0, _ := ret[0].(error)
	return ret0
}

// CrawlMacPage indicates an expected call of CrawlMacPage.
func (mr *MockCrawlerUseCaseMockRecorder) CrawlMacPage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrawlMacPage", reflect.TypeOf((*MockCrawlerUseCase)(nil).CrawlMacPage))
}

// CrawlWatchPage mocks base method.
func (m *MockCrawlerUseCase) CrawlWatchPage() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CrawlWatchPage")
	ret0, _ := ret[0].(error)
	return ret0
}

// CrawlWatchPage indicates an expected call of CrawlWatchPage.
func (mr *MockCrawlerUseCaseMockRecorder) CrawlWatchPage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CrawlWatchPage", reflect.TypeOf((*MockCrawlerUseCase)(nil).CrawlWatchPage))
}
