// Code generated by mockery; DO NOT EDIT.
// github.com/vektra/mockery
// template: testify

package service

import (
	"context"

	"github.com/arjunksofficial/kart-challenge/internal/core/serror"
	"github.com/arjunksofficial/kart-challenge/internal/entities/orders/models"
	mock "github.com/stretchr/testify/mock"
)

// NewMockService creates a new instance of MockService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockService {
	mock := &MockService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}

// MockService is an autogenerated mock type for the Service type
type MockService struct {
	mock.Mock
}

type MockService_Expecter struct {
	mock *mock.Mock
}

func (_m *MockService) EXPECT() *MockService_Expecter {
	return &MockService_Expecter{mock: &_m.Mock}
}

// CreateOrder provides a mock function for the type MockService
func (_mock *MockService) CreateOrder(ctx context.Context, req models.CreateOrderRequest) (models.CreateOrderResponse, *serror.ServiceError) {
	ret := _mock.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 models.CreateOrderResponse
	var r1 *serror.ServiceError
	if returnFunc, ok := ret.Get(0).(func(context.Context, models.CreateOrderRequest) (models.CreateOrderResponse, *serror.ServiceError)); ok {
		return returnFunc(ctx, req)
	}
	if returnFunc, ok := ret.Get(0).(func(context.Context, models.CreateOrderRequest) models.CreateOrderResponse); ok {
		r0 = returnFunc(ctx, req)
	} else {
		r0 = ret.Get(0).(models.CreateOrderResponse)
	}
	if returnFunc, ok := ret.Get(1).(func(context.Context, models.CreateOrderRequest) *serror.ServiceError); ok {
		r1 = returnFunc(ctx, req)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*serror.ServiceError)
		}
	}
	return r0, r1
}

// MockService_CreateOrder_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateOrder'
type MockService_CreateOrder_Call struct {
	*mock.Call
}

// CreateOrder is a helper method to define mock.On call
//   - ctx context.Context
//   - req models.CreateOrderRequest
func (_e *MockService_Expecter) CreateOrder(ctx interface{}, req interface{}) *MockService_CreateOrder_Call {
	return &MockService_CreateOrder_Call{Call: _e.mock.On("CreateOrder", ctx, req)}
}

func (_c *MockService_CreateOrder_Call) Run(run func(ctx context.Context, req models.CreateOrderRequest)) *MockService_CreateOrder_Call {
	_c.Call.Run(func(args mock.Arguments) {
		var arg0 context.Context
		if args[0] != nil {
			arg0 = args[0].(context.Context)
		}
		var arg1 models.CreateOrderRequest
		if args[1] != nil {
			arg1 = args[1].(models.CreateOrderRequest)
		}
		run(
			arg0,
			arg1,
		)
	})
	return _c
}

func (_c *MockService_CreateOrder_Call) Return(createOrderResponse models.CreateOrderResponse, serviceError *serror.ServiceError) *MockService_CreateOrder_Call {
	_c.Call.Return(createOrderResponse, serviceError)
	return _c
}

func (_c *MockService_CreateOrder_Call) RunAndReturn(run func(ctx context.Context, req models.CreateOrderRequest) (models.CreateOrderResponse, *serror.ServiceError)) *MockService_CreateOrder_Call {
	_c.Call.Return(run)
	return _c
}
