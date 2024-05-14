package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"micro/internal/biz"

	pb "micro/api/api/system/v1"
)

type SysConfigService struct {
	pb.UnimplementedSysConfigServer
	uc  *biz.SysConfigUsecase
	log *log.Helper
}

func NewSysConfigService(uc *biz.SysConfigUsecase, logger log.Logger) *SysConfigService {
	return &SysConfigService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysConfigService) CreateSysConfig(ctx context.Context, req *pb.CreateSysConfigRep) (*pb.EmptySysConfigReply, error) {
	return &pb.EmptySysConfigReply{}, nil
}
func (s *SysConfigService) UpdateSysConfig(ctx context.Context, req *pb.UpdateSysConfigRep) (*pb.EmptySysConfigReply, error) {
	return &pb.EmptySysConfigReply{}, nil
}
func (s *SysConfigService) DeleteSysConfig(ctx context.Context, req *pb.DeleteSysConfigRep) (*pb.EmptySysConfigReply, error) {
	return &pb.EmptySysConfigReply{}, nil
}
func (s *SysConfigService) GetSysConfig(ctx context.Context, req *pb.GetSysConfigRep) (*pb.GetSysConfigReply, error) {
	return &pb.GetSysConfigReply{}, nil
}
func (s *SysConfigService) ListSysConfig(ctx context.Context, req *pb.ListSysConfigRep) (*pb.ListSysConfigReply, error) {
	return s.uc.ListSysConfig(ctx, req)
}
func (s *SysConfigService) ConfigByKey(ctx context.Context, req *pb.ConfigByKeyReq) (*pb.ConfigByKeyReply, error) {
	return s.uc.GetConfigByKey(ctx, req.Key)
}
