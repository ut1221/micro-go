// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.1
// 	protoc        v5.28.2
// source: api/system/v1/auth.proto

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

type LogoutReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogoutReq) Reset() {
	*x = LogoutReq{}
	mi := &file_api_system_v1_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReq) ProtoMessage() {}

func (x *LogoutReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReq.ProtoReflect.Descriptor instead.
func (*LogoutReq) Descriptor() ([]byte, []int) {
	return file_api_system_v1_auth_proto_rawDescGZIP(), []int{0}
}

type LogoutReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LogoutReply) Reset() {
	*x = LogoutReply{}
	mi := &file_api_system_v1_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LogoutReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogoutReply) ProtoMessage() {}

func (x *LogoutReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogoutReply.ProtoReflect.Descriptor instead.
func (*LogoutReply) Descriptor() ([]byte, []int) {
	return file_api_system_v1_auth_proto_rawDescGZIP(), []int{1}
}

type CaptchaReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CaptchaReq) Reset() {
	*x = CaptchaReq{}
	mi := &file_api_system_v1_auth_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CaptchaReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptchaReq) ProtoMessage() {}

func (x *CaptchaReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1_auth_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptchaReq.ProtoReflect.Descriptor instead.
func (*CaptchaReq) Descriptor() ([]byte, []int) {
	return file_api_system_v1_auth_proto_rawDescGZIP(), []int{2}
}

type CaptchaReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CaptchaEnable bool   `protobuf:"varint,1,opt,name=captchaEnable,proto3" json:"captchaEnable"`
	Img           string `protobuf:"bytes,2,opt,name=img,proto3" json:"img"`
	Uuid          string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid"`
}

func (x *CaptchaReply) Reset() {
	*x = CaptchaReply{}
	mi := &file_api_system_v1_auth_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CaptchaReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CaptchaReply) ProtoMessage() {}

func (x *CaptchaReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1_auth_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CaptchaReply.ProtoReflect.Descriptor instead.
func (*CaptchaReply) Descriptor() ([]byte, []int) {
	return file_api_system_v1_auth_proto_rawDescGZIP(), []int{3}
}

func (x *CaptchaReply) GetCaptchaEnable() bool {
	if x != nil {
		return x.CaptchaEnable
	}
	return false
}

func (x *CaptchaReply) GetImg() string {
	if x != nil {
		return x.Img
	}
	return ""
}

func (x *CaptchaReply) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password"`
	Uuid     string `protobuf:"bytes,3,opt,name=uuid,proto3" json:"uuid"`
	Code     string `protobuf:"bytes,4,opt,name=code,proto3" json:"code"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	mi := &file_api_system_v1_auth_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1_auth_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_api_system_v1_auth_proto_rawDescGZIP(), []int{4}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *LoginReq) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *LoginReq) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type LoginReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token"`
}

func (x *LoginReply) Reset() {
	*x = LoginReply{}
	mi := &file_api_system_v1_auth_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LoginReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReply) ProtoMessage() {}

func (x *LoginReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1_auth_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReply.ProtoReflect.Descriptor instead.
func (*LoginReply) Descriptor() ([]byte, []int) {
	return file_api_system_v1_auth_proto_rawDescGZIP(), []int{5}
}

func (x *LoginReply) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type UserInfoReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UserInfoReq) Reset() {
	*x = UserInfoReq{}
	mi := &file_api_system_v1_auth_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoReq) ProtoMessage() {}

func (x *UserInfoReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1_auth_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoReq.ProtoReflect.Descriptor instead.
func (*UserInfoReq) Descriptor() ([]byte, []int) {
	return file_api_system_v1_auth_proto_rawDescGZIP(), []int{6}
}

type UserInfoReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Permissions []string   `protobuf:"bytes,1,rep,name=permissions,proto3" json:"permissions"`
	Roles       []string   `protobuf:"bytes,2,rep,name=roles,proto3" json:"roles"`
	User        *UserReply `protobuf:"bytes,3,opt,name=user,proto3" json:"user"`
}

