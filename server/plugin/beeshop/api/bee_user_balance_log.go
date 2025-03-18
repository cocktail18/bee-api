package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/service"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
	beeUtils "github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BeeUserBalanceLogApi struct{}

var beeUserBalanceLogService = service.ServiceGroupApp.BeeServiceGroup.BeeUserBalanceLogService

// CreateBeeUserBalanceLog 创建用户消费记录
// @Tags BeeUserBalanceLog
// @Summary 创建用户消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserBalanceLog true "创建用户消费记录"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeUserBalanceLog/createBeeUserBalanceLog [post]
func (beeUserBalanceLogApi *BeeUserBalanceLogApi) CreateBeeUserBalanceLog(c *gin.Context) {
	var beeUserBalanceLog bee.BeeUserBalanceLog
	err := c.ShouldBindJSON(&beeUserBalanceLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUserBalanceLog.UserId = &shopUserId

	if err := beeUserBalanceLogService.CreateBeeUserBalanceLog(&beeUserBalanceLog); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeUserBalanceLog 删除用户消费记录
// @Tags BeeUserBalanceLog
// @Summary 删除用户消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserBalanceLog true "删除用户消费记录"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeUserBalanceLog/deleteBeeUserBalanceLog [delete]
func (beeUserBalanceLogApi *BeeUserBalanceLogApi) DeleteBeeUserBalanceLog(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUserBalanceLogService.DeleteBeeUserBalanceLog(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeUserBalanceLogByIds 批量删除用户消费记录
// @Tags BeeUserBalanceLog
// @Summary 批量删除用户消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeUserBalanceLog/deleteBeeUserBalanceLogByIds [delete]
func (beeUserBalanceLogApi *BeeUserBalanceLogApi) DeleteBeeUserBalanceLogByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUserBalanceLogService.DeleteBeeUserBalanceLogByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeUserBalanceLog 更新用户消费记录
// @Tags BeeUserBalanceLog
// @Summary 更新用户消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserBalanceLog true "更新用户消费记录"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeUserBalanceLog/updateBeeUserBalanceLog [put]
func (beeUserBalanceLogApi *BeeUserBalanceLogApi) UpdateBeeUserBalanceLog(c *gin.Context) {
	var beeUserBalanceLog bee.BeeUserBalanceLog
	err := c.ShouldBindJSON(&beeUserBalanceLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUserBalanceLog.UserId = &shopUserId
	if err := beeUserBalanceLogService.UpdateBeeUserBalanceLog(beeUserBalanceLog, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeUserBalanceLog 用id查询用户消费记录
// @Tags BeeUserBalanceLog
// @Summary 用id查询用户消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeUserBalanceLog true "用id查询用户消费记录"
// @Success 200 {object} response.Response{data=object{rebeeUserBalanceLog=bee.BeeUserBalanceLog},msg=string} "查询成功"
// @Router /beeUserBalanceLog/findBeeUserBalanceLog [get]
func (beeUserBalanceLogApi *BeeUserBalanceLogApi) FindBeeUserBalanceLog(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeUserBalanceLog, err := beeUserBalanceLogService.GetBeeUserBalanceLog(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeUserBalanceLog, c)
	}
}

// GetBeeUserBalanceLogList 分页获取用户消费记录列表
// @Tags BeeUserBalanceLog
// @Summary 分页获取用户消费记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUserBalanceLogSearch true "分页获取用户消费记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeUserBalanceLog/getBeeUserBalanceLogList [get]
func (beeUserBalanceLogApi *BeeUserBalanceLogApi) GetBeeUserBalanceLogList(c *gin.Context) {
	var pageInfo beeReq.BeeUserBalanceLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	userId := beeUtils.GetUserID(c)
	if list, total, err := beeUserBalanceLogService.GetBeeUserBalanceLogInfoList(pageInfo, shopUserId, userId); err != nil {
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

// GetBeeUserBalanceLogPublic 不需要鉴权的用户消费记录接口
// @Tags BeeUserBalanceLog
// @Summary 不需要鉴权的用户消费记录接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUserBalanceLogSearch true "分页获取用户消费记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeUserBalanceLog/getBeeUserBalanceLogPublic [get]
func (beeUserBalanceLogApi *BeeUserBalanceLogApi) GetBeeUserBalanceLogPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的用户消费记录接口信息",
	}, "获取成功", c)
}
