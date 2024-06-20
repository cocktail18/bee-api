package api

import (
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/kit"
	"gitee.com/stuinfer/bee-api/model"
	"github.com/gin-gonic/gin"
	"strings"
)

type ConfigApi struct {
	BaseApi
}

func (api ConfigApi) Values(c *gin.Context) {
	keys := c.Query("keys")
	keysArr := strings.Split(keys, ",")
	var list []*model.BeeConfig
	err := db.GetDB().WithContext(c.Request.Context()).Where("`key` IN (?) and user_id = ?", keysArr, kit.GetUserId(c)).
		Find(&list).Error
	if err != nil {
		api.Fail(c, enum.ResCodeFail, "系统错误")
		return
	}
	api.Success(c, list)
}
