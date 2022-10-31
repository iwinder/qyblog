// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: api/qycms_bff/admin/v1/qy_admin_menus_admin.proto

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

// QyAdminMenusAdminClient is the client API for QyAdminMenusAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QyAdminMenusAdminClient interface {
	CreateQyAdminMenusAdmin(ctx context.Context, in *CreateQyAdminMenusAdminRequest, opts ...grpc.CallOption) (*CreateQyAdminMenusAdminReply, error)
	UpdateQyAdminMenusAdmin(ctx context.Context, in *UpdateQyAdminMenusAdminRequest, opts ...grpc.CallOption) (*UpdateQyAdminMenusAdminReply, error)
	DeleteQyAdminMenusAdmin(ctx context.Context, in *DeleteQyAdminMenusAdminRequest, opts ...grpc.CallOption) (*DeleteQyAdminMenusAdminReply, error)
	DeleteQyAdminMenusAdmins(ctx context.Context, in *DeleteQyAdminMenusAdminsRequest, opts ...grpc.CallOption) (*DeleteQyAdminMenusAdminsReply, error)
	GetQyAdminMenusAdmin(ctx context.Context, in *GetQyAdminMenusAdminRequest, opts ...grpc.CallOption) (*GetQyAdminMenusAdminReply, error)
	GetMyMenusAdminInfo(ctx context.Context, in *GetMyMenusAdminInfoReq, opts ...grpc.CallOption) (*GetMyMenusAdminInfoReply, error)
	ListQyAdminMenusAdmin(ctx context.Context, in *ListQyAdminMenusAdminRequest, opts ...grpc.CallOption) (*ListQyAdminMenusAdminReply, error)
}

type qyAdminMenusAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewQyAdminMenusAdminClient(cc grpc.ClientConnInterface) QyAdminMenusAdminClient {
	return &qyAdminMenusAdminClient{cc}
}

