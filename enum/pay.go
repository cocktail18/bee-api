package enum

// PayLogStatus 0 未支付；1 已支付；2 失败
type PayLogStatus int32
type PayGate string

var PayLogStatusMap = map[PayLogStatus]string{
	PayLogStatusUnPaid: "未支付",
	PayLogStatusPaid:   "已支付",
	PayLogStatusFail:   "失败",
}

var PayGateMap = map[PayGate]string{
	PayGateWXAPP: "微信小程序",
}

const (
	PayLogStatusUnPaid PayLogStatus = 0
	PayLogStatusPaid   PayLogStatus = 1
	PayLogStatusFail   PayLogStatus = 2

	PayGateWXAPP = "WXAPP"
)
