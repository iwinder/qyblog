// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: api/qycms_bff/admin/v1/qy_admin_menus.proto

package v1

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

// QyAdminMenusClient is the client API for QyAdminMenus service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QyAdminMenusClient interface {
	// 新增
	CreateQyAdminMenus(ctx context.Context, in *CreateQyAdminMenusRequest, opts ...grpc.CallOption) (*CreateQyAdminMenusReply, error)
	// 更新
	UpdateQyAdminMenus(ctx context.Context, in *UpdateQyAdminMenusRequest, opts ...grpc.CallOption) (*UpdateQyAdminMenusReply, error)
	// 删除
	DeleteQyAdminMenus(ctx context.Context, in *DeleteQyAdminMenusRequest, opts ...grpc.CallOption) (*DeleteQyAdminMenusReply, error)
	GetQyAdminMenus(ctx context.Context, in *GetQyAdminMenusRequest, opts ...grpc.CallOption) (*GetQyAdminMenusReply, error)
	// 列表
	ListQyAdminMenus(ctx context.Context, in *ListQyAdminMenusRequest, opts ...grpc.CallOption) (*ListQyAdminMenusReply, error)
}

type qyAdminMenusClient struct {
	cc grpc.ClientConnInterface
}

func NewQyAdminMenusClient(cc grpc.ClientConnInterface) QyAdminMenusClient {
	return &qyAdminMenusClient{cc}
}

