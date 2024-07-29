package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	
)

type BeeOrderLogSearch struct{
    OrderId  *int `json:"orderId" form:"orderId" `
    Type  *int `json:"type" form:"type" `
    request.PageInfo
}
