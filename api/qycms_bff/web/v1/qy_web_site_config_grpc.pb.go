// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: api/qycms_bff/web/v1/qy_web_site_config.proto

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

// QyWebSiteConfigClient is the client API for QyWebSiteConfig service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QyWebSiteConfigClient interface {
	CreateQyWebSiteConfig(ctx context.Context, in *CreateQyWebSiteConfigRequest, opts ...grpc.CallOption) (*CreateQyWebSiteConfigReply, error)
	UpdateQyWebSiteConfig(ctx context.Context, in *UpdateQyWebSiteConfigRequest, opts ...grpc.CallOption) (*UpdateQyWebSiteConfigReply, error)
	DeleteQyWebSiteConfig(ctx context.Context, in *DeleteQyWebSiteConfigRequest, opts ...grpc.CallOption) (*DeleteQyWebSiteConfigReply, error)
	GetQyWebSiteConfig(ctx context.Context, in *GetQyWebSiteConfigRequest, opts ...grpc.CallOption) (*GetQyWebSiteConfigReply, error)
	ListQyWebSiteConfig(ctx context.Context, in *ListQyWebSiteConfigRequest, opts ...grpc.CallOption) (*ListQyWebSiteConfigReply, error)
	ListQyBaseSiteConfig(ctx context.Context, in *ListQyWebSiteConfigRequest, opts ...grpc.CallOption) (*ListQyWebSiteConfigReply, error)
}

type qyWebSiteConfigClient struct {
	cc grpc.ClientConnInterface
}

func NewQyWebSiteConfigClient(cc grpc.ClientConnInterface) QyWebSiteConfigClient {
	return &qyWebSiteConfigClient{cc}
}

