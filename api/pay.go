package api

import (
	"gitee.com/stuinfer/bee-api/enum"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
)

type PayApi struct {
	BaseApi
}

func (api PayApi) WxApp(c *gin.Context) {
	money, err := decimal.NewFromString(c.PostForm("money"))
	if err != nil {
		api.Fail(c, enum.ResCodeFail, err.Error())
		return
	}
	remark := c.PostForm("remark")
	nextAction := c.PostForm("nextAction") // {"type":4,"uid":6803950,"money":"123"}
	payName := c.PostForm("payName")
	userInfo := api.GetUserInfo(c)
	logrus.Infof("WxApp userId:%v money:%v remark:%v nextAction:%v payName:%v ", userInfo.UserId, money, remark, nextAction, payName)
	//	@todo
	api.Fail(c, enum.ResCodeEmpty, "todo")
}
