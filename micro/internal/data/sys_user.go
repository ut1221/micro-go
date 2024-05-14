package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"google.golang.org/grpc/status"
	"micro/internal/biz"
	"micro/internal/model"
	"micro/internal/pkg"
)

type sysUserRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysUserRepo(data *Data, logger log.Logger) biz.SysUserRepo {
	return &sysUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s sysUserRepo) GetUserByUserName(ctx context.Context, username string) (*model.BizSysUser, error) {
	var sysUser model.SysUser
	tx := s.data.Db.Where(&model.SysUser{UserName: username}).First(&sysUser)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, status.Error(404, "用户不存在")
	}
	return sysUser.EntityToBiz(), nil
}

func (s sysUserRepo) GetUserById(ctx context.Context, userId string) (*model.BizSysUser, error) {
	var sysUser model.SysUser
	db := s.data.Db.Preload("SysDept")
	tx := db.Where(&model.SysUser{UserID: userId}).First(&sysUser)
	if tx.Error != nil {
		return nil, tx.Error
	}
	if tx.RowsAffected == 0 {
		return nil, status.Error(404, "用户不存在")
	}
	var bizUser *model.BizSysUser
	bizUser = sysUser.EntityToBiz()
	bizUser.Dept = *sysUser.SysDept.EntityToBiz()
	return bizUser, nil
}

func (s sysUserRepo) GetPageSet(ctx context.Context, pageNum int64, pageSize int64, query model.SysUserQuery) ([]*model.BizSysUser, int64, error) {
	var sysUsers []*model.SysUser
	var total int64
	db := s.data.Db.Preload("SysDept")
	if query.DeptId != "" {
		db.Where("FIND_IN_SET(dept_id, GetDeptSubordinates(?))", query.DeptId)
	}
	if query.UserId != "" {
		db.Where("user_id = ?", query.UserId)
	}
	if query.UserName != "" {
		db.Where("user_name = ?", query.UserName)
	}
	if query.Status != 0 {
		db.Where("status = ?", query.Status)
	}
	if query.PhoneNumber != "" {
		db.Where("phone_number = ?", query.PhoneNumber)
	}
	if query.BeginTime != "" {
		db.Where("created_at >= ?", query.BeginTime)
	}
	if query.EndTime != "" {
		db.Where("created_at <= ?", query.EndTime)
	}
	db.Model(&model.SysUser{}).Count(&total)
	tx := db.Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&sysUsers)
	if tx.Error != nil {
		log.Error(tx.Error)
		return nil, 0, tx.Error
	}
	var bizUsers []*model.BizSysUser
	for _, sysUser := range sysUsers {
		bizUser := sysUser.EntityToBiz()
		bizUser.Dept = *sysUser.SysDept.EntityToBiz()
		bizUsers = append(bizUsers, bizUser)
	}
	return bizUsers, total, nil
}

func (s sysUserRepo) Save(ctx context.Context, user *model.SysUser) error {
	return s.data.Db.Save(&user).Error
}

func (s sysUserRepo) Update(ctx context.Context, user *model.BizSysUser) error {
	db := s.data.Db.Model(&model.SysUser{}).Where("user_id = ?", user.UserID)
	if user.UserName != "" {
		db.Update("user_name", user.UserName)
	}
	if user.NickName != "" {
		db.Update("nick_name", user.NickName)
	}
	if !user.LoginDate.IsZero() {
		db.Update("login_data", user.LoginDate)
	}
	if user.LoginIP != "" {
		db.Update("login_ip", user.LoginIP)
	}
	if user.DeptID != "" {
		db.Update("dept_id", user.DeptID)
	}
	if user.Email != "" {
		db.Update("email", user.Email)
	}
	if user.PhoneNumber != "" {
		db.Update("phone_number", user.PhoneNumber)
	}
	if user.Password != "" {
		db.Update("password", pkg.Encrypt(user.Password))
	}
	return db.Error
}

func (s sysUserRepo) Delete(ctx context.Context, userId string) error {
	return s.data.Db.Where("user_id = ?", userId).Delete(&model.SysUser{}).Error
}

// GetRoleUser 获取角色用户
// authUser 是否有权限
func (s sysUserRepo) GetRoleUser(ctx context.Context, pageNum int64, pageSize int64, authUser bool, query model.SysUserQuery) ([]*model.BizSysUser, int64, error) {
	var sysUsers []*model.SysUser
	var total int64
	db := s.data.Db.Model(&model.SysUser{}).Where("status = '0'")
	if authUser {
		db.Where("user_id in (?)", s.data.Db.Model(&model.SysUserRole{}).Select("user_id").Where("role_id = ?", query.RoleId).Find(&model.SysUserRole{}))
	} else {
		db.Where("user_id not in (?)", s.data.Db.Model(&model.SysUserRole{}).Select("user_id").Where("role_id = ?", query.RoleId).Find(&model.SysUserRole{}))
	}
	if err := db.Count(&total).Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&sysUsers).Error; err != nil {
		return nil, 0, err
	}
	var bizUsers []*model.BizSysUser
	for _, sysUser := range sysUsers {
		bizUser := sysUser.EntityToBiz()
		bizUser.Dept = *sysUser.SysDept.EntityToBiz()
		bizUsers = append(bizUsers, bizUser)
	}
	return bizUsers, total, nil
}
