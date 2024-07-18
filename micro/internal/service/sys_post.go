package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	pb "micro/api/system/v1"
	"micro/internal/biz"
)

type SysPostService struct {
	pb.UnimplementedSysPostServer
	uc  *biz.SysPostUsecase
	log *log.Helper
}

func NewSysPostService(uc *biz.SysPostUsecase, logger log.Logger) *SysPostService {
	return &SysPostService{uc: uc, log: log.NewHelper(logger)}
}

func (s *SysPostService) CreateSysPost(ctx context.Context, req *pb.SysPostRep) (*pb.EmptyReply, error) {
	return s.uc.CreateSysPost(ctx, req)
}
func (s *SysPostService) UpdateSysPost(ctx context.Context, req *pb.SysPostRep) (*pb.EmptyReply, error) {
	return s.uc.UpdateSysPost(ctx, req)
}
func (s *SysPostService) DeleteSysPost(ctx context.Context, req *pb.DeleteSysPostRep) (*pb.EmptyReply, error) {
	return s.uc.DeleteSysPost(ctx, req)
}
func (s *SysPostService) GetSysPost(ctx context.Context, req *pb.GetSysPostRep) (*pb.GetSysPostReply, error) {
	return s.uc.GetSysPost(ctx, req)
}
func (s *SysPostService) ListSysPost(ctx context.Context, req *pb.ListSysPostRep) (*pb.ListSysPostReply, error) {
	return s.uc.GetPageSet(ctx, req)
}
