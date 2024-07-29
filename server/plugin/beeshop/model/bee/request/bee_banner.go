package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeeBannerSearch struct{
    ShopId  *int `json:"shopId" form:"shopId" `
    Status  *int `json:"status" form:"status" `
    Type  string `json:"type" form:"type" `
    request.PageInfo
    Sort  string `json:"sort" form:"sort"`
    Order string `json:"order" form:"order"`
}
