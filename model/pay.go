package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/shopspring/decimal"
)

type BeePayLog struct {
	common.BaseModel
	Money      decimal.Decimal   `gorm:"column:money;type:decimal(10,2);comment:总金额" json:"money"`
	NextAction string            `gorm:"column:next_action;type:varchar(100);comment:next_action" json:"nextAction"`
	OrderNo    string            `gorm:"column:order_no;type:varchar(100);comment:订单号" json:"orderNo"`
	PayGate    enum.PayGate      `gorm:"column:pay_gate;type:varchar(100);comment:pay_gate" json:"payGate"`
	PayGateStr string            `gorm:"-" json:"payGateStr"`
	Remark     string            `gorm:"column:remark;type:varchar(100);comment:备注" json:"remark"`
	Status     enum.PayLogStatus `gorm:"column:status;type:int(11);comment:状态" json:"status"`
	StatusStr  string            `gorm:"-" json:"statusStr"`
	Uid        int64             `gorm:"column:uid;type:bigint(11);comment:用户id" json:"uid"`
}

func (b *BeePayLog) TableName() string {
	return "bee_pay_log"
}

func (b *BeePayLog) FillData() {
	b.StatusStr = enum.PayLogStatusMap[b.Status]
	b.PayGateStr = enum.PayGateMap[b.PayGate]
}
