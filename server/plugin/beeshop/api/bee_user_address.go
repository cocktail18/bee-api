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

type BeeUserAddressApi struct{}

var beeUserAddressService = service.ServiceGroupApp.BeeServiceGroup.BeeUserAddressService

// CreateBeeUserAddress 创建用户地址
// @Tags BeeUserAddress
// @Summary 创建用户地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserAddress true "创建用户地址"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeUserAddress/createBeeUserAddress [post]
func (beeUserAddressApi *BeeUserAddressApi) CreateBeeUserAddress(c *gin.Context) {
	var beeUserAddress bee.BeeUserAddress
	err := c.ShouldBindJSON(&beeUserAddress)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUserAddress.UserId = &shopUserId
	if err := beeUserAddressService.CreateBeeUserAddress(&beeUserAddress); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeUserAddress 删除用户地址
// @Tags BeeUserAddress
// @Summary 删除用户地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserAddress true "删除用户地址"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeUserAddress/deleteBeeUserAddress [delete]
func (beeUserAddressApi *BeeUserAddressApi) DeleteBeeUserAddress(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUserAddressService.DeleteBeeUserAddress(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeUserAddressByIds 批量删除用户地址
// @Tags BeeUserAddress
// @Summary 批量删除用户地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeUserAddress/deleteBeeUserAddressByIds [delete]
func (beeUserAddressApi *BeeUserAddressApi) DeleteBeeUserAddressByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUserAddressService.DeleteBeeUserAddressByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeUserAddress 更新用户地址
// @Tags BeeUserAddress
// @Summary 更新用户地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserAddress true "更新用户地址"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeUserAddress/updateBeeUserAddress [put]
func (beeUserAddressApi *BeeUserAddressApi) UpdateBeeUserAddress(c *gin.Context) {
	var beeUserAddress bee.BeeUserAddress
	err := c.ShouldBindJSON(&beeUserAddress)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUserAddress.UserId = &shopUserId
	if err := beeUserAddressService.UpdateBeeUserAddress(beeUserAddress, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeUserAddress 用id查询用户地址
// @Tags BeeUserAddress
// @Summary 用id查询用户地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeUserAddress true "用id查询用户地址"
// @Success 200 {object} response.Response{data=object{rebeeUserAddress=bee.BeeUserAddress},msg=string} "查询成功"
// @Router /beeUserAddress/findBeeUserAddress [get]
func (beeUserAddressApi *BeeUserAddressApi) FindBeeUserAddress(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeUserAddress, err := beeUserAddressService.GetBeeUserAddress(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeUserAddress, c)
	}
}

// GetBeeUserAddressList 分页获取用户地址列表
// @Tags BeeUserAddress
// @Summary 分页获取用户地址列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUserAddressSearch true "分页获取用户地址列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeUserAddress/getBeeUserAddressList [get]
func (beeUserAddressApi *BeeUserAddressApi) GetBeeUserAddressList(c *gin.Context) {
	var pageInfo beeReq.BeeUserAddressSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeUserAddressService.GetBeeUserAddressInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeUserAddressPublic 不需要鉴权的用户地址接口
// @Tags BeeUserAddress
// @Summary 不需要鉴权的用户地址接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUserAddressSearch true "分页获取用户地址列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeUserAddress/getBeeUserAddressPublic [get]
func (beeUserAddressApi *BeeUserAddressApi) GetBeeUserAddressPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的用户地址接口信息",
	}, "获取成功", c)
}
