// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: system/v1/sys_config.proto

package v1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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

type CreateSysConfigRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateSysConfigRep) Reset() {
	*x = CreateSysConfigRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_config_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSysConfigRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSysConfigRep) ProtoMessage() {}

func (x *CreateSysConfigRep) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_config_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSysConfigRep.ProtoReflect.Descriptor instead.
func (*CreateSysConfigRep) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_config_proto_rawDescGZIP(), []int{0}
}

type EmptySysConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EmptySysConfigReply) Reset() {
	*x = EmptySysConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_config_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EmptySysConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EmptySysConfigReply) ProtoMessage() {}

func (x *EmptySysConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_config_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EmptySysConfigReply.ProtoReflect.Descriptor instead.
func (*EmptySysConfigReply) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_config_proto_rawDescGZIP(), []int{1}
}

type UpdateSysConfigRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateSysConfigRep) Reset() {
	*x = UpdateSysConfigRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_config_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSysConfigRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSysConfigRep) ProtoMessage() {}

func (x *UpdateSysConfigRep) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_config_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSysConfigRep.ProtoReflect.Descriptor instead.
func (*UpdateSysConfigRep) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_config_proto_rawDescGZIP(), []int{2}
}

type DeleteSysConfigRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteSysConfigRep) Reset() {
	*x = DeleteSysConfigRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_config_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSysConfigRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSysConfigRep) ProtoMessage() {}

func (x *DeleteSysConfigRep) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_config_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSysConfigRep.ProtoReflect.Descriptor instead.
func (*DeleteSysConfigRep) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_config_proto_rawDescGZIP(), []int{3}
}

type GetSysConfigRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *GetSysConfigRep) Reset() {
	*x = GetSysConfigRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_config_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSysConfigRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSysConfigRep) ProtoMessage() {}

func (x *GetSysConfigRep) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_config_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSysConfigRep.ProtoReflect.Descriptor instead.
func (*GetSysConfigRep) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_config_proto_rawDescGZIP(), []int{4}
}

func (x *GetSysConfigRep) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSysConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *ConfigReply `protobuf:"bytes,1,opt,name=config,proto3" json:"config"`
}

func (x *GetSysConfigReply) Reset() {
	*x = GetSysConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_config_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSysConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSysConfigReply) ProtoMessage() {}

func (x *GetSysConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_config_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSysConfigReply.ProtoReflect.Descriptor instead.
func (*GetSysConfigReply) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_config_proto_rawDescGZIP(), []int{5}
}

func (x *GetSysConfigReply) GetConfig() *ConfigReply {
	if x != nil {
		return x.Config
	}
	return nil
}

type ListSysConfigRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pageNum
	PageNum int64 `protobuf:"varint,1,opt,name=pageNum,proto3" json:"pageNum"`
	// 页面大小
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize"`
}

func (x *ListSysConfigRep) Reset() {
	*x = ListSysConfigRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_config_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSysConfigRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSysConfigRep) ProtoMessage() {}

func (x *ListSysConfigRep) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_config_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSysConfigRep.ProtoReflect.Descriptor instead.
func (*ListSysConfigRep) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_config_proto_rawDescGZIP(), []int{6}
}

func (x *ListSysConfigRep) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListSysConfigRep) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListSysConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Rows  []*ConfigReply `protobuf:"bytes,1,rep,name=rows,proto3" json:"rows"`
	Total int64          `protobuf:"varint,2,opt,name=total,proto3" json:"total"`
}

func (x *ListSysConfigReply) Reset() {
	*x = ListSysConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_config_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSysConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSysConfigReply) ProtoMessage() {}

func (x *ListSysConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_config_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSysConfigReply.ProtoReflect.Descriptor instead.
func (*ListSysConfigReply) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_config_proto_rawDescGZIP(), []int{7}
}

func (x *ListSysConfigReply) GetRows() []*ConfigReply {
	if x != nil {
		return x.Rows
	}
	return nil
}

func (x *ListSysConfigReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

type ConfigByKeyReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key"`
}

func (x *ConfigByKeyReq) Reset() {
	*x = ConfigByKeyReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_config_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigByKeyReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigByKeyReq) ProtoMessage() {}

func (x *ConfigByKeyReq) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_config_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigByKeyReq.ProtoReflect.Descriptor instead.
func (*ConfigByKeyReq) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_config_proto_rawDescGZIP(), []int{8}
}

func (x *ConfigByKeyReq) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

type ConfigByKeyReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value"`
}

func (x *ConfigByKeyReply) Reset() {
	*x = ConfigByKeyReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_config_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigByKeyReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigByKeyReply) ProtoMessage() {}

