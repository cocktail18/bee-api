package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BeeOrderPeisongLogSearch struct {
	OrderId        *int   `json:"orderId" form:"orderId" `
	PeisongOrderNo string `json:"peisongOrderNo" form:"peisongOrderNo" `
	request.PageInfo
}
