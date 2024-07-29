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

type BeeNoticeApi struct{}

var beeNoticeService = service.ServiceGroupApp.BeeServiceGroup.BeeNoticeService

// CreateBeeNotice 创建系统公告
// @Tags BeeNotice
// @Summary 创建系统公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeNotice true "创建系统公告"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeNotice/createBeeNotice [post]
func (beeNoticeApi *BeeNoticeApi) CreateBeeNotice(c *gin.Context) {
	var beeNotice bee.BeeNotice
	err := c.ShouldBindJSON(&beeNotice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeNotice.UserId = &shopUserId
	if err := beeNoticeService.CreateBeeNotice(&beeNotice); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeNotice 删除系统公告
// @Tags BeeNotice
// @Summary 删除系统公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeNotice true "删除系统公告"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeNotice/deleteBeeNotice [delete]
func (beeNoticeApi *BeeNoticeApi) DeleteBeeNotice(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeNoticeService.DeleteBeeNotice(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeNoticeByIds 批量删除系统公告
// @Tags BeeNotice
// @Summary 批量删除系统公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeNotice/deleteBeeNoticeByIds [delete]
func (beeNoticeApi *BeeNoticeApi) DeleteBeeNoticeByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeNoticeService.DeleteBeeNoticeByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeNotice 更新系统公告
// @Tags BeeNotice
// @Summary 更新系统公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeNotice true "更新系统公告"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeNotice/updateBeeNotice [put]
func (beeNoticeApi *BeeNoticeApi) UpdateBeeNotice(c *gin.Context) {
	var beeNotice bee.BeeNotice
	err := c.ShouldBindJSON(&beeNotice)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeNotice.UserId = &shopUserId
	if err := beeNoticeService.UpdateBeeNotice(beeNotice, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeNotice 用id查询系统公告
// @Tags BeeNotice
// @Summary 用id查询系统公告
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeNotice true "用id查询系统公告"
// @Success 200 {object} response.Response{data=object{rebeeNotice=bee.BeeNotice},msg=string} "查询成功"
// @Router /beeNotice/findBeeNotice [get]
func (beeNoticeApi *BeeNoticeApi) FindBeeNotice(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeNotice, err := beeNoticeService.GetBeeNotice(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeNotice, c)
	}
}

// GetBeeNoticeList 分页获取系统公告列表
// @Tags BeeNotice
// @Summary 分页获取系统公告列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeNoticeSearch true "分页获取系统公告列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeNotice/getBeeNoticeList [get]
func (beeNoticeApi *BeeNoticeApi) GetBeeNoticeList(c *gin.Context) {
	var pageInfo beeReq.BeeNoticeSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeNoticeService.GetBeeNoticeInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeNoticePublic 不需要鉴权的系统公告接口
// @Tags BeeNotice
// @Summary 不需要鉴权的系统公告接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeNoticeSearch true "分页获取系统公告列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeNotice/getBeeNoticePublic [get]
func (beeNoticeApi *BeeNoticeApi) GetBeeNoticePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的系统公告接口信息",
	}, "获取成功", c)
}
