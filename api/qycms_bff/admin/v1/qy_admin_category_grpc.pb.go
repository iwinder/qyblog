// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: api/qycms_bff/admin/v1/qy_admin_category.proto

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

// QyAdminCategoryClient is the client API for QyAdminCategory service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QyAdminCategoryClient interface {
	// 新增
	CreateQyAdminCategory(ctx context.Context, in *CreateQyAdminCategoryRequest, opts ...grpc.CallOption) (*CreateQyAdminCategoryReply, error)
	// 更新
	UpdateQyAdminCategory(ctx context.Context, in *UpdateQyAdminCategoryRequest, opts ...grpc.CallOption) (*UpdateQyAdminCategoryReply, error)
	// 删除
	DeleteQyAdminCategory(ctx context.Context, in *DeleteQyAdminCategoryRequest, opts ...grpc.CallOption) (*DeleteQyAdminCategoryReply, error)
	GetQyAdminCategory(ctx context.Context, in *GetQyAdminCategoryRequest, opts ...grpc.CallOption) (*GetQyAdminCategoryReply, error)
	// 列表
	ListQyAdminCategory(ctx context.Context, in *ListQyAdminCategoryRequest, opts ...grpc.CallOption) (*ListQyAdminCategoryReply, error)
}

type qyAdminCategoryClient struct {
	cc grpc.ClientConnInterface
}

func NewQyAdminCategoryClient(cc grpc.ClientConnInterface) QyAdminCategoryClient {
	return &qyAdminCategoryClient{cc}
}

