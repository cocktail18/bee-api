package api

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee"
	beeReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/model/bee/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/service"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/utils"
	beeUtils "github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BeeShopInfoApi struct{}

var beeShopInfoService = service.ServiceGroupApp.BeeServiceGroup.BeeShopInfoService

// CreateBeeShopInfo 创建商店信息
// @Tags BeeShopInfo
// @Summary 创建商店信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopInfo true "创建商店信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeShopInfo/createBeeShopInfo [post]
func (beeShopInfoApi *BeeShopInfoApi) CreateBeeShopInfo(c *gin.Context) {
	var beeShopInfo bee.BeeShopInfo
	err := c.ShouldBindJSON(&beeShopInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopInfo.UserId = &shopUserId
	if err := beeShopInfoService.CreateBeeShopInfo(&beeShopInfo); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeShopInfo 删除商店信息
// @Tags BeeShopInfo
// @Summary 删除商店信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopInfo true "删除商店信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeShopInfo/deleteBeeShopInfo [delete]
func (beeShopInfoApi *BeeShopInfoApi) DeleteBeeShopInfo(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopInfoService.DeleteBeeShopInfo(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeShopInfoByIds 批量删除商店信息
// @Tags BeeShopInfo
// @Summary 批量删除商店信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeShopInfo/deleteBeeShopInfoByIds [delete]
func (beeShopInfoApi *BeeShopInfoApi) DeleteBeeShopInfoByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeShopInfoService.DeleteBeeShopInfoByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeShopInfo 更新商店信息
// @Tags BeeShopInfo
// @Summary 更新商店信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeShopInfo true "更新商店信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeShopInfo/updateBeeShopInfo [put]
func (beeShopInfoApi *BeeShopInfoApi) UpdateBeeShopInfo(c *gin.Context) {
	var beeShopInfo bee.BeeShopInfo
	err := c.ShouldBindJSON(&beeShopInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeShopInfo.UserId = &shopUserId
	if err := beeShopInfoService.UpdateBeeShopInfo(beeShopInfo, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeShopInfo 用id查询商店信息
// @Tags BeeShopInfo
// @Summary 用id查询商店信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeShopInfo true "用id查询商店信息"
// @Success 200 {object} response.Response{data=object{rebeeShopInfo=bee.BeeShopInfo},msg=string} "查询成功"
// @Router /beeShopInfo/findBeeShopInfo [get]
func (beeShopInfoApi *BeeShopInfoApi) FindBeeShopInfo(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeShopInfo, err := beeShopInfoService.GetBeeShopInfo(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeShopInfo, c)
	}
}

// GetBeeShopInfoList 分页获取商店信息列表
// @Tags BeeShopInfo
// @Summary 分页获取商店信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopInfoSearch true "分页获取商店信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeShopInfo/getBeeShopInfoList [get]
func (beeShopInfoApi *BeeShopInfoApi) GetBeeShopInfoList(c *gin.Context) {
	var pageInfo beeReq.BeeShopInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeShopInfoService.GetBeeShopInfoInfoList(pageInfo, shopUserId, 0); err != nil {
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

// GetBeeShopInfoList 分页获取商店信息列表, 根据我的角色查询
// @Tags BeeShopInfo
// @Summary 分页获取商店信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopInfoSearch true "分页获取商店信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeShopInfo/getMyBeeShopInfoList [get]
func (beeShopInfoApi *BeeShopInfoApi) GetMyBeeShopInfoList(c *gin.Context) {
	var pageInfo beeReq.BeeShopInfoSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	userId := beeUtils.GetUserID(c)

	if list, total, err := beeShopInfoService.GetBeeShopInfoInfoList(pageInfo, shopUserId, userId); err != nil {
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

// GetBeeShopInfoPublic 不需要鉴权的商店信息接口
// @Tags BeeShopInfo
// @Summary 不需要鉴权的商店信息接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeShopInfoSearch true "分页获取商店信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeShopInfo/getBeeShopInfoPublic [get]
func (beeShopInfoApi *BeeShopInfoApi) GetBeeShopInfoPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的商店信息接口信息",
	}, "获取成功", c)
}

// getAllBeeShopInfos 获取商店信息列表
// @Tags BeeShopInfo
// @Summary 获取商店信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{data=response.Response,msg=string} "获取成功"
// @Router /beeShopInfo/getAllBeeShopInfos [get]
func (api *BeeShopInfoApi) GetAllBeeShopInfos(c *gin.Context) {

	if list, err := beeShopInfoService.GetAllBeeShopInfos(); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.Response{
			Data: list,
		}, "获取成功", c)
	}
}
