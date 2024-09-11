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

type BeeRechargeSendRuleApi struct{}

var beeRechargeSendRuleService = service.ServiceGroupApp.BeeServiceGroup.BeeRechargeSendRuleService

// CreateBeeRechargeSendRule 创建beeRechargeSendRule表
// @Tags BeeRechargeSendRule
// @Summary 创建beeRechargeSendRule表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeRechargeSendRule true "创建beeRechargeSendRule表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeRechargeSendRule/createBeeRechargeSendRule [post]
func (beeRechargeSendRuleApi *BeeRechargeSendRuleApi) CreateBeeRechargeSendRule(c *gin.Context) {
	var beeRechargeSendRule bee.BeeRechargeSendRule
	err := c.ShouldBindJSON(&beeRechargeSendRule)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeRechargeSendRule.UserId = &shopUserId
	if err := beeRechargeSendRuleService.CreateBeeRechargeSendRule(&beeRechargeSendRule); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeRechargeSendRule 删除beeRechargeSendRule表
// @Tags BeeRechargeSendRule
// @Summary 删除beeRechargeSendRule表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeRechargeSendRule true "删除beeRechargeSendRule表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeRechargeSendRule/deleteBeeRechargeSendRule [delete]
func (beeRechargeSendRuleApi *BeeRechargeSendRuleApi) DeleteBeeRechargeSendRule(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeRechargeSendRuleService.DeleteBeeRechargeSendRule(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeRechargeSendRuleByIds 批量删除beeRechargeSendRule表
// @Tags BeeRechargeSendRule
// @Summary 批量删除beeRechargeSendRule表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeRechargeSendRule/deleteBeeRechargeSendRuleByIds [delete]
func (beeRechargeSendRuleApi *BeeRechargeSendRuleApi) DeleteBeeRechargeSendRuleByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeRechargeSendRuleService.DeleteBeeRechargeSendRuleByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeRechargeSendRule 更新beeRechargeSendRule表
// @Tags BeeRechargeSendRule
// @Summary 更新beeRechargeSendRule表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeRechargeSendRule true "更新beeRechargeSendRule表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeRechargeSendRule/updateBeeRechargeSendRule [put]
func (beeRechargeSendRuleApi *BeeRechargeSendRuleApi) UpdateBeeRechargeSendRule(c *gin.Context) {
	var beeRechargeSendRule bee.BeeRechargeSendRule
	err := c.ShouldBindJSON(&beeRechargeSendRule)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeRechargeSendRule.UserId = &shopUserId
	if err := beeRechargeSendRuleService.UpdateBeeRechargeSendRule(beeRechargeSendRule, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeRechargeSendRule 用id查询beeRechargeSendRule表
// @Tags BeeRechargeSendRule
// @Summary 用id查询beeRechargeSendRule表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeRechargeSendRule true "用id查询beeRechargeSendRule表"
// @Success 200 {object} response.Response{data=object{rebeeRechargeSendRule=bee.BeeRechargeSendRule},msg=string} "查询成功"
// @Router /beeRechargeSendRule/findBeeRechargeSendRule [get]
func (beeRechargeSendRuleApi *BeeRechargeSendRuleApi) FindBeeRechargeSendRule(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeRechargeSendRule, err := beeRechargeSendRuleService.GetBeeRechargeSendRule(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeRechargeSendRule, c)
	}
}

// GetBeeRechargeSendRuleList 分页获取beeRechargeSendRule表列表
// @Tags BeeRechargeSendRule
// @Summary 分页获取beeRechargeSendRule表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeRechargeSendRuleSearch true "分页获取beeRechargeSendRule表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeRechargeSendRule/getBeeRechargeSendRuleList [get]
func (beeRechargeSendRuleApi *BeeRechargeSendRuleApi) GetBeeRechargeSendRuleList(c *gin.Context) {
	var pageInfo beeReq.BeeRechargeSendRuleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeRechargeSendRuleService.GetBeeRechargeSendRuleInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeRechargeSendRulePublic 不需要鉴权的beeRechargeSendRule表接口
// @Tags BeeRechargeSendRule
// @Summary 不需要鉴权的beeRechargeSendRule表接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeRechargeSendRuleSearch true "分页获取beeRechargeSendRule表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeRechargeSendRule/getBeeRechargeSendRulePublic [get]
func (beeRechargeSendRuleApi *BeeRechargeSendRuleApi) GetBeeRechargeSendRulePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的beeRechargeSendRule表接口信息",
	}, "获取成功", c)
}
