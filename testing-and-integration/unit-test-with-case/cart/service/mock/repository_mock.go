// Code generated by MockGen. DO NOT EDIT.
// Source: unit-test-case-cart/cart/service (interfaces: RepositoryManager)
//
// Generated by this command:
//
//	mockgen -build_flags=--mod=mod -destination=mock/repository_mock.go -package=mock . RepositoryManager
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockRepositoryManager is a mock of RepositoryManager interface.
type MockRepositoryManager struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryManagerMockRecorder
	isgomock struct{}
}

// MockRepositoryManagerMockRecorder is the mock recorder for MockRepositoryManager.
type MockRepositoryManagerMockRecorder struct {
	mock *MockRepositoryManager
}

// NewMockRepositoryManager creates a new mock instance.
func NewMockRepositoryManager(ctrl *gomock.Controller) *MockRepositoryManager {
	mock := &MockRepositoryManager{ctrl: ctrl}
	mock.recorder = &MockRepositoryManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepositoryManager) EXPECT() *MockRepositoryManagerMockRecorder {
	return m.recorder
}

// AddToCart mocks base method.
func (m *MockRepositoryManager) AddToCart(ctx context.Context, userID, productID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToCart", ctx, userID, productID)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToCart indicates an expected call of AddToCart.
func (mr *MockRepositoryManagerMockRecorder) AddToCart(ctx, userID, productID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToCart", reflect.TypeOf((*MockRepositoryManager)(nil).AddToCart), ctx, userID, productID)
}
