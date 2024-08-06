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
	Amount            decimal.Decimal        `gorm:"column:amount;type:decimal(10,2);comment:商品金额" json:"amount"` // 总价格
	AmountCard        decimal.Decimal        `gorm:"column:amount_card;type:decimal(10,2);comment:优惠卡抵扣" json:"amountCard"`
	AmountCoupons     decimal.Decimal        `gorm:"column:amount_coupons;type:decimal(10,2);comment:优惠券抵扣" json:"amountCoupons"`
	AmountLogistics   decimal.Decimal        `gorm:"column:amount_logistics;type:decimal(10,2);comment:运费" json:"amountLogistics"`
	AmountBalance     decimal.Decimal        `gorm:"column:amount_balance;type:decimal(10,2);comment:余额抵扣" json:"amountBalance"`
	AmountReal        decimal.Decimal        `gorm:"column:amount_real;type:decimal(10,2);comment:付款金额" json:"amountReal"` // 除去优惠、余额抵扣，实际支付金额
	AmountRefundTotal decimal.Decimal        `gorm:"column:amount_refund_total;type:decimal(10,2);comment:退款总金额" json:"amountRefundTotal"`
	AmountTax         decimal.Decimal        `gorm:"column:amount_tax;type:decimal(10,2);comment:总税金额" json:"amountTax"`
	AmountTaxGst      decimal.Decimal        `gorm:"column:amount_tax_gst;type:decimal(10,2);comment:消费税" json:"amountTaxGst"`
	AmountTaxService  decimal.Decimal        `gorm:"column:amount_tax_service;type:decimal(10,2);comment:税服务费用" json:"amountTaxService"`
	AutoDeliverStatus enum.AutoDeliverStatus `gorm:"column:auto_deliver_status;type:bigint(11);comment:自动发货状态" json:"autoDeliverStatus"`
	DateClose         common.JsonTime        `gorm:"column:date_close;type:datetime(3);comment:关闭订单时间" json:"dateClose"`
	DatePay           common.JsonTime        `gorm:"column:date_pay;type:datetime(3);comment:支付订单时间" json:"datePay"`
	GoodsNumber       int64                  `gorm:"column:goods_number;type:bigint(11);comment:商品总数量" json:"goodsNumber"`
	HasRefund         bool                   `gorm:"column:has_refund;type:tinyint(1);comment:是否退款" json:"hasRefund"`
	HxNumber          string                 `gorm:"uniqueIndex;column:hx_number;type:varchar(100);comment:核销码" json:"hxNumber"`
	Ip                string                 `gorm:"column:ip;type:varchar(100);comment:下单ip" json:"ip"`
	IsCanHx           bool                   `gorm:"column:is_can_hx;type:tinyint(1);comment:是否可以核销" json:"isCanHx"`
	IsDelUser         bool                   `gorm:"column:is_del_user;type:tinyint(1);comment:用户删除" json:"isDelUser"`
	IsEnd             bool                   `gorm:"column:is_end;type:tinyint(1);comment:订单已经结束" json:"isEnd"`
	IsHasBenefit      bool                   `gorm:"column:is_has_benefit;type:tinyint(1);comment:是否有收益" json:"isHasBenefit"`
	IsNeedLogistics   bool                   `gorm:"column:is_need_logistics;type:tinyint(1);comment:需要配送" json:"isNeedLogistics"`
	IsPay             bool                   `gorm:"column:is_pay;type:tinyint(1);comment:是否已经支付" json:"isPay"`
	IsScoreOrder      bool                   `gorm:"column:is_score_order;type:tinyint(1);comment:是否有积分" json:"isScoreOrder"`
	IsSuccessPingtuan bool                   `gorm:"column:is_success_pingtuan;type:tinyint(1);comment:是否拼团成功" json:"isSuccessPingtuan"`
	OrderNumber       string                 `gorm:"uniqueIndex;column:order_number;type:varchar(100);comment:订单号" json:"orderNumber"`
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
	ExtJsonStr        string                 `gorm:"column:ext_json_str;type:varchar(1000);comment:扩展信息" json:"extJsonStr"`
}

func (m *BeeOrder) TableName() string {
	return "bee_order"
}

