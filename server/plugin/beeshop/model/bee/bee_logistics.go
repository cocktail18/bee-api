// 自动生成模板BeeLogistics
package bee

import (
	"time"
)

// 运费模版 结构体  BeeLogistics
type BeeLogistics struct {
	Id                  *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`         //id字段
	UserId              *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"` //商店用户ID
	Name                string     `json:"name" form:"name" gorm:"column:name;comment:模板名;size:100;"`          //模板名
	DetailsJsonStr      string     `gorm:"column:details_json;type:text" json:"detailsJson" form:"detailsJson" `
	FreeShippingSetting string     `gorm:"column:free_shipping_setting;type:text" json:"freeShippingSetting" form:"freeShippingSetting"`
	FeeType             *int       `json:"feeType" form:"feeType" gorm:"column:fee_type;comment:计价方式，1按件，2按重量，3按金额;size:10;"` //计价方式，1按件，2按重量，3按金额
	IsFree              *bool      `json:"isFree" form:"isFree" gorm:"column:is_free;comment:是否包邮;"`                          //是否包邮
	IsDeleted           *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`                  //已删除
	DateAdd             *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                       //创建时间
	DateUpdate          *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`              //更新时间
	DateDelete          *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`              //删除时间
}

// TableName 运费模版 BeeLogistics自定义表名 bee_logistics
func (BeeLogistics) TableName() string {
	return "bee_logistics"
}
