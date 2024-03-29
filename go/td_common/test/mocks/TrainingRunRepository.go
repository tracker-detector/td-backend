// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/Tracking-Detector/td-backend/go/td_common/model"
	mock "github.com/stretchr/testify/mock"
)

// TrainingRunRepository is an autogenerated mock type for the TrainingRunRepository type
type TrainingRunRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx
func (_m *TrainingRunRepository) Count(ctx context.Context) (int64, error) {
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

// CountByModelID provides a mock function with given fields: ctx, modelId
func (_m *TrainingRunRepository) CountByModelID(ctx context.Context, modelId string) (int64, error) {
	ret := _m.Called(ctx, modelId)

	if len(ret) == 0 {
		panic("no return value specified for CountByModelID")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, modelId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, modelId)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, modelId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CountByName provides a mock function with given fields: ctx, modelName
func (_m *TrainingRunRepository) CountByName(ctx context.Context, modelName string) (int64, error) {
	ret := _m.Called(ctx, modelName)

	if len(ret) == 0 {
		panic("no return value specified for CountByName")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, modelName)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, modelName)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, modelName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAll provides a mock function with given fields: ctx
func (_m *TrainingRunRepository) DeleteAll(ctx context.Context) error {
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

// DeleteAllByModelID provides a mock function with given fields: ctx, id
func (_m *TrainingRunRepository) DeleteAllByModelID(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAllByModelID")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteByID provides a mock function with given fields: ctx, id
func (_m *TrainingRunRepository) DeleteByID(ctx context.Context, id string) error {
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
func (_m *TrainingRunRepository) FindAll(ctx context.Context) ([]*model.TrainingRun, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []*model.TrainingRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*model.TrainingRun, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*model.TrainingRun); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TrainingRun)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *TrainingRunRepository) FindByID(ctx context.Context, id string) (*model.TrainingRun, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *model.TrainingRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.TrainingRun, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.TrainingRun); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TrainingRun)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByModelID provides a mock function with given fields: ctx, modelId
func (_m *TrainingRunRepository) FindByModelID(ctx context.Context, modelId string) ([]*model.TrainingRun, error) {
	ret := _m.Called(ctx, modelId)

	if len(ret) == 0 {
		panic("no return value specified for FindByModelID")
	}

	var r0 []*model.TrainingRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*model.TrainingRun, error)); ok {
		return rf(ctx, modelId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.TrainingRun); ok {
		r0 = rf(ctx, modelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TrainingRun)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, modelId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByModelName provides a mock function with given fields: ctx, modelId
func (_m *TrainingRunRepository) FindByModelName(ctx context.Context, modelId string) ([]*model.TrainingRun, error) {
	ret := _m.Called(ctx, modelId)

	if len(ret) == 0 {
		panic("no return value specified for FindByModelName")
	}

	var r0 []*model.TrainingRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*model.TrainingRun, error)); ok {
		return rf(ctx, modelId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.TrainingRun); ok {
		r0 = rf(ctx, modelId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TrainingRun)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, modelId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InTransaction provides a mock function with given fields: ctx, fn
func (_m *TrainingRunRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
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
func (_m *TrainingRunRepository) Save(ctx context.Context, entity *model.TrainingRun) (*model.TrainingRun, error) {
	ret := _m.Called(ctx, entity)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 *model.TrainingRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.TrainingRun) (*model.TrainingRun, error)); ok {
		return rf(ctx, entity)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.TrainingRun) *model.TrainingRun); ok {
		r0 = rf(ctx, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.TrainingRun)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.TrainingRun) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveAll provides a mock function with given fields: ctx, entity
func (_m *TrainingRunRepository) SaveAll(ctx context.Context, entity []*model.TrainingRun) ([]*model.TrainingRun, error) {
	ret := _m.Called(ctx, entity)

	if len(ret) == 0 {
		panic("no return value specified for SaveAll")
	}

	var r0 []*model.TrainingRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*model.TrainingRun) ([]*model.TrainingRun, error)); ok {
		return rf(ctx, entity)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*model.TrainingRun) []*model.TrainingRun); ok {
		r0 = rf(ctx, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.TrainingRun)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*model.TrainingRun) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StreamAll provides a mock function with given fields: ctx
func (_m *TrainingRunRepository) StreamAll(ctx context.Context) (<-chan *model.TrainingRun, <-chan error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for StreamAll")
	}

	var r0 <-chan *model.TrainingRun
	var r1 <-chan error
	if rf, ok := ret.Get(0).(func(context.Context) (<-chan *model.TrainingRun, <-chan error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) <-chan *model.TrainingRun); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *model.TrainingRun)
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

// NewTrainingRunRepository creates a new instance of TrainingRunRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewTrainingRunRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *TrainingRunRepository {
	mock := &TrainingRunRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
