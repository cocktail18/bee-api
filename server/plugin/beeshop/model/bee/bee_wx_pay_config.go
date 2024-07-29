// 自动生成模板BeeWxPayConfig
package bee

import (
	"time"
)

// 微信支付配置 结构体  BeeWxPayConfig
type BeeWxPayConfig struct {
	Id          *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                        //id字段
	UserId      *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`                //商店用户ID
	Mchid       string     `json:"mchid" form:"mchid" gorm:"column:mchid;comment:;size:100;"`                         //mchid字段
	AppId       string     `json:"appId" form:"appId" gorm:"column:app_id;comment:;size:100;"`                        //appId字段
	AppSecret   string     `json:"appSecret" form:"appSecret" gorm:"column:app_secret;comment:;size:100;"`            //appSecret字段
	Token       string     `json:"token" form:"token" gorm:"column:token;comment:;size:100;"`                         //token字段
	PrivateCert string     `json:"privateCert" form:"privateCert" gorm:"column:private_cert;comment:v3私钥;size:2000;"` //v3私钥
	NotifyUrl   string     `json:"notifyUrl" form:"notifyUrl" gorm:"column:notify_url;comment:;size:200;"`
	Debug       *bool      `json:"debug" form:"debug" gorm:"column:debug;comment:;size:1;"`
	IsDeleted   *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`     //已删除
	DateAdd     *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`          //创建时间
	DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"` //更新时间
	DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"` //删除时间
}

// TableName 微信支付配置 BeeWxPayConfig自定义表名 bee_wx_pay_config
func (BeeWxPayConfig) TableName() string {
	return "bee_wx_pay_config"
}
