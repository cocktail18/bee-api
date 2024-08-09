package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BeeShopGoodsSearch struct {
	BarCode            string     `json:"barCode" form:"barCode" `
	CategoryId         *int       `json:"categoryId" form:"categoryId" `
	LogisticsId        *int       `json:"logisticsId" form:"logisticsId" `
	Name               string     `json:"name" form:"name" `
	ShopId             *int       `json:"shopId" form:"shopId" `
	StartDateAdd       *time.Time `json:"startDateAdd" form:"startDateAdd"`
	EndDateAdd         *time.Time `json:"endDateAdd" form:"endDateAdd"`
	StartSellBeginTime *time.Time `json:"startSellBeginTime" form:"startSellBeginTime"`
	EndSellBeginTime   *time.Time `json:"endSellBeginTime" form:"endSellBeginTime"`
	StartSellEndTime   *time.Time `json:"startSellEndTime" form:"startSellEndTime"`
	EndSellEndTime     *time.Time `json:"endSellEndTime" form:"endSellEndTime"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
