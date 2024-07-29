package enum

type BeeGoodsFeeType int32

var (
	BeeGoodsFeeTypeMap = map[BeeGoodsFeeType]string{
		BeeGoodsFeeTypeNum:    "按件",
		BeeGoodsFeeTypeWeight: "按重量",
		BeeGoodsFeeTypeAmount: "按金额",
	}
)

const (
	BeeGoodsFeeTypeNum    BeeGoodsFeeType = 1
	BeeGoodsFeeTypeWeight BeeGoodsFeeType = 2
	BeeGoodsFeeTypeAmount BeeGoodsFeeType = 3
)
