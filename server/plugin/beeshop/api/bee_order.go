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
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type BeeOrderApi struct{}

var beeOrderService = service.ServiceGroupApp.BeeServiceGroup.BeeOrderService

// CreateBeeOrder 创建用户订单
// @Tags BeeOrder
// @Summary 创建用户订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrder true "创建用户订单"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeOrder/createBeeOrder [post]
func (beeOrderApi *BeeOrderApi) CreateBeeOrder(c *gin.Context) {
	var beeOrder bee.BeeOrder
	err := c.ShouldBindJSON(&beeOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeOrder.UserId = &shopUserId
	if err := beeOrderService.CreateBeeOrder(&beeOrder); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeOrder 删除用户订单
// @Tags BeeOrder
// @Summary 删除用户订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrder true "删除用户订单"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeOrder/deleteBeeOrder [delete]
func (beeOrderApi *BeeOrderApi) DeleteBeeOrder(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderService.DeleteBeeOrder(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeOrderByIds 批量删除用户订单
// @Tags BeeOrder
// @Summary 批量删除用户订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeOrder/deleteBeeOrderByIds [delete]
func (beeOrderApi *BeeOrderApi) DeleteBeeOrderByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderService.DeleteBeeOrderByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeOrder 更新用户订单
// @Tags BeeOrder
// @Summary 更新用户订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrder true "更新用户订单"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeOrder/updateBeeOrder [put]
func (beeOrderApi *BeeOrderApi) UpdateBeeOrder(c *gin.Context) {
	var beeOrder bee.BeeOrder
	err := c.ShouldBindJSON(&beeOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeOrder.UserId = &shopUserId
	if err := beeOrderService.UpdateBeeOrder(beeOrder, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// UpdateBeeOrderExtJsonStr 更新用户订单 extJsonStr 字段
// @Tags BeeOrder
// @Summary 更新用户订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrder true "更新用户订单"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeOrder/UpdateBeeOrderExtJsonStr [put]
func (beeOrderApi *BeeOrderApi) UpdateBeeOrderExtJsonStr(c *gin.Context) {
	var beeOrder bee.BeeOrder
	err := c.ShouldBindJSON(&beeOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeOrder.UserId = &shopUserId
	if err := beeOrderService.UpdateBeeOrderExtJsonStr(beeOrder, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// UpdateBeeOrderStatus 更新用户订单 status 字段
// @Tags BeeOrder
// @Summary 更新用户订单 status
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrder true "更新用户订单 status"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeOrder/UpdateBeeOrderExtJsonStr [put]
func (beeOrderApi *BeeOrderApi) UpdateBeeOrderStatus(c *gin.Context) {
	var beeOrder bee.BeeOrder
	err := c.ShouldBindJSON(&beeOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeOrder.UserId = &shopUserId
	if err := beeOrderService.UpdateBeeOrderStatus(beeOrder, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeOrder 用id查询用户订单
// @Tags BeeOrder
// @Summary 用id查询用户订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeOrder true "用id查询用户订单"
// @Success 200 {object} response.Response{data=object{rebeeOrder=bee.BeeOrder},msg=string} "查询成功"
// @Router /beeOrder/findBeeOrder [get]
func (beeOrderApi *BeeOrderApi) FindBeeOrder(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeOrder, err := beeOrderService.GetBeeOrder(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeOrder, c)
	}
}

// GetBeeOrderList 分页获取用户订单列表
// @Tags BeeOrder
// @Summary 分页获取用户订单列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeOrderSearch true "分页获取用户订单列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeOrder/getBeeOrderList [get]
func (beeOrderApi *BeeOrderApi) GetBeeOrderList(c *gin.Context) {
	var pageInfo beeReq.BeeOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	loginUserId := beeUtils.GetUserID(c)

	if list, total, _, err := beeOrderService.GetBeeOrderInfoList(pageInfo, shopUserId, loginUserId); err != nil {
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

// GetBeeOrderPublic 不需要鉴权的用户订单接口
// @Tags BeeOrder
// @Summary 不需要鉴权的用户订单接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeOrderSearch true "分页获取用户订单列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeOrder/getBeeOrderPublic [get]
func (beeOrderApi *BeeOrderApi) GetBeeOrderPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的用户订单接口信息",
	}, "获取成功", c)
}

// MarkBeeOrderPaid 设置为已支付订单
// @Tags BeeOrder
// @Summary 设置为已支付订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "设置成功"
// @Router /beeOrder/markBeeOrderPaid [put]
func (beeOrderApi *BeeOrderApi) MarkBeeOrderPaid(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderService.MarkBeeOrderPaid(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

// MarkBeeOrderDone 设置为已完成订单
// @Tags BeeOrder
// @Summary 设置为已完成订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "设置成功"
// @Router /beeOrder/markBeeOrderPaid [put]
func (beeOrderApi *BeeOrderApi) MarkBeeOrderDone(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderService.MarkBeeOrderDone(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

// ShippedBeeOrder 发货订单
// @Tags BeeOrder
// @Summary 设置为已发货订单
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "设置成功"
// @Router /beeOrder/markBeeOrderPaid [post]
func (beeOrderApi *BeeOrderApi) ShippedBeeOrder(c *gin.Context) {
	var beeOrder bee.BeeOrder
	err := c.ShouldBindJSON(&beeOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderService.ShippedBeeOrder(cast.ToInt64(beeOrder.Id), shopUserId); err != nil {
		global.GVA_LOG.Error("设置失败!", zap.Error(err))
		response.FailWithMessage("设置失败", c)
	} else {
		response.OkWithMessage("设置成功", c)
	}
}

// ShippedBeeOrder 门店订单列表
// @Tags BeeOrder
// @Summary 门店订单列表查询
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "设置成功"
// @Router /beeOrder/orderList [get]
func (api *BeeOrderApi) OrderList(c *gin.Context) {
	var param beeReq.BeeOrderSearch

	if err := c.ShouldBindQuery(&param); err != nil {
		response.FailWithMessage("请求参数异常", c)
	} else {
		loginUserId := beeUtils.GetUserID(c)
		if list, total, sum, err := beeOrderService.GetBeeOrderInfoList(param, 100, loginUserId); err != nil {
			response.FailWithMessage("查询异常", c)
		} else {
			response.OkWithData(response.PageResult{
				List:  list,
				Total: total,
				Sum:   sum,
			}, c)
		}
	}
}
