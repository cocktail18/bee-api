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

type BeeWxPayConfigApi struct{}

var beeWxPayConfigService = service.ServiceGroupApp.BeeServiceGroup.BeeWxPayConfigService

// CreateBeeWxPayConfig 创建微信支付配置
// @Tags BeeWxPayConfig
// @Summary 创建微信支付配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeWxPayConfig true "创建微信支付配置"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeWxPayConfig/createBeeWxPayConfig [post]
func (beeWxPayConfigApi *BeeWxPayConfigApi) CreateBeeWxPayConfig(c *gin.Context) {
	var beeWxPayConfig bee.BeeWxPayConfig
	err := c.ShouldBindJSON(&beeWxPayConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeWxPayConfig.UserId = &shopUserId
	if err := beeWxPayConfigService.CreateBeeWxPayConfig(&beeWxPayConfig); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeWxPayConfig 删除微信支付配置
// @Tags BeeWxPayConfig
// @Summary 删除微信支付配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeWxPayConfig true "删除微信支付配置"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeWxPayConfig/deleteBeeWxPayConfig [delete]
func (beeWxPayConfigApi *BeeWxPayConfigApi) DeleteBeeWxPayConfig(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeWxPayConfigService.DeleteBeeWxPayConfig(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeWxPayConfigByIds 批量删除微信支付配置
// @Tags BeeWxPayConfig
// @Summary 批量删除微信支付配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeWxPayConfig/deleteBeeWxPayConfigByIds [delete]
func (beeWxPayConfigApi *BeeWxPayConfigApi) DeleteBeeWxPayConfigByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeWxPayConfigService.DeleteBeeWxPayConfigByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeWxPayConfig 更新微信支付配置
// @Tags BeeWxPayConfig
// @Summary 更新微信支付配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeWxPayConfig true "更新微信支付配置"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeWxPayConfig/updateBeeWxPayConfig [put]
func (beeWxPayConfigApi *BeeWxPayConfigApi) UpdateBeeWxPayConfig(c *gin.Context) {
	var beeWxPayConfig bee.BeeWxPayConfig
	err := c.ShouldBindJSON(&beeWxPayConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeWxPayConfig.UserId = &shopUserId
	if err := beeWxPayConfigService.UpdateBeeWxPayConfig(beeWxPayConfig, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeWxPayConfig 用id查询微信支付配置
// @Tags BeeWxPayConfig
// @Summary 用id查询微信支付配置
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeWxPayConfig true "用id查询微信支付配置"
// @Success 200 {object} response.Response{data=object{rebeeWxPayConfig=bee.BeeWxPayConfig},msg=string} "查询成功"
// @Router /beeWxPayConfig/findBeeWxPayConfig [get]
func (beeWxPayConfigApi *BeeWxPayConfigApi) FindBeeWxPayConfig(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeWxPayConfig, err := beeWxPayConfigService.GetBeeWxPayConfig(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeWxPayConfig, c)
	}
}

// GetBeeWxPayConfigList 分页获取微信支付配置列表
// @Tags BeeWxPayConfig
// @Summary 分页获取微信支付配置列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeWxPayConfigSearch true "分页获取微信支付配置列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeWxPayConfig/getBeeWxPayConfigList [get]
func (beeWxPayConfigApi *BeeWxPayConfigApi) GetBeeWxPayConfigList(c *gin.Context) {
	var pageInfo beeReq.BeeWxPayConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeWxPayConfigService.GetBeeWxPayConfigInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeWxPayConfigPublic 不需要鉴权的微信支付配置接口
// @Tags BeeWxPayConfig
// @Summary 不需要鉴权的微信支付配置接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeWxPayConfigSearch true "分页获取微信支付配置列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeWxPayConfig/getBeeWxPayConfigPublic [get]
func (beeWxPayConfigApi *BeeWxPayConfigApi) GetBeeWxPayConfigPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的微信支付配置接口信息",
	}, "获取成功", c)
}
