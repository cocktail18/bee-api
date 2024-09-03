package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BeeUserQueueSearch struct {
	Uid     *int `json:"uid" form:"uid" `
	QueueId *int `json:"queueId" form:"queueId" `
	Number  *int `json:"number" form:"number" `
	Status  *int `json:"status" form:"status" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
