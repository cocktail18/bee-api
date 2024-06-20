package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/shopspring/decimal"
)

type BeeShopGoodsSku struct {
	common.BaseModel
	GoodsId          int64           `gorm:"column:goods_id;type:bigint(20);comment:商品id" json:"goodsId"`
	Code             string          `gorm:"column:code;type:varchar(100);comment:sku编号" json:"code"`
	FxType           enum.FxType     `gorm:"column:fx_type;type:int(11);comment:返现类型" json:"fxType"`
	OriginalPrice    decimal.Decimal `gorm:"column:original_price;type:decimal(10,2);comment:市场价" json:"originalPrice"`
	PingtuanPrice    decimal.Decimal `gorm:"column:pingtuan_price;type:decimal(10,2);comment:拼团价" json:"pingtuanPrice"`
	Price            decimal.Decimal `gorm:"column:price;type:decimal(10,2);comment:价格" json:"price"`
	PropertyChildIds string          `gorm:"column:property_child_ids;type:varchar(1000);comment:属性关系" json:"propertyChildIds"`
	Score            decimal.Decimal `gorm:"column:score;type:decimal(10,2);comment:需要积分" json:"score"`
	Stores           int64           `gorm:"column:stores;type:bigint(20);comment:库存" json:"stores"`
	Weight           decimal.Decimal `gorm:"column:weight;type:decimal(10,2);comment:重量" json:"weight"`

	PropertyChildNames string `gorm:"-" json:"propertyChildNames"`
}

func (m *BeeShopGoodsSku) TableName() string {
	return "bee_shop_goods_sku"
}
