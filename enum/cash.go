package enum

// BeeCashLogBehavior 0 收入 1 支出
type BeeCashLogBehavior int32

// BeeCashLogType 交易类型 0 充值 11 提现申请 12 提现失败 1 提现成功 2 支付订单 3 退款 4支付预约报名费 5 知识付费 6分销返佣
// 7 分享商品奖励； 8 优惠买单； 100 押金；101 押金退还； 110 购买会员； 120 货款收入； 130 分润收入； 140 积分兑换
type BeeCashLogType int32

var BeeCashLogTypeMap = map[BeeCashLogType]string{
	BeeCashLogTypeRecharge:    "充值",
	BeeCashLogTypePayOrder:    "支付订单",
	BeeCashLogTypePayRefund:   "退款",
	BeeCashLogTypePayExchange: "积分兑换",
}

const (
	BeeCashLogBehaviorIncome  BeeCashLogBehavior = 0
	BeeCashLogBehaviorConsume BeeCashLogBehavior = 1

	BeeCashLogTypeRecharge    BeeCashLogType = 0
	BeeCashLogTypePayOrder    BeeCashLogType = 2
	BeeCashLogTypePayRefund   BeeCashLogType = 3
	BeeCashLogTypePayExchange BeeCashLogType = 140
)
