package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeConfigRouter struct{}

// InitBeeConfigRouter 初始化 公用配置表 路由信息
func (s *BeeConfigRouter) InitBeeConfigRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeConfigRouter := Router.Group("beeConfig").Use(middleware.OperationRecord())
	beeConfigRouterWithoutRecord := Router.Group("beeConfig")
	beeConfigRouterWithoutAuth := PublicRouter.Group("beeConfig")

	var beeConfigApi = api.BeeConfigApi{}
	{
		beeConfigRouter.POST("createBeeConfig", beeConfigApi.CreateBeeConfig)             // 新建公用配置表
		beeConfigRouter.DELETE("deleteBeeConfig", beeConfigApi.DeleteBeeConfig)           // 删除公用配置表
		beeConfigRouter.DELETE("deleteBeeConfigByIds", beeConfigApi.DeleteBeeConfigByIds) // 批量删除公用配置表
		beeConfigRouter.PUT("updateBeeConfig", beeConfigApi.UpdateBeeConfig)              // 更新公用配置表
	}
	{
		beeConfigRouterWithoutRecord.GET("findBeeConfig", beeConfigApi.FindBeeConfig)       // 根据ID获取公用配置表
		beeConfigRouterWithoutRecord.GET("getBeeConfigList", beeConfigApi.GetBeeConfigList) // 获取公用配置表列表
	}
	{
		beeConfigRouterWithoutAuth.GET("getBeeConfigPublic", beeConfigApi.GetBeeConfigPublic) // 获取公用配置表列表
	}
}
