package api

import (
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/gin-gonic/gin"
)

type QueuingApi struct {
	BaseApi
}

func (a QueuingApi) Types(c *gin.Context) {
	a.Res(c, nil, enum.ErrEmpty)
}

func (a QueuingApi) My(c *gin.Context) {
	a.Res(c, nil, enum.ErrEmpty)
}
