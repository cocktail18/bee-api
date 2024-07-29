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

type BeeCmsInfoApi struct{}

var beeCmsInfoService = service.ServiceGroupApp.BeeServiceGroup.BeeCmsInfoService

// CreateBeeCmsInfo 创建cms信息
// @Tags BeeCmsInfo
// @Summary 创建cms信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeCmsInfo true "创建cms信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeCmsInfo/createBeeCmsInfo [post]
func (beeCmsInfoApi *BeeCmsInfoApi) CreateBeeCmsInfo(c *gin.Context) {
	var beeCmsInfo bee.BeeCmsInfo
	err := c.ShouldBindJSON(&beeCmsInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeCmsInfo.UserId = &shopUserId
	if err := beeCmsInfoService.CreateBeeCmsInfo(&beeCmsInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeCmsInfo 删除cms信息
// @Tags BeeCmsInfo
// @Summary 删除cms信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeCmsInfo true "删除cms信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeCmsInfo/deleteBeeCmsInfo [delete]
func (beeCmsInfoApi *BeeCmsInfoApi) DeleteBeeCmsInfo(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeCmsInfoService.DeleteBeeCmsInfo(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeCmsInfoByIds 批量删除cms信息
// @Tags BeeCmsInfo
// @Summary 批量删除cms信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeCmsInfo/deleteBeeCmsInfoByIds [delete]
func (beeCmsInfoApi *BeeCmsInfoApi) DeleteBeeCmsInfoByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeCmsInfoService.DeleteBeeCmsInfoByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeCmsInfo 更新cms信息
// @Tags BeeCmsInfo
// @Summary 更新cms信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeCmsInfo true "更新cms信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeCmsInfo/updateBeeCmsInfo [put]
func (beeCmsInfoApi *BeeCmsInfoApi) UpdateBeeCmsInfo(c *gin.Context) {
	var beeCmsInfo bee.BeeCmsInfo
	err := c.ShouldBindJSON(&beeCmsInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeCmsInfo.UserId = &shopUserId
	if err := beeCmsInfoService.UpdateBeeCmsInfo(beeCmsInfo, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeCmsInfo 用id查询cms信息
// @Tags BeeCmsInfo
// @Summary 用id查询cms信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeCmsInfo true "用id查询cms信息"
// @Success 200 {object} response.Response{data=object{rebeeCmsInfo=bee.BeeCmsInfo},msg=string} "查询成功"
// @Router /beeCmsInfo/findBeeCmsInfo [get]
func (beeCmsInfoApi *BeeCmsInfoApi) FindBeeCmsInfo(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeCmsInfo, err := beeCmsInfoService.GetBeeCmsInfo(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeCmsInfo, c)
	}
}

// GetBeeCmsInfoList 分页获取cms信息列表
// @Tags BeeCmsInfo
// @Summary 分页获取cms信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeCmsInfoSearch true "分页获取cms信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeCmsInfo/getBeeCmsInfoList [get]
func (beeCmsInfoApi *BeeCmsInfoApi) GetBeeCmsInfoList(c *gin.Context) {
	var pageInfo beeReq.BeeCmsInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeCmsInfoService.GetBeeCmsInfoInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeCmsInfoPublic 不需要鉴权的cms信息接口
// @Tags BeeCmsInfo
// @Summary 不需要鉴权的cms信息接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeCmsInfoSearch true "分页获取cms信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeCmsInfo/getBeeCmsInfoPublic [get]
func (beeCmsInfoApi *BeeCmsInfoApi) GetBeeCmsInfoPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的cms信息接口信息",
	}, "获取成功", c)
}
