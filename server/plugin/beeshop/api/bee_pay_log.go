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

type BeePayLogApi struct{}

var beePayLogService = service.ServiceGroupApp.BeeServiceGroup.BeePayLogService

// CreateBeePayLog 创建支付流水
// @Tags BeePayLog
// @Summary 创建支付流水
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeePayLog true "创建支付流水"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beePayLog/createBeePayLog [post]
func (beePayLogApi *BeePayLogApi) CreateBeePayLog(c *gin.Context) {
	var beePayLog bee.BeePayLog
	err := c.ShouldBindJSON(&beePayLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beePayLog.UserId = &shopUserId
	if err := beePayLogService.CreateBeePayLog(&beePayLog); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeePayLog 删除支付流水
// @Tags BeePayLog
// @Summary 删除支付流水
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeePayLog true "删除支付流水"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beePayLog/deleteBeePayLog [delete]
func (beePayLogApi *BeePayLogApi) DeleteBeePayLog(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beePayLogService.DeleteBeePayLog(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeePayLogByIds 批量删除支付流水
// @Tags BeePayLog
// @Summary 批量删除支付流水
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beePayLog/deleteBeePayLogByIds [delete]
func (beePayLogApi *BeePayLogApi) DeleteBeePayLogByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beePayLogService.DeleteBeePayLogByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeePayLog 更新支付流水
// @Tags BeePayLog
// @Summary 更新支付流水
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeePayLog true "更新支付流水"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beePayLog/updateBeePayLog [put]
func (beePayLogApi *BeePayLogApi) UpdateBeePayLog(c *gin.Context) {
	var beePayLog bee.BeePayLog
	err := c.ShouldBindJSON(&beePayLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beePayLog.UserId = &shopUserId
	if err := beePayLogService.UpdateBeePayLog(beePayLog, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeePayLog 用id查询支付流水
// @Tags BeePayLog
// @Summary 用id查询支付流水
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeePayLog true "用id查询支付流水"
// @Success 200 {object} response.Response{data=object{rebeePayLog=bee.BeePayLog},msg=string} "查询成功"
// @Router /beePayLog/findBeePayLog [get]
func (beePayLogApi *BeePayLogApi) FindBeePayLog(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeePayLog, err := beePayLogService.GetBeePayLog(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeePayLog, c)
	}
}

// GetBeePayLogList 分页获取支付流水列表
// @Tags BeePayLog
// @Summary 分页获取支付流水列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeePayLogSearch true "分页获取支付流水列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beePayLog/getBeePayLogList [get]
func (beePayLogApi *BeePayLogApi) GetBeePayLogList(c *gin.Context) {
	var pageInfo beeReq.BeePayLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beePayLogService.GetBeePayLogInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeePayTotal 分页获取支付流水列表
// @Tags BeePayLog
// @Summary 获取支付总额
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeePayLogSearch true "分页获取支付流水列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beePayLog/getBeePayTotal [get]
func (beePayLogApi *BeePayLogApi) GetBeePayTotal(c *gin.Context) {
	var pageInfo beeReq.BeePayLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if total, err := beePayLogService.GetBeePayTotal(pageInfo, shopUserId); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(map[string]interface{}{
			"total": total,
		}, c)
	}
}

// GetBeePayLogPublic 不需要鉴权的支付流水接口
// @Tags BeePayLog
// @Summary 不需要鉴权的支付流水接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeePayLogSearch true "分页获取支付流水列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beePayLog/getBeePayLogPublic [get]
func (beePayLogApi *BeePayLogApi) GetBeePayLogPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的支付流水接口信息",
	}, "获取成功", c)
}
