// 自动生成模板BeeOrderGoods
package bee

import (
	"time"
)

// 订单商品信息 结构体  BeeOrderGoods
type BeeOrderGoods struct {
    Id  *int `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`  //id字段 
    UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`  //商店用户ID 
    IsDeleted  *bool `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`  //已删除 
    DateAdd  *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:添加时间;size:100;"`  //添加时间 
    DateUpdate  *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`  //更新时间 
    DateDelete  *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`  //删除时间 
    AfterSale  string `json:"afterSale" form:"afterSale" gorm:"column:after_sale;comment:售后类型;size:100;"`  //售后类型 
    Amount  *float64 `json:"amount" form:"amount" gorm:"column:amount;comment:总金额;size:10;"`  //总金额 
    AmountCoupon  *float64 `json:"amountCoupon" form:"amountCoupon" gorm:"column:amount_coupon;comment:优惠券金额;size:10;"`  //优惠券金额 
    AmountSingle  *float64 `json:"amountSingle" form:"amountSingle" gorm:"column:amount_single;comment:单价;size:10;"`  //单价 
    AmountSingleBase  *float64 `json:"amountSingleBase" form:"amountSingleBase" gorm:"column:amount_single_base;comment:基本单价;size:10;"`  //基本单价 
    BuyRewardEnd  *bool `json:"buyRewardEnd" form:"buyRewardEnd" gorm:"column:buy_reward_end;comment:奖励发放结束;"`  //奖励发放结束 
    CategoryId  *int `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:分类;size:19;"`  //分类 
    CyTableStatus  *int `json:"cyTableStatus" form:"cyTableStatus" gorm:"column:cy_table_status;comment:餐饮桌子状态;size:19;"`  //餐饮桌子状态 
    FxType  *int `json:"fxType" form:"fxType" gorm:"column:fx_type;comment:返现类型;size:19;"`  //返现类型 
    GoodsId  *int `json:"goodsId" form:"goodsId" gorm:"column:goods_id;comment:商品id;size:19;"`  //商品id 
    GoodsName  string `json:"goodsName" form:"goodsName" gorm:"column:goods_name;comment:商品名;size:100;"`  //商品名 
    IsScoreOrder  *bool `json:"isScoreOrder" form:"isScoreOrder" gorm:"column:is_score_order;comment:是否积分订单;"`  //是否积分订单 
    Number  *int `json:"number" form:"number" gorm:"column:number;comment:数量;size:19;"`  //数量 
    NumberNoFahuo  *int `json:"numberNoFahuo" form:"numberNoFahuo" gorm:"column:number_no_fahuo;comment:未发货数量;size:19;"`  //未发货数量 
    OrderId  *int `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单id;size:19;"`  //订单id 
    Persion  *int `json:"persion" form:"persion" gorm:"column:persion;comment:人数;size:19;"`  //人数 
    Pic  string `json:"pic" form:"pic" gorm:"column:pic;comment:商品图;size:100;"`  //商品图 
    PriceId  *int `json:"priceId" form:"priceId" gorm:"column:price_id;comment:价格id;size:19;"`  //价格id 
    Property  string `json:"property" form:"property" gorm:"column:property;comment:属性;size:100;"`  //属性 
    PropertyChildIds  string `json:"propertyChildIds" form:"propertyChildIds" gorm:"column:property_child_ids;comment:属性id;size:100;"`  //属性id 
    Purchase  string `json:"purchase" form:"purchase" gorm:"column:purchase;comment:是否已支付;size:100;"`  //是否已支付 
    RefundStatus  *int `json:"refundStatus" form:"refundStatus" gorm:"column:refund_status;comment:退款状态;size:19;"`  //退款状态 
    Score  *int `json:"score" form:"score" gorm:"column:score;comment:积分;size:19;"`  //积分 
    ShopId  *int `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:商店id;size:19;"`  //商店id 
    Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态;size:19;"`  //状态 
    Type  *int `json:"type" form:"type" gorm:"column:type;comment:类型;size:19;"`  //类型 
    Uid  *int `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;size:19;"`  //用户id 
    Unit  string `json:"unit" form:"unit" gorm:"column:unit;comment:单位;size:100;"`  //单位 
}


// TableName 订单商品信息 BeeOrderGoods自定义表名 bee_order_goods
func (BeeOrderGoods) TableName() string {
    return "bee_order_goods"
}

