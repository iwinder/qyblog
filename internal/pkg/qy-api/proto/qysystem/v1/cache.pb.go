package v1

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/runtime/protoimpl"
)

type ListSecretsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset *int64 `json:"offset,omitempty" protobuf:"varint,1,opt,name=offset,proto3,oneof"`
	Limit  *int64 `json:"limit,omitempty" protobuf:"varint,2,opt,name=limit,proto3,oneof"`
}

type ListPoliciesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	offset *int64 `json:"offset,omitempty" protobuf:"varint,1,opt,name=offset,proto3,oneof"`
	Limit  *int64 `json:"limit,omitempty" protobuf:"varint,2,opt,name=limit,proto3,oneof"`
}

type SecretInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name        string `json:"name,omitempty" protobuf:"bytes,1,opt,name=name,proto3"`
	SecretId    string `json:"secret_id,omitempty" protobuf:"bytes,2,opt,name=secret_id,json=secretId,proto3"`
	Username    string `json:"username,omitempty" protobuf:"bytes,3,opt,name=username,proto3"`
	SecretKey   string `json:"secret_key" protobuf:"bytes,4,opt,name=secret_key,json=secretKey,proto3"`
	Expires     int64  `json:"expires" protobuf:"varint,5,opt,name=expires,proto3"`
	Description string `json:"description" protobuf:"bytes,6,opt,name=description,proto3"`
	CreatedAt   string `json:"createdAt" protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3"`
	UpdatedAt   string `json:"updatedAt" protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3"`
}

type ListSecretsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64         `json:"total_count,omitempty" protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3"`
	Items      []*SecretInfo `json:"items,omitempty" protobuf:"bytes,2,rep,name=items,proto3"`
}

type ListPoliciesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TotalCount int64         `json:"total_count,omitempty" protobuf:"varint,1,opt,name=total_count,json=totalCount,proto3"`
	Items      []*SecretInfo `json:"items,omitempty" protobuf:"bytes,2,rep,name=items,proto3"`
}

type CacheClient interface {
	ListSecrets(ctx context.Context, in *ListSecretsRequest, opts ...grpc.CallOption) (*ListSecretsResponse, error)
	ListPolicies(ctx context.Context, in *ListPoliciesRequest, opts ...grpc.CallOption) (*ListPoliciesResponse, error)
}
type CacheServer interface {
	ListSecrets(context.Context, *ListSecretsRequest) (*ListSecretsResponse, error)
	ListPolicies(context.Context, *ListPoliciesRequest) (*ListPoliciesResponse, error)
}

func RegisterCacheServer(s *grpc.Server, srv CacheServer) {
	s.RegisterService(&_Cache_serviceDesc, srv)
}

func _Cache_ListSecrets_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSecretsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).ListSecrets(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Cache/ListSecrets",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).ListSecrets(ctx, req.(*ListSecretsRequest))
	}
	return interceptor(ctx, in, info, handler)
}
func _Cache_ListPolicies_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListPoliciesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).ListPolicies(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Cache/ListPolicies",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).ListPolicies(ctx, req.(*ListPoliciesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListSecrets",
			Handler:    _Cache_ListSecrets_Handler,
		}, {
			MethodName: "ListPolicies",
			Handler:    _Cache_ListPolicies_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/apiserver/v1/cache.proto",
}
