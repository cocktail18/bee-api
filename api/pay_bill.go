package api

import (
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/gin-gonic/gin"
)

type PayBillApi struct {
	BaseApi
}

func (a PayBillApi) Discounts(c *gin.Context) {
	a.Res(c, nil, enum.ErrNotImplement)
}

func (a PayBillApi) Pay(c *gin.Context) {
	a.Res(c, nil, enum.ErrNotImplement)
}
