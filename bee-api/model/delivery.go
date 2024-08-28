package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/enum"
)

// BeeDelivery 配送配置
type BeeDelivery struct {
	common.BaseModel
	Type      enum.DeliveryType `gorm:"column:type;type:bigint(11);comment:类型" json:"type"`
	Name      string            `gorm:"column:name;type:varchar(100);comment:名称" json:"name"`
	AppKey    string            `gorm:"column:app_key;type:varchar(100);comment:app_key" json:"app_key"`
	AppSecret string            `gorm:"column:app_secret;type:varchar(100);comment:app_secret" json:"app_secret"`
	SourceId  string            `gorm:"column:source_id;type:varchar(100);comment:source_id" json:"source_id"`
	Debug     bool              `gorm:"column:debug;type:tinyint(1);comment:debug模式" json:"debug"`
}

func (b *BeeDelivery) TableName() string {
	return "bee_delivery"
}
