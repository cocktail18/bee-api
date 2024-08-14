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

type BeePrinterApi struct{}

var beePrinterService = service.ServiceGroupApp.BeeServiceGroup.BeePrinterService

// CreateBeePrinter 创建beePrinter表
// @Tags BeePrinter
// @Summary 创建beePrinter表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeePrinter true "创建beePrinter表"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beePrinter/createBeePrinter [post]
func (beePrinterApi *BeePrinterApi) CreateBeePrinter(c *gin.Context) {
	var beePrinter bee.BeePrinter
	err := c.ShouldBindJSON(&beePrinter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beePrinter.UserId = &shopUserId
	if err := beePrinterService.CreateBeePrinter(&beePrinter); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeePrinter 删除beePrinter表
// @Tags BeePrinter
// @Summary 删除beePrinter表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeePrinter true "删除beePrinter表"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beePrinter/deleteBeePrinter [delete]
func (beePrinterApi *BeePrinterApi) DeleteBeePrinter(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beePrinterService.DeleteBeePrinter(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeePrinterByIds 批量删除beePrinter表
// @Tags BeePrinter
// @Summary 批量删除beePrinter表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beePrinter/deleteBeePrinterByIds [delete]
func (beePrinterApi *BeePrinterApi) DeleteBeePrinterByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beePrinterService.DeleteBeePrinterByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeePrinter 更新beePrinter表
// @Tags BeePrinter
// @Summary 更新beePrinter表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeePrinter true "更新beePrinter表"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beePrinter/updateBeePrinter [put]
func (beePrinterApi *BeePrinterApi) UpdateBeePrinter(c *gin.Context) {
	var beePrinter bee.BeePrinter
	err := c.ShouldBindJSON(&beePrinter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beePrinter.UserId = &shopUserId
	if err := beePrinterService.UpdateBeePrinter(beePrinter, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// TestBeePrinter 检查配置是否正确
// @Tags BeePrinter
// @Summary 检查配置是否正确
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeePrinter true "检查配置是否正确"
// @Success 200 {object} response.Response{msg=string} "成功"
// @Router /beePrinter/updateBeePrinter [post]
func (beePrinterApi *BeePrinterApi) TestBeePrinter(c *gin.Context) {
	var beePrinter bee.BeePrinter
	err := c.ShouldBindJSON(&beePrinter)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beePrinter.UserId = &shopUserId
	if err := beePrinterService.TestBeePrinter(beePrinter, shopUserId); err != nil {
		response.FailWithMessage("测试失败:"+err.Error(), c)
	} else {
		response.OkWithMessage("测试成功", c)
	}
}

// FindBeePrinter 用id查询beePrinter表
// @Tags BeePrinter
// @Summary 用id查询beePrinter表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeePrinter true "用id查询beePrinter表"
// @Success 200 {object} response.Response{data=object{rebeePrinter=bee.BeePrinter},msg=string} "查询成功"
// @Router /beePrinter/findBeePrinter [get]
func (beePrinterApi *BeePrinterApi) FindBeePrinter(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeePrinter, err := beePrinterService.GetBeePrinter(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeePrinter, c)
	}
}

// GetBeePrinterList 分页获取beePrinter表列表
// @Tags BeePrinter
// @Summary 分页获取beePrinter表列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeePrinterSearch true "分页获取beePrinter表列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beePrinter/getBeePrinterList [get]
func (beePrinterApi *BeePrinterApi) GetBeePrinterList(c *gin.Context) {
	var pageInfo beeReq.BeePrinterSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beePrinterService.GetBeePrinterInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeePrinterPublic 不需要鉴权的beePrinter表接口
// @Tags BeePrinter
// @Summary 不需要鉴权的beePrinter表接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeePrinterSearch true "分页获取beePrinter表列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beePrinter/getBeePrinterPublic [get]
func (beePrinterApi *BeePrinterApi) GetBeePrinterPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的beePrinter表接口信息",
	}, "获取成功", c)
}
