package proto

import (
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/model"
	"github.com/golang-module/carbon/v2"
	"github.com/shopspring/decimal"
	"time"
)

type UserCouponResp struct {
	DateAdd       string               `json:"dateAdd"`
	DateEnd       string               `json:"dateEnd"`
	DateStart     string               `json:"dateStart"`
	ExpiryMillis  int64                `json:"expiryMillis"`
	Id            int64                `json:"id"`
	Money         decimal.Decimal      `json:"money"`
	MoneyHreshold decimal.Decimal      `json:"moneyHreshold"`
	MoneyType     enum.CouponMoneyType `json:"moneyType"`
	Name          string               `json:"name"`
	OnlyFreight   bool                 `json:"onlyFreight"`
	Pid           int64                `json:"pid"`
	ShopId        int64                `json:"shopId"`
	Source        string               `json:"source"`
	Status        enum.CouponStatus    `json:"status"`
	StatusStr     string               `json:"statusStr"`
}

func NewUserCouponResp(coupon *model.BeeUserCoupon) *UserCouponResp {
	return &UserCouponResp{
		DateStart:     time.Time(coupon.DateStart).Format(carbon.DateTimeLayout),
		ExpiryMillis:  coupon.ExpiryMillis,
		Id:            coupon.Id,
		Money:         coupon.Money,
		MoneyHreshold: coupon.MoneyHreshold,
		MoneyType:     coupon.MoneyType,
		Name:          coupon.Name,
		OnlyFreight:   coupon.OnlyFreight,
		Pid:           coupon.Pid,
		ShopId:        coupon.ShopId,
		Source:        coupon.Source,
		Status:        coupon.Status,
		StatusStr:     enum.CouponStatusMap[coupon.Status],
	}
}

type MyCouponStatisticsResp struct {
	CanFetch int `json:"canFetch"`
	CanUse   int `json:"canUse"`
	Invalid  int `json:"invalid"`
	Used     int `json:"used"`
}

type CouponDetailResp struct {
	Info *model.BeeCoupon `json:"info"`
}
