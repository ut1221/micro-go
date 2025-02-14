// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.1
// - protoc             v5.28.2
// source: api/system/v1/sys_role.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationSysRoleAllocatedList = "/api.system.v1.SysRole/AllocatedList"
const OperationSysRoleAuthUserCancel = "/api.system.v1.SysRole/AuthUserCancel"
const OperationSysRoleAuthUserCancelAll = "/api.system.v1.SysRole/AuthUserCancelAll"
const OperationSysRoleAuthUserSelectAll = "/api.system.v1.SysRole/AuthUserSelectAll"
const OperationSysRoleChangeStatusSysRole = "/api.system.v1.SysRole/ChangeStatusSysRole"
const OperationSysRoleCreateSysRole = "/api.system.v1.SysRole/CreateSysRole"
const OperationSysRoleDataScopeSysRole = "/api.system.v1.SysRole/DataScopeSysRole"
const OperationSysRoleDeleteSysRole = "/api.system.v1.SysRole/DeleteSysRole"
const OperationSysRoleGetSysRole = "/api.system.v1.SysRole/GetSysRole"
const OperationSysRoleListSysRole = "/api.system.v1.SysRole/ListSysRole"
const OperationSysRoleUnAllocatedList = "/api.system.v1.SysRole/UnAllocatedList"
const OperationSysRoleUpdateSysRole = "/api.system.v1.SysRole/UpdateSysRole"

type SysRoleHTTPServer interface {
	AllocatedList(context.Context, *IsAllocatedListRep) (*IsAllocatedListReply, error)
	AuthUserCancel(context.Context, *AuthUserCancelRep) (*EmptyReply, error)
	AuthUserCancelAll(context.Context, *AuthUserSelectAllRep) (*EmptyReply, error)
	AuthUserSelectAll(context.Context, *AuthUserSelectAllRep) (*EmptyReply, error)
	ChangeStatusSysRole(context.Context, *ChangeStatusSysRoleRep) (*EmptyReply, error)
	CreateSysRole(context.Context, *ModifySysRoleRep) (*EmptyReply, error)
	DataScopeSysRole(context.Context, *DataScopeSysRoleRep) (*EmptyReply, error)
	DeleteSysRole(context.Context, *DeleteSysRoleRep) (*EmptyReply, error)
	GetSysRole(context.Context, *GetSysRoleRep) (*GetSysRoleReply, error)
	ListSysRole(context.Context, *ListSysRoleRep) (*ListSysRoleReply, error)
	UnAllocatedList(context.Context, *IsAllocatedListRep) (*IsAllocatedListReply, error)
	UpdateSysRole(context.Context, *ModifySysRoleRep) (*EmptyReply, error)
}

func RegisterSysRoleHTTPServer(s *http.Server, srv SysRoleHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/role/save", _SysRole_CreateSysRole0_HTTP_Handler(srv))
	r.PUT("/v1/role/update", _SysRole_UpdateSysRole0_HTTP_Handler(srv))
	r.PUT("/v1/role/dataScope", _SysRole_DataScopeSysRole0_HTTP_Handler(srv))
	r.PUT("/v1/role/changeStatus", _SysRole_ChangeStatusSysRole0_HTTP_Handler(srv))
	r.DELETE("/v1/role/delete/{id}", _SysRole_DeleteSysRole0_HTTP_Handler(srv))
	r.GET("/v1/role/info/{roleId}", _SysRole_GetSysRole0_HTTP_Handler(srv))
	r.POST("/v1/role/list", _SysRole_ListSysRole0_HTTP_Handler(srv))
	r.POST("/v1/role/authUser/allocatedList", _SysRole_AllocatedList0_HTTP_Handler(srv))
	r.POST("/v1/role/authUser/unallocatedList", _SysRole_UnAllocatedList0_HTTP_Handler(srv))
	r.PUT("/v1/role/authUser/cancel", _SysRole_AuthUserCancel0_HTTP_Handler(srv))
	r.PUT("/v1/role/authUser/cancelAll", _SysRole_AuthUserCancelAll0_HTTP_Handler(srv))
	r.PUT("/v1/role/authUser/selectAll", _SysRole_AuthUserSelectAll0_HTTP_Handler(srv))
}

