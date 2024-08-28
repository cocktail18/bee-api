package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeOrderRouter struct{}

// InitBeeOrderRouter 初始化 用户订单 路由信息
func (s *BeeOrderRouter) InitBeeOrderRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeOrderRouter := Router.Group("beeOrder").Use(middleware.OperationRecord())
	beeOrderRouterWithoutRecord := Router.Group("beeOrder")
	beeOrderRouterWithoutAuth := PublicRouter.Group("beeOrder")

	var beeOrderApi = api.BeeOrderApi{}
	{
		beeOrderRouter.POST("createBeeOrder", beeOrderApi.CreateBeeOrder)                    // 新建用户订单
		beeOrderRouter.DELETE("deleteBeeOrder", beeOrderApi.DeleteBeeOrder)                  // 删除用户订单
		beeOrderRouter.DELETE("deleteBeeOrderByIds", beeOrderApi.DeleteBeeOrderByIds)        // 批量删除用户订单
		beeOrderRouter.PUT("updateBeeOrder", beeOrderApi.UpdateBeeOrder)                     // 更新用户订单
		beeOrderRouter.PUT("updateBeeOrderExtJsonStr", beeOrderApi.UpdateBeeOrderExtJsonStr) // 更新用户订单extJsonStr字段
		beeOrderRouter.PUT("updateBeeOrderStatus", beeOrderApi.UpdateBeeOrderStatus)         // 更新用户订单status字段
		beeOrderRouter.PUT("markBeeOrderDone", beeOrderApi.MarkBeeOrderDone)                 // 设置为已完成订单
		beeOrderRouter.PUT("markBeeOrderPaid", beeOrderApi.MarkBeeOrderPaid)                 // 设置为已支付订单
		beeOrderRouter.PUT("shippedBeeOrder", beeOrderApi.ShippedBeeOrder)                   // 设置为已发货订单
	}
	{
		beeOrderRouterWithoutRecord.GET("findBeeOrder", beeOrderApi.FindBeeOrder)       // 根据ID获取用户订单
		beeOrderRouterWithoutRecord.GET("getBeeOrderList", beeOrderApi.GetBeeOrderList) // 获取用户订单列表
	}
	{
		beeOrderRouterWithoutAuth.GET("getBeeOrderPublic", beeOrderApi.GetBeeOrderPublic) // 获取用户订单列表
	}
}
