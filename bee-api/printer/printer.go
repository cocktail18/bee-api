package printer

import (
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/model"
)

type Printer interface {
	Print(config *model.BeePrinter, voice, content string) error
	AddPrinter(config *model.BeePrinter) error
	DelPrinter(config *model.BeePrinter, codes []string) error
}

var brand2printer = map[enum.PrinterBrand]Printer{
	enum.PrinterBrandDaQu: NewDaQu(),
}

func GetPrinter(config *model.BeePrinter) Printer {
	return brand2printer[config.Brand]
}

func GetPrinterByBrand(brand enum.PrinterBrand) Printer {
	return brand2printer[brand]
}
