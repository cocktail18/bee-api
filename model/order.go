package model

import (
	"gitee.com/stuinfer/bee-api/common"
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	"strings"
)

// 订单
type BeeOrder struct {
	common.BaseModel
	Amount            decimal.Decimal        `gorm:"column:amount;type:decimal(10,2);comment:商品金额" json:"amount"`
	AmountCard        decimal.Decimal        `gorm:"column:amount_card;type:decimal(10,2);comment:优惠卡抵扣" json:"amountCard"`
	AmountCoupons     decimal.Decimal        `gorm:"column:amount_coupons;type:decimal(10,2);comment:优惠券抵扣" json:"amountCoupons"`
	AmountLogistics   decimal.Decimal        `gorm:"column:amount_logistics;type:decimal(10,2);comment:运费" json:"amountLogistics"`
	AmountReal        decimal.Decimal        `gorm:"column:amount_real;type:decimal(10,2);comment:付款金额" json:"amountReal"`
	AmountRefundTotal decimal.Decimal        `gorm:"column:amount_refund_total;type:decimal(10,2);comment:退款总金额" json:"amountRefundTotal"`
	AmountTax         decimal.Decimal        `gorm:"column:amount_tax;type:decimal(10,2);comment:总税金额" json:"amountTax"`
	AmountTaxGst      decimal.Decimal        `gorm:"column:amount_tax_gst;type:decimal(10,2);comment:消费税" json:"amountTaxGst"`
	AmountTaxService  decimal.Decimal        `gorm:"column:amount_tax_service;type:decimal(10,2);comment:税服务费用" json:"amountTaxService"`
	AutoDeliverStatus enum.AutoDeliverStatus `gorm:"column:auto_deliver_status;type:bigint(11);comment:自动发货状态" json:"autoDeliverStatus"`
	DateClose         string                 `gorm:"column:date_close;type:varchar(100);comment:关闭订单时间" json:"dateClose"`
	DatePay           string                 `gorm:"column:date_pay;type:varchar(100);comment:支付订单时间" json:"datePay"`
	GoodsNumber       int64                  `gorm:"column:goods_number;type:bigint(11);comment:商品总数量" json:"goodsNumber"`
	HasRefund         bool                   `gorm:"column:has_refund;type:tinyint(1);comment:是否退款" json:"hasRefund"`
	HxNumber          string                 `gorm:"column:hx_number;type:varchar(100);comment:核销码" json:"hxNumber"`
	Ip                string                 `gorm:"column:ip;type:varchar(100);comment:下单ip" json:"ip"`
	IsCanHx           bool                   `gorm:"column:is_can_hx;type:tinyint(1);comment:是否可以核销" json:"isCanHx"`
	IsDelUser         bool                   `gorm:"column:is_del_user;type:tinyint(1);comment:用户删除" json:"isDelUser"`
	IsEnd             bool                   `gorm:"column:is_end;type:tinyint(1);comment:订单已经结束" json:"isEnd"`
	IsHasBenefit      bool                   `gorm:"column:is_has_benefit;type:tinyint(1);comment:是否有收益" json:"isHasBenefit"`
	IsNeedLogistics   bool                   `gorm:"column:is_need_logistics;type:tinyint(1);comment:需要配送" json:"isNeedLogistics"`
	IsPay             bool                   `gorm:"column:is_pay;type:tinyint(1);comment:是否已经支付" json:"isPay"`
	IsScoreOrder      bool                   `gorm:"column:is_score_order;type:tinyint(1);comment:是否有积分" json:"isScoreOrder"`
	IsSuccessPingtuan bool                   `gorm:"column:is_success_pingtuan;type:tinyint(1);comment:是否拼团成功" json:"isSuccessPingtuan"`
	OrderNumber       string                 `gorm:"column:order_number;type:varchar(100);comment:订单号" json:"orderNumber"`
	OrderType         enum.OrderType         `gorm:"column:order_type;type:bigint(11);comment:订单类型" json:"orderType"`
	Pid               int64                  `gorm:"column:pid;type:bigint(20);comment:父订单号" json:"pid"`
	Qudanhao          string                 `gorm:"column:qudanhao;type:varchar(100);comment:取单号" json:"qudanhao"`
	RefundStatus      enum.OrderRefundStatus `gorm:"column:refund_status;type:bigint(11);comment:退款状态" json:"refundStatus"`
	Remark            string                 `gorm:"column:remark;type:varchar(100);comment:备注" json:"remark"`
	Score             int64                  `gorm:"column:score;type:bigint(11);comment:获得积分" json:"score"`
	ScoreDeduction    int64                  `gorm:"column:score_deduction;type:bigint(11);comment:扣除积分" json:"scoreDeduction"`
	ShopId            int64                  `gorm:"column:shop_id;type:bigint(11);comment:商店id" json:"shopId"`
	ShopIdZt          int64                  `gorm:"column:shop_id_zt;type:bigint(11);comment:自提商店id" json:"shopIdZt"`
	ShopNameZt        string                 `gorm:"column:shop_name_zt;type:varchar(100);comment:自提商店名称" json:"shopNameZt"`
	Status            enum.OrderStatus       `gorm:"column:status;type:bigint(11);comment:订单状态，1未发货" json:"status"`
	Trips             decimal.Decimal        `gorm:"column:trips;type:bigint(11);comment:小费金额" json:"trips"`
	Type              enum.OrderType         `gorm:"column:type;type:bigint(11);comment:订单类型" json:"type"`
	Uid               int64                  `gorm:"column:uid;type:bigint(11);comment:用户id" json:"uid"`
}

