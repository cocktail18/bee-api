package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeeShopGoodsContentSearch struct{
    GoodsId  *int `json:"goodsId" form:"goodsId" `
    request.PageInfo
}
