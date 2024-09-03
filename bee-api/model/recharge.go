package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"github.com/shopspring/decimal"
)

// RechargeSendRule 充值赠送规则
type RechargeSendRule struct {
	common.BaseModel
	Confine decimal.Decimal `gorm:"column:confine;type:decimal(10,2);comment:充值金额" json:"confine"`
	Send    decimal.Decimal `gorm:"column:send;type:decimal(10,2);comment:赠送金额" json:"send"`
	Loop    bool            `gorm:"column:loop;type:tinyint(1);comment:循环赠送" json:"loop"` // 每满多少送多少
}

func (RechargeSendRule) TableName() string {
	return "bee_recharge_send_rule"
}
