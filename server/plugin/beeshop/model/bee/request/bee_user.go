package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BeeUserSearch struct {
	Id           *int       `json:"id" form:"id" `
	ShowUid      *int       `json:"showUid" form:"showUid" `
	CardNumber   string     `json:"cardNumber" form:"cardNumber" `
	IsVirtual    *bool      `json:"isVirtual" form:"isVirtual" `
	Status       *int       `json:"status" form:"status" `
	StartDateAdd *time.Time `json:"startDateAdd" form:"startDateAdd"`
	EndDateAdd   *time.Time `json:"endDateAdd" form:"endDateAdd"`
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
