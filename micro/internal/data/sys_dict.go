package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"micro/internal/biz"
	"micro/internal/model"
)

type sysDictRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysDictRepo(data *Data, logger log.Logger) biz.SysDictRepo {
	return &sysDictRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s sysDictRepo) GetSysDictDataByType(ctx context.Context, dictType string) ([]*model.BizSysDictData, error) {
	var sysDictType *model.SysDictType
	if err := s.data.Db.Preload("SysDictData").Where("dict_type = ?", dictType).First(&sysDictType).Error; err != nil {
		return nil, err
	}
	var bizDictData []*model.BizSysDictData
	for _, sysDictData := range sysDictType.SysDictData {
		toBiz := sysDictData.EntityToBiz()
		bizDictData = append(bizDictData, toBiz)
	}
	return bizDictData, nil
}

func (s sysDictRepo) SaveType(ctx context.Context, dictType *model.SysDictType) error {
	if err := s.data.Db.Save(&dictType).Error; err != nil {
		return err
	}
	return nil
}

func (s sysDictRepo) UpdateType(ctx context.Context, dictType *model.SysDictType) error {
	if err := s.data.Db.Updates(&dictType).Error; err != nil {
		return err
	}
	return nil
}

func (s sysDictRepo) DeleteType(ctx context.Context, dictId string) error {
	if err := s.data.Db.Where("dict_id = ?", dictId).Delete(&model.SysDictType{}).Error; err != nil {
		return err
	}
	return nil
}

func (s sysDictRepo) GetTypeInfo(ctx context.Context, query *model.SysDictTypeQuery) (*model.BizSysDictType, error) {
	var dictType model.SysDictType
	db := s.data.Db.Model(&model.SysDictType{})
	if query.DictID != "" {
		db.Where("dict_id = ?", query.DictID)
	}
	if query.DictType != "" {
		db.Where("dict_type = ?", query.DictType)
	}
	if err := db.First(&dictType).Error; err != nil {
		return nil, err
	}
	return dictType.EntityToBiz(), nil
}

func (s sysDictRepo) GetTypePageSet(ctx context.Context, pageNum int64, pageSize int64, query *model.SysDictTypeQuery) ([]*model.BizSysDictType, int64, error) {
	var dictTypes []*model.SysDictType
	var total int64
	db := s.data.Db.Model(&model.SysDictType{})
	if query.DictName != "" {
		db.Where("dict_name = ?", query.DictName)
	}
	tx := db.Count(&total).Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&dictTypes)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	var bizDictType []*model.BizSysDictType
	for _, sysDictType := range dictTypes {
		toBiz := sysDictType.EntityToBiz()
		bizDictType = append(bizDictType, toBiz)
	}
	return bizDictType, total, nil
}

func (s sysDictRepo) SaveData(ctx context.Context, dictData *model.SysDictData) error {
	if err := s.data.Db.Save(&dictData).Error; err != nil {
		return err
	}
	return nil
}

func (s sysDictRepo) UpdateData(ctx context.Context, dictData *model.SysDictData) error {
	if err := s.data.Db.Updates(&dictData).Error; err != nil {
		return err
	}
	return nil
}

func (s sysDictRepo) DeleteData(ctx context.Context, dictCode string) error {
	if err := s.data.Db.Where("dict_code = ?", dictCode).Delete(&model.SysDictData{}).Error; err != nil {
		return err
	}
	return nil
}

func (s sysDictRepo) DeleteDataByType(ctx context.Context, dictType string) error {
	return s.data.Db.Where("dict_type = ?", dictType).Delete(&model.SysDictData{}).Error
}

func (s sysDictRepo) GetDataInfo(ctx context.Context, query *model.SysDictDataQuery) (*model.BizSysDictData, error) {
	var dictData model.SysDictData
	db := s.data.Db.Model(&model.SysDictData{})
	if query.DictCode != "" {
		db.Where("dict_code = ?", query.DictCode)
	}
	if query.DictLabel != "" {
		db.Where("dict_label = ?", query.DictLabel)
	}
	if err := db.First(&dictData).Error; err != nil {
		return nil, err
	}
	return dictData.EntityToBiz(), nil
}

func (s sysDictRepo) GetDataPageSet(ctx context.Context, pageNum int64, pageSize int64, query *model.SysDictDataQuery) ([]*model.BizSysDictData, int64, error) {
	var dictDatas []*model.SysDictData
	var total int64
	db := s.data.Db.Model(&model.SysDictData{})
	if query.DictType != "" {
		db.Where("dict_type = ?", query.DictType)
	}
	tx := db.Count(&total).Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&dictDatas)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	var bizDictData []*model.BizSysDictData
	for _, sysDictData := range dictDatas {
		toBiz := sysDictData.EntityToBiz()
		bizDictData = append(bizDictData, toBiz)
	}
	return bizDictData, total, nil
}

func (s sysDictRepo) OptionSelectType(ctx context.Context) ([]*model.BizSysDictType, error) {
	var sysDictTypes []*model.SysDictType
	if err := s.data.Db.Model(&model.SysDictType{}).Find(&sysDictTypes).Error; err != nil {
		return nil, err
	}
	var bizDictType []*model.BizSysDictType
	for _, sysDictType := range sysDictTypes {
		toBiz := sysDictType.EntityToBiz()
		bizDictType = append(bizDictType, toBiz)
	}
	return bizDictType, nil
}
