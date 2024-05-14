package service

import (
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAuthService,
	NewSysUserService,
	NewSysRoleService,
	NewSysMenuService,
	NewSysDeptService,
	NewSysDictService,
	NewSysConfigService,
	NewSysPostService,
	NewMonitorService,
)