func (m *BeeOrder) TableName() string {
	return "bee_order"
}

// BeeOrderGoods 订单商品
type BeeOrderGoods struct {
	Id               int64                        `gorm:"column:id;type:bigint(100);primary_key;AUTO_INCREMENT;comment:id" json:"id"`
	AfterSale        string                       `gorm:"column:after_sale;type:varchar(100);comment:售后类型" json:"afterSale"`
	Amount           decimal.Decimal              `gorm:"column:amount;type:decimal(10,2);comment:总金额" json:"amount"`
	AmountCoupon     decimal.Decimal              `gorm:"column:amount_coupon;type:decimal(10,2);comment:优惠券金额" json:"amountCoupon"`
	AmountSingle     decimal.Decimal              `gorm:"column:amount_single;type:decimal(10,2);comment:单价" json:"amountSingle"`
	AmountSingleBase decimal.Decimal              `gorm:"column:amount_single_base;type:decimal(10,2);comment:基本单价" json:"amountSingleBase"`
	BuyRewardEnd     int                          `gorm:"column:buy_reward_end;type:tinyint(1);comment:奖励发放结束" json:"buyRewardEnd"`
	CategoryId       int64                        `gorm:"column:category_id;type:bigint(20);comment:分类" json:"categoryId"`
	CyTableStatus    enum.OrderGoodsCyTableStatus `gorm:"column:cy_table_status;type:bigint(11);comment:餐饮桌子状态" json:"cyTableStatus"`
	DateAdd          string                       `gorm:"column:date_add;type:varchar(100);comment:添加时间" json:"dateAdd"`
	FxType           enum.OrderGoodsFxType        `gorm:"column:fx_type;type:bigint(11);comment:返现类型" json:"fxType"`
	GoodsId          int64                        `gorm:"column:goods_id;type:bigint(11);comment:商品id" json:"goodsId"`
	GoodsName        string                       `gorm:"column:goods_name;type:varchar(100);comment:商品名" json:"goodsName"`
	GoodsSubName     string                       `gorm:"column:goods_name;type:varchar(100);comment:商品名称副标题" json:"goodsSubName"`
	IsScoreOrder     int                          `gorm:"column:is_score_order;type:tinyint(1);comment:是否积分订单" json:"isScoreOrder"`
	Number           int64                        `gorm:"column:number;type:bigint(11);comment:数量" json:"number"`
	NumberNoFahuo    int64                        `gorm:"column:number_no_fahuo;type:bigint(11);comment:未发货数量" json:"numberNoFahuo"`
	OrderId          int64                        `gorm:"column:order_id;type:bigint(11);comment:订单id" json:"orderId"`
	Persion          int64                        `gorm:"column:persion;type:bigint(11);comment:人数" json:"persion"`
	Pic              string                       `gorm:"column:pic;type:varchar(100);comment:商品图" json:"pic"`
	PriceId          int64                        `gorm:"column:price_id;type:bigint(11);comment:价格id" json:"priceId"`
	Property         string                       `gorm:"column:property;type:varchar(100);comment:属性" json:"property"`
	PropertyChildIds string                       `gorm:"column:property_child_ids;type:varchar(100);comment:属性id" json:"propertyChildIds"`
	Purchase         string                       `gorm:"column:purchase;type:varchar(100);comment:是否已支付" json:"purchase"`
	RefundStatus     int64                        `gorm:"column:refund_status;type:bigint(11);comment:退款状态" json:"refundStatus"`
	Score            int64                        `gorm:"column:score;type:bigint(11);comment:积分" json:"score"`
	ShopId           int64                        `gorm:"column:shop_id;type:bigint(11);comment:商店id" json:"shopId"`
	Status           enum.OrderGoodsStatus        `gorm:"column:status;type:bigint(11);comment:状态" json:"status"`
	Type             int64                        `gorm:"column:type;type:bigint(11);comment:类型" json:"type"`
	Uid              int64                        `gorm:"column:uid;type:bigint(11);comment:用户id" json:"uid"`
	Unit             string                       `gorm:"column:unit;type:varchar(100);comment:单位" json:"unit"`
}

func (m *BeeOrderGoods) TableName() string {
	return "bee_order_goods"
}

// BeeOrderLogistics 订单收货地址
type BeeOrderLogistics struct {
	BeeUserAddress
	OrderId int64 `gorm:"column:order_id;type:bigint(11);comment:订单id" json:"orderId"`
}

func (m *BeeOrderLogistics) TableName() string {
	return "bee_order_logistics"
}

type BeeOrderReputation struct {
	common.BaseModel
	Uid        int64                `gorm:"column:uid;type:bigint(20);comment:用户id" json:"uid"`
	OrderId    string               `gorm:"column:order_id;type:varchar(100);comment:订单id" json:"orderId"`
	Reputation enum.OrderReputation `gorm:"column:reputation;type:bigint(11);comment:评价" json:"reputation"`
	Remark     string               `gorm:"column:remark;type:varchar(100);comment:备注" json:"remark"`
	PicsStr    string               `gorm:"column:pics_str;type:varchar(100);comment:图片列表" json:"picsStr"`
	Pics       []string             `gorm:"-" json:"pics"`
}

func (m *BeeOrderReputation) TableName() string {
	return "bee_order_reputation"
}

func (b *BeeOrderReputation) Decode() {
	if b.PicsStr != "" {
		b.Pics = lo.Filter(strings.Split(b.PicsStr, ","), func(item string, index int) bool {
			return item != ""
		})
	}
}
