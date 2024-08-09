package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeePayLogRouter struct{}

// InitBeePayLogRouter 初始化 支付流水 路由信息
func (s *BeePayLogRouter) InitBeePayLogRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beePayLogRouter := Router.Group("beePayLog").Use(middleware.OperationRecord())
	beePayLogRouterWithoutRecord := Router.Group("beePayLog")
	beePayLogRouterWithoutAuth := PublicRouter.Group("beePayLog")

	var beePayLogApi = api.BeePayLogApi{}
	{
		beePayLogRouter.POST("createBeePayLog", beePayLogApi.CreateBeePayLog)             // 新建支付流水
		beePayLogRouter.DELETE("deleteBeePayLog", beePayLogApi.DeleteBeePayLog)           // 删除支付流水
		beePayLogRouter.DELETE("deleteBeePayLogByIds", beePayLogApi.DeleteBeePayLogByIds) // 批量删除支付流水
		beePayLogRouter.PUT("updateBeePayLog", beePayLogApi.UpdateBeePayLog)              // 更新支付流水
	}
	{
		beePayLogRouterWithoutRecord.GET("findBeePayLog", beePayLogApi.FindBeePayLog)       // 根据ID获取支付流水
		beePayLogRouterWithoutRecord.GET("getBeePayLogList", beePayLogApi.GetBeePayLogList) // 获取支付流水列表
		beePayLogRouterWithoutRecord.GET("getBeePayTotal", beePayLogApi.GetBeePayTotal)     // 获取支付总额
	}
	{
		beePayLogRouterWithoutAuth.GET("getBeePayLogPublic", beePayLogApi.GetBeePayLogPublic) // 获取支付流水列表
	}
}
