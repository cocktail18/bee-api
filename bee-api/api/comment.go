package api

import (
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
)

type CommentApi struct {
	BaseApi
}

func (api CommentApi) Add(c *gin.Context) {
	var req proto.CommentAddReq
	if err := c.Bind(&req); err != nil {
		api.Res(c, nil, err)
		return
	}
	req.Ip = c.ClientIP()
	err := service.GetCommentSrv().Add(c, &req)
	api.Res(c, "success", err)
}
