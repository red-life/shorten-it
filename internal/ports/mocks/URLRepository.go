// Code generated by mockery v2.32.4. DO NOT EDIT.

package ports

import (
	context "context"

	models "github.com/red-life/shorten-it/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// URLRepository is an autogenerated mock type for the URLRepository type
type URLRepository struct {
	mock.Mock
}

type URLRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *URLRepository) EXPECT() *URLRepository_Expecter {
	return &URLRepository_Expecter{mock: &_m.Mock}
}

// GetKeyByLong provides a mock function with given fields: ctx, longURL
func (_m *URLRepository) GetKeyByLong(ctx context.Context, longURL string) (string, error) {
	ret := _m.Called(ctx, longURL)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, longURL)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, longURL)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, longURL)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URLRepository_GetKeyByLong_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetKeyByLong'
type URLRepository_GetKeyByLong_Call struct {
	*mock.Call
}

// GetKeyByLong is a helper method to define mock.On call
//   - ctx context.Context
//   - longURL string
func (_e *URLRepository_Expecter) GetKeyByLong(ctx interface{}, longURL interface{}) *URLRepository_GetKeyByLong_Call {
	return &URLRepository_GetKeyByLong_Call{Call: _e.mock.On("GetKeyByLong", ctx, longURL)}
}

func (_c *URLRepository_GetKeyByLong_Call) Run(run func(ctx context.Context, longURL string)) *URLRepository_GetKeyByLong_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *URLRepository_GetKeyByLong_Call) Return(_a0 string, _a1 error) *URLRepository_GetKeyByLong_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *URLRepository_GetKeyByLong_Call) RunAndReturn(run func(context.Context, string) (string, error)) *URLRepository_GetKeyByLong_Call {
	_c.Call.Return(run)
	return _c
}

// GetLongByKey provides a mock function with given fields: ctx, key
func (_m *URLRepository) GetLongByKey(ctx context.Context, key string) (string, error) {
	ret := _m.Called(ctx, key)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, key)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, key)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// URLRepository_GetLongByKey_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetLongByKey'
type URLRepository_GetLongByKey_Call struct {
	*mock.Call
}

// GetLongByKey is a helper method to define mock.On call
//   - ctx context.Context
//   - key string
func (_e *URLRepository_Expecter) GetLongByKey(ctx interface{}, key interface{}) *URLRepository_GetLongByKey_Call {
	return &URLRepository_GetLongByKey_Call{Call: _e.mock.On("GetLongByKey", ctx, key)}
}

func (_c *URLRepository_GetLongByKey_Call) Run(run func(ctx context.Context, key string)) *URLRepository_GetLongByKey_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *URLRepository_GetLongByKey_Call) Return(_a0 string, _a1 error) *URLRepository_GetLongByKey_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *URLRepository_GetLongByKey_Call) RunAndReturn(run func(context.Context, string) (string, error)) *URLRepository_GetLongByKey_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ctx, url
func (_m *URLRepository) Save(ctx context.Context, url models.URL) error {
	ret := _m.Called(ctx, url)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.URL) error); ok {
		r0 = rf(ctx, url)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// URLRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type URLRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ctx context.Context
//   - url models.URL
func (_e *URLRepository_Expecter) Save(ctx interface{}, url interface{}) *URLRepository_Save_Call {
	return &URLRepository_Save_Call{Call: _e.mock.On("Save", ctx, url)}
}

func (_c *URLRepository_Save_Call) Run(run func(ctx context.Context, url models.URL)) *URLRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.URL))
	})
	return _c
}

func (_c *URLRepository_Save_Call) Return(_a0 error) *URLRepository_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *URLRepository_Save_Call) RunAndReturn(run func(context.Context, models.URL) error) *URLRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// NewURLRepository creates a new instance of URLRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewURLRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *URLRepository {
	mock := &URLRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}