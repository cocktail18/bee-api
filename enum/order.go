package enum

// OrderStatus https://www.yuque.com/apifm/nu0f75/uwggsm
// -1 已关闭 0 待支付 1 已支付待发货 2 已发货待确认 3 确认收货待评价 4 已评价
type OrderStatus int32
type AutoDeliverStatus int32
type OrderReputation int32 // 评价状态，0 差评 1 中评 2 好评

// OrderType 0 普通订单 1 周期订单 2 扫码点餐订单 3 京东vop订单 4 从区管进货
type OrderType int32
type OrderRefundStatus int32
type OrderGoodsStatus int32
type OrderLogType int32
type OrderPaisongStatus int32

var OrderReputationMap = map[OrderReputation]string{
	OrderReputationNone:   "无",
	OrderReputationBad:    "差评",
	OrderReputationMiddle: "中评",
	OrderReputationGood:   "好评",
}

var OrderTypeMap = map[OrderType]string{
	OrderTypeNormal:    "普通订单",
	OrderTypePeriod:    "周期订单",
	OrderTypeScan:      "扫码点餐订单",
	OrderTypeJdVop:     "京东vop订单",
	OrderTypeReceiving: "从区管进货",
}

var OrderStatusMap = map[OrderStatus]string{
	OrderStatusClose:          "已关闭",
	OrderStatusUnPaid:         "待支付",
	OrderStatusPaid:           "已支付待发货",
	OrderStatusShipped:        "已发货待确认",
	OrderStatusConfirmShipped: "确认收货待评价",
	OrderStatusHadComment:     "已评价",
}

var OrderLogTypeMap = map[OrderLogType]string{
	OrderLogTypePayByOthers:    "代付",
	OrderLogTypeCreate:         "下单",
	OrderLogTypeDeliver:        "发货",
	OrderLogTypeRefund:         "退款",
	OrderLogTypeChangePrice:    "改价",
	OrderLogTypeConfirmDeliver: "确认发货",
	OrderLogTypeReputation:     "评价",
	OrderLogTypeClose:          "关闭",
	OrderLogTypePay:            "支付",
	OrderLogTypeHx:             "核销",
	OrderLogTypeAddGoods:       "添加商品",
	OrderLogTypeSetComplete:    "商家设置完成",
	OrderLogTypeCreateSub:      "创建子订单",
}

const (
	OrderStatusClose          OrderStatus = -1
	OrderStatusUnPaid         OrderStatus = 0
	OrderStatusPaid           OrderStatus = 1
	OrderStatusShipped        OrderStatus = 2 // 已发货待确认
	OrderStatusConfirmShipped OrderStatus = 3 // 确认收货待评价
	OrderStatusHadComment     OrderStatus = 4 // 已评价

	OrderGoodsStatusNone           OrderGoodsStatus = 0 // 无
	OrderGoodsStatusShipped        OrderGoodsStatus = 2 // 已发货待确认
	OrderGoodsStatusConfirmShipped OrderGoodsStatus = 3 // 确认收货待评价
	OrderGoodsStatusHadComment     OrderGoodsStatus = 4 // 已评价

	OrderRefundStatusNone      = 0
	OrderRefundStatusHadRefund = 1

	OrderTypeNormal    = 0
	OrderTypePeriod    = 1
	OrderTypeScan      = 2
	OrderTypeJdVop     = 3
	OrderTypeReceiving = 4

	OrderReputationNone   = -1
	OrderReputationBad    = 0
	OrderReputationMiddle = 1
	OrderReputationGood   = 2

	OrderLogTypeClose          OrderLogType = -1  // 关闭订单
	OrderLogTypeCreate         OrderLogType = 0   // 0 下单
	OrderLogTypePay            OrderLogType = 1   // 1 支付
	OrderLogTypeDeliver        OrderLogType = 2   // 2 发货
	OrderLogTypeConfirmDeliver OrderLogType = 3   // 3 确认收货
	OrderLogTypeReputation     OrderLogType = 4   // 4 评价
	OrderLogTypeChangePrice    OrderLogType = 5   // 5 卖家修改价格
	OrderLogTypeSetComplete    OrderLogType = 6   // 6 商家设置交易完成
	OrderLogTypeRefund         OrderLogType = 7   // 7退款
	OrderLogTypeCreateSub      OrderLogType = 8   // 8 创建子订单
	OrderLogTypeAddGoods       OrderLogType = 100 // 100 增加商品
	OrderLogTypePayByOthers    OrderLogType = 11  // 11 代支付
	OrderLogTypeHx             OrderLogType = 21  // 11 核销

	OrderPaisongStatusNone OrderPaisongStatus = 0
)
