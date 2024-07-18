// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.3
// - protoc             v4.25.3
// source: system/v1/sys_user.proto

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

const OperationSysUserAuthRoleSysUser = "/api.system.v1.SysUser/AuthRoleSysUser"
const OperationSysUserCreateSysUser = "/api.system.v1.SysUser/CreateSysUser"
const OperationSysUserDeleteSysUser = "/api.system.v1.SysUser/DeleteSysUser"
const OperationSysUserGetAuthRoleSysUser = "/api.system.v1.SysUser/GetAuthRoleSysUser"
const OperationSysUserGetOtherInfo = "/api.system.v1.SysUser/GetOtherInfo"
const OperationSysUserGetSysUser = "/api.system.v1.SysUser/GetSysUser"
const OperationSysUserListSysUser = "/api.system.v1.SysUser/ListSysUser"
const OperationSysUserProfile = "/api.system.v1.SysUser/Profile"
const OperationSysUserResetPwd = "/api.system.v1.SysUser/ResetPwd"
const OperationSysUserSaveSysUser = "/api.system.v1.SysUser/SaveSysUser"
const OperationSysUserUpdateSysUser = "/api.system.v1.SysUser/UpdateSysUser"

type SysUserHTTPServer interface {
	AuthRoleSysUser(context.Context, *AuthRoleSysUserRep) (*AuthRoleSysUserReply, error)
	CreateSysUser(context.Context, *CreateSysUserRep) (*CreateSysUserReply, error)
	DeleteSysUser(context.Context, *DeleteSysUserRep) (*DeleteSysUserReply, error)
	GetAuthRoleSysUser(context.Context, *GetAuthRoleSysUserRep) (*GetAuthRoleSysUserReply, error)
	// GetOtherInfo获取角色以及岗位详细信息
	GetOtherInfo(context.Context, *GetOtherInfoRep) (*GetOtherInfoReply, error)
	// GetSysUser根据用户ID获取详细信息
	GetSysUser(context.Context, *GetSysUserRep) (*GetSysUserReply, error)
	ListSysUser(context.Context, *ListSysUserRep) (*ListSysUserReply, error)
	Profile(context.Context, *ProfileRep) (*ProfileReply, error)
	ResetPwd(context.Context, *ResetPwdRep) (*ResetPwdReply, error)
	SaveSysUser(context.Context, *SaveSysUserRep) (*SaveSysUserReply, error)
	UpdateSysUser(context.Context, *UpdateSysUserRep) (*UpdateSysUserReply, error)
}

func RegisterSysUserHTTPServer(s *http.Server, srv SysUserHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/user/create", _SysUser_CreateSysUser0_HTTP_Handler(srv))
	r.PUT("/v1/user/update", _SysUser_UpdateSysUser0_HTTP_Handler(srv))
	r.PUT("/v1/user/resetPwd", _SysUser_ResetPwd0_HTTP_Handler(srv))
	r.DELETE("/v1/user/delete/{id}", _SysUser_DeleteSysUser0_HTTP_Handler(srv))
	r.POST("/v1/user/save", _SysUser_SaveSysUser0_HTTP_Handler(srv))
	r.GET("/v1/user/info/{id}", _SysUser_GetSysUser0_HTTP_Handler(srv))
	r.GET("/v1/user/otherInfo", _SysUser_GetOtherInfo0_HTTP_Handler(srv))
	r.POST("/v1/user/list", _SysUser_ListSysUser0_HTTP_Handler(srv))
	r.GET("/v1/user/profile", _SysUser_Profile0_HTTP_Handler(srv))
	r.GET("/v1/user/authRole/{id}", _SysUser_GetAuthRoleSysUser0_HTTP_Handler(srv))
	r.PUT("/v1/user/authRole", _SysUser_AuthRoleSysUser0_HTTP_Handler(srv))
}

func _SysUser_CreateSysUser0_HTTP_Handler(srv SysUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateSysUserRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysUserCreateSysUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.CreateSysUser(ctx, req.(*CreateSysUserRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateSysUserReply)
		return ctx.Result(200, reply)
	}
}

