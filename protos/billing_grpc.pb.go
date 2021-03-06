// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protos

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BillingClient is the client API for Billing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BillingClient interface {
	BillUser(ctx context.Context, in *BillRequest, opts ...grpc.CallOption) (*BillResponse, error)
	UserBalance(ctx context.Context, in *BillableEntity, opts ...grpc.CallOption) (*BillResponse, error)
	FindOrCreateBillingEntity(ctx context.Context, in *BillableEntity, opts ...grpc.CallOption) (*BillResponse, error)
	CalculateMessageRate(ctx context.Context, in *BillRequest, opts ...grpc.CallOption) (*BillResponse, error)
	RefundUser(ctx context.Context, in *RefundRequest, opts ...grpc.CallOption) (*BillResponse, error)
}

type billingClient struct {
	cc grpc.ClientConnInterface
}

func NewBillingClient(cc grpc.ClientConnInterface) BillingClient {
	return &billingClient{cc}
}

func (c *billingClient) BillUser(ctx context.Context, in *BillRequest, opts ...grpc.CallOption) (*BillResponse, error) {
	out := new(BillResponse)
	err := c.cc.Invoke(ctx, "/grpcpoc.Billing/BillUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) UserBalance(ctx context.Context, in *BillableEntity, opts ...grpc.CallOption) (*BillResponse, error) {
	out := new(BillResponse)
	err := c.cc.Invoke(ctx, "/grpcpoc.Billing/UserBalance", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) FindOrCreateBillingEntity(ctx context.Context, in *BillableEntity, opts ...grpc.CallOption) (*BillResponse, error) {
	out := new(BillResponse)
	err := c.cc.Invoke(ctx, "/grpcpoc.Billing/FindOrCreateBillingEntity", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) CalculateMessageRate(ctx context.Context, in *BillRequest, opts ...grpc.CallOption) (*BillResponse, error) {
	out := new(BillResponse)
	err := c.cc.Invoke(ctx, "/grpcpoc.Billing/CalculateMessageRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *billingClient) RefundUser(ctx context.Context, in *RefundRequest, opts ...grpc.CallOption) (*BillResponse, error) {
	out := new(BillResponse)
	err := c.cc.Invoke(ctx, "/grpcpoc.Billing/RefundUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BillingServer is the server API for Billing service.
// All implementations must embed UnimplementedBillingServer
// for forward compatibility
type BillingServer interface {
	BillUser(context.Context, *BillRequest) (*BillResponse, error)
	UserBalance(context.Context, *BillableEntity) (*BillResponse, error)
	FindOrCreateBillingEntity(context.Context, *BillableEntity) (*BillResponse, error)
	CalculateMessageRate(context.Context, *BillRequest) (*BillResponse, error)
	RefundUser(context.Context, *RefundRequest) (*BillResponse, error)
	mustEmbedUnimplementedBillingServer()
}

// UnimplementedBillingServer must be embedded to have forward compatible implementations.
type UnimplementedBillingServer struct {
}

func (UnimplementedBillingServer) BillUser(context.Context, *BillRequest) (*BillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BillUser not implemented")
}
func (UnimplementedBillingServer) UserBalance(context.Context, *BillableEntity) (*BillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UserBalance not implemented")
}
func (UnimplementedBillingServer) FindOrCreateBillingEntity(context.Context, *BillableEntity) (*BillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindOrCreateBillingEntity not implemented")
}
func (UnimplementedBillingServer) CalculateMessageRate(context.Context, *BillRequest) (*BillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CalculateMessageRate not implemented")
}
func (UnimplementedBillingServer) RefundUser(context.Context, *RefundRequest) (*BillResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefundUser not implemented")
}
func (UnimplementedBillingServer) mustEmbedUnimplementedBillingServer() {}

// UnsafeBillingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BillingServer will
// result in compilation errors.
type UnsafeBillingServer interface {
	mustEmbedUnimplementedBillingServer()
}

func RegisterBillingServer(s grpc.ServiceRegistrar, srv BillingServer) {
	s.RegisterService(&Billing_ServiceDesc, srv)
}

func _Billing_BillUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).BillUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpoc.Billing/BillUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).BillUser(ctx, req.(*BillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_UserBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BillableEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).UserBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpoc.Billing/UserBalance",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).UserBalance(ctx, req.(*BillableEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_FindOrCreateBillingEntity_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BillableEntity)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).FindOrCreateBillingEntity(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpoc.Billing/FindOrCreateBillingEntity",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).FindOrCreateBillingEntity(ctx, req.(*BillableEntity))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_CalculateMessageRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BillRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).CalculateMessageRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpoc.Billing/CalculateMessageRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).CalculateMessageRate(ctx, req.(*BillRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Billing_RefundUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BillingServer).RefundUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpcpoc.Billing/RefundUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BillingServer).RefundUser(ctx, req.(*RefundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Billing_ServiceDesc is the grpc.ServiceDesc for Billing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Billing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpcpoc.Billing",
	HandlerType: (*BillingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "BillUser",
			Handler:    _Billing_BillUser_Handler,
		},
		{
			MethodName: "UserBalance",
			Handler:    _Billing_UserBalance_Handler,
		},
		{
			MethodName: "FindOrCreateBillingEntity",
			Handler:    _Billing_FindOrCreateBillingEntity_Handler,
		},
		{
			MethodName: "CalculateMessageRate",
			Handler:    _Billing_CalculateMessageRate_Handler,
		},
		{
			MethodName: "RefundUser",
			Handler:    _Billing_RefundUser_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/billing.proto",
}
