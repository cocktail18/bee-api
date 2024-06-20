package api

import (
	"gitee.com/stuinfer/bee-api/proto"
	"gitee.com/stuinfer/bee-api/service"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type UserApi struct {
	BaseApi
}

func (api UserApi) My(c *gin.Context) {

}

func (api UserApi) Amount(c *gin.Context) {
	//我的资产信息
	userInfo := api.GetUserInfo(c)
	data, err := service.GetUserSrv().Amount(c, userInfo.UserId)
	api.Res(c, data, err)
}

func (api UserApi) CheckToken(c *gin.Context) {
	api.Success(c, "success")
}

func (api UserApi) WxAppAuthorize(c *gin.Context) {
	code := c.PostForm("code")
	resp, err := service.GetUserSrv().Login(c, code)
	api.Res(c, resp, err)
}

func (api UserApi) Modify(c *gin.Context) {
	userInfo := api.GetUserInfo(c)
	nick := c.PostForm("nick")
	avatarUrl := c.PostForm("avatarUrl")
	city := c.PostForm("city")
	province := c.PostForm("province")
	gender := cast.ToInt(c.PostForm("gender"))
	userInfo.Nick = nick
	userInfo.AvatarUrl = avatarUrl
	userInfo.City = city
	userInfo.Province = province
	userInfo.Gender = gender
	err := service.GetUserSrv().Modify(userInfo)
	api.Res(c, "success", err)
}

func (api UserApi) Detail(c *gin.Context) {
	userInfo := api.GetUserInfo(c)
	api.Success(c, &proto.GetUserDetailResp{Base: userInfo})
}

func (api UserApi) CashLog(c *gin.Context) {
	userInfo := api.GetUserInfo(c)
	var req proto.CashLogReq
	if err := c.Bind(&req); err != nil {
		api.Res(c, nil, err)
		return
	}
	resp, err := service.GetUserSrv().CashLog(c, userInfo.UserId, &req)
	api.Res(c, resp, err)
}

func (api UserApi) RechargeSendRule(c *gin.Context) {
	resp, err := service.GetUserSrv().RechargeSendRule(c)
	api.Res(c, resp, err)
}

// PayLogs 获取充值/支付记录
func (api UserApi) PayLogs(c *gin.Context) {
	userInfo := api.GetUserInfo(c)
	var req proto.PayLogsReq
	if err := c.Bind(&req); err != nil {
		api.Res(c, nil, err)
		return
	}
	resp, err := service.GetUserSrv().PayLogs(c, userInfo.UserId, &req)
	api.Res(c, resp, err)
}

func (api UserApi) BindMobile(c *gin.Context) {
	//@todo
	api.Res(c, nil, nil)
}

// BindSeller 重新绑定分销商（抢客）
func (api UserApi) BindSeller(c *gin.Context) {
	//@todo
	api.Res(c, nil, nil)
}

func (api UserApi) LevelList(c *gin.Context) {
	resp, err := service.GetUserSrv().LevelList(c)
	api.Res(c, resp, err)
}

// GetDynamicUserCode 获取会员码
func (api UserApi) GetDynamicUserCode(c *gin.Context) {
	code, err := service.GetUserSrv().GetDynamicUserCode(c)
	api.Res(c, code, err)
}
