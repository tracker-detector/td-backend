// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	context "context"

	model "github.com/Tracking-Detector/td-backend/go/td_common/model"
	mock "github.com/stretchr/testify/mock"
)

// ExportRunRepository is an autogenerated mock type for the ExportRunRepository type
type ExportRunRepository struct {
	mock.Mock
}

// Count provides a mock function with given fields: ctx
func (_m *ExportRunRepository) Count(ctx context.Context) (int64, error) {
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

// CountByExporterID provides a mock function with given fields: ctx, exporterId
func (_m *ExportRunRepository) CountByExporterID(ctx context.Context, exporterId string) (int64, error) {
	ret := _m.Called(ctx, exporterId)

	if len(ret) == 0 {
		panic("no return value specified for CountByExporterID")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (int64, error)); ok {
		return rf(ctx, exporterId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) int64); ok {
		r0 = rf(ctx, exporterId)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, exporterId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteAll provides a mock function with given fields: ctx
func (_m *ExportRunRepository) DeleteAll(ctx context.Context) error {
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

// DeleteAllByExporterID provides a mock function with given fields: ctx, id
func (_m *ExportRunRepository) DeleteAllByExporterID(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteAllByExporterID")
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
func (_m *ExportRunRepository) DeleteByID(ctx context.Context, id string) error {
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

// ExistByExporterIDAndRecducer provides a mock function with given fields: ctx, exporterId, reducer
func (_m *ExportRunRepository) ExistByExporterIDAndRecducer(ctx context.Context, exporterId string, reducer string) (bool, error) {
	ret := _m.Called(ctx, exporterId, reducer)

	if len(ret) == 0 {
		panic("no return value specified for ExistByExporterIDAndRecducer")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) (bool, error)); ok {
		return rf(ctx, exporterId, reducer)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string) bool); ok {
		r0 = rf(ctx, exporterId, reducer)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, exporterId, reducer)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAll provides a mock function with given fields: ctx
func (_m *ExportRunRepository) FindAll(ctx context.Context) ([]*model.ExportRun, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for FindAll")
	}

	var r0 []*model.ExportRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]*model.ExportRun, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []*model.ExportRun); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ExportRun)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByExporterID provides a mock function with given fields: ctx, exporterId
func (_m *ExportRunRepository) FindByExporterID(ctx context.Context, exporterId string) ([]*model.ExportRun, error) {
	ret := _m.Called(ctx, exporterId)

	if len(ret) == 0 {
		panic("no return value specified for FindByExporterID")
	}

	var r0 []*model.ExportRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) ([]*model.ExportRun, error)); ok {
		return rf(ctx, exporterId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) []*model.ExportRun); ok {
		r0 = rf(ctx, exporterId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ExportRun)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, exporterId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByID provides a mock function with given fields: ctx, id
func (_m *ExportRunRepository) FindByID(ctx context.Context, id string) (*model.ExportRun, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *model.ExportRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.ExportRun, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.ExportRun); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ExportRun)
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
func (_m *ExportRunRepository) InTransaction(ctx context.Context, fn func(context.Context) error) error {
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
func (_m *ExportRunRepository) Save(ctx context.Context, entity *model.ExportRun) (*model.ExportRun, error) {
	ret := _m.Called(ctx, entity)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 *model.ExportRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.ExportRun) (*model.ExportRun, error)); ok {
		return rf(ctx, entity)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.ExportRun) *model.ExportRun); ok {
		r0 = rf(ctx, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.ExportRun)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.ExportRun) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SaveAll provides a mock function with given fields: ctx, entity
func (_m *ExportRunRepository) SaveAll(ctx context.Context, entity []*model.ExportRun) ([]*model.ExportRun, error) {
	ret := _m.Called(ctx, entity)

	if len(ret) == 0 {
		panic("no return value specified for SaveAll")
	}

	var r0 []*model.ExportRun
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []*model.ExportRun) ([]*model.ExportRun, error)); ok {
		return rf(ctx, entity)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []*model.ExportRun) []*model.ExportRun); ok {
		r0 = rf(ctx, entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.ExportRun)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []*model.ExportRun) error); ok {
		r1 = rf(ctx, entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StreamAll provides a mock function with given fields: ctx
func (_m *ExportRunRepository) StreamAll(ctx context.Context) (<-chan *model.ExportRun, <-chan error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for StreamAll")
	}

	var r0 <-chan *model.ExportRun
	var r1 <-chan error
	if rf, ok := ret.Get(0).(func(context.Context) (<-chan *model.ExportRun, <-chan error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) <-chan *model.ExportRun); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *model.ExportRun)
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

// NewExportRunRepository creates a new instance of ExportRunRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewExportRunRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *ExportRunRepository {
	mock := &ExportRunRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
