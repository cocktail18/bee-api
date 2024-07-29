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

type BeeShopGoodsApi struct{}

var beeShopGoodsService = service.ServiceGroupApp.BeeServiceGroup.BeeShopGoodsService

// CreateBeeShopGoods 创建商品表
// @Tags BeeShopGoods
// @Summary 创建商品表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoods true "创建商品表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeShopGoods/createBeeShopGoods [post]
func (beeShopGoodsApi *BeeShopGoodsApi) CreateBeeShopGoods(c *gin.Context) {
	var beeShopGoods bee.BeeShopGoods
	err := c.ShouldBindJSON(&beeShopGoods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopGoods.UserId = &shopUserId
	if err := beeShopGoodsService.CreateBeeShopGoods(&beeShopGoods); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeShopGoods 删除商品表
// @Tags BeeShopGoods
// @Summary 删除商品表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoods true "删除商品表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeShopGoods/deleteBeeShopGoods [delete]
func (beeShopGoodsApi *BeeShopGoodsApi) DeleteBeeShopGoods(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopGoodsService.DeleteBeeShopGoods(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeShopGoodsByIds 批量删除商品表
// @Tags BeeShopGoods
// @Summary 批量删除商品表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeShopGoods/deleteBeeShopGoodsByIds [delete]
func (beeShopGoodsApi *BeeShopGoodsApi) DeleteBeeShopGoodsByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopGoodsService.DeleteBeeShopGoodsByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeShopGoods 更新商品表
// @Tags BeeShopGoods
// @Summary 更新商品表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoods true "更新商品表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeShopGoods/updateBeeShopGoods [put]
func (beeShopGoodsApi *BeeShopGoodsApi) UpdateBeeShopGoods(c *gin.Context) {
	var beeShopGoods bee.BeeShopGoods
	err := c.ShouldBindJSON(&beeShopGoods)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopGoods.UserId = &shopUserId
	if err := beeShopGoodsService.UpdateBeeShopGoods(beeShopGoods, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeShopGoods 用id查询商品表
// @Tags BeeShopGoods
// @Summary 用id查询商品表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeShopGoods true "用id查询商品表"
// @Success 200 {object} response.Response{data=object{rebeeShopGoods=bee.BeeShopGoods},msg=string} "查询成功"
// @Router /beeShopGoods/findBeeShopGoods [get]
func (beeShopGoodsApi *BeeShopGoodsApi) FindBeeShopGoods(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeShopGoods, err := beeShopGoodsService.GetBeeShopGoods(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeShopGoods, c)
	}
}

// GetBeeShopGoodsList 分页获取商品表列表
// @Tags BeeShopGoods
// @Summary 分页获取商品表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopGoodsSearch true "分页获取商品表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeShopGoods/getBeeShopGoodsList [get]
func (beeShopGoodsApi *BeeShopGoodsApi) GetBeeShopGoodsList(c *gin.Context) {
	var pageInfo beeReq.BeeShopGoodsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeShopGoodsService.GetBeeShopGoodsInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeShopGoodsPublic 不需要鉴权的商品表接口
// @Tags BeeShopGoods
// @Summary 不需要鉴权的商品表接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopGoodsSearch true "分页获取商品表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeShopGoods/getBeeShopGoodsPublic [get]
func (beeShopGoodsApi *BeeShopGoodsApi) GetBeeShopGoodsPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的商品表接口信息",
	}, "获取成功", c)
}
