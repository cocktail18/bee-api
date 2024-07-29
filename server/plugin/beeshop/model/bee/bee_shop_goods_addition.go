// 自动生成模板BeeShopGoodsAddition
package bee

import (
	"time"
)

// 商品额外信息 结构体  BeeShopGoodsAddition
type BeeShopGoodsAddition struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`           //id字段
	UserId     *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`   //商店用户ID
	IsDeleted  *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`     //已删除
	DateAdd    *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`          //创建时间
	DateUpdate *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"` //更新时间
	DateDelete *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"` //删除时间
	GoodsId    *int       `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:;size:19;"`      //goodsId字段
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:;size:100;"`               //name字段
	Require    *bool      `json:"require" form:"require" gorm:"column:require;comment:;"`               //require字段
	Type       *int       `json:"type" form:"type" gorm:"column:type;comment:;size:10;"`                //type字段
	Pid        *int       `json:"pid" form:"pid" gorm:"column:pid;comment:;size:19;"`                   //pid字段
	Price      *float64   `json:"price" form:"price" gorm:"column:price;comment:价格;size:10;"`           //价格
}

// TableName 商品额外信息 BeeShopGoodsAddition自定义表名 bee_shop_goods_addition
func (BeeShopGoodsAddition) TableName() string {
	return "bee_shop_goods_addition"
}
