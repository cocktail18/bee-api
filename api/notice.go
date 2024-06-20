package api

import (
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type NoticeApi struct {
	BaseApi
}

func (api NoticeApi) LastOne(c *gin.Context) {
	info, err := service.GetNoticeSrv().GetLastOne()
	if err != nil {
		api.Fail(c, enum.ResCodeFail, err.Error())
		return
	}
	api.Success(c, info)
}

func (api NoticeApi) Detail(c *gin.Context) {
	id := cast.ToInt64(c.Query("id"))
	info, err := service.GetNoticeSrv().GetDetail(id)
	if err != nil {
		api.Fail(c, enum.ResCodeFail, err.Error())
		return
	}
	api.Success(c, info)
}
