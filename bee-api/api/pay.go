package api

import (
	"encoding/json"
	"fmt"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/logger"
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
	"github.com/go-pay/gopay/wechat/v3"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type PayApi struct {
	BaseApi
}

func (api PayApi) WxApp(c *gin.Context) {
	money, err := decimal.NewFromString(c.PostForm("money"))
	if err != nil {
		api.Fail(c, enum.ResCodeFail, err.Error())
		return
	}
	remark := c.PostForm("remark")
	nextAction := c.PostForm("nextAction") // {"type":4,"uid":6803950,"money":"123"}, 充值的时候为空
	payName := c.PostForm("payName")

	res, err := service.GetPaySrv().GetWxAppPayInfo(c, money, remark, nextAction, payName)
	api.Res(c, res, err)
}

func (api PayApi) WxPayCallBack(c *gin.Context) {
	//支付回调
	logger.GetLogger().Info("微信回调",
		zap.Any("uri", c.Request.RequestURI),
		zap.Any("userId", kit.GetUserId(c)))

	req, err := api.getWxV3NotifyReq(c)
	if err != nil {
		logger.GetLogger().Error("微信回调处理失败",
			zap.Any("id", req.Id),
			zap.Any("userId", kit.GetUserId(c)),
			zap.Error(err))
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	err = service.GetPaySrv().WxNotify(c, c.ClientIP(), req)
	if err != nil {
		logger.GetLogger().Error("微信回调处理失败",
			zap.Any("req", req), zap.Error(err))
		c.AbortWithStatusJSON(http.StatusInternalServerError, &proto.WxPayNotifyResp{
			Code:    "FAIL",
			Message: "失败",
		})
		return
	}
	c.JSON(http.StatusOK, &proto.WxPayNotifyResp{
		Code: "SUCCESS",
	})
	return
}

func (api PayApi) getWxV3NotifyReq(c *gin.Context) (*wechat.V3NotifyReq, error) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return nil, err
	}
	si := &wechat.SignInfo{
		HeaderTimestamp: c.GetHeader("Wechatpay-Timestamp"),
		HeaderNonce:     c.GetHeader("Wechatpay-Nonce"),
		HeaderSignature: c.GetHeader("Wechatpay-Signature"),
		HeaderSerial:    c.GetHeader("Wechatpay-Serial"),
		SignBody:        string(body),
	}
	notifyReq := &wechat.V3NotifyReq{SignInfo: si}
	if err := json.Unmarshal(body, notifyReq); err != nil {
		return nil, fmt.Errorf("json.Unmarshal(%s, %+v)：%w", string(body), notifyReq, err)
	}
	return notifyReq, nil
}

func (api PayApi) RechargeSendRule(c *gin.Context) {
	resp, err := service.GetPaySrv().RechargeSendRule(c)
	api.Res(c, resp, err)
}
