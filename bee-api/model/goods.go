package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/shopspring/decimal"
)

// 商品分类
type BeeShopGoodsCategory struct {
	common.BaseModel
	IsUse  bool   `gorm:"column:is_use;type:tinyint(1);comment:是否启用" json:"isUse"`
	Level  int64  `gorm:"column:level;type:bigint(11);comment:等级" json:"level"`
	Name   string `gorm:"column:name;type:varchar(100);comment:名称" json:"name"`
	Paixu  int64  `gorm:"column:paixu;type:bigint(11);comment:排序" json:"paixu"`
	Pid    int64  `gorm:"column:pid;type:bigint(11);comment:父id" json:"pid"`
	ShopId int64  `gorm:"column:shop_id;type:bigint(11);comment:店铺id" json:"shopId"`
}

func (m *BeeShopGoodsCategory) TableName() string {
	return "bee_shop_goods_category"
}

// 商品
type BeeShopGoods struct {
	common.BaseModel
	BarCode              string           `gorm:"column:bar_code;type:varchar(100);comment:条码" json:"barCode"`
	AfterSale            string           `gorm:"column:after_sale;type:varchar(100);comment:售后说明,1仅退款，2退款退货，3换货" json:"afterSale"`
	CategoryId           int64            `gorm:"column:category_id;type:bigint(11);comment:分类id" json:"categoryId"`
	Commission           decimal.Decimal  `gorm:"column:commission;type:decimal(10,2);comment:积分奖励" json:"commission"`
	CommissionSettleType int64            `gorm:"column:commission_settle_type;type:bigint(11);comment:积分奖励结算类型" json:"commissionSettleType"`
	CommissionType       int64            `gorm:"column:commission_type;type:bigint(11);comment:积分奖励类型" json:"commissionType"`
	CommissionUserType   int64            `gorm:"column:commission_user_type;type:bigint(11);comment:积分奖励用户类型" json:"commissionUserType"`
	HasAddition          bool             `gorm:"column:has_addition;type:tinyint(1);comment:有附加" json:"hasAddition"`
	HasTourJourney       bool             `gorm:"column:has_tour_journey;type:tinyint(1);comment:有团购" json:"hasTourJourney"`
	Hidden               int64            `gorm:"column:hidden;type:bigint(11);comment:隐藏" json:"hidden"`
	Kanjia               bool             `gorm:"column:kanjia;type:tinyint(1);comment:允许砍价" json:"kanjia"`
	KanjiaPrice          decimal.Decimal  `gorm:"column:kanjia_price;type:decimal(10,2);comment:砍价价格" json:"kanjiaPrice"`
	VipPrice             decimal.Decimal  `gorm:"column:vip_price;type:decimal(10,2);comment:会员价格" json:"vipPrice"`
	Limitation           bool             `gorm:"column:limitation;type:tinyint(1);comment:限制" json:"limitation"`
	LogisticsId          int64            `gorm:"column:logistics_id;type:bigint(11);comment:物流id" json:"logisticsId"`
	MaxCoupons           int64            `gorm:"column:max_coupons;type:bigint(11);comment:最大使用优惠券数量" json:"maxCoupons"`
	Miaosha              bool             `gorm:"column:miaosha;type:tinyint(1);comment:秒杀" json:"miaosha"`
	MinBuyNumber         int64            `gorm:"column:min_buy_number;type:bigint(11);comment:最低购买数量" json:"minBuyNumber"`
	MinPrice             decimal.Decimal  `gorm:"column:min_price;type:decimal(10,2);comment:最低价格" json:"minPrice"`
	MinScore             decimal.Decimal  `gorm:"column:min_score;type:decimal(10,2);comment:需要积分数量" json:"minScore"`
	Name                 string           `gorm:"column:name;type:varchar(100);comment:商品名字" json:"name"`
	NumberFav            int64            `gorm:"column:number_fav;type:bigint(11);comment:喜欢数量" json:"numberFav"`
	NumberGoodReputation decimal.Decimal  `gorm:"column:number_good_reputation;type:decimal(10,2);comment:商品评分" json:"numberGoodReputation"`
	NumberOrders         int64            `gorm:"column:number_orders;type:bigint(11);comment:订单数量" json:"numberOrders"`
	NumberReputation     int64            `gorm:"column:number_reputation;type:bigint(11);comment:评分数量" json:"numberReputation"`
	NumberSells          int64            `gorm:"column:number_sells;type:bigint(11);comment:销售数量" json:"numberSells"`
	OriginalPrice        decimal.Decimal  `gorm:"column:original_price;type:decimal(10,2);comment:原价" json:"originalPrice"`
	Overseas             bool             `gorm:"column:overseas;type:tinyint(1);comment:海外直邮" json:"overseas"`
	Paixu                int64            `gorm:"column:paixu;type:bigint(11);comment:排序" json:"paixu"`
	Persion              int64            `gorm:"column:persion;type:bigint(11);comment:用户" json:"persion"`
	Pic                  string           `gorm:"column:pic;type:varchar(100);comment:图片" json:"pic"`
	Pingtuan             bool             `gorm:"column:pingtuan;type:tinyint(1);comment:允许拼团" json:"pingtuan"`
	PingtuanPrice        decimal.Decimal  `gorm:"column:pingtuan_price;type:decimal(10,2);comment:拼团价格" json:"pingtuanPrice"`
	PropertyIds          string           `gorm:"column:property_ids;type:varchar(1000);comment:属性id" json:"propertyIds"`
	RecommendStatus      int64            `gorm:"column:recommend_status;type:bigint(11);comment:推荐状态" json:"recommendStatus"`
	SeckillBuyNumber     int64            `gorm:"column:seckill_buy_number;type:bigint(11);comment:秒杀最低购买数量" json:"seckillBuyNumber"`
	ShopId               int64            `gorm:"column:shop_id;type:bigint(11);comment:商店id" json:"shopId"`
	Status               enum.GoodsStatus `gorm:"column:status;type:bigint(11);comment:商品状态" json:"status"`
	Stores               int64            `gorm:"column:stores;type:bigint(11);comment:库存" json:"stores"`
	Stores0Unsale        bool             `gorm:"column:stores0_unsale;type:tinyint(1);comment:自动下架" json:"stores0Unsale"`
	SellBeginTime        common.JsonTime  `gorm:"column:sell_begin_time;type:datetime;comment:开始售卖时间;default:null" json:"sellBeginTime"`
	SellEndTime          common.JsonTime  `gorm:"column:sell_end_time;type:datetime;comment:结束售卖时间;default:null" json:"sellEndTime"`
	Weight               decimal.Decimal  `gorm:"column:weight;type:decimal(10,2);default:0.00;comment:重量，kg" json:"weight"`
	Type                 int64            `gorm:"column:type;type:bigint(11);comment:类型" json:"type"`
	Unit                 string           `gorm:"column:unit;type:varchar(100);comment:单位" json:"unit"`
}

