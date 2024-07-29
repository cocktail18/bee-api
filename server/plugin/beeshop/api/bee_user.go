package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/service"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BeeUserApi struct{}

var beeUserService = service.ServiceGroupApp.BeeServiceGroup.BeeUserService

// CreateBeeUser 创建beeUser表
// @Tags BeeUser
// @Summary 创建beeUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUser true "创建beeUser表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeUser/createBeeUser [post]
func (beeUserApi *BeeUserApi) CreateBeeUser(c *gin.Context) {
	var beeUser bee.BeeUser
	err := c.ShouldBindJSON(&beeUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUser.UserId = &shopUserId
	if err := beeUserService.CreateBeeUser(&beeUser); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeUser 删除beeUser表
// @Tags BeeUser
// @Summary 删除beeUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUser true "删除beeUser表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeUser/deleteBeeUser [delete]
func (beeUserApi *BeeUserApi) DeleteBeeUser(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUserService.DeleteBeeUser(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeUserByIds 批量删除beeUser表
// @Tags BeeUser
// @Summary 批量删除beeUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeUser/deleteBeeUserByIds [delete]
func (beeUserApi *BeeUserApi) DeleteBeeUserByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUserService.DeleteBeeUserByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeUser 更新beeUser表
// @Tags BeeUser
// @Summary 更新beeUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUser true "更新beeUser表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeUser/updateBeeUser [put]
func (beeUserApi *BeeUserApi) UpdateBeeUser(c *gin.Context) {
	var beeUser bee.BeeUser
	err := c.ShouldBindJSON(&beeUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUser.UserId = &shopUserId
	if err := beeUserService.UpdateBeeUser(beeUser, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeUser 用id查询beeUser表
// @Tags BeeUser
// @Summary 用id查询beeUser表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeUser true "用id查询beeUser表"
// @Success 200 {object} response.Response{data=object{rebeeUser=bee.BeeUser},msg=string} "查询成功"
// @Router /beeUser/findBeeUser [get]
func (beeUserApi *BeeUserApi) FindBeeUser(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeUser, err := beeUserService.GetBeeUser(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeUser, c)
	}
}

// GetBeeUserList 分页获取beeUser表列表
// @Tags BeeUser
// @Summary 分页获取beeUser表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUserSearch true "分页获取beeUser表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeUser/getBeeUserList [get]
func (beeUserApi *BeeUserApi) GetBeeUserList(c *gin.Context) {
	var pageInfo beeReq.BeeUserSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeUserService.GetBeeUserInfoList(pageInfo, shopUserId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetBeeUserPublic 不需要鉴权的beeUser表接口
// @Tags BeeUser
// @Summary 不需要鉴权的beeUser表接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUserSearch true "分页获取beeUser表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeUser/getBeeUserPublic [get]
func (beeUserApi *BeeUserApi) GetBeeUserPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的beeUser表接口信息",
	}, "获取成功", c)
}
