// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: api/system/v1/sys_dict.proto

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
	SysDict_SysDictDataType_FullMethodName     = "/api.system.v1.SysDict/SysDictDataType"
	SysDict_CreateSysDictType_FullMethodName   = "/api.system.v1.SysDict/CreateSysDictType"
	SysDict_UpdateSysDictType_FullMethodName   = "/api.system.v1.SysDict/UpdateSysDictType"
	SysDict_DeleteSysDictType_FullMethodName   = "/api.system.v1.SysDict/DeleteSysDictType"
	SysDict_GetSysDictType_FullMethodName      = "/api.system.v1.SysDict/GetSysDictType"
	SysDict_ListSysDictType_FullMethodName     = "/api.system.v1.SysDict/ListSysDictType"
	SysDict_OptionSelectType_FullMethodName    = "/api.system.v1.SysDict/OptionSelectType"
	SysDict_CreateSysDictData_FullMethodName   = "/api.system.v1.SysDict/CreateSysDictData"
	SysDict_UpdateSysDictData_FullMethodName   = "/api.system.v1.SysDict/UpdateSysDictData"
	SysDict_DeleteSysDictData_FullMethodName   = "/api.system.v1.SysDict/DeleteSysDictData"
	SysDict_GetSysDictData_FullMethodName      = "/api.system.v1.SysDict/GetSysDictData"
	SysDict_ListSysDictData_FullMethodName     = "/api.system.v1.SysDict/ListSysDictData"
	SysDict_RefreshCacheSysDict_FullMethodName = "/api.system.v1.SysDict/RefreshCacheSysDict"
)

// SysDictClient is the client API for SysDict service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SysDictClient interface {
	SysDictDataType(ctx context.Context, in *SysDictDataTypeReq, opts ...grpc.CallOption) (*SysDictDataTypeReply, error)
	CreateSysDictType(ctx context.Context, in *SysDictTypeRep, opts ...grpc.CallOption) (*EmptySysDictReply, error)
	UpdateSysDictType(ctx context.Context, in *SysDictTypeRep, opts ...grpc.CallOption) (*EmptySysDictReply, error)
	DeleteSysDictType(ctx context.Context, in *DeleteSysDictTypeRep, opts ...grpc.CallOption) (*EmptySysDictReply, error)
	GetSysDictType(ctx context.Context, in *GetSysDictRep, opts ...grpc.CallOption) (*GetSysDictTypeReply, error)
	ListSysDictType(ctx context.Context, in *ListSysDictTypeRep, opts ...grpc.CallOption) (*ListSysDictTypeReply, error)
	OptionSelectType(ctx context.Context, in *OptionSelectTypeRep, opts ...grpc.CallOption) (*OptionSelectTypeReply, error)
	CreateSysDictData(ctx context.Context, in *SysDictDataRep, opts ...grpc.CallOption) (*EmptySysDictReply, error)
	UpdateSysDictData(ctx context.Context, in *SysDictDataRep, opts ...grpc.CallOption) (*EmptySysDictReply, error)
	DeleteSysDictData(ctx context.Context, in *DeleteSysDictDataRep, opts ...grpc.CallOption) (*EmptySysDictReply, error)
	GetSysDictData(ctx context.Context, in *GetSysDictRep, opts ...grpc.CallOption) (*GetSysDictDataReply, error)
	ListSysDictData(ctx context.Context, in *ListSysDictDataRep, opts ...grpc.CallOption) (*ListSysDictDataReply, error)
	RefreshCacheSysDict(ctx context.Context, in *CacheSysDicReq, opts ...grpc.CallOption) (*EmptySysDictReply, error)
}

type sysDictClient struct {
	cc grpc.ClientConnInterface
}

func NewSysDictClient(cc grpc.ClientConnInterface) SysDictClient {
	return &sysDictClient{cc}
}

