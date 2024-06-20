package api

import (
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
)

type CommonApi struct {
	BaseApi
}

func (api CommonApi) Province(c *gin.Context) {
	list, err := service.GetRegionSrv().GetAllProvince()
	api.Res(c, list, err)
}

func (api CommonApi) Child(c *gin.Context) {
	pid := c.Query("pid")
	list, err := service.GetRegionSrv().GetAllChild(pid)
	api.Res(c, list, err)
}

func (api CommonApi) MapQQDistance(c *gin.Context) {
	//key: TJ3BZ-6LVCF-C25JP-JOA3O-EWZFJ-5FBMI
	//mode: bicycling
	//from: 19.65325,110.96756
	//to: 23.12463,113.36199
	key := c.PostForm("key")
	mode := c.PostForm("mode")
	from := c.PostForm("from")
	to := c.PostForm("to")
	resp, err := service.GetCommonSrv().GetMapQQDistance(key, mode, from, to)
	api.Res(c, resp, err)
}
