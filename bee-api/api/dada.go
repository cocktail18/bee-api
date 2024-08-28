package api

import (
	dadasdk "dada/gosdk"
	"gitee.com/stuinfer/bee-api/logger"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type DadaAPi struct {
	BaseApi
}

func (a DadaAPi) Notify(c *gin.Context) {
	var param dadasdk.CallbackParam
	err := c.ShouldBindBodyWithJSON(&param)
	if param.OrderId == "" || err != nil {
		var paramV2 dadasdk.CallbackParamV2
		if err := c.ShouldBindBodyWithJSON(&paramV2); err != nil {
			c.JSON(http.StatusOK, map[string]interface{}{"code": 200})
			return
		}
		param.ClientId = paramV2.ClientId
		param.OrderId = paramV2.OrderId
		param.OrderStatus = paramV2.OrderStatus
		param.RepeatReasonType = paramV2.RepeatReasonType
		param.CancelReason = paramV2.CancelReason
		param.CancelFrom = paramV2.CancelFrom
		param.UpdateTime = paramV2.UpdateTime
		param.Signature = paramV2.Signature
		param.DmId = paramV2.DmId
		param.DmName = paramV2.DmName
		param.DmMobile = paramV2.DmMobile
		param.FinishCode = paramV2.FinishCode
	}
	if param.OrderId == "" {
		c.JSON(http.StatusOK, map[string]interface{}{"code": 200})
		return
	}
	if cb, ok := c.Get(gin.BodyBytesKey); ok {
		if cbb, ok := cb.([]byte); ok {
			logger.GetLogger().Info("dada notify", zap.Any("param", string(cbb)))
		}
	}

	err = service.GetDadaSrv().Notify(c, &param)
	if err != nil {
		logger.GetLogger().Error("dada notify err", zap.Any("param", param), zap.Error(err))
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": 500})
		return
	}
	logger.GetLogger().Info("dada notify success", zap.Any("param", param))
	c.JSON(http.StatusOK, map[string]interface{}{"code": 200})
}
