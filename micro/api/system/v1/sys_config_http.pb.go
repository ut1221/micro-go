// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.1
// - protoc             v5.28.2
// source: api/system/v1/sys_config.proto

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

const OperationSysConfigConfigByKey = "/api.system.v1.SysConfig/ConfigByKey"
const OperationSysConfigCreateSysConfig = "/api.system.v1.SysConfig/CreateSysConfig"
const OperationSysConfigListSysConfig = "/api.system.v1.SysConfig/ListSysConfig"

type SysConfigHTTPServer interface {
	ConfigByKey(context.Context, *ConfigByKeyReq) (*ConfigByKeyReply, error)
	CreateSysConfig(context.Context, *CreateSysConfigRep) (*EmptySysConfigReply, error)
	ListSysConfig(context.Context, *ListSysConfigRep) (*ListSysConfigReply, error)
}

func RegisterSysConfigHTTPServer(s *http.Server, srv SysConfigHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/config/create", _SysConfig_CreateSysConfig0_HTTP_Handler(srv))
	r.POST("/v1/config/list", _SysConfig_ListSysConfig0_HTTP_Handler(srv))
	r.GET("/v1/config/configKey/{key}", _SysConfig_ConfigByKey0_HTTP_Handler(srv))
}

func _SysConfig_CreateSysConfig0_HTTP_Handler(srv SysConfigHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateSysConfigRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysConfigCreateSysConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSysConfig(ctx, req.(*CreateSysConfigRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*EmptySysConfigReply)
		return ctx.Result(200, reply)
	}
}

func _SysConfig_ListSysConfig0_HTTP_Handler(srv SysConfigHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListSysConfigRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysConfigListSysConfig)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListSysConfig(ctx, req.(*ListSysConfigRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSysConfigReply)
		return ctx.Result(200, reply)
	}
}

func _SysConfig_ConfigByKey0_HTTP_Handler(srv SysConfigHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ConfigByKeyReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysConfigConfigByKey)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ConfigByKey(ctx, req.(*ConfigByKeyReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ConfigByKeyReply)
		return ctx.Result(200, reply)
	}
}

type SysConfigHTTPClient interface {
	ConfigByKey(ctx context.Context, req *ConfigByKeyReq, opts ...http.CallOption) (rsp *ConfigByKeyReply, err error)
	CreateSysConfig(ctx context.Context, req *CreateSysConfigRep, opts ...http.CallOption) (rsp *EmptySysConfigReply, err error)
	ListSysConfig(ctx context.Context, req *ListSysConfigRep, opts ...http.CallOption) (rsp *ListSysConfigReply, err error)
}

type SysConfigHTTPClientImpl struct {
	cc *http.Client
}

func NewSysConfigHTTPClient(client *http.Client) SysConfigHTTPClient {
	return &SysConfigHTTPClientImpl{client}
}

func (c *SysConfigHTTPClientImpl) ConfigByKey(ctx context.Context, in *ConfigByKeyReq, opts ...http.CallOption) (*ConfigByKeyReply, error) {
	var out ConfigByKeyReply
	pattern := "/v1/config/configKey/{key}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSysConfigConfigByKey))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysConfigHTTPClientImpl) CreateSysConfig(ctx context.Context, in *CreateSysConfigRep, opts ...http.CallOption) (*EmptySysConfigReply, error) {
	var out EmptySysConfigReply
	pattern := "/v1/config/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysConfigCreateSysConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysConfigHTTPClientImpl) ListSysConfig(ctx context.Context, in *ListSysConfigRep, opts ...http.CallOption) (*ListSysConfigReply, error) {
	var out ListSysConfigReply
	pattern := "/v1/config/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysConfigListSysConfig))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
