package api

import (
	"gitee.com/stuinfer/bee-api/logger"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	yunlaba_sdk "yunlaba/gosdk"
)

type YunlabaAPi struct {
	BaseApi
}

func (a YunlabaAPi) Notify(c *gin.Context) {
	var param yunlaba_sdk.NotifyData
	resp, err := service.GetYunlabaSrv().Notify(c, &param)
	if err != nil {
		logger.GetLogger().Error("云喇叭回调失败", zap.Error(err), zap.Any("param", param))
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"code": 500})
		return
	}
	logger.GetLogger().Info("云喇叭回调返回", zap.Any("resp", resp), zap.Any("param", param))
	c.JSON(http.StatusOK, resp)
}
