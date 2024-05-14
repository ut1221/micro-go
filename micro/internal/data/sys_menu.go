package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"micro/internal/biz"
	"micro/internal/model"
	"micro/internal/pkg/constants"
)

type sysMenuRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysMenuRepo(data *Data, logger log.Logger) biz.SysMenuRepo {
	return &sysMenuRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s sysMenuRepo) GetMenuByUserId(ctx context.Context, userId string, isAdmin bool) ([]*model.BizSysMenu, error) {
	var sysMenus []*model.SysMenu
	if isAdmin {
		if err := s.data.Db.Model(&model.SysMenu{}).
			Where("menu_type in ?", []string{constants.TypeDir, constants.TypeMenu}).
			Where("status = '0'").
			Order("parent_id, order_num").
			Find(&sysMenus).Error; err != nil {
			return nil, err
		}
	} else {
		db := s.data.Db.Model(&model.SysMenu{}).Distinct().
			Joins("left join sys_role_menu on sys_menu.menu_id = sys_role_menu.menu_id").
			Joins("left join sys_user_role on sys_role_menu.role_id = sys_user_role.role_id").
			Joins("left join sys_role on sys_user_role.role_id = sys_role.role_id").
			Joins("left join sys_user on sys_user_role.user_id = sys_user.user_id").
			Where("sys_user.user_id = ? and sys_menu.menu_type in ? and sys_menu.status = 0 AND sys_role.status = 0 AND sys_role.deleted_at is null", userId, []string{constants.TypeDir, constants.TypeMenu})
		if tx := db.Order("sys_menu.parent_id, sys_menu.order_num").Find(&sysMenus); tx.Error != nil {
			return nil, tx.Error
		}
	}
	var menus []*model.BizSysMenu
	for _, sysMenu := range sysMenus {
		menus = append(menus, sysMenu.EntityToBiz())
	}
	return menus, nil
}

func (s sysMenuRepo) GetMenusList(ctx context.Context, query *model.SysMenuQuery) ([]*model.BizSysMenu, error) {
	var sysMenus []*model.SysMenu
	db := s.data.Db.Model(&model.SysMenu{})
	if query.RoleId != "" {
		db.Joins("INNER JOIN sys_role_menu ON sys_menu.menu_id = sys_role_menu.menu_id and sys_role_menu.deleted_at is null").
			Where("sys_menu.status = 0 and sys_role_menu.role_id = ?", query.RoleId)
	}
	if err := db.Find(&sysMenus).Error; err != nil {
		return nil, err
	}
	var bizMenus []*model.BizSysMenu
	for _, sysMenu := range sysMenus {
		toV1 := sysMenu.EntityToBiz()
		bizMenus = append(bizMenus, toV1)
	}
	return bizMenus, nil
}

func (s sysMenuRepo) ModifyMenuForRole(ctx context.Context, roleId string, menuIds []string) error {
	if err := s.data.Db.Where("role_id = ?", roleId).Delete(&model.SysRoleMenu{}).Error; err != nil {
		return err
	}
	if len(menuIds) == 0 {
		return nil
	}
	var sysRoleMenus []model.SysRoleMenu
	for _, menuId := range menuIds {
		sysRoleMenus = append(sysRoleMenus, model.SysRoleMenu{
			RoleID: roleId,
			MenuID: menuId,
		})
	}
	return s.data.Db.CreateInBatches(&sysRoleMenus, 100).Error
}

func (s sysMenuRepo) GetMenuInfo(ctx context.Context, query model.SysMenuQuery) (*model.BizSysMenu, error) {
	var menu model.SysMenu
	db := s.data.Db.Model(&model.SysMenu{})
	if query.MenuID != "" {
		db.Where("menu_id = ?", query.MenuID)
	}
	if err := db.First(&menu).Error; err != nil {
		return nil, err
	}
	return menu.EntityToBiz(), nil
}

func (s sysMenuRepo) Save(ctx context.Context, menu *model.SysMenu) error {
	return s.data.Db.Save(&menu).Error
}

func (s sysMenuRepo) Update(ctx context.Context, menu *model.SysMenu) error {
	return s.data.Db.Model(&model.SysMenu{}).Where("menu_id = ?", menu.MenuID).Updates(&menu).Error
}

func (s sysMenuRepo) Delete(ctx context.Context, menuId string) error {
	return s.data.Db.Where("menu_id = ?", menuId).Delete(&model.SysMenu{}).Error
}

func (s sysMenuRepo) DeleteRoleMenu(ctx context.Context, menuId string, roleId string) error {
	if menuId != "" {
		return s.data.Db.Where("menu_id = ?", menuId).Delete(&model.SysRoleMenu{}).Error
	}
	if roleId != "" {
		return s.data.Db.Where("role_id = ?", roleId).Delete(&model.SysRoleMenu{}).Error
	}
	return errors.New(constants.DeleteError)
}
