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

type BeeUserMapperApi struct{}

var beeUserMapperService = service.ServiceGroupApp.BeeServiceGroup.BeeUserMapperService

// CreateBeeUserMapper 创建beeUserMapper表
// @Tags BeeUserMapper
// @Summary 创建beeUserMapper表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserMapper true "创建beeUserMapper表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeUserMapper/createBeeUserMapper [post]
func (beeUserMapperApi *BeeUserMapperApi) CreateBeeUserMapper(c *gin.Context) {
	var beeUserMapper bee.BeeUserMapper
	err := c.ShouldBindJSON(&beeUserMapper)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUserMapper.UserId = &shopUserId
	if err := beeUserMapperService.CreateBeeUserMapper(&beeUserMapper); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeUserMapper 删除beeUserMapper表
// @Tags BeeUserMapper
// @Summary 删除beeUserMapper表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserMapper true "删除beeUserMapper表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeUserMapper/deleteBeeUserMapper [delete]
func (beeUserMapperApi *BeeUserMapperApi) DeleteBeeUserMapper(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUserMapperService.DeleteBeeUserMapper(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeUserMapperByIds 批量删除beeUserMapper表
// @Tags BeeUserMapper
// @Summary 批量删除beeUserMapper表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeUserMapper/deleteBeeUserMapperByIds [delete]
func (beeUserMapperApi *BeeUserMapperApi) DeleteBeeUserMapperByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUserMapperService.DeleteBeeUserMapperByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeUserMapper 更新beeUserMapper表
// @Tags BeeUserMapper
// @Summary 更新beeUserMapper表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUserMapper true "更新beeUserMapper表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeUserMapper/updateBeeUserMapper [put]
func (beeUserMapperApi *BeeUserMapperApi) UpdateBeeUserMapper(c *gin.Context) {
	var beeUserMapper bee.BeeUserMapper
	err := c.ShouldBindJSON(&beeUserMapper)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUserMapper.UserId = &shopUserId
	if err := beeUserMapperService.UpdateBeeUserMapper(beeUserMapper, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeUserMapper 用id查询beeUserMapper表
// @Tags BeeUserMapper
// @Summary 用id查询beeUserMapper表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeUserMapper true "用id查询beeUserMapper表"
// @Success 200 {object} response.Response{data=object{rebeeUserMapper=bee.BeeUserMapper},msg=string} "查询成功"
// @Router /beeUserMapper/findBeeUserMapper [get]
func (beeUserMapperApi *BeeUserMapperApi) FindBeeUserMapper(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeUserMapper, err := beeUserMapperService.GetBeeUserMapper(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeUserMapper, c)
	}
}

// GetBeeUserMapperList 分页获取beeUserMapper表列表
// @Tags BeeUserMapper
// @Summary 分页获取beeUserMapper表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUserMapperSearch true "分页获取beeUserMapper表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeUserMapper/getBeeUserMapperList [get]
func (beeUserMapperApi *BeeUserMapperApi) GetBeeUserMapperList(c *gin.Context) {
	var pageInfo beeReq.BeeUserMapperSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeUserMapperService.GetBeeUserMapperInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeUserMapperPublic 不需要鉴权的beeUserMapper表接口
// @Tags BeeUserMapper
// @Summary 不需要鉴权的beeUserMapper表接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUserMapperSearch true "分页获取beeUserMapper表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeUserMapper/getBeeUserMapperPublic [get]
func (beeUserMapperApi *BeeUserMapperApi) GetBeeUserMapperPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的beeUserMapper表接口信息",
	}, "获取成功", c)
}
