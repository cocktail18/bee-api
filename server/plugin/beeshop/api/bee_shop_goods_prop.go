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

type BeeShopGoodsPropApi struct{}

var beeShopGoodsPropService = service.ServiceGroupApp.BeeServiceGroup.BeeShopGoodsPropService

// CreateBeeShopGoodsProp 创建商品属性
// @Tags BeeShopGoodsProp
// @Summary 创建商品属性
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsProp true "创建商品属性"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeShopGoodsProp/createBeeShopGoodsProp [post]
func (beeShopGoodsPropApi *BeeShopGoodsPropApi) CreateBeeShopGoodsProp(c *gin.Context) {
	var beeShopGoodsProp bee.BeeShopGoodsProp
	err := c.ShouldBindJSON(&beeShopGoodsProp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopGoodsProp.UserId = &shopUserId
	if err := beeShopGoodsPropService.CreateBeeShopGoodsProp(&beeShopGoodsProp); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeShopGoodsProp 删除商品属性
// @Tags BeeShopGoodsProp
// @Summary 删除商品属性
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsProp true "删除商品属性"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeShopGoodsProp/deleteBeeShopGoodsProp [delete]
func (beeShopGoodsPropApi *BeeShopGoodsPropApi) DeleteBeeShopGoodsProp(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopGoodsPropService.DeleteBeeShopGoodsProp(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeShopGoodsPropByIds 批量删除商品属性
// @Tags BeeShopGoodsProp
// @Summary 批量删除商品属性
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeShopGoodsProp/deleteBeeShopGoodsPropByIds [delete]
func (beeShopGoodsPropApi *BeeShopGoodsPropApi) DeleteBeeShopGoodsPropByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopGoodsPropService.DeleteBeeShopGoodsPropByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeShopGoodsProp 更新商品属性
// @Tags BeeShopGoodsProp
// @Summary 更新商品属性
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopGoodsProp true "更新商品属性"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeShopGoodsProp/updateBeeShopGoodsProp [put]
func (beeShopGoodsPropApi *BeeShopGoodsPropApi) UpdateBeeShopGoodsProp(c *gin.Context) {
	var beeShopGoodsProp bee.BeeShopGoodsProp
	err := c.ShouldBindJSON(&beeShopGoodsProp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopGoodsProp.UserId = &shopUserId
	if err := beeShopGoodsPropService.UpdateBeeShopGoodsProp(beeShopGoodsProp, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeShopGoodsProp 用id查询商品属性
// @Tags BeeShopGoodsProp
// @Summary 用id查询商品属性
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeShopGoodsProp true "用id查询商品属性"
// @Success 200 {object} response.Response{data=object{rebeeShopGoodsProp=bee.BeeShopGoodsProp},msg=string} "查询成功"
// @Router /beeShopGoodsProp/findBeeShopGoodsProp [get]
func (beeShopGoodsPropApi *BeeShopGoodsPropApi) FindBeeShopGoodsProp(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeShopGoodsProp, err := beeShopGoodsPropService.GetBeeShopGoodsProp(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeShopGoodsProp, c)
	}
}

// GetBeeShopGoodsPropList 分页获取商品属性列表
// @Tags BeeShopGoodsProp
// @Summary 分页获取商品属性列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopGoodsPropSearch true "分页获取商品属性列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeShopGoodsProp/getBeeShopGoodsPropList [get]
func (beeShopGoodsPropApi *BeeShopGoodsPropApi) GetBeeShopGoodsPropList(c *gin.Context) {
	var pageInfo beeReq.BeeShopGoodsPropSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeShopGoodsPropService.GetBeeShopGoodsPropInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeShopGoodsPropPublic 不需要鉴权的商品属性接口
// @Tags BeeShopGoodsProp
// @Summary 不需要鉴权的商品属性接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopGoodsPropSearch true "分页获取商品属性列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeShopGoodsProp/getBeeShopGoodsPropPublic [get]
func (beeShopGoodsPropApi *BeeShopGoodsPropApi) GetBeeShopGoodsPropPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的商品属性接口信息",
	}, "获取成功", c)
}
