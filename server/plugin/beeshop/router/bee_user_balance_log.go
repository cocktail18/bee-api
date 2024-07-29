package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeUserBalanceLogRouter struct{}

// InitBeeUserBalanceLogRouter 初始化 用户消费记录 路由信息
func (s *BeeUserBalanceLogRouter) InitBeeUserBalanceLogRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeUserBalanceLogRouter := Router.Group("beeUserBalanceLog").Use(middleware.OperationRecord())
	beeUserBalanceLogRouterWithoutRecord := Router.Group("beeUserBalanceLog")
	beeUserBalanceLogRouterWithoutAuth := PublicRouter.Group("beeUserBalanceLog")

	var beeUserBalanceLogApi = api.BeeUserBalanceLogApi{}
	{
		beeUserBalanceLogRouter.POST("createBeeUserBalanceLog", beeUserBalanceLogApi.CreateBeeUserBalanceLog)             // 新建用户消费记录
		beeUserBalanceLogRouter.DELETE("deleteBeeUserBalanceLog", beeUserBalanceLogApi.DeleteBeeUserBalanceLog)           // 删除用户消费记录
		beeUserBalanceLogRouter.DELETE("deleteBeeUserBalanceLogByIds", beeUserBalanceLogApi.DeleteBeeUserBalanceLogByIds) // 批量删除用户消费记录
		beeUserBalanceLogRouter.PUT("updateBeeUserBalanceLog", beeUserBalanceLogApi.UpdateBeeUserBalanceLog)              // 更新用户消费记录
	}
	{
		beeUserBalanceLogRouterWithoutRecord.GET("findBeeUserBalanceLog", beeUserBalanceLogApi.FindBeeUserBalanceLog)       // 根据ID获取用户消费记录
		beeUserBalanceLogRouterWithoutRecord.GET("getBeeUserBalanceLogList", beeUserBalanceLogApi.GetBeeUserBalanceLogList) // 获取用户消费记录列表
	}
	{
		beeUserBalanceLogRouterWithoutAuth.GET("getBeeUserBalanceLogPublic", beeUserBalanceLogApi.GetBeeUserBalanceLogPublic) // 获取用户消费记录列表
	}
}
