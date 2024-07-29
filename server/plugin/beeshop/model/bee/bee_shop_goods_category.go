// 自动生成模板BeeShopGoodsCategory
package bee

import (
	"time"
)

// 商品分类 结构体  BeeShopGoodsCategory
type BeeShopGoodsCategory struct {
	Id         *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`           //id字段
	UserId     *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`   //商店用户ID
	IsDeleted  *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`     //已删除
	DateAdd    *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`          //创建时间
	DateUpdate *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"` //更新时间
	DateDelete *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"` //删除时间
	IsUse      *bool      `json:"isUse" form:"isUse" gorm:"column:is_use;comment:是否启用;"`                //是否启用
	Level      *int       `json:"level" form:"level" gorm:"column:level;comment:等级;size:19;"`           //等级
	Name       string     `json:"name" form:"name" gorm:"column:name;comment:名称;size:100;"`             //名称
	Paixu      *int       `json:"paixu" form:"paixu" gorm:"column:paixu;comment:排序;size:19;"`           //排序
	Pid        *int       `json:"pid" form:"pid" gorm:"column:pid;comment:父id;size:19;"`                //父id
	ShopId     *int       `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:店铺id;size:19;"`     //店铺id
}

// TableName 商品分类 BeeShopGoodsCategory自定义表名 bee_shop_goods_category
func (BeeShopGoodsCategory) TableName() string {
	return "bee_shop_goods_category"
}
