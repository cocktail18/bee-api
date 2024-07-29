package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/service"
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"math"
)

type BeeRegionApi struct{}

var beeRegionService = service.ServiceGroupApp.BeeServiceGroup.BeeRegionService

// CreateBeeRegion 创建地址库
// @Tags BeeRegion
// @Summary 创建地址库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeRegion true "创建地址库"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeRegion/createBeeRegion [post]
func (beeRegionApi *BeeRegionApi) CreateBeeRegion(c *gin.Context) {
	var beeRegion bee.BeeRegion
	err := c.ShouldBindJSON(&beeRegion)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := beeRegionService.CreateBeeRegion(&beeRegion); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeRegion 删除地址库
// @Tags BeeRegion
// @Summary 删除地址库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeRegion true "删除地址库"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeRegion/deleteBeeRegion [delete]
func (beeRegionApi *BeeRegionApi) DeleteBeeRegion(c *gin.Context) {
	id := c.Query("id")
	if err := beeRegionService.DeleteBeeRegion(id); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeRegionByIds 批量删除地址库
// @Tags BeeRegion
// @Summary 批量删除地址库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeRegion/deleteBeeRegionByIds [delete]
func (beeRegionApi *BeeRegionApi) DeleteBeeRegionByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	if err := beeRegionService.DeleteBeeRegionByIds(ids); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeRegion 更新地址库
// @Tags BeeRegion
// @Summary 更新地址库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeRegion true "更新地址库"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeRegion/updateBeeRegion [put]
func (beeRegionApi *BeeRegionApi) UpdateBeeRegion(c *gin.Context) {
	var beeRegion bee.BeeRegion
	err := c.ShouldBindJSON(&beeRegion)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := beeRegionService.UpdateBeeRegion(beeRegion); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeRegion 用id查询地址库
// @Tags BeeRegion
// @Summary 用id查询地址库
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeRegion true "用id查询地址库"
// @Success 200 {object} response.Response{data=object{rebeeRegion=bee.BeeRegion},msg=string} "查询成功"
// @Router /beeRegion/findBeeRegion [get]
func (beeRegionApi *BeeRegionApi) FindBeeRegion(c *gin.Context) {
	id := c.Query("id")
	if rebeeRegion, err := beeRegionService.GetBeeRegion(id); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeRegion, c)
	}
}

// GetBeeRegionList 分页获取地址库列表
// @Tags BeeRegion
// @Summary 分页获取地址库列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeRegionSearch true "分页获取地址库列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeRegion/getBeeRegionList [get]
func (beeRegionApi *BeeRegionApi) GetBeeRegionList(c *gin.Context) {
	var pageInfo beeReq.BeeRegionSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := beeRegionService.GetBeeRegionInfoList(pageInfo); err != nil {
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
func (beeRegionApi *BeeRegionApi) buildTree(items []*bee.BeeRegion, regionByPid map[string][]*bee.BeeRegion) {
	for _, item := range items {
		item.Childs = regionByPid[cast.ToString(item.Id)]
		beeRegionApi.buildTree(item.Childs, regionByPid)
	}
}

// GetBeeRegionCityTree 获取省市区树
// @Tags BeeRegion
// @Summary 获取省市区树
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeRegion/getBeeRegionCityTree [get]
func (beeRegionApi *BeeRegionApi) GetBeeRegionCityTree(c *gin.Context) {
	var pageInfo beeReq.BeeRegionSearch
	f := false
	pageInfo.LevelList = []int{1, 2, 3}
	pageInfo.IsDeleted = &f
	pageInfo.Page = 1
	pageInfo.PageSize = math.MaxInt
	if list, total, err := beeRegionService.GetBeeRegionInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		listPtr := lo.Map(list, func(item bee.BeeRegion, index int) *bee.BeeRegion {
			return &item
		})
		regionByPid := lo.GroupBy(listPtr, func(item *bee.BeeRegion) string {
			return item.Pid
		})
		beeRegionApi.buildTree(regionByPid[""], regionByPid)
		response.OkWithDetailed(response.PageResult{
			List:     regionByPid[""],
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// GetBeeRegionPublic 不需要鉴权的地址库接口
// @Tags BeeRegion
// @Summary 不需要鉴权的地址库接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeRegionSearch true "分页获取地址库列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeRegion/getBeeRegionPublic [get]
func (beeRegionApi *BeeRegionApi) GetBeeRegionPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的地址库接口信息",
	}, "获取成功", c)
}
