package enum

type BeeLogisticsType int32

var (
	BeeLogisticsTypeMap = map[BeeLogisticsType]string{
		BeeLogisticsTypeNum:    "按件",
		BeeLogisticsTypeWeight: "按重量",
		BeeLogisticsTypeAmount: "按金额",
	}
	BeeLogisticsUnitMap = map[BeeLogisticsType]string{
		BeeLogisticsTypeNum:    "件",
		BeeLogisticsTypeWeight: "KG",
		BeeLogisticsTypeAmount: "元",
	}
)

const (
	BeeLogisticsTypeNum    BeeLogisticsType = 0
	BeeLogisticsTypeAmount BeeLogisticsType = 3
	BeeLogisticsTypeWeight BeeLogisticsType = 1
)
