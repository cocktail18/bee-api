package api

import (
	"github.com/gin-gonic/gin"
)

type QrcodeApi struct {
	BaseApi
}

func (a QrcodeApi) WxaUnLimit(c *gin.Context) {
	//@todo 获取小程序码
}