func (c *qyAdminMenusClient) CreateQyAdminMenus(ctx context.Context, in *CreateQyAdminMenusRequest, opts ...grpc.CallOption) (*CreateQyAdminMenusReply, error) {
	out := new(CreateQyAdminMenusReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminMenus/CreateQyAdminMenus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminMenusClient) UpdateQyAdminMenus(ctx context.Context, in *UpdateQyAdminMenusRequest, opts ...grpc.CallOption) (*UpdateQyAdminMenusReply, error) {
	out := new(UpdateQyAdminMenusReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminMenus/UpdateQyAdminMenus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminMenusClient) DeleteQyAdminMenus(ctx context.Context, in *DeleteQyAdminMenusRequest, opts ...grpc.CallOption) (*DeleteQyAdminMenusReply, error) {
	out := new(DeleteQyAdminMenusReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminMenus/DeleteQyAdminMenus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminMenusClient) GetQyAdminMenus(ctx context.Context, in *GetQyAdminMenusRequest, opts ...grpc.CallOption) (*GetQyAdminMenusReply, error) {
	out := new(GetQyAdminMenusReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminMenus/GetQyAdminMenus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminMenusClient) ListQyAdminMenus(ctx context.Context, in *ListQyAdminMenusRequest, opts ...grpc.CallOption) (*ListQyAdminMenusReply, error) {
	out := new(ListQyAdminMenusReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminMenus/ListQyAdminMenus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QyAdminMenusServer is the server API for QyAdminMenus service.
// All implementations must embed UnimplementedQyAdminMenusServer
// for forward compatibility
type QyAdminMenusServer interface {
	// 新增
	CreateQyAdminMenus(context.Context, *CreateQyAdminMenusRequest) (*CreateQyAdminMenusReply, error)
	// 更新
	UpdateQyAdminMenus(context.Context, *UpdateQyAdminMenusRequest) (*UpdateQyAdminMenusReply, error)
	// 删除
	DeleteQyAdminMenus(context.Context, *DeleteQyAdminMenusRequest) (*DeleteQyAdminMenusReply, error)
	GetQyAdminMenus(context.Context, *GetQyAdminMenusRequest) (*GetQyAdminMenusReply, error)
	// 列表
	ListQyAdminMenus(context.Context, *ListQyAdminMenusRequest) (*ListQyAdminMenusReply, error)
	mustEmbedUnimplementedQyAdminMenusServer()
}

// UnimplementedQyAdminMenusServer must be embedded to have forward compatible implementations.
type UnimplementedQyAdminMenusServer struct {
}

func (UnimplementedQyAdminMenusServer) CreateQyAdminMenus(context.Context, *CreateQyAdminMenusRequest) (*CreateQyAdminMenusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQyAdminMenus not implemented")
}
func (UnimplementedQyAdminMenusServer) UpdateQyAdminMenus(context.Context, *UpdateQyAdminMenusRequest) (*UpdateQyAdminMenusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQyAdminMenus not implemented")
}
func (UnimplementedQyAdminMenusServer) DeleteQyAdminMenus(context.Context, *DeleteQyAdminMenusRequest) (*DeleteQyAdminMenusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQyAdminMenus not implemented")
}
func (UnimplementedQyAdminMenusServer) GetQyAdminMenus(context.Context, *GetQyAdminMenusRequest) (*GetQyAdminMenusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQyAdminMenus not implemented")
}
func (UnimplementedQyAdminMenusServer) ListQyAdminMenus(context.Context, *ListQyAdminMenusRequest) (*ListQyAdminMenusReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQyAdminMenus not implemented")
}
func (UnimplementedQyAdminMenusServer) mustEmbedUnimplementedQyAdminMenusServer() {}

// UnsafeQyAdminMenusServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QyAdminMenusServer will
// result in compilation errors.
type UnsafeQyAdminMenusServer interface {
	mustEmbedUnimplementedQyAdminMenusServer()
}

func RegisterQyAdminMenusServer(s grpc.ServiceRegistrar, srv QyAdminMenusServer) {
	s.RegisterService(&QyAdminMenus_ServiceDesc, srv)
}

func _QyAdminMenus_CreateQyAdminMenus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQyAdminMenusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminMenusServer).CreateQyAdminMenus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminMenus/CreateQyAdminMenus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminMenusServer).CreateQyAdminMenus(ctx, req.(*CreateQyAdminMenusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminMenus_UpdateQyAdminMenus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQyAdminMenusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminMenusServer).UpdateQyAdminMenus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminMenus/UpdateQyAdminMenus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminMenusServer).UpdateQyAdminMenus(ctx, req.(*UpdateQyAdminMenusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminMenus_DeleteQyAdminMenus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQyAdminMenusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminMenusServer).DeleteQyAdminMenus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminMenus/DeleteQyAdminMenus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminMenusServer).DeleteQyAdminMenus(ctx, req.(*DeleteQyAdminMenusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminMenus_GetQyAdminMenus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQyAdminMenusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminMenusServer).GetQyAdminMenus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminMenus/GetQyAdminMenus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminMenusServer).GetQyAdminMenus(ctx, req.(*GetQyAdminMenusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminMenus_ListQyAdminMenus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQyAdminMenusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminMenusServer).ListQyAdminMenus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminMenus/ListQyAdminMenus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminMenusServer).ListQyAdminMenus(ctx, req.(*ListQyAdminMenusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QyAdminMenus_ServiceDesc is the grpc.ServiceDesc for QyAdminMenus service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QyAdminMenus_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.qycms_bff.admin.v1.QyAdminMenus",
	HandlerType: (*QyAdminMenusServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQyAdminMenus",
			Handler:    _QyAdminMenus_CreateQyAdminMenus_Handler,
		},
		{
			MethodName: "UpdateQyAdminMenus",
			Handler:    _QyAdminMenus_UpdateQyAdminMenus_Handler,
		},
		{
			MethodName: "DeleteQyAdminMenus",
			Handler:    _QyAdminMenus_DeleteQyAdminMenus_Handler,
		},
		{
			MethodName: "GetQyAdminMenus",
			Handler:    _QyAdminMenus_GetQyAdminMenus_Handler,
		},
		{
			MethodName: "ListQyAdminMenus",
			Handler:    _QyAdminMenus_ListQyAdminMenus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/qycms_bff/admin/v1/qy_admin_menus.proto",
}
