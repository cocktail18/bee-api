package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeQueueRouter struct{}

// InitBeeQueueRouter 初始化 beeQueue表 路由信息
func (s *BeeQueueRouter) InitBeeQueueRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeQueueRouter := Router.Group("beeQueue").Use(middleware.OperationRecord())
	beeQueueRouterWithoutRecord := Router.Group("beeQueue")
	beeQueueRouterWithoutAuth := PublicRouter.Group("beeQueue")

	var beeQueueApi = api.BeeQueueApi{}
	{
		beeQueueRouter.POST("createBeeQueue", beeQueueApi.CreateBeeQueue)             // 新建beeQueue表
		beeQueueRouter.DELETE("deleteBeeQueue", beeQueueApi.DeleteBeeQueue)           // 删除beeQueue表
		beeQueueRouter.DELETE("deleteBeeQueueByIds", beeQueueApi.DeleteBeeQueueByIds) // 批量删除beeQueue表
		beeQueueRouter.PUT("updateBeeQueue", beeQueueApi.UpdateBeeQueue)              // 更新beeQueue表
		beeQueueRouter.PUT("callBeeQueue", beeQueueApi.CallBeeQueue)                  // 叫号
		beeQueueRouter.PUT("reCallBeeQueue", beeQueueApi.ReCallBeeQueue)              // 重新叫号
		beeQueueRouter.PUT("passBeeQueue", beeQueueApi.PassBeeQueue)                  // 过号
		beeQueueRouter.PUT("nextBeeQueue", beeQueueApi.NextBeeQueue)                  // 下一号
	}
	{
		beeQueueRouterWithoutRecord.GET("findBeeQueue", beeQueueApi.FindBeeQueue)       // 根据ID获取beeQueue表
		beeQueueRouterWithoutRecord.GET("getBeeQueueList", beeQueueApi.GetBeeQueueList) // 获取beeQueue表列表
	}
	{
		beeQueueRouterWithoutAuth.GET("getBeeQueuePublic", beeQueueApi.GetBeeQueuePublic) // 获取beeQueue表列表
	}
}
