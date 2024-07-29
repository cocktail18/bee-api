package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BeeShopGoodsSkuSearch struct {
	Id          *int     `json:"id" form:"id" `
	GoodsId     *int     `json:"goodsId" form:"goodsId" `
	Code        string   `json:"code" form:"code" `
	StartPrice  *float64 `json:"startPrice" form:"startPrice"`
	EndPrice    *float64 `json:"endPrice" form:"endPrice"`
	StartStores *int     `json:"startStores" form:"startStores"`
	EndStores   *int     `json:"endStores" form:"endStores"`
	StartWeight *float64 `json:"startWeight" form:"startWeight"`
	EndWeight   *float64 `json:"endWeight" form:"endWeight"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
