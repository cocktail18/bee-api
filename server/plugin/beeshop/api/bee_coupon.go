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

type BeeCouponApi struct{}

var beeCouponService = service.ServiceGroupApp.BeeServiceGroup.BeeCouponService

// CreateBeeCoupon 创建优惠券
// @Tags BeeCoupon
// @Summary 创建优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeCoupon true "创建优惠券"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeCoupon/createBeeCoupon [post]
func (beeCouponApi *BeeCouponApi) CreateBeeCoupon(c *gin.Context) {
	var beeCoupon bee.BeeCoupon
	err := c.ShouldBindJSON(&beeCoupon)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeCoupon.UserId = &shopUserId
	if err := beeCouponService.CreateBeeCoupon(&beeCoupon); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeCoupon 删除优惠券
// @Tags BeeCoupon
// @Summary 删除优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeCoupon true "删除优惠券"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeCoupon/deleteBeeCoupon [delete]
func (beeCouponApi *BeeCouponApi) DeleteBeeCoupon(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeCouponService.DeleteBeeCoupon(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeCouponByIds 批量删除优惠券
// @Tags BeeCoupon
// @Summary 批量删除优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeCoupon/deleteBeeCouponByIds [delete]
func (beeCouponApi *BeeCouponApi) DeleteBeeCouponByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeCouponService.DeleteBeeCouponByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeCoupon 更新优惠券
// @Tags BeeCoupon
// @Summary 更新优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeCoupon true "更新优惠券"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeCoupon/updateBeeCoupon [put]
func (beeCouponApi *BeeCouponApi) UpdateBeeCoupon(c *gin.Context) {
	var beeCoupon bee.BeeCoupon
	err := c.ShouldBindJSON(&beeCoupon)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeCoupon.UserId = &shopUserId
	if err := beeCouponService.UpdateBeeCoupon(beeCoupon, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeCoupon 用id查询优惠券
// @Tags BeeCoupon
// @Summary 用id查询优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeCoupon true "用id查询优惠券"
// @Success 200 {object} response.Response{data=object{rebeeCoupon=bee.BeeCoupon},msg=string} "查询成功"
// @Router /beeCoupon/findBeeCoupon [get]
func (beeCouponApi *BeeCouponApi) FindBeeCoupon(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeCoupon, err := beeCouponService.GetBeeCoupon(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeCoupon, c)
	}
}

// GetBeeCouponList 分页获取优惠券列表
// @Tags BeeCoupon
// @Summary 分页获取优惠券列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeCouponSearch true "分页获取优惠券列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeCoupon/getBeeCouponList [get]
func (beeCouponApi *BeeCouponApi) GetBeeCouponList(c *gin.Context) {
	var pageInfo beeReq.BeeCouponSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeCouponService.GetBeeCouponInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeCouponPublic 不需要鉴权的优惠券接口
// @Tags BeeCoupon
// @Summary 不需要鉴权的优惠券接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeCouponSearch true "分页获取优惠券列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeCoupon/getBeeCouponPublic [get]
func (beeCouponApi *BeeCouponApi) GetBeeCouponPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的优惠券接口信息",
	}, "获取成功", c)
}
