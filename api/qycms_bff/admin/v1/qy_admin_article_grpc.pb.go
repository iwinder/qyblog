// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.1
// source: api/qycms_bff/admin/v1/qy_admin_article.proto

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

// QyAdminArticleClient is the client API for QyAdminArticle service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QyAdminArticleClient interface {
	// 新增
	CreateQyAdminArticle(ctx context.Context, in *CreateQyAdminArticleRequest, opts ...grpc.CallOption) (*CreateQyAdminArticleReply, error)
	// 更新
	UpdateQyAdminArticle(ctx context.Context, in *UpdateQyAdminArticleRequest, opts ...grpc.CallOption) (*UpdateQyAdminArticleReply, error)
	// 删除
	DeleteQyAdminArticle(ctx context.Context, in *DeleteQyAdminArticleRequest, opts ...grpc.CallOption) (*DeleteQyAdminArticleReply, error)
	// 获取详情
	GetQyAdminArticle(ctx context.Context, in *GetQyAdminArticleRequest, opts ...grpc.CallOption) (*GetQyAdminArticleReply, error)
	// 生成文章链接
	InitQyAdminArticlePermaLink(ctx context.Context, in *InitQyAdminArticlePermaLinkRequest, opts ...grpc.CallOption) (*InitQyAdminArticlePermaLinkReply, error)
	// 列表
	ListQyAdminArticle(ctx context.Context, in *ListQyAdminArticleRequest, opts ...grpc.CallOption) (*ListQyAdminArticleReply, error)
}

type qyAdminArticleClient struct {
	cc grpc.ClientConnInterface
}

func NewQyAdminArticleClient(cc grpc.ClientConnInterface) QyAdminArticleClient {
	return &qyAdminArticleClient{cc}
}

