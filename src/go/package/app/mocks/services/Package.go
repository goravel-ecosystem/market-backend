// Code generated by mockery v2.34.2. DO NOT EDIT.

package mocks

import (
	base "market.goravel.dev/proto/base"
	_package "market.goravel.dev/proto/package"

	context "context"

	mock "github.com/stretchr/testify/mock"

	models "market.goravel.dev/package/app/models"
)

// Package is an autogenerated mock type for the Package type
type Package struct {
	mock.Mock
}

// CreatePackage provides a mock function with given fields: req
func (_m *Package) CreatePackage(req *_package.CreatePackageRequest) (*models.Package, error) {
	ret := _m.Called(req)

	var r0 *models.Package
	var r1 error
	if rf, ok := ret.Get(0).(func(*_package.CreatePackageRequest) (*models.Package, error)); ok {
		return rf(req)
	}
	if rf, ok := ret.Get(0).(func(*_package.CreatePackageRequest) *models.Package); ok {
		r0 = rf(req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Package)
		}
	}

	if rf, ok := ret.Get(1).(func(*_package.CreatePackageRequest) error); ok {
		r1 = rf(req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPackageByID provides a mock function with given fields: id
func (_m *Package) GetPackageByID(id string) (*models.Package, error) {
	ret := _m.Called(id)

	var r0 *models.Package
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*models.Package, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *models.Package); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Package)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetPackages provides a mock function with given fields: query, pagination
func (_m *Package) GetPackages(query *_package.PackagesQuery, pagination *base.Pagination) ([]*models.Package, int64, error) {
	ret := _m.Called(query, pagination)

	var r0 []*models.Package
	var r1 int64
	var r2 error
	if rf, ok := ret.Get(0).(func(*_package.PackagesQuery, *base.Pagination) ([]*models.Package, int64, error)); ok {
		return rf(query, pagination)
	}
	if rf, ok := ret.Get(0).(func(*_package.PackagesQuery, *base.Pagination) []*models.Package); ok {
		r0 = rf(query, pagination)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Package)
		}
	}

	if rf, ok := ret.Get(1).(func(*_package.PackagesQuery, *base.Pagination) int64); ok {
		r1 = rf(query, pagination)
	} else {
		r1 = ret.Get(1).(int64)
	}

	if rf, ok := ret.Get(2).(func(*_package.PackagesQuery, *base.Pagination) error); ok {
		r2 = rf(query, pagination)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// UpdatePackage provides a mock function with given fields: ctx, req
func (_m *Package) UpdatePackage(ctx context.Context, req *_package.UpdatePackageRequest) (*models.Package, error) {
	ret := _m.Called(ctx, req)

	var r0 *models.Package
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *_package.UpdatePackageRequest) (*models.Package, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *_package.UpdatePackageRequest) *models.Package); ok {
		r0 = rf(ctx, req)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Package)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *_package.UpdatePackageRequest) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewPackage creates a new instance of Package. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewPackage(t interface {
	mock.TestingT
	Cleanup(func())
}) *Package {
	mock := &Package{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
