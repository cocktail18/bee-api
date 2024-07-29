package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BeeShopGoodsAdditionSearch struct {
	Id      *int `json:"id" form:"id" `
	GoodsId *int `json:"goodsId" form:"goodsId" `
	Pid     *int `json:"pid" form:"pid" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
