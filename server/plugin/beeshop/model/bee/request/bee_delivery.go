package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BeeDeliverySearch struct {
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}

type BeeDeliveryBindYunlabaShopReq struct {
	Source string `json:"source" form:"source"`
	ShopId int    `json:"shopId" form:"shopId"`
	State  string `json:"state" form:"state"`
}
