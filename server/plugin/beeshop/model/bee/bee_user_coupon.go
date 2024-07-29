// 自动生成模板BeeUserCoupon
package bee

import (
	"time"
)

// 用户优惠券 结构体  BeeUserCoupon
type BeeUserCoupon struct {
	Id            *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                             //id字段
	UserId        *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`                     //商店用户ID
	IsDeleted     *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`                       //已删除
	DateAdd       *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                            //创建时间
	DateUpdate    *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`                   //更新时间
	DateDelete    *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`                   //删除时间
	Uid           *int       `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;size:19;"`                                 //用户id
	DateStart     *time.Time `json:"dateStart" form:"dateStart" gorm:"column:date_start;comment:可以使用时间;"`                    //可以使用时间
	ExpiryMillis  *int       `json:"expiryMillis" form:"expiryMillis" gorm:"column:expiry_millis;comment:过期时间（毫秒）;size:19;"` //过期时间（毫秒）
	Money         *float64   `json:"money" form:"money" gorm:"column:money;comment:金额;size:20;"`                             //金额
	MoneyHreshold *int       `json:"moneyHreshold" form:"moneyHreshold" gorm:"column:money_hreshold;comment:满xx可用;size:19;"` //满xx可用
	MoneyMin      *int       `json:"moneyMin" form:"moneyMin" gorm:"column:money_min;comment:优惠券面额范围;size:19;"`              //优惠券面额范围
	MoneyMax      *int       `json:"moneyMax" form:"moneyMax" gorm:"column:money_max;comment:优惠券面额范围;size:19;"`              //优惠券面额范围
	MoneyType     *int       `json:"moneyType" form:"moneyType" gorm:"column:money_type;comment:优惠券类型;size:19;"`             //优惠券类型
	Name          string     `json:"name" form:"name" gorm:"column:name;comment:优惠券名字;size:100;"`                            //优惠券名字
	OnlyFreight   *bool      `json:"onlyFreight" form:"onlyFreight" gorm:"column:only_freight;comment:仅抵扣运费;"`               //仅抵扣运费
	Pid           *int       `json:"pid" form:"pid" gorm:"column:pid;comment:优惠券id;size:19;"`                                //优惠券id
	ShopId        *int       `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:店铺id;size:19;"`                       //店铺id
	Source        string     `json:"source" form:"source" gorm:"column:source;comment:管理员批量发放;size:100;"`                    //管理员批量发放
	Status        *int       `json:"status" form:"status" gorm:"column:status;comment:使用状态，0未使用，1使用中，2已使用;size:10;"`         //使用状态，0未使用，1使用中，2已使用
}

// TableName 用户优惠券 BeeUserCoupon自定义表名 bee_user_coupon
func (BeeUserCoupon) TableName() string {
	return "bee_user_coupon"
}
