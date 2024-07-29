package model

import (
	"gitee.com/stuinfer/bee-api/common"
)

type BeeConfig struct {
	common.BaseModel
	Key    string `gorm:"column:key;type:varchar(100);default:''" json:"key"`
	Remark string `gorm:"column:remark;type:varchar(100);NOT NULL;default:'';comment:备注" json:"remark"`
	Value  string `gorm:"column:value;type:varchar(1000);NOT NULL;default:'';comment:值" json:"value"`
}

func (m *BeeConfig) TableName() string {
	return "bee_config"
}

type BeeWxConfig struct {
	common.BaseModel
	AppId     string `gorm:"column:app_id;type:varchar(100);default:'';NOT NULL" json:"appId"`
	AppSecret string `gorm:"column:app_secret;type:varchar(100);default:'';NOT NULL" json:"appSecret"`
	Token     string `gorm:"column:token;type:varchar(100);default:'';NOT NULL" json:"token"`
}

func (m *BeeWxConfig) TableName() string {
	return "bee_wx_config"
}

type BeeWxPayConfig struct {
	common.BaseModel
	MchId       string `gorm:"column:mchid;type:varchar(100);NOT NULL;default:''" json:"mchid"`
	AppId       string `gorm:"column:app_id;type:varchar(100);NOT NULL;default:''" json:"appId"`
	AppSecret   string `gorm:"column:app_secret;type:varchar(100);NOT NULL;default:''" json:"appSecret"`
	Token       string `gorm:"column:token;type:varchar(100);NOT NULL;default:''" json:"token"`
	PrivateCert string `gorm:"column:private_cert;type:varchar(2000);NOT NULL;default:'';comment:v3私钥" json:"privateCert"` // 私钥
	NotifyUrl   string `gorm:"column:notify_url;type:varchar(255);NOT NULL;default:'';comment:回调地址" json:"notifyUrl"`
	Debug       bool   `gorm:"column:debug;type:tinyint(1);NOT NULL;default:0" json:"debug"`
}

func (m *BeeWxPayConfig) TableName() string {
	return "bee_wx_pay_config"
}
