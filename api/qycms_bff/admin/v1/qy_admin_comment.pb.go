// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: api/qycms_bff/admin/v1/qy_admin_comment.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type CreateQyAdminCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateQyAdminCommentRequest) Reset() {
	*x = CreateQyAdminCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQyAdminCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQyAdminCommentRequest) ProtoMessage() {}

func (x *CreateQyAdminCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQyAdminCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateQyAdminCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescGZIP(), []int{0}
}

type CreateQyAdminCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateQyAdminCommentReply) Reset() {
	*x = CreateQyAdminCommentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQyAdminCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQyAdminCommentReply) ProtoMessage() {}

func (x *CreateQyAdminCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQyAdminCommentReply.ProtoReflect.Descriptor instead.
func (*CreateQyAdminCommentReply) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescGZIP(), []int{1}
}

type UpdateQyAdminCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateQyAdminCommentRequest) Reset() {
	*x = UpdateQyAdminCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQyAdminCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQyAdminCommentRequest) ProtoMessage() {}

func (x *UpdateQyAdminCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQyAdminCommentRequest.ProtoReflect.Descriptor instead.
func (*UpdateQyAdminCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescGZIP(), []int{2}
}

type UpdateQyAdminCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateQyAdminCommentReply) Reset() {
	*x = UpdateQyAdminCommentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQyAdminCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQyAdminCommentReply) ProtoMessage() {}

func (x *UpdateQyAdminCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQyAdminCommentReply.ProtoReflect.Descriptor instead.
func (*UpdateQyAdminCommentReply) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescGZIP(), []int{3}
}

type DeleteQyAdminCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteQyAdminCommentRequest) Reset() {
	*x = DeleteQyAdminCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQyAdminCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQyAdminCommentRequest) ProtoMessage() {}

func (x *DeleteQyAdminCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQyAdminCommentRequest.ProtoReflect.Descriptor instead.
func (*DeleteQyAdminCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescGZIP(), []int{4}
}

type DeleteQyAdminCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteQyAdminCommentReply) Reset() {
	*x = DeleteQyAdminCommentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQyAdminCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQyAdminCommentReply) ProtoMessage() {}

func (x *DeleteQyAdminCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQyAdminCommentReply.ProtoReflect.Descriptor instead.
func (*DeleteQyAdminCommentReply) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescGZIP(), []int{5}
}

type GetQyAdminCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetQyAdminCommentRequest) Reset() {
	*x = GetQyAdminCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQyAdminCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQyAdminCommentRequest) ProtoMessage() {}

func (x *GetQyAdminCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQyAdminCommentRequest.ProtoReflect.Descriptor instead.
func (*GetQyAdminCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescGZIP(), []int{6}
}

type GetQyAdminCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetQyAdminCommentReply) Reset() {
	*x = GetQyAdminCommentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQyAdminCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQyAdminCommentReply) ProtoMessage() {}

func (x *GetQyAdminCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQyAdminCommentReply.ProtoReflect.Descriptor instead.
func (*GetQyAdminCommentReply) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescGZIP(), []int{7}
}

type ListQyAdminCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListQyAdminCommentRequest) Reset() {
	*x = ListQyAdminCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQyAdminCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQyAdminCommentRequest) ProtoMessage() {}

func (x *ListQyAdminCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQyAdminCommentRequest.ProtoReflect.Descriptor instead.
func (*ListQyAdminCommentRequest) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescGZIP(), []int{8}
}

type ListQyAdminCommentReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListQyAdminCommentReply) Reset() {
	*x = ListQyAdminCommentReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQyAdminCommentReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQyAdminCommentReply) ProtoMessage() {}

func (x *ListQyAdminCommentReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQyAdminCommentReply.ProtoReflect.Descriptor instead.
func (*ListQyAdminCommentReply) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescGZIP(), []int{9}
}

var File_api_qycms_bff_admin_v1_qy_admin_comment_proto protoreflect.FileDescriptor

var file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x79, 0x5f, 0x61, 0x64, 0x6d, 0x69,
	0x6e, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x16, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x22, 0x1d, 0x0a, 0x1b, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1b, 0x0a, 0x19, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1d, 0x0a, 0x1b, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x79, 0x41,
	0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x1b, 0x0a, 0x19, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x79, 0x41, 0x64,
	0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x1d, 0x0a, 0x1b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1b,
	0x0a, 0x19, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x0a, 0x18, 0x47,
	0x65, 0x74, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x18, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x51, 0x79,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x22, 0x1b, 0x0a, 0x19, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x19,
	0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0x81, 0x05, 0x0a, 0x0e, 0x51, 0x79,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x7e, 0x0a, 0x14,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73,
	0x5f, 0x62, 0x66, 0x66, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x7e, 0x0a, 0x14,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73,
	0x5f, 0x62, 0x66, 0x66, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x7e, 0x0a, 0x14,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x12, 0x33, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73,
	0x5f, 0x62, 0x66, 0x66, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x75, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66,
	0x66, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51, 0x79,
	0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f,
	0x62, 0x66, 0x66, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x12, 0x78, 0x0a, 0x12, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x79, 0x41, 0x64, 0x6d,
	0x69, 0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x43, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x79, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x42, 0x52, 0x0a,
	0x16, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x61,
	0x64, 0x6d, 0x69, 0x6e, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x36, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x77, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x71, 0x69,
	0x6e, 0x67, 0x79, 0x75, 0x63, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x79, 0x63, 0x6d,
	0x73, 0x5f, 0x62, 0x66, 0x66, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x76,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescOnce sync.Once
	file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescData = file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDesc
)

