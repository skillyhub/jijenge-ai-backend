// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.12
// source: pb/business.proto

package gen

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

const (
	BusinessService_CreateBusiness_FullMethodName   = "/finance.BusinessService/CreateBusiness"
	BusinessService_GetBusiness_FullMethodName      = "/finance.BusinessService/GetBusiness"
	BusinessService_UpdateBusiness_FullMethodName   = "/finance.BusinessService/UpdateBusiness"
	BusinessService_DeleteBusiness_FullMethodName   = "/finance.BusinessService/DeleteBusiness"
	BusinessService_ListBusinesses_FullMethodName   = "/finance.BusinessService/ListBusinesses"
	BusinessService_SearchBusinesses_FullMethodName = "/finance.BusinessService/SearchBusinesses"
)

// BusinessServiceClient is the client API for BusinessService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BusinessServiceClient interface {
	CreateBusiness(ctx context.Context, in *CreateBusinessRequest, opts ...grpc.CallOption) (*BusinessResponse, error)
	GetBusiness(ctx context.Context, in *GetBusinessRequest, opts ...grpc.CallOption) (*BusinessResponse, error)
	UpdateBusiness(ctx context.Context, in *UpdateBusinessRequest, opts ...grpc.CallOption) (*BusinessResponse, error)
	DeleteBusiness(ctx context.Context, in *DeleteBusinessRequest, opts ...grpc.CallOption) (*DeleteBusinessResponse, error)
	ListBusinesses(ctx context.Context, in *ListBusinessesRequest, opts ...grpc.CallOption) (*ListBusinessesResponse, error)
	SearchBusinesses(ctx context.Context, in *SearchBusinessesRequest, opts ...grpc.CallOption) (*ListBusinessesResponse, error)
}

type businessServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBusinessServiceClient(cc grpc.ClientConnInterface) BusinessServiceClient {
	return &businessServiceClient{cc}
}

