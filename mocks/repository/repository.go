// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	context "context"

	models "github.com/keremdokumaci/goql/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository[T models.Modeler] struct {
	mock.Mock
}

// Get provides a mock function with given fields: ctx, ID
func (_m *Repository[T]) Get(ctx context.Context, ID int) (T, error) {
	ret := _m.Called(ctx, ID)

	var r0 T
	if rf, ok := ret.Get(0).(func(context.Context, int) T); ok {
		r0 = rf(ctx, ID)
	} else {
		r0 = ret.Get(0).(T)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUniqueField provides a mock function with given fields: ctx, field, value
func (_m *Repository[T]) GetByUniqueField(ctx context.Context, field string, value interface{}) (T, error) {
	ret := _m.Called(ctx, field, value)

	var r0 T
	if rf, ok := ret.Get(0).(func(context.Context, string, interface{}) T); ok {
		r0 = rf(ctx, field, value)
	} else {
		r0 = ret.Get(0).(T)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, interface{}) error); ok {
		r1 = rf(ctx, field, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository[T models.Modeler](t mockConstructorTestingTNewRepository) *Repository[T] {
	mock := &Repository[T]{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