func (x *UserInfoReply) Reset() {
	*x = UserInfoReply{}
	mi := &file_api_system_v1_auth_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *UserInfoReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfoReply) ProtoMessage() {}

func (x *UserInfoReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1_auth_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfoReply.ProtoReflect.Descriptor instead.
func (*UserInfoReply) Descriptor() ([]byte, []int) {
	return file_api_system_v1_auth_proto_rawDescGZIP(), []int{7}
}

func (x *UserInfoReply) GetPermissions() []string {
	if x != nil {
		return x.Permissions
	}
	return nil
}

func (x *UserInfoReply) GetRoles() []string {
	if x != nil {
		return x.Roles
	}
	return nil
}

func (x *UserInfoReply) GetUser() *UserReply {
	if x != nil {
		return x.User
	}
	return nil
}

type RoutersReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RoutersReq) Reset() {
	*x = RoutersReq{}
	mi := &file_api_system_v1_auth_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoutersReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutersReq) ProtoMessage() {}

func (x *RoutersReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1_auth_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutersReq.ProtoReflect.Descriptor instead.
func (*RoutersReq) Descriptor() ([]byte, []int) {
	return file_api_system_v1_auth_proto_rawDescGZIP(), []int{8}
}

type RoutersReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Routers []*Router `protobuf:"bytes,1,rep,name=routers,proto3" json:"routers"`
}

func (x *RoutersReply) Reset() {
	*x = RoutersReply{}
	mi := &file_api_system_v1_auth_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RoutersReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RoutersReply) ProtoMessage() {}

func (x *RoutersReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1_auth_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RoutersReply.ProtoReflect.Descriptor instead.
func (*RoutersReply) Descriptor() ([]byte, []int) {
	return file_api_system_v1_auth_proto_rawDescGZIP(), []int{9}
}

func (x *RoutersReply) GetRouters() []*Router {
	if x != nil {
		return x.Routers
	}
	return nil
}

type Router struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string       `protobuf:"bytes,1,opt,name=name,proto3" json:"name"`
	Path      string       `protobuf:"bytes,2,opt,name=path,proto3" json:"path"`
	Redirect  string       `protobuf:"bytes,4,opt,name=redirect,proto3" json:"redirect"`
	Component string       `protobuf:"bytes,5,opt,name=component,proto3" json:"component"`
	Meta      *Router_Meta `protobuf:"bytes,7,opt,name=meta,proto3" json:"meta"`
	Children  []*Router    `protobuf:"bytes,8,rep,name=children,proto3" json:"children"`
	MenuId    string       `protobuf:"bytes,9,opt,name=menuId,proto3" json:"menuId"`
	ParentId  string       `protobuf:"bytes,10,opt,name=parentId,proto3" json:"parentId"`
}

func (x *Router) Reset() {
	*x = Router{}
	mi := &file_api_system_v1_auth_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Router) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Router) ProtoMessage() {}

func (x *Router) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1_auth_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Router.ProtoReflect.Descriptor instead.
func (*Router) Descriptor() ([]byte, []int) {
	return file_api_system_v1_auth_proto_rawDescGZIP(), []int{10}
}

func (x *Router) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Router) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Router) GetRedirect() string {
	if x != nil {
		return x.Redirect
	}
	return ""
}

func (x *Router) GetComponent() string {
	if x != nil {
		return x.Component
	}
	return ""
}

func (x *Router) GetMeta() *Router_Meta {
	if x != nil {
		return x.Meta
	}
	return nil
}

func (x *Router) GetChildren() []*Router {
	if x != nil {
		return x.Children
	}
	return nil
}

func (x *Router) GetMenuId() string {
	if x != nil {
		return x.MenuId
	}
	return ""
}

func (x *Router) GetParentId() string {
	if x != nil {
		return x.ParentId
	}
	return ""
}

type RegisterReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname,proto3" json:"nickname"`
	Mobile   string `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile"`
	Password string `protobuf:"bytes,4,opt,name=password,proto3" json:"password"`
}

