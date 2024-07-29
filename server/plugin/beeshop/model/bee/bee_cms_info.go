// 自动生成模板BeeCmsInfo
package bee

import (
	"time"
)

// cms信息 结构体  BeeCmsInfo
type BeeCmsInfo struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`           //id字段
	UserId     *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`   //商店用户ID
	IsDeleted  *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`     //已删除
	DateAdd    *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`          //创建时间
	DateUpdate *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"` //更新时间
	DateDelete *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"` //删除时间
	Content    string     `json:"content" form:"content" gorm:"column:content;comment:内容;size:100;"`    //内容
	Key        string     `json:"key" form:"key" gorm:"column:key;comment:key;size:100;"`               //key
	Title      string     `json:"title" form:"title" gorm:"column:title;comment:标题;size:100;"`          //标题
	Type       *int       `json:"type" form:"type" gorm:"column:type;comment:类型;size:19;"`              //类型
}

// TableName cms信息 BeeCmsInfo自定义表名 bee_cms_info
func (BeeCmsInfo) TableName() string {
	return "bee_cms_info"
}
