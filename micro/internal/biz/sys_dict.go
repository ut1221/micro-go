package biz

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	v1 "micro/api/system/v1"
	"micro/internal/model"
	"micro/internal/pkg"
	"micro/internal/pkg/constants"
)

type SysDictRepo interface {
	GetSysDictDataByType(ctx context.Context, dictType string) ([]*model.BizSysDictData, error)
	SaveType(ctx context.Context, dictType *model.SysDictType) error
	UpdateType(ctx context.Context, dictType *model.SysDictType) error
	DeleteType(ctx context.Context, dictId string) error
	GetTypeInfo(ctx context.Context, query *model.SysDictTypeQuery) (*model.BizSysDictType, error)
	GetTypePageSet(ctx context.Context, pageNum int64, pageSize int64, query *model.SysDictTypeQuery) ([]*model.BizSysDictType, int64, error)

	SaveData(ctx context.Context, dictData *model.SysDictData) error
	UpdateData(ctx context.Context, dictData *model.SysDictData) error
	DeleteData(ctx context.Context, dictCode string) error
	DeleteDataByType(ctx context.Context, dictType string) error
	GetDataInfo(ctx context.Context, query *model.SysDictDataQuery) (*model.BizSysDictData, error)
	GetDataPageSet(ctx context.Context, pageNum int64, pageSize int64, query *model.SysDictDataQuery) ([]*model.BizSysDictData, int64, error)
	OptionSelectType(ctx context.Context) ([]*model.BizSysDictType, error)
}

type SysDictUsecase struct {
	repo     SysDictRepo
	coreRepo CoreRepo
	log      *log.Helper
}

func NewSysDictUsecase(repo SysDictRepo, coreRepo CoreRepo, logger log.Logger) *SysDictUsecase {
	return &SysDictUsecase{repo: repo, coreRepo: coreRepo, log: log.NewHelper(logger)}
}

func (uc *SysDictUsecase) GetSysDictDataByType(ctx context.Context, dictType string) ([]*v1.DictDataReply, error) {
	//查询redis中是否存在
	var dataDict []*model.BizSysDictData
	res, _ := uc.coreRepo.RExistsData(ctx, constants.CacheSysDict+dictType)
	if res == 0 {
		dataDict, _ = uc.repo.GetSysDictDataByType(ctx, dictType)
	} else {
		data, err := uc.coreRepo.RGetObj(ctx, constants.CacheSysDict+dictType)
		if err != nil {
			return nil, err
		}
		if err = json.Unmarshal(data, &dataDict); err != nil {
			return nil, err
		}
	}
	var dictData []*v1.DictDataReply
	for _, bizData := range dataDict {
		toV1 := bizData.BizToV1()
		dictData = append(dictData, toV1)
	}
	return dictData, nil
}

func (uc *SysDictUsecase) ListSysDictType(ctx context.Context, req *v1.ListSysDictTypeRep) (*v1.ListSysDictTypeReply, error) {
	info, total, err := uc.repo.GetTypePageSet(ctx, req.GetPageNum(), req.GetPageSize(), &model.SysDictTypeQuery{})
	if err != nil {
		return nil, err
	}
	var rows []*v1.DictTypeReply
	for _, dictType := range info {
		toV1 := dictType.BizToV1()
		rows = append(rows, toV1)
	}
	return &v1.ListSysDictTypeReply{Rows: rows, Total: total}, err
}

func (uc *SysDictUsecase) CreateSysDictType(ctx context.Context, req *v1.SysDictTypeRep) (*v1.EmptySysDictReply, error) {
	if _, err := uc.repo.GetTypeInfo(ctx, &model.SysDictTypeQuery{DictType: req.GetDictType()}); !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("该编码已存在")
	}
	err := uc.repo.SaveType(ctx, &model.SysDictType{
		DictID:   pkg.GetID(),
		DictName: req.GetDictName(),
		DictType: req.GetDictType(),
		Remark:   req.GetRemark(),
		Status:   req.GetStatus(),
	})
	return &v1.EmptySysDictReply{}, err
}

func (uc *SysDictUsecase) UpdateSysDictType(ctx context.Context, req *v1.SysDictTypeRep) (*v1.EmptySysDictReply, error) {
	err := uc.repo.UpdateType(ctx, &model.SysDictType{
		DictID:   req.DictId,
		DictName: req.GetDictName(),
		DictType: req.GetDictType(),
		Remark:   req.GetRemark(),
		Status:   req.GetStatus(),
	})
	return &v1.EmptySysDictReply{}, err
}

