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

type BeeOrderPeisongApi struct{}

var beeOrderPeisongService = service.ServiceGroupApp.BeeServiceGroup.BeeOrderPeisongService

// CreateBeeOrderPeisong 创建beeOrderPeisong表
// @Tags BeeOrderPeisong
// @Summary 创建beeOrderPeisong表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderPeisong true "创建beeOrderPeisong表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeOrderPeisong/createBeeOrderPeisong [post]
func (beeOrderPeisongApi *BeeOrderPeisongApi) CreateBeeOrderPeisong(c *gin.Context) {
	var beeOrderPeisong bee.BeeOrderPeisong
	err := c.ShouldBindJSON(&beeOrderPeisong)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeOrderPeisong.UserId = &shopUserId
	if err := beeOrderPeisongService.CreateBeeOrderPeisong(&beeOrderPeisong); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeOrderPeisong 删除beeOrderPeisong表
// @Tags BeeOrderPeisong
// @Summary 删除beeOrderPeisong表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderPeisong true "删除beeOrderPeisong表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeOrderPeisong/deleteBeeOrderPeisong [delete]
func (beeOrderPeisongApi *BeeOrderPeisongApi) DeleteBeeOrderPeisong(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderPeisongService.DeleteBeeOrderPeisong(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeOrderPeisongByIds 批量删除beeOrderPeisong表
// @Tags BeeOrderPeisong
// @Summary 批量删除beeOrderPeisong表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeOrderPeisong/deleteBeeOrderPeisongByIds [delete]
func (beeOrderPeisongApi *BeeOrderPeisongApi) DeleteBeeOrderPeisongByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderPeisongService.DeleteBeeOrderPeisongByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeOrderPeisong 更新beeOrderPeisong表
// @Tags BeeOrderPeisong
// @Summary 更新beeOrderPeisong表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderPeisong true "更新beeOrderPeisong表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeOrderPeisong/updateBeeOrderPeisong [put]
func (beeOrderPeisongApi *BeeOrderPeisongApi) UpdateBeeOrderPeisong(c *gin.Context) {
	var beeOrderPeisong bee.BeeOrderPeisong
	err := c.ShouldBindJSON(&beeOrderPeisong)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeOrderPeisong.UserId = &shopUserId
	if err := beeOrderPeisongService.UpdateBeeOrderPeisong(beeOrderPeisong, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// CancelBeeOrderPeisong 取消配送
// @Tags BeeOrderPeisong
// @Summary 取消配送
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderPeisong true "更新beeOrderPeisong表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeOrderPeisong/cancelBeeOrderPeisong [post]
func (beeOrderPeisongApi *BeeOrderPeisongApi) CancelBeeOrderPeisong(c *gin.Context) {
	var beeOrderPeisong beeReq.BeeOrderPeisongCancel
	err := c.ShouldBindJSON(&beeOrderPeisong)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeOrderPeisongService.Cancel(beeOrderPeisong, shopUserId); err != nil {
		global.GVA_LOG.Error("取消失败!", zap.Error(err))
		response.FailWithMessage("取消失败:"+err.Error(), c)
	} else {
		response.OkWithMessage("取消成功", c)
	}
}

// NotifyBeeOrderPeisong 通知供应商进行配送
// @Tags BeeOrderPeisong
// @Summary 取消配送
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeOrderPeisong true "更新beeOrderPeisong表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeOrderPeisong/notifyBeeOrderPeisong [post]
func (beeOrderPeisongApi *BeeOrderPeisongApi) NotifyBeeOrderPeisong(c *gin.Context) {
	var beeOrderPeisong bee.BeeOrderPeisong
	err := c.ShouldBindJSON(&beeOrderPeisong)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeOrderPeisong.UserId = &shopUserId
	if err := beeOrderPeisongService.Notify(beeOrderPeisong, shopUserId); err != nil {
		global.GVA_LOG.Error("通知失败!", zap.Error(err))
		response.FailWithMessage("通知失败", c)
	} else {
		response.OkWithMessage("通知成功", c)
	}
}

// FindBeeOrderPeisong 用id查询beeOrderPeisong表
// @Tags BeeOrderPeisong
// @Summary 用id查询beeOrderPeisong表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeOrderPeisong true "用id查询beeOrderPeisong表"
// @Success 200 {object} response.Response{data=object{rebeeOrderPeisong=bee.BeeOrderPeisong},msg=string} "查询成功"
// @Router /beeOrderPeisong/findBeeOrderPeisong [get]
func (beeOrderPeisongApi *BeeOrderPeisongApi) FindBeeOrderPeisong(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeOrderPeisong, err := beeOrderPeisongService.GetBeeOrderPeisong(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeOrderPeisong, c)
	}
}

// GetBeeOrderPeisongList 分页获取beeOrderPeisong表列表
// @Tags BeeOrderPeisong
// @Summary 分页获取beeOrderPeisong表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeOrderPeisongSearch true "分页获取beeOrderPeisong表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeOrderPeisong/getBeeOrderPeisongList [get]
func (beeOrderPeisongApi *BeeOrderPeisongApi) GetBeeOrderPeisongList(c *gin.Context) {
	var pageInfo beeReq.BeeOrderPeisongSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeOrderPeisongService.GetBeeOrderPeisongInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeOrderPeisongDetail 获取当前配送信息
// @Tags BeeOrderPeisong
// @Summary 不需要鉴权的beeOrderPeisong表接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeOrderPeisongQueryDetail true "查询配送详情"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeOrderPeisong/getBeeOrderPeisongDetail [get]
func (beeOrderPeisongApi *BeeOrderPeisongApi) GetBeeOrderPeisongDetail(c *gin.Context) {
	var info beeReq.BeeOrderPeisongQueryDetail
	err := c.ShouldBindQuery(&info)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if detail, err := beeOrderPeisongService.GetBeeOrderPeisongDetail(info.Id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(detail, c)
	}
}

// GetBeeOrderPeisongPublic 不需要鉴权的beeOrderPeisong表接口
// @Tags BeeOrderPeisong
// @Summary 不需要鉴权的beeOrderPeisong表接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeOrderPeisongSearch true "分页获取beeOrderPeisong表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeOrderPeisong/getBeeOrderPeisongPublic [get]
func (beeOrderPeisongApi *BeeOrderPeisongApi) GetBeeOrderPeisongPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的beeOrderPeisong表接口信息",
	}, "获取成功", c)
}
