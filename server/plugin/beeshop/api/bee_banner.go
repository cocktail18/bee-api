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

type BeeBannerApi struct{}

var beeBannerService = service.ServiceGroupApp.BeeServiceGroup.BeeBannerService

// CreateBeeBanner 创建商店banner
// @Tags BeeBanner
// @Summary 创建商店banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeBanner true "创建商店banner"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeBanner/createBeeBanner [post]
func (beeBannerApi *BeeBannerApi) CreateBeeBanner(c *gin.Context) {
	var beeBanner bee.BeeBanner
	err := c.ShouldBindJSON(&beeBanner)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeBanner.UserId = &shopUserId
	if err := beeBannerService.CreateBeeBanner(&beeBanner); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeBanner 删除商店banner
// @Tags BeeBanner
// @Summary 删除商店banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeBanner true "删除商店banner"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeBanner/deleteBeeBanner [delete]
func (beeBannerApi *BeeBannerApi) DeleteBeeBanner(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeBannerService.DeleteBeeBanner(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeBannerByIds 批量删除商店banner
// @Tags BeeBanner
// @Summary 批量删除商店banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeBanner/deleteBeeBannerByIds [delete]
func (beeBannerApi *BeeBannerApi) DeleteBeeBannerByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeBannerService.DeleteBeeBannerByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeBanner 更新商店banner
// @Tags BeeBanner
// @Summary 更新商店banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeBanner true "更新商店banner"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeBanner/updateBeeBanner [put]
func (beeBannerApi *BeeBannerApi) UpdateBeeBanner(c *gin.Context) {
	var beeBanner bee.BeeBanner
	err := c.ShouldBindJSON(&beeBanner)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeBanner.UserId = &shopUserId
	if err := beeBannerService.UpdateBeeBanner(beeBanner, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeBanner 用id查询商店banner
// @Tags BeeBanner
// @Summary 用id查询商店banner
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeBanner true "用id查询商店banner"
// @Success 200 {object} response.Response{data=object{rebeeBanner=bee.BeeBanner},msg=string} "查询成功"
// @Router /beeBanner/findBeeBanner [get]
func (beeBannerApi *BeeBannerApi) FindBeeBanner(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeBanner, err := beeBannerService.GetBeeBanner(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeBanner, c)
	}
}

// GetBeeBannerList 分页获取商店banner列表
// @Tags BeeBanner
// @Summary 分页获取商店banner列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeBannerSearch true "分页获取商店banner列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeBanner/getBeeBannerList [get]
func (beeBannerApi *BeeBannerApi) GetBeeBannerList(c *gin.Context) {
	var pageInfo beeReq.BeeBannerSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeBannerService.GetBeeBannerInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeBannerPublic 不需要鉴权的商店banner接口
// @Tags BeeBanner
// @Summary 不需要鉴权的商店banner接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeBannerSearch true "分页获取商店banner列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeBanner/getBeeBannerPublic [get]
func (beeBannerApi *BeeBannerApi) GetBeeBannerPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的商店banner接口信息",
	}, "获取成功", c)
}
