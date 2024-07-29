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

type BeeShopGoodsContentApi struct{}

var beeShopGoodsContentService = service.ServiceGroupApp.BeeServiceGroup.BeeShopGoodsContentService

// CreateBeeShopGoodsContent 创建商品详情描述
// @Tags BeeShopGoodsContent
// @Summary 创建商品详情描述
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsContent true "创建商品详情描述"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeShopGoodsContent/createBeeShopGoodsContent [post]
func (beeShopGoodsContentApi *BeeShopGoodsContentApi) CreateBeeShopGoodsContent(c *gin.Context) {
	var beeShopGoodsContent bee.BeeShopGoodsContent
	err := c.ShouldBindJSON(&beeShopGoodsContent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopGoodsContent.UserId = &shopUserId
	if err := beeShopGoodsContentService.CreateBeeShopGoodsContent(&beeShopGoodsContent); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeShopGoodsContent 删除商品详情描述
// @Tags BeeShopGoodsContent
// @Summary 删除商品详情描述
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsContent true "删除商品详情描述"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeShopGoodsContent/deleteBeeShopGoodsContent [delete]
func (beeShopGoodsContentApi *BeeShopGoodsContentApi) DeleteBeeShopGoodsContent(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopGoodsContentService.DeleteBeeShopGoodsContent(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeShopGoodsContentByIds 批量删除商品详情描述
// @Tags BeeShopGoodsContent
// @Summary 批量删除商品详情描述
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeShopGoodsContent/deleteBeeShopGoodsContentByIds [delete]
func (beeShopGoodsContentApi *BeeShopGoodsContentApi) DeleteBeeShopGoodsContentByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopGoodsContentService.DeleteBeeShopGoodsContentByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeShopGoodsContent 更新商品详情描述
// @Tags BeeShopGoodsContent
// @Summary 更新商品详情描述
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsContent true "更新商品详情描述"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeShopGoodsContent/updateBeeShopGoodsContent [put]
func (beeShopGoodsContentApi *BeeShopGoodsContentApi) UpdateBeeShopGoodsContent(c *gin.Context) {
	var beeShopGoodsContent bee.BeeShopGoodsContent
	err := c.ShouldBindJSON(&beeShopGoodsContent)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopGoodsContent.UserId = &shopUserId
	if err := beeShopGoodsContentService.UpdateBeeShopGoodsContent(beeShopGoodsContent, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeShopGoodsContent 用id查询商品详情描述
// @Tags BeeShopGoodsContent
// @Summary 用id查询商品详情描述
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeShopGoodsContent true "用id查询商品详情描述"
// @Success 200 {object} response.Response{data=object{rebeeShopGoodsContent=bee.BeeShopGoodsContent},msg=string} "查询成功"
// @Router /beeShopGoodsContent/findBeeShopGoodsContent [get]
func (beeShopGoodsContentApi *BeeShopGoodsContentApi) FindBeeShopGoodsContent(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeShopGoodsContent, err := beeShopGoodsContentService.GetBeeShopGoodsContent(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeShopGoodsContent, c)
	}
}

// GetBeeShopGoodsContentList 分页获取商品详情描述列表
// @Tags BeeShopGoodsContent
// @Summary 分页获取商品详情描述列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopGoodsContentSearch true "分页获取商品详情描述列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeShopGoodsContent/getBeeShopGoodsContentList [get]
func (beeShopGoodsContentApi *BeeShopGoodsContentApi) GetBeeShopGoodsContentList(c *gin.Context) {
	var pageInfo beeReq.BeeShopGoodsContentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeShopGoodsContentService.GetBeeShopGoodsContentInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeShopGoodsContentPublic 不需要鉴权的商品详情描述接口
// @Tags BeeShopGoodsContent
// @Summary 不需要鉴权的商品详情描述接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopGoodsContentSearch true "分页获取商品详情描述列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeShopGoodsContent/getBeeShopGoodsContentPublic [get]
func (beeShopGoodsContentApi *BeeShopGoodsContentApi) GetBeeShopGoodsContentPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的商品详情描述接口信息",
	}, "获取成功", c)
}
