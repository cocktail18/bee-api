package api

import (
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type QrcodeApi struct {
	BaseApi
}

func (a QrcodeApi) WxaUnLimit(c *gin.Context) {
	scene := c.PostForm("scene")
	page := c.PostForm("page")
	isHyaline := cast.ToBool("is_hyaline")
	autoColor := cast.ToBool("autoColor")
	expireHours := cast.ToBool("expireHours")
	url, err := service.GetQrCodeSrv().WxaUnLimit(c, scene, page, isHyaline, autoColor, expireHours)
	a.Res(c, url, err)
}

func (a QrcodeApi) Content(c *gin.Context) {
	content := c.PostForm("content")
	width := cast.ToInt(c.PostForm("width"))
	expireHours := cast.ToInt64(c.PostForm("expireHours"))
	url, err := service.GetQrCodeSrv().GenByContent(c, content, width, expireHours)
	a.Res(c, url, err)
}
