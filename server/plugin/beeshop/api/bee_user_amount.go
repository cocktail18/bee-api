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

type BeeUserAmountApi struct{}

var beeUserAmountService = service.ServiceGroupApp.BeeServiceGroup.BeeUserAmountService

// CreateBeeUserAmount 创建用户资产
// @Tags BeeUserAmount
// @Summary 创建用户资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserAmount true "创建用户资产"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeUserAmount/createBeeUserAmount [post]
func (beeUserAmountApi *BeeUserAmountApi) CreateBeeUserAmount(c *gin.Context) {
	var beeUserAmount bee.BeeUserAmount
	err := c.ShouldBindJSON(&beeUserAmount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUserAmount.UserId = &shopUserId
	if err := beeUserAmountService.CreateBeeUserAmount(&beeUserAmount); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeUserAmount 删除用户资产
// @Tags BeeUserAmount
// @Summary 删除用户资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserAmount true "删除用户资产"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeUserAmount/deleteBeeUserAmount [delete]
func (beeUserAmountApi *BeeUserAmountApi) DeleteBeeUserAmount(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUserAmountService.DeleteBeeUserAmount(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeUserAmountByIds 批量删除用户资产
// @Tags BeeUserAmount
// @Summary 批量删除用户资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeUserAmount/deleteBeeUserAmountByIds [delete]
func (beeUserAmountApi *BeeUserAmountApi) DeleteBeeUserAmountByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUserAmountService.DeleteBeeUserAmountByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeUserAmount 更新用户资产
// @Tags BeeUserAmount
// @Summary 更新用户资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserAmount true "更新用户资产"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeUserAmount/updateBeeUserAmount [put]
func (beeUserAmountApi *BeeUserAmountApi) UpdateBeeUserAmount(c *gin.Context) {
	var beeUserAmount bee.BeeUserAmount
	err := c.ShouldBindJSON(&beeUserAmount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUserAmount.UserId = &shopUserId
	if err := beeUserAmountService.UpdateBeeUserAmount(beeUserAmount, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeUserAmount 用id查询用户资产
// @Tags BeeUserAmount
// @Summary 用id查询用户资产
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeUserAmount true "用id查询用户资产"
// @Success 200 {object} response.Response{data=object{rebeeUserAmount=bee.BeeUserAmount},msg=string} "查询成功"
// @Router /beeUserAmount/findBeeUserAmount [get]
func (beeUserAmountApi *BeeUserAmountApi) FindBeeUserAmount(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeUserAmount, err := beeUserAmountService.GetBeeUserAmount(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeUserAmount, c)
	}
}

// GetBeeUserAmountList 分页获取用户资产列表
// @Tags BeeUserAmount
// @Summary 分页获取用户资产列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUserAmountSearch true "分页获取用户资产列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeUserAmount/getBeeUserAmountList [get]
func (beeUserAmountApi *BeeUserAmountApi) GetBeeUserAmountList(c *gin.Context) {
	var pageInfo beeReq.BeeUserAmountSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeUserAmountService.GetBeeUserAmountInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeUserAmountPublic 不需要鉴权的用户资产接口
// @Tags BeeUserAmount
// @Summary 不需要鉴权的用户资产接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUserAmountSearch true "分页获取用户资产列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeUserAmount/getBeeUserAmountPublic [get]
func (beeUserAmountApi *BeeUserAmountApi) GetBeeUserAmountPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的用户资产接口信息",
	}, "获取成功", c)
}
