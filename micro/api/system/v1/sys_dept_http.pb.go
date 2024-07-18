// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.25.3
// source: system/v1/sys_dept.proto

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

const OperationSysDeptCreateSysDept = "/api.system.v1.SysDept/CreateSysDept"
const OperationSysDeptDeleteSysDept = "/api.system.v1.SysDept/DeleteSysDept"
const OperationSysDeptExcludeDept = "/api.system.v1.SysDept/ExcludeDept"
const OperationSysDeptGetSysDept = "/api.system.v1.SysDept/GetSysDept"
const OperationSysDeptGetSysRoleDept = "/api.system.v1.SysDept/GetSysRoleDept"
const OperationSysDeptListSysDept = "/api.system.v1.SysDept/ListSysDept"
const OperationSysDeptUpdateSysDept = "/api.system.v1.SysDept/UpdateSysDept"

type SysDeptHTTPServer interface {
	CreateSysDept(context.Context, *SysDeptRep) (*EmptyReply, error)
	DeleteSysDept(context.Context, *DeleteSysDeptRep) (*EmptyReply, error)
	ExcludeDept(context.Context, *ExcludeDeptRep) (*ListSysDeptReply, error)
	GetSysDept(context.Context, *GetSysDeptRep) (*GetSysDeptReply, error)
	GetSysRoleDept(context.Context, *GetSysRoleDeptRep) (*GetSysRoleDeptReply, error)
	ListSysDept(context.Context, *ListSysDeptRep) (*ListSysDeptReply, error)
	UpdateSysDept(context.Context, *SysDeptRep) (*EmptyReply, error)
}

func RegisterSysDeptHTTPServer(s *http.Server, srv SysDeptHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/dept/save", _SysDept_CreateSysDept0_HTTP_Handler(srv))
	r.PUT("/v1/dept/update", _SysDept_UpdateSysDept0_HTTP_Handler(srv))
	r.DELETE("/v1/dept/delete/{id}", _SysDept_DeleteSysDept0_HTTP_Handler(srv))
	r.GET("/v1/dept/info/{id}", _SysDept_GetSysDept0_HTTP_Handler(srv))
	r.POST("/v1/dept/list", _SysDept_ListSysDept0_HTTP_Handler(srv))
	r.GET("/v1/dept/list/exclude/{id}", _SysDept_ExcludeDept0_HTTP_Handler(srv))
	r.GET("/v1/dept/deptTreeByRoleId/{roleId}", _SysDept_GetSysRoleDept0_HTTP_Handler(srv))
}

func _SysDept_CreateSysDept0_HTTP_Handler(srv SysDeptHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysDeptRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysDeptCreateSysDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSysDept(ctx, req.(*SysDeptRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _SysDept_UpdateSysDept0_HTTP_Handler(srv SysDeptHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SysDeptRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysDeptUpdateSysDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateSysDept(ctx, req.(*SysDeptRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _SysDept_DeleteSysDept0_HTTP_Handler(srv SysDeptHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteSysDeptRep
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysDeptDeleteSysDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteSysDept(ctx, req.(*DeleteSysDeptRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptyReply)
		return ctx.Result(200, reply)
	}
}

func _SysDept_GetSysDept0_HTTP_Handler(srv SysDeptHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSysDeptRep
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysDeptGetSysDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSysDept(ctx, req.(*GetSysDeptRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSysDeptReply)
		return ctx.Result(200, reply)
	}
}

func _SysDept_ListSysDept0_HTTP_Handler(srv SysDeptHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListSysDeptRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysDeptListSysDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListSysDept(ctx, req.(*ListSysDeptRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSysDeptReply)
		return ctx.Result(200, reply)
	}
}

func _SysDept_ExcludeDept0_HTTP_Handler(srv SysDeptHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExcludeDeptRep
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysDeptExcludeDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ExcludeDept(ctx, req.(*ExcludeDeptRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSysDeptReply)
		return ctx.Result(200, reply)
	}
}

func _SysDept_GetSysRoleDept0_HTTP_Handler(srv SysDeptHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSysRoleDeptRep
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysDeptGetSysRoleDept)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSysRoleDept(ctx, req.(*GetSysRoleDeptRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSysRoleDeptReply)
		return ctx.Result(200, reply)
	}
}

type SysDeptHTTPClient interface {
	CreateSysDept(ctx context.Context, req *SysDeptRep, opts ...http.CallOption) (rsp *EmptyReply, err error)
	DeleteSysDept(ctx context.Context, req *DeleteSysDeptRep, opts ...http.CallOption) (rsp *EmptyReply, err error)
	ExcludeDept(ctx context.Context, req *ExcludeDeptRep, opts ...http.CallOption) (rsp *ListSysDeptReply, err error)
	GetSysDept(ctx context.Context, req *GetSysDeptRep, opts ...http.CallOption) (rsp *GetSysDeptReply, err error)
	GetSysRoleDept(ctx context.Context, req *GetSysRoleDeptRep, opts ...http.CallOption) (rsp *GetSysRoleDeptReply, err error)
	ListSysDept(ctx context.Context, req *ListSysDeptRep, opts ...http.CallOption) (rsp *ListSysDeptReply, err error)
	UpdateSysDept(ctx context.Context, req *SysDeptRep, opts ...http.CallOption) (rsp *EmptyReply, err error)
}

type SysDeptHTTPClientImpl struct {
	cc *http.Client
}

func NewSysDeptHTTPClient(client *http.Client) SysDeptHTTPClient {
	return &SysDeptHTTPClientImpl{client}
}

func (c *SysDeptHTTPClientImpl) CreateSysDept(ctx context.Context, in *SysDeptRep, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/dept/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysDeptCreateSysDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysDeptHTTPClientImpl) DeleteSysDept(ctx context.Context, in *DeleteSysDeptRep, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/dept/delete/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSysDeptDeleteSysDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysDeptHTTPClientImpl) ExcludeDept(ctx context.Context, in *ExcludeDeptRep, opts ...http.CallOption) (*ListSysDeptReply, error) {
	var out ListSysDeptReply
	pattern := "/v1/dept/list/exclude/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSysDeptExcludeDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysDeptHTTPClientImpl) GetSysDept(ctx context.Context, in *GetSysDeptRep, opts ...http.CallOption) (*GetSysDeptReply, error) {
	var out GetSysDeptReply
	pattern := "/v1/dept/info/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSysDeptGetSysDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysDeptHTTPClientImpl) GetSysRoleDept(ctx context.Context, in *GetSysRoleDeptRep, opts ...http.CallOption) (*GetSysRoleDeptReply, error) {
	var out GetSysRoleDeptReply
	pattern := "/v1/dept/deptTreeByRoleId/{roleId}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSysDeptGetSysRoleDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysDeptHTTPClientImpl) ListSysDept(ctx context.Context, in *ListSysDeptRep, opts ...http.CallOption) (*ListSysDeptReply, error) {
	var out ListSysDeptReply
	pattern := "/v1/dept/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysDeptListSysDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysDeptHTTPClientImpl) UpdateSysDept(ctx context.Context, in *SysDeptRep, opts ...http.CallOption) (*EmptyReply, error) {
	var out EmptyReply
	pattern := "/v1/dept/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysDeptUpdateSysDept))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
