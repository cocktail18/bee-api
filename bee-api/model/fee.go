package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/shopspring/decimal"
)

// 配送费
type BeePeiSong struct {
	common.BaseModel

	Distance   decimal.Decimal     `gorm:"column:distance;type:decimal(10,2);comment:距离，单位km" json:"distance"`
	Fwf1Min    decimal.Decimal     `gorm:"column:fwf1_min;type:decimal(10,2);comment:最低收费金额" json:"fwf1Min"` // ● 最低收费金额
	Fwf1Name   string              `gorm:"column:fwf1_name;type:decimal(10,2);comment:收费项目名称" json:"fwf1Name"`
	Fwf1Number decimal.Decimal     `gorm:"column:fwf1_number;type:decimal(10,2);comment:比例或者金额" json:"fwf1Number"` //收取的比例或者固定金额
	Fwf1Type   enum.BeePeiSongType `gorm:"column:fwf1_type;type:int(10);comment:类型" json:"fwf1Type"`               //0 固定金额； 1 百分比

	ZtDiscounts decimal.Decimal `gorm:"column:zt_discounts;type:decimal(10,2);comment:自提补贴" json:"ztDiscounts"` //自提补贴，如果用户选择自提，订单金额减少
}

func (m *BeePeiSong) TableName() string {
	return "bee_pei_song"
}
