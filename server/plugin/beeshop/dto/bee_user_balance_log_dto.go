package dto

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/bee"
	"github.com/shopspring/decimal"
)

// 用户余额日志 Dto
type BeeUserBalanceLogDto struct {
	bee.BeeUserBalanceLog
	ShopName string          `json:"shopName" gorm:"column:shopName"`
	Amount   decimal.Decimal `json:"amount" gorm:"column:amount"`
}
