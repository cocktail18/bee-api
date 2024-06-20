package api

import (
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
)

type FeeAPi struct {
	BaseApi
}

func (p FeeAPi) ListPeiSong(c *gin.Context) {
	resp, err := service.GetFeeSrv().ListPeiSong(c)
	p.Res(c, resp, err)
}
