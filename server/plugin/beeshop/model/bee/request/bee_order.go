package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type BeeOrderSearch struct {
	ID                   *int       `json:"id" form:"id" `
	IsDeleted            *bool      `json:"isDeleted" form:"isDeleted" `
	OrderNumber          *string    `json:"orderNumber" form:"orderNumber" `
	HxNumber             *string    `json:"hxNumber" form:"hxNumber" `
	StartDateAdd         *time.Time `json:"startDateAdd" form:"startDateAdd"`
	EndDateAdd           *time.Time `json:"endDateAdd" form:"endDateAdd"`
	StartDateUpdate      *time.Time `json:"startDateUpdate" form:"startDateUpdate"`
	EndDateUpdate        *time.Time `json:"endDateUpdate" form:"endDateUpdate"`
	StartDateDelete      *time.Time `json:"startDateDelete" form:"startDateDelete"`
	EndDateDelete        *time.Time `json:"endDateDelete" form:"endDateDelete"`
	StartAmountLogistics *float64   `json:"startAmountLogistics" form:"startAmountLogistics"`
	EndAmountLogistics   *float64   `json:"endAmountLogistics" form:"endAmountLogistics"`
	StartAmountReal      *float64   `json:"startAmountReal" form:"startAmountReal"`
	EndAmountReal        *float64   `json:"endAmountReal" form:"endAmountReal"`
	AutoDeliverStatus    *int       `json:"autoDeliverStatus" form:"autoDeliverStatus" `
	StartDateClose       *time.Time `json:"startDateClose" form:"startDateClose"`
	EndDateClose         *time.Time `json:"endDateClose" form:"endDateClose"`
	StartDatePay         *time.Time `json:"startDatePay" form:"startDatePay"`
	EndDatePay           *time.Time `json:"endDatePay" form:"endDatePay"`
	HasRefund            *bool      `json:"hasRefund" form:"hasRefund" `
	IsEnd                *bool      `json:"isEnd" form:"isEnd" `
	Pid                  *int       `json:"pid" form:"pid" `
	Status               *int       `json:"status" form:"status" `
	StartTrips           *int       `json:"startTrips" form:"startTrips"`
	EndTrips             *int       `json:"endTrips" form:"endTrips"`
	Type                 *int       `json:"type" form:"type" `
	Uid                  *int       `json:"uid" form:"uid" `
	request.PageInfo
	Sort  string `json:"sort" form:"sort"`
	Order string `json:"order" form:"order"`
}
