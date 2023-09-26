// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: api/qycms_bff/admin/v1/qy_admin_link.proto

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

// QyAdminLinkClient is the client API for QyAdminLink service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QyAdminLinkClient interface {
	// 新增友链
	CreateQyAdminLink(ctx context.Context, in *CreateQyAdminLinkRequest, opts ...grpc.CallOption) (*CreateQyAdminLinkReply, error)
	// 更新友链
	UpdateQyAdminLink(ctx context.Context, in *UpdateQyAdminLinkRequest, opts ...grpc.CallOption) (*UpdateQyAdminLinkReply, error)
	DeleteQyAdminLink(ctx context.Context, in *DeleteQyAdminLinkRequest, opts ...grpc.CallOption) (*DeleteQyAdminLinkReply, error)
	// 批量删除友链
	DeleteQyAdminLinks(ctx context.Context, in *DeleteQyAdminLinksRequest, opts ...grpc.CallOption) (*DeleteQyAdminLinkReply, error)
	GetQyAdminLink(ctx context.Context, in *GetQyAdminLinkRequest, opts ...grpc.CallOption) (*GetQyAdminLinkReply, error)
	// 友链列表
	ListQyAdminLink(ctx context.Context, in *ListQyAdminLinkRequest, opts ...grpc.CallOption) (*ListQyAdminLinkReply, error)
}

type qyAdminLinkClient struct {
	cc grpc.ClientConnInterface
}

func NewQyAdminLinkClient(cc grpc.ClientConnInterface) QyAdminLinkClient {
	return &qyAdminLinkClient{cc}
}