func (uc *SysDictUsecase) DeleteSysDictType(ctx context.Context, req *v1.DeleteSysDictTypeRep) (*v1.EmptySysDictReply, error) {
	typeInfo, _ := uc.repo.GetTypeInfo(ctx, &model.SysDictTypeQuery{DictID: req.GetId()})
	if err := uc.repo.DeleteType(ctx, req.GetId()); err != nil {
		return nil, err
	}
	if err := uc.repo.DeleteDataByType(ctx, typeInfo.DictType); err != nil {
		return nil, err
	}
	return &v1.EmptySysDictReply{}, nil
}

func (uc *SysDictUsecase) GetSysDictType(ctx context.Context, req *v1.GetSysDictRep) (*v1.GetSysDictTypeReply, error) {
	r, err := uc.repo.GetTypeInfo(ctx, &model.SysDictTypeQuery{DictID: req.GetId()})
	if err != nil {
		return nil, err
	}
	return &v1.GetSysDictTypeReply{
		Dict: r.BizToV1(),
	}, nil
}

func (uc *SysDictUsecase) CreateSysDictData(ctx context.Context, req *v1.SysDictDataRep) (*v1.EmptySysDictReply, error) {
	err := uc.repo.SaveData(ctx, &model.SysDictData{
		DictCode:  pkg.GetID(),
		DictSort:  req.GetDictSort(),
		DictLabel: req.GetDictLabel(),
		DictValue: req.GetDictValue(),
		DictType:  req.GetDictType(),
		CSSClass:  req.GetCssClass(),
		ListClass: req.GetListClass(),
		IsDefault: req.GetIsDefault(),
		Status:    req.GetStatus(),
		Remark:    req.GetRemark(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.EmptySysDictReply{}, nil
}

func (uc *SysDictUsecase) UpdateSysDictData(ctx context.Context, req *v1.SysDictDataRep) (*v1.EmptySysDictReply, error) {
	err := uc.repo.UpdateData(ctx, &model.SysDictData{
		DictCode:  req.DictCode,
		DictSort:  req.GetDictSort(),
		DictLabel: req.GetDictLabel(),
		DictValue: req.GetDictValue(),
		DictType:  req.GetDictType(),
		CSSClass:  req.GetCssClass(),
		ListClass: req.GetListClass(),
		IsDefault: req.GetIsDefault(),
		Status:    req.GetStatus(),
		Remark:    req.GetRemark(),
	})
	if err != nil {
		return nil, err
	}
	return &v1.EmptySysDictReply{}, nil
}

func (uc *SysDictUsecase) DeleteSysDictData(ctx context.Context, req *v1.DeleteSysDictDataRep) (*v1.EmptySysDictReply, error) {
	err := uc.repo.DeleteData(ctx, req.GetId())
	return &v1.EmptySysDictReply{}, err
}

func (uc *SysDictUsecase) GetSysDictData(ctx context.Context, req *v1.GetSysDictRep) (*v1.GetSysDictDataReply, error) {
	info, err := uc.repo.GetDataInfo(ctx, &model.SysDictDataQuery{DictCode: req.GetId()})
	if err != nil {
		return nil, err
	}
	return &v1.GetSysDictDataReply{Dict: info.BizToV1()}, nil
}

func (uc *SysDictUsecase) ListSysDictData(ctx context.Context, req *v1.ListSysDictDataRep) (*v1.ListSysDictDataReply, error) {
	list, total, err := uc.repo.GetDataPageSet(ctx, req.GetPageNum(), req.GetPageSize(), &model.SysDictDataQuery{DictType: req.GetDictType()})
	if err != nil {
		return nil, err
	}
	var rows []*v1.DictDataReply
	for _, dictData := range list {
		rows = append(rows, dictData.BizToV1())
	}
	return &v1.ListSysDictDataReply{Rows: rows, Total: total}, err
}

func (uc *SysDictUsecase) OptionSelectType(ctx context.Context, req *v1.OptionSelectTypeRep) (*v1.OptionSelectTypeReply, error) {
	selectType, err := uc.repo.OptionSelectType(ctx)
	var dicts []*v1.DictTypeReply
	for _, dictType := range selectType {
		dicts = append(dicts, dictType.BizToV1())
	}
	return &v1.OptionSelectTypeReply{Dict: dicts}, err
}

func (uc *SysDictUsecase) RefreshCacheSysDict(ctx context.Context, req *v1.CacheSysDicReq) (*v1.EmptySysDictReply, error) {
	dictType, _ := uc.repo.OptionSelectType(ctx)
	for _, sysDictType := range dictType {
		if err := uc.coreRepo.RSetObj(ctx, constants.CacheSysDictTag+sysDictType.DictType, sysDictType); err != nil {
			return nil, err
		}
		byType, _ := uc.repo.GetSysDictDataByType(ctx, sysDictType.DictType)
		if err := uc.coreRepo.RSetObj(ctx, constants.CacheSysDict+sysDictType.DictType, byType); err != nil {
			return nil, err
		}
	}
	return &v1.EmptySysDictReply{}, nil
}
