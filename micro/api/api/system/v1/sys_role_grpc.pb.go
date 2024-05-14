// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/system/v1/sys_role.proto

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

const (
	SysRole_CreateSysRole_FullMethodName       = "/api.system.v1.SysRole/CreateSysRole"
	SysRole_UpdateSysRole_FullMethodName       = "/api.system.v1.SysRole/UpdateSysRole"
	SysRole_DataScopeSysRole_FullMethodName    = "/api.system.v1.SysRole/DataScopeSysRole"
	SysRole_ChangeStatusSysRole_FullMethodName = "/api.system.v1.SysRole/ChangeStatusSysRole"
	SysRole_DeleteSysRole_FullMethodName       = "/api.system.v1.SysRole/DeleteSysRole"
	SysRole_GetSysRole_FullMethodName          = "/api.system.v1.SysRole/GetSysRole"
	SysRole_ListSysRole_FullMethodName         = "/api.system.v1.SysRole/ListSysRole"
	SysRole_AllocatedList_FullMethodName       = "/api.system.v1.SysRole/AllocatedList"
	SysRole_UnAllocatedList_FullMethodName     = "/api.system.v1.SysRole/UnAllocatedList"
	SysRole_AuthUserCancel_FullMethodName      = "/api.system.v1.SysRole/AuthUserCancel"
	SysRole_AuthUserCancelAll_FullMethodName   = "/api.system.v1.SysRole/AuthUserCancelAll"
	SysRole_AuthUserSelectAll_FullMethodName   = "/api.system.v1.SysRole/AuthUserSelectAll"
)

// SysRoleClient is the client API for SysRole service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysRoleClient interface {
	CreateSysRole(ctx context.Context, in *ModifySysRoleRep, opts ...grpc.CallOption) (*EmptyReply, error)
	UpdateSysRole(ctx context.Context, in *ModifySysRoleRep, opts ...grpc.CallOption) (*EmptyReply, error)
	DataScopeSysRole(ctx context.Context, in *DataScopeSysRoleRep, opts ...grpc.CallOption) (*EmptyReply, error)
	ChangeStatusSysRole(ctx context.Context, in *ChangeStatusSysRoleRep, opts ...grpc.CallOption) (*EmptyReply, error)
	DeleteSysRole(ctx context.Context, in *DeleteSysRoleRep, opts ...grpc.CallOption) (*EmptyReply, error)
	GetSysRole(ctx context.Context, in *GetSysRoleRep, opts ...grpc.CallOption) (*GetSysRoleReply, error)
	ListSysRole(ctx context.Context, in *ListSysRoleRep, opts ...grpc.CallOption) (*ListSysRoleReply, error)
	AllocatedList(ctx context.Context, in *IsAllocatedListRep, opts ...grpc.CallOption) (*IsAllocatedListReply, error)
	UnAllocatedList(ctx context.Context, in *IsAllocatedListRep, opts ...grpc.CallOption) (*IsAllocatedListReply, error)
	AuthUserCancel(ctx context.Context, in *AuthUserCancelRep, opts ...grpc.CallOption) (*EmptyReply, error)
	AuthUserCancelAll(ctx context.Context, in *AuthUserSelectAllRep, opts ...grpc.CallOption) (*EmptyReply, error)
	AuthUserSelectAll(ctx context.Context, in *AuthUserSelectAllRep, opts ...grpc.CallOption) (*EmptyReply, error)
}

type sysRoleClient struct {
	cc grpc.ClientConnInterface
}

func NewSysRoleClient(cc grpc.ClientConnInterface) SysRoleClient {
	return &sysRoleClient{cc}
}

