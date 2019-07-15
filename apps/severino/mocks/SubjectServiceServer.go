// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import proto "backend/proto"

// SubjectServiceServer is an autogenerated mock type for the SubjectServiceServer type
type SubjectServiceServer struct {
	mock.Mock
}

// Create provides a mock function with given fields: _a0, _a1
func (_m *SubjectServiceServer) Create(_a0 context.Context, _a1 *proto.SubjectRequest) (*proto.SubjectResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.SubjectResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.SubjectRequest) *proto.SubjectResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.SubjectResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.SubjectRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: _a0, _a1
func (_m *SubjectServiceServer) Delete(_a0 context.Context, _a1 *proto.SubjectRequest) (*proto.SubjectResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.SubjectResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.SubjectRequest) *proto.SubjectResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.SubjectResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.SubjectRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadAll provides a mock function with given fields: _a0, _a1
func (_m *SubjectServiceServer) ReadAll(_a0 context.Context, _a1 *proto.SearchRequest) (*proto.SubjectsResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.SubjectsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.SearchRequest) *proto.SubjectsResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.SubjectsResponse)
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
func (_m *SubjectServiceServer) ReadOne(_a0 context.Context, _a1 *proto.SubjectRequest) (*proto.SubjectResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.SubjectResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.SubjectRequest) *proto.SubjectResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.SubjectResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.SubjectRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: _a0, _a1
func (_m *SubjectServiceServer) Update(_a0 context.Context, _a1 *proto.SubjectRequest) (*proto.SubjectResponse, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *proto.SubjectResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.SubjectRequest) *proto.SubjectResponse); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.SubjectResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.SubjectRequest) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
