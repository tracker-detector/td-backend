// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/Tracking-Detector/td-backend/go/td_common/model"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx
func (_m *UserRepository) Count(ctx context.Context) (int64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Count")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (int64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAll provides a mock function with given fields: ctx
func (_m *UserRepository) DeleteAll(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAll")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByID provides a mock function with given fields: ctx, id
func (_m *UserRepository) DeleteByID(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindAll provides a mock function with given fields: ctx
func (_m *UserRepository) FindAll(ctx context.Context) ([]*model.UserData, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []*model.UserData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*model.UserData, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*model.UserData); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepository) FindByEmail(ctx context.Context, email string) (*model.UserData, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for FindByEmail")
	}

	var r0 *model.UserData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.UserData, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.UserData); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *UserRepository) FindByID(ctx context.Context, id string) (*model.UserData, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *model.UserData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.UserData, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.UserData); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InTransaction provides a mock function with given fields: ctx, fn
func (_m *UserRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
	ret := _m.Called(ctx, fn)

	if len(ret) == 0 {
		panic("no return value specified for InTransaction")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, func(context.Context) error) error); ok {
		r0 = rf(ctx, fn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Save provides a mock function with given fields: ctx, entity
func (_m *UserRepository) Save(ctx context.Context, entity *model.UserData) (*model.UserData, error) {
	ret := _m.Called(ctx, entity)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 *model.UserData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserData) (*model.UserData, error)); ok {
		return rf(ctx, entity)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserData) *model.UserData); ok {
		r0 = rf(ctx, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.UserData) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveAll provides a mock function with given fields: ctx, entity
func (_m *UserRepository) SaveAll(ctx context.Context, entity []*model.UserData) ([]*model.UserData, error) {
	ret := _m.Called(ctx, entity)

	if len(ret) == 0 {
		panic("no return value specified for SaveAll")
	}

	var r0 []*model.UserData
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*model.UserData) ([]*model.UserData, error)); ok {
		return rf(ctx, entity)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*model.UserData) []*model.UserData); ok {
		r0 = rf(ctx, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.UserData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*model.UserData) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StreamAll provides a mock function with given fields: ctx
func (_m *UserRepository) StreamAll(ctx context.Context) (<-chan *model.UserData, <-chan error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for StreamAll")
	}

	var r0 <-chan *model.UserData
	var r1 <-chan error
	if rf, ok := ret.Get(0).(func(context.Context) (<-chan *model.UserData, <-chan error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) <-chan *model.UserData); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *model.UserData)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) <-chan error); ok {
		r1 = rf(ctx)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(<-chan error)
		}
	}

	return r0, r1
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
