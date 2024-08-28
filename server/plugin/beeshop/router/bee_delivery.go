package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeDeliveryRouter struct{}

// InitBeeDeliveryRouter 初始化 beeDelivery表 路由信息
func (s *BeeDeliveryRouter) InitBeeDeliveryRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeDeliveryRouter := Router.Group("beeDelivery").Use(middleware.OperationRecord())
	beeDeliveryRouterWithoutRecord := Router.Group("beeDelivery")
	beeDeliveryRouterWithoutAuth := PublicRouter.Group("beeDelivery")

	var beeDeliveryApi = api.BeeDeliveryApi{}
	{
		beeDeliveryRouter.POST("createBeeDelivery", beeDeliveryApi.CreateBeeDelivery)             // 新建beeDelivery表
		beeDeliveryRouter.DELETE("deleteBeeDelivery", beeDeliveryApi.DeleteBeeDelivery)           // 删除beeDelivery表
		beeDeliveryRouter.DELETE("deleteBeeDeliveryByIds", beeDeliveryApi.DeleteBeeDeliveryByIds) // 批量删除beeDelivery表
		beeDeliveryRouter.PUT("updateBeeDelivery", beeDeliveryApi.UpdateBeeDelivery)              // 更新beeDelivery表
	}
	{
		beeDeliveryRouterWithoutRecord.GET("findBeeDelivery", beeDeliveryApi.FindBeeDelivery)       // 根据ID获取beeDelivery表
		beeDeliveryRouterWithoutRecord.GET("getBeeDeliveryList", beeDeliveryApi.GetBeeDeliveryList) // 获取beeDelivery表列表
	}
	{
		beeDeliveryRouterWithoutAuth.GET("getBeeDeliveryPublic", beeDeliveryApi.GetBeeDeliveryPublic) // 获取beeDelivery表列表
	}
}
