package api

import (
	"github.com/gin-gonic/gin"
)

type CardApi struct {
	BaseApi
}

func (a CardApi) List(c *gin.Context) {
	a.Res(c, nil, nil)
}

func (a CardApi) Get(c *gin.Context) {
	a.Res(c, nil, nil)
}

func (a CardApi) Buy(c *gin.Context) {
	a.Res(c, nil, nil)
}

func (a CardApi) Logs(c *gin.Context) {
	a.Res(c, nil, nil)
}

func (a CardApi) My(c *gin.Context) {
	a.Res(c, nil, nil)
}

func (a CardApi) Exchange(c *gin.Context) {
	a.Res(c, nil, nil)
}

func (a CardApi) ShareOpen(c *gin.Context) {
	a.Res(c, nil, nil)
}

func (a CardApi) ShareFetch(c *gin.Context) {
	a.Res(c, nil, nil)
}

func (a CardApi) ShareClose(c *gin.Context) {
	a.Res(c, nil, nil)
}