func (x *ConfigByKeyReply) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_config_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigByKeyReply.ProtoReflect.Descriptor instead.
func (*ConfigByKeyReply) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_config_proto_rawDescGZIP(), []int{9}
}

func (x *ConfigByKeyReply) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ConfigReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigId    string `protobuf:"bytes,1,opt,name=configId,proto3" json:"configId"`
	ConfigName  string `protobuf:"bytes,2,opt,name=configName,proto3" json:"configName"`
	ConfigKey   string `protobuf:"bytes,3,opt,name=configKey,proto3" json:"configKey"`
	ConfigValue string `protobuf:"bytes,4,opt,name=configValue,proto3" json:"configValue"`
	ConfigType  string `protobuf:"bytes,5,opt,name=configType,proto3" json:"configType"`
	Remark      string `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark"`
	CreatedAt   string `protobuf:"bytes,7,opt,name=createdAt,proto3" json:"createdAt"`
}

func (x *ConfigReply) Reset() {
	*x = ConfigReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_config_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ConfigReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ConfigReply) ProtoMessage() {}

func (x *ConfigReply) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_config_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ConfigReply.ProtoReflect.Descriptor instead.
func (*ConfigReply) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_config_proto_rawDescGZIP(), []int{10}
}

func (x *ConfigReply) GetConfigId() string {
	if x != nil {
		return x.ConfigId
	}
	return ""
}

func (x *ConfigReply) GetConfigName() string {
	if x != nil {
		return x.ConfigName
	}
	return ""
}

func (x *ConfigReply) GetConfigKey() string {
	if x != nil {
		return x.ConfigKey
	}
	return ""
}

func (x *ConfigReply) GetConfigValue() string {
	if x != nil {
		return x.ConfigValue
	}
	return ""
}

func (x *ConfigReply) GetConfigType() string {
	if x != nil {
		return x.ConfigType
	}
	return ""
}

func (x *ConfigReply) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

func (x *ConfigReply) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

var File_system_v1_sys_config_proto protoreflect.FileDescriptor

var file_system_v1_sys_config_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f,
	0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x22, 0x15, 0x0a, 0x13, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x14, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x65, 0x70, 0x22, 0x14, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53,
	0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x22, 0x2a, 0x0a, 0x0f, 0x47,
	0x65, 0x74, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x12, 0x17,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x47, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x53, 0x79,
	0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x32, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x22, 0x5a, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x70, 0x12, 0x21, 0x0a, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x07,
	0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x23, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02,
	0x20, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x22, 0x5a, 0x0a, 0x12,
	0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x2e, 0x0a, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x04, 0x72, 0x6f,
	0x77, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x22, 0x22, 0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x71, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x22, 0x28, 0x0a, 0x10,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xdf, 0x01, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x4b, 0x65, 0x79,
	0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79, 0x70, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72, 0x6b, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x32, 0xf1, 0x04, 0x0a, 0x09, 0x53, 0x79, 0x73,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x76, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x1a, 0x22, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70,
	0x74, 0x79, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16, 0x3a, 0x01, 0x2a, 0x22, 0x11, 0x2f, 0x76, 0x31,
	0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x58,
	0x0a, 0x0f, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x70, 0x1a, 0x22, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x58, 0x0a, 0x0f, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x21, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x1a, 0x22,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x12, 0x50, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x70, 0x1a, 0x20, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x12, 0x6f, 0x0a, 0x0d, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74,
	0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x1a, 0x21, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02,
	0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x2f, 0x6c, 0x69, 0x73, 0x74, 0x12, 0x75, 0x0a, 0x0b, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42,
	0x79, 0x4b, 0x65, 0x79, 0x12, 0x1d, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x4b, 0x65, 0x79,
	0x52, 0x65, 0x71, 0x1a, 0x1f, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x4b, 0x65, 0x79, 0x52,
	0x65, 0x70, 0x6c, 0x79, 0x22, 0x26, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x20, 0x12, 0x1e, 0x2f, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x4b, 0x65, 0x79, 0x2f, 0x7b, 0x6b, 0x65, 0x79, 0x7d, 0x42, 0x29, 0x0a, 0x0d,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a,
	0x16, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_v1_sys_config_proto_rawDescOnce sync.Once
	file_system_v1_sys_config_proto_rawDescData = file_system_v1_sys_config_proto_rawDesc
)

func file_system_v1_sys_config_proto_rawDescGZIP() []byte {
	file_system_v1_sys_config_proto_rawDescOnce.Do(func() {
		file_system_v1_sys_config_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_v1_sys_config_proto_rawDescData)
	})
	return file_system_v1_sys_config_proto_rawDescData
}

