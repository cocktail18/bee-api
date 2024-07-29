package test

import (
	config2 "gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/model"
	"gitee.com/stuinfer/bee-api/model/sys"
	"golang.org/x/net/context"
)

func GetTestContext() context.Context {
	ctx := context.Background()
	sysUserInfo := &sys.SysUserModel{}
	db.GetDB().Order("id asc").First(sysUserInfo)
	ctx = context.WithValue(ctx, string(enum.CtxKeySysUser), sysUserInfo)

	userInfo := &model.BeeUser{}
	db.GetDB().Order("id asc").First(userInfo)
	ctx = context.WithValue(ctx, enum.UserInfoKey, userInfo)
	return ctx
}

func GetTestContextByUserInfo(userInfo *model.BeeUser) context.Context {
	ctx := context.Background()
	sysUserInfo := &sys.SysUserModel{}
	db.GetDB().Order("id asc").First(sysUserInfo)
	ctx = context.WithValue(ctx, string(enum.CtxKeySysUser), sysUserInfo)

	ctx = context.WithValue(ctx, enum.UserInfoKey, userInfo)
	return ctx
}

func InitTestConfig() {
	config2.AppConfigIns.Default = &config2.DefaultConfig{
		SysUser: &config2.SysUser{
			Username: "admin",
			Password: "123456",
		},
		Wx: &config2.WxConfig{
			AppId:  "wx0a0a0a0a0a0a0a0a",
			Secret: "0a0a0a0a0a0a0a0a0a0a0a0a0a0a0a0a",
		},
	}
	config2.AppConfigIns.DB = &config2.AppDBConfig{
		Database: "bee",
		Host:     "127.0.0.1",
	}
}
