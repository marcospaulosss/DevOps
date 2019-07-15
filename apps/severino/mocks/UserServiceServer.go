// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import proto "backend/proto"

// UserServiceServer is an autogenerated mock type for the UserServiceServer type
type UserServiceServer struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) Create(_a0 context.Context, _a1 *proto.UserRequest) (*proto.UserResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.UserResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.UserRequest) *proto.UserResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.UserResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.UserRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadAll provides a mock function with given fields: _a0, _a1
func (_m *UserServiceServer) ReadAll(_a0 context.Context, _a1 *proto.SearchRequest) (*proto.UsersResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.UsersResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.SearchRequest) *proto.UsersResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.UsersResponse)
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
