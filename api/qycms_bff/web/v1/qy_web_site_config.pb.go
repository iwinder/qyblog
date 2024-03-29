// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.1
// source: api/qycms_bff/web/v1/qy_web_site_config.proto

package v1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type SiteConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigKey   string `protobuf:"bytes,2,opt,name=configKey,proto3" json:"configKey,omitempty"`
	ConfigValue string `protobuf:"bytes,3,opt,name=configValue,proto3" json:"configValue,omitempty"`
}

func (x *SiteConfigResponse) Reset() {
	*x = SiteConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SiteConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SiteConfigResponse) ProtoMessage() {}

func (x *SiteConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SiteConfigResponse.ProtoReflect.Descriptor instead.
func (*SiteConfigResponse) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescGZIP(), []int{0}
}

func (x *SiteConfigResponse) GetConfigKey() string {
	if x != nil {
		return x.ConfigKey
	}
	return ""
}

func (x *SiteConfigResponse) GetConfigValue() string {
	if x != nil {
		return x.ConfigValue
	}
	return ""
}

type CreateQyWebSiteConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateQyWebSiteConfigRequest) Reset() {
	*x = CreateQyWebSiteConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQyWebSiteConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQyWebSiteConfigRequest) ProtoMessage() {}

func (x *CreateQyWebSiteConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQyWebSiteConfigRequest.ProtoReflect.Descriptor instead.
func (*CreateQyWebSiteConfigRequest) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescGZIP(), []int{1}
}

type CreateQyWebSiteConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateQyWebSiteConfigReply) Reset() {
	*x = CreateQyWebSiteConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateQyWebSiteConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateQyWebSiteConfigReply) ProtoMessage() {}

func (x *CreateQyWebSiteConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateQyWebSiteConfigReply.ProtoReflect.Descriptor instead.
func (*CreateQyWebSiteConfigReply) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescGZIP(), []int{2}
}

type UpdateQyWebSiteConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateQyWebSiteConfigRequest) Reset() {
	*x = UpdateQyWebSiteConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQyWebSiteConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQyWebSiteConfigRequest) ProtoMessage() {}

func (x *UpdateQyWebSiteConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQyWebSiteConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateQyWebSiteConfigRequest) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescGZIP(), []int{3}
}

type UpdateQyWebSiteConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateQyWebSiteConfigReply) Reset() {
	*x = UpdateQyWebSiteConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQyWebSiteConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQyWebSiteConfigReply) ProtoMessage() {}

func (x *UpdateQyWebSiteConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQyWebSiteConfigReply.ProtoReflect.Descriptor instead.
func (*UpdateQyWebSiteConfigReply) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescGZIP(), []int{4}
}

type DeleteQyWebSiteConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteQyWebSiteConfigRequest) Reset() {
	*x = DeleteQyWebSiteConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQyWebSiteConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQyWebSiteConfigRequest) ProtoMessage() {}

func (x *DeleteQyWebSiteConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQyWebSiteConfigRequest.ProtoReflect.Descriptor instead.
func (*DeleteQyWebSiteConfigRequest) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescGZIP(), []int{5}
}

type DeleteQyWebSiteConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteQyWebSiteConfigReply) Reset() {
	*x = DeleteQyWebSiteConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteQyWebSiteConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteQyWebSiteConfigReply) ProtoMessage() {}

func (x *DeleteQyWebSiteConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteQyWebSiteConfigReply.ProtoReflect.Descriptor instead.
func (*DeleteQyWebSiteConfigReply) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescGZIP(), []int{6}
}

type GetQyWebSiteConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetQyWebSiteConfigRequest) Reset() {
	*x = GetQyWebSiteConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQyWebSiteConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQyWebSiteConfigRequest) ProtoMessage() {}

func (x *GetQyWebSiteConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQyWebSiteConfigRequest.ProtoReflect.Descriptor instead.
func (*GetQyWebSiteConfigRequest) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescGZIP(), []int{7}
}

type GetQyWebSiteConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetQyWebSiteConfigReply) Reset() {
	*x = GetQyWebSiteConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetQyWebSiteConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetQyWebSiteConfigReply) ProtoMessage() {}

func (x *GetQyWebSiteConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetQyWebSiteConfigReply.ProtoReflect.Descriptor instead.
func (*GetQyWebSiteConfigReply) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescGZIP(), []int{8}
}

type ListQyWebSiteConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ftypes string `protobuf:"bytes,1,opt,name=ftypes,proto3" json:"ftypes,omitempty"`
}

func (x *ListQyWebSiteConfigRequest) Reset() {
	*x = ListQyWebSiteConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQyWebSiteConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQyWebSiteConfigRequest) ProtoMessage() {}

func (x *ListQyWebSiteConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQyWebSiteConfigRequest.ProtoReflect.Descriptor instead.
func (*ListQyWebSiteConfigRequest) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescGZIP(), []int{9}
}

func (x *ListQyWebSiteConfigRequest) GetFtypes() string {
	if x != nil {
		return x.Ftypes
	}
	return ""
}

type ListQyWebSiteConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Items []*SiteConfigResponse `protobuf:"bytes,1,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *ListQyWebSiteConfigReply) Reset() {
	*x = ListQyWebSiteConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListQyWebSiteConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListQyWebSiteConfigReply) ProtoMessage() {}

func (x *ListQyWebSiteConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListQyWebSiteConfigReply.ProtoReflect.Descriptor instead.
func (*ListQyWebSiteConfigReply) Descriptor() ([]byte, []int) {
	return file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescGZIP(), []int{10}
}

func (x *ListQyWebSiteConfigReply) GetItems() []*SiteConfigResponse {
	if x != nil {
		return x.Items
	}
	return nil
}

var File_api_qycms_bff_web_v1_qy_web_site_config_proto protoreflect.FileDescriptor

var file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2f,
	0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x71, 0x79, 0x5f, 0x77, 0x65, 0x62, 0x5f, 0x73, 0x69,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77,
	0x65, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x54, 0x0a, 0x12, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x4b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1e, 0x0a, 0x1c, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1c, 0x0a, 0x1a, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x0a, 0x1c, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1c, 0x0a, 0x1a, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1e, 0x0a, 0x1c, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51,
	0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x1c, 0x0a, 0x1a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51,
	0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x1b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53,
	0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x19, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x34, 0x0a, 0x1a, 0x4c,
	0x69, 0x73, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x74, 0x79, 0x70, 0x65,
	0x73, 0x22, 0x5a, 0x0a, 0x18, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x3e, 0x0a,
	0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x32, 0xbf, 0x07,
	0x0a, 0x0f, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x7d, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62,
	0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77,
	0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65,
	0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x7d, 0x0a, 0x15, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53,
	0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62,
	0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x7d, 0x0a, 0x15, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x32, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71,
	0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x30, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53,
	0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x74,
	0x0a, 0x12, 0x47, 0x65, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x12, 0x2f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73,
	0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x51,
	0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d,
	0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x77, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x79, 0x57, 0x65,
	0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x9d, 0x01,
	0x0a, 0x14, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x79, 0x42, 0x61, 0x73, 0x65, 0x53, 0x69, 0x74, 0x65,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63,
	0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71,
	0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x69, 0x73, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x23, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x1d,
	0x12, 0x1b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x69,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x12, 0x9f, 0x01,
	0x0a, 0x15, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x79, 0x4f, 0x74, 0x68, 0x65, 0x72, 0x53, 0x69, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x30, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79,
	0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x69, 0x73, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66, 0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31,
	0x2e, 0x4c, 0x69, 0x73, 0x74, 0x51, 0x79, 0x57, 0x65, 0x62, 0x53, 0x69, 0x74, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x24, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x1e, 0x12, 0x1c, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x73,
	0x69, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x6f, 0x74, 0x68, 0x65, 0x72, 0x42,
	0x4e, 0x0a, 0x14, 0x61, 0x70, 0x69, 0x2e, 0x71, 0x79, 0x63, 0x6d, 0x73, 0x5f, 0x62, 0x66, 0x66,
	0x2e, 0x77, 0x65, 0x62, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x69, 0x77, 0x69, 0x6e, 0x64, 0x65, 0x72, 0x2f, 0x71, 0x69,
	0x6e, 0x67, 0x79, 0x75, 0x63, 0x6d, 0x73, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x71, 0x79, 0x63, 0x6d,
	0x73, 0x5f, 0x62, 0x66, 0x66, 0x2f, 0x77, 0x65, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescOnce sync.Once
	file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescData = file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDesc
)

func file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescGZIP() []byte {
	file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescOnce.Do(func() {
		file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescData)
	})
	return file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDescData
}

var file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_api_qycms_bff_web_v1_qy_web_site_config_proto_goTypes = []interface{}{
	(*SiteConfigResponse)(nil),           // 0: api.qycms_bff.web.v1.SiteConfigResponse
	(*CreateQyWebSiteConfigRequest)(nil), // 1: api.qycms_bff.web.v1.CreateQyWebSiteConfigRequest
	(*CreateQyWebSiteConfigReply)(nil),   // 2: api.qycms_bff.web.v1.CreateQyWebSiteConfigReply
	(*UpdateQyWebSiteConfigRequest)(nil), // 3: api.qycms_bff.web.v1.UpdateQyWebSiteConfigRequest
	(*UpdateQyWebSiteConfigReply)(nil),   // 4: api.qycms_bff.web.v1.UpdateQyWebSiteConfigReply
	(*DeleteQyWebSiteConfigRequest)(nil), // 5: api.qycms_bff.web.v1.DeleteQyWebSiteConfigRequest
	(*DeleteQyWebSiteConfigReply)(nil),   // 6: api.qycms_bff.web.v1.DeleteQyWebSiteConfigReply
	(*GetQyWebSiteConfigRequest)(nil),    // 7: api.qycms_bff.web.v1.GetQyWebSiteConfigRequest
	(*GetQyWebSiteConfigReply)(nil),      // 8: api.qycms_bff.web.v1.GetQyWebSiteConfigReply
	(*ListQyWebSiteConfigRequest)(nil),   // 9: api.qycms_bff.web.v1.ListQyWebSiteConfigRequest
	(*ListQyWebSiteConfigReply)(nil),     // 10: api.qycms_bff.web.v1.ListQyWebSiteConfigReply
}
var file_api_qycms_bff_web_v1_qy_web_site_config_proto_depIdxs = []int32{
	0,  // 0: api.qycms_bff.web.v1.ListQyWebSiteConfigReply.items:type_name -> api.qycms_bff.web.v1.SiteConfigResponse
	1,  // 1: api.qycms_bff.web.v1.QyWebSiteConfig.CreateQyWebSiteConfig:input_type -> api.qycms_bff.web.v1.CreateQyWebSiteConfigRequest
	3,  // 2: api.qycms_bff.web.v1.QyWebSiteConfig.UpdateQyWebSiteConfig:input_type -> api.qycms_bff.web.v1.UpdateQyWebSiteConfigRequest
	5,  // 3: api.qycms_bff.web.v1.QyWebSiteConfig.DeleteQyWebSiteConfig:input_type -> api.qycms_bff.web.v1.DeleteQyWebSiteConfigRequest
	7,  // 4: api.qycms_bff.web.v1.QyWebSiteConfig.GetQyWebSiteConfig:input_type -> api.qycms_bff.web.v1.GetQyWebSiteConfigRequest
	9,  // 5: api.qycms_bff.web.v1.QyWebSiteConfig.ListQyWebSiteConfig:input_type -> api.qycms_bff.web.v1.ListQyWebSiteConfigRequest
	9,  // 6: api.qycms_bff.web.v1.QyWebSiteConfig.ListQyBaseSiteConfig:input_type -> api.qycms_bff.web.v1.ListQyWebSiteConfigRequest
	9,  // 7: api.qycms_bff.web.v1.QyWebSiteConfig.ListQyOtherSiteConfig:input_type -> api.qycms_bff.web.v1.ListQyWebSiteConfigRequest
	2,  // 8: api.qycms_bff.web.v1.QyWebSiteConfig.CreateQyWebSiteConfig:output_type -> api.qycms_bff.web.v1.CreateQyWebSiteConfigReply
	4,  // 9: api.qycms_bff.web.v1.QyWebSiteConfig.UpdateQyWebSiteConfig:output_type -> api.qycms_bff.web.v1.UpdateQyWebSiteConfigReply
	6,  // 10: api.qycms_bff.web.v1.QyWebSiteConfig.DeleteQyWebSiteConfig:output_type -> api.qycms_bff.web.v1.DeleteQyWebSiteConfigReply
	8,  // 11: api.qycms_bff.web.v1.QyWebSiteConfig.GetQyWebSiteConfig:output_type -> api.qycms_bff.web.v1.GetQyWebSiteConfigReply
	10, // 12: api.qycms_bff.web.v1.QyWebSiteConfig.ListQyWebSiteConfig:output_type -> api.qycms_bff.web.v1.ListQyWebSiteConfigReply
	10, // 13: api.qycms_bff.web.v1.QyWebSiteConfig.ListQyBaseSiteConfig:output_type -> api.qycms_bff.web.v1.ListQyWebSiteConfigReply
	10, // 14: api.qycms_bff.web.v1.QyWebSiteConfig.ListQyOtherSiteConfig:output_type -> api.qycms_bff.web.v1.ListQyWebSiteConfigReply
	8,  // [8:15] is the sub-list for method output_type
	1,  // [1:8] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_api_qycms_bff_web_v1_qy_web_site_config_proto_init() }
func file_api_qycms_bff_web_v1_qy_web_site_config_proto_init() {
	if File_api_qycms_bff_web_v1_qy_web_site_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SiteConfigResponse); i {
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
		file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQyWebSiteConfigRequest); i {
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
		file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateQyWebSiteConfigReply); i {
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
		file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQyWebSiteConfigRequest); i {
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
		file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQyWebSiteConfigReply); i {
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
		file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQyWebSiteConfigRequest); i {
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
		file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteQyWebSiteConfigReply); i {
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
		file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQyWebSiteConfigRequest); i {
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
		file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetQyWebSiteConfigReply); i {
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
		file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQyWebSiteConfigRequest); i {
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
		file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListQyWebSiteConfigReply); i {
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
			RawDescriptor: file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_qycms_bff_web_v1_qy_web_site_config_proto_goTypes,
		DependencyIndexes: file_api_qycms_bff_web_v1_qy_web_site_config_proto_depIdxs,
		MessageInfos:      file_api_qycms_bff_web_v1_qy_web_site_config_proto_msgTypes,
	}.Build()
	File_api_qycms_bff_web_v1_qy_web_site_config_proto = out.File
	file_api_qycms_bff_web_v1_qy_web_site_config_proto_rawDesc = nil
	file_api_qycms_bff_web_v1_qy_web_site_config_proto_goTypes = nil
	file_api_qycms_bff_web_v1_qy_web_site_config_proto_depIdxs = nil
}
