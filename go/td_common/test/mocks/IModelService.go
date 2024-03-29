// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/Tracking-Detector/td-backend/go/td_common/model"
	mock "github.com/stretchr/testify/mock"
)

// IModelService is an autogenerated mock type for the IModelService type
type IModelService struct {
	mock.Mock
}

// DeleteModelByID provides a mock function with given fields: ctx, id
func (_m *IModelService) DeleteModelByID(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteModelByID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllModels provides a mock function with given fields: ctx
func (_m *IModelService) GetAllModels(ctx context.Context) ([]*model.Model, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetAllModels")
	}

	var r0 []*model.Model
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*model.Model, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*model.Model); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Model)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetModelById provides a mock function with given fields: ctx, id
func (_m *IModelService) GetModelById(ctx context.Context, id string) (*model.Model, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetModelById")
	}

	var r0 *model.Model
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Model, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Model); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Model)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetModelByName provides a mock function with given fields: ctx, name
func (_m *IModelService) GetModelByName(ctx context.Context, name string) (*model.Model, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetModelByName")
	}

	var r0 *model.Model
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Model, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Model); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Model)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: ctx, _a1
func (_m *IModelService) Save(ctx context.Context, _a1 *model.Model) (*model.Model, error) {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 *model.Model
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Model) (*model.Model, error)); ok {
		return rf(ctx, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Model) *model.Model); ok {
		r0 = rf(ctx, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Model)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Model) error); ok {
		r1 = rf(ctx, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewIModelService creates a new instance of IModelService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIModelService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IModelService {
	mock := &IModelService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
