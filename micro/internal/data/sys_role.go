package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"micro/internal/biz"
	"micro/internal/model"
)

type sysRoleRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysRoleRepo(data *Data, logger log.Logger) biz.SysRoleRepo {
	return &sysRoleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s sysRoleRepo) GetRoleInfo(ctx context.Context, query *model.SysRoleQuery) (*model.BizSysRole, error) {
	var sysRole model.SysRole
	db := s.data.Db.Model(&model.SysRole{})
	if query.RoleID != "" {
		db.Where("role_id = ?", query.RoleID)
	}
	if err := db.First(&sysRole).Error; err != nil {
		return nil, err
	}
	return sysRole.EntityToBiz(), nil
}

func (s sysRoleRepo) GetListRoles(ctx context.Context, query model.SysRoleQuery) ([]*model.BizSysRole, error) {
	var sysRoles []model.SysRole
	db := s.data.Db.Table("sys_role")
	if query.UserId != "" {
		db.Joins("INNER JOIN sys_user_role ON sys_role.role_id = sys_user_role.role_id and sys_user_role.deleted_at is null").
			Where("sys_user_role.user_id = ?", query.UserId)
	}
	if err := db.Find(&sysRoles).Error; err != nil {
		return nil, err
	}
	var bizRoles []*model.BizSysRole
	for _, sysRole := range sysRoles {
		bizRole := sysRole.EntityToBiz()
		bizRoles = append(bizRoles, bizRole)
	}
	return bizRoles, nil
}

func (s sysRoleRepo) ModifyRoleForUser(ctx context.Context, userId string, roleIds []string) error {
	tx := s.data.Db.Where("user_id = ?", userId).Delete(&model.SysUserRole{})
	if tx.Error != nil {
		return tx.Error
	}
	if len(roleIds) == 0 {
		return nil
	}
	var sysUserRoles = make([]model.SysUserRole, 0)
	for _, roleId := range roleIds {
		sysUserRoles = append(sysUserRoles, model.SysUserRole{
			UserID: userId,
			RoleID: roleId,
		})
	}
	return s.data.Db.CreateInBatches(sysUserRoles, 100).Error
}

func (s sysRoleRepo) AuthUserSelectAll(ctx context.Context, roleId string, userIds []string) error {
	//tx := s.data.Db.Where("role_id = ?", roleId).Delete(&model.SysUserRole{})
	//if tx.Error != nil {
	//	return tx.Error
	//}
	if len(userIds) == 0 {
		return nil
	}
	var sysUserRoles = make([]model.SysUserRole, 0)
	for _, userId := range userIds {
		sysUserRoles = append(sysUserRoles, model.SysUserRole{
			UserID: userId,
			RoleID: roleId,
		})
	}
	return s.data.Db.CreateInBatches(sysUserRoles, 100).Error
}

func (s sysRoleRepo) AuthUserCancel(ctx context.Context, roleId string, userIds []string) error {
	return s.data.Db.Where("role_id = ?", roleId).
		Where("user_id in ?", userIds).
		Delete(&model.SysUserRole{}).Error
}
func (s sysRoleRepo) GetPageSet(ctx context.Context, pageNum int64, pageSize int64, query model.SysRoleQuery) ([]*model.BizSysRole, int64, error) {
	var sysRoles []*model.SysRole
	var total int64
	db := s.data.Db.Model(&model.SysRole{}).Count(&total)
	tx := db.Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&sysRoles)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	var bizRoles []*model.BizSysRole
	for _, sysRole := range sysRoles {
		toBiz := sysRole.EntityToBiz()
		bizRoles = append(bizRoles, toBiz)
	}
	return bizRoles, total, nil
}

func (s sysRoleRepo) Save(ctx context.Context, role *model.SysRole) error {
	return s.data.Db.Save(&role).Error
}

func (s sysRoleRepo) Update(ctx context.Context, role *model.SysRole) error {
	return s.data.Db.Model(&model.SysRole{}).Where("role_id = ?", role.RoleID).Updates(&role).Error
}

func (s sysRoleRepo) Delete(ctx context.Context, roleId string) error {
	return s.data.Db.Where("role_id = ?", roleId).Delete(&model.SysRole{}).Error
}
