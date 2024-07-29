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

type BeeOrderLogisticsApi struct{}

var beeOrderLogisticsService = service.ServiceGroupApp.BeeServiceGroup.BeeOrderLogisticsService

// CreateBeeOrderLogistics 创建用户订单地址
// @Tags BeeOrderLogistics
// @Summary 创建用户订单地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderLogistics true "创建用户订单地址"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeOrderLogistics/createBeeOrderLogistics [post]
func (beeOrderLogisticsApi *BeeOrderLogisticsApi) CreateBeeOrderLogistics(c *gin.Context) {
	var beeOrderLogistics bee.BeeOrderLogistics
	err := c.ShouldBindJSON(&beeOrderLogistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeOrderLogistics.UserId = &shopUserId
	if err := beeOrderLogisticsService.CreateBeeOrderLogistics(&beeOrderLogistics); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeOrderLogistics 删除用户订单地址
// @Tags BeeOrderLogistics
// @Summary 删除用户订单地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderLogistics true "删除用户订单地址"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeOrderLogistics/deleteBeeOrderLogistics [delete]
func (beeOrderLogisticsApi *BeeOrderLogisticsApi) DeleteBeeOrderLogistics(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderLogisticsService.DeleteBeeOrderLogistics(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeOrderLogisticsByIds 批量删除用户订单地址
// @Tags BeeOrderLogistics
// @Summary 批量删除用户订单地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeOrderLogistics/deleteBeeOrderLogisticsByIds [delete]
func (beeOrderLogisticsApi *BeeOrderLogisticsApi) DeleteBeeOrderLogisticsByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderLogisticsService.DeleteBeeOrderLogisticsByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeOrderLogistics 更新用户订单地址
// @Tags BeeOrderLogistics
// @Summary 更新用户订单地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderLogistics true "更新用户订单地址"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeOrderLogistics/updateBeeOrderLogistics [put]
func (beeOrderLogisticsApi *BeeOrderLogisticsApi) UpdateBeeOrderLogistics(c *gin.Context) {
	var beeOrderLogistics bee.BeeOrderLogistics
	err := c.ShouldBindJSON(&beeOrderLogistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeOrderLogistics.UserId = &shopUserId
	if err := beeOrderLogisticsService.UpdateBeeOrderLogistics(beeOrderLogistics, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeOrderLogistics 用id查询用户订单地址
// @Tags BeeOrderLogistics
// @Summary 用id查询用户订单地址
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeOrderLogistics true "用id查询用户订单地址"
// @Success 200 {object} response.Response{data=object{rebeeOrderLogistics=bee.BeeOrderLogistics},msg=string} "查询成功"
// @Router /beeOrderLogistics/findBeeOrderLogistics [get]
func (beeOrderLogisticsApi *BeeOrderLogisticsApi) FindBeeOrderLogistics(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeOrderLogistics, err := beeOrderLogisticsService.GetBeeOrderLogistics(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeOrderLogistics, c)
	}
}

// GetBeeOrderLogisticsList 分页获取用户订单地址列表
// @Tags BeeOrderLogistics
// @Summary 分页获取用户订单地址列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeOrderLogisticsSearch true "分页获取用户订单地址列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeOrderLogistics/getBeeOrderLogisticsList [get]
func (beeOrderLogisticsApi *BeeOrderLogisticsApi) GetBeeOrderLogisticsList(c *gin.Context) {
	var pageInfo beeReq.BeeOrderLogisticsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeOrderLogisticsService.GetBeeOrderLogisticsInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeOrderLogisticsPublic 不需要鉴权的用户订单地址接口
// @Tags BeeOrderLogistics
// @Summary 不需要鉴权的用户订单地址接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeOrderLogisticsSearch true "分页获取用户订单地址列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeOrderLogistics/getBeeOrderLogisticsPublic [get]
func (beeOrderLogisticsApi *BeeOrderLogisticsApi) GetBeeOrderLogisticsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的用户订单地址接口信息",
	}, "获取成功", c)
}
