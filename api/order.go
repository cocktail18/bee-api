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
	userId := api.GetUserInfo(c).UserId
	var req proto.ListOrderReq
	if err := c.Bind(&req); err != nil {
		api.Res(c, nil, err)
		return
	}
	rsp, err := service.GetOrderSrv().List(c, userId, &req)
	api.Res(c, rsp, err)
}

func (api OrderApi) Create(c *gin.Context) {

	goodsJsonStr := c.PostForm("goodsJsonStr")
	remark := c.PostForm("remark")
	peisongType := c.PostForm("peisongType")
	isCanHx := cast.ToBool(c.PostForm("isCanHx"))
	shopIdZt := cast.ToInt64(c.PostForm("shopIdZt"))
	couponId := cast.ToInt64(c.PostForm("couponId"))
	extJsonStr := c.PostForm("extJsonStr")
	calculate := cast.ToBool(c.PostForm("calculate"))

	orderGoodsList := make([]*proto.BeeOrderGoods, 0, 10)
	err := json.Unmarshal([]byte(goodsJsonStr), &orderGoodsList)
	if err != nil {
		api.Res(c, nil, err)
		return
	}
	resp, err := service.GetOrderSrv().Create(c, orderGoodsList, remark, peisongType, isCanHx, shopIdZt, couponId, extJsonStr, calculate)
	api.Res(c, resp, err)
}

func (api OrderApi) Close(c *gin.Context) {
	userId := api.GetUserInfo(c).UserId
	orderId := cast.ToInt64(c.PostForm("orderId"))

	err := service.GetOrderSrv().Close(userId, orderId, "用户主动关闭")
	api.Res(c, nil, err)
}

func (api OrderApi) Pay(c *gin.Context) {
	api.Res(c, nil, enum.ErrNotImplement)
}

func (api OrderApi) Delete(c *gin.Context) {
	userId := api.GetUserInfo(c).UserId
	orderId := cast.ToInt64(c.PostForm("orderId"))

	err := service.GetOrderSrv().Delete(c, userId, orderId)
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
