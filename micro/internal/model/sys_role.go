package model

import (
	"gorm.io/gorm"
	v1 "micro/api/api/system/v1"
	"micro/internal/pkg"
	"time"
)

const TableNameSysUserRole = "sys_user_role"

// SysUserRole 角色信息表
type SysUserRole struct {
	RoleID    string         `gorm:"column:role_id;type:bigint;comment:角色ID" json:"role_id"` // 角色ID
	UserID    string         `gorm:"column:user_id;type:bigint;comment:用户ID" json:"user_id"` // 角色ID
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName SysRole's table name
func (*SysUserRole) TableName() string {
	return TableNameSysUserRole
}

const TableNameSysRole = "sys_role"

// SysRole 角色信息表
type SysRole struct {
	RoleID            string `gorm:"column:role_id;type:bigint;primaryKey;autoIncrement:true;comment:角色ID" json:"role_id"`                              // 角色ID
	RoleName          string `gorm:"column:role_name;type:varchar(30);not null;comment:角色名称" json:"role_name"`                                          // 角色名称
	RoleKey           string `gorm:"column:role_key;type:varchar(100);not null;comment:角色权限字符串" json:"role_key"`                                        // 角色权限字符串
	RoleSort          int64  `gorm:"column:role_sort;type:int;not null;comment:显示顺序" json:"role_sort"`                                                  // 显示顺序
	DataScope         string `gorm:"column:data_scope;type:char(1);default:1;comment:数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）" json:"data_scope"` // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly int64  `gorm:"column:menu_check_strictly;type:tinyint(1);default:1;comment:菜单树选择项是否关联显示" json:"menu_check_strictly"`              // 菜单树选择项是否关联显示
	DeptCheckStrictly int64  `gorm:"column:dept_check_strictly;type:tinyint(1);default:1;comment:部门树选择项是否关联显示" json:"dept_check_strictly"`              // 部门树选择项是否关联显示
	Status            string `gorm:"column:status;type:char(1);not null;comment:角色状态（0正常 1停用）" json:"status"`                                           // 角色状态（0正常 1停用）
	Remark            string `gorm:"column:remark;type:varchar(500);comment:备注" json:"remark"`                                                          // 备注
	pkg.MicroUser
}

// TableName SysRole's table name
func (*SysRole) TableName() string {
	return TableNameSysRole
}

type SysRoleQuery struct {
	UserId    string
	RoleID    string // 角色ID
	RoleName  string // 角色名称
	RoleKey   string // 角色权限字符串
	DataScope string // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	Status    string // 角色状态（0正常 1停用）
}

type BizSysRole struct {
	RoleID            string // 角色ID
	RoleName          string // 角色名称
	RoleKey           string // 角色权限字符串
	RoleSort          int64  // 显示顺序
	DataScope         string // 数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）
	MenuCheckStrictly int64  // 菜单树选择项是否关联显示
	DeptCheckStrictly int64  // 部门树选择项是否关联显示
	Status            string // 角色状态（0正常 1停用）
	Remark            string // 备注
	CreatedAt         time.Time
	UpdatedAt         time.Time
	CreateBy          string
	UpdateBy          string
}

func (role *SysRole) EntityToBiz() *BizSysRole {
	return &BizSysRole{
		RoleID:            role.RoleID,
		RoleName:          role.RoleName,
		RoleKey:           role.RoleKey,
		RoleSort:          role.RoleSort,
		DataScope:         role.DataScope,
		MenuCheckStrictly: role.MenuCheckStrictly,
		DeptCheckStrictly: role.DeptCheckStrictly,
		Status:            role.Status,
		Remark:            role.Remark,
		CreatedAt:         role.CreatedAt,
	}
}

func (role *BizSysRole) BizToV1() *v1.RoleReply {
	return &v1.RoleReply{
		RoleId:            role.RoleID,
		RoleName:          role.RoleName,
		RoleKey:           role.RoleKey,
		RoleSort:          role.RoleSort,
		DataScope:         role.DataScope,
		MenuCheckStrictly: coverBool(role.MenuCheckStrictly),
		DeptCheckStrictly: coverBool(role.DeptCheckStrictly),
		Status:            role.Status,
		Remark:            role.Remark,
		CreatedAt:         role.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func coverBool(num int64) bool {
	return num == 1
}
