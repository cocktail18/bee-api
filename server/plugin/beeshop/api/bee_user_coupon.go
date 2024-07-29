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

type BeeUserCouponApi struct{}

var beeUserCouponService = service.ServiceGroupApp.BeeServiceGroup.BeeUserCouponService

// CreateBeeUserCoupon 创建用户优惠券
// @Tags BeeUserCoupon
// @Summary 创建用户优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserCoupon true "创建用户优惠券"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeUserCoupon/createBeeUserCoupon [post]
func (beeUserCouponApi *BeeUserCouponApi) CreateBeeUserCoupon(c *gin.Context) {
	var beeUserCoupon bee.BeeUserCoupon
	err := c.ShouldBindJSON(&beeUserCoupon)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUserCoupon.UserId = &shopUserId
	if err := beeUserCouponService.CreateBeeUserCoupon(&beeUserCoupon); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeUserCoupon 删除用户优惠券
// @Tags BeeUserCoupon
// @Summary 删除用户优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserCoupon true "删除用户优惠券"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeUserCoupon/deleteBeeUserCoupon [delete]
func (beeUserCouponApi *BeeUserCouponApi) DeleteBeeUserCoupon(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUserCouponService.DeleteBeeUserCoupon(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeUserCouponByIds 批量删除用户优惠券
// @Tags BeeUserCoupon
// @Summary 批量删除用户优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeUserCoupon/deleteBeeUserCouponByIds [delete]
func (beeUserCouponApi *BeeUserCouponApi) DeleteBeeUserCouponByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUserCouponService.DeleteBeeUserCouponByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeUserCoupon 更新用户优惠券
// @Tags BeeUserCoupon
// @Summary 更新用户优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserCoupon true "更新用户优惠券"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeUserCoupon/updateBeeUserCoupon [put]
func (beeUserCouponApi *BeeUserCouponApi) UpdateBeeUserCoupon(c *gin.Context) {
	var beeUserCoupon bee.BeeUserCoupon
	err := c.ShouldBindJSON(&beeUserCoupon)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUserCoupon.UserId = &shopUserId
	if err := beeUserCouponService.UpdateBeeUserCoupon(beeUserCoupon, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeUserCoupon 用id查询用户优惠券
// @Tags BeeUserCoupon
// @Summary 用id查询用户优惠券
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeUserCoupon true "用id查询用户优惠券"
// @Success 200 {object} response.Response{data=object{rebeeUserCoupon=bee.BeeUserCoupon},msg=string} "查询成功"
// @Router /beeUserCoupon/findBeeUserCoupon [get]
func (beeUserCouponApi *BeeUserCouponApi) FindBeeUserCoupon(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeUserCoupon, err := beeUserCouponService.GetBeeUserCoupon(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeUserCoupon, c)
	}
}

// GetBeeUserCouponList 分页获取用户优惠券列表
// @Tags BeeUserCoupon
// @Summary 分页获取用户优惠券列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUserCouponSearch true "分页获取用户优惠券列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeUserCoupon/getBeeUserCouponList [get]
func (beeUserCouponApi *BeeUserCouponApi) GetBeeUserCouponList(c *gin.Context) {
	var pageInfo beeReq.BeeUserCouponSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeUserCouponService.GetBeeUserCouponInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeUserCouponPublic 不需要鉴权的用户优惠券接口
// @Tags BeeUserCoupon
// @Summary 不需要鉴权的用户优惠券接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUserCouponSearch true "分页获取用户优惠券列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeUserCoupon/getBeeUserCouponPublic [get]
func (beeUserCouponApi *BeeUserCouponApi) GetBeeUserCouponPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的用户优惠券接口信息",
	}, "获取成功", c)
}
