package api

import (
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/gin-gonic/gin"
)

type KanjiaApi struct {
	BaseApi
}

func (api KanjiaApi) Info(c *gin.Context) {
	api.Res(c, nil, enum.ErrNotImplement)
}

func (api KanjiaApi) Help(c *gin.Context) {
	api.Res(c, nil, enum.ErrNotImplement)
}

func (api KanjiaApi) Clear(c *gin.Context) {
	api.Res(c, nil, enum.ErrNotImplement)
}

func (api KanjiaApi) My(c *gin.Context) {
	api.Res(c, nil, enum.ErrNotImplement)
}

func (api KanjiaApi) MyHelp(c *gin.Context) {
	api.Res(c, nil, enum.ErrNotImplement)
}

func (api KanjiaApi) Set(c *gin.Context) {
	api.Res(c, nil, enum.ErrNotImplement)
}

func (api KanjiaApi) Join(c *gin.Context) {
	api.Res(c, nil, enum.ErrNotImplement)
}
