package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BeeShopGoodsPropSearch struct {
	Id         *int   `json:"id" form:"id" `
	PropertyId *int   `json:"propertyId" form:"propertyId" `
	Name       string `json:"name" form:"name" `
	request.PageInfo

	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