func _SysUser_UpdateSysUser0_HTTP_Handler(srv SysUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateSysUserRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysUserUpdateSysUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateSysUser(ctx, req.(*UpdateSysUserRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateSysUserReply)
		return ctx.Result(200, reply)
	}
}

func _SysUser_ResetPwd0_HTTP_Handler(srv SysUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ResetPwdRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysUserResetPwd)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ResetPwd(ctx, req.(*ResetPwdRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ResetPwdReply)
		return ctx.Result(200, reply)
	}
}

func _SysUser_DeleteSysUser0_HTTP_Handler(srv SysUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteSysUserRep
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysUserDeleteSysUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.DeleteSysUser(ctx, req.(*DeleteSysUserRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteSysUserReply)
		return ctx.Result(200, reply)
	}
}

func _SysUser_SaveSysUser0_HTTP_Handler(srv SysUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in SaveSysUserRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysUserSaveSysUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SaveSysUser(ctx, req.(*SaveSysUserRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*SaveSysUserReply)
		return ctx.Result(200, reply)
	}
}

func _SysUser_GetSysUser0_HTTP_Handler(srv SysUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetSysUserRep
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysUserGetSysUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetSysUser(ctx, req.(*GetSysUserRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetSysUserReply)
		return ctx.Result(200, reply)
	}
}

func _SysUser_GetOtherInfo0_HTTP_Handler(srv SysUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetOtherInfoRep
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysUserGetOtherInfo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetOtherInfo(ctx, req.(*GetOtherInfoRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetOtherInfoReply)
		return ctx.Result(200, reply)
	}
}

func _SysUser_ListSysUser0_HTTP_Handler(srv SysUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListSysUserRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysUserListSysUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListSysUser(ctx, req.(*ListSysUserRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListSysUserReply)
		return ctx.Result(200, reply)
	}
}

func _SysUser_Profile0_HTTP_Handler(srv SysUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ProfileRep
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysUserProfile)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Profile(ctx, req.(*ProfileRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ProfileReply)
		return ctx.Result(200, reply)
	}
}

func _SysUser_GetAuthRoleSysUser0_HTTP_Handler(srv SysUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetAuthRoleSysUserRep
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysUserGetAuthRoleSysUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetAuthRoleSysUser(ctx, req.(*GetAuthRoleSysUserRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetAuthRoleSysUserReply)
		return ctx.Result(200, reply)
	}
}

func _SysUser_AuthRoleSysUser0_HTTP_Handler(srv SysUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuthRoleSysUserRep
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationSysUserAuthRoleSysUser)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.AuthRoleSysUser(ctx, req.(*AuthRoleSysUserRep))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthRoleSysUserReply)
		return ctx.Result(200, reply)
	}
}

type SysUserHTTPClient interface {
	AuthRoleSysUser(ctx context.Context, req *AuthRoleSysUserRep, opts ...http.CallOption) (rsp *AuthRoleSysUserReply, err error)
	CreateSysUser(ctx context.Context, req *CreateSysUserRep, opts ...http.CallOption) (rsp *CreateSysUserReply, err error)
	DeleteSysUser(ctx context.Context, req *DeleteSysUserRep, opts ...http.CallOption) (rsp *DeleteSysUserReply, err error)
	GetAuthRoleSysUser(ctx context.Context, req *GetAuthRoleSysUserRep, opts ...http.CallOption) (rsp *GetAuthRoleSysUserReply, err error)
	GetOtherInfo(ctx context.Context, req *GetOtherInfoRep, opts ...http.CallOption) (rsp *GetOtherInfoReply, err error)
	GetSysUser(ctx context.Context, req *GetSysUserRep, opts ...http.CallOption) (rsp *GetSysUserReply, err error)
	ListSysUser(ctx context.Context, req *ListSysUserRep, opts ...http.CallOption) (rsp *ListSysUserReply, err error)
	Profile(ctx context.Context, req *ProfileRep, opts ...http.CallOption) (rsp *ProfileReply, err error)
	ResetPwd(ctx context.Context, req *ResetPwdRep, opts ...http.CallOption) (rsp *ResetPwdReply, err error)
	SaveSysUser(ctx context.Context, req *SaveSysUserRep, opts ...http.CallOption) (rsp *SaveSysUserReply, err error)
	UpdateSysUser(ctx context.Context, req *UpdateSysUserRep, opts ...http.CallOption) (rsp *UpdateSysUserReply, err error)
}

