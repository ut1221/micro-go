package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "micro/api/api/system/v1"
	"micro/internal/biz"
)

type SysDictService struct {
	v1.UnimplementedSysDictServer
	uc  *biz.SysDictUsecase
	log *log.Helper
}

func NewSysDictService(uc *biz.SysDictUsecase, logger log.Logger) *SysDictService {
	return &SysDictService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysDictService) SysDictDataType(ctx context.Context, req *v1.SysDictDataTypeReq) (*v1.SysDictDataTypeReply, error) {
	data, err := s.uc.GetSysDictDataByType(ctx, req.Type)
	if err != nil {
		return nil, err
	}
	return &v1.SysDictDataTypeReply{Dict: data}, nil
}
func (s *SysDictService) CreateSysDictType(ctx context.Context, req *v1.SysDictTypeRep) (*v1.EmptySysDictReply, error) {
	return s.uc.CreateSysDictType(ctx, req)
}
func (s *SysDictService) UpdateSysDictType(ctx context.Context, req *v1.SysDictTypeRep) (*v1.EmptySysDictReply, error) {
	return s.uc.UpdateSysDictType(ctx, req)
}
func (s *SysDictService) DeleteSysDictType(ctx context.Context, req *v1.DeleteSysDictTypeRep) (*v1.EmptySysDictReply, error) {
	return s.uc.DeleteSysDictType(ctx, req)
}
func (s *SysDictService) GetSysDictType(ctx context.Context, req *v1.GetSysDictRep) (*v1.GetSysDictTypeReply, error) {
	return s.uc.GetSysDictType(ctx, req)
}
func (s *SysDictService) ListSysDictType(ctx context.Context, req *v1.ListSysDictTypeRep) (*v1.ListSysDictTypeReply, error) {
	return s.uc.ListSysDictType(ctx, req)
}
func (s *SysDictService) OptionSelectType(ctx context.Context, req *v1.OptionSelectTypeRep) (*v1.OptionSelectTypeReply, error) {
	return s.uc.OptionSelectType(ctx, req)
}
func (s *SysDictService) CreateSysDictData(ctx context.Context, req *v1.SysDictDataRep) (*v1.EmptySysDictReply, error) {
	return s.uc.CreateSysDictData(ctx, req)
}
func (s *SysDictService) UpdateSysDictData(ctx context.Context, req *v1.SysDictDataRep) (*v1.EmptySysDictReply, error) {
	return s.uc.UpdateSysDictData(ctx, req)
}
func (s *SysDictService) DeleteSysDictData(ctx context.Context, req *v1.DeleteSysDictDataRep) (*v1.EmptySysDictReply, error) {
	return s.uc.DeleteSysDictData(ctx, req)
}
func (s *SysDictService) GetSysDictData(ctx context.Context, req *v1.GetSysDictRep) (*v1.GetSysDictDataReply, error) {
	return s.uc.GetSysDictData(ctx, req)
}
func (s *SysDictService) ListSysDictData(ctx context.Context, req *v1.ListSysDictDataRep) (*v1.ListSysDictDataReply, error) {
	return s.uc.ListSysDictData(ctx, req)
}
func (s *SysDictService) RefreshCacheSysDict(ctx context.Context, req *v1.CacheSysDicReq) (*v1.EmptySysDictReply, error) {

	return s.uc.RefreshCacheSysDict(ctx, req)
}
