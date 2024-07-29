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

type BeeWxConfigApi struct{}

var beeWxConfigService = service.ServiceGroupApp.BeeServiceGroup.BeeWxConfigService

// CreateBeeWxConfig 创建微信配置
// @Tags BeeWxConfig
// @Summary 创建微信配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeWxConfig true "创建微信配置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeWxConfig/createBeeWxConfig [post]
func (beeWxConfigApi *BeeWxConfigApi) CreateBeeWxConfig(c *gin.Context) {
	var beeWxConfig bee.BeeWxConfig
	err := c.ShouldBindJSON(&beeWxConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeWxConfig.UserId = &shopUserId
	if err := beeWxConfigService.CreateBeeWxConfig(&beeWxConfig); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeWxConfig 删除微信配置
// @Tags BeeWxConfig
// @Summary 删除微信配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeWxConfig true "删除微信配置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeWxConfig/deleteBeeWxConfig [delete]
func (beeWxConfigApi *BeeWxConfigApi) DeleteBeeWxConfig(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeWxConfigService.DeleteBeeWxConfig(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeWxConfigByIds 批量删除微信配置
// @Tags BeeWxConfig
// @Summary 批量删除微信配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeWxConfig/deleteBeeWxConfigByIds [delete]
func (beeWxConfigApi *BeeWxConfigApi) DeleteBeeWxConfigByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeWxConfigService.DeleteBeeWxConfigByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeWxConfig 更新微信配置
// @Tags BeeWxConfig
// @Summary 更新微信配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeWxConfig true "更新微信配置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeWxConfig/updateBeeWxConfig [put]
func (beeWxConfigApi *BeeWxConfigApi) UpdateBeeWxConfig(c *gin.Context) {
	var beeWxConfig bee.BeeWxConfig
	err := c.ShouldBindJSON(&beeWxConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeWxConfig.UserId = &shopUserId
	if err := beeWxConfigService.UpdateBeeWxConfig(beeWxConfig, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeWxConfig 用id查询微信配置
// @Tags BeeWxConfig
// @Summary 用id查询微信配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeWxConfig true "用id查询微信配置"
// @Success 200 {object} response.Response{data=object{beeWxConfig=bee.BeeWxConfig},msg=string} "查询成功"
// @Router /beeWxConfig/findBeeWxConfig [get]
func (beeWxConfigApi *BeeWxConfigApi) FindBeeWxConfig(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeWxConfig, err := beeWxConfigService.GetBeeWxConfig(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeWxConfig, c)
	}
}

// GetBeeWxConfigList 分页获取微信配置列表
// @Tags BeeWxConfig
// @Summary 分页获取微信配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeWxConfigSearch true "分页获取微信配置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeWxConfig/getBeeWxConfigList [get]
func (beeWxConfigApi *BeeWxConfigApi) GetBeeWxConfigList(c *gin.Context) {
	var pageInfo beeReq.BeeWxConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeWxConfigService.GetBeeWxConfigInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeWxConfigPublic 不需要鉴权的微信配置接口
// @Tags BeeWxConfig
// @Summary 不需要鉴权的微信配置接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeWxConfigSearch true "分页获取微信配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeWxConfig/getBeeWxConfigPublic [get]
func (beeWxConfigApi *BeeWxConfigApi) GetBeeWxConfigPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的微信配置接口信息",
	}, "获取成功", c)
}
