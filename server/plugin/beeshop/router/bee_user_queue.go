package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeUserQueueRouter struct{}

// InitBeeUserQueueRouter 初始化 beeUserQueue表 路由信息
func (s *BeeUserQueueRouter) InitBeeUserQueueRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeUserQueueRouter := Router.Group("beeUserQueue").Use(middleware.OperationRecord())
	beeUserQueueRouterWithoutRecord := Router.Group("beeUserQueue")
	beeUserQueueRouterWithoutAuth := PublicRouter.Group("beeUserQueue")

	var beeUserQueueApi = api.BeeUserQueueApi{}
	{
		beeUserQueueRouter.POST("createBeeUserQueue", beeUserQueueApi.CreateBeeUserQueue)             // 新建beeUserQueue表
		beeUserQueueRouter.DELETE("deleteBeeUserQueue", beeUserQueueApi.DeleteBeeUserQueue)           // 删除beeUserQueue表
		beeUserQueueRouter.DELETE("deleteBeeUserQueueByIds", beeUserQueueApi.DeleteBeeUserQueueByIds) // 批量删除beeUserQueue表
		beeUserQueueRouter.PUT("updateBeeUserQueue", beeUserQueueApi.UpdateBeeUserQueue)              // 更新beeUserQueue表
	}
	{
		beeUserQueueRouterWithoutRecord.GET("findBeeUserQueue", beeUserQueueApi.FindBeeUserQueue)       // 根据ID获取beeUserQueue表
		beeUserQueueRouterWithoutRecord.GET("getBeeUserQueueList", beeUserQueueApi.GetBeeUserQueueList) // 获取beeUserQueue表列表
	}
	{
		beeUserQueueRouterWithoutAuth.GET("getBeeUserQueuePublic", beeUserQueueApi.GetBeeUserQueuePublic) // 获取beeUserQueue表列表
	}
}
