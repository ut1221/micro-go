package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "micro/api/system/v1"
	"micro/internal/biz"
	"micro/internal/pkg"
)

type SysDeptService struct {
	v1.UnimplementedSysDeptServer
	uc  *biz.SysDeptUsecase
	log *log.Helper
}

func NewSysDeptService(uc *biz.SysDeptUsecase, logger log.Logger) *SysDeptService {
	return &SysDeptService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysDeptService) DeptTree(ctx context.Context, req *v1.DeptTreeReq) (*v1.DeptTreeReply, error) {
	tree, err := s.uc.GetDeptTree(ctx)
	if err != nil {
		return nil, err
	}
	return &v1.DeptTreeReply{
		Dept: tree,
	}, nil
}

func (s *SysDeptService) CreateSysDept(ctx context.Context, req *v1.SysDeptRep) (*v1.EmptyReply, error) {
	req.DeptId = pkg.GetID()
	return s.uc.Sava(ctx, req)
}
func (s *SysDeptService) UpdateSysDept(ctx context.Context, req *v1.SysDeptRep) (*v1.EmptyReply, error) {
	return s.uc.Update(ctx, req)
}
func (s *SysDeptService) DeleteSysDept(ctx context.Context, req *v1.DeleteSysDeptRep) (*v1.EmptyReply, error) {
	return s.uc.Delete(ctx, req)
}
func (s *SysDeptService) GetSysDept(ctx context.Context, req *v1.GetSysDeptRep) (*v1.GetSysDeptReply, error) {
	return s.uc.GetSysDeptData(ctx, req)
}
func (s *SysDeptService) ExcludeDept(ctx context.Context, req *v1.ExcludeDeptRep) (*v1.ListSysDeptReply, error) {
	return s.uc.ExcludeDept(ctx, req)
}
func (s *SysDeptService) ListSysDept(ctx context.Context, req *v1.ListSysDeptRep) (*v1.ListSysDeptReply, error) {
	return s.uc.GetDeptList(ctx, req)
}
func (s *SysDeptService) GetSysRoleDept(ctx context.Context, req *v1.GetSysRoleDeptRep) (*v1.GetSysRoleDeptReply, error) {
	return s.uc.GetSysRoleDept(ctx, req)
}
