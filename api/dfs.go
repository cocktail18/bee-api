package api

import (
	"gitee.com/stuinfer/bee-api/service"
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
	dst := "uploads/" + subDomain + "/" + file.Filename
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		a.Res(c, "", err)
		return
	}
	resp, err := service.GetDfsSrv().SaveUploadedFile(c, subDomain, file.Filename, dst, expireHours)
	a.Res(c, resp, err)
}
