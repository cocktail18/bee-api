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

type BeeOrderPeisongLogApi struct{}

var beeOrderPeisongLogService = service.ServiceGroupApp.BeeServiceGroup.BeeOrderPeisongLogService

// CreateBeeOrderPeisongLog 创建beeOrderPeisongLog表
// @Tags BeeOrderPeisongLog
// @Summary 创建beeOrderPeisongLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderPeisongLog true "创建beeOrderPeisongLog表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeOrderPeisongLog/createBeeOrderPeisongLog [post]
func (beeOrderPeisongLogApi *BeeOrderPeisongLogApi) CreateBeeOrderPeisongLog(c *gin.Context) {
	var beeOrderPeisongLog bee.BeeOrderPeisongLog
	err := c.ShouldBindJSON(&beeOrderPeisongLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeOrderPeisongLog.UserId = &shopUserId
	if err := beeOrderPeisongLogService.CreateBeeOrderPeisongLog(&beeOrderPeisongLog); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeOrderPeisongLog 删除beeOrderPeisongLog表
// @Tags BeeOrderPeisongLog
// @Summary 删除beeOrderPeisongLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderPeisongLog true "删除beeOrderPeisongLog表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeOrderPeisongLog/deleteBeeOrderPeisongLog [delete]
func (beeOrderPeisongLogApi *BeeOrderPeisongLogApi) DeleteBeeOrderPeisongLog(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderPeisongLogService.DeleteBeeOrderPeisongLog(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeOrderPeisongLogByIds 批量删除beeOrderPeisongLog表
// @Tags BeeOrderPeisongLog
// @Summary 批量删除beeOrderPeisongLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeOrderPeisongLog/deleteBeeOrderPeisongLogByIds [delete]
func (beeOrderPeisongLogApi *BeeOrderPeisongLogApi) DeleteBeeOrderPeisongLogByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderPeisongLogService.DeleteBeeOrderPeisongLogByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeOrderPeisongLog 更新beeOrderPeisongLog表
// @Tags BeeOrderPeisongLog
// @Summary 更新beeOrderPeisongLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderPeisongLog true "更新beeOrderPeisongLog表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeOrderPeisongLog/updateBeeOrderPeisongLog [put]
func (beeOrderPeisongLogApi *BeeOrderPeisongLogApi) UpdateBeeOrderPeisongLog(c *gin.Context) {
	var beeOrderPeisongLog bee.BeeOrderPeisongLog
	err := c.ShouldBindJSON(&beeOrderPeisongLog)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeOrderPeisongLog.UserId = &shopUserId
	if err := beeOrderPeisongLogService.UpdateBeeOrderPeisongLog(beeOrderPeisongLog, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeOrderPeisongLog 用id查询beeOrderPeisongLog表
// @Tags BeeOrderPeisongLog
// @Summary 用id查询beeOrderPeisongLog表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeOrderPeisongLog true "用id查询beeOrderPeisongLog表"
// @Success 200 {object} response.Response{data=object{rebeeOrderPeisongLog=bee.BeeOrderPeisongLog},msg=string} "查询成功"
// @Router /beeOrderPeisongLog/findBeeOrderPeisongLog [get]
func (beeOrderPeisongLogApi *BeeOrderPeisongLogApi) FindBeeOrderPeisongLog(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeOrderPeisongLog, err := beeOrderPeisongLogService.GetBeeOrderPeisongLog(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeOrderPeisongLog, c)
	}
}

// GetBeeOrderPeisongLogList 分页获取beeOrderPeisongLog表列表
// @Tags BeeOrderPeisongLog
// @Summary 分页获取beeOrderPeisongLog表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeOrderPeisongLogSearch true "分页获取beeOrderPeisongLog表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeOrderPeisongLog/getBeeOrderPeisongLogList [get]
func (beeOrderPeisongLogApi *BeeOrderPeisongLogApi) GetBeeOrderPeisongLogList(c *gin.Context) {
	var pageInfo beeReq.BeeOrderPeisongLogSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeOrderPeisongLogService.GetBeeOrderPeisongLogInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeOrderPeisongLogPublic 不需要鉴权的beeOrderPeisongLog表接口
// @Tags BeeOrderPeisongLog
// @Summary 不需要鉴权的beeOrderPeisongLog表接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeOrderPeisongLogSearch true "分页获取beeOrderPeisongLog表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeOrderPeisongLog/getBeeOrderPeisongLogPublic [get]
func (beeOrderPeisongLogApi *BeeOrderPeisongLogApi) GetBeeOrderPeisongLogPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的beeOrderPeisongLog表接口信息",
	}, "获取成功", c)
}
