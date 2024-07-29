// 自动生成模板BeeWxConfig
package bee

import (
	"time"
)

// 微信配置 结构体  BeeWxConfig
type BeeWxConfig struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`             //id字段
	UserId     *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`     //商店用户ID
	IsDeleted  *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`       //已删除
	DateAdd    *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`            //创建时间
	DateUpdate *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`   //更新时间
	DateDelete *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`   //删除时间
	AppId      string     `json:"appId" form:"appId" gorm:"column:app_id;comment:;size:100;"`             //appId字段
	AppSecret  string     `json:"appSecret" form:"appSecret" gorm:"column:app_secret;comment:;size:100;"` //appSecret字段
}

// TableName 微信配置 BeeWxConfig自定义表名 bee_wx_config
func (BeeWxConfig) TableName() string {
	return "bee_wx_config"
}
