// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// WhiteLister is an autogenerated mock type for the WhiteLister type
type WhiteLister struct {
	mock.Mock
}

// OperationAllowed provides a mock function with given fields: ctx, operationName
func (_m *WhiteLister) OperationAllowed(ctx context.Context, operationName string) bool {
	ret := _m.Called(ctx, operationName)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, operationName)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

type mockConstructorTestingTNewWhiteLister interface {
	mock.TestingT
	Cleanup(func())
}

// NewWhiteLister creates a new instance of WhiteLister. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWhiteLister(t mockConstructorTestingTNewWhiteLister) *WhiteLister {
	mock := &WhiteLister{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}