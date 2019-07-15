// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import grpc "google.golang.org/grpc"
import mock "github.com/stretchr/testify/mock"
import proto "backend/proto"

// ProductServiceClient is an autogenerated mock type for the ProductServiceClient type
type ProductServiceClient struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, in, opts
func (_m *ProductServiceClient) Create(ctx context.Context, in *proto.ProductRequest, opts ...grpc.CallOption) (*proto.ProductResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.ProductResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ProductRequest, ...grpc.CallOption) *proto.ProductResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ProductResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.ProductRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadAll provides a mock function with given fields: ctx, in, opts
func (_m *ProductServiceClient) ReadAll(ctx context.Context, in *proto.SearchRequest, opts ...grpc.CallOption) (*proto.ProductsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.ProductsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.SearchRequest, ...grpc.CallOption) *proto.ProductsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ProductsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.SearchRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ReadOne provides a mock function with given fields: ctx, in, opts
func (_m *ProductServiceClient) ReadOne(ctx context.Context, in *proto.ProductRequest, opts ...grpc.CallOption) (*proto.ProductResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.ProductResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ProductRequest, ...grpc.CallOption) *proto.ProductResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ProductResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.ProductRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, in, opts
func (_m *ProductServiceClient) Update(ctx context.Context, in *proto.ProductRequest, opts ...grpc.CallOption) (*proto.ProductResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *proto.ProductResponse
	if rf, ok := ret.Get(0).(func(context.Context, *proto.ProductRequest, ...grpc.CallOption) *proto.ProductResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*proto.ProductResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *proto.ProductRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
