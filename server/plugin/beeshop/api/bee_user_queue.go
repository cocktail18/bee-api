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

type BeeUserQueueApi struct{}

var beeUserQueueService = service.ServiceGroupApp.BeeServiceGroup.BeeUserQueueService

// CreateBeeUserQueue 创建beeUserQueue表
// @Tags BeeUserQueue
// @Summary 创建beeUserQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserQueue true "创建beeUserQueue表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeUserQueue/createBeeUserQueue [post]
func (beeUserQueueApi *BeeUserQueueApi) CreateBeeUserQueue(c *gin.Context) {
	var beeUserQueue bee.BeeUserQueue
	err := c.ShouldBindJSON(&beeUserQueue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUserQueue.UserId = &shopUserId
	if err := beeUserQueueService.CreateBeeUserQueue(&beeUserQueue); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeUserQueue 删除beeUserQueue表
// @Tags BeeUserQueue
// @Summary 删除beeUserQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserQueue true "删除beeUserQueue表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeUserQueue/deleteBeeUserQueue [delete]
func (beeUserQueueApi *BeeUserQueueApi) DeleteBeeUserQueue(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUserQueueService.DeleteBeeUserQueue(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeUserQueueByIds 批量删除beeUserQueue表
// @Tags BeeUserQueue
// @Summary 批量删除beeUserQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeUserQueue/deleteBeeUserQueueByIds [delete]
func (beeUserQueueApi *BeeUserQueueApi) DeleteBeeUserQueueByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUserQueueService.DeleteBeeUserQueueByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeUserQueue 更新beeUserQueue表
// @Tags BeeUserQueue
// @Summary 更新beeUserQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserQueue true "更新beeUserQueue表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeUserQueue/updateBeeUserQueue [put]
func (beeUserQueueApi *BeeUserQueueApi) UpdateBeeUserQueue(c *gin.Context) {
	var beeUserQueue bee.BeeUserQueue
	err := c.ShouldBindJSON(&beeUserQueue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUserQueue.UserId = &shopUserId
	if err := beeUserQueueService.UpdateBeeUserQueue(beeUserQueue, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeUserQueue 用id查询beeUserQueue表
// @Tags BeeUserQueue
// @Summary 用id查询beeUserQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeUserQueue true "用id查询beeUserQueue表"
// @Success 200 {object} response.Response{data=object{rebeeUserQueue=bee.BeeUserQueue},msg=string} "查询成功"
// @Router /beeUserQueue/findBeeUserQueue [get]
func (beeUserQueueApi *BeeUserQueueApi) FindBeeUserQueue(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeUserQueue, err := beeUserQueueService.GetBeeUserQueue(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeUserQueue, c)
	}
}

// GetBeeUserQueueList 分页获取beeUserQueue表列表
// @Tags BeeUserQueue
// @Summary 分页获取beeUserQueue表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUserQueueSearch true "分页获取beeUserQueue表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeUserQueue/getBeeUserQueueList [get]
func (beeUserQueueApi *BeeUserQueueApi) GetBeeUserQueueList(c *gin.Context) {
	var pageInfo beeReq.BeeUserQueueSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeUserQueueService.GetBeeUserQueueInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeUserQueuePublic 不需要鉴权的beeUserQueue表接口
// @Tags BeeUserQueue
// @Summary 不需要鉴权的beeUserQueue表接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUserQueueSearch true "分页获取beeUserQueue表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeUserQueue/getBeeUserQueuePublic [get]
func (beeUserQueueApi *BeeUserQueueApi) GetBeeUserQueuePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的beeUserQueue表接口信息",
	}, "获取成功", c)
}
