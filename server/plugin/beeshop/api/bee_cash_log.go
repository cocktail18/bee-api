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

type BeeCashLogApi struct{}

var beeCashLogService = service.ServiceGroupApp.BeeServiceGroup.BeeCashLogService

// CreateBeeCashLog 创建用户现金消费记录
// @Tags BeeCashLog
// @Summary 创建用户现金消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeCashLog true "创建用户现金消费记录"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeCashLog/createBeeCashLog [post]
func (beeCashLogApi *BeeCashLogApi) CreateBeeCashLog(c *gin.Context) {
	var beeCashLog bee.BeeCashLog
	err := c.ShouldBindJSON(&beeCashLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeCashLog.UserId = &shopUserId
	if err := beeCashLogService.CreateBeeCashLog(&beeCashLog); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeCashLog 删除用户现金消费记录
// @Tags BeeCashLog
// @Summary 删除用户现金消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeCashLog true "删除用户现金消费记录"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeCashLog/deleteBeeCashLog [delete]
func (beeCashLogApi *BeeCashLogApi) DeleteBeeCashLog(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeCashLogService.DeleteBeeCashLog(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeCashLogByIds 批量删除用户现金消费记录
// @Tags BeeCashLog
// @Summary 批量删除用户现金消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeCashLog/deleteBeeCashLogByIds [delete]
func (beeCashLogApi *BeeCashLogApi) DeleteBeeCashLogByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeCashLogService.DeleteBeeCashLogByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeCashLog 更新用户现金消费记录
// @Tags BeeCashLog
// @Summary 更新用户现金消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeCashLog true "更新用户现金消费记录"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeCashLog/updateBeeCashLog [put]
func (beeCashLogApi *BeeCashLogApi) UpdateBeeCashLog(c *gin.Context) {
	var beeCashLog bee.BeeCashLog
	err := c.ShouldBindJSON(&beeCashLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeCashLog.UserId = &shopUserId
	if err := beeCashLogService.UpdateBeeCashLog(beeCashLog, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeCashLog 用id查询用户现金消费记录
// @Tags BeeCashLog
// @Summary 用id查询用户现金消费记录
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeCashLog true "用id查询用户现金消费记录"
// @Success 200 {object} response.Response{data=object{rebeeCashLog=bee.BeeCashLog},msg=string} "查询成功"
// @Router /beeCashLog/findBeeCashLog [get]
func (beeCashLogApi *BeeCashLogApi) FindBeeCashLog(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeCashLog, err := beeCashLogService.GetBeeCashLog(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeCashLog, c)
	}
}

// GetBeeCashLogList 分页获取用户现金消费记录列表
// @Tags BeeCashLog
// @Summary 分页获取用户现金消费记录列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeCashLogSearch true "分页获取用户现金消费记录列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeCashLog/getBeeCashLogList [get]
func (beeCashLogApi *BeeCashLogApi) GetBeeCashLogList(c *gin.Context) {
	var pageInfo beeReq.BeeCashLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeCashLogService.GetBeeCashLogInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeCashLogPublic 不需要鉴权的用户现金消费记录接口
// @Tags BeeCashLog
// @Summary 不需要鉴权的用户现金消费记录接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeCashLogSearch true "分页获取用户现金消费记录列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeCashLog/getBeeCashLogPublic [get]
func (beeCashLogApi *BeeCashLogApi) GetBeeCashLogPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的用户现金消费记录接口信息",
	}, "获取成功", c)
}
