// 自动生成模板BeeRechargeSendRule
package bee

import (
	"github.com/shopspring/decimal"
	"time"
)

// beeRechargeSendRule表 结构体  BeeRechargeSendRule
type BeeRechargeSendRule struct {
	Id         *int            `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`           //id字段
	UserId     *int            `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`   //商店用户ID
	IsDeleted  *bool           `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`     //已删除
	DateAdd    *time.Time      `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`          //创建时间
	DateUpdate *time.Time      `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"` //更新时间
	DateDelete *time.Time      `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"` //删除时间
	Confine    decimal.Decimal `json:"confine" form:"confine" gorm:"column:confine;comment:充值金额;size:10;"`   //充值金额
	Send       decimal.Decimal `json:"send" form:"send" gorm:"column:send;comment:赠送金额;size:10;"`            //赠送金额
	Loop       *bool           `json:"loop" form:"loop" gorm:"column:loop;comment:循环赠送;"`                    //循环赠送
}

// TableName beeRechargeSendRule表 BeeRechargeSendRule自定义表名 bee_recharge_send_rule
func (BeeRechargeSendRule) TableName() string {
	return "bee_recharge_send_rule"
}
