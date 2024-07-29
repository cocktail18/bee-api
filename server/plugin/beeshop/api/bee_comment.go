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

type BeeCommentApi struct{}

var beeCommentService = service.ServiceGroupApp.BeeServiceGroup.BeeCommentService

// CreateBeeComment 创建评论表
// @Tags BeeComment
// @Summary 创建评论表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeComment true "创建评论表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeComment/createBeeComment [post]
func (beeCommentApi *BeeCommentApi) CreateBeeComment(c *gin.Context) {
	var beeComment bee.BeeComment
	err := c.ShouldBindJSON(&beeComment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeComment.UserId = &shopUserId
	if err := beeCommentService.CreateBeeComment(&beeComment); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeComment 删除评论表
// @Tags BeeComment
// @Summary 删除评论表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeComment true "删除评论表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeComment/deleteBeeComment [delete]
func (beeCommentApi *BeeCommentApi) DeleteBeeComment(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeCommentService.DeleteBeeComment(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeCommentByIds 批量删除评论表
// @Tags BeeComment
// @Summary 批量删除评论表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeComment/deleteBeeCommentByIds [delete]
func (beeCommentApi *BeeCommentApi) DeleteBeeCommentByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeCommentService.DeleteBeeCommentByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeComment 更新评论表
// @Tags BeeComment
// @Summary 更新评论表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeComment true "更新评论表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeComment/updateBeeComment [put]
func (beeCommentApi *BeeCommentApi) UpdateBeeComment(c *gin.Context) {
	var beeComment bee.BeeComment
	err := c.ShouldBindJSON(&beeComment)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeComment.UserId = &shopUserId
	if err := beeCommentService.UpdateBeeComment(beeComment, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeComment 用id查询评论表
// @Tags BeeComment
// @Summary 用id查询评论表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeComment true "用id查询评论表"
// @Success 200 {object} response.Response{data=object{rebeeComment=bee.BeeComment},msg=string} "查询成功"
// @Router /beeComment/findBeeComment [get]
func (beeCommentApi *BeeCommentApi) FindBeeComment(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeComment, err := beeCommentService.GetBeeComment(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeComment, c)
	}
}

// GetBeeCommentList 分页获取评论表列表
// @Tags BeeComment
// @Summary 分页获取评论表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeCommentSearch true "分页获取评论表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeComment/getBeeCommentList [get]
func (beeCommentApi *BeeCommentApi) GetBeeCommentList(c *gin.Context) {
	var pageInfo beeReq.BeeCommentSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeCommentService.GetBeeCommentInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeCommentPublic 不需要鉴权的评论表接口
// @Tags BeeComment
// @Summary 不需要鉴权的评论表接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeCommentSearch true "分页获取评论表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeComment/getBeeCommentPublic [get]
func (beeCommentApi *BeeCommentApi) GetBeeCommentPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的评论表接口信息",
	}, "获取成功", c)
}
