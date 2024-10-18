package model

import (
	"gorm.io/gorm"
	v1 "micro/api/system/v1"
	"micro/internal/pkg"
	"time"
)

const TableNameRoleDept = "sys_role_dept"

// SysRoleDept 角色信息表
type SysRoleDept struct {
	DeptID    string         `gorm:"column:dept_id;type:bigint;comment:部门id" json:"dept_id"` // 部门id
	RoleID    string         `gorm:"column:role_id;type:bigint;comment:角色ID" json:"role_id"` // 角色ID
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName SysRoleDept table name
func (*SysRoleDept) TableName() string {
	return TableNameRoleDept
}

const TableNameSysDept = "sys_dept"

// SysDept 部门表
type SysDept struct {
	DeptID    string `gorm:"column:dept_id;type:bigint;primaryKey;comment:部门id" json:"dept_id"`        // 部门id
	ParentID  string `gorm:"column:parent_id;type:bigint;comment:父部门id" json:"parent_id"`              // 父部门id
	Ancestors string `gorm:"column:ancestors;type:varchar(50);comment:祖级列表" json:"ancestors"`          // 祖级列表
	DeptName  string `gorm:"column:dept_name;type:varchar(30);comment:部门名称" json:"dept_name"`          // 部门名称
	OrderNum  int64  `gorm:"column:order_num;type:int;comment:显示顺序" json:"order_num"`                  // 显示顺序
	Leader    string `gorm:"column:leader;type:varchar(20);comment:负责人" json:"leader"`                 // 负责人
	Phone     string `gorm:"column:phone;type:varchar(11);comment:联系电话" json:"phone"`                  // 联系电话
	Email     string `gorm:"column:email;type:varchar(50);comment:邮箱" json:"email"`                    // 邮箱
	Status    string `gorm:"column:status;type:char(1);default:0;comment:部门状态（0正常 1停用）" json:"status"` // 部门状态（0正常 1停用）
	pkg.MicroUser
}

// TableName SysDept's table name
func (*SysDept) TableName() string {
	return TableNameSysDept
}

type SysDeptQuery struct {
	RoleId   string
	DeptId   string
	DeptName string
	ParentId string
}

type BizSysDept struct {
	DeptID    string    // 部门id
	ParentID  string    // 父部门id
	Ancestors string    // 祖级列表
	DeptName  string    // 部门名称
	OrderNum  int64     // 显示顺序
	Leader    string    // 负责人
	Phone     string    // 联系电话
	Email     string    // 邮箱
	Status    string    // 部门状态（0正常 1停用）
	CreateBy  string    // 创建者
	CreateAt  time.Time // 创建时间
}

func (s *SysDept) EntityToBiz() *BizSysDept {
	return &BizSysDept{
		DeptID:    s.DeptID,
		ParentID:  s.ParentID,
		Ancestors: s.Ancestors,
		DeptName:  s.DeptName,
		OrderNum:  s.OrderNum,
		Leader:    s.Leader,
		Phone:     s.Phone,
		Email:     s.Email,
		Status:    s.Status,
		CreateAt:  s.CreatedAt,
	}
}

func (b *BizSysDept) BizToV1() *v1.DeptReply {
	return &v1.DeptReply{
		DeptId:    b.DeptID,
		ParentId:  b.ParentID,
		Ancestors: b.Ancestors,
		DeptName:  b.DeptName,
		OrderNum:  b.OrderNum,
		Leader:    b.Leader,
		Phone:     b.Phone,
		Email:     b.Email,
		Status:    b.Status,
		CreatedAt: b.CreateAt.Format("2006-01-02 15:04:05"),
	}
}