func (x *RegisterReq) Reset() {
	*x = RegisterReq{}
	mi := &file_api_system_v1_auth_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReq) ProtoMessage() {}

func (x *RegisterReq) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1_auth_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReq.ProtoReflect.Descriptor instead.
func (*RegisterReq) Descriptor() ([]byte, []int) {
	return file_api_system_v1_auth_proto_rawDescGZIP(), []int{11}
}

func (x *RegisterReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *RegisterReq) GetNickname() string {
	if x != nil {
		return x.Nickname
	}
	return ""
}

func (x *RegisterReq) GetMobile() string {
	if x != nil {
		return x.Mobile
	}
	return ""
}

func (x *RegisterReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type RegisterReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RegisterReply) Reset() {
	*x = RegisterReply{}
	mi := &file_api_system_v1_auth_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RegisterReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterReply) ProtoMessage() {}

func (x *RegisterReply) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1_auth_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterReply.ProtoReflect.Descriptor instead.
func (*RegisterReply) Descriptor() ([]byte, []int) {
	return file_api_system_v1_auth_proto_rawDescGZIP(), []int{12}
}

type Router_Meta struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title   string `protobuf:"bytes,1,opt,name=title,proto3" json:"title"`
	Icon    string `protobuf:"bytes,2,opt,name=icon,proto3" json:"icon"`
	NoCache bool   `protobuf:"varint,3,opt,name=noCache,proto3" json:"noCache"`
	Link    string `protobuf:"bytes,5,opt,name=link,proto3" json:"link"`
}

func (x *Router_Meta) Reset() {
	*x = Router_Meta{}
	mi := &file_api_system_v1_auth_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Router_Meta) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Router_Meta) ProtoMessage() {}

func (x *Router_Meta) ProtoReflect() protoreflect.Message {
	mi := &file_api_system_v1_auth_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Router_Meta.ProtoReflect.Descriptor instead.
func (*Router_Meta) Descriptor() ([]byte, []int) {
	return file_api_system_v1_auth_proto_rawDescGZIP(), []int{10, 0}
}

func (x *Router_Meta) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Router_Meta) GetIcon() string {
	if x != nil {
		return x.Icon
	}
	return ""
}

func (x *Router_Meta) GetNoCache() bool {
	if x != nil {
		return x.NoCache
	}
	return false
}

func (x *Router_Meta) GetLink() string {
	if x != nil {
		return x.Link
	}
	return ""
}

var File_api_system_v1_auth_proto protoreflect.FileDescriptor

