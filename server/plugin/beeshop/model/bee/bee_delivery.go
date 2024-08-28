// 自动生成模板BeeDelivery
package bee

import (
	"time"
)

// beeDelivery表 结构体  BeeDelivery
type BeeDelivery struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                       //id字段
	UserId     *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`               //商店用户ID
	IsDeleted  *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`                 //已删除
	DateAdd    *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                      //创建时间
	DateUpdate *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`             //更新时间
	DateDelete *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`             //删除时间
	Type       *int       `json:"type" form:"type" gorm:"column:type;comment:类型;size:19;"`                          //类型
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:名称;size:100;"`                         //名称
	AppKey     string     `json:"appKey" form:"appKey" gorm:"column:app_key;comment:app_key;size:100;"`             //app_key
	AppSecret  string     `json:"appSecret" form:"appSecret" gorm:"column:app_secret;comment:app_secret;size:100;"` //app_secret
	SourceId   string     `json:"sourceId" form:"sourceId" gorm:"column:source_id;comment:source_id;size:100;"`     //source_id
	Debug      *bool      `json:"debug" form:"debug" gorm:"column:debug;comment:debug模式;"`                          //debug模式
}

// TableName beeDelivery表 BeeDelivery自定义表名 bee_delivery
func (BeeDelivery) TableName() string {
	return "bee_delivery"
}
