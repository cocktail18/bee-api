package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/enum"
)

// BeePrinter 打印机配置
type BeePrinter struct {
	common.BaseModel
	Appid     string                `gorm:"column:appid;type:varchar(100);comment:appid" json:"appid"`
	AppSecret string                `gorm:"column:app_secret;type:varchar(100);comment:app_secret" json:"app_secret"`
	Name      string                `gorm:"column:name;type:varchar(100);comment:打印机名称" json:"name"`
	Key       string                `gorm:"column:key;type:varchar(100);comment:打印机密钥" json:"key"`
	Brand     enum.PrinterBrand     `gorm:"column:brand;type:int(11);comment:品牌" json:"brand"`
	Condition enum.PrinterCondition `gorm:"column:condition;type:int(11);comment:打印条件" json:"condition"`
	Code      string                `gorm:"column:code;type:varchar(100);comment:打印机编号" json:"code"`
	Num       int                   `gorm:"column:num;type:int(11);comment:打印份数" json:"num"`
	ShopId    int64                 `gorm:"column:shop_id;type:bigint(11);comment:店铺id" json:"shopId"`
	Template  string                `gorm:"column:template;type:varchar(100);comment:打印模板" json:"template"`
	Voice     string                `gorm:"column:voice;type:varchar(100);comment:音源" json:"voice"`
}

func (m *BeePrinter) TableName() string {
	return "bee_printer"
}
