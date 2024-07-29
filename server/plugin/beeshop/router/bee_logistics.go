package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeLogisticsRouter struct{}

// InitBeeLogisticsRouter 初始化 运费模版 路由信息
func (s *BeeLogisticsRouter) InitBeeLogisticsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeLogisticsRouter := Router.Group("beeLogistics").Use(middleware.OperationRecord())
	beeLogisticsRouterWithoutRecord := Router.Group("beeLogistics")
	beeLogisticsRouterWithoutAuth := PublicRouter.Group("beeLogistics")

	var beeLogisticsApi = api.BeeLogisticsApi{}
	{
		beeLogisticsRouter.POST("createBeeLogistics", beeLogisticsApi.CreateBeeLogistics)             // 新建运费模版
		beeLogisticsRouter.DELETE("deleteBeeLogistics", beeLogisticsApi.DeleteBeeLogistics)           // 删除运费模版
		beeLogisticsRouter.DELETE("deleteBeeLogisticsByIds", beeLogisticsApi.DeleteBeeLogisticsByIds) // 批量删除运费模版
		beeLogisticsRouter.PUT("updateBeeLogistics", beeLogisticsApi.UpdateBeeLogistics)              // 更新运费模版
	}
	{
		beeLogisticsRouterWithoutRecord.GET("findBeeLogistics", beeLogisticsApi.FindBeeLogistics)       // 根据ID获取运费模版
		beeLogisticsRouterWithoutRecord.GET("getBeeLogisticsList", beeLogisticsApi.GetBeeLogisticsList) // 获取运费模版列表
	}
	{
		beeLogisticsRouterWithoutAuth.GET("getBeeLogisticsPublic", beeLogisticsApi.GetBeeLogisticsPublic) // 获取运费模版列表
	}
}
