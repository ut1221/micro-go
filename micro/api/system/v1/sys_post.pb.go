// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: system/v1/sys_post.proto

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

type SysPostRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PostId   string `protobuf:"bytes,1,opt,name=postId,proto3" json:"postId"`
	PostCode string `protobuf:"bytes,2,opt,name=postCode,proto3" json:"postCode"`
	PostName string `protobuf:"bytes,3,opt,name=postName,proto3" json:"postName"`
	PostSort int64  `protobuf:"varint,4,opt,name=postSort,proto3" json:"postSort"`
	Status   string `protobuf:"bytes,5,opt,name=status,proto3" json:"status"`
	Remark   string `protobuf:"bytes,6,opt,name=remark,proto3" json:"remark"`
}

func (x *SysPostRep) Reset() {
	*x = SysPostRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_post_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SysPostRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SysPostRep) ProtoMessage() {}

func (x *SysPostRep) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_post_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SysPostRep.ProtoReflect.Descriptor instead.
func (*SysPostRep) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_post_proto_rawDescGZIP(), []int{0}
}

func (x *SysPostRep) GetPostId() string {
	if x != nil {
		return x.PostId
	}
	return ""
}

func (x *SysPostRep) GetPostCode() string {
	if x != nil {
		return x.PostCode
	}
	return ""
}

func (x *SysPostRep) GetPostName() string {
	if x != nil {
		return x.PostName
	}
	return ""
}

func (x *SysPostRep) GetPostSort() int64 {
	if x != nil {
		return x.PostSort
	}
	return 0
}

func (x *SysPostRep) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *SysPostRep) GetRemark() string {
	if x != nil {
		return x.Remark
	}
	return ""
}

type DeleteSysPostRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *DeleteSysPostRep) Reset() {
	*x = DeleteSysPostRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_post_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSysPostRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSysPostRep) ProtoMessage() {}

func (x *DeleteSysPostRep) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_post_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSysPostRep.ProtoReflect.Descriptor instead.
func (*DeleteSysPostRep) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_post_proto_rawDescGZIP(), []int{1}
}

func (x *DeleteSysPostRep) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSysPostRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id"`
}

func (x *GetSysPostRep) Reset() {
	*x = GetSysPostRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_post_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSysPostRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSysPostRep) ProtoMessage() {}

func (x *GetSysPostRep) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_post_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSysPostRep.ProtoReflect.Descriptor instead.
func (*GetSysPostRep) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_post_proto_rawDescGZIP(), []int{2}
}

func (x *GetSysPostRep) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetSysPostReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Post *PostReply `protobuf:"bytes,1,opt,name=post,proto3" json:"post"`
}

func (x *GetSysPostReply) Reset() {
	*x = GetSysPostReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_post_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSysPostReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSysPostReply) ProtoMessage() {}

func (x *GetSysPostReply) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_post_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSysPostReply.ProtoReflect.Descriptor instead.
func (*GetSysPostReply) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_post_proto_rawDescGZIP(), []int{3}
}

func (x *GetSysPostReply) GetPost() *PostReply {
	if x != nil {
		return x.Post
	}
	return nil
}

type ListSysPostRep struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// pageNum
	PageNum int64 `protobuf:"varint,1,opt,name=pageNum,proto3" json:"pageNum"`
	// 页面大小
	PageSize int64 `protobuf:"varint,2,opt,name=pageSize,proto3" json:"pageSize"`
}

func (x *ListSysPostRep) Reset() {
	*x = ListSysPostRep{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_post_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSysPostRep) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSysPostRep) ProtoMessage() {}

func (x *ListSysPostRep) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_post_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSysPostRep.ProtoReflect.Descriptor instead.
func (*ListSysPostRep) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_post_proto_rawDescGZIP(), []int{4}
}

func (x *ListSysPostRep) GetPageNum() int64 {
	if x != nil {
		return x.PageNum
	}
	return 0
}

func (x *ListSysPostRep) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

type ListSysPostReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64        `protobuf:"varint,1,opt,name=total,proto3" json:"total"`
	Rows  []*PostReply `protobuf:"bytes,2,rep,name=rows,proto3" json:"rows"`
}

func (x *ListSysPostReply) Reset() {
	*x = ListSysPostReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_system_v1_sys_post_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListSysPostReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListSysPostReply) ProtoMessage() {}

func (x *ListSysPostReply) ProtoReflect() protoreflect.Message {
	mi := &file_system_v1_sys_post_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListSysPostReply.ProtoReflect.Descriptor instead.
func (*ListSysPostReply) Descriptor() ([]byte, []int) {
	return file_system_v1_sys_post_proto_rawDescGZIP(), []int{5}
}

func (x *ListSysPostReply) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *ListSysPostReply) GetRows() []*PostReply {
	if x != nil {
		return x.Rows
	}
	return nil
}

var File_system_v1_sys_post_proto protoreflect.FileDescriptor

var file_system_v1_sys_post_proto_rawDesc = []byte{
	0x0a, 0x18, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x79, 0x73, 0x5f,
	0x70, 0x6f, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74,
	0x65, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x1a, 0x14, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa8, 0x01, 0x0a, 0x0a, 0x53, 0x79, 0x73, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a,
	0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x70, 0x6f, 0x73, 0x74, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6f, 0x73,
	0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x53, 0x6f, 0x72,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x74, 0x53, 0x6f, 0x72,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x6d,
	0x61, 0x72, 0x6b, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x72, 0x65, 0x6d, 0x61, 0x72,
	0x6b, 0x22, 0x2b, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x12, 0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x28,
	0x0a, 0x0d, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x12,
	0x17, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x10, 0x01, 0x52, 0x02, 0x69, 0x64, 0x22, 0x3f, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x53,
	0x79, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2c, 0x0a, 0x04, 0x70,
	0x6f, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x52, 0x04, 0x70, 0x6f, 0x73, 0x74, 0x22, 0x58, 0x0a, 0x0e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x79, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x12, 0x21, 0x0a, 0x07, 0x70,
	0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x42, 0x07, 0xfa, 0x42,
	0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x07, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x12, 0x23,
	0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x42, 0x07, 0xfa, 0x42, 0x04, 0x22, 0x02, 0x20, 0x00, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53,
	0x69, 0x7a, 0x65, 0x22, 0x56, 0x0a, 0x10, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x50, 0x6f,
	0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x2c, 0x0a,
	0x04, 0x72, 0x6f, 0x77, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x04, 0x72, 0x6f, 0x77, 0x73, 0x32, 0x89, 0x04, 0x0a, 0x07,
	0x53, 0x79, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x5f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x79, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x50, 0x6f, 0x73, 0x74,
	0x52, 0x65, 0x70, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d,
	0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18,
	0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x6f, 0x73, 0x74, 0x2f, 0x73, 0x61, 0x76, 0x65, 0x12, 0x61, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x79, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x79, 0x73, 0x50, 0x6f, 0x73,
	0x74, 0x52, 0x65, 0x70, 0x1a, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x1a, 0x0f, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x69, 0x0a, 0x0d, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x53, 0x79, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x1a, 0x19, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x1c, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x16,
	0x2a, 0x14, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f, 0x73, 0x74, 0x2f, 0x64, 0x65, 0x6c, 0x65, 0x74,
	0x65, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x66, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73,
	0x50, 0x6f, 0x73, 0x74, 0x12, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x52,
	0x65, 0x70, 0x1a, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x79, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x76, 0x31, 0x2f,
	0x70, 0x6f, 0x73, 0x74, 0x2f, 0x69, 0x6e, 0x66, 0x6f, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x12, 0x67,
	0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x79, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x12, 0x1d, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x53, 0x79, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x1a, 0x1f, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x73,
	0x74, 0x53, 0x79, 0x73, 0x50, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x18, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x12, 0x3a, 0x01, 0x2a, 0x22, 0x0d, 0x2f, 0x76, 0x31, 0x2f, 0x70, 0x6f,
	0x73, 0x74, 0x2f, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x29, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x16, 0x6d, 0x69, 0x63, 0x72,
	0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x3b,
	0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_system_v1_sys_post_proto_rawDescOnce sync.Once
	file_system_v1_sys_post_proto_rawDescData = file_system_v1_sys_post_proto_rawDesc
)

