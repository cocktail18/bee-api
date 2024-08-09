// 自动生成模板BeeShopGoods
package bee

import (
	"time"
)

// 商品表 结构体  BeeShopGoods
type BeeShopGoods struct {
	Id                   *int       `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                                                      //id字段
	UserId               *int       `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`                                              //商店用户ID
	IsDeleted            *bool      `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`                                                //已删除
	DateAdd              *time.Time `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                                                     //创建时间
	DateUpdate           *time.Time `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`                                            //更新时间
	DateDelete           *time.Time `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`                                            //删除时间
	BarCode              string     `json:"barCode" form:"barCode" gorm:"column:bar_code;comment:条码;size:100;"`                                              //条码
	AfterSale            string     `json:"afterSale" form:"afterSale" gorm:"column:after_sale;comment:售后说明,1仅退款，2退款退货，3换货;size:100;"`                       //售后说明,1仅退款，2退款退货，3换货
	CategoryId           *int       `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:分类id;size:19;"`                                    //分类id
	Commission           *float64   `json:"commission" form:"commission" gorm:"column:commission;comment:积分奖励;size:10;"`                                     //积分奖励
	CommissionSettleType *int       `json:"commissionSettleType" form:"commissionSettleType" gorm:"column:commission_settle_type;comment:积分奖励结算类型;size:19;"` //积分奖励结算类型
	CommissionType       *int       `json:"commissionType" form:"commissionType" gorm:"column:commission_type;comment:积分奖励类型;size:19;"`                      //积分奖励类型
	CommissionUserType   *int       `json:"commissionUserType" form:"commissionUserType" gorm:"column:commission_user_type;comment:积分奖励用户类型;size:19;"`       //积分奖励用户类型
	HasAddition          *bool      `json:"hasAddition" form:"hasAddition" gorm:"column:has_addition;comment:有附加;"`                                          //有附加
	HasTourJourney       *bool      `json:"hasTourJourney" form:"hasTourJourney" gorm:"column:has_tour_journey;comment:有团购;"`                                //有团购
	Hidden               *bool      `json:"hidden" form:"hidden" gorm:"column:hidden;comment:隐藏;size:19;"`                                                   //隐藏
	Kanjia               *bool      `json:"kanjia" form:"kanjia" gorm:"column:kanjia;comment:允许砍价;"`                                                         //允许砍价
	KanjiaPrice          *float64   `json:"kanjiaPrice" form:"kanjiaPrice" gorm:"column:kanjia_price;comment:砍价价格;size:10;"`                                 //砍价价格
	Limitation           *bool      `json:"limitation" form:"limitation" gorm:"column:limitation;comment:限制;"`                                               //限制
	LogisticsId          *int       `json:"logisticsId" form:"logisticsId" gorm:"column:logistics_id;comment:物流id;size:19;"`                                 //物流id
	MaxCoupons           *int       `json:"maxCoupons" form:"maxCoupons" gorm:"column:max_coupons;comment:最大使用优惠券数量;size:19;"`                               //最大使用优惠券数量
	Miaosha              *bool      `json:"miaosha" form:"miaosha" gorm:"column:miaosha;comment:秒杀;"`                                                        //秒杀
	MinBuyNumber         *int       `json:"minBuyNumber" form:"minBuyNumber" gorm:"column:min_buy_number;comment:最低购买数量;size:19;"`                           //最低购买数量
	MinPrice             *float64   `json:"minPrice" form:"minPrice" gorm:"column:min_price;comment:最低价格;size:10;"`                                          //最低价格
	MinScore             *float64   `json:"minScore" form:"minScore" gorm:"column:min_score;comment:需要积分数量;size:10;"`                                        //需要积分数量
	Name                 string     `json:"name" form:"name" gorm:"column:name;comment:商品名字;size:100;"`                                                      //商品名字
	NumberFav            *int       `json:"numberFav" form:"numberFav" gorm:"column:number_fav;comment:喜欢数量;size:19;"`                                       //喜欢数量
	NumberGoodReputation *float64   `json:"numberGoodReputation" form:"numberGoodReputation" gorm:"column:number_good_reputation;comment:商品评分;size:10;"`     //商品评分
	NumberOrders         *int       `json:"numberOrders" form:"numberOrders" gorm:"column:number_orders;comment:订单数量;size:19;"`                              //订单数量
	NumberReputation     *int       `json:"numberReputation" form:"numberReputation" gorm:"column:number_reputation;comment:评分数量;size:19;"`                  //评分数量
	NumberSells          *int       `json:"numberSells" form:"numberSells" gorm:"column:number_sells;comment:销售数量;size:19;"`                                 //销售数量
	OriginalPrice        *float64   `json:"originalPrice" form:"originalPrice" gorm:"column:original_price;comment:原价;size:10;"`                             //原价
	Overseas             *bool      `json:"overseas" form:"overseas" gorm:"column:overseas;comment:海外直邮;"`                                                   //海外直邮
	Paixu                *int       `json:"paixu" form:"paixu" gorm:"column:paixu;comment:排序;size:19;"`                                                      //排序
	Persion              *int       `json:"persion" form:"persion" gorm:"column:persion;comment:用户;size:19;"`                                                //用户
	Pic                  string     `json:"pic" form:"pic" gorm:"column:pic;comment:图片;size:100;"`                                                           //图片
	Pingtuan             *bool      `json:"pingtuan" form:"pingtuan" gorm:"column:pingtuan;comment:允许拼团;"`                                                   //允许拼团
	PingtuanPrice        *float64   `json:"pingtuanPrice" form:"pingtuanPrice" gorm:"column:pingtuan_price;comment:拼团价格;size:10;"`                           //拼团价格
	PropertyIds          string     `json:"propertyIds" form:"propertyIds" gorm:"column:property_ids;comment:属性id;size:1000;"`                               //属性id
	RecommendStatus      *int       `json:"recommendStatus" form:"recommendStatus" gorm:"column:recommend_status;comment:推荐状态;size:19;"`                     //推荐状态
	SeckillBuyNumber     *int       `json:"seckillBuyNumber" form:"seckillBuyNumber" gorm:"column:seckill_buy_number;comment:秒杀最低购买数量;size:19;"`             //秒杀最低购买数量
	ShopId               *int       `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:商店id;size:19;"`                                                //商店id
	Status               *int       `json:"status" form:"status" gorm:"column:status;comment:商品状态;size:19;"`                                                 //商品状态
	Stores               *int       `json:"stores" form:"stores" gorm:"column:stores;comment:库存;size:19;"`                                                   //库存
	Stores0Unsale        *bool      `json:"stores0Unsale" form:"stores0Unsale" gorm:"column:stores0_unsale;comment:自动下架;"`                                   //自动下架
	SellBeginTime        *time.Time `json:"sellBeginTime" form:"sellBeginTime" gorm:"column:sell_begin_time;comment:开始售卖时间;"`                                //开始售卖时间
	SellEndTime          *time.Time `json:"sellEndTime" form:"sellEndTime" gorm:"column:sell_end_time;comment:结束售卖时间;"`                                      //结束售卖时间
	Weight               *float64   `json:"weight" form:"weight" gorm:"column:weight;comment:重量，kg;size:10;"`                                                //重量，kg
	Type                 *int       `json:"type" form:"type" gorm:"column:type;comment:类型;size:19;"`                                                         //类型
	Unit                 string     `json:"unit" form:"unit" gorm:"column:unit;comment:单位;size:100;"`                                                        //单位
}

// TableName 商品表 BeeShopGoods自定义表名 bee_shop_goods
func (BeeShopGoods) TableName() string {
	return "bee_shop_goods"
}
