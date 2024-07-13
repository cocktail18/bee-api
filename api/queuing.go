package api

import (
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type QueuingApi struct {
	BaseApi
}

func (a QueuingApi) Types(c *gin.Context) {
	status := c.PostForm("status")
	res, err := service.GetQueueSrv().GetQueueTypes(c, status)
	a.Res(c, res, err)
}

func (a QueuingApi) My(c *gin.Context) {
	typeId := c.PostForm("typeId")
	status := c.PostForm("status")
	res, err := service.GetQueueSrv().GetMyQueueInfo(c, cast.ToInt64(typeId), status)
	a.Res(c, res, err)
}

func (a QueuingApi) Get(c *gin.Context) {
	typeId := c.PostForm("typeId")
	mobile := c.PostForm("mobile")
	err := service.GetQueueSrv().Get(c, cast.ToInt64(typeId), mobile)
	a.Res(c, "success", err)
}
