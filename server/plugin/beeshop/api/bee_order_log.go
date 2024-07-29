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

type BeeOrderLogApi struct{}

var beeOrderLogService = service.ServiceGroupApp.BeeServiceGroup.BeeOrderLogService

// CreateBeeOrderLog 创建订单日志
// @Tags BeeOrderLog
// @Summary 创建订单日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderLog true "创建订单日志"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeOrderLog/createBeeOrderLog [post]
func (beeOrderLogApi *BeeOrderLogApi) CreateBeeOrderLog(c *gin.Context) {
	var beeOrderLog bee.BeeOrderLog
	err := c.ShouldBindJSON(&beeOrderLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeOrderLog.UserId = &shopUserId
	if err := beeOrderLogService.CreateBeeOrderLog(&beeOrderLog); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeOrderLog 删除订单日志
// @Tags BeeOrderLog
// @Summary 删除订单日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderLog true "删除订单日志"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeOrderLog/deleteBeeOrderLog [delete]
func (beeOrderLogApi *BeeOrderLogApi) DeleteBeeOrderLog(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderLogService.DeleteBeeOrderLog(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeOrderLogByIds 批量删除订单日志
// @Tags BeeOrderLog
// @Summary 批量删除订单日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeOrderLog/deleteBeeOrderLogByIds [delete]
func (beeOrderLogApi *BeeOrderLogApi) DeleteBeeOrderLogByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderLogService.DeleteBeeOrderLogByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeOrderLog 更新订单日志
// @Tags BeeOrderLog
// @Summary 更新订单日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderLog true "更新订单日志"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeOrderLog/updateBeeOrderLog [put]
func (beeOrderLogApi *BeeOrderLogApi) UpdateBeeOrderLog(c *gin.Context) {
	var beeOrderLog bee.BeeOrderLog
	err := c.ShouldBindJSON(&beeOrderLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeOrderLog.UserId = &shopUserId
	if err := beeOrderLogService.UpdateBeeOrderLog(beeOrderLog, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeOrderLog 用id查询订单日志
// @Tags BeeOrderLog
// @Summary 用id查询订单日志
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeOrderLog true "用id查询订单日志"
// @Success 200 {object} response.Response{data=object{rebeeOrderLog=bee.BeeOrderLog},msg=string} "查询成功"
// @Router /beeOrderLog/findBeeOrderLog [get]
func (beeOrderLogApi *BeeOrderLogApi) FindBeeOrderLog(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeOrderLog, err := beeOrderLogService.GetBeeOrderLog(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeOrderLog, c)
	}
}

// GetBeeOrderLogList 分页获取订单日志列表
// @Tags BeeOrderLog
// @Summary 分页获取订单日志列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeOrderLogSearch true "分页获取订单日志列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeOrderLog/getBeeOrderLogList [get]
func (beeOrderLogApi *BeeOrderLogApi) GetBeeOrderLogList(c *gin.Context) {
	var pageInfo beeReq.BeeOrderLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeOrderLogService.GetBeeOrderLogInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeOrderLogPublic 不需要鉴权的订单日志接口
// @Tags BeeOrderLog
// @Summary 不需要鉴权的订单日志接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeOrderLogSearch true "分页获取订单日志列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeOrderLog/getBeeOrderLogPublic [get]
func (beeOrderLogApi *BeeOrderLogApi) GetBeeOrderLogPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的订单日志接口信息",
	}, "获取成功", c)
}
