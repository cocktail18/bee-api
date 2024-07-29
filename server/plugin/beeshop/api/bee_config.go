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

type BeeConfigApi struct{}

var beeConfigService = service.ServiceGroupApp.BeeServiceGroup.BeeConfigService

// CreateBeeConfig 创建公用配置表
// @Tags BeeConfig
// @Summary 创建公用配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeConfig true "创建公用配置表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeConfig/createBeeConfig [post]
func (beeConfigApi *BeeConfigApi) CreateBeeConfig(c *gin.Context) {
	var beeConfig bee.BeeConfig
	err := c.ShouldBindJSON(&beeConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeConfig.UserId = &shopUserId
	if err := beeConfigService.CreateBeeConfig(&beeConfig); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeConfig 删除公用配置表
// @Tags BeeConfig
// @Summary 删除公用配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeConfig true "删除公用配置表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeConfig/deleteBeeConfig [delete]
func (beeConfigApi *BeeConfigApi) DeleteBeeConfig(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeConfigService.DeleteBeeConfig(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeConfigByIds 批量删除公用配置表
// @Tags BeeConfig
// @Summary 批量删除公用配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeConfig/deleteBeeConfigByIds [delete]
func (beeConfigApi *BeeConfigApi) DeleteBeeConfigByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeConfigService.DeleteBeeConfigByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeConfig 更新公用配置表
// @Tags BeeConfig
// @Summary 更新公用配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeConfig true "更新公用配置表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeConfig/updateBeeConfig [put]
func (beeConfigApi *BeeConfigApi) UpdateBeeConfig(c *gin.Context) {
	var beeConfig bee.BeeConfig
	err := c.ShouldBindJSON(&beeConfig)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeConfig.UserId = &shopUserId
	if err := beeConfigService.UpdateBeeConfig(beeConfig, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeConfig 用id查询公用配置表
// @Tags BeeConfig
// @Summary 用id查询公用配置表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeConfig true "用id查询公用配置表"
// @Success 200 {object} response.Response{data=object{rebeeConfig=bee.BeeConfig},msg=string} "查询成功"
// @Router /beeConfig/findBeeConfig [get]
func (beeConfigApi *BeeConfigApi) FindBeeConfig(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeConfig, err := beeConfigService.GetBeeConfig(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeConfig, c)
	}
}

// GetBeeConfigList 分页获取公用配置表列表
// @Tags BeeConfig
// @Summary 分页获取公用配置表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeConfigSearch true "分页获取公用配置表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeConfig/getBeeConfigList [get]
func (beeConfigApi *BeeConfigApi) GetBeeConfigList(c *gin.Context) {
	var pageInfo beeReq.BeeConfigSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeConfigService.GetBeeConfigInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeConfigPublic 不需要鉴权的公用配置表接口
// @Tags BeeConfig
// @Summary 不需要鉴权的公用配置表接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeConfigSearch true "分页获取公用配置表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeConfig/getBeeConfigPublic [get]
func (beeConfigApi *BeeConfigApi) GetBeeConfigPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的公用配置表接口信息",
	}, "获取成功", c)
}