func (c *qyAdminMenusAdminClient) CreateQyAdminMenusAdmin(ctx context.Context, in *CreateQyAdminMenusAdminRequest, opts ...grpc.CallOption) (*CreateQyAdminMenusAdminReply, error) {
	out := new(CreateQyAdminMenusAdminReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/CreateQyAdminMenusAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminMenusAdminClient) UpdateQyAdminMenusAdmin(ctx context.Context, in *UpdateQyAdminMenusAdminRequest, opts ...grpc.CallOption) (*UpdateQyAdminMenusAdminReply, error) {
	out := new(UpdateQyAdminMenusAdminReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/UpdateQyAdminMenusAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminMenusAdminClient) DeleteQyAdminMenusAdmin(ctx context.Context, in *DeleteQyAdminMenusAdminRequest, opts ...grpc.CallOption) (*DeleteQyAdminMenusAdminReply, error) {
	out := new(DeleteQyAdminMenusAdminReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/DeleteQyAdminMenusAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminMenusAdminClient) DeleteQyAdminMenusAdmins(ctx context.Context, in *DeleteQyAdminMenusAdminsRequest, opts ...grpc.CallOption) (*DeleteQyAdminMenusAdminsReply, error) {
	out := new(DeleteQyAdminMenusAdminsReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/DeleteQyAdminMenusAdmins", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminMenusAdminClient) GetQyAdminMenusAdmin(ctx context.Context, in *GetQyAdminMenusAdminRequest, opts ...grpc.CallOption) (*GetQyAdminMenusAdminReply, error) {
	out := new(GetQyAdminMenusAdminReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/GetQyAdminMenusAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminMenusAdminClient) GetMyMenusAdminInfo(ctx context.Context, in *GetMyMenusAdminInfoReq, opts ...grpc.CallOption) (*GetMyMenusAdminInfoReply, error) {
	out := new(GetMyMenusAdminInfoReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/GetMyMenusAdminInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminMenusAdminClient) ListQyAdminMenusAdmin(ctx context.Context, in *ListQyAdminMenusAdminRequest, opts ...grpc.CallOption) (*ListQyAdminMenusAdminReply, error) {
	out := new(ListQyAdminMenusAdminReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/ListQyAdminMenusAdmin", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QyAdminMenusAdminServer is the server API for QyAdminMenusAdmin service.
// All implementations must embed UnimplementedQyAdminMenusAdminServer
// for forward compatibility
type QyAdminMenusAdminServer interface {
	CreateQyAdminMenusAdmin(context.Context, *CreateQyAdminMenusAdminRequest) (*CreateQyAdminMenusAdminReply, error)
	UpdateQyAdminMenusAdmin(context.Context, *UpdateQyAdminMenusAdminRequest) (*UpdateQyAdminMenusAdminReply, error)
	DeleteQyAdminMenusAdmin(context.Context, *DeleteQyAdminMenusAdminRequest) (*DeleteQyAdminMenusAdminReply, error)
	DeleteQyAdminMenusAdmins(context.Context, *DeleteQyAdminMenusAdminsRequest) (*DeleteQyAdminMenusAdminsReply, error)
	GetQyAdminMenusAdmin(context.Context, *GetQyAdminMenusAdminRequest) (*GetQyAdminMenusAdminReply, error)
	GetMyMenusAdminInfo(context.Context, *GetMyMenusAdminInfoReq) (*GetMyMenusAdminInfoReply, error)
	ListQyAdminMenusAdmin(context.Context, *ListQyAdminMenusAdminRequest) (*ListQyAdminMenusAdminReply, error)
	mustEmbedUnimplementedQyAdminMenusAdminServer()
}

// UnimplementedQyAdminMenusAdminServer must be embedded to have forward compatible implementations.
type UnimplementedQyAdminMenusAdminServer struct {
}

func (UnimplementedQyAdminMenusAdminServer) CreateQyAdminMenusAdmin(context.Context, *CreateQyAdminMenusAdminRequest) (*CreateQyAdminMenusAdminReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQyAdminMenusAdmin not implemented")
}
func (UnimplementedQyAdminMenusAdminServer) UpdateQyAdminMenusAdmin(context.Context, *UpdateQyAdminMenusAdminRequest) (*UpdateQyAdminMenusAdminReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQyAdminMenusAdmin not implemented")
}
func (UnimplementedQyAdminMenusAdminServer) DeleteQyAdminMenusAdmin(context.Context, *DeleteQyAdminMenusAdminRequest) (*DeleteQyAdminMenusAdminReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQyAdminMenusAdmin not implemented")
}
func (UnimplementedQyAdminMenusAdminServer) DeleteQyAdminMenusAdmins(context.Context, *DeleteQyAdminMenusAdminsRequest) (*DeleteQyAdminMenusAdminsReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQyAdminMenusAdmins not implemented")
}
func (UnimplementedQyAdminMenusAdminServer) GetQyAdminMenusAdmin(context.Context, *GetQyAdminMenusAdminRequest) (*GetQyAdminMenusAdminReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQyAdminMenusAdmin not implemented")
}
func (UnimplementedQyAdminMenusAdminServer) GetMyMenusAdminInfo(context.Context, *GetMyMenusAdminInfoReq) (*GetMyMenusAdminInfoReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMyMenusAdminInfo not implemented")
}
func (UnimplementedQyAdminMenusAdminServer) ListQyAdminMenusAdmin(context.Context, *ListQyAdminMenusAdminRequest) (*ListQyAdminMenusAdminReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQyAdminMenusAdmin not implemented")
}
func (UnimplementedQyAdminMenusAdminServer) mustEmbedUnimplementedQyAdminMenusAdminServer() {}

// UnsafeQyAdminMenusAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QyAdminMenusAdminServer will
// result in compilation errors.
type UnsafeQyAdminMenusAdminServer interface {
	mustEmbedUnimplementedQyAdminMenusAdminServer()
}

func RegisterQyAdminMenusAdminServer(s grpc.ServiceRegistrar, srv QyAdminMenusAdminServer) {
	s.RegisterService(&QyAdminMenusAdmin_ServiceDesc, srv)
}

func _QyAdminMenusAdmin_CreateQyAdminMenusAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQyAdminMenusAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminMenusAdminServer).CreateQyAdminMenusAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/CreateQyAdminMenusAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminMenusAdminServer).CreateQyAdminMenusAdmin(ctx, req.(*CreateQyAdminMenusAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminMenusAdmin_UpdateQyAdminMenusAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQyAdminMenusAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminMenusAdminServer).UpdateQyAdminMenusAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/UpdateQyAdminMenusAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminMenusAdminServer).UpdateQyAdminMenusAdmin(ctx, req.(*UpdateQyAdminMenusAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminMenusAdmin_DeleteQyAdminMenusAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQyAdminMenusAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminMenusAdminServer).DeleteQyAdminMenusAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/DeleteQyAdminMenusAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminMenusAdminServer).DeleteQyAdminMenusAdmin(ctx, req.(*DeleteQyAdminMenusAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminMenusAdmin_DeleteQyAdminMenusAdmins_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQyAdminMenusAdminsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminMenusAdminServer).DeleteQyAdminMenusAdmins(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/DeleteQyAdminMenusAdmins",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminMenusAdminServer).DeleteQyAdminMenusAdmins(ctx, req.(*DeleteQyAdminMenusAdminsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminMenusAdmin_GetQyAdminMenusAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQyAdminMenusAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminMenusAdminServer).GetQyAdminMenusAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/GetQyAdminMenusAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminMenusAdminServer).GetQyAdminMenusAdmin(ctx, req.(*GetQyAdminMenusAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminMenusAdmin_GetMyMenusAdminInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetMyMenusAdminInfoReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminMenusAdminServer).GetMyMenusAdminInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/GetMyMenusAdminInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminMenusAdminServer).GetMyMenusAdminInfo(ctx, req.(*GetMyMenusAdminInfoReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminMenusAdmin_ListQyAdminMenusAdmin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQyAdminMenusAdminRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminMenusAdminServer).ListQyAdminMenusAdmin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminMenusAdmin/ListQyAdminMenusAdmin",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminMenusAdminServer).ListQyAdminMenusAdmin(ctx, req.(*ListQyAdminMenusAdminRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QyAdminMenusAdmin_ServiceDesc is the grpc.ServiceDesc for QyAdminMenusAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QyAdminMenusAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.qycms_bff.admin.v1.QyAdminMenusAdmin",
	HandlerType: (*QyAdminMenusAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQyAdminMenusAdmin",
			Handler:    _QyAdminMenusAdmin_CreateQyAdminMenusAdmin_Handler,
		},
		{
			MethodName: "UpdateQyAdminMenusAdmin",
			Handler:    _QyAdminMenusAdmin_UpdateQyAdminMenusAdmin_Handler,
		},
		{
			MethodName: "DeleteQyAdminMenusAdmin",
			Handler:    _QyAdminMenusAdmin_DeleteQyAdminMenusAdmin_Handler,
		},
		{
			MethodName: "DeleteQyAdminMenusAdmins",
			Handler:    _QyAdminMenusAdmin_DeleteQyAdminMenusAdmins_Handler,
		},
		{
			MethodName: "GetQyAdminMenusAdmin",
			Handler:    _QyAdminMenusAdmin_GetQyAdminMenusAdmin_Handler,
		},
		{
			MethodName: "GetMyMenusAdminInfo",
			Handler:    _QyAdminMenusAdmin_GetMyMenusAdminInfo_Handler,
		},
		{
			MethodName: "ListQyAdminMenusAdmin",
			Handler:    _QyAdminMenusAdmin_ListQyAdminMenusAdmin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/qycms_bff/admin/v1/qy_admin_menus_admin.proto",
}