package api

import (
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
)

type PayBillApi struct {
	BaseApi
}

func (a PayBillApi) Discounts(c *gin.Context) {
	a.Res(c, make([]interface{}, 0), nil)
}

func (a PayBillApi) Pay(c *gin.Context) {
	pwd := c.PostForm("pwd")
	smsCode := c.PostForm("smsCode")
	money, err := decimal.NewFromString(c.PostForm("money"))
	if err != nil {
		a.Res(c, nil, enum.ErrParamError)
		return
	}
	shopId := cast.ToInt64(c.PostForm("shop_id"))
	calculate := cast.ToBool(c.PostForm("calculate"))
	res, err := service.GetPaySrv().PayBill(c, shopId, money, pwd, calculate, smsCode)
	a.Res(c, res, err)
}
