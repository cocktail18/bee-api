package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeOrderLogisticsRouter struct{}

// InitBeeOrderLogisticsRouter 初始化 用户订单地址 路由信息
func (s *BeeOrderLogisticsRouter) InitBeeOrderLogisticsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeOrderLogisticsRouter := Router.Group("beeOrderLogistics").Use(middleware.OperationRecord())
	beeOrderLogisticsRouterWithoutRecord := Router.Group("beeOrderLogistics")
	beeOrderLogisticsRouterWithoutAuth := PublicRouter.Group("beeOrderLogistics")

	var beeOrderLogisticsApi = api.BeeOrderLogisticsApi{}
	{
		beeOrderLogisticsRouter.POST("createBeeOrderLogistics", beeOrderLogisticsApi.CreateBeeOrderLogistics)             // 新建用户订单地址
		beeOrderLogisticsRouter.DELETE("deleteBeeOrderLogistics", beeOrderLogisticsApi.DeleteBeeOrderLogistics)           // 删除用户订单地址
		beeOrderLogisticsRouter.DELETE("deleteBeeOrderLogisticsByIds", beeOrderLogisticsApi.DeleteBeeOrderLogisticsByIds) // 批量删除用户订单地址
		beeOrderLogisticsRouter.PUT("updateBeeOrderLogistics", beeOrderLogisticsApi.UpdateBeeOrderLogistics)              // 更新用户订单地址
	}
	{
		beeOrderLogisticsRouterWithoutRecord.GET("findBeeOrderLogistics", beeOrderLogisticsApi.FindBeeOrderLogistics)       // 根据ID获取用户订单地址
		beeOrderLogisticsRouterWithoutRecord.GET("getBeeOrderLogisticsList", beeOrderLogisticsApi.GetBeeOrderLogisticsList) // 获取用户订单地址列表
	}
	{
		beeOrderLogisticsRouterWithoutAuth.GET("getBeeOrderLogisticsPublic", beeOrderLogisticsApi.GetBeeOrderLogisticsPublic) // 获取用户订单地址列表
	}
}
