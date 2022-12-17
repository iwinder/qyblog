// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: api/qycms_bff/admin/v1/qy_admin_home.proto

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

// QyAdminHomeClient is the client API for QyAdminHome service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QyAdminHomeClient interface {
	UpdateContentCountJobsQyAdminHome(ctx context.Context, in *CreateQyAdminHomeRequest, opts ...grpc.CallOption) (*CreateQyAdminHomeReply, error)
	GeneratorMapJobQyAdminHome(ctx context.Context, in *CreateQyAdminHomeRequest, opts ...grpc.CallOption) (*CreateQyAdminHomeReply, error)
	UpdateAllPostsCountJobsQyAdminHome(ctx context.Context, in *CreateQyAdminHomeRequest, opts ...grpc.CallOption) (*CreateQyAdminHomeReply, error)
	EmailToNotSendCountJobsQyAdminHome(ctx context.Context, in *CreateQyAdminHomeRequest, opts ...grpc.CallOption) (*CreateQyAdminHomeReply, error)
	GetQyAdminHome(ctx context.Context, in *GetQyAdminHomeRequest, opts ...grpc.CallOption) (*GetQyAdminHomeReply, error)
	ListQyAdminHome(ctx context.Context, in *ListQyAdminHomeRequest, opts ...grpc.CallOption) (*ListQyAdminHomeReply, error)
}

type qyAdminHomeClient struct {
	cc grpc.ClientConnInterface
}

func NewQyAdminHomeClient(cc grpc.ClientConnInterface) QyAdminHomeClient {
	return &qyAdminHomeClient{cc}
}

