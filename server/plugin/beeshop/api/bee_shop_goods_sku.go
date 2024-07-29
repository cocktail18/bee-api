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

type BeeShopGoodsSkuApi struct{}

var beeShopGoodsSkuService = service.ServiceGroupApp.BeeServiceGroup.BeeShopGoodsSkuService

// CreateBeeShopGoodsSku 创建商品sku
// @Tags BeeShopGoodsSku
// @Summary 创建商品sku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsSku true "创建商品sku"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeShopGoodsSku/createBeeShopGoodsSku [post]
func (beeShopGoodsSkuApi *BeeShopGoodsSkuApi) CreateBeeShopGoodsSku(c *gin.Context) {
	var beeShopGoodsSku bee.BeeShopGoodsSku
	err := c.ShouldBindJSON(&beeShopGoodsSku)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopGoodsSku.UserId = &shopUserId
	if err := beeShopGoodsSkuService.CreateBeeShopGoodsSku(&beeShopGoodsSku); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeShopGoodsSku 删除商品sku
// @Tags BeeShopGoodsSku
// @Summary 删除商品sku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsSku true "删除商品sku"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeShopGoodsSku/deleteBeeShopGoodsSku [delete]
func (beeShopGoodsSkuApi *BeeShopGoodsSkuApi) DeleteBeeShopGoodsSku(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopGoodsSkuService.DeleteBeeShopGoodsSku(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeShopGoodsSkuByIds 批量删除商品sku
// @Tags BeeShopGoodsSku
// @Summary 批量删除商品sku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeShopGoodsSku/deleteBeeShopGoodsSkuByIds [delete]
func (beeShopGoodsSkuApi *BeeShopGoodsSkuApi) DeleteBeeShopGoodsSkuByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopGoodsSkuService.DeleteBeeShopGoodsSkuByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeShopGoodsSku 更新商品sku
// @Tags BeeShopGoodsSku
// @Summary 更新商品sku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsSku true "更新商品sku"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeShopGoodsSku/updateBeeShopGoodsSku [put]
func (beeShopGoodsSkuApi *BeeShopGoodsSkuApi) UpdateBeeShopGoodsSku(c *gin.Context) {
	var beeShopGoodsSku bee.BeeShopGoodsSku
	err := c.ShouldBindJSON(&beeShopGoodsSku)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopGoodsSku.UserId = &shopUserId
	if err := beeShopGoodsSkuService.UpdateBeeShopGoodsSku(beeShopGoodsSku, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeShopGoodsSku 用id查询商品sku
// @Tags BeeShopGoodsSku
// @Summary 用id查询商品sku
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeShopGoodsSku true "用id查询商品sku"
// @Success 200 {object} response.Response{data=object{rebeeShopGoodsSku=bee.BeeShopGoodsSku},msg=string} "查询成功"
// @Router /beeShopGoodsSku/findBeeShopGoodsSku [get]
func (beeShopGoodsSkuApi *BeeShopGoodsSkuApi) FindBeeShopGoodsSku(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeShopGoodsSku, err := beeShopGoodsSkuService.GetBeeShopGoodsSku(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeShopGoodsSku, c)
	}
}

// GetBeeShopGoodsSkuList 分页获取商品sku列表
// @Tags BeeShopGoodsSku
// @Summary 分页获取商品sku列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopGoodsSkuSearch true "分页获取商品sku列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeShopGoodsSku/getBeeShopGoodsSkuList [get]
func (beeShopGoodsSkuApi *BeeShopGoodsSkuApi) GetBeeShopGoodsSkuList(c *gin.Context) {
	var pageInfo beeReq.BeeShopGoodsSkuSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeShopGoodsSkuService.GetBeeShopGoodsSkuInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeShopGoodsSkuPublic 不需要鉴权的商品sku接口
// @Tags BeeShopGoodsSku
// @Summary 不需要鉴权的商品sku接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopGoodsSkuSearch true "分页获取商品sku列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeShopGoodsSku/getBeeShopGoodsSkuPublic [get]
func (beeShopGoodsSkuApi *BeeShopGoodsSkuApi) GetBeeShopGoodsSkuPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的商品sku接口信息",
	}, "获取成功", c)
}
