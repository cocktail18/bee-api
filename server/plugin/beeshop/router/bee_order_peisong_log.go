package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeOrderPeisongLogRouter struct{}

// InitBeeOrderPeisongLogRouter 初始化 beeOrderPeisongLog表 路由信息
func (s *BeeOrderPeisongLogRouter) InitBeeOrderPeisongLogRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeOrderPeisongLogRouter := Router.Group("beeOrderPeisongLog").Use(middleware.OperationRecord())
	beeOrderPeisongLogRouterWithoutRecord := Router.Group("beeOrderPeisongLog")
	beeOrderPeisongLogRouterWithoutAuth := PublicRouter.Group("beeOrderPeisongLog")

	var beeOrderPeisongLogApi = api.BeeOrderPeisongLogApi{}
	{
		beeOrderPeisongLogRouter.POST("createBeeOrderPeisongLog", beeOrderPeisongLogApi.CreateBeeOrderPeisongLog)             // 新建beeOrderPeisongLog表
		beeOrderPeisongLogRouter.DELETE("deleteBeeOrderPeisongLog", beeOrderPeisongLogApi.DeleteBeeOrderPeisongLog)           // 删除beeOrderPeisongLog表
		beeOrderPeisongLogRouter.DELETE("deleteBeeOrderPeisongLogByIds", beeOrderPeisongLogApi.DeleteBeeOrderPeisongLogByIds) // 批量删除beeOrderPeisongLog表
		beeOrderPeisongLogRouter.PUT("updateBeeOrderPeisongLog", beeOrderPeisongLogApi.UpdateBeeOrderPeisongLog)              // 更新beeOrderPeisongLog表
	}
	{
		beeOrderPeisongLogRouterWithoutRecord.GET("findBeeOrderPeisongLog", beeOrderPeisongLogApi.FindBeeOrderPeisongLog)       // 根据ID获取beeOrderPeisongLog表
		beeOrderPeisongLogRouterWithoutRecord.GET("getBeeOrderPeisongLogList", beeOrderPeisongLogApi.GetBeeOrderPeisongLogList) // 获取beeOrderPeisongLog表列表
	}
	{
		beeOrderPeisongLogRouterWithoutAuth.GET("getBeeOrderPeisongLogPublic", beeOrderPeisongLogApi.GetBeeOrderPeisongLogPublic) // 获取beeOrderPeisongLog表列表
	}
}
