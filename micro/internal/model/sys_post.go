package model

import (
	"gorm.io/gorm"
	v1 "micro/api/system/v1"
	"micro/internal/pkg"
	"time"
)

const TableNameSysUserPost = "sys_user_post"

// SysUserPost 角色信息表
type SysUserPost struct {
	PostID    string         `gorm:"column:post_id;type:bigint;primaryKey;comment:岗位ID" json:"post_id"` // 岗位ID
	UserID    string         `gorm:"column:user_id;type:bigint;primaryKey;comment:用户ID" json:"user_id"` // 角色ID
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

// TableName SysRole's table name
func (*SysUserPost) TableName() string {
	return TableNameSysUserPost
}

const TableNameSysPost = "sys_post"

// SysPost 岗位信息表
type SysPost struct {
	PostID   string `gorm:"column:post_id;type:bigint;primaryKey;autoIncrement:true;comment:岗位ID" json:"post_id"` // 岗位ID
	PostCode string `gorm:"column:post_code;type:varchar(64);not null;comment:岗位编码" json:"post_code"`             // 岗位编码
	PostName string `gorm:"column:post_name;type:varchar(50);not null;comment:岗位名称" json:"post_name"`             // 岗位名称
	PostSort int64  `gorm:"column:post_sort;type:int;not null;comment:显示顺序" json:"post_sort"`                     // 显示顺序
	Status   string `gorm:"column:status;type:char(1);not null;comment:状态（0正常 1停用）" json:"status"`                // 状态（0正常 1停用）
	Remark   string `gorm:"column:remark;type:varchar(500);comment:备注" json:"remark"`                             // 备注
	pkg.MicroUser
}

// TableName SysPost's table name
func (*SysPost) TableName() string {
	return TableNameSysPost
}

type SysPostQuery struct {
	UserId   string
	PostID   string
	PostCode string
	PostName string
}

type BizSysPost struct {
	PostID    string
	PostCode  string
	PostName  string
	PostSort  int64
	Status    string
	Remark    string
	CreatedAt time.Time
}

func (uc *SysPost) EntityToBiz() *BizSysPost {
	return &BizSysPost{
		PostID:    uc.PostID,
		PostCode:  uc.PostCode,
		PostName:  uc.PostName,
		PostSort:  uc.PostSort,
		Status:    uc.Status,
		Remark:    uc.Remark,
		CreatedAt: uc.CreatedAt,
	}
}

func (uc *BizSysPost) BizToV1() *v1.PostReply {
	return &v1.PostReply{
		PostId:    uc.PostID,
		PostCode:  uc.PostCode,
		PostName:  uc.PostName,
		PostSort:  uc.PostSort,
		Status:    uc.Status,
		Remark:    uc.Remark,
		CreatedAt: uc.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
