package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeWxConfigRouter struct{}

// InitBeeWxConfigRouter 初始化 微信配置 路由信息
func (s *BeeWxConfigRouter) InitBeeWxConfigRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeWxConfigRouter := Router.Group("beeWxConfig").Use(middleware.OperationRecord())
	beeWxConfigRouterWithoutRecord := Router.Group("beeWxConfig")
	beeWxConfigRouterWithoutAuth := PublicRouter.Group("beeWxConfig")

	var beeWxConfigApi = api.BeeWxConfigApi{}
	{
		beeWxConfigRouter.POST("createBeeWxConfig", beeWxConfigApi.CreateBeeWxConfig)             // 新建微信配置
		beeWxConfigRouter.DELETE("deleteBeeWxConfig", beeWxConfigApi.DeleteBeeWxConfig)           // 删除微信配置
		beeWxConfigRouter.DELETE("deleteBeeWxConfigByIds", beeWxConfigApi.DeleteBeeWxConfigByIds) // 批量删除微信配置
		beeWxConfigRouter.PUT("updateBeeWxConfig", beeWxConfigApi.UpdateBeeWxConfig)              // 更新微信配置
	}
	{
		beeWxConfigRouterWithoutRecord.GET("findBeeWxConfig", beeWxConfigApi.FindBeeWxConfig)       // 根据ID获取微信配置
		beeWxConfigRouterWithoutRecord.GET("getBeeWxConfigList", beeWxConfigApi.GetBeeWxConfigList) // 获取微信配置列表
	}
	{
		beeWxConfigRouterWithoutAuth.GET("getBeeWxConfigPublic", beeWxConfigApi.GetBeeWxConfigPublic) // 获取微信配置列表
	}
}
