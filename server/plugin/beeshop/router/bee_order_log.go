package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeOrderLogRouter struct{}

// InitBeeOrderLogRouter 初始化 订单日志 路由信息
func (s *BeeOrderLogRouter) InitBeeOrderLogRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeOrderLogRouter := Router.Group("beeOrderLog").Use(middleware.OperationRecord())
	beeOrderLogRouterWithoutRecord := Router.Group("beeOrderLog")
	beeOrderLogRouterWithoutAuth := PublicRouter.Group("beeOrderLog")

	var beeOrderLogApi = api.BeeOrderLogApi{}
	{
		beeOrderLogRouter.POST("createBeeOrderLog", beeOrderLogApi.CreateBeeOrderLog)             // 新建订单日志
		beeOrderLogRouter.DELETE("deleteBeeOrderLog", beeOrderLogApi.DeleteBeeOrderLog)           // 删除订单日志
		beeOrderLogRouter.DELETE("deleteBeeOrderLogByIds", beeOrderLogApi.DeleteBeeOrderLogByIds) // 批量删除订单日志
		beeOrderLogRouter.PUT("updateBeeOrderLog", beeOrderLogApi.UpdateBeeOrderLog)              // 更新订单日志
	}
	{
		beeOrderLogRouterWithoutRecord.GET("findBeeOrderLog", beeOrderLogApi.FindBeeOrderLog)       // 根据ID获取订单日志
		beeOrderLogRouterWithoutRecord.GET("getBeeOrderLogList", beeOrderLogApi.GetBeeOrderLogList) // 获取订单日志列表
	}
	{
		beeOrderLogRouterWithoutAuth.GET("getBeeOrderLogPublic", beeOrderLogApi.GetBeeOrderLogPublic) // 获取订单日志列表
	}
}
