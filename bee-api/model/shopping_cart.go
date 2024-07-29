package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"github.com/shopspring/decimal"
)

// BeeShoppingCart 购物车
type BeeShoppingCart struct {
	common.BaseModel
	Uid              int64           `gorm:"column:uid;type:bigint(11);comment:uid" json:"uid"`
	GoodsId          int64           `gorm:"column:goods_id;type:bigint(11);comment:商品id" json:"goodsId"`
	CategoryId       int64           `gorm:"column:category_id;type:bigint(11);comment:分类id" json:"categoryId"`
	Key              string          `gorm:"column:key;type:varchar(100);comment:唯一key" json:"key"`
	LogisticsId      int64           `gorm:"column:logistics_id;type:bigint(11);comment:物流id" json:"logisticsId"`
	Name             string          `gorm:"column:name;type:varchar(100);comment:商品名" json:"name"`
	Number           int64           `gorm:"column:number;type:bigint(11);comment:数量" json:"number"`
	Price            decimal.Decimal `gorm:"column:price;type:decimal(10,2);comment:加入购物车价格" json:"price"`
	Selected         bool            `gorm:"column:selected;type:tinyint(1);comment:勾选" json:"selected"`
	ShopId           int64           `gorm:"column:shop_id;type:bigint(11);comment:店铺id" json:"shopId"`
	SkuId            int64           `gorm:"column:sku_id;type:bigint(20);comment:sku id" json:"skuId"`
	PropertyChildIds string          `gorm:"column:property_child_ids;type:varchar(100);comment:属性" json:"propertyChildIds"`
}

func (m *BeeShoppingCart) TableName() string {
	return "bee_shopping_cart"
}
