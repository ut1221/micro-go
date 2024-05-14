package model

import (
	v1 "micro/api/api/system/v1"
	"micro/internal/pkg"
	"time"
)

const TableNameSysUser = "sys_user"

// SysUser 用户信息表
type SysUser struct {
	UserID      string    `gorm:"column:user_id;type:bigint;primaryKey;autoIncrement:true;comment:用户ID" json:"user_id"` // 用户ID
	DeptID      string    `gorm:"column:dept_id;type:bigint;comment:部门ID" json:"dept_id"`                               // 部门ID
	UserName    string    `gorm:"column:user_name;type:varchar(30);not null;comment:用户账号" json:"user_name"`             // 用户账号
	NickName    string    `gorm:"column:nick_name;type:varchar(30);not null;comment:用户昵称" json:"nick_name"`             // 用户昵称
	UserType    string    `gorm:"column:user_type;type:varchar(2);default:00;comment:用户类型（00系统用户）" json:"user_type"`    // 用户类型（00系统用户）
	Email       string    `gorm:"column:email;type:varchar(50);comment:用户邮箱" json:"email"`                              // 用户邮箱
	PhoneNumber string    `gorm:"column:phone_number;type:varchar(11);comment:手机号码" json:"phone_number"`                // 手机号码
	Sex         string    `gorm:"column:sex;type:char(1);default:0;comment:用户性别（0男 1女 2未知）" json:"sex"`                 // 用户性别（0男 1女 2未知）
	Avatar      string    `gorm:"column:avatar;type:varchar(100);comment:头像地址" json:"avatar"`                           // 头像地址
	Password    string    `gorm:"column:password;type:varchar(100);comment:密码" json:"password"`                         // 密码
	Status      string    `gorm:"column:status;type:char(1);default:0;comment:帐号状态（0正常 1停用）" json:"status"`             // 帐号状态（0正常 1停用）
	LoginIP     string    `gorm:"column:login_ip;type:varchar(128);comment:最后登录IP" json:"login_ip"`                     // 最后登录IP
	LoginDate   time.Time `gorm:"column:login_date;type:datetime;comment:最后登录时间" json:"login_date"`                     // 最后登录时间
	Remark      string    `gorm:"column:remark;type:varchar(500);comment:备注" json:"remark"`                             // 备注
	pkg.MicroUser
	SysRole []SysRole `gorm:"-" json:"roles"`
	SysDept SysDept   `gorm:"foreignKey:DeptID;references:DeptID" json:"sys_dept"`
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}

type SysUserQuery struct {
	UserId      string
	UserName    string
	Status      int64
	DeptId      string
	RoleId      string
	PhoneNumber string
	BeginTime   string
	EndTime     string
}

type BizSysUser struct {
	UserID      string    // 用户ID
	DeptID      string    // 部门ID
	UserName    string    // 用户账号
	NickName    string    // 用户昵称
	UserType    string    // 用户类型（00系统用户）
	Email       string    // 用户邮箱
	PhoneNumber string    // 手机号码
	Sex         string    // 用户性别（0男 1女 2未知）
	Avatar      string    // 头像地址
	Password    string    // 密码
	Status      string    // 帐号状态（0正常 1停用）
	LoginIP     string    // 最后登录IP
	LoginDate   time.Time // 最后登录时间
	CreateBy    string    // 创建者
	CreatedAt   time.Time // 创建时间
	UpdateBy    string    // 更新者
	UpdatedAt   time.Time // 更新时间
	Remark      string    // 备注
	Roles       []BizSysRole
	Dept        BizSysDept
}

func (user *SysUser) EntityToBiz() *BizSysUser {
	return &BizSysUser{
		UserID:      user.UserID,
		DeptID:      user.DeptID,
		UserName:    user.UserName,
		NickName:    user.NickName,
		UserType:    user.UserType,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Sex:         user.Sex,
		Avatar:      user.Avatar,
		Password:    user.Password,
		Status:      user.Status,
		LoginIP:     user.LoginIP,
		LoginDate:   user.LoginDate,
		CreatedAt:   user.CreatedAt,
		Remark:      user.Remark,
	}
}

func (user *BizSysUser) BizToV1() *v1.UserReply {
	return &v1.UserReply{
		UserId:      user.UserID,
		DeptId:      user.DeptID,
		UserName:    user.UserName,
		NickName:    user.NickName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Sex:         user.Sex,
		Avatar:      user.Avatar,
		Password:    user.Password,
		Status:      user.Status,
		LoginIp:     user.LoginIP,
		LoginDate:   user.LoginDate.Format("2006-01-02 15:04:05"),
		Remark:      user.Remark,
		CreatedAt:   user.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}
