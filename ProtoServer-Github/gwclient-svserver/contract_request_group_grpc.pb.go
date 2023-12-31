// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.2
// source: contract_request_group.proto

package gwclient_svserver

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

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	// Auth
	GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error)
	// Main
	GetGroupDetails(ctx context.Context, in *GetGroupDetailsRequest, opts ...grpc.CallOption) (*GetGroupDetailsResponse, error)
	GetRequestDetails(ctx context.Context, in *GetWithdrawsRequest, opts ...grpc.CallOption) (*GetWithdrawsResponse, error)
	GetContractDetails(ctx context.Context, in *GetContractRequest, opts ...grpc.CallOption) (*GetContractResponse, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) GetUser(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*GetUserResponse, error) {
	out := new(GetUserResponse)
	err := c.cc.Invoke(ctx, "/protoGRPCagw.Service/GetUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetGroupDetails(ctx context.Context, in *GetGroupDetailsRequest, opts ...grpc.CallOption) (*GetGroupDetailsResponse, error) {
	out := new(GetGroupDetailsResponse)
	err := c.cc.Invoke(ctx, "/protoGRPCagw.Service/GetGroupDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetRequestDetails(ctx context.Context, in *GetWithdrawsRequest, opts ...grpc.CallOption) (*GetWithdrawsResponse, error) {
	out := new(GetWithdrawsResponse)
	err := c.cc.Invoke(ctx, "/protoGRPCagw.Service/GetRequestDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *serviceClient) GetContractDetails(ctx context.Context, in *GetContractRequest, opts ...grpc.CallOption) (*GetContractResponse, error) {
	out := new(GetContractResponse)
	err := c.cc.Invoke(ctx, "/protoGRPCagw.Service/GetContractDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	// Auth
	GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error)
	// Main
	GetGroupDetails(context.Context, *GetGroupDetailsRequest) (*GetGroupDetailsResponse, error)
	GetRequestDetails(context.Context, *GetWithdrawsRequest) (*GetWithdrawsResponse, error)
	GetContractDetails(context.Context, *GetContractRequest) (*GetContractResponse, error)
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) GetUser(context.Context, *GetUserRequest) (*GetUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUser not implemented")
}
func (UnimplementedServiceServer) GetGroupDetails(context.Context, *GetGroupDetailsRequest) (*GetGroupDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGroupDetails not implemented")
}
func (UnimplementedServiceServer) GetRequestDetails(context.Context, *GetWithdrawsRequest) (*GetWithdrawsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRequestDetails not implemented")
}
func (UnimplementedServiceServer) GetContractDetails(context.Context, *GetContractRequest) (*GetContractResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetContractDetails not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_GetUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoGRPCagw.Service/GetUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetUser(ctx, req.(*GetUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetGroupDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetGroupDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetGroupDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoGRPCagw.Service/GetGroupDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetGroupDetails(ctx, req.(*GetGroupDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetRequestDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetWithdrawsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetRequestDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoGRPCagw.Service/GetRequestDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetRequestDetails(ctx, req.(*GetWithdrawsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Service_GetContractDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetContractRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceServer).GetContractDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/protoGRPCagw.Service/GetContractDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceServer).GetContractDetails(ctx, req.(*GetContractRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "protoGRPCagw.Service",
	HandlerType: (*ServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetUser",
			Handler:    _Service_GetUser_Handler,
		},
		{
			MethodName: "GetGroupDetails",
			Handler:    _Service_GetGroupDetails_Handler,
		},
		{
			MethodName: "GetRequestDetails",
			Handler:    _Service_GetRequestDetails_Handler,
		},
		{
			MethodName: "GetContractDetails",
			Handler:    _Service_GetContractDetails_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "contract_request_group.proto",
}
