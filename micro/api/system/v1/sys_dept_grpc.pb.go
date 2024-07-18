// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: system/v1/sys_dept.proto

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
	SysDept_CreateSysDept_FullMethodName  = "/api.system.v1.SysDept/CreateSysDept"
	SysDept_UpdateSysDept_FullMethodName  = "/api.system.v1.SysDept/UpdateSysDept"
	SysDept_DeleteSysDept_FullMethodName  = "/api.system.v1.SysDept/DeleteSysDept"
	SysDept_GetSysDept_FullMethodName     = "/api.system.v1.SysDept/GetSysDept"
	SysDept_ListSysDept_FullMethodName    = "/api.system.v1.SysDept/ListSysDept"
	SysDept_ExcludeDept_FullMethodName    = "/api.system.v1.SysDept/ExcludeDept"
	SysDept_DeptTree_FullMethodName       = "/api.system.v1.SysDept/DeptTree"
	SysDept_GetSysRoleDept_FullMethodName = "/api.system.v1.SysDept/GetSysRoleDept"
)

// SysDeptClient is the client API for SysDept service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysDeptClient interface {
	CreateSysDept(ctx context.Context, in *SysDeptRep, opts ...grpc.CallOption) (*EmptyReply, error)
	UpdateSysDept(ctx context.Context, in *SysDeptRep, opts ...grpc.CallOption) (*EmptyReply, error)
	DeleteSysDept(ctx context.Context, in *DeleteSysDeptRep, opts ...grpc.CallOption) (*EmptyReply, error)
	GetSysDept(ctx context.Context, in *GetSysDeptRep, opts ...grpc.CallOption) (*GetSysDeptReply, error)
	ListSysDept(ctx context.Context, in *ListSysDeptRep, opts ...grpc.CallOption) (*ListSysDeptReply, error)
	ExcludeDept(ctx context.Context, in *ExcludeDeptRep, opts ...grpc.CallOption) (*ListSysDeptReply, error)
	DeptTree(ctx context.Context, in *DeptTreeReq, opts ...grpc.CallOption) (*DeptTreeReply, error)
	GetSysRoleDept(ctx context.Context, in *GetSysRoleDeptRep, opts ...grpc.CallOption) (*GetSysRoleDeptReply, error)
}

type sysDeptClient struct {
	cc grpc.ClientConnInterface
}

func NewSysDeptClient(cc grpc.ClientConnInterface) SysDeptClient {
	return &sysDeptClient{cc}
}

