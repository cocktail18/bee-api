// 自动生成模板BeeOrderLog
package bee

import (
	"time"
)

// 订单日志 结构体  BeeOrderLog
type BeeOrderLog struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`           //id字段
	UserId     *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`   //商店用户ID
	OrderId    *int       `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单id;size:19;"`  //订单id
	Type       *int       `json:"type" form:"type" gorm:"column:type;comment:类型;size:19;"`              //类型
	Remark     string     `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:100;"`       //备注
	IsDeleted  *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`     //已删除
	DateAdd    *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`          //创建时间
	DateUpdate *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"` //更新时间
	DateDelete *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"` //删除时间
}

// TableName 订单日志 BeeOrderLog自定义表名 bee_order_log
func (BeeOrderLog) TableName() string {
	return "bee_order_log"
}
