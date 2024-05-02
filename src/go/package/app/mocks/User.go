// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	user "market.goravel.dev/proto/user"
)

// User is an autogenerated mock type for the User type
type User struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: ctx, userID
func (_m *User) GetUser(ctx context.Context, userID uint64) (*user.User, error) {
	ret := _m.Called(ctx, userID)

	var r0 *user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, uint64) (*user.User, error)); ok {
		return rf(ctx, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *user.User); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUsers provides a mock function with given fields: ctx, userIDs
func (_m *User) GetUsers(ctx context.Context, userIDs []string) ([]*user.User, error) {
	ret := _m.Called(ctx, userIDs)

	var r0 []*user.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, []string) ([]*user.User, error)); ok {
		return rf(ctx, userIDs)
	}
	if rf, ok := ret.Get(0).(func(context.Context, []string) []*user.User); ok {
		r0 = rf(ctx, userIDs)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*user.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, []string) error); ok {
		r1 = rf(ctx, userIDs)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUser creates a new instance of User. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUser(t interface {
	mock.TestingT
	Cleanup(func())
}) *User {
	mock := &User{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
