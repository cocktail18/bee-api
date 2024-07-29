package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeeOrderLogisticsSearch struct{
    Uid  *int `json:"uid" form:"uid" `
    OrderId  *int `json:"orderId" form:"orderId" `
    request.PageInfo
}
