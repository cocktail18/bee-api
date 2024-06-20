package kit

import (
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/inter"
	"gitee.com/stuinfer/bee-api/model/sys"
	"github.com/gin-gonic/gin"
)

// GetUid 真实用户id
func GetUid(c *gin.Context) int64 {
	data, _ := c.Get(enum.UserInfoKey)
	userInfo := data.(inter.User)
	return userInfo.GetUid()
}

// GetUserId 主账号id
func GetUserId(c *gin.Context) int64 {
	if c == nil {
		return 0
	}
	sysUserInfo := c.MustGet(string(enum.CtxKeySysUser))
	userInfo := sysUserInfo.(*sys.SysUserModel)
	return userInfo.Id
}
