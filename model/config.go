package model

import (
	"gitee.com/stuinfer/bee-api/common"
)

type BeeConfig struct {
	common.BaseModel
	Key    string `gorm:"column:key;type:varchar(100)" json:"key"`
	Remark string `gorm:"column:remark;type:varchar(100);NOT NULL" json:"remark"`
	Value  string `gorm:"column:value;type:varchar(1000);NOT NULL" json:"value"`
}

func (m *BeeConfig) TableName() string {
	return "bee_config"
}

type BeeWxConfig struct {
	common.BaseModel
	AppId     string `gorm:"column:app_id;type:varchar(100);NOT NULL" json:"appId"`
	AppSecret string `gorm:"column:app_secret;type:varchar(100);NOT NULL" json:"appSecret"`
}

func (m *BeeWxConfig) TableName() string {
	return "bee_wx_config"
}
