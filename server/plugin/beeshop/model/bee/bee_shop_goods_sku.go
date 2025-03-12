// 自动生成模板BeeShopGoodsSku
package bee

import (
	"time"
)

// 商品sku 结构体  BeeShopGoodsSku
type BeeShopGoodsSku struct {
	Id               *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                                            //id字段
	UserId           *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`                                //商店用户ID
	IsDeleted        *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`                                   //已删除
	DateAdd          *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                                       //创建时间
	DateUpdate       *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`                              //更新时间
	DateDelete       *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`                              //删除时间
	GoodsId          *int       `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品id;size:19;"`                                 //商品id
	Code             string     `json:"code" form:"code" gorm:"column:code;comment:sku编号;size:100;"`                                         //sku编号
	FxType           *int       `json:"fxType" form:"fxType" gorm:"column:fx_type;comment:返现类型;size:10;"`                                  //返现类型
	OriginalPrice    *float64   `json:"originalPrice" form:"originalPrice" gorm:"column:original_price;comment:市场价;size:10;"`               //市场价
	PingtuanPrice    *float64   `json:"pingtuanPrice" form:"pingtuanPrice" gorm:"column:pingtuan_price;comment:拼团价;size:10;"`               //拼团价
	Price            *float64   `json:"price" form:"price" gorm:"column:price;comment:价格;size:10;"`                                          //价格
	VipPrice         *float64   `json:"vipPrice" form:"vipPrice" gorm:"column:vip_price;comment:会员价格;size:10;"`                            //价格
	PropertyChildIds string     `json:"propertyChildIds" form:"propertyChildIds" gorm:"column:property_child_ids;comment:属性关系;size:1000;"` //属性关系
	Score            *float64   `json:"score" form:"score" gorm:"column:score;comment:需要积分;size:10;"`                                      //需要积分
	Stores           *int       `json:"stores" form:"stores" gorm:"column:stores;comment:库存;size:19;"`                                       //库存
	Weight           *float64   `json:"weight" form:"weight" gorm:"column:weight;comment:重量;size:10;"`                                       //重量
}

// TableName 商品sku BeeShopGoodsSku自定义表名 bee_shop_goods_sku
func (BeeShopGoodsSku) TableName() string {
	return "bee_shop_goods_sku"
}
