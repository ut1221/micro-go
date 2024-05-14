package model

import (
	v1 "micro/api/api/system/v1"
	"micro/internal/pkg"
	"time"
)

// SysDictType 字典类型表
type SysDictType struct {
	DictID   string `gorm:"column:dict_id;type:bigint;primaryKey;autoIncrement:true;comment:字典主键" json:"dict_id"` // 字典主键
	DictName string `gorm:"column:dict_name;type:varchar(100);comment:字典名称" json:"dict_name"`                     // 字典名称
	DictType string `gorm:"column:dict_type;type:varchar(100);comment:字典类型" json:"dict_type"`                     // 字典类型
	Status   string `gorm:"column:status;type:char(1);default:0;comment:状态（0正常 1停用）" json:"status"`               // 状态（0正常 1停用）
	Remark   string `gorm:"column:remark;type:varchar(500);comment:备注" json:"remark"`                             // 备注
	pkg.MicroUser
	SysDictData []SysDictData `gorm:"foreignKey:DictType;references:DictType" json:"sys_dict_data"`
}

// TableName SysDictType's table name
func (*SysDictType) TableName() string {
	return "sys_dict_type"
}

type SysDictTypeQuery struct {
	DictID   string
	DictName string
	DictType string
}

// SysDictData 字典数据表
type SysDictData struct {
	DictCode  string `gorm:"column:dict_code;type:bigint;primaryKey;autoIncrement:true;comment:字典编码" json:"dict_code"` // 字典编码
	DictSort  int64  `gorm:"column:dict_sort;type:int;comment:字典排序" json:"dict_sort"`                                  // 字典排序
	DictLabel string `gorm:"column:dict_label;type:varchar(100);comment:字典标签" json:"dict_label"`                       // 字典标签
	DictValue string `gorm:"column:dict_value;type:varchar(100);comment:字典键值" json:"dict_value"`                       // 字典键值
	DictType  string `gorm:"column:dict_type;type:varchar(100);comment:字典类型" json:"dict_type"`                         // 字典类型
	CSSClass  string `gorm:"column:css_class;type:varchar(100);comment:样式属性（其他样式扩展）" json:"css_class"`                 // 样式属性（其他样式扩展）
	ListClass string `gorm:"column:list_class;type:varchar(100);comment:表格回显样式" json:"list_class"`                     // 表格回显样式
	IsDefault string `gorm:"column:is_default;type:char(1);default:N;comment:是否默认（Y是 N否）" json:"is_default"`           // 是否默认（Y是 N否）
	Status    string `gorm:"column:status;type:char(1);default:0;comment:状态（0正常 1停用）" json:"status"`                   // 状态（0正常 1停用）
	Remark    string `gorm:"column:remark;type:varchar(500);comment:备注" json:"remark"`                                 // 备注
	pkg.MicroUser
}

// TableName SysDictDatum's table name
func (*SysDictData) TableName() string {
	return "sys_dict_data"
}

type SysDictDataQuery struct {
	DictCode  string
	DictLabel string
	DictType  string
}

type BizSysDictType struct {
	DictID    string
	DictName  string
	DictType  string
	Status    string
	Remark    string
	CreatedAt time.Time
}
type BizSysDictData struct {
	DictCode  string // 字典编码
	DictSort  int64  // 字典排序
	DictLabel string // 字典标签
	DictValue string // 字典键值
	DictType  string // 字典类型
	CSSClass  string // 样式属性（其他样式扩展）
	ListClass string // 表格回显样式
	IsDefault string // 是否默认（Y是 N否）
	Status    string // 状态（0正常 1停用）
	CreatedAt time.Time
	Remark    string // 备注
}

func (s *SysDictType) EntityToBiz() *BizSysDictType {
	return &BizSysDictType{
		DictID:    s.DictID,
		DictName:  s.DictName,
		DictType:  s.DictType,
		Status:    s.Status,
		Remark:    s.Remark,
		CreatedAt: s.CreatedAt,
	}
}

func (s *BizSysDictType) BizToV1() *v1.DictTypeReply {
	return &v1.DictTypeReply{
		DictId:    s.DictID,
		DictName:  s.DictName,
		DictType:  s.DictType,
		Status:    s.Status,
		Remark:    s.Remark,
		CreatedAt: s.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (s *SysDictData) EntityToBiz() *BizSysDictData {
	return &BizSysDictData{
		DictCode:  s.DictCode,
		DictSort:  s.DictSort,
		DictLabel: s.DictLabel,
		DictValue: s.DictValue,
		DictType:  s.DictType,
		CSSClass:  s.CSSClass,
		ListClass: s.ListClass,
		IsDefault: s.IsDefault,
		Status:    s.Status,
		Remark:    s.Remark,
		CreatedAt: s.CreatedAt,
	}
}

func (s *BizSysDictData) BizToV1() *v1.DictDataReply {
	return &v1.DictDataReply{
		DictCode:  s.DictCode,
		DictSort:  s.DictSort,
		DictLabel: s.DictLabel,
		DictValue: s.DictValue,
		DictType:  s.DictType,
		CssClass:  s.CSSClass,
		ListClass: s.ListClass,
		IsDefault: s.IsDefault,
		CreatedAt: s.CreatedAt.Format("2006-01-02 15:04:05"),
		Status:    s.Status,
		Remark:    s.Remark,
	}
}
