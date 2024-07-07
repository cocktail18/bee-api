package api

import (
	"encoding/json"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type OrderApi struct {
	BaseApi
}

func (api OrderApi) List(c *gin.Context) {
	var req proto.ListOrderReq
	if err := c.Bind(&req); err != nil {
		api.Res(c, nil, err)
		return
	}
	rsp, err := service.GetOrderSrv().List(c, &req)
	api.Res(c, rsp, err)
}

func (api OrderApi) Create(c *gin.Context) {

	var req proto.CreateOrderReq
	if err := c.Bind(&req); err != nil {
		api.Res(c, nil, err)
		return
	}
	resp, err := service.GetOrderSrv().Create(c, c.ClientIP(), &req)
	api.Res(c, resp, err)
}

func (api OrderApi) Close(c *gin.Context) {
	orderId := cast.ToInt64(c.PostForm("orderId"))
	err := service.GetOrderSrv().Close(c, orderId, "用户主动关闭")
	api.Res(c, nil, err)
}

func (api OrderApi) Pay(c *gin.Context) {
	orderId := c.PostForm("orderId") // 订单id, 多个订单之间用英文逗号分隔
	smsCode := c.PostForm("smsCode") // 短信验证码
	pwd := c.PostForm("pwd")         // 密码
	err := service.GetOrderSrv().PayByBalance(c, c.ClientIP(), orderId, smsCode, pwd)
	api.Res(c, nil, err)
}

func (api OrderApi) Delete(c *gin.Context) {
	orderId := cast.ToInt64(c.PostForm("orderId"))

	err := service.GetOrderSrv().Delete(c, orderId)
	api.Res(c, nil, err)
}

// Delivery 确认收货
func (api OrderApi) Delivery(c *gin.Context) {
	orderId := cast.ToInt64(c.PostForm("orderId"))

	err := service.GetOrderSrv().Delivery(c, orderId)
	api.Res(c, nil, err)
}

func (api OrderApi) Reputation(c *gin.Context) {
	postJsonString := c.PostForm("postJsonString")
	var item proto.BeeOrderReputationReq
	if err := json.Unmarshal([]byte(postJsonString), &item); err != nil {
		api.Res(c, nil, err)
		return
	}
	err := service.GetOrderSrv().Reputation(c, &item)
	api.Res(c, nil, err)
}

func (api OrderApi) Hx(c *gin.Context) {
	hxNumber := c.PostForm("hxNumber")
	err := service.GetOrderSrv().Hx(c, hxNumber)
	api.Res(c, nil, err)
}

func (api OrderApi) Detail(c *gin.Context) {
	hxNumber := c.PostForm("hxNumber")
	orderId := c.PostForm("orderId")
	peisongOrderId := c.PostForm("peisongOrderId")
	if peisongOrderId != "" {
		api.Res(c, nil, enum.ErrNotImplement)
		return
	}
	rsp, err := service.GetOrderSrv().Detail(c, cast.ToInt64(orderId), hxNumber)
	api.Res(c, rsp, err)
}
