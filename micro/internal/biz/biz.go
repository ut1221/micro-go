package biz

import (
	"github.com/google/wire"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewCoreUsecase,
	NewSysUserUsecase,
	NewSysRoleUsecase,
	NewSysMenuUsecase,
	NewSysDeptUsecase,
	NewSysDictUsecase,
	NewSysConfigUsecase,
	NewSysPostUsecase,
)

//type Transaction interface {
//	InTx(context.Context, func(ctx context.Context) error) error
//}
