package api

import (
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type ScoreApi struct {
	BaseApi
}

func (api ScoreApi) SignLogs(c *gin.Context) {
	list, err := service.GetScoreSrv().GetSignLogs(c)
	api.Res(c, list, err)
}

func (api ScoreApi) Sign(c *gin.Context) {
	err := service.GetScoreSrv().Sign(c)
	api.Res(c, "success", err)
}

// WxaGroup 转发微信群赠送用户积分
func (api ScoreApi) WxaGroup(c *gin.Context) {
	err := service.GetScoreSrv().WxaGroup(c)
	api.Res(c, nil, err)
}

func (api ScoreApi) Logs(c *gin.Context) {
	page := cast.ToInt64(c.PostForm("page"))
	pageSize := cast.ToInt64(c.PostForm("pageSize"))
	res, err := service.GetScoreSrv().GetScoreLogs(c, page, pageSize)
	api.Res(c, res, err)
}
