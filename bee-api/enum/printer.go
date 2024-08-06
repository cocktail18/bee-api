package enum

type PrinterBrand int32
type PrinterCondition int32

const (
	// PrinterBrandDaQu 大趋打印机 https://www.trendit.cn/pages_9/
	PrinterBrandDaQu PrinterBrand = 1

	PrinterConditionAfterPaySuccess PrinterCondition = 1 // 支付后打印
	PrinterConditionAfterShip       PrinterCondition = 2 // 发货后打印
)

var PrinterBrandMap = map[PrinterBrand]string{
	PrinterBrandDaQu: "大趋",
}

var PrinterConditionMap = map[PrinterCondition]string{
	PrinterConditionAfterPaySuccess: "支付后打印",
	//PrinterConditionAfterShip:       "发货后打印",
}