func (c *businessServiceClient) CreateBusiness(ctx context.Context, in *CreateBusinessRequest, opts ...grpc.CallOption) (*BusinessResponse, error) {
	out := new(BusinessResponse)
	err := c.cc.Invoke(ctx, BusinessService_CreateBusiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessServiceClient) GetBusiness(ctx context.Context, in *GetBusinessRequest, opts ...grpc.CallOption) (*BusinessResponse, error) {
	out := new(BusinessResponse)
	err := c.cc.Invoke(ctx, BusinessService_GetBusiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessServiceClient) UpdateBusiness(ctx context.Context, in *UpdateBusinessRequest, opts ...grpc.CallOption) (*BusinessResponse, error) {
	out := new(BusinessResponse)
	err := c.cc.Invoke(ctx, BusinessService_UpdateBusiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessServiceClient) DeleteBusiness(ctx context.Context, in *DeleteBusinessRequest, opts ...grpc.CallOption) (*DeleteBusinessResponse, error) {
	out := new(DeleteBusinessResponse)
	err := c.cc.Invoke(ctx, BusinessService_DeleteBusiness_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessServiceClient) ListBusinesses(ctx context.Context, in *ListBusinessesRequest, opts ...grpc.CallOption) (*ListBusinessesResponse, error) {
	out := new(ListBusinessesResponse)
	err := c.cc.Invoke(ctx, BusinessService_ListBusinesses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *businessServiceClient) SearchBusinesses(ctx context.Context, in *SearchBusinessesRequest, opts ...grpc.CallOption) (*ListBusinessesResponse, error) {
	out := new(ListBusinessesResponse)
	err := c.cc.Invoke(ctx, BusinessService_SearchBusinesses_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BusinessServiceServer is the server API for BusinessService service.
// All implementations must embed UnimplementedBusinessServiceServer
// for forward compatibility
type BusinessServiceServer interface {
	CreateBusiness(context.Context, *CreateBusinessRequest) (*BusinessResponse, error)
	GetBusiness(context.Context, *GetBusinessRequest) (*BusinessResponse, error)
	UpdateBusiness(context.Context, *UpdateBusinessRequest) (*BusinessResponse, error)
	DeleteBusiness(context.Context, *DeleteBusinessRequest) (*DeleteBusinessResponse, error)
	ListBusinesses(context.Context, *ListBusinessesRequest) (*ListBusinessesResponse, error)
	SearchBusinesses(context.Context, *SearchBusinessesRequest) (*ListBusinessesResponse, error)
	mustEmbedUnimplementedBusinessServiceServer()
}

// UnimplementedBusinessServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBusinessServiceServer struct {
}

func (UnimplementedBusinessServiceServer) CreateBusiness(context.Context, *CreateBusinessRequest) (*BusinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateBusiness not implemented")
}
func (UnimplementedBusinessServiceServer) GetBusiness(context.Context, *GetBusinessRequest) (*BusinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBusiness not implemented")
}
func (UnimplementedBusinessServiceServer) UpdateBusiness(context.Context, *UpdateBusinessRequest) (*BusinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateBusiness not implemented")
}
func (UnimplementedBusinessServiceServer) DeleteBusiness(context.Context, *DeleteBusinessRequest) (*DeleteBusinessResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteBusiness not implemented")
}
func (UnimplementedBusinessServiceServer) ListBusinesses(context.Context, *ListBusinessesRequest) (*ListBusinessesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListBusinesses not implemented")
}
func (UnimplementedBusinessServiceServer) SearchBusinesses(context.Context, *SearchBusinessesRequest) (*ListBusinessesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchBusinesses not implemented")
}
func (UnimplementedBusinessServiceServer) mustEmbedUnimplementedBusinessServiceServer() {}

// UnsafeBusinessServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BusinessServiceServer will
// result in compilation errors.
type UnsafeBusinessServiceServer interface {
	mustEmbedUnimplementedBusinessServiceServer()
}

func RegisterBusinessServiceServer(s grpc.ServiceRegistrar, srv BusinessServiceServer) {
	s.RegisterService(&BusinessService_ServiceDesc, srv)
}

func _BusinessService_CreateBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServiceServer).CreateBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusinessService_CreateBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServiceServer).CreateBusiness(ctx, req.(*CreateBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessService_GetBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServiceServer).GetBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusinessService_GetBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServiceServer).GetBusiness(ctx, req.(*GetBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessService_UpdateBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServiceServer).UpdateBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusinessService_UpdateBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServiceServer).UpdateBusiness(ctx, req.(*UpdateBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessService_DeleteBusiness_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteBusinessRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServiceServer).DeleteBusiness(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusinessService_DeleteBusiness_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServiceServer).DeleteBusiness(ctx, req.(*DeleteBusinessRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessService_ListBusinesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListBusinessesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServiceServer).ListBusinesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusinessService_ListBusinesses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServiceServer).ListBusinesses(ctx, req.(*ListBusinessesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BusinessService_SearchBusinesses_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchBusinessesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BusinessServiceServer).SearchBusinesses(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BusinessService_SearchBusinesses_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BusinessServiceServer).SearchBusinesses(ctx, req.(*SearchBusinessesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BusinessService_ServiceDesc is the grpc.ServiceDesc for BusinessService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BusinessService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "finance.BusinessService",
	HandlerType: (*BusinessServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateBusiness",
			Handler:    _BusinessService_CreateBusiness_Handler,
		},
		{
			MethodName: "GetBusiness",
			Handler:    _BusinessService_GetBusiness_Handler,
		},
		{
			MethodName: "UpdateBusiness",
			Handler:    _BusinessService_UpdateBusiness_Handler,
		},
		{
			MethodName: "DeleteBusiness",
			Handler:    _BusinessService_DeleteBusiness_Handler,
		},
		{
			MethodName: "ListBusinesses",
			Handler:    _BusinessService_ListBusinesses_Handler,
		},
		{
			MethodName: "SearchBusinesses",
			Handler:    _BusinessService_SearchBusinesses_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/business.proto",
}