func (c *qyAdminArticleClient) CreateQyAdminArticle(ctx context.Context, in *CreateQyAdminArticleRequest, opts ...grpc.CallOption) (*CreateQyAdminArticleReply, error) {
	out := new(CreateQyAdminArticleReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminArticle/CreateQyAdminArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminArticleClient) UpdateQyAdminArticle(ctx context.Context, in *UpdateQyAdminArticleRequest, opts ...grpc.CallOption) (*UpdateQyAdminArticleReply, error) {
	out := new(UpdateQyAdminArticleReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminArticle/UpdateQyAdminArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminArticleClient) DeleteQyAdminArticle(ctx context.Context, in *DeleteQyAdminArticleRequest, opts ...grpc.CallOption) (*DeleteQyAdminArticleReply, error) {
	out := new(DeleteQyAdminArticleReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminArticle/DeleteQyAdminArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminArticleClient) GetQyAdminArticle(ctx context.Context, in *GetQyAdminArticleRequest, opts ...grpc.CallOption) (*GetQyAdminArticleReply, error) {
	out := new(GetQyAdminArticleReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminArticle/GetQyAdminArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminArticleClient) InitQyAdminArticlePermaLink(ctx context.Context, in *InitQyAdminArticlePermaLinkRequest, opts ...grpc.CallOption) (*InitQyAdminArticlePermaLinkReply, error) {
	out := new(InitQyAdminArticlePermaLinkReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminArticle/InitQyAdminArticlePermaLink", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *qyAdminArticleClient) ListQyAdminArticle(ctx context.Context, in *ListQyAdminArticleRequest, opts ...grpc.CallOption) (*ListQyAdminArticleReply, error) {
	out := new(ListQyAdminArticleReply)
	err := c.cc.Invoke(ctx, "/api.qycms_bff.admin.v1.QyAdminArticle/ListQyAdminArticle", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QyAdminArticleServer is the server API for QyAdminArticle service.
// All implementations must embed UnimplementedQyAdminArticleServer
// for forward compatibility
type QyAdminArticleServer interface {
	// 新增
	CreateQyAdminArticle(context.Context, *CreateQyAdminArticleRequest) (*CreateQyAdminArticleReply, error)
	// 更新
	UpdateQyAdminArticle(context.Context, *UpdateQyAdminArticleRequest) (*UpdateQyAdminArticleReply, error)
	// 删除
	DeleteQyAdminArticle(context.Context, *DeleteQyAdminArticleRequest) (*DeleteQyAdminArticleReply, error)
	// 获取详情
	GetQyAdminArticle(context.Context, *GetQyAdminArticleRequest) (*GetQyAdminArticleReply, error)
	// 生成文章链接
	InitQyAdminArticlePermaLink(context.Context, *InitQyAdminArticlePermaLinkRequest) (*InitQyAdminArticlePermaLinkReply, error)
	// 列表
	ListQyAdminArticle(context.Context, *ListQyAdminArticleRequest) (*ListQyAdminArticleReply, error)
	mustEmbedUnimplementedQyAdminArticleServer()
}

// UnimplementedQyAdminArticleServer must be embedded to have forward compatible implementations.
type UnimplementedQyAdminArticleServer struct {
}

func (UnimplementedQyAdminArticleServer) CreateQyAdminArticle(context.Context, *CreateQyAdminArticleRequest) (*CreateQyAdminArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateQyAdminArticle not implemented")
}
func (UnimplementedQyAdminArticleServer) UpdateQyAdminArticle(context.Context, *UpdateQyAdminArticleRequest) (*UpdateQyAdminArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQyAdminArticle not implemented")
}
func (UnimplementedQyAdminArticleServer) DeleteQyAdminArticle(context.Context, *DeleteQyAdminArticleRequest) (*DeleteQyAdminArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteQyAdminArticle not implemented")
}
func (UnimplementedQyAdminArticleServer) GetQyAdminArticle(context.Context, *GetQyAdminArticleRequest) (*GetQyAdminArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQyAdminArticle not implemented")
}
func (UnimplementedQyAdminArticleServer) InitQyAdminArticlePermaLink(context.Context, *InitQyAdminArticlePermaLinkRequest) (*InitQyAdminArticlePermaLinkReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InitQyAdminArticlePermaLink not implemented")
}
func (UnimplementedQyAdminArticleServer) ListQyAdminArticle(context.Context, *ListQyAdminArticleRequest) (*ListQyAdminArticleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListQyAdminArticle not implemented")
}
func (UnimplementedQyAdminArticleServer) mustEmbedUnimplementedQyAdminArticleServer() {}

// UnsafeQyAdminArticleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QyAdminArticleServer will
// result in compilation errors.
type UnsafeQyAdminArticleServer interface {
	mustEmbedUnimplementedQyAdminArticleServer()
}

func RegisterQyAdminArticleServer(s grpc.ServiceRegistrar, srv QyAdminArticleServer) {
	s.RegisterService(&QyAdminArticle_ServiceDesc, srv)
}

func _QyAdminArticle_CreateQyAdminArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateQyAdminArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminArticleServer).CreateQyAdminArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminArticle/CreateQyAdminArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminArticleServer).CreateQyAdminArticle(ctx, req.(*CreateQyAdminArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminArticle_UpdateQyAdminArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQyAdminArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminArticleServer).UpdateQyAdminArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminArticle/UpdateQyAdminArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminArticleServer).UpdateQyAdminArticle(ctx, req.(*UpdateQyAdminArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminArticle_DeleteQyAdminArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteQyAdminArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminArticleServer).DeleteQyAdminArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminArticle/DeleteQyAdminArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminArticleServer).DeleteQyAdminArticle(ctx, req.(*DeleteQyAdminArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminArticle_GetQyAdminArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQyAdminArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminArticleServer).GetQyAdminArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminArticle/GetQyAdminArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminArticleServer).GetQyAdminArticle(ctx, req.(*GetQyAdminArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminArticle_InitQyAdminArticlePermaLink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InitQyAdminArticlePermaLinkRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminArticleServer).InitQyAdminArticlePermaLink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminArticle/InitQyAdminArticlePermaLink",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminArticleServer).InitQyAdminArticlePermaLink(ctx, req.(*InitQyAdminArticlePermaLinkRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QyAdminArticle_ListQyAdminArticle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListQyAdminArticleRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QyAdminArticleServer).ListQyAdminArticle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.qycms_bff.admin.v1.QyAdminArticle/ListQyAdminArticle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QyAdminArticleServer).ListQyAdminArticle(ctx, req.(*ListQyAdminArticleRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QyAdminArticle_ServiceDesc is the grpc.ServiceDesc for QyAdminArticle service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QyAdminArticle_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.qycms_bff.admin.v1.QyAdminArticle",
	HandlerType: (*QyAdminArticleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateQyAdminArticle",
			Handler:    _QyAdminArticle_CreateQyAdminArticle_Handler,
		},
		{
			MethodName: "UpdateQyAdminArticle",
			Handler:    _QyAdminArticle_UpdateQyAdminArticle_Handler,
		},
		{
			MethodName: "DeleteQyAdminArticle",
			Handler:    _QyAdminArticle_DeleteQyAdminArticle_Handler,
		},
		{
			MethodName: "GetQyAdminArticle",
			Handler:    _QyAdminArticle_GetQyAdminArticle_Handler,
		},
		{
			MethodName: "InitQyAdminArticlePermaLink",
			Handler:    _QyAdminArticle_InitQyAdminArticlePermaLink_Handler,
		},
		{
			MethodName: "ListQyAdminArticle",
			Handler:    _QyAdminArticle_ListQyAdminArticle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/qycms_bff/admin/v1/qy_admin_article.proto",
}
