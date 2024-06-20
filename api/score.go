package api

import (
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
)

type ScoreApi struct {
	BaseApi
}

func (api ScoreApi) SignLogs(c *gin.Context) {
	userInfo := api.GetUserInfo(c)
	list, err := service.GetScoreSrv().GetSignLogs(userInfo.UserId)
	api.Res(c, list, err)
}

func (api ScoreApi) Sign(c *gin.Context) {
	err := service.GetScoreSrv().Sign(c)
	api.Res(c, "success", err)
}

// WxaGroup 转发微信群赠送用户积分
func (api ScoreApi) WxaGroup(c *gin.Context) {
	//@todo
	api.Res(c, "success", nil)
}

func (api ScoreApi) Logs(c *gin.Context) {
	api.Res(c, nil, enum.ErrNotImplement)
}
