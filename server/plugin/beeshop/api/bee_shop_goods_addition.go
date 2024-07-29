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

type BeeShopGoodsAdditionApi struct{}

var beeShopGoodsAdditionService = service.ServiceGroupApp.BeeServiceGroup.BeeShopGoodsAdditionService

// CreateBeeShopGoodsAddition 创建商品额外信息
// @Tags BeeShopGoodsAddition
// @Summary 创建商品额外信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsAddition true "创建商品额外信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeShopGoodsAddition/createBeeShopGoodsAddition [post]
func (beeShopGoodsAdditionApi *BeeShopGoodsAdditionApi) CreateBeeShopGoodsAddition(c *gin.Context) {
	var beeShopGoodsAddition bee.BeeShopGoodsAddition
	err := c.ShouldBindJSON(&beeShopGoodsAddition)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopGoodsAddition.UserId = &shopUserId
	if err := beeShopGoodsAdditionService.CreateBeeShopGoodsAddition(&beeShopGoodsAddition); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeShopGoodsAddition 删除商品额外信息
// @Tags BeeShopGoodsAddition
// @Summary 删除商品额外信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsAddition true "删除商品额外信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeShopGoodsAddition/deleteBeeShopGoodsAddition [delete]
func (beeShopGoodsAdditionApi *BeeShopGoodsAdditionApi) DeleteBeeShopGoodsAddition(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopGoodsAdditionService.DeleteBeeShopGoodsAddition(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeShopGoodsAdditionByIds 批量删除商品额外信息
// @Tags BeeShopGoodsAddition
// @Summary 批量删除商品额外信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeShopGoodsAddition/deleteBeeShopGoodsAdditionByIds [delete]
func (beeShopGoodsAdditionApi *BeeShopGoodsAdditionApi) DeleteBeeShopGoodsAdditionByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopGoodsAdditionService.DeleteBeeShopGoodsAdditionByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeShopGoodsAddition 更新商品额外信息
// @Tags BeeShopGoodsAddition
// @Summary 更新商品额外信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsAddition true "更新商品额外信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeShopGoodsAddition/updateBeeShopGoodsAddition [put]
func (beeShopGoodsAdditionApi *BeeShopGoodsAdditionApi) UpdateBeeShopGoodsAddition(c *gin.Context) {
	var beeShopGoodsAddition bee.BeeShopGoodsAddition
	err := c.ShouldBindJSON(&beeShopGoodsAddition)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopGoodsAddition.UserId = &shopUserId
	if err := beeShopGoodsAdditionService.UpdateBeeShopGoodsAddition(beeShopGoodsAddition, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeShopGoodsAddition 用id查询商品额外信息
// @Tags BeeShopGoodsAddition
// @Summary 用id查询商品额外信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeShopGoodsAddition true "用id查询商品额外信息"
// @Success 200 {object} response.Response{data=object{rebeeShopGoodsAddition=bee.BeeShopGoodsAddition},msg=string} "查询成功"
// @Router /beeShopGoodsAddition/findBeeShopGoodsAddition [get]
func (beeShopGoodsAdditionApi *BeeShopGoodsAdditionApi) FindBeeShopGoodsAddition(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeShopGoodsAddition, err := beeShopGoodsAdditionService.GetBeeShopGoodsAddition(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeShopGoodsAddition, c)
	}
}

// GetBeeShopGoodsAdditionList 分页获取商品额外信息列表
// @Tags BeeShopGoodsAddition
// @Summary 分页获取商品额外信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopGoodsAdditionSearch true "分页获取商品额外信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeShopGoodsAddition/getBeeShopGoodsAdditionList [get]
func (beeShopGoodsAdditionApi *BeeShopGoodsAdditionApi) GetBeeShopGoodsAdditionList(c *gin.Context) {
	var pageInfo beeReq.BeeShopGoodsAdditionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeShopGoodsAdditionService.GetBeeShopGoodsAdditionInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeShopGoodsAdditionPublic 不需要鉴权的商品额外信息接口
// @Tags BeeShopGoodsAddition
// @Summary 不需要鉴权的商品额外信息接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopGoodsAdditionSearch true "分页获取商品额外信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeShopGoodsAddition/getBeeShopGoodsAdditionPublic [get]
func (beeShopGoodsAdditionApi *BeeShopGoodsAdditionApi) GetBeeShopGoodsAdditionPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的商品额外信息接口信息",
	}, "获取成功", c)
}
