// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"

	users "github.com/binodluitel/api/pkg/models/users"
)

// UsersService is an autogenerated mock type for the UsersService type
type UsersService struct {
	mock.Mock
}

// CreateUser provides a mock function with given fields: ctx, request
func (_m *UsersService) CreateUser(ctx *gin.Context, request *users.CreateRequest) (*users.User, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context, *users.CreateRequest) (*users.User, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context, *users.CreateRequest) *users.User); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*gin.Context, *users.CreateRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUser provides a mock function with given fields: ctx, id
func (_m *UsersService) DeleteUser(ctx *gin.Context, id string) (*users.User, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context, string) (*users.User, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context, string) *users.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*gin.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUser provides a mock function with given fields: ctx, id, filters
func (_m *UsersService) GetUser(ctx *gin.Context, id string, filters string) (*users.User, error) {
	ret := _m.Called(ctx, id, filters)

	if len(ret) == 0 {
		panic("no return value specified for GetUser")
	}

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context, string, string) (*users.User, error)); ok {
		return rf(ctx, id, filters)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context, string, string) *users.User); ok {
		r0 = rf(ctx, id, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*gin.Context, string, string) error); ok {
		r1 = rf(ctx, id, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUsers provides a mock function with given fields: ctx, filters
func (_m *UsersService) ListUsers(ctx *gin.Context, filters string) ([]*users.User, error) {
	ret := _m.Called(ctx, filters)

	if len(ret) == 0 {
		panic("no return value specified for ListUsers")
	}

	var r0 []*users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context, string) ([]*users.User, error)); ok {
		return rf(ctx, filters)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context, string) []*users.User); ok {
		r0 = rf(ctx, filters)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*gin.Context, string) error); ok {
		r1 = rf(ctx, filters)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateUser provides a mock function with given fields: ctx, id, request
func (_m *UsersService) UpdateUser(ctx *gin.Context, id string, request *users.UpdateRequest) (*users.User, error) {
	ret := _m.Called(ctx, id, request)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 *users.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*gin.Context, string, *users.UpdateRequest) (*users.User, error)); ok {
		return rf(ctx, id, request)
	}
	if rf, ok := ret.Get(0).(func(*gin.Context, string, *users.UpdateRequest) *users.User); ok {
		r0 = rf(ctx, id, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*users.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*gin.Context, string, *users.UpdateRequest) error); ok {
		r1 = rf(ctx, id, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUsersService creates a new instance of UsersService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUsersService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UsersService {
	mock := &UsersService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
