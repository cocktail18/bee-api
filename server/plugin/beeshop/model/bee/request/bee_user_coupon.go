package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BeeUserCouponSearch struct {
	StartDateAdd *time.Time `json:"startDateAdd" form:"startDateAdd"`
	EndDateAdd   *time.Time `json:"endDateAdd" form:"endDateAdd"`
	Uid          *int       `json:"uid" form:"uid" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
