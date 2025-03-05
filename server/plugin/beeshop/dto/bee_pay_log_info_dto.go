package dto

import "github.com/flipped-aurora/gin-vue-admin/server/model/bee"

type BeePayLogInfoDto struct {
	bee.BeePayLog
	ShopName string `json:"shopName" gorm:"shopName"`
}