var file_api_system_v1_auth_proto_rawDesc = []byte{
	0x0a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x18, 0x61, 0x70, 0x69, 0x2f, 0x73, 0x79, 0x73,
	0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x61, 0x73, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x0b, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x22, 0x0d,
	0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x0c, 0x0a,
	0x0a, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x71, 0x22, 0x5a, 0x0a, 0x0c, 0x43,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x24, 0x0a, 0x0d, 0x63,
	0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x45, 0x6e, 0x61, 0x62, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x0d, 0x63, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x45, 0x6e, 0x61, 0x62, 0x6c,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x69, 0x6d, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x69, 0x6d, 0x67, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x22, 0x6a, 0x0a, 0x08, 0x4c, 0x6f, 0x67, 0x69, 0x6e,
	0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x22, 0x22, 0x0a, 0x0a, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x70, 0x6c,
	0x79, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x0d, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x49,
	0x6e, 0x66, 0x6f, 0x52, 0x65, 0x71, 0x22, 0x75, 0x0a, 0x0d, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x65, 0x72, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0b, 0x70, 0x65,
	0x72, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x6f, 0x6c,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x05, 0x72, 0x6f, 0x6c, 0x65, 0x73, 0x12,
	0x2c, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x0c, 0x0a,
	0x0a, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x22, 0x3f, 0x0a, 0x0c, 0x52,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x2f, 0x0a, 0x07, 0x72,
	0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75,
	0x74, 0x65, 0x72, 0x52, 0x07, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x22, 0xe1, 0x02, 0x0a,
	0x06, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70,
	0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12,
	0x1a, 0x0a, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x72, 0x65, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x63,
	0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x63, 0x6f, 0x6d, 0x70, 0x6f, 0x6e, 0x65, 0x6e, 0x74, 0x12, 0x2e, 0x0a, 0x04, 0x6d, 0x65, 0x74,
	0x61, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x2e, 0x4d,
	0x65, 0x74, 0x61, 0x52, 0x04, 0x6d, 0x65, 0x74, 0x61, 0x12, 0x31, 0x0a, 0x08, 0x63, 0x68, 0x69,
	0x6c, 0x64, 0x72, 0x65, 0x6e, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74,
	0x65, 0x72, 0x52, 0x08, 0x63, 0x68, 0x69, 0x6c, 0x64, 0x72, 0x65, 0x6e, 0x12, 0x16, 0x0a, 0x06,
	0x6d, 0x65, 0x6e, 0x75, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x65,
	0x6e, 0x75, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x49, 0x64,
	0x1a, 0x5e, 0x0a, 0x04, 0x4d, 0x65, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x69, 0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x69, 0x63,
	0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07, 0x6e, 0x6f, 0x43, 0x61, 0x63, 0x68, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x07, 0x6e, 0x6f, 0x43, 0x61, 0x63, 0x68, 0x65, 0x12, 0x12, 0x0a, 0x04,
	0x6c, 0x69, 0x6e, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6c, 0x69, 0x6e, 0x6b,
	0x22, 0x79, 0x0a, 0x0b, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6e,
	0x69, 0x63, 0x6b, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c,
	0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x6d, 0x6f, 0x62, 0x69, 0x6c, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x0f, 0x0a, 0x0d, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x32, 0xa2, 0x04, 0x0a,
	0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x55, 0x0a, 0x06, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12,
	0x18, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e,
	0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x6f, 0x75, 0x74,
	0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x15, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0f, 0x3a, 0x01, 0x2a,
	0x22, 0x0a, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x6f, 0x75, 0x74, 0x12, 0x5b, 0x0a, 0x07,
	0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79,
	0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52,
	0x65, 0x71, 0x1a, 0x1b, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x61, 0x70, 0x74, 0x63, 0x68, 0x61, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22,
	0x18, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x12, 0x12, 0x10, 0x2f, 0x76, 0x31, 0x2f, 0x63, 0x61, 0x70,
	0x74, 0x63, 0x68, 0x61, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x51, 0x0a, 0x05, 0x4c, 0x6f, 0x67,
	0x69, 0x6e, 0x12, 0x17, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x67, 0x69,
	0x6e, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x14, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x3a, 0x01,
	0x2a, 0x22, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x59, 0x0a, 0x08,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1a, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73,
	0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65,
	0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x65, 0x70,
	0x6c, 0x79, 0x22, 0x13, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0d, 0x12, 0x0b, 0x2f, 0x76, 0x31, 0x2f,
	0x67, 0x65, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x59, 0x0a, 0x07, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x1b, 0x2e,
	0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x6f,
	0x75, 0x74, 0x65, 0x72, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93,
	0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x67, 0x65, 0x74, 0x52, 0x6f, 0x75, 0x74, 0x65,
	0x72, 0x73, 0x12, 0x5d, 0x0a, 0x08, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x1a,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52,
	0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x1c, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x17, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x11,
	0x3a, 0x01, 0x2a, 0x22, 0x0c, 0x2f, 0x76, 0x31, 0x2f, 0x72, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65,
	0x72, 0x42, 0x29, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2e,
	0x76, 0x31, 0x50, 0x01, 0x5a, 0x16, 0x6d, 0x69, 0x63, 0x72, 0x6f, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x2f, 0x76, 0x31, 0x3b, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_system_v1_auth_proto_rawDescOnce sync.Once
	file_api_system_v1_auth_proto_rawDescData = file_api_system_v1_auth_proto_rawDesc
)

