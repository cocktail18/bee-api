package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type BeeOrderPeisongSearch struct {
	OrderId        *int   `json:"orderId" form:"orderId" `
	PeisongOrderId string `json:"peisongOrderId" form:"peisongOrderId" `
	Status         *int   `json:"status" form:"status" `
	Type           *int   `json:"type" form:"type" `
	ThirdStatus    string `json:"thirdStatus" form:"thirdStatus" `
	RetryTimes     *int   `json:"retryTimes" form:"retryTimes" `
	LastRetryUnix  *int   `json:"lastRetryUnix" form:"lastRetryUnix" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}

type BeeOrderPeisongCancel struct {
	Id             int    `json:"id" form:"id"`
	PeisongOrderNo string `json:"peisongOrderNo" form:"peisongOrderNo"`
	ReasonId       int    `json:"reasonId" form:"reasonId"`
	Reason         string `json:"reason" form:"reason"`
}

type BeeOrderPeisongQueryDetail struct {
	Id int `json:"id" form:"id"`
}