func (c *qyAdminCategoryClient) CreateQyAdminCategory(ctx context.Context, in *CreateQyAdminCategoryRequest, opts ...grpc.CallOption) (*CreateQyAdminCategoryReply, error) {
	out := new(CreateQyAdminCategoryReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminCategory/CreateQyAdminCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminCategoryClient) UpdateQyAdminCategory(ctx context.Context, in *UpdateQyAdminCategoryRequest, opts ...grpc.CallOption) (*UpdateQyAdminCategoryReply, error) {
	out := new(UpdateQyAdminCategoryReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminCategory/UpdateQyAdminCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminCategoryClient) DeleteQyAdminCategory(ctx context.Context, in *DeleteQyAdminCategoryRequest, opts ...grpc.CallOption) (*DeleteQyAdminCategoryReply, error) {
	out := new(DeleteQyAdminCategoryReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminCategory/DeleteQyAdminCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminCategoryClient) GetQyAdminCategory(ctx context.Context, in *GetQyAdminCategoryRequest, opts ...grpc.CallOption) (*GetQyAdminCategoryReply, error) {
	out := new(GetQyAdminCategoryReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminCategory/GetQyAdminCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminCategoryClient) ListQyAdminCategory(ctx context.Context, in *ListQyAdminCategoryRequest, opts ...grpc.CallOption) (*ListQyAdminCategoryReply, error) {
	out := new(ListQyAdminCategoryReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminCategory/ListQyAdminCategory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QyAdminCategoryServer is the server API for QyAdminCategory service.
// All implementations must embed UnimplementedQyAdminCategoryServer
// for forward compatibility
type QyAdminCategoryServer interface {
	// 新增
	CreateQyAdminCategory(context.Context, *CreateQyAdminCategoryRequest) (*CreateQyAdminCategoryReply, error)
	// 更新
	UpdateQyAdminCategory(context.Context, *UpdateQyAdminCategoryRequest) (*UpdateQyAdminCategoryReply, error)
	// 删除
	DeleteQyAdminCategory(context.Context, *DeleteQyAdminCategoryRequest) (*DeleteQyAdminCategoryReply, error)
	GetQyAdminCategory(context.Context, *GetQyAdminCategoryRequest) (*GetQyAdminCategoryReply, error)
	// 列表
	ListQyAdminCategory(context.Context, *ListQyAdminCategoryRequest) (*ListQyAdminCategoryReply, error)
	mustEmbedUnimplementedQyAdminCategoryServer()
}

// UnimplementedQyAdminCategoryServer must be embedded to have forward compatible implementations.
type UnimplementedQyAdminCategoryServer struct {
}

func (UnimplementedQyAdminCategoryServer) CreateQyAdminCategory(context.Context, *CreateQyAdminCategoryRequest) (*CreateQyAdminCategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQyAdminCategory not implemented")
}
func (UnimplementedQyAdminCategoryServer) UpdateQyAdminCategory(context.Context, *UpdateQyAdminCategoryRequest) (*UpdateQyAdminCategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQyAdminCategory not implemented")
}
func (UnimplementedQyAdminCategoryServer) DeleteQyAdminCategory(context.Context, *DeleteQyAdminCategoryRequest) (*DeleteQyAdminCategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQyAdminCategory not implemented")
}
func (UnimplementedQyAdminCategoryServer) GetQyAdminCategory(context.Context, *GetQyAdminCategoryRequest) (*GetQyAdminCategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQyAdminCategory not implemented")
}
func (UnimplementedQyAdminCategoryServer) ListQyAdminCategory(context.Context, *ListQyAdminCategoryRequest) (*ListQyAdminCategoryReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQyAdminCategory not implemented")
}
func (UnimplementedQyAdminCategoryServer) mustEmbedUnimplementedQyAdminCategoryServer() {}

// UnsafeQyAdminCategoryServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QyAdminCategoryServer will
// result in compilation errors.
type UnsafeQyAdminCategoryServer interface {
	mustEmbedUnimplementedQyAdminCategoryServer()
}

func RegisterQyAdminCategoryServer(s grpc.ServiceRegistrar, srv QyAdminCategoryServer) {
	s.RegisterService(&QyAdminCategory_ServiceDesc, srv)
}

func _QyAdminCategory_CreateQyAdminCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQyAdminCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminCategoryServer).CreateQyAdminCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminCategory/CreateQyAdminCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminCategoryServer).CreateQyAdminCategory(ctx, req.(*CreateQyAdminCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminCategory_UpdateQyAdminCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQyAdminCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminCategoryServer).UpdateQyAdminCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminCategory/UpdateQyAdminCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminCategoryServer).UpdateQyAdminCategory(ctx, req.(*UpdateQyAdminCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminCategory_DeleteQyAdminCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQyAdminCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminCategoryServer).DeleteQyAdminCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminCategory/DeleteQyAdminCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminCategoryServer).DeleteQyAdminCategory(ctx, req.(*DeleteQyAdminCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminCategory_GetQyAdminCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQyAdminCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminCategoryServer).GetQyAdminCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminCategory/GetQyAdminCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminCategoryServer).GetQyAdminCategory(ctx, req.(*GetQyAdminCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminCategory_ListQyAdminCategory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQyAdminCategoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminCategoryServer).ListQyAdminCategory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminCategory/ListQyAdminCategory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminCategoryServer).ListQyAdminCategory(ctx, req.(*ListQyAdminCategoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QyAdminCategory_ServiceDesc is the grpc.ServiceDesc for QyAdminCategory service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QyAdminCategory_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.qycms_bff.admin.v1.QyAdminCategory",
	HandlerType: (*QyAdminCategoryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQyAdminCategory",
			Handler:    _QyAdminCategory_CreateQyAdminCategory_Handler,
		},
		{
			MethodName: "UpdateQyAdminCategory",
			Handler:    _QyAdminCategory_UpdateQyAdminCategory_Handler,
		},
		{
			MethodName: "DeleteQyAdminCategory",
			Handler:    _QyAdminCategory_DeleteQyAdminCategory_Handler,
		},
		{
			MethodName: "GetQyAdminCategory",
			Handler:    _QyAdminCategory_GetQyAdminCategory_Handler,
		},
		{
			MethodName: "ListQyAdminCategory",
			Handler:    _QyAdminCategory_ListQyAdminCategory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/qycms_bff/admin/v1/qy_admin_category.proto",
}