func file_api_system_v1_auth_proto_rawDescGZIP() []byte {
	file_api_system_v1_auth_proto_rawDescOnce.Do(func() {
		file_api_system_v1_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_system_v1_auth_proto_rawDescData)
	})
	return file_api_system_v1_auth_proto_rawDescData
}

var file_api_system_v1_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_api_system_v1_auth_proto_goTypes = []any{
	(*LogoutReq)(nil),     // 0: api.system.v1.LogoutReq
	(*LogoutReply)(nil),   // 1: api.system.v1.LogoutReply
	(*CaptchaReq)(nil),    // 2: api.system.v1.CaptchaReq
	(*CaptchaReply)(nil),  // 3: api.system.v1.CaptchaReply
	(*LoginReq)(nil),      // 4: api.system.v1.LoginReq
	(*LoginReply)(nil),    // 5: api.system.v1.LoginReply
	(*UserInfoReq)(nil),   // 6: api.system.v1.UserInfoReq
	(*UserInfoReply)(nil), // 7: api.system.v1.UserInfoReply
	(*RoutersReq)(nil),    // 8: api.system.v1.RoutersReq
	(*RoutersReply)(nil),  // 9: api.system.v1.RoutersReply
	(*Router)(nil),        // 10: api.system.v1.Router
	(*RegisterReq)(nil),   // 11: api.system.v1.RegisterReq
	(*RegisterReply)(nil), // 12: api.system.v1.RegisterReply
	(*Router_Meta)(nil),   // 13: api.system.v1.Router.Meta
	(*UserReply)(nil),     // 14: api.system.v1.UserReply
}
var file_api_system_v1_auth_proto_depIdxs = []int32{
	14, // 0: api.system.v1.UserInfoReply.user:type_name -> api.system.v1.UserReply
	10, // 1: api.system.v1.RoutersReply.routers:type_name -> api.system.v1.Router
	13, // 2: api.system.v1.Router.meta:type_name -> api.system.v1.Router.Meta
	10, // 3: api.system.v1.Router.children:type_name -> api.system.v1.Router
	0,  // 4: api.system.v1.Auth.Logout:input_type -> api.system.v1.LogoutReq
	2,  // 5: api.system.v1.Auth.Captcha:input_type -> api.system.v1.CaptchaReq
	4,  // 6: api.system.v1.Auth.Login:input_type -> api.system.v1.LoginReq
	6,  // 7: api.system.v1.Auth.UserInfo:input_type -> api.system.v1.UserInfoReq
	8,  // 8: api.system.v1.Auth.Routers:input_type -> api.system.v1.RoutersReq
	11, // 9: api.system.v1.Auth.Register:input_type -> api.system.v1.RegisterReq
	1,  // 10: api.system.v1.Auth.Logout:output_type -> api.system.v1.LogoutReply
	3,  // 11: api.system.v1.Auth.Captcha:output_type -> api.system.v1.CaptchaReply
	5,  // 12: api.system.v1.Auth.Login:output_type -> api.system.v1.LoginReply
	7,  // 13: api.system.v1.Auth.UserInfo:output_type -> api.system.v1.UserInfoReply
	9,  // 14: api.system.v1.Auth.Routers:output_type -> api.system.v1.RoutersReply
	12, // 15: api.system.v1.Auth.Register:output_type -> api.system.v1.RegisterReply
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_api_system_v1_auth_proto_init() }
func file_api_system_v1_auth_proto_init() {
	if File_api_system_v1_auth_proto != nil {
		return
	}
	file_api_system_v1_base_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_system_v1_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_system_v1_auth_proto_goTypes,
		DependencyIndexes: file_api_system_v1_auth_proto_depIdxs,
		MessageInfos:      file_api_system_v1_auth_proto_msgTypes,
	}.Build()
	File_api_system_v1_auth_proto = out.File
	file_api_system_v1_auth_proto_rawDesc = nil
	file_api_system_v1_auth_proto_goTypes = nil
	file_api_system_v1_auth_proto_depIdxs = nil
}
