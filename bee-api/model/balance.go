package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/shopspring/decimal"
)

type BeeUserBalanceLog struct {
	common.BaseModel
	OrderId     string           `gorm:"column:order_id;type:varchar(150);comment:订单id" json:"order_id"`
	BalanceType enum.BalanceType `gorm:"column:balance_type;type:int(11);comment:1金币，2积分" json:"balance_type"`
	Num         decimal.Decimal  `gorm:"column:num;type:decimal(20,2);comment:数量" json:"num"`
	Uid         int64            `gorm:"column:uid;type:bigint(20);comment:用户id" json:"uid"`
	Mark        string           `gorm:"column:mark;type:varchar(200);comment:备注" json:"mark"`
}

func (m *BeeUserBalanceLog) TableName() string {
	return "bee_user_balance_log"
}