type SysUserHTTPClientImpl struct {
	cc *http.Client
}

func NewSysUserHTTPClient(client *http.Client) SysUserHTTPClient {
	return &SysUserHTTPClientImpl{client}
}

func (c *SysUserHTTPClientImpl) AuthRoleSysUser(ctx context.Context, in *AuthRoleSysUserRep, opts ...http.CallOption) (*AuthRoleSysUserReply, error) {
	var out AuthRoleSysUserReply
	pattern := "/v1/user/authRole"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysUserAuthRoleSysUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysUserHTTPClientImpl) CreateSysUser(ctx context.Context, in *CreateSysUserRep, opts ...http.CallOption) (*CreateSysUserReply, error) {
	var out CreateSysUserReply
	pattern := "/v1/user/create"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysUserCreateSysUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysUserHTTPClientImpl) DeleteSysUser(ctx context.Context, in *DeleteSysUserRep, opts ...http.CallOption) (*DeleteSysUserReply, error) {
	var out DeleteSysUserReply
	pattern := "/v1/user/delete/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSysUserDeleteSysUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysUserHTTPClientImpl) GetAuthRoleSysUser(ctx context.Context, in *GetAuthRoleSysUserRep, opts ...http.CallOption) (*GetAuthRoleSysUserReply, error) {
	var out GetAuthRoleSysUserReply
	pattern := "/v1/user/authRole/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSysUserGetAuthRoleSysUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysUserHTTPClientImpl) GetOtherInfo(ctx context.Context, in *GetOtherInfoRep, opts ...http.CallOption) (*GetOtherInfoReply, error) {
	var out GetOtherInfoReply
	pattern := "/v1/user/otherInfo"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSysUserGetOtherInfo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysUserHTTPClientImpl) GetSysUser(ctx context.Context, in *GetSysUserRep, opts ...http.CallOption) (*GetSysUserReply, error) {
	var out GetSysUserReply
	pattern := "/v1/user/info/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSysUserGetSysUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysUserHTTPClientImpl) ListSysUser(ctx context.Context, in *ListSysUserRep, opts ...http.CallOption) (*ListSysUserReply, error) {
	var out ListSysUserReply
	pattern := "/v1/user/list"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysUserListSysUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysUserHTTPClientImpl) Profile(ctx context.Context, in *ProfileRep, opts ...http.CallOption) (*ProfileReply, error) {
	var out ProfileReply
	pattern := "/v1/user/profile"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationSysUserProfile))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysUserHTTPClientImpl) ResetPwd(ctx context.Context, in *ResetPwdRep, opts ...http.CallOption) (*ResetPwdReply, error) {
	var out ResetPwdReply
	pattern := "/v1/user/resetPwd"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysUserResetPwd))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysUserHTTPClientImpl) SaveSysUser(ctx context.Context, in *SaveSysUserRep, opts ...http.CallOption) (*SaveSysUserReply, error) {
	var out SaveSysUserReply
	pattern := "/v1/user/save"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysUserSaveSysUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *SysUserHTTPClientImpl) UpdateSysUser(ctx context.Context, in *UpdateSysUserRep, opts ...http.CallOption) (*UpdateSysUserReply, error) {
	var out UpdateSysUserReply
	pattern := "/v1/user/update"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationSysUserUpdateSysUser))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
