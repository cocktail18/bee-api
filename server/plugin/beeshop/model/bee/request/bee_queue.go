package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BeeQueueSearch struct {
	Id *int `json:"id" form:"id" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
