package server

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"
	"golang.org/x/net/context"
	v1 "micro/api/system/v1"
	"micro/internal/conf"
	"micro/internal/pkg"
	"micro/internal/pkg/middleware"
	"micro/internal/service"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, logger log.Logger,
	as *service.AuthService,
	sysUser *service.SysUserService,
	sysRole *service.SysRoleService,
	sysMenu *service.SysMenuService,
	sysDept *service.SysDeptService,
	sysDict *service.SysDictService,
	sysConfig *service.SysConfigService,
	sysPost *service.SysPostService,
	monitor *service.MonitorService,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			validate.Validator(),
			recovery.Recovery(),
			selector.Server(middleware.Auth()).Match(NewWhiteListMatcher()).Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	opts = append(opts, http.ResponseEncoder(pkg.EncoderResponse()))
	opts = append(opts, http.ErrorEncoder(pkg.EncoderError()))
	srv := http.NewServer(opts...)
	v1.RegisterAuthHTTPServer(srv, as)
	v1.RegisterSysUserHTTPServer(srv, sysUser)
	v1.RegisterSysRoleHTTPServer(srv, sysRole)
	v1.RegisterSysMenuHTTPServer(srv, sysMenu)
	v1.RegisterSysDeptHTTPServer(srv, sysDept)
	v1.RegisterSysDictHTTPServer(srv, sysDict)
	v1.RegisterSysConfigHTTPServer(srv, sysConfig)
	v1.RegisterSysPostHTTPServer(srv, sysPost)
	v1.RegisterMonitorHTTPServer(srv, monitor)
	return srv
}

// NewWhiteListMatcher 设置白名单，不需要 token 验证的接口
func NewWhiteListMatcher() selector.MatchFunc {
	whiteList := make(map[string]struct{})
	whiteList["/api.system.v1.Auth/Login"] = struct{}{}
	whiteList["/api.system.v1.Auth/Captcha"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