func (c *qyAdminHomeClient) UpdateContentCountJobsQyAdminHome(ctx context.Context, in *CreateQyAdminHomeRequest, opts ...grpc.CallOption) (*CreateQyAdminHomeReply, error) {
	out := new(CreateQyAdminHomeReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminHome/UpdateContentCountJobsQyAdminHome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminHomeClient) GeneratorMapJobQyAdminHome(ctx context.Context, in *CreateQyAdminHomeRequest, opts ...grpc.CallOption) (*CreateQyAdminHomeReply, error) {
	out := new(CreateQyAdminHomeReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminHome/GeneratorMapJobQyAdminHome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminHomeClient) UpdateAllPostsCountJobsQyAdminHome(ctx context.Context, in *CreateQyAdminHomeRequest, opts ...grpc.CallOption) (*CreateQyAdminHomeReply, error) {
	out := new(CreateQyAdminHomeReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminHome/UpdateAllPostsCountJobsQyAdminHome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminHomeClient) EmailToNotSendCountJobsQyAdminHome(ctx context.Context, in *CreateQyAdminHomeRequest, opts ...grpc.CallOption) (*CreateQyAdminHomeReply, error) {
	out := new(CreateQyAdminHomeReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminHome/EmailToNotSendCountJobsQyAdminHome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminHomeClient) GetQyAdminHome(ctx context.Context, in *GetQyAdminHomeRequest, opts ...grpc.CallOption) (*GetQyAdminHomeReply, error) {
	out := new(GetQyAdminHomeReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminHome/GetQyAdminHome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminHomeClient) ListQyAdminHome(ctx context.Context, in *ListQyAdminHomeRequest, opts ...grpc.CallOption) (*ListQyAdminHomeReply, error) {
	out := new(ListQyAdminHomeReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminHome/ListQyAdminHome", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QyAdminHomeServer is the server API for QyAdminHome service.
// All implementations must embed UnimplementedQyAdminHomeServer
// for forward compatibility
type QyAdminHomeServer interface {
	UpdateContentCountJobsQyAdminHome(context.Context, *CreateQyAdminHomeRequest) (*CreateQyAdminHomeReply, error)
	GeneratorMapJobQyAdminHome(context.Context, *CreateQyAdminHomeRequest) (*CreateQyAdminHomeReply, error)
	UpdateAllPostsCountJobsQyAdminHome(context.Context, *CreateQyAdminHomeRequest) (*CreateQyAdminHomeReply, error)
	EmailToNotSendCountJobsQyAdminHome(context.Context, *CreateQyAdminHomeRequest) (*CreateQyAdminHomeReply, error)
	GetQyAdminHome(context.Context, *GetQyAdminHomeRequest) (*GetQyAdminHomeReply, error)
	ListQyAdminHome(context.Context, *ListQyAdminHomeRequest) (*ListQyAdminHomeReply, error)
	mustEmbedUnimplementedQyAdminHomeServer()
}

// UnimplementedQyAdminHomeServer must be embedded to have forward compatible implementations.
type UnimplementedQyAdminHomeServer struct {
}

func (UnimplementedQyAdminHomeServer) UpdateContentCountJobsQyAdminHome(context.Context, *CreateQyAdminHomeRequest) (*CreateQyAdminHomeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateContentCountJobsQyAdminHome not implemented")
}
func (UnimplementedQyAdminHomeServer) GeneratorMapJobQyAdminHome(context.Context, *CreateQyAdminHomeRequest) (*CreateQyAdminHomeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GeneratorMapJobQyAdminHome not implemented")
}
func (UnimplementedQyAdminHomeServer) UpdateAllPostsCountJobsQyAdminHome(context.Context, *CreateQyAdminHomeRequest) (*CreateQyAdminHomeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateAllPostsCountJobsQyAdminHome not implemented")
}
func (UnimplementedQyAdminHomeServer) EmailToNotSendCountJobsQyAdminHome(context.Context, *CreateQyAdminHomeRequest) (*CreateQyAdminHomeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EmailToNotSendCountJobsQyAdminHome not implemented")
}
func (UnimplementedQyAdminHomeServer) GetQyAdminHome(context.Context, *GetQyAdminHomeRequest) (*GetQyAdminHomeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQyAdminHome not implemented")
}
func (UnimplementedQyAdminHomeServer) ListQyAdminHome(context.Context, *ListQyAdminHomeRequest) (*ListQyAdminHomeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQyAdminHome not implemented")
}
func (UnimplementedQyAdminHomeServer) mustEmbedUnimplementedQyAdminHomeServer() {}

// UnsafeQyAdminHomeServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QyAdminHomeServer will
// result in compilation errors.
type UnsafeQyAdminHomeServer interface {
	mustEmbedUnimplementedQyAdminHomeServer()
}

func RegisterQyAdminHomeServer(s grpc.ServiceRegistrar, srv QyAdminHomeServer) {
	s.RegisterService(&QyAdminHome_ServiceDesc, srv)
}

func _QyAdminHome_UpdateContentCountJobsQyAdminHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQyAdminHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminHomeServer).UpdateContentCountJobsQyAdminHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminHome/UpdateContentCountJobsQyAdminHome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminHomeServer).UpdateContentCountJobsQyAdminHome(ctx, req.(*CreateQyAdminHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminHome_GeneratorMapJobQyAdminHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQyAdminHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminHomeServer).GeneratorMapJobQyAdminHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminHome/GeneratorMapJobQyAdminHome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminHomeServer).GeneratorMapJobQyAdminHome(ctx, req.(*CreateQyAdminHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminHome_UpdateAllPostsCountJobsQyAdminHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQyAdminHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminHomeServer).UpdateAllPostsCountJobsQyAdminHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminHome/UpdateAllPostsCountJobsQyAdminHome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminHomeServer).UpdateAllPostsCountJobsQyAdminHome(ctx, req.(*CreateQyAdminHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminHome_EmailToNotSendCountJobsQyAdminHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQyAdminHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminHomeServer).EmailToNotSendCountJobsQyAdminHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminHome/EmailToNotSendCountJobsQyAdminHome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminHomeServer).EmailToNotSendCountJobsQyAdminHome(ctx, req.(*CreateQyAdminHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminHome_GetQyAdminHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQyAdminHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminHomeServer).GetQyAdminHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminHome/GetQyAdminHome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminHomeServer).GetQyAdminHome(ctx, req.(*GetQyAdminHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminHome_ListQyAdminHome_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQyAdminHomeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminHomeServer).ListQyAdminHome(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminHome/ListQyAdminHome",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminHomeServer).ListQyAdminHome(ctx, req.(*ListQyAdminHomeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QyAdminHome_ServiceDesc is the grpc.ServiceDesc for QyAdminHome service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QyAdminHome_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.qycms_bff.admin.v1.QyAdminHome",
	HandlerType: (*QyAdminHomeServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateContentCountJobsQyAdminHome",
			Handler:    _QyAdminHome_UpdateContentCountJobsQyAdminHome_Handler,
		},
		{
			MethodName: "GeneratorMapJobQyAdminHome",
			Handler:    _QyAdminHome_GeneratorMapJobQyAdminHome_Handler,
		},
		{
			MethodName: "UpdateAllPostsCountJobsQyAdminHome",
			Handler:    _QyAdminHome_UpdateAllPostsCountJobsQyAdminHome_Handler,
		},
		{
			MethodName: "EmailToNotSendCountJobsQyAdminHome",
			Handler:    _QyAdminHome_EmailToNotSendCountJobsQyAdminHome_Handler,
		},
		{
			MethodName: "GetQyAdminHome",
			Handler:    _QyAdminHome_GetQyAdminHome_Handler,
		},
		{
			MethodName: "ListQyAdminHome",
			Handler:    _QyAdminHome_ListQyAdminHome_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/qycms_bff/admin/v1/qy_admin_home.proto",
}