func (c *qyAdminLinkClient) CreateQyAdminLink(ctx context.Context, in *CreateQyAdminLinkRequest, opts ...grpc.CallOption) (*CreateQyAdminLinkReply, error) {
	out := new(CreateQyAdminLinkReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminLink/CreateQyAdminLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminLinkClient) UpdateQyAdminLink(ctx context.Context, in *UpdateQyAdminLinkRequest, opts ...grpc.CallOption) (*UpdateQyAdminLinkReply, error) {
	out := new(UpdateQyAdminLinkReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminLink/UpdateQyAdminLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminLinkClient) DeleteQyAdminLink(ctx context.Context, in *DeleteQyAdminLinkRequest, opts ...grpc.CallOption) (*DeleteQyAdminLinkReply, error) {
	out := new(DeleteQyAdminLinkReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminLink/DeleteQyAdminLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminLinkClient) DeleteQyAdminLinks(ctx context.Context, in *DeleteQyAdminLinksRequest, opts ...grpc.CallOption) (*DeleteQyAdminLinkReply, error) {
	out := new(DeleteQyAdminLinkReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminLink/DeleteQyAdminLinks", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminLinkClient) GetQyAdminLink(ctx context.Context, in *GetQyAdminLinkRequest, opts ...grpc.CallOption) (*GetQyAdminLinkReply, error) {
	out := new(GetQyAdminLinkReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminLink/GetQyAdminLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminLinkClient) ListQyAdminLink(ctx context.Context, in *ListQyAdminLinkRequest, opts ...grpc.CallOption) (*ListQyAdminLinkReply, error) {
	out := new(ListQyAdminLinkReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminLink/ListQyAdminLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QyAdminLinkServer is the server API for QyAdminLink service.
// All implementations must embed UnimplementedQyAdminLinkServer
// for forward compatibility
type QyAdminLinkServer interface {
	// 新增友链
	CreateQyAdminLink(context.Context, *CreateQyAdminLinkRequest) (*CreateQyAdminLinkReply, error)
	// 更新友链
	UpdateQyAdminLink(context.Context, *UpdateQyAdminLinkRequest) (*UpdateQyAdminLinkReply, error)
	DeleteQyAdminLink(context.Context, *DeleteQyAdminLinkRequest) (*DeleteQyAdminLinkReply, error)
	// 批量删除友链
	DeleteQyAdminLinks(context.Context, *DeleteQyAdminLinksRequest) (*DeleteQyAdminLinkReply, error)
	GetQyAdminLink(context.Context, *GetQyAdminLinkRequest) (*GetQyAdminLinkReply, error)
	// 友链列表
	ListQyAdminLink(context.Context, *ListQyAdminLinkRequest) (*ListQyAdminLinkReply, error)
	mustEmbedUnimplementedQyAdminLinkServer()
}

// UnimplementedQyAdminLinkServer must be embedded to have forward compatible implementations.
type UnimplementedQyAdminLinkServer struct {
}

func (UnimplementedQyAdminLinkServer) CreateQyAdminLink(context.Context, *CreateQyAdminLinkRequest) (*CreateQyAdminLinkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQyAdminLink not implemented")
}
func (UnimplementedQyAdminLinkServer) UpdateQyAdminLink(context.Context, *UpdateQyAdminLinkRequest) (*UpdateQyAdminLinkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQyAdminLink not implemented")
}
func (UnimplementedQyAdminLinkServer) DeleteQyAdminLink(context.Context, *DeleteQyAdminLinkRequest) (*DeleteQyAdminLinkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQyAdminLink not implemented")
}
func (UnimplementedQyAdminLinkServer) DeleteQyAdminLinks(context.Context, *DeleteQyAdminLinksRequest) (*DeleteQyAdminLinkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQyAdminLinks not implemented")
}
func (UnimplementedQyAdminLinkServer) GetQyAdminLink(context.Context, *GetQyAdminLinkRequest) (*GetQyAdminLinkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQyAdminLink not implemented")
}
func (UnimplementedQyAdminLinkServer) ListQyAdminLink(context.Context, *ListQyAdminLinkRequest) (*ListQyAdminLinkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQyAdminLink not implemented")
}
func (UnimplementedQyAdminLinkServer) mustEmbedUnimplementedQyAdminLinkServer() {}

// UnsafeQyAdminLinkServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QyAdminLinkServer will
// result in compilation errors.
type UnsafeQyAdminLinkServer interface {
	mustEmbedUnimplementedQyAdminLinkServer()
}

func RegisterQyAdminLinkServer(s grpc.ServiceRegistrar, srv QyAdminLinkServer) {
	s.RegisterService(&QyAdminLink_ServiceDesc, srv)
}

func _QyAdminLink_CreateQyAdminLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQyAdminLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminLinkServer).CreateQyAdminLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminLink/CreateQyAdminLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminLinkServer).CreateQyAdminLink(ctx, req.(*CreateQyAdminLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminLink_UpdateQyAdminLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQyAdminLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminLinkServer).UpdateQyAdminLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminLink/UpdateQyAdminLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminLinkServer).UpdateQyAdminLink(ctx, req.(*UpdateQyAdminLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminLink_DeleteQyAdminLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQyAdminLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminLinkServer).DeleteQyAdminLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminLink/DeleteQyAdminLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminLinkServer).DeleteQyAdminLink(ctx, req.(*DeleteQyAdminLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminLink_DeleteQyAdminLinks_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQyAdminLinksRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminLinkServer).DeleteQyAdminLinks(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminLink/DeleteQyAdminLinks",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminLinkServer).DeleteQyAdminLinks(ctx, req.(*DeleteQyAdminLinksRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminLink_GetQyAdminLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQyAdminLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminLinkServer).GetQyAdminLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminLink/GetQyAdminLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminLinkServer).GetQyAdminLink(ctx, req.(*GetQyAdminLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminLink_ListQyAdminLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQyAdminLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminLinkServer).ListQyAdminLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminLink/ListQyAdminLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminLinkServer).ListQyAdminLink(ctx, req.(*ListQyAdminLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QyAdminLink_ServiceDesc is the grpc.ServiceDesc for QyAdminLink service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QyAdminLink_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.qycms_bff.admin.v1.QyAdminLink",
	HandlerType: (*QyAdminLinkServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQyAdminLink",
			Handler:    _QyAdminLink_CreateQyAdminLink_Handler,
		},
		{
			MethodName: "UpdateQyAdminLink",
			Handler:    _QyAdminLink_UpdateQyAdminLink_Handler,
		},
		{
			MethodName: "DeleteQyAdminLink",
			Handler:    _QyAdminLink_DeleteQyAdminLink_Handler,
		},
		{
			MethodName: "DeleteQyAdminLinks",
			Handler:    _QyAdminLink_DeleteQyAdminLinks_Handler,
		},
		{
			MethodName: "GetQyAdminLink",
			Handler:    _QyAdminLink_GetQyAdminLink_Handler,
		},
		{
			MethodName: "ListQyAdminLink",
			Handler:    _QyAdminLink_ListQyAdminLink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/qycms_bff/admin/v1/qy_admin_link.proto",
}