func (c *qyWebSiteConfigClient) CreateQyWebSiteConfig(ctx context.Context, in *CreateQyWebSiteConfigRequest, opts ...grpc.CallOption) (*CreateQyWebSiteConfigReply, error) {
	out := new(CreateQyWebSiteConfigReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.web.v1.QyWebSiteConfig/CreateQyWebSiteConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyWebSiteConfigClient) UpdateQyWebSiteConfig(ctx context.Context, in *UpdateQyWebSiteConfigRequest, opts ...grpc.CallOption) (*UpdateQyWebSiteConfigReply, error) {
	out := new(UpdateQyWebSiteConfigReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.web.v1.QyWebSiteConfig/UpdateQyWebSiteConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyWebSiteConfigClient) DeleteQyWebSiteConfig(ctx context.Context, in *DeleteQyWebSiteConfigRequest, opts ...grpc.CallOption) (*DeleteQyWebSiteConfigReply, error) {
	out := new(DeleteQyWebSiteConfigReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.web.v1.QyWebSiteConfig/DeleteQyWebSiteConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyWebSiteConfigClient) GetQyWebSiteConfig(ctx context.Context, in *GetQyWebSiteConfigRequest, opts ...grpc.CallOption) (*GetQyWebSiteConfigReply, error) {
	out := new(GetQyWebSiteConfigReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.web.v1.QyWebSiteConfig/GetQyWebSiteConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyWebSiteConfigClient) ListQyWebSiteConfig(ctx context.Context, in *ListQyWebSiteConfigRequest, opts ...grpc.CallOption) (*ListQyWebSiteConfigReply, error) {
	out := new(ListQyWebSiteConfigReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.web.v1.QyWebSiteConfig/ListQyWebSiteConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyWebSiteConfigClient) ListQyBaseSiteConfig(ctx context.Context, in *ListQyWebSiteConfigRequest, opts ...grpc.CallOption) (*ListQyWebSiteConfigReply, error) {
	out := new(ListQyWebSiteConfigReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.web.v1.QyWebSiteConfig/ListQyBaseSiteConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QyWebSiteConfigServer is the server API for QyWebSiteConfig service.
// All implementations must embed UnimplementedQyWebSiteConfigServer
// for forward compatibility
type QyWebSiteConfigServer interface {
	CreateQyWebSiteConfig(context.Context, *CreateQyWebSiteConfigRequest) (*CreateQyWebSiteConfigReply, error)
	UpdateQyWebSiteConfig(context.Context, *UpdateQyWebSiteConfigRequest) (*UpdateQyWebSiteConfigReply, error)
	DeleteQyWebSiteConfig(context.Context, *DeleteQyWebSiteConfigRequest) (*DeleteQyWebSiteConfigReply, error)
	GetQyWebSiteConfig(context.Context, *GetQyWebSiteConfigRequest) (*GetQyWebSiteConfigReply, error)
	ListQyWebSiteConfig(context.Context, *ListQyWebSiteConfigRequest) (*ListQyWebSiteConfigReply, error)
	ListQyBaseSiteConfig(context.Context, *ListQyWebSiteConfigRequest) (*ListQyWebSiteConfigReply, error)
	mustEmbedUnimplementedQyWebSiteConfigServer()
}

// UnimplementedQyWebSiteConfigServer must be embedded to have forward compatible implementations.
type UnimplementedQyWebSiteConfigServer struct {
}

func (UnimplementedQyWebSiteConfigServer) CreateQyWebSiteConfig(context.Context, *CreateQyWebSiteConfigRequest) (*CreateQyWebSiteConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQyWebSiteConfig not implemented")
}
func (UnimplementedQyWebSiteConfigServer) UpdateQyWebSiteConfig(context.Context, *UpdateQyWebSiteConfigRequest) (*UpdateQyWebSiteConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQyWebSiteConfig not implemented")
}
func (UnimplementedQyWebSiteConfigServer) DeleteQyWebSiteConfig(context.Context, *DeleteQyWebSiteConfigRequest) (*DeleteQyWebSiteConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQyWebSiteConfig not implemented")
}
func (UnimplementedQyWebSiteConfigServer) GetQyWebSiteConfig(context.Context, *GetQyWebSiteConfigRequest) (*GetQyWebSiteConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQyWebSiteConfig not implemented")
}
func (UnimplementedQyWebSiteConfigServer) ListQyWebSiteConfig(context.Context, *ListQyWebSiteConfigRequest) (*ListQyWebSiteConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQyWebSiteConfig not implemented")
}
func (UnimplementedQyWebSiteConfigServer) ListQyBaseSiteConfig(context.Context, *ListQyWebSiteConfigRequest) (*ListQyWebSiteConfigReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQyBaseSiteConfig not implemented")
}
func (UnimplementedQyWebSiteConfigServer) mustEmbedUnimplementedQyWebSiteConfigServer() {}

// UnsafeQyWebSiteConfigServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QyWebSiteConfigServer will
// result in compilation errors.
type UnsafeQyWebSiteConfigServer interface {
	mustEmbedUnimplementedQyWebSiteConfigServer()
}

func RegisterQyWebSiteConfigServer(s grpc.ServiceRegistrar, srv QyWebSiteConfigServer) {
	s.RegisterService(&QyWebSiteConfig_ServiceDesc, srv)
}

func _QyWebSiteConfig_CreateQyWebSiteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQyWebSiteConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyWebSiteConfigServer).CreateQyWebSiteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.web.v1.QyWebSiteConfig/CreateQyWebSiteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyWebSiteConfigServer).CreateQyWebSiteConfig(ctx, req.(*CreateQyWebSiteConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyWebSiteConfig_UpdateQyWebSiteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQyWebSiteConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyWebSiteConfigServer).UpdateQyWebSiteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.web.v1.QyWebSiteConfig/UpdateQyWebSiteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyWebSiteConfigServer).UpdateQyWebSiteConfig(ctx, req.(*UpdateQyWebSiteConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyWebSiteConfig_DeleteQyWebSiteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQyWebSiteConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyWebSiteConfigServer).DeleteQyWebSiteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.web.v1.QyWebSiteConfig/DeleteQyWebSiteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyWebSiteConfigServer).DeleteQyWebSiteConfig(ctx, req.(*DeleteQyWebSiteConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyWebSiteConfig_GetQyWebSiteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQyWebSiteConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyWebSiteConfigServer).GetQyWebSiteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.web.v1.QyWebSiteConfig/GetQyWebSiteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyWebSiteConfigServer).GetQyWebSiteConfig(ctx, req.(*GetQyWebSiteConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyWebSiteConfig_ListQyWebSiteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQyWebSiteConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyWebSiteConfigServer).ListQyWebSiteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.web.v1.QyWebSiteConfig/ListQyWebSiteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyWebSiteConfigServer).ListQyWebSiteConfig(ctx, req.(*ListQyWebSiteConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyWebSiteConfig_ListQyBaseSiteConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQyWebSiteConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyWebSiteConfigServer).ListQyBaseSiteConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.web.v1.QyWebSiteConfig/ListQyBaseSiteConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyWebSiteConfigServer).ListQyBaseSiteConfig(ctx, req.(*ListQyWebSiteConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QyWebSiteConfig_ServiceDesc is the grpc.ServiceDesc for QyWebSiteConfig service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QyWebSiteConfig_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.qycms_bff.web.v1.QyWebSiteConfig",
	HandlerType: (*QyWebSiteConfigServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQyWebSiteConfig",
			Handler:    _QyWebSiteConfig_CreateQyWebSiteConfig_Handler,
		},
		{
			MethodName: "UpdateQyWebSiteConfig",
			Handler:    _QyWebSiteConfig_UpdateQyWebSiteConfig_Handler,
		},
		{
			MethodName: "DeleteQyWebSiteConfig",
			Handler:    _QyWebSiteConfig_DeleteQyWebSiteConfig_Handler,
		},
		{
			MethodName: "GetQyWebSiteConfig",
			Handler:    _QyWebSiteConfig_GetQyWebSiteConfig_Handler,
		},
		{
			MethodName: "ListQyWebSiteConfig",
			Handler:    _QyWebSiteConfig_ListQyWebSiteConfig_Handler,
		},
		{
			MethodName: "ListQyBaseSiteConfig",
			Handler:    _QyWebSiteConfig_ListQyBaseSiteConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/qycms_bff/web/v1/qy_web_site_config.proto",
}