func file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescGZIP() []byte {
	file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescOnce.Do(func() {
		file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescData)
	})
	return file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDescData
}

var file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_api_qycms_bff_admin_v1_qy_admin_comment_proto_goTypes = []interface{}{
	(*CreateQyAdminCommentRequest)(nil), // 0: api.qycms_bff.admin.v1.CreateQyAdminCommentRequest
	(*CreateQyAdminCommentReply)(nil),   // 1: api.qycms_bff.admin.v1.CreateQyAdminCommentReply
	(*UpdateQyAdminCommentRequest)(nil), // 2: api.qycms_bff.admin.v1.UpdateQyAdminCommentRequest
	(*UpdateQyAdminCommentReply)(nil),   // 3: api.qycms_bff.admin.v1.UpdateQyAdminCommentReply
	(*DeleteQyAdminCommentRequest)(nil), // 4: api.qycms_bff.admin.v1.DeleteQyAdminCommentRequest
	(*DeleteQyAdminCommentReply)(nil),   // 5: api.qycms_bff.admin.v1.DeleteQyAdminCommentReply
	(*GetQyAdminCommentRequest)(nil),    // 6: api.qycms_bff.admin.v1.GetQyAdminCommentRequest
	(*GetQyAdminCommentReply)(nil),      // 7: api.qycms_bff.admin.v1.GetQyAdminCommentReply
	(*ListQyAdminCommentRequest)(nil),   // 8: api.qycms_bff.admin.v1.ListQyAdminCommentRequest
	(*ListQyAdminCommentReply)(nil),     // 9: api.qycms_bff.admin.v1.ListQyAdminCommentReply
}
var file_api_qycms_bff_admin_v1_qy_admin_comment_proto_depIdxs = []int32{
	0, // 0: api.qycms_bff.admin.v1.QyAdminComment.CreateQyAdminComment:input_type -> api.qycms_bff.admin.v1.CreateQyAdminCommentRequest
	2, // 1: api.qycms_bff.admin.v1.QyAdminComment.UpdateQyAdminComment:input_type -> api.qycms_bff.admin.v1.UpdateQyAdminCommentRequest
	4, // 2: api.qycms_bff.admin.v1.QyAdminComment.DeleteQyAdminComment:input_type -> api.qycms_bff.admin.v1.DeleteQyAdminCommentRequest
	6, // 3: api.qycms_bff.admin.v1.QyAdminComment.GetQyAdminComment:input_type -> api.qycms_bff.admin.v1.GetQyAdminCommentRequest
	8, // 4: api.qycms_bff.admin.v1.QyAdminComment.ListQyAdminComment:input_type -> api.qycms_bff.admin.v1.ListQyAdminCommentRequest
	1, // 5: api.qycms_bff.admin.v1.QyAdminComment.CreateQyAdminComment:output_type -> api.qycms_bff.admin.v1.CreateQyAdminCommentReply
	3, // 6: api.qycms_bff.admin.v1.QyAdminComment.UpdateQyAdminComment:output_type -> api.qycms_bff.admin.v1.UpdateQyAdminCommentReply
	5, // 7: api.qycms_bff.admin.v1.QyAdminComment.DeleteQyAdminComment:output_type -> api.qycms_bff.admin.v1.DeleteQyAdminCommentReply
	7, // 8: api.qycms_bff.admin.v1.QyAdminComment.GetQyAdminComment:output_type -> api.qycms_bff.admin.v1.GetQyAdminCommentReply
	9, // 9: api.qycms_bff.admin.v1.QyAdminComment.ListQyAdminComment:output_type -> api.qycms_bff.admin.v1.ListQyAdminCommentReply
	5, // [5:10] is the sub-list for method output_type
	0, // [0:5] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_qycms_bff_admin_v1_qy_admin_comment_proto_init() }
func file_api_qycms_bff_admin_v1_qy_admin_comment_proto_init() {
	if File_api_qycms_bff_admin_v1_qy_admin_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQyAdminCommentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQyAdminCommentReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQyAdminCommentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQyAdminCommentReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQyAdminCommentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQyAdminCommentReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQyAdminCommentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQyAdminCommentReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQyAdminCommentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQyAdminCommentReply); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_qycms_bff_admin_v1_qy_admin_comment_proto_goTypes,
		DependencyIndexes: file_api_qycms_bff_admin_v1_qy_admin_comment_proto_depIdxs,
		MessageInfos:      file_api_qycms_bff_admin_v1_qy_admin_comment_proto_msgTypes,
	}.Build()
	File_api_qycms_bff_admin_v1_qy_admin_comment_proto = out.File
	file_api_qycms_bff_admin_v1_qy_admin_comment_proto_rawDesc = nil
	file_api_qycms_bff_admin_v1_qy_admin_comment_proto_goTypes = nil
	file_api_qycms_bff_admin_v1_qy_admin_comment_proto_depIdxs = nil
}