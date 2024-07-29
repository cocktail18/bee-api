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

type BeeUploadFileApi struct{}

var beeUploadFileService = service.ServiceGroupApp.BeeServiceGroup.BeeUploadFileService

// CreateBeeUploadFile 创建用户上传文件
// @Tags BeeUploadFile
// @Summary 创建用户上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUploadFile true "创建用户上传文件"
// @Success 200 {object} response.Response{msg=string} "创建成功"
// @Router /beeUploadFile/createBeeUploadFile [post]
func (beeUploadFileApi *BeeUploadFileApi) CreateBeeUploadFile(c *gin.Context) {
	var beeUploadFile bee.BeeUploadFile
	err := c.ShouldBindJSON(&beeUploadFile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUploadFile.UserId = &shopUserId
	if err := beeUploadFileService.CreateBeeUploadFile(&beeUploadFile); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteBeeUploadFile 删除用户上传文件
// @Tags BeeUploadFile
// @Summary 删除用户上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUploadFile true "删除用户上传文件"
// @Success 200 {object} response.Response{msg=string} "删除成功"
// @Router /beeUploadFile/deleteBeeUploadFile [delete]
func (beeUploadFileApi *BeeUploadFileApi) DeleteBeeUploadFile(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUploadFileService.DeleteBeeUploadFile(id, shopUserId); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteBeeUploadFileByIds 批量删除用户上传文件
// @Tags BeeUploadFile
// @Summary 批量删除用户上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Success 200 {object} response.Response{msg=string} "批量删除成功"
// @Router /beeUploadFile/deleteBeeUploadFileByIds [delete]
func (beeUploadFileApi *BeeUploadFileApi) DeleteBeeUploadFileByIds(c *gin.Context) {
	ids := c.QueryArray("ids[]")
	shopUserId := int(utils.GetShopUserID(c))
	if err := beeUploadFileService.DeleteBeeUploadFileByIds(ids, shopUserId); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateBeeUploadFile 更新用户上传文件
// @Tags BeeUploadFile
// @Summary 更新用户上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body bee.BeeUploadFile true "更新用户上传文件"
// @Success 200 {object} response.Response{msg=string} "更新成功"
// @Router /beeUploadFile/updateBeeUploadFile [put]
func (beeUploadFileApi *BeeUploadFileApi) UpdateBeeUploadFile(c *gin.Context) {
	var beeUploadFile bee.BeeUploadFile
	err := c.ShouldBindJSON(&beeUploadFile)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	beeUploadFile.UserId = &shopUserId
	if err := beeUploadFileService.UpdateBeeUploadFile(beeUploadFile, shopUserId); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindBeeUploadFile 用id查询用户上传文件
// @Tags BeeUploadFile
// @Summary 用id查询用户上传文件
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query bee.BeeUploadFile true "用id查询用户上传文件"
// @Success 200 {object} response.Response{data=object{rebeeUploadFile=bee.BeeUploadFile},msg=string} "查询成功"
// @Router /beeUploadFile/findBeeUploadFile [get]
func (beeUploadFileApi *BeeUploadFileApi) FindBeeUploadFile(c *gin.Context) {
	id := c.Query("id")
	shopUserId := int(utils.GetShopUserID(c))
	if rebeeUploadFile, err := beeUploadFileService.GetBeeUploadFile(id, shopUserId); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(rebeeUploadFile, c)
	}
}

// GetBeeUploadFileList 分页获取用户上传文件列表
// @Tags BeeUploadFile
// @Summary 分页获取用户上传文件列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUploadFileSearch true "分页获取用户上传文件列表"
// @Success 200 {object} response.Response{data=response.PageResult,msg=string} "获取成功"
// @Router /beeUploadFile/getBeeUploadFileList [get]
func (beeUploadFileApi *BeeUploadFileApi) GetBeeUploadFileList(c *gin.Context) {
	var pageInfo beeReq.BeeUploadFileSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	shopUserId := int(utils.GetShopUserID(c))
	if list, total, err := beeUploadFileService.GetBeeUploadFileInfoList(pageInfo, shopUserId); err != nil {
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

// GetBeeUploadFilePublic 不需要鉴权的用户上传文件接口
// @Tags BeeUploadFile
// @Summary 不需要鉴权的用户上传文件接口
// @accept application/json
// @Produce application/json
// @Param data query beeReq.BeeUploadFileSearch true "分页获取用户上传文件列表"
// @Success 200 {object} response.Response{data=object,msg=string} "获取成功"
// @Router /beeUploadFile/getBeeUploadFilePublic [get]
func (beeUploadFileApi *BeeUploadFileApi) GetBeeUploadFilePublic(c *gin.Context) {
	// 此接口不需要鉴权
	// 示例为返回了一个固定的消息接口，一般本接口用于C端服务，需要自己实现业务逻辑
	response.OkWithDetailed(gin.H{
		"info": "不需要鉴权的用户上传文件接口信息",
	}, "获取成功", c)
}
