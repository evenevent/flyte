// Code generated by mockery v2.40.3. DO NOT EDIT.

package mocks

import (
	context "context"

	admin "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/admin"

	mock "github.com/stretchr/testify/mock"
)

// MetricsInterface is an autogenerated mock type for the MetricsInterface type
type MetricsInterface struct {
	mock.Mock
}

type MetricsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MetricsInterface) EXPECT() *MetricsInterface_Expecter {
	return &MetricsInterface_Expecter{mock: &_m.Mock}
}

// GetExecutionMetrics provides a mock function with given fields: ctx, request
func (_m *MetricsInterface) GetExecutionMetrics(ctx context.Context, request *admin.WorkflowExecutionGetMetricsRequest) (*admin.WorkflowExecutionGetMetricsResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetExecutionMetrics")
	}

	var r0 *admin.WorkflowExecutionGetMetricsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *admin.WorkflowExecutionGetMetricsRequest) (*admin.WorkflowExecutionGetMetricsResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *admin.WorkflowExecutionGetMetricsRequest) *admin.WorkflowExecutionGetMetricsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*admin.WorkflowExecutionGetMetricsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *admin.WorkflowExecutionGetMetricsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MetricsInterface_GetExecutionMetrics_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetExecutionMetrics'
type MetricsInterface_GetExecutionMetrics_Call struct {
	*mock.Call
}

// GetExecutionMetrics is a helper method to define mock.On call
//   - ctx context.Context
//   - request *admin.WorkflowExecutionGetMetricsRequest
func (_e *MetricsInterface_Expecter) GetExecutionMetrics(ctx interface{}, request interface{}) *MetricsInterface_GetExecutionMetrics_Call {
	return &MetricsInterface_GetExecutionMetrics_Call{Call: _e.mock.On("GetExecutionMetrics", ctx, request)}
}

func (_c *MetricsInterface_GetExecutionMetrics_Call) Run(run func(ctx context.Context, request *admin.WorkflowExecutionGetMetricsRequest)) *MetricsInterface_GetExecutionMetrics_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*admin.WorkflowExecutionGetMetricsRequest))
	})
	return _c
}

func (_c *MetricsInterface_GetExecutionMetrics_Call) Return(_a0 *admin.WorkflowExecutionGetMetricsResponse, _a1 error) *MetricsInterface_GetExecutionMetrics_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MetricsInterface_GetExecutionMetrics_Call) RunAndReturn(run func(context.Context, *admin.WorkflowExecutionGetMetricsRequest) (*admin.WorkflowExecutionGetMetricsResponse, error)) *MetricsInterface_GetExecutionMetrics_Call {
	_c.Call.Return(run)
	return _c
}

// NewMetricsInterface creates a new instance of MetricsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMetricsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MetricsInterface {
	mock := &MetricsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
