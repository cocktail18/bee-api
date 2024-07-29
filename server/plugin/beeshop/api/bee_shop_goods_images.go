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

type BeeShopGoodsImagesApi struct{}

var beeShopGoodsImagesService = service.ServiceGroupApp.BeeServiceGroup.BeeShopGoodsImagesService

// CreateBeeShopGoodsImages 创建商品图
// @Tags BeeShopGoodsImages
// @Summary 创建商品图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsImages true "创建商品图"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeShopGoodsImages/createBeeShopGoodsImages [post]
func (beeShopGoodsImagesApi *BeeShopGoodsImagesApi) CreateBeeShopGoodsImages(c *gin.Context) {
	var beeShopGoodsImages bee.BeeShopGoodsImages
	err := c.ShouldBindJSON(&beeShopGoodsImages)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopGoodsImages.UserId = &shopUserId
	if err := beeShopGoodsImagesService.CreateBeeShopGoodsImages(&beeShopGoodsImages); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeShopGoodsImages 删除商品图
// @Tags BeeShopGoodsImages
// @Summary 删除商品图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsImages true "删除商品图"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeShopGoodsImages/deleteBeeShopGoodsImages [delete]
func (beeShopGoodsImagesApi *BeeShopGoodsImagesApi) DeleteBeeShopGoodsImages(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopGoodsImagesService.DeleteBeeShopGoodsImages(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeShopGoodsImagesByIds 批量删除商品图
// @Tags BeeShopGoodsImages
// @Summary 批量删除商品图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeShopGoodsImages/deleteBeeShopGoodsImagesByIds [delete]
func (beeShopGoodsImagesApi *BeeShopGoodsImagesApi) DeleteBeeShopGoodsImagesByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopGoodsImagesService.DeleteBeeShopGoodsImagesByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeShopGoodsImages 更新商品图
// @Tags BeeShopGoodsImages
// @Summary 更新商品图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsImages true "更新商品图"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeShopGoodsImages/updateBeeShopGoodsImages [put]
func (beeShopGoodsImagesApi *BeeShopGoodsImagesApi) UpdateBeeShopGoodsImages(c *gin.Context) {
	var beeShopGoodsImages bee.BeeShopGoodsImages
	err := c.ShouldBindJSON(&beeShopGoodsImages)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopGoodsImages.UserId = &shopUserId
	if err := beeShopGoodsImagesService.UpdateBeeShopGoodsImages(beeShopGoodsImages, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeShopGoodsImages 用id查询商品图
// @Tags BeeShopGoodsImages
// @Summary 用id查询商品图
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeShopGoodsImages true "用id查询商品图"
// @Success 200 {object} response.Response{data=object{rebeeShopGoodsImages=bee.BeeShopGoodsImages},msg=string} "查询成功"
// @Router /beeShopGoodsImages/findBeeShopGoodsImages [get]
func (beeShopGoodsImagesApi *BeeShopGoodsImagesApi) FindBeeShopGoodsImages(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeShopGoodsImages, err := beeShopGoodsImagesService.GetBeeShopGoodsImages(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeShopGoodsImages, c)
	}
}

// GetBeeShopGoodsImagesList 分页获取商品图列表
// @Tags BeeShopGoodsImages
// @Summary 分页获取商品图列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopGoodsImagesSearch true "分页获取商品图列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeShopGoodsImages/getBeeShopGoodsImagesList [get]
func (beeShopGoodsImagesApi *BeeShopGoodsImagesApi) GetBeeShopGoodsImagesList(c *gin.Context) {
	var pageInfo beeReq.BeeShopGoodsImagesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeShopGoodsImagesService.GetBeeShopGoodsImagesInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeShopGoodsImagesPublic 不需要鉴权的商品图接口
// @Tags BeeShopGoodsImages
// @Summary 不需要鉴权的商品图接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopGoodsImagesSearch true "分页获取商品图列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeShopGoodsImages/getBeeShopGoodsImagesPublic [get]
func (beeShopGoodsImagesApi *BeeShopGoodsImagesApi) GetBeeShopGoodsImagesPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的商品图接口信息",
	}, "获取成功", c)
}
