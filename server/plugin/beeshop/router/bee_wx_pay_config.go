package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeWxPayConfigRouter struct{}

// InitBeeWxPayConfigRouter 初始化 微信支付配置 路由信息
func (s *BeeWxPayConfigRouter) InitBeeWxPayConfigRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeWxPayConfigRouter := Router.Group("beeWxPayConfig").Use(middleware.OperationRecord())
	beeWxPayConfigRouterWithoutRecord := Router.Group("beeWxPayConfig")
	beeWxPayConfigRouterWithoutAuth := PublicRouter.Group("beeWxPayConfig")

	var beeWxPayConfigApi = api.BeeWxPayConfigApi{}
	{
		beeWxPayConfigRouter.POST("createBeeWxPayConfig", beeWxPayConfigApi.CreateBeeWxPayConfig)             // 新建微信支付配置
		beeWxPayConfigRouter.DELETE("deleteBeeWxPayConfig", beeWxPayConfigApi.DeleteBeeWxPayConfig)           // 删除微信支付配置
		beeWxPayConfigRouter.DELETE("deleteBeeWxPayConfigByIds", beeWxPayConfigApi.DeleteBeeWxPayConfigByIds) // 批量删除微信支付配置
		beeWxPayConfigRouter.PUT("updateBeeWxPayConfig", beeWxPayConfigApi.UpdateBeeWxPayConfig)              // 更新微信支付配置
	}
	{
		beeWxPayConfigRouterWithoutRecord.GET("findBeeWxPayConfig", beeWxPayConfigApi.FindBeeWxPayConfig)       // 根据ID获取微信支付配置
		beeWxPayConfigRouterWithoutRecord.GET("getBeeWxPayConfigList", beeWxPayConfigApi.GetBeeWxPayConfigList) // 获取微信支付配置列表
	}
	{
		beeWxPayConfigRouterWithoutAuth.GET("getBeeWxPayConfigPublic", beeWxPayConfigApi.GetBeeWxPayConfigPublic) // 获取微信支付配置列表
	}
}
