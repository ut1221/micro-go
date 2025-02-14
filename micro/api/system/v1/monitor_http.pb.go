// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.1
// - protoc             v5.28.2
// source: api/system/v1/monitor.proto

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

const OperationMonitorMonitorServer = "/api.system.v1.Monitor/MonitorServer"

type MonitorHTTPServer interface {
	MonitorServer(context.Context, *MonitorServerReq) (*MonitorServerReply, error)
}

func RegisterMonitorHTTPServer(s *http.Server, srv MonitorHTTPServer) {
	r := s.Route("/")
	r.GET("/v1/monitor/server", _Monitor_MonitorServer0_HTTP_Handler(srv))
}

func _Monitor_MonitorServer0_HTTP_Handler(srv MonitorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in MonitorServerReq
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationMonitorMonitorServer)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.MonitorServer(ctx, req.(*MonitorServerReq))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*MonitorServerReply)
		return ctx.Result(200, reply)
	}
}

type MonitorHTTPClient interface {
	MonitorServer(ctx context.Context, req *MonitorServerReq, opts ...http.CallOption) (rsp *MonitorServerReply, err error)
}

type MonitorHTTPClientImpl struct {
	cc *http.Client
}

func NewMonitorHTTPClient(client *http.Client) MonitorHTTPClient {
	return &MonitorHTTPClientImpl{client}
}

func (c *MonitorHTTPClientImpl) MonitorServer(ctx context.Context, in *MonitorServerReq, opts ...http.CallOption) (*MonitorServerReply, error) {
	var out MonitorServerReply
	pattern := "/v1/monitor/server"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationMonitorMonitorServer))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
