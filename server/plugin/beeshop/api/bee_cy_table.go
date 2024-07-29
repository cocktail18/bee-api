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

type BeeCyTableApi struct{}

var beeCyTableService = service.ServiceGroupApp.BeeServiceGroup.BeeCyTableService

// CreateBeeCyTable 创建桌号信息
// @Tags BeeCyTable
// @Summary 创建桌号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeCyTable true "创建桌号信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeCyTable/createBeeCyTable [post]
func (beeCyTableApi *BeeCyTableApi) CreateBeeCyTable(c *gin.Context) {
	var beeCyTable bee.BeeCyTable
	err := c.ShouldBindJSON(&beeCyTable)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeCyTable.UserId = &shopUserId
	if err := beeCyTableService.CreateBeeCyTable(&beeCyTable); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeCyTable 删除桌号信息
// @Tags BeeCyTable
// @Summary 删除桌号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeCyTable true "删除桌号信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeCyTable/deleteBeeCyTable [delete]
func (beeCyTableApi *BeeCyTableApi) DeleteBeeCyTable(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeCyTableService.DeleteBeeCyTable(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeCyTableByIds 批量删除桌号信息
// @Tags BeeCyTable
// @Summary 批量删除桌号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeCyTable/deleteBeeCyTableByIds [delete]
func (beeCyTableApi *BeeCyTableApi) DeleteBeeCyTableByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeCyTableService.DeleteBeeCyTableByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeCyTable 更新桌号信息
// @Tags BeeCyTable
// @Summary 更新桌号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeCyTable true "更新桌号信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeCyTable/updateBeeCyTable [put]
func (beeCyTableApi *BeeCyTableApi) UpdateBeeCyTable(c *gin.Context) {
	var beeCyTable bee.BeeCyTable
	err := c.ShouldBindJSON(&beeCyTable)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeCyTable.UserId = &shopUserId
	if err := beeCyTableService.UpdateBeeCyTable(beeCyTable, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeCyTable 用id查询桌号信息
// @Tags BeeCyTable
// @Summary 用id查询桌号信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeCyTable true "用id查询桌号信息"
// @Success 200 {object} response.Response{data=object{rebeeCyTable=bee.BeeCyTable},msg=string} "查询成功"
// @Router /beeCyTable/findBeeCyTable [get]
func (beeCyTableApi *BeeCyTableApi) FindBeeCyTable(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeCyTable, err := beeCyTableService.GetBeeCyTable(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeCyTable, c)
	}
}

// GetBeeCyTableList 分页获取桌号信息列表
// @Tags BeeCyTable
// @Summary 分页获取桌号信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeCyTableSearch true "分页获取桌号信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeCyTable/getBeeCyTableList [get]
func (beeCyTableApi *BeeCyTableApi) GetBeeCyTableList(c *gin.Context) {
	var pageInfo beeReq.BeeCyTableSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeCyTableService.GetBeeCyTableInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeCyTablePublic 不需要鉴权的桌号信息接口
// @Tags BeeCyTable
// @Summary 不需要鉴权的桌号信息接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeCyTableSearch true "分页获取桌号信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeCyTable/getBeeCyTablePublic [get]
func (beeCyTableApi *BeeCyTableApi) GetBeeCyTablePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的桌号信息接口信息",
	}, "获取成功", c)
}