// BeeOrderGoods 订单商品
type BeeOrderGoods struct {
	common.BaseModel
	AfterSale        string                 `gorm:"column:after_sale;type:varchar(100);comment:售后类型" json:"afterSale"`
	HadPayAmount     decimal.Decimal        `gorm:"column:had_pay_amount;type:decimal(10,2);comment:已支付金额" json:"hadPayAmount"`
	Amount           decimal.Decimal        `gorm:"column:amount;type:decimal(10,2);comment:总金额" json:"amount"`
	AmountCoupon     decimal.Decimal        `gorm:"column:amount_coupon;type:decimal(10,2);comment:优惠券金额" json:"amountCoupon"`
	AmountSingle     decimal.Decimal        `gorm:"column:amount_single;type:decimal(10,2);comment:单价" json:"amountSingle"`
	AmountSingleBase decimal.Decimal        `gorm:"column:amount_single_base;type:decimal(10,2);comment:基本单价" json:"amountSingleBase"`
	BuyRewardEnd     bool                   `gorm:"column:buy_reward_end;type:tinyint(1);comment:奖励发放结束" json:"buyRewardEnd"`
	CategoryId       int64                  `gorm:"column:category_id;type:bigint(20);comment:分类" json:"categoryId"`
	CyTableStatus    enum.CyTableStatus     `gorm:"column:cy_table_status;type:bigint(11);comment:餐饮桌子状态" json:"cyTableStatus"`
	FxType           enum.FxType            `gorm:"column:fx_type;type:bigint(11);comment:返现类型" json:"fxType"`
	GoodsId          int64                  `gorm:"column:goods_id;type:bigint(11);comment:商品id" json:"goodsId"`
	GoodsName        string                 `gorm:"column:goods_name;type:varchar(100);comment:商品名" json:"goodsName"`
	GoodsSubName     string                 `gorm:"column:goods_name;type:varchar(100);comment:商品名称副标题" json:"goodsSubName"`
	IsScoreOrder     bool                   `gorm:"column:is_score_order;type:tinyint(1);comment:是否积分订单" json:"isScoreOrder"`
	Number           int64                  `gorm:"column:number;type:bigint(11);comment:数量" json:"number"`
	NumberNoFahuo    int64                  `gorm:"column:number_no_fahuo;type:bigint(11);comment:未发货数量" json:"numberNoFahuo"`
	OrderId          int64                  `gorm:"column:order_id;type:bigint(11);comment:订单id" json:"orderId"`
	Persion          int64                  `gorm:"column:persion;type:bigint(11);comment:人数" json:"persion"`
	Pic              string                 `gorm:"column:pic;type:varchar(100);comment:商品图" json:"pic"`
	PriceId          int64                  `gorm:"column:price_id;type:bigint(11);comment:价格id" json:"priceId"`
	Property         string                 `gorm:"column:property;type:varchar(100);comment:属性" json:"property"`
	PropertyChildIds string                 `gorm:"column:property_child_ids;type:varchar(100);comment:属性id" json:"propertyChildIds"`
	Purchase         bool                   `gorm:"column:purchase;type:tinyint(1);comment:是否已支付" json:"purchase"`
	RefundStatus     enum.OrderRefundStatus `gorm:"column:refund_status;type:bigint(11);comment:退款状态" json:"refundStatus"`
	Score            decimal.Decimal        `gorm:"column:score;type:bigint(11);comment:积分" json:"score"`
	ShopId           int64                  `gorm:"column:shop_id;type:bigint(11);comment:商店id" json:"shopId"`
	Status           enum.OrderGoodsStatus  `gorm:"column:status;type:bigint(11);comment:状态" json:"status"`
	Type             int64                  `gorm:"column:type;type:bigint(11);comment:类型" json:"type"`
	Uid              int64                  `gorm:"column:uid;type:bigint(11);comment:用户id" json:"uid"`
	Unit             string                 `gorm:"column:unit;type:varchar(100);comment:单位" json:"unit"`
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
	OrderId    int64                `gorm:"column:order_id;type:bigint(100);comment:订单id" json:"orderId"`
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

type BeeOrderQuDanHao struct {
	common.BaseModel `json:"common.BaseModel"`
	ShopId           int64  `gorm:"column:shop_id;type:bigint(20);comment:商店id" json:"shopId" json:"shopId,omitempty"`
	Type             string `gorm:"column:type;type:varchar(100);comment:类型" json:"type" json:"type,omitempty"`
	Num              int64  `gorm:"column:num;type:bigint(11);comment:数量" json:"num" json:"num,omitempty"`
}

func (m *BeeOrderQuDanHao) TableName() string {
	return "bee_order_qu_dan_hao"
}

// BeeOrderPayLog 支付流水，一个订单可以分多次支付
type BeeOrderPayLog struct {
	common.BaseModel
	OrderId int64             `gorm:"column:order_id;type:bigint(100);comment:订单id" json:"orderId"`
	PayGate enum.PayGate      `gorm:"column:pay_gate;type:varchar(100);comment:支付网关" json:"payGate"`
	Amount  decimal.Decimal   `gorm:"column:amount;type:decimal(10,2);comment:金额" json:"amount"`
	Status  enum.PayLogStatus `gorm:"column:status;type:bigint(11);comment:状态" json:"status"`
	Remark  string            `gorm:"column:remark;type:varchar(100);comment:备注" json:"remark"`
	ThirdId string            `gorm:"column:third_id;type:varchar(100);comment:第三方支付id" json:"thirdId"`
	Extra   string            `gorm:"column:extra;type:varchar(2000);comment:额外信息" json:"extra"`
}

func (m *BeeOrderPayLog) TableName() string {
	return "bee_order_pay_log"
}

type BeeOrderPeisong struct {
	common.BaseModel
	OrderId        int64                   `gorm:"column:order_id;type:bigint(100);comment:订单id" json:"orderId"`
	PeisongOrderId string                  `gorm:"column:peisong_order_id;type:varchar(100);comment:配送订单id" json:"peisongOrderId"`
	Status         enum.OrderPaisongStatus `gorm:"column:status;type:bigint(11);comment:状态" json:"status"`
}

func (m *BeeOrderPeisong) TableName() string {
	return "bee_order_peisong"
}

type BeeOrderLog struct {
	common.BaseModel
	OrderId int64             `gorm:"column:order_id;type:bigint(100);comment:订单id" json:"orderId"`
	Type    enum.OrderLogType `gorm:"column:type;type:bigint(11);comment:类型" json:"type"`
	TypeStr string            `gorm:"-" json:"typeStr"`
	Remark  string            `gorm:"column:remark;type:varchar(100);comment:备注" json:"remark"`
}

func (b *BeeOrderLog) TableName() string {
	return "bee_order_log"
}

func (b *BeeOrderLog) FillData() {
	b.TypeStr = enum.OrderLogTypeMap[b.Type]
}

type BeeOrderCoupon struct {
	common.BaseModel
	OrderId  int64 `gorm:"column:order_id;type:bigint(100);comment:订单id" json:"orderId"`
	CouponId int64 `gorm:"column:coupon_id;type:bigint(100);comment:优惠券id" json:"couponId"`
	Uid      int64 `gorm:"column:uid;type:bigint(20);comment:用户id" json:"uid"`
}

func (m *BeeOrderCoupon) TableName() string {
	return "bee_order_coupon"
}

type BeeOrderPrintLog struct {
	common.BaseModel
	OrderId     int64                   `gorm:"column:order_id;type:bigint(100);comment:订单id" json:"orderId"`
	Condition   enum.PrinterCondition   `gorm:"column:condition;type:int(11);comment:打印条件" json:"condition"`
	Status      enum.OrderPrinterStatus `gorm:"column:status;type:int(11);comment:状态" json:"status"`
	ErrMsg      string                  `gorm:"column:err_msg;type:varchar(100);comment:错误信息" json:"errMsg"`
	TryTimes    int                     `gorm:"column:try_times;type:int(11);comment:尝试次数" json:"tryTimes"`
	LastTryUnix int64                   `gorm:"column:last_try_unix;type:bigint(20);comment:最后尝试时间" json:"lastTryUnix"`
}

func (m *BeeOrderPrintLog) TableName() string {
	return "bee_order_print_log"
}
