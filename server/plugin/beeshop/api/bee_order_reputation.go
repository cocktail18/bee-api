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

type BeeOrderReputationApi struct{}

var beeOrderReputationService = service.ServiceGroupApp.BeeServiceGroup.BeeOrderReputationService

// CreateBeeOrderReputation 创建商品评价
// @Tags BeeOrderReputation
// @Summary 创建商品评价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderReputation true "创建商品评价"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeOrderReputation/createBeeOrderReputation [post]
func (beeOrderReputationApi *BeeOrderReputationApi) CreateBeeOrderReputation(c *gin.Context) {
	var beeOrderReputation bee.BeeOrderReputation
	err := c.ShouldBindJSON(&beeOrderReputation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeOrderReputation.UserId = &shopUserId
	if err := beeOrderReputationService.CreateBeeOrderReputation(&beeOrderReputation); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeOrderReputation 删除商品评价
// @Tags BeeOrderReputation
// @Summary 删除商品评价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderReputation true "删除商品评价"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeOrderReputation/deleteBeeOrderReputation [delete]
func (beeOrderReputationApi *BeeOrderReputationApi) DeleteBeeOrderReputation(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderReputationService.DeleteBeeOrderReputation(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeOrderReputationByIds 批量删除商品评价
// @Tags BeeOrderReputation
// @Summary 批量删除商品评价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeOrderReputation/deleteBeeOrderReputationByIds [delete]
func (beeOrderReputationApi *BeeOrderReputationApi) DeleteBeeOrderReputationByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderReputationService.DeleteBeeOrderReputationByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeOrderReputation 更新商品评价
// @Tags BeeOrderReputation
// @Summary 更新商品评价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderReputation true "更新商品评价"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeOrderReputation/updateBeeOrderReputation [put]
func (beeOrderReputationApi *BeeOrderReputationApi) UpdateBeeOrderReputation(c *gin.Context) {
	var beeOrderReputation bee.BeeOrderReputation
	err := c.ShouldBindJSON(&beeOrderReputation)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeOrderReputation.UserId = &shopUserId
	if err := beeOrderReputationService.UpdateBeeOrderReputation(beeOrderReputation, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeOrderReputation 用id查询商品评价
// @Tags BeeOrderReputation
// @Summary 用id查询商品评价
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeOrderReputation true "用id查询商品评价"
// @Success 200 {object} response.Response{data=object{rebeeOrderReputation=bee.BeeOrderReputation},msg=string} "查询成功"
// @Router /beeOrderReputation/findBeeOrderReputation [get]
func (beeOrderReputationApi *BeeOrderReputationApi) FindBeeOrderReputation(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeOrderReputation, err := beeOrderReputationService.GetBeeOrderReputation(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeOrderReputation, c)
	}
}

// GetBeeOrderReputationList 分页获取商品评价列表
// @Tags BeeOrderReputation
// @Summary 分页获取商品评价列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeOrderReputationSearch true "分页获取商品评价列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeOrderReputation/getBeeOrderReputationList [get]
func (beeOrderReputationApi *BeeOrderReputationApi) GetBeeOrderReputationList(c *gin.Context) {
	var pageInfo beeReq.BeeOrderReputationSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeOrderReputationService.GetBeeOrderReputationInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeOrderReputationPublic 不需要鉴权的商品评价接口
// @Tags BeeOrderReputation
// @Summary 不需要鉴权的商品评价接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeOrderReputationSearch true "分页获取商品评价列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeOrderReputation/getBeeOrderReputationPublic [get]
func (beeOrderReputationApi *BeeOrderReputationApi) GetBeeOrderReputationPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的商品评价接口信息",
	}, "获取成功", c)
}
