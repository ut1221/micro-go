package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"micro/internal/biz"
	"micro/internal/model"
)

type sysDeptRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysDeptRepo(data *Data, logger log.Logger) biz.SysDeptRepo {
	return &sysDeptRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s sysDeptRepo) GetSysDeptInfo(ctx context.Context, query model.SysDeptQuery) (*model.BizSysDept, error) {
	var sysDept model.SysDept
	db := s.data.Db.Model(&model.SysDept{})
	if query.DeptId != "" {
		db = db.Where("dept_id = ?", query.DeptId)
	}
	if query.DeptName != "" {
		db = db.Where("dept_name = ?", query.DeptName)
	}
	if err := db.First(&sysDept).Error; err != nil {
		return nil, err
	}
	return sysDept.EntityToBiz(), nil
}

func (s sysDeptRepo) GetDeptList(ctx context.Context, query model.SysDeptQuery) ([]*model.BizSysDept, error) {
	var sysDepths []*model.SysDept
	db := s.data.Db.Model(&sysDepths)
	if query.RoleId != "" {
		db.Joins("inner join sys_role_dept on sys_dept.dept_id = sys_role_dept.dept_id and sys_role_dept.deleted_at is null")
		db.Where("sys_role_dept.role_id = ?", query.RoleId)
	}
	if err := db.Find(&sysDepths).Error; err != nil {
		return nil, err
	}
	var bizDepths []*model.BizSysDept
	for _, dept := range sysDepths {
		toBiz := dept.EntityToBiz()
		bizDepths = append(bizDepths, toBiz)
	}
	return bizDepths, nil
}

func (s sysDeptRepo) ModifyDeptForRole(ctx context.Context, roleId string, deptIds []string) error {
	if err := s.data.Db.Where("role_id = ?", roleId).Delete(&model.SysRoleDept{}).Error; err != nil {
		return err
	}
	if len(deptIds) == 0 {
		return nil
	}
	var sysRoleDepts []model.SysRoleDept
	for _, deptId := range deptIds {
		sysRoleDepts = append(sysRoleDepts, model.SysRoleDept{DeptID: deptId, RoleID: roleId})
	}
	return s.data.Db.CreateInBatches(&sysRoleDepts, 100).Error
}

func (s sysDeptRepo) Save(ctx context.Context, dept *model.SysDept) error {
	return s.data.Db.Save(&dept).Error
}

func (s sysDeptRepo) Update(ctx context.Context, dept *model.SysDept) error {
	return s.data.Db.Model(&model.SysDept{}).Where("dept_id = ?", dept.DeptID).Updates(&dept).Error
}

func (s sysDeptRepo) Delete(ctx context.Context, deptId string) error {
	return s.data.Db.Where("dept_id = ?", deptId).Delete(&model.SysDept{}).Error
}

func (s sysDeptRepo) DeleteSysDept(ctx context.Context, deptId string) error {
	if err := s.data.Db.Where("dept_id = ?", deptId).Delete(&model.SysRoleDept{}).Error; err != nil {
		return err
	}
	return nil
}
