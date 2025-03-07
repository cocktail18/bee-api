// 自动生成模板BeeUserBalanceLog
package bee

import (
	"time"
)

// 用户消费记录 结构体  BeeUserBalanceLog
type BeeUserBalanceLog struct {
	Id          *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                         //id字段
	UserId      *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`                 //商店用户ID
	IsDeleted   *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`                   //已删除
	DateAdd     *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                        //创建时间
	DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`               //更新时间
	DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`               //删除时间
	OrderId     string     `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单id;size:150;"`               //订单id
	BalanceType *int       `json:"balanceType" form:"balanceType" gorm:"column:balance_type;comment:1金币，2积分;size:10;"` //1金币，2积分
	Num         *float64   `json:"num" form:"num" gorm:"column:num;comment:数量;size:20;"`                               //数量
	Uid         *int       `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;size:19;"`                             //用户id
	Mark        string     `json:"mark" form:"mark" gorm:"column:mark;comment:备注;size:200;"`                           //备注
}

// TableName 用户消费记录 BeeUserBalanceLog自定义表名 bee_user_balance_log
func (BeeUserBalanceLog) TableName() string {
	return "bee_user_balance_log"
}
