// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	models "github.com/keremdokumaci/goql/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// WhitelistRepository is an autogenerated mock type for the WhitelistRepository type
type WhitelistRepository struct {
	mock.Mock
}

// GetWhitelistByQueryName provides a mock function with given fields: queryName
func (_m *WhitelistRepository) GetWhitelistByQueryName(queryName string) (*models.Whitelist, error) {
	ret := _m.Called(queryName)

	var r0 *models.Whitelist
	if rf, ok := ret.Get(0).(func(string) *models.Whitelist); ok {
		r0 = rf(queryName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Whitelist)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(queryName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewWhitelistRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewWhitelistRepository creates a new instance of WhitelistRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewWhitelistRepository(t mockConstructorTestingTNewWhitelistRepository) *WhitelistRepository {
	mock := &WhitelistRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
