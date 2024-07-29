package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeCashLogRouter struct{}

// InitBeeCashLogRouter 初始化 用户现金消费记录 路由信息
func (s *BeeCashLogRouter) InitBeeCashLogRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeCashLogRouter := Router.Group("beeCashLog").Use(middleware.OperationRecord())
	beeCashLogRouterWithoutRecord := Router.Group("beeCashLog")
	beeCashLogRouterWithoutAuth := PublicRouter.Group("beeCashLog")

	var beeCashLogApi = api.BeeCashLogApi{}
	{
		beeCashLogRouter.POST("createBeeCashLog", beeCashLogApi.CreateBeeCashLog)             // 新建用户现金消费记录
		beeCashLogRouter.DELETE("deleteBeeCashLog", beeCashLogApi.DeleteBeeCashLog)           // 删除用户现金消费记录
		beeCashLogRouter.DELETE("deleteBeeCashLogByIds", beeCashLogApi.DeleteBeeCashLogByIds) // 批量删除用户现金消费记录
		beeCashLogRouter.PUT("updateBeeCashLog", beeCashLogApi.UpdateBeeCashLog)              // 更新用户现金消费记录
	}
	{
		beeCashLogRouterWithoutRecord.GET("findBeeCashLog", beeCashLogApi.FindBeeCashLog)       // 根据ID获取用户现金消费记录
		beeCashLogRouterWithoutRecord.GET("getBeeCashLogList", beeCashLogApi.GetBeeCashLogList) // 获取用户现金消费记录列表
	}
	{
		beeCashLogRouterWithoutAuth.GET("getBeeCashLogPublic", beeCashLogApi.GetBeeCashLogPublic) // 获取用户现金消费记录列表
	}
}
