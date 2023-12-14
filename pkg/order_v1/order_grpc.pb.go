// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: order.proto

package calendarv1

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

// OrderV1Client is the client API for OrderV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type OrderV1Client interface {
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
}

type orderV1Client struct {
	cc grpc.ClientConnInterface
}

func NewOrderV1Client(cc grpc.ClientConnInterface) OrderV1Client {
	return &orderV1Client{cc}
}

func (c *orderV1Client) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, "/api.order_v1.OrderV1/Create", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderV1Server is the server API for OrderV1 service.
// All implementations must embed UnimplementedOrderV1Server
// for forward compatibility
type OrderV1Server interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	mustEmbedUnimplementedOrderV1Server()
}

// UnimplementedOrderV1Server must be embedded to have forward compatible implementations.
type UnimplementedOrderV1Server struct {
}

func (UnimplementedOrderV1Server) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedOrderV1Server) mustEmbedUnimplementedOrderV1Server() {}

// UnsafeOrderV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to OrderV1Server will
// result in compilation errors.
type UnsafeOrderV1Server interface {
	mustEmbedUnimplementedOrderV1Server()
}

func RegisterOrderV1Server(s grpc.ServiceRegistrar, srv OrderV1Server) {
	s.RegisterService(&OrderV1_ServiceDesc, srv)
}

func _OrderV1_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderV1Server).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.order_v1.OrderV1/Create",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderV1Server).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// OrderV1_ServiceDesc is the grpc.ServiceDesc for OrderV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var OrderV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.order_v1.OrderV1",
	HandlerType: (*OrderV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Create",
			Handler:    _OrderV1_Create_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "order.proto",
}