func (m *BeeShopGoods) TableName() string {
	return "bee_shop_goods"
}

type BeeShopGoodsImages struct {
	common.BaseModel
	GoodsId int64  `gorm:"column:goods_id;type:bigint(20)" json:"goodsId"`
	Pic     string `gorm:"column:pic;type:varchar(255)" json:"pic"`
}

func (m *BeeShopGoodsImages) TableName() string {
	return "bee_shop_goods_images"
}

type BeeShopGoodsProp struct {
	common.BaseModel
	PropertyId int64  `gorm:"column:property_id;type:bigint(20)" json:"propertyId"` //父属性id
	Name       string `gorm:"column:name;type:varchar(100)" json:"name"`
	Paixu      int64  `gorm:"column:paixu;type:bigint(11)" json:"paixu"`

	ChildsCurGoods []*BeeShopGoodsProp `gorm:"-" json:"childsCurGoods"`
}

func (m *BeeShopGoodsProp) TableName() string {
	return "bee_shop_goods_prop"
}

type BeeShopGoodsContent struct {
	common.BaseModel
	GoodsId int64  `gorm:"column:goods_id;type:bigint(20)" json:"goodsId"`
	Content string `gorm:"column:content;type:longblob" json:"content"`
}

func (m *BeeShopGoodsContent) TableName() string {
	return "bee_shop_goods_content"
}

type BeeShopGoodsAddition struct {
	common.BaseModel
	GoodsId int64                         `gorm:"column:goods_id;type:bigint(20)" json:"goodsId"`
	Name    string                        `gorm:"column:name;type:varchar(100)" json:"name"`
	Require bool                          `gorm:"column:require;type:tinyint(1)" json:"require"` //是否必选
	Type    enum.BeeShopGoodsAdditionType `gorm:"column:type;type:int(11)" json:"type"`
	Pid     int64                         `gorm:"column:pid;type:bigint(20)" json:"pid"`
	Items   []*BeeShopGoodsAddition       `gorm:"-" json:"items"`
	Price   decimal.Decimal               `gorm:"column:price;type:decimal(10,2);default:0.00;comment:价格" json:"price"` //选择之后增加费用
}

func (m *BeeShopGoodsAddition) TableName() string {
	return "bee_shop_goods_addition"
}
