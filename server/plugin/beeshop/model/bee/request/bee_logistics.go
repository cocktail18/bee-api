package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BeeLogisticsSearch struct {
	request.PageInfo
	Id *int `json:"id" form:"id" `
}
