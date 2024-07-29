package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeNoticeRouter struct{}

// InitBeeNoticeRouter 初始化 系统公告 路由信息
func (s *BeeNoticeRouter) InitBeeNoticeRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeNoticeRouter := Router.Group("beeNotice").Use(middleware.OperationRecord())
	beeNoticeRouterWithoutRecord := Router.Group("beeNotice")
	beeNoticeRouterWithoutAuth := PublicRouter.Group("beeNotice")

	var beeNoticeApi = api.BeeNoticeApi{}
	{
		beeNoticeRouter.POST("createBeeNotice", beeNoticeApi.CreateBeeNotice)             // 新建系统公告
		beeNoticeRouter.DELETE("deleteBeeNotice", beeNoticeApi.DeleteBeeNotice)           // 删除系统公告
		beeNoticeRouter.DELETE("deleteBeeNoticeByIds", beeNoticeApi.DeleteBeeNoticeByIds) // 批量删除系统公告
		beeNoticeRouter.PUT("updateBeeNotice", beeNoticeApi.UpdateBeeNotice)              // 更新系统公告
	}
	{
		beeNoticeRouterWithoutRecord.GET("findBeeNotice", beeNoticeApi.FindBeeNotice)       // 根据ID获取系统公告
		beeNoticeRouterWithoutRecord.GET("getBeeNoticeList", beeNoticeApi.GetBeeNoticeList) // 获取系统公告列表
	}
	{
		beeNoticeRouterWithoutAuth.GET("getBeeNoticePublic", beeNoticeApi.GetBeeNoticePublic) // 获取系统公告列表
	}
}