func (c *sysDictClient) SysDictDataType(ctx context.Context, in *SysDictDataTypeReq, opts ...grpc.CallOption) (*SysDictDataTypeReply, error) {
	out := new(SysDictDataTypeReply)
	err := c.cc.Invoke(ctx, SysDict_SysDictDataType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDictClient) CreateSysDictType(ctx context.Context, in *SysDictTypeRep, opts ...grpc.CallOption) (*EmptySysDictReply, error) {
	out := new(EmptySysDictReply)
	err := c.cc.Invoke(ctx, SysDict_CreateSysDictType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDictClient) UpdateSysDictType(ctx context.Context, in *SysDictTypeRep, opts ...grpc.CallOption) (*EmptySysDictReply, error) {
	out := new(EmptySysDictReply)
	err := c.cc.Invoke(ctx, SysDict_UpdateSysDictType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDictClient) DeleteSysDictType(ctx context.Context, in *DeleteSysDictTypeRep, opts ...grpc.CallOption) (*EmptySysDictReply, error) {
	out := new(EmptySysDictReply)
	err := c.cc.Invoke(ctx, SysDict_DeleteSysDictType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDictClient) GetSysDictType(ctx context.Context, in *GetSysDictRep, opts ...grpc.CallOption) (*GetSysDictTypeReply, error) {
	out := new(GetSysDictTypeReply)
	err := c.cc.Invoke(ctx, SysDict_GetSysDictType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDictClient) ListSysDictType(ctx context.Context, in *ListSysDictTypeRep, opts ...grpc.CallOption) (*ListSysDictTypeReply, error) {
	out := new(ListSysDictTypeReply)
	err := c.cc.Invoke(ctx, SysDict_ListSysDictType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDictClient) OptionSelectType(ctx context.Context, in *OptionSelectTypeRep, opts ...grpc.CallOption) (*OptionSelectTypeReply, error) {
	out := new(OptionSelectTypeReply)
	err := c.cc.Invoke(ctx, SysDict_OptionSelectType_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDictClient) CreateSysDictData(ctx context.Context, in *SysDictDataRep, opts ...grpc.CallOption) (*EmptySysDictReply, error) {
	out := new(EmptySysDictReply)
	err := c.cc.Invoke(ctx, SysDict_CreateSysDictData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDictClient) UpdateSysDictData(ctx context.Context, in *SysDictDataRep, opts ...grpc.CallOption) (*EmptySysDictReply, error) {
	out := new(EmptySysDictReply)
	err := c.cc.Invoke(ctx, SysDict_UpdateSysDictData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDictClient) DeleteSysDictData(ctx context.Context, in *DeleteSysDictDataRep, opts ...grpc.CallOption) (*EmptySysDictReply, error) {
	out := new(EmptySysDictReply)
	err := c.cc.Invoke(ctx, SysDict_DeleteSysDictData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDictClient) GetSysDictData(ctx context.Context, in *GetSysDictRep, opts ...grpc.CallOption) (*GetSysDictDataReply, error) {
	out := new(GetSysDictDataReply)
	err := c.cc.Invoke(ctx, SysDict_GetSysDictData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDictClient) ListSysDictData(ctx context.Context, in *ListSysDictDataRep, opts ...grpc.CallOption) (*ListSysDictDataReply, error) {
	out := new(ListSysDictDataReply)
	err := c.cc.Invoke(ctx, SysDict_ListSysDictData_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *sysDictClient) RefreshCacheSysDict(ctx context.Context, in *CacheSysDicReq, opts ...grpc.CallOption) (*EmptySysDictReply, error) {
	out := new(EmptySysDictReply)
	err := c.cc.Invoke(ctx, SysDict_RefreshCacheSysDict_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SysDictServer is the server API for SysDict service.
// All implementations must embed UnimplementedSysDictServer
// for forward compatibility
type SysDictServer interface {
	SysDictDataType(context.Context, *SysDictDataTypeReq) (*SysDictDataTypeReply, error)
	CreateSysDictType(context.Context, *SysDictTypeRep) (*EmptySysDictReply, error)
	UpdateSysDictType(context.Context, *SysDictTypeRep) (*EmptySysDictReply, error)
	DeleteSysDictType(context.Context, *DeleteSysDictTypeRep) (*EmptySysDictReply, error)
	GetSysDictType(context.Context, *GetSysDictRep) (*GetSysDictTypeReply, error)
	ListSysDictType(context.Context, *ListSysDictTypeRep) (*ListSysDictTypeReply, error)
	OptionSelectType(context.Context, *OptionSelectTypeRep) (*OptionSelectTypeReply, error)
	CreateSysDictData(context.Context, *SysDictDataRep) (*EmptySysDictReply, error)
	UpdateSysDictData(context.Context, *SysDictDataRep) (*EmptySysDictReply, error)
	DeleteSysDictData(context.Context, *DeleteSysDictDataRep) (*EmptySysDictReply, error)
	GetSysDictData(context.Context, *GetSysDictRep) (*GetSysDictDataReply, error)
	ListSysDictData(context.Context, *ListSysDictDataRep) (*ListSysDictDataReply, error)
	RefreshCacheSysDict(context.Context, *CacheSysDicReq) (*EmptySysDictReply, error)
	mustEmbedUnimplementedSysDictServer()
}

// UnimplementedSysDictServer must be embedded to have forward compatible implementations.
type UnimplementedSysDictServer struct {
}

func (UnimplementedSysDictServer) SysDictDataType(context.Context, *SysDictDataTypeReq) (*SysDictDataTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SysDictDataType not implemented")
}
func (UnimplementedSysDictServer) CreateSysDictType(context.Context, *SysDictTypeRep) (*EmptySysDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSysDictType not implemented")
}
func (UnimplementedSysDictServer) UpdateSysDictType(context.Context, *SysDictTypeRep) (*EmptySysDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSysDictType not implemented")
}
func (UnimplementedSysDictServer) DeleteSysDictType(context.Context, *DeleteSysDictTypeRep) (*EmptySysDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSysDictType not implemented")
}
func (UnimplementedSysDictServer) GetSysDictType(context.Context, *GetSysDictRep) (*GetSysDictTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysDictType not implemented")
}
func (UnimplementedSysDictServer) ListSysDictType(context.Context, *ListSysDictTypeRep) (*ListSysDictTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSysDictType not implemented")
}
func (UnimplementedSysDictServer) OptionSelectType(context.Context, *OptionSelectTypeRep) (*OptionSelectTypeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method OptionSelectType not implemented")
}
func (UnimplementedSysDictServer) CreateSysDictData(context.Context, *SysDictDataRep) (*EmptySysDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSysDictData not implemented")
}
func (UnimplementedSysDictServer) UpdateSysDictData(context.Context, *SysDictDataRep) (*EmptySysDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSysDictData not implemented")
}
func (UnimplementedSysDictServer) DeleteSysDictData(context.Context, *DeleteSysDictDataRep) (*EmptySysDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSysDictData not implemented")
}
func (UnimplementedSysDictServer) GetSysDictData(context.Context, *GetSysDictRep) (*GetSysDictDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSysDictData not implemented")
}
func (UnimplementedSysDictServer) ListSysDictData(context.Context, *ListSysDictDataRep) (*ListSysDictDataReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListSysDictData not implemented")
}
func (UnimplementedSysDictServer) RefreshCacheSysDict(context.Context, *CacheSysDicReq) (*EmptySysDictReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshCacheSysDict not implemented")
}
func (UnimplementedSysDictServer) mustEmbedUnimplementedSysDictServer() {}

// UnsafeSysDictServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SysDictServer will
// result in compilation errors.
type UnsafeSysDictServer interface {
	mustEmbedUnimplementedSysDictServer()
}

func RegisterSysDictServer(s grpc.ServiceRegistrar, srv SysDictServer) {
	s.RegisterService(&SysDict_ServiceDesc, srv)
}

func _SysDict_SysDictDataType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDictDataTypeReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDictServer).SysDictDataType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDict_SysDictDataType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDictServer).SysDictDataType(ctx, req.(*SysDictDataTypeReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDict_CreateSysDictType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDictTypeRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDictServer).CreateSysDictType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDict_CreateSysDictType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDictServer).CreateSysDictType(ctx, req.(*SysDictTypeRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDict_UpdateSysDictType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDictTypeRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDictServer).UpdateSysDictType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDict_UpdateSysDictType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDictServer).UpdateSysDictType(ctx, req.(*SysDictTypeRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDict_DeleteSysDictType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSysDictTypeRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDictServer).DeleteSysDictType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDict_DeleteSysDictType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDictServer).DeleteSysDictType(ctx, req.(*DeleteSysDictTypeRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDict_GetSysDictType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSysDictRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDictServer).GetSysDictType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDict_GetSysDictType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDictServer).GetSysDictType(ctx, req.(*GetSysDictRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDict_ListSysDictType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSysDictTypeRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDictServer).ListSysDictType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDict_ListSysDictType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDictServer).ListSysDictType(ctx, req.(*ListSysDictTypeRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDict_OptionSelectType_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OptionSelectTypeRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDictServer).OptionSelectType(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDict_OptionSelectType_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDictServer).OptionSelectType(ctx, req.(*OptionSelectTypeRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDict_CreateSysDictData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDictDataRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDictServer).CreateSysDictData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDict_CreateSysDictData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDictServer).CreateSysDictData(ctx, req.(*SysDictDataRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDict_UpdateSysDictData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SysDictDataRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDictServer).UpdateSysDictData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDict_UpdateSysDictData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDictServer).UpdateSysDictData(ctx, req.(*SysDictDataRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDict_DeleteSysDictData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteSysDictDataRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDictServer).DeleteSysDictData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDict_DeleteSysDictData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDictServer).DeleteSysDictData(ctx, req.(*DeleteSysDictDataRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDict_GetSysDictData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetSysDictRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDictServer).GetSysDictData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDict_GetSysDictData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDictServer).GetSysDictData(ctx, req.(*GetSysDictRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDict_ListSysDictData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListSysDictDataRep)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDictServer).ListSysDictData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDict_ListSysDictData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDictServer).ListSysDictData(ctx, req.(*ListSysDictDataRep))
	}
	return interceptor(ctx, in, info, handler)
}

func _SysDict_RefreshCacheSysDict_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CacheSysDicReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SysDictServer).RefreshCacheSysDict(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: SysDict_RefreshCacheSysDict_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SysDictServer).RefreshCacheSysDict(ctx, req.(*CacheSysDicReq))
	}
	return interceptor(ctx, in, info, handler)
}

// SysDict_ServiceDesc is the grpc.ServiceDesc for SysDict service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SysDict_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.system.v1.SysDict",
	HandlerType: (*SysDictServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SysDictDataType",
			Handler:    _SysDict_SysDictDataType_Handler,
		},
		{
			MethodName: "CreateSysDictType",
			Handler:    _SysDict_CreateSysDictType_Handler,
		},
		{
			MethodName: "UpdateSysDictType",
			Handler:    _SysDict_UpdateSysDictType_Handler,
		},
		{
			MethodName: "DeleteSysDictType",
			Handler:    _SysDict_DeleteSysDictType_Handler,
		},
		{
			MethodName: "GetSysDictType",
			Handler:    _SysDict_GetSysDictType_Handler,
		},
		{
			MethodName: "ListSysDictType",
			Handler:    _SysDict_ListSysDictType_Handler,
		},
		{
			MethodName: "OptionSelectType",
			Handler:    _SysDict_OptionSelectType_Handler,
		},
		{
			MethodName: "CreateSysDictData",
			Handler:    _SysDict_CreateSysDictData_Handler,
		},
		{
			MethodName: "UpdateSysDictData",
			Handler:    _SysDict_UpdateSysDictData_Handler,
		},
		{
			MethodName: "DeleteSysDictData",
			Handler:    _SysDict_DeleteSysDictData_Handler,
		},
		{
			MethodName: "GetSysDictData",
			Handler:    _SysDict_GetSysDictData_Handler,
		},
		{
			MethodName: "ListSysDictData",
			Handler:    _SysDict_ListSysDictData_Handler,
		},
		{
			MethodName: "RefreshCacheSysDict",
			Handler:    _SysDict_RefreshCacheSysDict_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/system/v1/sys_dict.proto",
}
