package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"micro/internal/biz"

	v1 "micro/api/api/system/v1"
)

type SysRoleService struct {
	v1.UnimplementedSysRoleServer
	uc  *biz.SysRoleUsecase
	log *log.Helper
}

func NewSysRoleService(uc *biz.SysRoleUsecase, logger log.Logger) *SysRoleService {
	return &SysRoleService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysRoleService) CreateSysRole(ctx context.Context, req *v1.ModifySysRoleRep) (*v1.EmptyReply, error) {
	return s.uc.CreateSysRole(ctx, req)
}
func (s *SysRoleService) UpdateSysRole(ctx context.Context, req *v1.ModifySysRoleRep) (*v1.EmptyReply, error) {
	return s.uc.UpdateSysRole(ctx, req)
}
func (s *SysRoleService) DeleteSysRole(ctx context.Context, req *v1.DeleteSysRoleRep) (*v1.EmptyReply, error) {
	return s.uc.Delete(ctx, req.Id)
}
func (s *SysRoleService) DataScopeSysRole(ctx context.Context, req *v1.DataScopeSysRoleRep) (*v1.EmptyReply, error) {
	return s.uc.DataScope(ctx, req)
}
func (s *SysRoleService) ChangeStatusSysRole(ctx context.Context, req *v1.ChangeStatusSysRoleRep) (*v1.EmptyReply, error) {
	return s.uc.ChangeStatus(ctx, req)
}
func (s *SysRoleService) GetSysRole(ctx context.Context, req *v1.GetSysRoleRep) (*v1.GetSysRoleReply, error) {
	return s.uc.GetRoleById(ctx, req)
}
func (s *SysRoleService) ListSysRole(ctx context.Context, req *v1.ListSysRoleRep) (*v1.ListSysRoleReply, error) {
	return s.uc.GetPageSet(ctx, req)
}
func (s *SysRoleService) AllocatedList(ctx context.Context, req *v1.IsAllocatedListRep) (*v1.IsAllocatedListReply, error) {
	return s.uc.IsAllocatedList(ctx, req, true)
}
func (s *SysRoleService) UnAllocatedList(ctx context.Context, req *v1.IsAllocatedListRep) (*v1.IsAllocatedListReply, error) {
	return s.uc.IsAllocatedList(ctx, req, false)
}
func (s *SysRoleService) AuthUserCancel(ctx context.Context, req *v1.AuthUserCancelRep) (*v1.EmptyReply, error) {
	return s.uc.AuthUserCancel(ctx, req)
}
func (s *SysRoleService) AuthUserCancelAll(ctx context.Context, req *v1.AuthUserSelectAllRep) (*v1.EmptyReply, error) {
	return s.uc.AuthUserCancelAll(ctx, req)
}
func (s *SysRoleService) AuthUserSelectAll(ctx context.Context, req *v1.AuthUserSelectAllRep) (*v1.EmptyReply, error) {
	return s.uc.AuthUserSelectAll(ctx, req)
}
