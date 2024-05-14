package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "micro/api/api/system/v1"
	"micro/internal/model"
)

type SysConfigRepo interface {
	GetConfig(ctx context.Context, query *model.SysConfigQuery) (*model.BizSysConfig, error)
	Save(ctx context.Context, config *model.SysConfig) error
	Update(ctx context.Context, config *model.SysConfig) error
	Delete(ctx context.Context, configId string) error
	GetPageSet(ctx context.Context, num int64, size int64, query model.SysConfigQuery) ([]*model.BizSysConfig, int64, error)
}

type SysConfigUsecase struct {
	repo SysConfigRepo
	log  *log.Helper
}

func NewSysConfigUsecase(repo SysConfigRepo, logger log.Logger) *SysConfigUsecase {
	return &SysConfigUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysConfigUsecase) GetConfigByKey(ctx context.Context, key string) (*v1.ConfigByKeyReply, error) {
	config, err := uc.repo.GetConfig(ctx, &model.SysConfigQuery{ConfigKey: key})
	if err != nil {
		return nil, err
	}
	return &v1.ConfigByKeyReply{Value: config.ConfigValue}, err
}

func (uc *SysConfigUsecase) ListSysConfig(ctx context.Context, req *v1.ListSysConfigRep) (*v1.ListSysConfigReply, error) {
	bizConfigs, total, err := uc.repo.GetPageSet(ctx, req.GetPageNum(), req.GetPageSize(), model.SysConfigQuery{})
	if err != nil {
		return nil, err
	}
	list := make([]*v1.ConfigReply, len(bizConfigs))
	for i, config := range bizConfigs {
		list[i] = config.BizToV1()
	}
	return &v1.ListSysConfigReply{Rows: list, Total: total}, err
}
