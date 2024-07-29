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

type BeePeiSongApi struct{}

var beePeiSongService = service.ServiceGroupApp.BeeServiceGroup.BeePeiSongService

// CreateBeePeiSong 创建配送信息
// @Tags BeePeiSong
// @Summary 创建配送信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeePeiSong true "创建配送信息"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beePeiSong/createBeePeiSong [post]
func (beePeiSongApi *BeePeiSongApi) CreateBeePeiSong(c *gin.Context) {
	var beePeiSong bee.BeePeiSong
	err := c.ShouldBindJSON(&beePeiSong)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beePeiSong.UserId = &shopUserId
	if err := beePeiSongService.CreateBeePeiSong(&beePeiSong); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeePeiSong 删除配送信息
// @Tags BeePeiSong
// @Summary 删除配送信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeePeiSong true "删除配送信息"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beePeiSong/deleteBeePeiSong [delete]
func (beePeiSongApi *BeePeiSongApi) DeleteBeePeiSong(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beePeiSongService.DeleteBeePeiSong(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeePeiSongByIds 批量删除配送信息
// @Tags BeePeiSong
// @Summary 批量删除配送信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beePeiSong/deleteBeePeiSongByIds [delete]
func (beePeiSongApi *BeePeiSongApi) DeleteBeePeiSongByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beePeiSongService.DeleteBeePeiSongByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeePeiSong 更新配送信息
// @Tags BeePeiSong
// @Summary 更新配送信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeePeiSong true "更新配送信息"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beePeiSong/updateBeePeiSong [put]
func (beePeiSongApi *BeePeiSongApi) UpdateBeePeiSong(c *gin.Context) {
	var beePeiSong bee.BeePeiSong
	err := c.ShouldBindJSON(&beePeiSong)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beePeiSong.UserId = &shopUserId
	if err := beePeiSongService.UpdateBeePeiSong(beePeiSong, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeePeiSong 用id查询配送信息
// @Tags BeePeiSong
// @Summary 用id查询配送信息
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeePeiSong true "用id查询配送信息"
// @Success 200 {object} response.Response{data=object{rebeePeiSong=bee.BeePeiSong},msg=string} "查询成功"
// @Router /beePeiSong/findBeePeiSong [get]
func (beePeiSongApi *BeePeiSongApi) FindBeePeiSong(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeePeiSong, err := beePeiSongService.GetBeePeiSong(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeePeiSong, c)
	}
}

// GetBeePeiSongList 分页获取配送信息列表
// @Tags BeePeiSong
// @Summary 分页获取配送信息列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeePeiSongSearch true "分页获取配送信息列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beePeiSong/getBeePeiSongList [get]
func (beePeiSongApi *BeePeiSongApi) GetBeePeiSongList(c *gin.Context) {
	var pageInfo beeReq.BeePeiSongSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beePeiSongService.GetBeePeiSongInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeePeiSongPublic 不需要鉴权的配送信息接口
// @Tags BeePeiSong
// @Summary 不需要鉴权的配送信息接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeePeiSongSearch true "分页获取配送信息列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beePeiSong/getBeePeiSongPublic [get]
func (beePeiSongApi *BeePeiSongApi) GetBeePeiSongPublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的配送信息接口信息",
	}, "获取成功", c)
}
