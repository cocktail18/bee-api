package api

import (
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
)

type BannerApi struct {
	BaseApi
}

func (api BannerApi) List(c *gin.Context) {
	list, err := service.GetBannerSrv().GetBannerList()
	if err != nil {
		api.Fail(c, enum.ResCodeFail, err.Error())
		return
	}
	api.Success(c, list)
}
