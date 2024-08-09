package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BeePayLogSearch struct {
	StartDateAdd    *time.Time `json:"startDateAdd" form:"startDateAdd"`
	EndDateAdd      *time.Time `json:"endDateAdd" form:"endDateAdd"`
	StartDateUpdate *time.Time `json:"startDateUpdate" form:"startDateUpdate"`
	EndDateUpdate   *time.Time `json:"endDateUpdate" form:"endDateUpdate"`
	StartMoney      *float64   `json:"startMoney" form:"startMoney"`
	EndMoney        *float64   `json:"endMoney" form:"endMoney"`
	OrderNo         string     `json:"orderNo" form:"orderNo" `
	Status          string     `json:"status" form:"status" `
	Uid             *int       `json:"uid" form:"uid" `
	request.PageInfo
	Sort     string `json:"sort" form:"sort"`
	Order    string `json:"order" form:"order"`
	Distinct string `json:"distinct" form:"distinct"`
	Sum      string `json:"sum" form:"sum"`
}