var file_system_v1_sys_config_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_system_v1_sys_config_proto_goTypes = []interface{}{
	(*CreateSysConfigRep)(nil),  // 0: api.system.v1.CreateSysConfigRep
	(*EmptySysConfigReply)(nil), // 1: api.system.v1.EmptySysConfigReply
	(*UpdateSysConfigRep)(nil),  // 2: api.system.v1.UpdateSysConfigRep
	(*DeleteSysConfigRep)(nil),  // 3: api.system.v1.DeleteSysConfigRep
	(*GetSysConfigRep)(nil),     // 4: api.system.v1.GetSysConfigRep
	(*GetSysConfigReply)(nil),   // 5: api.system.v1.GetSysConfigReply
	(*ListSysConfigRep)(nil),    // 6: api.system.v1.ListSysConfigRep
	(*ListSysConfigReply)(nil),  // 7: api.system.v1.ListSysConfigReply
	(*ConfigByKeyReq)(nil),      // 8: api.system.v1.ConfigByKeyReq
	(*ConfigByKeyReply)(nil),    // 9: api.system.v1.ConfigByKeyReply
	(*ConfigReply)(nil),         // 10: api.system.v1.ConfigReply
}
var file_system_v1_sys_config_proto_depIdxs = []int32{
	10, // 0: api.system.v1.GetSysConfigReply.config:type_name -> api.system.v1.ConfigReply
	10, // 1: api.system.v1.ListSysConfigReply.rows:type_name -> api.system.v1.ConfigReply
	0,  // 2: api.system.v1.SysConfig.CreateSysConfig:input_type -> api.system.v1.CreateSysConfigRep
	2,  // 3: api.system.v1.SysConfig.UpdateSysConfig:input_type -> api.system.v1.UpdateSysConfigRep
	3,  // 4: api.system.v1.SysConfig.DeleteSysConfig:input_type -> api.system.v1.DeleteSysConfigRep
	4,  // 5: api.system.v1.SysConfig.GetSysConfig:input_type -> api.system.v1.GetSysConfigRep
	6,  // 6: api.system.v1.SysConfig.ListSysConfig:input_type -> api.system.v1.ListSysConfigRep
	8,  // 7: api.system.v1.SysConfig.ConfigByKey:input_type -> api.system.v1.ConfigByKeyReq
	1,  // 8: api.system.v1.SysConfig.CreateSysConfig:output_type -> api.system.v1.EmptySysConfigReply
	1,  // 9: api.system.v1.SysConfig.UpdateSysConfig:output_type -> api.system.v1.EmptySysConfigReply
	1,  // 10: api.system.v1.SysConfig.DeleteSysConfig:output_type -> api.system.v1.EmptySysConfigReply
	5,  // 11: api.system.v1.SysConfig.GetSysConfig:output_type -> api.system.v1.GetSysConfigReply
	7,  // 12: api.system.v1.SysConfig.ListSysConfig:output_type -> api.system.v1.ListSysConfigReply
	9,  // 13: api.system.v1.SysConfig.ConfigByKey:output_type -> api.system.v1.ConfigByKeyReply
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_system_v1_sys_config_proto_init() }
func file_system_v1_sys_config_proto_init() {
	if File_system_v1_sys_config_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_system_v1_sys_config_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSysConfigRep); i {
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
		file_system_v1_sys_config_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EmptySysConfigReply); i {
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
		file_system_v1_sys_config_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSysConfigRep); i {
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
		file_system_v1_sys_config_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSysConfigRep); i {
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
		file_system_v1_sys_config_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSysConfigRep); i {
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
		file_system_v1_sys_config_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSysConfigReply); i {
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
		file_system_v1_sys_config_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSysConfigRep); i {
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
		file_system_v1_sys_config_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSysConfigReply); i {
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
		file_system_v1_sys_config_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigByKeyReq); i {
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
		file_system_v1_sys_config_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigByKeyReply); i {
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
		file_system_v1_sys_config_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ConfigReply); i {
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
			RawDescriptor: file_system_v1_sys_config_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_v1_sys_config_proto_goTypes,
		DependencyIndexes: file_system_v1_sys_config_proto_depIdxs,
		MessageInfos:      file_system_v1_sys_config_proto_msgTypes,
	}.Build()
	File_system_v1_sys_config_proto = out.File
	file_system_v1_sys_config_proto_rawDesc = nil
	file_system_v1_sys_config_proto_goTypes = nil
	file_system_v1_sys_config_proto_depIdxs = nil
}
