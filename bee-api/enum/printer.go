package enum

type PrinterBrand int32
type PrinterCondition int32

const (
	// PrinterBrandDaQu 大趋打印机 https://www.trendit.cn/pages_9/
	PrinterBrandDaQu PrinterBrand = 1
	// PrinterBrandFeiE 飞蛾打印机 https://help.feieyun.com/home/doc/zh;nav=1-1
	PrinterBrandFeiE PrinterBrand = 2

	PrinterConditionAfterPaySuccess PrinterCondition = 1 // 支付后打印
	PrinterConditionAfterShip       PrinterCondition = 2 // 发货后打印
)

var PrinterBrandMap = map[PrinterBrand]string{
	PrinterBrandDaQu: "大趋",
	PrinterBrandFeiE: "飞蛾",
}

var PrinterConditionMap = map[PrinterCondition]string{
	PrinterConditionAfterPaySuccess: "支付后打印",
	//PrinterConditionAfterShip:       "发货后打印",
}
