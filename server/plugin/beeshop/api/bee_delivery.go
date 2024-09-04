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

type BeeDeliveryApi struct{}

var beeDeliveryService = service.ServiceGroupApp.BeeServiceGroup.BeeDeliveryService

// CreateBeeDelivery 创建beeDelivery表
// @Tags BeeDelivery
// @Summary 创建beeDelivery表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeDelivery true "创建beeDelivery表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeDelivery/createBeeDelivery [post]
func (beeDeliveryApi *BeeDeliveryApi) CreateBeeDelivery(c *gin.Context) {
	var beeDelivery bee.BeeDelivery
	err := c.ShouldBindJSON(&beeDelivery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeDelivery.UserId = &shopUserId
	if err := beeDeliveryService.CreateBeeDelivery(&beeDelivery); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeDelivery 删除beeDelivery表
// @Tags BeeDelivery
// @Summary 删除beeDelivery表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeDelivery true "删除beeDelivery表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeDelivery/deleteBeeDelivery [delete]
func (beeDeliveryApi *BeeDeliveryApi) DeleteBeeDelivery(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeDeliveryService.DeleteBeeDelivery(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeDeliveryByIds 批量删除beeDelivery表
// @Tags BeeDelivery
// @Summary 批量删除beeDelivery表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeDelivery/deleteBeeDeliveryByIds [delete]
func (beeDeliveryApi *BeeDeliveryApi) DeleteBeeDeliveryByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeDeliveryService.DeleteBeeDeliveryByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeDelivery 更新beeDelivery表
// @Tags BeeDelivery
// @Summary 更新beeDelivery表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeDelivery true "更新beeDelivery表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeDelivery/updateBeeDelivery [put]
func (beeDeliveryApi *BeeDeliveryApi) UpdateBeeDelivery(c *gin.Context) {
	var beeDelivery bee.BeeDelivery
	err := c.ShouldBindJSON(&beeDelivery)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeDelivery.UserId = &shopUserId
	if err := beeDeliveryService.UpdateBeeDelivery(beeDelivery, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// BindYunlabaShop 绑定云喇叭商户
// @Tags BeeDelivery
// @Summary 绑定云喇叭商户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeDelivery true "更新beeDelivery表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeDelivery/bindYunlabaShop [put]
func (beeDeliveryApi *BeeDeliveryApi) BindYunlabaShop(c *gin.Context) {
	var bindReq beeReq.BeeDeliveryBindYunlabaShopReq
	err := c.ShouldBindJSON(&bindReq)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeDeliveryService.BindYunlabaShop(bindReq, shopUserId); err != nil {
		global.GVA_LOG.Error("绑定失败!", zap.Error(err))
		response.FailWithMessage("绑定失败"+err.Error(), c)
	} else {
		response.OkWithMessage("绑定成功", c)
	}
}

// FindBeeDelivery 用id查询beeDelivery表
// @Tags BeeDelivery
// @Summary 用id查询beeDelivery表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeDelivery true "用id查询beeDelivery表"
// @Success 200 {object} response.Response{data=object{rebeeDelivery=bee.BeeDelivery},msg=string} "查询成功"
// @Router /beeDelivery/findBeeDelivery [get]
func (beeDeliveryApi *BeeDeliveryApi) FindBeeDelivery(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeDelivery, err := beeDeliveryService.GetBeeDelivery(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeDelivery, c)
	}
}

// GetBeeDeliveryList 分页获取beeDelivery表列表
// @Tags BeeDelivery
// @Summary 分页获取beeDelivery表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeDeliverySearch true "分页获取beeDelivery表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeDelivery/getBeeDeliveryList [get]
func (beeDeliveryApi *BeeDeliveryApi) GetBeeDeliveryList(c *gin.Context) {
	var pageInfo beeReq.BeeDeliverySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeDeliveryService.GetBeeDeliveryInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeDeliveryPublic 不需要鉴权的beeDelivery表接口
// @Tags BeeDelivery
// @Summary 不需要鉴权的beeDelivery表接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeDeliverySearch true "分页获取beeDelivery表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeDelivery/getBeeDeliveryPublic [get]
func (beeDeliveryApi *BeeDeliveryApi) GetBeeDeliveryPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的beeDelivery表接口信息",
	}, "获取成功", c)
}
