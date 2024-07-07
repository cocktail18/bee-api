package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"github.com/shopspring/decimal"
)

type BeeLevel struct {
	common.BaseModel
	Level         int64           `gorm:"column:level;type:int(10);comment:等级" json:"level"`
	Name          string          `gorm:"column:name;type:varchar(100);comment:等级名称" json:"name"`
	Paixu         int             `gorm:"column:paixu;type:int(10);comment:排序" json:"paixu"`
	UpgradeAmount decimal.Decimal `gorm:"column:upgrade_amount;type:decimal(10,2);comment:升级金额" json:"upgradeAmount"`
}

func (m *BeeLevel) TableName() string {
	return "bee_level"
}
