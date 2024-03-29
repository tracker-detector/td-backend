// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// BaseModelName is an autogenerated mock type for the BaseModelName type
type BaseModelName struct {
	mock.Mock
}

// GetID provides a mock function with given fields:
func (_m *BaseModelName) GetID() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetID")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetName provides a mock function with given fields:
func (_m *BaseModelName) GetName() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetName")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SetID provides a mock function with given fields: id
func (_m *BaseModelName) SetID(id string) {
	_m.Called(id)
}

// NewBaseModelName creates a new instance of BaseModelName. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBaseModelName(t interface {
	mock.TestingT
	Cleanup(func())
}) *BaseModelName {
	mock := &BaseModelName{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