func _SysRole_CreateSysRole0_HTTP_Handler(srv SysRoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ModifySysRoleRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysRoleCreateSysRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSysRole(ctx, req.(*ModifySysRoleRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _SysRole_UpdateSysRole0_HTTP_Handler(srv SysRoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ModifySysRoleRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysRoleUpdateSysRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateSysRole(ctx, req.(*ModifySysRoleRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _SysRole_DataScopeSysRole0_HTTP_Handler(srv SysRoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DataScopeSysRoleRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysRoleDataScopeSysRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DataScopeSysRole(ctx, req.(*DataScopeSysRoleRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _SysRole_ChangeStatusSysRole0_HTTP_Handler(srv SysRoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ChangeStatusSysRoleRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysRoleChangeStatusSysRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ChangeStatusSysRole(ctx, req.(*ChangeStatusSysRoleRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _SysRole_DeleteSysRole0_HTTP_Handler(srv SysRoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteSysRoleRep
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysRoleDeleteSysRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteSysRole(ctx, req.(*DeleteSysRoleRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _SysRole_GetSysRole0_HTTP_Handler(srv SysRoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSysRoleRep
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysRoleGetSysRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSysRole(ctx, req.(*GetSysRoleRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSysRoleReply)
		return ctx.Result(200, reply)
	}
}

func _SysRole_ListSysRole0_HTTP_Handler(srv SysRoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListSysRoleRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysRoleListSysRole)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListSysRole(ctx, req.(*ListSysRoleRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSysRoleReply)
		return ctx.Result(200, reply)
	}
}

func _SysRole_AllocatedList0_HTTP_Handler(srv SysRoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IsAllocatedListRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysRoleAllocatedList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AllocatedList(ctx, req.(*IsAllocatedListRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IsAllocatedListReply)
		return ctx.Result(200, reply)
	}
}

func _SysRole_UnAllocatedList0_HTTP_Handler(srv SysRoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IsAllocatedListRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysRoleUnAllocatedList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UnAllocatedList(ctx, req.(*IsAllocatedListRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IsAllocatedListReply)
		return ctx.Result(200, reply)
	}
}

func _SysRole_AuthUserCancel0_HTTP_Handler(srv SysRoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuthUserCancelRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysRoleAuthUserCancel)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AuthUserCancel(ctx, req.(*AuthUserCancelRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _SysRole_AuthUserCancelAll0_HTTP_Handler(srv SysRoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuthUserSelectAllRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysRoleAuthUserCancelAll)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AuthUserCancelAll(ctx, req.(*AuthUserSelectAllRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _SysRole_AuthUserSelectAll0_HTTP_Handler(srv SysRoleHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuthUserSelectAllRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysRoleAuthUserSelectAll)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AuthUserSelectAll(ctx, req.(*AuthUserSelectAllRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

type SysRoleHTTPClient interface {
	AllocatedList(ctx context.Context, req *IsAllocatedListRep, opts ...http.CallOption) (rsp *IsAllocatedListReply, err error)
	AuthUserCancel(ctx context.Context, req *AuthUserCancelRep, opts ...http.CallOption) (rsp *EmptyReply, err error)
	AuthUserCancelAll(ctx context.Context, req *AuthUserSelectAllRep, opts ...http.CallOption) (rsp *EmptyReply, err error)
	AuthUserSelectAll(ctx context.Context, req *AuthUserSelectAllRep, opts ...http.CallOption) (rsp *EmptyReply, err error)
	ChangeStatusSysRole(ctx context.Context, req *ChangeStatusSysRoleRep, opts ...http.CallOption) (rsp *EmptyReply, err error)
	CreateSysRole(ctx context.Context, req *ModifySysRoleRep, opts ...http.CallOption) (rsp *EmptyReply, err error)
	DataScopeSysRole(ctx context.Context, req *DataScopeSysRoleRep, opts ...http.CallOption) (rsp *EmptyReply, err error)
	DeleteSysRole(ctx context.Context, req *DeleteSysRoleRep, opts ...http.CallOption) (rsp *EmptyReply, err error)
	GetSysRole(ctx context.Context, req *GetSysRoleRep, opts ...http.CallOption) (rsp *GetSysRoleReply, err error)
	ListSysRole(ctx context.Context, req *ListSysRoleRep, opts ...http.CallOption) (rsp *ListSysRoleReply, err error)
	UnAllocatedList(ctx context.Context, req *IsAllocatedListRep, opts ...http.CallOption) (rsp *IsAllocatedListReply, err error)
	UpdateSysRole(ctx context.Context, req *ModifySysRoleRep, opts ...http.CallOption) (rsp *EmptyReply, err error)
}

type SysRoleHTTPClientImpl struct {
	cc *http.Client
}

func NewSysRoleHTTPClient(client *http.Client) SysRoleHTTPClient {
	return &SysRoleHTTPClientImpl{client}
}

func (c *SysRoleHTTPClientImpl) AllocatedList(ctx context.Context, in *IsAllocatedListRep, opts ...http.CallOption) (*IsAllocatedListReply, error) {
	var out IsAllocatedListReply
	pattern := "/v1/role/authUser/allocatedList"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysRoleAllocatedList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysRoleHTTPClientImpl) AuthUserCancel(ctx context.Context, in *AuthUserCancelRep, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/role/authUser/cancel"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysRoleAuthUserCancel))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysRoleHTTPClientImpl) AuthUserCancelAll(ctx context.Context, in *AuthUserSelectAllRep, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/role/authUser/cancelAll"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysRoleAuthUserCancelAll))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysRoleHTTPClientImpl) AuthUserSelectAll(ctx context.Context, in *AuthUserSelectAllRep, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/role/authUser/selectAll"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysRoleAuthUserSelectAll))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysRoleHTTPClientImpl) ChangeStatusSysRole(ctx context.Context, in *ChangeStatusSysRoleRep, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/role/changeStatus"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysRoleChangeStatusSysRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysRoleHTTPClientImpl) CreateSysRole(ctx context.Context, in *ModifySysRoleRep, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/role/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysRoleCreateSysRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysRoleHTTPClientImpl) DataScopeSysRole(ctx context.Context, in *DataScopeSysRoleRep, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/role/dataScope"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysRoleDataScopeSysRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysRoleHTTPClientImpl) DeleteSysRole(ctx context.Context, in *DeleteSysRoleRep, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/role/delete/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSysRoleDeleteSysRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysRoleHTTPClientImpl) GetSysRole(ctx context.Context, in *GetSysRoleRep, opts ...http.CallOption) (*GetSysRoleReply, error) {
	var out GetSysRoleReply
	pattern := "/v1/role/info/{roleId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSysRoleGetSysRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysRoleHTTPClientImpl) ListSysRole(ctx context.Context, in *ListSysRoleRep, opts ...http.CallOption) (*ListSysRoleReply, error) {
	var out ListSysRoleReply
	pattern := "/v1/role/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysRoleListSysRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysRoleHTTPClientImpl) UnAllocatedList(ctx context.Context, in *IsAllocatedListRep, opts ...http.CallOption) (*IsAllocatedListReply, error) {
	var out IsAllocatedListReply
	pattern := "/v1/role/authUser/unallocatedList"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysRoleUnAllocatedList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysRoleHTTPClientImpl) UpdateSysRole(ctx context.Context, in *ModifySysRoleRep, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/role/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysRoleUpdateSysRole))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
