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

type BeeQueueApi struct{}

var beeQueueService = service.ServiceGroupApp.BeeServiceGroup.BeeQueueService

// CreateBeeQueue 创建beeQueue表
// @Tags BeeQueue
// @Summary 创建beeQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeQueue true "创建beeQueue表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeQueue/createBeeQueue [post]
func (beeQueueApi *BeeQueueApi) CreateBeeQueue(c *gin.Context) {
	var beeQueue bee.BeeQueue
	err := c.ShouldBindJSON(&beeQueue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeQueue.UserId = &shopUserId
	if err := beeQueueService.CreateBeeQueue(&beeQueue); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeQueue 删除beeQueue表
// @Tags BeeQueue
// @Summary 删除beeQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeQueue true "删除beeQueue表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeQueue/deleteBeeQueue [delete]
func (beeQueueApi *BeeQueueApi) DeleteBeeQueue(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeQueueService.DeleteBeeQueue(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeQueueByIds 批量删除beeQueue表
// @Tags BeeQueue
// @Summary 批量删除beeQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeQueue/deleteBeeQueueByIds [delete]
func (beeQueueApi *BeeQueueApi) DeleteBeeQueueByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeQueueService.DeleteBeeQueueByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeQueue 更新beeQueue表
// @Tags BeeQueue
// @Summary 更新beeQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeQueue true "更新beeQueue表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeQueue/updateBeeQueue [put]
func (beeQueueApi *BeeQueueApi) UpdateBeeQueue(c *gin.Context) {
	var beeQueue bee.BeeQueue
	err := c.ShouldBindJSON(&beeQueue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeQueue.UserId = &shopUserId
	if err := beeQueueService.UpdateBeeQueue(beeQueue, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// CallBeeQueue 叫号
// @Tags BeeQueue
// @Summary 更新beeQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeQueue true "更新beeQueue表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeQueue/callBeeQueue [put]
func (beeQueueApi *BeeQueueApi) CallBeeQueue(c *gin.Context) {
	var beeQueue bee.BeeQueue
	err := c.ShouldBindJSON(&beeQueue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeQueue.UserId = &shopUserId
	if err := beeQueueService.CallBeeQueue(beeQueue, shopUserId); err != nil {
		global.GVA_LOG.Error("叫号失败!", zap.Error(err))
		response.FailWithMessage("叫号失败:"+err.Error(), c)
	} else {
		response.OkWithMessage("叫号成功", c)
	}
}

// ReCallBeeQueue 重新叫号
// @Tags BeeQueue
// @Summary 更新beeQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeQueue true "更新beeQueue表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeQueue/reCallBeeQueue [put]
func (beeQueueApi *BeeQueueApi) ReCallBeeQueue(c *gin.Context) {
	var beeQueue bee.BeeQueue
	err := c.ShouldBindJSON(&beeQueue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeQueue.UserId = &shopUserId
	if err := beeQueueService.ReCallBeeQueue(beeQueue, shopUserId); err != nil {
		global.GVA_LOG.Error("叫号失败!", zap.Error(err))
		response.FailWithMessage("叫号失败:"+err.Error(), c)
	} else {
		response.OkWithMessage("叫号成功", c)
	}
}

// PassBeeQueue 过号
// @Tags BeeQueue
// @Summary 更新beeQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeQueue true "更新beeQueue表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeQueue/passBeeQueue [put]
func (beeQueueApi *BeeQueueApi) PassBeeQueue(c *gin.Context) {
	var beeQueue bee.BeeQueue
	err := c.ShouldBindJSON(&beeQueue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeQueue.UserId = &shopUserId
	if err := beeQueueService.PassBeeQueue(beeQueue, shopUserId); err != nil {
		global.GVA_LOG.Error("过号失败!", zap.Error(err))
		response.FailWithMessage("过号失败:"+err.Error(), c)
	} else {
		response.OkWithMessage("过号成功", c)
	}
}

// NextBeeQueue 下一位
// @Tags BeeQueue
// @Summary 更新beeQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeQueue true "更新beeQueue表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeQueue/nextBeeQueue [put]
func (beeQueueApi *BeeQueueApi) NextBeeQueue(c *gin.Context) {
	var beeQueue bee.BeeQueue
	err := c.ShouldBindJSON(&beeQueue)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeQueue.UserId = &shopUserId
	if err := beeQueueService.NextBeeQueue(beeQueue, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败:"+err.Error(), c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeQueue 用id查询beeQueue表
// @Tags BeeQueue
// @Summary 用id查询beeQueue表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeQueue true "用id查询beeQueue表"
// @Success 200 {object} response.Response{data=object{rebeeQueue=bee.BeeQueue},msg=string} "查询成功"
// @Router /beeQueue/findBeeQueue [get]
func (beeQueueApi *BeeQueueApi) FindBeeQueue(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeQueue, err := beeQueueService.GetBeeQueue(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeQueue, c)
	}
}

// GetBeeQueueList 分页获取beeQueue表列表
// @Tags BeeQueue
// @Summary 分页获取beeQueue表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeQueueSearch true "分页获取beeQueue表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeQueue/getBeeQueueList [get]
func (beeQueueApi *BeeQueueApi) GetBeeQueueList(c *gin.Context) {
	var pageInfo beeReq.BeeQueueSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeQueueService.GetBeeQueueInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeQueuePublic 不需要鉴权的beeQueue表接口
// @Tags BeeQueue
// @Summary 不需要鉴权的beeQueue表接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeQueueSearch true "分页获取beeQueue表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeQueue/getBeeQueuePublic [get]
func (beeQueueApi *BeeQueueApi) GetBeeQueuePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的beeQueue表接口信息",
	}, "获取成功", c)
}
