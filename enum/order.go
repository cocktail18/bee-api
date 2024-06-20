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
type OrderGoodsFxType int32
type OrderGoodsCyTableStatus int32 // 餐饮桌子状态

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
)
