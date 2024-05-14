package model

import (
	v1 "micro/api/api/system/v1"
	"micro/internal/pkg"
	"time"
)

const TableNameSysConfig = "sys_config"

// SysConfig 参数配置表
type SysConfig struct {
	ConfigID    string `gorm:"column:config_id;type:int;primaryKey;autoIncrement:true;comment:参数主键" json:"config_id"` // 参数主键
	ConfigName  string `gorm:"column:config_name;type:varchar(100);comment:参数名称" json:"config_name"`                  // 参数名称
	ConfigKey   string `gorm:"column:config_key;type:varchar(100);comment:参数键名" json:"config_key"`                    // 参数键名
	ConfigValue string `gorm:"column:config_value;type:varchar(500);comment:参数键值" json:"config_value"`                // 参数键值
	ConfigType  string `gorm:"column:config_type;type:char(1);default:N;comment:系统内置（Y是 N否）" json:"config_type"`      // 系统内置（Y是 N否）
	Remark      string `gorm:"column:remark;type:varchar(500);comment:备注" json:"remark"`                              // 备注
	pkg.MicroUser
}

// TableName SysConfig's table name
func (*SysConfig) TableName() string {
	return TableNameSysConfig
}

type SysConfigQuery struct {
	ConfigId  string
	ConfigKey string
}

type BizSysConfig struct {
	ConfigID    string // 参数主键
	ConfigName  string // 参数名称
	ConfigKey   string // 参数键名
	ConfigValue string // 参数键值
	ConfigType  string // 系统内置（Y是 N否）
	Remark      string // 备注
	CreatedAt   time.Time
}

func (s *SysConfig) EntityToBiz() *BizSysConfig {
	return &BizSysConfig{
		ConfigID:    s.ConfigID,
		ConfigName:  s.ConfigName,
		ConfigKey:   s.ConfigKey,
		ConfigValue: s.ConfigValue,
		ConfigType:  s.ConfigType,
		Remark:      s.Remark,
		CreatedAt:   s.CreatedAt,
	}
}

func (s *BizSysConfig) BizToV1() *v1.ConfigReply {
	return &v1.ConfigReply{
		ConfigId:    s.ConfigID,
		ConfigName:  s.ConfigName,
		ConfigKey:   s.ConfigKey,
		ConfigType:  s.ConfigType,
		CreatedAt:   s.CreatedAt.Format("2006-01-02 15:04:05"),
		ConfigValue: s.ConfigValue,
		Remark:      s.Remark,
	}
}
