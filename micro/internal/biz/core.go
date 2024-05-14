package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type CoreRepo interface {
	RSetData(ctx context.Context, key string, value interface{}) error
	RSetObj(ctx context.Context, key string, value interface{}) error
	RGetData(ctx context.Context, key string) (interface{}, error)
	RGetObj(ctx context.Context, key string) ([]byte, error)
	RRemoveData(ctx context.Context, key string) error
	RExistsData(ctx context.Context, key string) (int64, error)
}

type CoreUsecase struct {
	repo CoreRepo
	log  *log.Helper
}

func NewCoreUsecase(repo CoreRepo, logger log.Logger) *CoreUsecase {
	return &CoreUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *CoreUsecase) RedisSetData(ctx context.Context, key string, value interface{}) error {
	return uc.repo.RSetData(ctx, key, value)
}

func (uc *CoreUsecase) RedisGetData(ctx context.Context, key string) (interface{}, error) {
	return uc.repo.RGetData(ctx, key)
}

func (uc *CoreUsecase) RedisRemoveData(ctx context.Context, key string) error {
	return uc.repo.RRemoveData(ctx, key)
}

func (uc *CoreUsecase) RedisExistsData(ctx context.Context, key string) (int64, error) {
	return uc.repo.RExistsData(ctx, key)
}