func file_system_v1_sys_post_proto_rawDescGZIP() []byte {
	file_system_v1_sys_post_proto_rawDescOnce.Do(func() {
		file_system_v1_sys_post_proto_rawDescData = protoimpl.X.CompressGZIP(file_system_v1_sys_post_proto_rawDescData)
	})
	return file_system_v1_sys_post_proto_rawDescData
}

var file_system_v1_sys_post_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_system_v1_sys_post_proto_goTypes = []interface{}{
	(*SysPostRep)(nil),       // 0: api.system.v1.SysPostRep
	(*DeleteSysPostRep)(nil), // 1: api.system.v1.DeleteSysPostRep
	(*GetSysPostRep)(nil),    // 2: api.system.v1.GetSysPostRep
	(*GetSysPostReply)(nil),  // 3: api.system.v1.GetSysPostReply
	(*ListSysPostRep)(nil),   // 4: api.system.v1.ListSysPostRep
	(*ListSysPostReply)(nil), // 5: api.system.v1.ListSysPostReply
	(*PostReply)(nil),        // 6: api.system.v1.PostReply
	(*EmptyReply)(nil),       // 7: api.system.v1.EmptyReply
}
var file_system_v1_sys_post_proto_depIdxs = []int32{
	6, // 0: api.system.v1.GetSysPostReply.post:type_name -> api.system.v1.PostReply
	6, // 1: api.system.v1.ListSysPostReply.rows:type_name -> api.system.v1.PostReply
	0, // 2: api.system.v1.SysPost.CreateSysPost:input_type -> api.system.v1.SysPostRep
	0, // 3: api.system.v1.SysPost.UpdateSysPost:input_type -> api.system.v1.SysPostRep
	1, // 4: api.system.v1.SysPost.DeleteSysPost:input_type -> api.system.v1.DeleteSysPostRep
	2, // 5: api.system.v1.SysPost.GetSysPost:input_type -> api.system.v1.GetSysPostRep
	4, // 6: api.system.v1.SysPost.ListSysPost:input_type -> api.system.v1.ListSysPostRep
	7, // 7: api.system.v1.SysPost.CreateSysPost:output_type -> api.system.v1.EmptyReply
	7, // 8: api.system.v1.SysPost.UpdateSysPost:output_type -> api.system.v1.EmptyReply
	7, // 9: api.system.v1.SysPost.DeleteSysPost:output_type -> api.system.v1.EmptyReply
	3, // 10: api.system.v1.SysPost.GetSysPost:output_type -> api.system.v1.GetSysPostReply
	5, // 11: api.system.v1.SysPost.ListSysPost:output_type -> api.system.v1.ListSysPostReply
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_system_v1_sys_post_proto_init() }
func file_system_v1_sys_post_proto_init() {
	if File_system_v1_sys_post_proto != nil {
		return
	}
	file_system_v1_base_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_system_v1_sys_post_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SysPostRep); i {
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
		file_system_v1_sys_post_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSysPostRep); i {
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
		file_system_v1_sys_post_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSysPostRep); i {
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
		file_system_v1_sys_post_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSysPostReply); i {
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
		file_system_v1_sys_post_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSysPostRep); i {
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
		file_system_v1_sys_post_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListSysPostReply); i {
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
			RawDescriptor: file_system_v1_sys_post_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_system_v1_sys_post_proto_goTypes,
		DependencyIndexes: file_system_v1_sys_post_proto_depIdxs,
		MessageInfos:      file_system_v1_sys_post_proto_msgTypes,
	}.Build()
	File_system_v1_sys_post_proto = out.File
	file_system_v1_sys_post_proto_rawDesc = nil
	file_system_v1_sys_post_proto_goTypes = nil
	file_system_v1_sys_post_proto_depIdxs = nil
}
