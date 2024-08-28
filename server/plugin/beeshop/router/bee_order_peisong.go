package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeOrderPeisongRouter struct{}

// InitBeeOrderPeisongRouter 初始化 beeOrderPeisong表 路由信息
func (s *BeeOrderPeisongRouter) InitBeeOrderPeisongRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeOrderPeisongRouter := Router.Group("beeOrderPeisong").Use(middleware.OperationRecord())
	beeOrderPeisongRouterWithoutRecord := Router.Group("beeOrderPeisong")
	beeOrderPeisongRouterWithoutAuth := PublicRouter.Group("beeOrderPeisong")

	var beeOrderPeisongApi = api.BeeOrderPeisongApi{}
	{
		beeOrderPeisongRouter.POST("createBeeOrderPeisong", beeOrderPeisongApi.CreateBeeOrderPeisong)             // 新建beeOrderPeisong表
		beeOrderPeisongRouter.POST("cancelBeeOrderPeisong", beeOrderPeisongApi.CancelBeeOrderPeisong)             // 取消配送
		beeOrderPeisongRouter.POST("notifyBeeOrderPeisong", beeOrderPeisongApi.NotifyBeeOrderPeisong)             // 通知供应商配送
		beeOrderPeisongRouter.DELETE("deleteBeeOrderPeisong", beeOrderPeisongApi.DeleteBeeOrderPeisong)           // 删除beeOrderPeisong表
		beeOrderPeisongRouter.DELETE("deleteBeeOrderPeisongByIds", beeOrderPeisongApi.DeleteBeeOrderPeisongByIds) // 批量删除beeOrderPeisong表
		beeOrderPeisongRouter.PUT("updateBeeOrderPeisong", beeOrderPeisongApi.UpdateBeeOrderPeisong)              // 更新beeOrderPeisong表
	}
	{
		beeOrderPeisongRouterWithoutRecord.GET("findBeeOrderPeisong", beeOrderPeisongApi.FindBeeOrderPeisong)           // 根据ID获取beeOrderPeisong表
		beeOrderPeisongRouterWithoutRecord.GET("getBeeOrderPeisongList", beeOrderPeisongApi.GetBeeOrderPeisongList)     // 获取beeOrderPeisong表列表
		beeOrderPeisongRouterWithoutRecord.GET("getBeeOrderPeisongDetail", beeOrderPeisongApi.GetBeeOrderPeisongDetail) // 获取beeOrderPeisong表列表
	}
	{
		beeOrderPeisongRouterWithoutAuth.GET("getBeeOrderPeisongPublic", beeOrderPeisongApi.GetBeeOrderPeisongPublic) // 获取beeOrderPeisong表列表
	}
}
