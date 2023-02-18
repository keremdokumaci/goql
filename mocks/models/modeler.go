// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// Modeler is an autogenerated mock type for the Modeler type
type Modeler struct {
	mock.Mock
}

// TableName provides a mock function with given fields:
func (_m *Modeler) TableName() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

type mockConstructorTestingTNewModeler interface {
	mock.TestingT
	Cleanup(func())
}

// NewModeler creates a new instance of Modeler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewModeler(t mockConstructorTestingTNewModeler) *Modeler {
	mock := &Modeler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}