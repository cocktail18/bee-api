package api

import (
	"gitee.com/stuinfer/bee-api/service"
	"gitee.com/stuinfer/bee-api/service/upload"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type DfsApi struct {
	BaseApi
}

func (a DfsApi) UploadFile(c *gin.Context) {
	subDomain := c.PostForm("subDomain")
	expireHours := cast.ToInt64(c.PostForm("expireHours"))
	file, err := c.FormFile("upfile")
	if err != nil {
		a.Res(c, "", err)
		return
	}

	// 保存文件到指定路径
	// @todo 定时删除过期文件
	filepath, filename, err := upload.NewOss().UploadFile(file)
	if err != nil {
		a.Res(c, "", err)
		return
	}
	resp, err := service.GetDfsSrv().SaveUploadedFile(c, c.Request.Host, subDomain, filename, filepath, expireHours)
	a.Res(c, resp, err)
}
