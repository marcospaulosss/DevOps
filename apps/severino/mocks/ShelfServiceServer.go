// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import proto "backend/proto"

// ShelfServiceServer is an autogenerated mock type for the ShelfServiceServer type
type ShelfServiceServer struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *ShelfServiceServer) Create(_a0 context.Context, _a1 *proto.ShelfRequest) (*proto.ShelfResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.ShelfResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ShelfRequest) *proto.ShelfResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ShelfResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.ShelfRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *ShelfServiceServer) Delete(_a0 context.Context, _a1 *proto.ShelfRequest) (*proto.ShelfResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.ShelfResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ShelfRequest) *proto.ShelfResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ShelfResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.ShelfRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadAll provides a mock function with given fields: _a0, _a1
func (_m *ShelfServiceServer) ReadAll(_a0 context.Context, _a1 *proto.SearchRequest) (*proto.ShelvesResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.ShelvesResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.SearchRequest) *proto.ShelvesResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ShelvesResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.SearchRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadOne provides a mock function with given fields: _a0, _a1
func (_m *ShelfServiceServer) ReadOne(_a0 context.Context, _a1 *proto.ShelfRequest) (*proto.ShelfResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.ShelfResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ShelfRequest) *proto.ShelfResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ShelfResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.ShelfRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *ShelfServiceServer) Update(_a0 context.Context, _a1 *proto.ShelfRequest) (*proto.ShelfResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.ShelfResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ShelfRequest) *proto.ShelfResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ShelfResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.ShelfRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
