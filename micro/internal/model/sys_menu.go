package model

import (
	"gorm.io/gorm"
	v1 "micro/api/system/v1"
	"micro/internal/pkg"
	"strconv"
	"time"
)

const TableNameRoleMenu = "sys_role_menu"

// SysRoleMenu 角色信息表
type SysRoleMenu struct {
	MenuID    string         `gorm:"column:menu_id;type:bigint;comment:菜单ID" json:"menu_id"` // 菜单ID
	RoleID    string         `gorm:"column:role_id;type:bigint;comment:角色ID" json:"role_id"` // 角色ID
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName SysRoleMenu table name
func (*SysRoleMenu) TableName() string {
	return TableNameRoleMenu
}

const TableNameSysMenu = "sys_menu"

// SysMenu 菜单权限表
type SysMenu struct {
	MenuID    string `gorm:"column:menu_id;type:bigint;primaryKey;comment:菜单ID" json:"menu_id"`          // 菜单ID
	MenuName  string `gorm:"column:menu_name;type:varchar(50);not null;comment:菜单名称" json:"menu_name"`   // 菜单名称
	ParentID  string `gorm:"column:parent_id;type:bigint;comment:父菜单ID" json:"parent_id"`                // 父菜单ID
	OrderNum  int64  `gorm:"column:order_num;type:int;comment:显示顺序" json:"order_num"`                    // 显示顺序
	Path      string `gorm:"column:path;type:varchar(200);comment:路由地址" json:"path"`                     // 路由地址
	Component string `gorm:"column:component;type:varchar(255);comment:组件路径" json:"component"`           // 组件路径
	Query     string `gorm:"column:query;type:varchar(255);comment:路由参数" json:"query"`                   // 路由参数
	IsFrame   int64  `gorm:"column:is_frame;type:int;default:1;comment:是否为外链（0是 1否）" json:"is_frame"`    // 是否为外链（0是 1否）
	IsCache   int64  `gorm:"column:is_cache;type:int;comment:是否缓存（0缓存 1不缓存）" json:"is_cache"`            // 是否缓存（0缓存 1不缓存）
	MenuType  string `gorm:"column:menu_type;type:char(1);comment:菜单类型（M目录 C菜单 F按钮）" json:"menu_type"`   // 菜单类型（M目录 C菜单 F按钮）
	Visible   string `gorm:"column:visible;type:char(1);default:0;comment:菜单状态（0显示 1隐藏）" json:"visible"` // 菜单状态（0显示 1隐藏）
	Status    string `gorm:"column:status;type:char(1);default:0;comment:菜单状态（0正常 1停用）" json:"status"`   // 菜单状态（0正常 1停用）
	Perms     string `gorm:"column:perms;type:varchar(100);comment:权限标识" json:"perms"`                   // 权限标识
	Icon      string `gorm:"column:icon;type:varchar(100);default:#;comment:菜单图标" json:"icon"`           // 菜单图标
	Remark    string `gorm:"column:remark;type:varchar(500);comment:备注" json:"remark"`                   // 备注
	pkg.MicroUser
}

// TableName SysMenu's table name
func (*SysMenu) TableName() string {
	return TableNameSysMenu
}

type SysMenuQuery struct {
	RoleId string
	MenuID string
}

type BizSysMenu struct {
	MenuID    string // 菜单ID
	MenuName  string // 菜单名称
	ParentID  string // 父菜单ID
	OrderNum  int64  // 显示顺序
	Path      string // 路由地址
	Component string // 组件路径
	Query     string // 路由参数
	IsFrame   int64  // 是否为外链（0是 1否）
	IsCache   int64  // 是否缓存（0缓存 1不缓存）
	MenuType  string // 菜单类型（M目录 C菜单 F按钮）
	Visible   string // 菜单状态（0显示 1隐藏）
	Status    string // 菜单状态（0正常 1停用）
	Perms     string // 权限标识
	Icon      string // 菜单图标
	Remark    string // 备注
	CreatedAt time.Time
}

func (menu *SysMenu) EntityToBiz() *BizSysMenu {
	return &BizSysMenu{
		MenuID:    menu.MenuID,
		MenuName:  menu.MenuName,
		ParentID:  menu.ParentID,
		OrderNum:  menu.OrderNum,
		Path:      menu.Path,
		Component: menu.Component,
		Query:     menu.Query,
		IsFrame:   menu.IsFrame,
		IsCache:   menu.IsCache,
		MenuType:  menu.MenuType,
		Visible:   menu.Visible,
		Status:    menu.Status,
		Perms:     menu.Perms,
		Icon:      menu.Icon,
		Remark:    menu.Remark,
		CreatedAt: menu.CreatedAt,
	}
}

func (m *BizSysMenu) BizToV1() *v1.MenuReply {
	return &v1.MenuReply{
		MenuId:    m.MenuID,
		MenuName:  m.MenuName,
		ParentId:  m.ParentID,
		OrderNum:  m.OrderNum,
		Path:      m.Path,
		Component: m.Component,
		Query:     m.Query,
		IsFrame:   strconv.FormatInt(m.IsFrame, 10),
		IsCache:   strconv.FormatInt(m.IsCache, 10),
		MenuType:  m.MenuType,
		Visible:   m.Visible,
		Status:    m.Status,
		Perms:     m.Perms,
		Icon:      m.Icon,
		CreatedAt: m.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