func (c *sysRoleClient) CreateSysRole(ctx context.Context, in *ModifySysRoleRep, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, SysRole_CreateSysRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRoleClient) UpdateSysRole(ctx context.Context, in *ModifySysRoleRep, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, SysRole_UpdateSysRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRoleClient) DataScopeSysRole(ctx context.Context, in *DataScopeSysRoleRep, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, SysRole_DataScopeSysRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRoleClient) ChangeStatusSysRole(ctx context.Context, in *ChangeStatusSysRoleRep, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, SysRole_ChangeStatusSysRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRoleClient) DeleteSysRole(ctx context.Context, in *DeleteSysRoleRep, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, SysRole_DeleteSysRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRoleClient) GetSysRole(ctx context.Context, in *GetSysRoleRep, opts ...grpc.CallOption) (*GetSysRoleReply, error) {
	out := new(GetSysRoleReply)
	err := c.cc.Invoke(ctx, SysRole_GetSysRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRoleClient) ListSysRole(ctx context.Context, in *ListSysRoleRep, opts ...grpc.CallOption) (*ListSysRoleReply, error) {
	out := new(ListSysRoleReply)
	err := c.cc.Invoke(ctx, SysRole_ListSysRole_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRoleClient) AllocatedList(ctx context.Context, in *IsAllocatedListRep, opts ...grpc.CallOption) (*IsAllocatedListReply, error) {
	out := new(IsAllocatedListReply)
	err := c.cc.Invoke(ctx, SysRole_AllocatedList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRoleClient) UnAllocatedList(ctx context.Context, in *IsAllocatedListRep, opts ...grpc.CallOption) (*IsAllocatedListReply, error) {
	out := new(IsAllocatedListReply)
	err := c.cc.Invoke(ctx, SysRole_UnAllocatedList_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRoleClient) AuthUserCancel(ctx context.Context, in *AuthUserCancelRep, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, SysRole_AuthUserCancel_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRoleClient) AuthUserCancelAll(ctx context.Context, in *AuthUserSelectAllRep, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, SysRole_AuthUserCancelAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysRoleClient) AuthUserSelectAll(ctx context.Context, in *AuthUserSelectAllRep, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, SysRole_AuthUserSelectAll_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysRoleServer is the server API for SysRole service.
// All implementations must embed UnimplementedSysRoleServer
// for forward compatibility
type SysRoleServer interface {
	CreateSysRole(context.Context, *ModifySysRoleRep) (*EmptyReply, error)
	UpdateSysRole(context.Context, *ModifySysRoleRep) (*EmptyReply, error)
	DataScopeSysRole(context.Context, *DataScopeSysRoleRep) (*EmptyReply, error)
	ChangeStatusSysRole(context.Context, *ChangeStatusSysRoleRep) (*EmptyReply, error)
	DeleteSysRole(context.Context, *DeleteSysRoleRep) (*EmptyReply, error)
	GetSysRole(context.Context, *GetSysRoleRep) (*GetSysRoleReply, error)
	ListSysRole(context.Context, *ListSysRoleRep) (*ListSysRoleReply, error)
	AllocatedList(context.Context, *IsAllocatedListRep) (*IsAllocatedListReply, error)
	UnAllocatedList(context.Context, *IsAllocatedListRep) (*IsAllocatedListReply, error)
	AuthUserCancel(context.Context, *AuthUserCancelRep) (*EmptyReply, error)
	AuthUserCancelAll(context.Context, *AuthUserSelectAllRep) (*EmptyReply, error)
	AuthUserSelectAll(context.Context, *AuthUserSelectAllRep) (*EmptyReply, error)
	mustEmbedUnimplementedSysRoleServer()
}

// UnimplementedSysRoleServer must be embedded to have forward compatible implementations.
type UnimplementedSysRoleServer struct {
}

func (UnimplementedSysRoleServer) CreateSysRole(context.Context, *ModifySysRoleRep) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSysRole not implemented")
}
func (UnimplementedSysRoleServer) UpdateSysRole(context.Context, *ModifySysRoleRep) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSysRole not implemented")
}
func (UnimplementedSysRoleServer) DataScopeSysRole(context.Context, *DataScopeSysRoleRep) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DataScopeSysRole not implemented")
}
func (UnimplementedSysRoleServer) ChangeStatusSysRole(context.Context, *ChangeStatusSysRoleRep) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ChangeStatusSysRole not implemented")
}
func (UnimplementedSysRoleServer) DeleteSysRole(context.Context, *DeleteSysRoleRep) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSysRole not implemented")
}
func (UnimplementedSysRoleServer) GetSysRole(context.Context, *GetSysRoleRep) (*GetSysRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysRole not implemented")
}
func (UnimplementedSysRoleServer) ListSysRole(context.Context, *ListSysRoleRep) (*ListSysRoleReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSysRole not implemented")
}
func (UnimplementedSysRoleServer) AllocatedList(context.Context, *IsAllocatedListRep) (*IsAllocatedListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AllocatedList not implemented")
}
func (UnimplementedSysRoleServer) UnAllocatedList(context.Context, *IsAllocatedListRep) (*IsAllocatedListReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UnAllocatedList not implemented")
}
func (UnimplementedSysRoleServer) AuthUserCancel(context.Context, *AuthUserCancelRep) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthUserCancel not implemented")
}
func (UnimplementedSysRoleServer) AuthUserCancelAll(context.Context, *AuthUserSelectAllRep) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthUserCancelAll not implemented")
}
func (UnimplementedSysRoleServer) AuthUserSelectAll(context.Context, *AuthUserSelectAllRep) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AuthUserSelectAll not implemented")
}
func (UnimplementedSysRoleServer) mustEmbedUnimplementedSysRoleServer() {}

// UnsafeSysRoleServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysRoleServer will
// result in compilation errors.
type UnsafeSysRoleServer interface {
	mustEmbedUnimplementedSysRoleServer()
}

func RegisterSysRoleServer(s grpc.ServiceRegistrar, srv SysRoleServer) {
	s.RegisterService(&SysRole_ServiceDesc, srv)
}

func _SysRole_CreateSysRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySysRoleRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysRoleServer).CreateSysRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysRole_CreateSysRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysRoleServer).CreateSysRole(ctx, req.(*ModifySysRoleRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysRole_UpdateSysRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ModifySysRoleRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysRoleServer).UpdateSysRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysRole_UpdateSysRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysRoleServer).UpdateSysRole(ctx, req.(*ModifySysRoleRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysRole_DataScopeSysRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DataScopeSysRoleRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysRoleServer).DataScopeSysRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysRole_DataScopeSysRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysRoleServer).DataScopeSysRole(ctx, req.(*DataScopeSysRoleRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysRole_ChangeStatusSysRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChangeStatusSysRoleRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysRoleServer).ChangeStatusSysRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysRole_ChangeStatusSysRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysRoleServer).ChangeStatusSysRole(ctx, req.(*ChangeStatusSysRoleRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysRole_DeleteSysRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSysRoleRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysRoleServer).DeleteSysRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysRole_DeleteSysRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysRoleServer).DeleteSysRole(ctx, req.(*DeleteSysRoleRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysRole_GetSysRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSysRoleRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysRoleServer).GetSysRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysRole_GetSysRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysRoleServer).GetSysRole(ctx, req.(*GetSysRoleRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysRole_ListSysRole_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSysRoleRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysRoleServer).ListSysRole(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysRole_ListSysRole_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysRoleServer).ListSysRole(ctx, req.(*ListSysRoleRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysRole_AllocatedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAllocatedListRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysRoleServer).AllocatedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysRole_AllocatedList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysRoleServer).AllocatedList(ctx, req.(*IsAllocatedListRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysRole_UnAllocatedList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(IsAllocatedListRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysRoleServer).UnAllocatedList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysRole_UnAllocatedList_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysRoleServer).UnAllocatedList(ctx, req.(*IsAllocatedListRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysRole_AuthUserCancel_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUserCancelRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysRoleServer).AuthUserCancel(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysRole_AuthUserCancel_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysRoleServer).AuthUserCancel(ctx, req.(*AuthUserCancelRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysRole_AuthUserCancelAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUserSelectAllRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysRoleServer).AuthUserCancelAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysRole_AuthUserCancelAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysRoleServer).AuthUserCancelAll(ctx, req.(*AuthUserSelectAllRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysRole_AuthUserSelectAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthUserSelectAllRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysRoleServer).AuthUserSelectAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysRole_AuthUserSelectAll_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysRoleServer).AuthUserSelectAll(ctx, req.(*AuthUserSelectAllRep))
	}
	return interceptor(ctx, in, info, handler)
}

// SysRole_ServiceDesc is the grpc.ServiceDesc for SysRole service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SysRole_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.system.v1.SysRole",
	HandlerType: (*SysRoleServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSysRole",
			Handler:    _SysRole_CreateSysRole_Handler,
		},
		{
			MethodName: "UpdateSysRole",
			Handler:    _SysRole_UpdateSysRole_Handler,
		},
		{
			MethodName: "DataScopeSysRole",
			Handler:    _SysRole_DataScopeSysRole_Handler,
		},
		{
			MethodName: "ChangeStatusSysRole",
			Handler:    _SysRole_ChangeStatusSysRole_Handler,
		},
		{
			MethodName: "DeleteSysRole",
			Handler:    _SysRole_DeleteSysRole_Handler,
		},
		{
			MethodName: "GetSysRole",
			Handler:    _SysRole_GetSysRole_Handler,
		},
		{
			MethodName: "ListSysRole",
			Handler:    _SysRole_ListSysRole_Handler,
		},
		{
			MethodName: "AllocatedList",
			Handler:    _SysRole_AllocatedList_Handler,
		},
		{
			MethodName: "UnAllocatedList",
			Handler:    _SysRole_UnAllocatedList_Handler,
		},
		{
			MethodName: "AuthUserCancel",
			Handler:    _SysRole_AuthUserCancel_Handler,
		},
		{
			MethodName: "AuthUserCancelAll",
			Handler:    _SysRole_AuthUserCancelAll_Handler,
		},
		{
			MethodName: "AuthUserSelectAll",
			Handler:    _SysRole_AuthUserSelectAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/system/v1/sys_role.proto",
}
