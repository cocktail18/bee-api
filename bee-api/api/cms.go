package api

import (
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
)

type CmsApi struct {
	BaseApi
}

func (api CmsApi) Info(c *gin.Context) {
	key := c.Query("key")
	item, err := service.GetCmsSrv().GetCmsInfo(key)
	api.Res(c, item, err)
}
