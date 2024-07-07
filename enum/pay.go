package enum

// PayLogStatus 0 未支付；1 已支付；2 失败
type PayLogStatus int32
type PayGate string
type PayNextActionType int32

var PayLogStatusMap = map[PayLogStatus]string{
	PayLogStatusUnPaid: "未支付",
	PayLogStatusPaid:   "已支付",
	PayLogStatusFail:   "失败",
}

var PayGateMap = map[PayGate]string{
	PayGateWXAPP:   "微信小程序",
	PayGateBalance: "余额",
}

const (
	PayLogStatusUnPaid PayLogStatus = 0
	PayLogStatusPaid   PayLogStatus = 1
	PayLogStatusFail   PayLogStatus = 2

	PayGateWXAPP   PayGate = "WXAPP"
	PayGateBalance PayGate = "余额"

	PayNextActionTypeRecharge  PayNextActionType = -1 // 充值，nextAction 没有值
	PayNextActionTypePayOrder  PayNextActionType = 0  // 支付订单， { type: 0, id: orderId}
	PayNextActionTypePayDirect PayNextActionType = 4  // 优惠买单 {type: 4, uid: wx.getStorageSync('uid'), money: data.money}
)
