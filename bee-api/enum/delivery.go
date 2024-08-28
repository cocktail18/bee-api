package enum

type DeliveryType int
type DeliveryCancelReason int

const (
	DeliveryTypeNone DeliveryType = 0
	DeliveryTypeDada DeliveryType = 1
)

var DeliveryTypeMap = map[DeliveryType]string{
	DeliveryTypeDada: "达达",
}

var DeliveryStrMap = map[string]DeliveryType{
	"dada": DeliveryTypeDada,
}

var DeliveryCancelReasonMap = map[DeliveryCancelReason]string{
	1:     "没有配送员接单",
	2:     "配送员没来取货",
	3:     "配送员态度太差",
	4:     "顾客取消订单",
	5:     "订单填写错误",
	34:    "配送员让我取消此单",
	35:    "配送员不愿上门取货",
	36:    "我不需要配送了",
	37:    "配送员以各种理由表示无法完成订单",
	10000: "其他",
}
