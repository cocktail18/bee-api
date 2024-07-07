package kit

import (
	"context"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/inter"
	"gitee.com/stuinfer/bee-api/model/sys"
)

// GetUid 真实用户id
func GetUid(c context.Context) int64 {
	data := c.Value(enum.UserInfoKey)
	userInfo := data.(inter.User)
	return userInfo.GetUid()
}

// GetSessionKey 获取微信session key
func GetSessionKey(c context.Context) string {
	data := c.Value(enum.UserInfoKey)
	userInfo := data.(inter.User)
	return userInfo.GetSessionKey()
}

// GetUserId 主账号id
func GetUserId(c context.Context) int64 {
	if c == nil {
		return 0
	}
	sysUserInfo := c.Value(string(enum.CtxKeySysUser))
	userInfo := sysUserInfo.(*sys.SysUserModel)
	return userInfo.Id
}

// GetDomain 获取域名
func GetDomain(c context.Context) string {
	if c == nil {
		return ""
	}
	sysUserInfo := c.Value(string(enum.CtxKeySysUser))
	userInfo := sysUserInfo.(*sys.SysUserModel)
	return userInfo.Domain
}

// GetUserInfo 用户信息
func GetUserInfo(c context.Context) *sys.SysUserModel {
	if c == nil {
		return nil
	}
	sysUserInfo := c.Value(string(enum.CtxKeySysUser))
	userInfo := sysUserInfo.(*sys.SysUserModel)
	return userInfo
}
