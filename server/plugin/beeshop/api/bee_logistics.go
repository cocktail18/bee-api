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

type BeeLogisticsApi struct{}

var beeLogisticsService = service.ServiceGroupApp.BeeServiceGroup.BeeLogisticsService

// CreateBeeLogistics 创建运费模版
// @Tags BeeLogistics
// @Summary 创建运费模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeLogistics true "创建运费模版"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeLogistics/createBeeLogistics [post]
func (beeLogisticsApi *BeeLogisticsApi) CreateBeeLogistics(c *gin.Context) {
	var beeLogistics bee.BeeLogistics
	err := c.ShouldBindJSON(&beeLogistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeLogistics.UserId = &shopUserId
	if err := beeLogisticsService.CreateBeeLogistics(&beeLogistics); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeLogistics 删除运费模版
// @Tags BeeLogistics
// @Summary 删除运费模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeLogistics true "删除运费模版"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeLogistics/deleteBeeLogistics [delete]
func (beeLogisticsApi *BeeLogisticsApi) DeleteBeeLogistics(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeLogisticsService.DeleteBeeLogistics(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeLogisticsByIds 批量删除运费模版
// @Tags BeeLogistics
// @Summary 批量删除运费模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeLogistics/deleteBeeLogisticsByIds [delete]
func (beeLogisticsApi *BeeLogisticsApi) DeleteBeeLogisticsByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeLogisticsService.DeleteBeeLogisticsByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeLogistics 更新运费模版
// @Tags BeeLogistics
// @Summary 更新运费模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeLogistics true "更新运费模版"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeLogistics/updateBeeLogistics [put]
func (beeLogisticsApi *BeeLogisticsApi) UpdateBeeLogistics(c *gin.Context) {
	var beeLogistics bee.BeeLogistics
	err := c.ShouldBindJSON(&beeLogistics)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeLogistics.UserId = &shopUserId
	if err := beeLogisticsService.UpdateBeeLogistics(beeLogistics, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeLogistics 用id查询运费模版
// @Tags BeeLogistics
// @Summary 用id查询运费模版
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeLogistics true "用id查询运费模版"
// @Success 200 {object} response.Response{data=object{rebeeLogistics=bee.BeeLogistics},msg=string} "查询成功"
// @Router /beeLogistics/findBeeLogistics [get]
func (beeLogisticsApi *BeeLogisticsApi) FindBeeLogistics(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeLogistics, err := beeLogisticsService.GetBeeLogistics(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeLogistics, c)
	}
}

// GetBeeLogisticsList 分页获取运费模版列表
// @Tags BeeLogistics
// @Summary 分页获取运费模版列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeLogisticsSearch true "分页获取运费模版列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeLogistics/getBeeLogisticsList [get]
func (beeLogisticsApi *BeeLogisticsApi) GetBeeLogisticsList(c *gin.Context) {
	var pageInfo beeReq.BeeLogisticsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeLogisticsService.GetBeeLogisticsInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeLogisticsPublic 不需要鉴权的运费模版接口
// @Tags BeeLogistics
// @Summary 不需要鉴权的运费模版接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeLogisticsSearch true "分页获取运费模版列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeLogistics/getBeeLogisticsPublic [get]
func (beeLogisticsApi *BeeLogisticsApi) GetBeeLogisticsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的运费模版接口信息",
	}, "获取成功", c)
}
