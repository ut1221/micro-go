package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"micro/internal/biz"
	"micro/internal/model"
)

type sysConfigRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysConfigRepo(data *Data, logger log.Logger) biz.SysConfigRepo {
	return &sysConfigRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s sysConfigRepo) GetConfig(ctx context.Context, query *model.SysConfigQuery) (*model.BizSysConfig, error) {
	var sysConfig model.SysConfig
	db := s.data.Db.Model(&model.SysConfig{})
	if query.ConfigId != "" {
		db.Where("config_id", query.ConfigId)
	}
	if query.ConfigKey != "" {
		db.Where("config_key", query.ConfigKey)
	}
	if err := db.First(&sysConfig).Error; err != nil {
		return nil, err
	}
	toBiz := sysConfig.EntityToBiz()
	return toBiz, nil
}

func (s sysConfigRepo) Save(ctx context.Context, config *model.SysConfig) error {
	return s.data.Db.Save(&config).Error
}

func (s sysConfigRepo) Update(ctx context.Context, config *model.SysConfig) error {
	return s.data.Db.Updates(&config).Error
}

func (s sysConfigRepo) Delete(ctx context.Context, configId string) error {
	return s.data.Db.Where("config_id", configId).Delete(&model.SysConfig{}).Error
}

func (s sysConfigRepo) GetPageSet(ctx context.Context, pageNum int64, pageSize int64, query model.SysConfigQuery) ([]*model.BizSysConfig, int64, error) {
	var total int64
	var sysConfigs []*model.SysConfig
	db := s.data.Db.Model(&model.SysConfig{}).Count(&total)
	tx := db.Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&sysConfigs)

	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	toBiz := make([]*model.BizSysConfig, len(sysConfigs))
	for i, v := range sysConfigs {
		toBiz[i] = v.EntityToBiz()
	}
	return toBiz, total, nil
}