func (c *sysDeptClient) CreateSysDept(ctx context.Context, in *SysDeptRep, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, SysDept_CreateSysDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDeptClient) UpdateSysDept(ctx context.Context, in *SysDeptRep, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, SysDept_UpdateSysDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDeptClient) DeleteSysDept(ctx context.Context, in *DeleteSysDeptRep, opts ...grpc.CallOption) (*EmptyReply, error) {
	out := new(EmptyReply)
	err := c.cc.Invoke(ctx, SysDept_DeleteSysDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDeptClient) GetSysDept(ctx context.Context, in *GetSysDeptRep, opts ...grpc.CallOption) (*GetSysDeptReply, error) {
	out := new(GetSysDeptReply)
	err := c.cc.Invoke(ctx, SysDept_GetSysDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDeptClient) ListSysDept(ctx context.Context, in *ListSysDeptRep, opts ...grpc.CallOption) (*ListSysDeptReply, error) {
	out := new(ListSysDeptReply)
	err := c.cc.Invoke(ctx, SysDept_ListSysDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDeptClient) ExcludeDept(ctx context.Context, in *ExcludeDeptRep, opts ...grpc.CallOption) (*ListSysDeptReply, error) {
	out := new(ListSysDeptReply)
	err := c.cc.Invoke(ctx, SysDept_ExcludeDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDeptClient) DeptTree(ctx context.Context, in *DeptTreeReq, opts ...grpc.CallOption) (*DeptTreeReply, error) {
	out := new(DeptTreeReply)
	err := c.cc.Invoke(ctx, SysDept_DeptTree_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDeptClient) GetSysRoleDept(ctx context.Context, in *GetSysRoleDeptRep, opts ...grpc.CallOption) (*GetSysRoleDeptReply, error) {
	out := new(GetSysRoleDeptReply)
	err := c.cc.Invoke(ctx, SysDept_GetSysRoleDept_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysDeptServer is the server API for SysDept service.
// All implementations must embed UnimplementedSysDeptServer
// for forward compatibility
type SysDeptServer interface {
	CreateSysDept(context.Context, *SysDeptRep) (*EmptyReply, error)
	UpdateSysDept(context.Context, *SysDeptRep) (*EmptyReply, error)
	DeleteSysDept(context.Context, *DeleteSysDeptRep) (*EmptyReply, error)
	GetSysDept(context.Context, *GetSysDeptRep) (*GetSysDeptReply, error)
	ListSysDept(context.Context, *ListSysDeptRep) (*ListSysDeptReply, error)
	ExcludeDept(context.Context, *ExcludeDeptRep) (*ListSysDeptReply, error)
	DeptTree(context.Context, *DeptTreeReq) (*DeptTreeReply, error)
	GetSysRoleDept(context.Context, *GetSysRoleDeptRep) (*GetSysRoleDeptReply, error)
	mustEmbedUnimplementedSysDeptServer()
}

// UnimplementedSysDeptServer must be embedded to have forward compatible implementations.
type UnimplementedSysDeptServer struct {
}

func (UnimplementedSysDeptServer) CreateSysDept(context.Context, *SysDeptRep) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSysDept not implemented")
}
func (UnimplementedSysDeptServer) UpdateSysDept(context.Context, *SysDeptRep) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSysDept not implemented")
}
func (UnimplementedSysDeptServer) DeleteSysDept(context.Context, *DeleteSysDeptRep) (*EmptyReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSysDept not implemented")
}
func (UnimplementedSysDeptServer) GetSysDept(context.Context, *GetSysDeptRep) (*GetSysDeptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysDept not implemented")
}
func (UnimplementedSysDeptServer) ListSysDept(context.Context, *ListSysDeptRep) (*ListSysDeptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSysDept not implemented")
}
func (UnimplementedSysDeptServer) ExcludeDept(context.Context, *ExcludeDeptRep) (*ListSysDeptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExcludeDept not implemented")
}
func (UnimplementedSysDeptServer) DeptTree(context.Context, *DeptTreeReq) (*DeptTreeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeptTree not implemented")
}
func (UnimplementedSysDeptServer) GetSysRoleDept(context.Context, *GetSysRoleDeptRep) (*GetSysRoleDeptReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysRoleDept not implemented")
}
func (UnimplementedSysDeptServer) mustEmbedUnimplementedSysDeptServer() {}

// UnsafeSysDeptServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysDeptServer will
// result in compilation errors.
type UnsafeSysDeptServer interface {
	mustEmbedUnimplementedSysDeptServer()
}

func RegisterSysDeptServer(s grpc.ServiceRegistrar, srv SysDeptServer) {
	s.RegisterService(&SysDept_ServiceDesc, srv)
}

func _SysDept_CreateSysDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDeptRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDeptServer).CreateSysDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDept_CreateSysDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDeptServer).CreateSysDept(ctx, req.(*SysDeptRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDept_UpdateSysDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDeptRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDeptServer).UpdateSysDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDept_UpdateSysDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDeptServer).UpdateSysDept(ctx, req.(*SysDeptRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDept_DeleteSysDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSysDeptRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDeptServer).DeleteSysDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDept_DeleteSysDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDeptServer).DeleteSysDept(ctx, req.(*DeleteSysDeptRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDept_GetSysDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSysDeptRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDeptServer).GetSysDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDept_GetSysDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDeptServer).GetSysDept(ctx, req.(*GetSysDeptRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDept_ListSysDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSysDeptRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDeptServer).ListSysDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDept_ListSysDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDeptServer).ListSysDept(ctx, req.(*ListSysDeptRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDept_ExcludeDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExcludeDeptRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDeptServer).ExcludeDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDept_ExcludeDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDeptServer).ExcludeDept(ctx, req.(*ExcludeDeptRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDept_DeptTree_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeptTreeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDeptServer).DeptTree(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDept_DeptTree_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDeptServer).DeptTree(ctx, req.(*DeptTreeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDept_GetSysRoleDept_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSysRoleDeptRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDeptServer).GetSysRoleDept(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDept_GetSysRoleDept_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDeptServer).GetSysRoleDept(ctx, req.(*GetSysRoleDeptRep))
	}
	return interceptor(ctx, in, info, handler)
}

// SysDept_ServiceDesc is the grpc.ServiceDesc for SysDept service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SysDept_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.system.v1.SysDept",
	HandlerType: (*SysDeptServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateSysDept",
			Handler:    _SysDept_CreateSysDept_Handler,
		},
		{
			MethodName: "UpdateSysDept",
			Handler:    _SysDept_UpdateSysDept_Handler,
		},
		{
			MethodName: "DeleteSysDept",
			Handler:    _SysDept_DeleteSysDept_Handler,
		},
		{
			MethodName: "GetSysDept",
			Handler:    _SysDept_GetSysDept_Handler,
		},
		{
			MethodName: "ListSysDept",
			Handler:    _SysDept_ListSysDept_Handler,
		},
		{
			MethodName: "ExcludeDept",
			Handler:    _SysDept_ExcludeDept_Handler,
		},
		{
			MethodName: "DeptTree",
			Handler:    _SysDept_DeptTree_Handler,
		},
		{
			MethodName: "GetSysRoleDept",
			Handler:    _SysDept_GetSysRoleDept_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "system/v1/sys_dept.proto",
}
