package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeRechargeSendRuleRouter struct{}

// InitBeeRechargeSendRuleRouter 初始化 beeRechargeSendRule表 路由信息
func (s *BeeRechargeSendRuleRouter) InitBeeRechargeSendRuleRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeRechargeSendRuleRouter := Router.Group("beeRechargeSendRule").Use(middleware.OperationRecord())
	beeRechargeSendRuleRouterWithoutRecord := Router.Group("beeRechargeSendRule")
	beeRechargeSendRuleRouterWithoutAuth := PublicRouter.Group("beeRechargeSendRule")

	var beeRechargeSendRuleApi = api.BeeRechargeSendRuleApi{}
	{
		beeRechargeSendRuleRouter.POST("createBeeRechargeSendRule", beeRechargeSendRuleApi.CreateBeeRechargeSendRule)             // 新建beeRechargeSendRule表
		beeRechargeSendRuleRouter.DELETE("deleteBeeRechargeSendRule", beeRechargeSendRuleApi.DeleteBeeRechargeSendRule)           // 删除beeRechargeSendRule表
		beeRechargeSendRuleRouter.DELETE("deleteBeeRechargeSendRuleByIds", beeRechargeSendRuleApi.DeleteBeeRechargeSendRuleByIds) // 批量删除beeRechargeSendRule表
		beeRechargeSendRuleRouter.PUT("updateBeeRechargeSendRule", beeRechargeSendRuleApi.UpdateBeeRechargeSendRule)              // 更新beeRechargeSendRule表
	}
	{
		beeRechargeSendRuleRouterWithoutRecord.GET("findBeeRechargeSendRule", beeRechargeSendRuleApi.FindBeeRechargeSendRule)       // 根据ID获取beeRechargeSendRule表
		beeRechargeSendRuleRouterWithoutRecord.GET("getBeeRechargeSendRuleList", beeRechargeSendRuleApi.GetBeeRechargeSendRuleList) // 获取beeRechargeSendRule表列表
	}
	{
		beeRechargeSendRuleRouterWithoutAuth.GET("getBeeRechargeSendRulePublic", beeRechargeSendRuleApi.GetBeeRechargeSendRulePublic) // 获取beeRechargeSendRule表列表
	}
}
