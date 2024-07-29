// 自动生成模板BeeOrder
package bee

import (
	"github.com/shopspring/decimal"
	"time"
)

// 用户订单 结构体  BeeOrder
type BeeOrder struct {
	Id                *int            `json:"id" form:"id" gorm:"primarykey;column:id;comment:;size:20;"`                    //id字段
	UserId            *int            `json:"userId" form:"userId" gorm:"column:user_id;comment:商店用户ID;size:19;"`            //商店用户ID
	IsDeleted         *bool           `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:已删除;"`              //已删除
	DateAdd           *time.Time      `json:"dateAdd" form:"dateAdd" gorm:"column:date_add;comment:创建时间;"`                   //创建时间
	DateUpdate        *time.Time      `json:"dateUpdate" form:"dateUpdate" gorm:"column:date_update;comment:更新时间;"`          //更新时间
	DateDelete        *time.Time      `json:"dateDelete" form:"dateDelete" gorm:"column:date_delete;comment:删除时间;"`          //删除时间
	Amount            decimal.Decimal `json:"amount" form:"amount" gorm:"column:amount;comment:商品金额;size:10;"`               //商品金额
	AmountCard        decimal.Decimal `json:"amountCard" form:"amountCard" gorm:"column:amount_card;comment:优惠卡抵扣;size:10;"` //优惠卡抵扣
	AmountBalance     decimal.Decimal `gorm:"column:amount_balance;type:decimal(10,2);comment:余额抵扣" json:"amountBalance"`
	AmountCoupons     decimal.Decimal `json:"amountCoupons" form:"amountCoupons" gorm:"column:amount_coupons;comment:优惠券抵扣;size:10;"`               //优惠券抵扣
	AmountLogistics   decimal.Decimal `json:"amountLogistics" form:"amountLogistics" gorm:"column:amount_logistics;comment:运费;size:10;"`            //运费
	AmountReal        decimal.Decimal `json:"amountReal" form:"amountReal" gorm:"column:amount_real;comment:付款金额;size:10;"`                         //付款金额
	AmountRefundTotal decimal.Decimal `json:"amountRefundTotal" form:"amountRefundTotal" gorm:"column:amount_refund_total;comment:退款总金额;size:10;"`  //退款总金额
	AmountTax         decimal.Decimal `json:"amountTax" form:"amountTax" gorm:"column:amount_tax;comment:总税金额;size:10;"`                            //总税金额
	AmountTaxGst      decimal.Decimal `json:"amountTaxGst" form:"amountTaxGst" gorm:"column:amount_tax_gst;comment:消费税;size:10;"`                   //消费税
	AmountTaxService  decimal.Decimal `json:"amountTaxService" form:"amountTaxService" gorm:"column:amount_tax_service;comment:税服务费用;size:10;"`     //税服务费用
	AutoDeliverStatus *int            `json:"autoDeliverStatus" form:"autoDeliverStatus" gorm:"column:auto_deliver_status;comment:自动发货状态;size:19;"` //自动发货状态
	DateClose         *time.Time      `json:"dateClose" form:"dateClose" gorm:"column:date_close;comment:关闭订单时间;size:100;"`                         //关闭订单时间
	DatePay           *time.Time      `json:"datePay" form:"datePay" gorm:"column:date_pay;comment:支付订单时间;size:3;"`                                 //支付订单时间
	GoodsNumber       *int            `json:"goodsNumber" form:"goodsNumber" gorm:"column:goods_number;comment:商品总数量;size:19;"`                     //商品总数量
	HasRefund         *bool           `json:"hasRefund" form:"hasRefund" gorm:"column:has_refund;comment:是否退款;"`                                    //是否退款
	HxNumber          string          `json:"hxNumber" form:"hxNumber" gorm:"column:hx_number;comment:核销码;size:100;"`                               //核销码
	Ip                string          `json:"ip" form:"ip" gorm:"column:ip;comment:下单ip;size:100;"`                                                 //下单ip
	IsCanHx           *bool           `json:"isCanHx" form:"isCanHx" gorm:"column:is_can_hx;comment:是否可以核销;"`                                       //是否可以核销
	IsDelUser         *bool           `json:"isDelUser" form:"isDelUser" gorm:"column:is_del_user;comment:用户删除;"`                                   //用户删除
	IsEnd             *bool           `json:"isEnd" form:"isEnd" gorm:"column:is_end;comment:订单已经结束;"`                                              //订单已经结束
	IsHasBenefit      *bool           `json:"isHasBenefit" form:"isHasBenefit" gorm:"column:is_has_benefit;comment:是否有收益;"`                         //是否有收益
	IsNeedLogistics   *bool           `json:"isNeedLogistics" form:"isNeedLogistics" gorm:"column:is_need_logistics;comment:需要配送;"`                 //需要配送
	IsPay             *bool           `json:"isPay" form:"isPay" gorm:"column:is_pay;comment:是否已经支付;"`                                              //是否已经支付
	IsScoreOrder      *bool           `json:"isScoreOrder" form:"isScoreOrder" gorm:"column:is_score_order;comment:是否有积分;"`                         //是否有积分
	IsSuccessPingtuan *bool           `json:"isSuccessPingtuan" form:"isSuccessPingtuan" gorm:"column:is_success_pingtuan;comment:是否拼团成功;"`         //是否拼团成功
	OrderNumber       string          `json:"orderNumber" form:"orderNumber" gorm:"column:order_number;comment:订单号;size:100;"`                      //订单号
	OrderType         *int            `json:"orderType" form:"orderType" gorm:"column:order_type;comment:订单类型;size:19;"`                            //订单类型
	Pid               *int            `json:"pid" form:"pid" gorm:"column:pid;comment:父订单号;size:19;"`                                               //父订单号
	Qudanhao          string          `json:"qudanhao" form:"qudanhao" gorm:"column:qudanhao;comment:取单号;size:100;"`                                //取单号
	RefundStatus      *int            `json:"refundStatus" form:"refundStatus" gorm:"column:refund_status;comment:退款状态;size:19;"`                   //退款状态
	Remark            string          `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:100;"`                                       //备注
	Score             *int            `json:"score" form:"score" gorm:"column:score;comment:获得积分;size:19;"`                                         //获得积分
	ScoreDeduction    *int            `json:"scoreDeduction" form:"scoreDeduction" gorm:"column:score_deduction;comment:扣除积分;size:19;"`             //扣除积分
	ShopId            *int            `json:"shopId" form:"shopId" gorm:"column:shop_id;comment:商店id;size:19;"`                                     //商店id
	ShopIdZt          *int            `json:"shopIdZt" form:"shopIdZt" gorm:"column:shop_id_zt;comment:自提商店id;size:19;"`                            //自提商店id
	ShopNameZt        string          `json:"shopNameZt" form:"shopNameZt" gorm:"column:shop_name_zt;comment:自提商店名称;size:100;"`                     //自提商店名称
	Status            *int            `json:"status" form:"status" gorm:"column:status;comment:订单状态，1未发货;size:19;"`                                 //订单状态，1未发货
	Trips             *int            `json:"trips" form:"trips" gorm:"column:trips;comment:小费金额;size:19;"`                                         //小费金额
	Type              *int            `json:"type" form:"type" gorm:"column:type;comment:订单类型;size:19;"`                                            //订单类型
	Uid               *int            `json:"uid" form:"uid" gorm:"column:uid;comment:用户id;size:19;"`                                               //用户id
	ExtJsonStr        string          `gorm:"column:ext_json_str;type:varchar(1000);comment:扩展信息" json:"extJsonStr"`
}

// TableName 用户订单 BeeOrder自定义表名 bee_order
func (BeeOrder) TableName() string {
	return "bee_order"
}
