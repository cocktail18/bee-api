// 自动生成模板BeePrinter
package bee

import (
	"time"
)

// beePrinter表 结构体  BeePrinter
type BeePrinter struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                       //id字段
	UserId     *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`               //商店用户ID
	IsDeleted  *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`                 //已删除
	DateAdd    *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                      //创建时间
	DateUpdate *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`             //更新时间
	DateDelete *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`             //删除时间
	Appid      string     `json:"appid" form:"appid" gorm:"column:appid;comment:appid;size:100;"`                   //appid
	AppSecret  string     `json:"appSecret" form:"appSecret" gorm:"column:app_secret;comment:app_secret;size:100;"` //app_secret
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:打印机名称;size:100;"`                      //打印机名称
	Key        string     `json:"key" form:"key" gorm:"column:key;comment:打印机密钥;size:100;"`                         //打印机密钥
	Brand      *int       `json:"brand" form:"brand" gorm:"column:brand;comment:品牌;size:10;"`                       //品牌
	Condition  *int       `json:"condition" form:"condition" gorm:"column:condition;comment:打印条件;size:10;"`         //打印条件
	Code       string     `json:"code" form:"code" gorm:"column:code;comment:打印机编号;size:100;"`                      //打印机编号
	Num        *int       `json:"num" form:"num" gorm:"column:num;comment:打印份数;size:10;"`                           //打印份数
	ShopId     *int       `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:店铺id;size:19;"`                 //店铺id
	Template   string     `json:"template" form:"template" gorm:"column:template;comment:打印模板;size:100;"`           //打印模板
}

// TableName beePrinter表 BeePrinter自定义表名 bee_printer
func (BeePrinter) TableName() string {
	return "bee_printer"
}
