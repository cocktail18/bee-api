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

type BeeShopGoodsCategoryApi struct{}

var beeShopGoodsCategoryService = service.ServiceGroupApp.BeeServiceGroup.BeeShopGoodsCategoryService

// CreateBeeShopGoodsCategory 创建商品分类
// @Tags BeeShopGoodsCategory
// @Summary 创建商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsCategory true "创建商品分类"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeShopGoodsCategory/createBeeShopGoodsCategory [post]
func (beeShopGoodsCategoryApi *BeeShopGoodsCategoryApi) CreateBeeShopGoodsCategory(c *gin.Context) {
	var beeShopGoodsCategory bee.BeeShopGoodsCategory
	err := c.ShouldBindJSON(&beeShopGoodsCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopGoodsCategory.UserId = &shopUserId
	if err := beeShopGoodsCategoryService.CreateBeeShopGoodsCategory(&beeShopGoodsCategory); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeShopGoodsCategory 删除商品分类
// @Tags BeeShopGoodsCategory
// @Summary 删除商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsCategory true "删除商品分类"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeShopGoodsCategory/deleteBeeShopGoodsCategory [delete]
func (beeShopGoodsCategoryApi *BeeShopGoodsCategoryApi) DeleteBeeShopGoodsCategory(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopGoodsCategoryService.DeleteBeeShopGoodsCategory(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeShopGoodsCategoryByIds 批量删除商品分类
// @Tags BeeShopGoodsCategory
// @Summary 批量删除商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeShopGoodsCategory/deleteBeeShopGoodsCategoryByIds [delete]
func (beeShopGoodsCategoryApi *BeeShopGoodsCategoryApi) DeleteBeeShopGoodsCategoryByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopGoodsCategoryService.DeleteBeeShopGoodsCategoryByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeShopGoodsCategory 更新商品分类
// @Tags BeeShopGoodsCategory
// @Summary 更新商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsCategory true "更新商品分类"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeShopGoodsCategory/updateBeeShopGoodsCategory [put]
func (beeShopGoodsCategoryApi *BeeShopGoodsCategoryApi) UpdateBeeShopGoodsCategory(c *gin.Context) {
	var beeShopGoodsCategory bee.BeeShopGoodsCategory
	err := c.ShouldBindJSON(&beeShopGoodsCategory)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopGoodsCategory.UserId = &shopUserId
	if err := beeShopGoodsCategoryService.UpdateBeeShopGoodsCategory(beeShopGoodsCategory, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeShopGoodsCategory 用id查询商品分类
// @Tags BeeShopGoodsCategory
// @Summary 用id查询商品分类
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeShopGoodsCategory true "用id查询商品分类"
// @Success 200 {object} response.Response{data=object{rebeeShopGoodsCategory=bee.BeeShopGoodsCategory},msg=string} "查询成功"
// @Router /beeShopGoodsCategory/findBeeShopGoodsCategory [get]
func (beeShopGoodsCategoryApi *BeeShopGoodsCategoryApi) FindBeeShopGoodsCategory(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeShopGoodsCategory, err := beeShopGoodsCategoryService.GetBeeShopGoodsCategory(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeShopGoodsCategory, c)
	}
}

// GetBeeShopGoodsCategoryList 分页获取商品分类列表
// @Tags BeeShopGoodsCategory
// @Summary 分页获取商品分类列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopGoodsCategorySearch true "分页获取商品分类列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeShopGoodsCategory/getBeeShopGoodsCategoryList [get]
func (beeShopGoodsCategoryApi *BeeShopGoodsCategoryApi) GetBeeShopGoodsCategoryList(c *gin.Context) {
	var pageInfo beeReq.BeeShopGoodsCategorySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeShopGoodsCategoryService.GetBeeShopGoodsCategoryInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeShopGoodsCategoryPublic 不需要鉴权的商品分类接口
// @Tags BeeShopGoodsCategory
// @Summary 不需要鉴权的商品分类接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopGoodsCategorySearch true "分页获取商品分类列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeShopGoodsCategory/getBeeShopGoodsCategoryPublic [get]
func (beeShopGoodsCategoryApi *BeeShopGoodsCategoryApi) GetBeeShopGoodsCategoryPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的商品分类接口信息",
	}, "获取成功", c)
}
