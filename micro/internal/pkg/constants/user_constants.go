package constants

const (
	// SysUser 平台内系统用户的唯一标志
	SysUser = "SYS_USER"

	// Normal 正常状态
	Normal = "0"

	// Exception 异常状态
	Exception = "1"

	// UserDisable 用户封禁状态
	UserDisable = "1"

	// RoleDisable 角色封禁状态
	RoleDisable = "1"

	// DeptNormal 部门正常状态
	DeptNormal = "0"

	// DeptDisable 部门停用状态
	DeptDisable = "1"

	// DictNormal 字典正常状态
	DictNormal = "0"

	// Yes 是否为系统默认（是）
	Yes = "Y"

	// YesFrame 是否菜单外链（是）
	YesFrame = 0

	// NoFrame 是否菜单外链（否）
	NoFrame = 1

	// TypeDir 菜单类型（目录）
	TypeDir = "M"

	// TypeMenu 菜单类型（菜单）
	TypeMenu = "C"

	// TypeButton 菜单类型（按钮）
	TypeButton = "F"

	// Layout Layout组件标识
	Layout = "Layout"

	// ParentView ParentView组件标识
	ParentView = "ParentView"

	// InnerLink 组件标识
	InnerLink = "InnerLink"

	// Unique 校验是否唯一的返回标识
	Unique    = true
	NotUnique = false

	// UsernameMinLength 用户名长度限制最小值
	UsernameMinLength = 2
	UsernameMaxLength = 20

	// PasswordMinLength 密码长度限制最小值
	PasswordMinLength = 5
	PasswordMaxLength = 20
)
