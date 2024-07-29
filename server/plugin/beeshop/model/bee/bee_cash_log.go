// 自动生成模板BeeCashLog
package bee

import (
	"time"
)

// 用户现金消费记录 结构体  BeeCashLog
type BeeCashLog struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                 //id字段
	UserId     *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`         //商店用户ID
	IsDeleted  *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`           //已删除
	DateAdd    *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                //创建时间
	DateUpdate *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`       //更新时间
	DateDelete *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`       //删除时间
	Amount     *float64   `json:"amount" form:"amount" gorm:"column:amount;comment:金额;size:10;"`              //金额
	Balance    *float64   `json:"balance" form:"balance" gorm:"column:balance;comment:剩余可用余额;size:10;"`       //剩余可用余额
	Behavior   *int       `json:"behavior" form:"behavior" gorm:"column:behavior;comment:0 收入 1 支出;size:10;"` //0 收入 1 支出
	Freeze     *float64   `json:"freeze" form:"freeze" gorm:"column:freeze;comment:剩余冻结金额;size:10;"`          //剩余冻结金额
	Remark     string     `json:"remark" form:"remark" gorm:"column:remark;comment:remark;size:100;"`         //remark
	Type       *int       `json:"type" form:"type" gorm:"column:type;comment:类型;size:10;"`                    //类型
	Uid        *int       `json:"uid" form:"uid" gorm:"column:uid;comment:;size:19;"`                         //uid字段
}

// TableName 用户现金消费记录 BeeCashLog自定义表名 bee_cash_log
func (BeeCashLog) TableName() string {
	return "bee_cash_log"
}
