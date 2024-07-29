package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeUserAmountRouter struct{}

// InitBeeUserAmountRouter 初始化 用户资产 路由信息
func (s *BeeUserAmountRouter) InitBeeUserAmountRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeUserAmountRouter := Router.Group("beeUserAmount").Use(middleware.OperationRecord())
	beeUserAmountRouterWithoutRecord := Router.Group("beeUserAmount")
	beeUserAmountRouterWithoutAuth := PublicRouter.Group("beeUserAmount")

	var beeUserAmountApi = api.BeeUserAmountApi{}
	{
		beeUserAmountRouter.POST("createBeeUserAmount", beeUserAmountApi.CreateBeeUserAmount)             // 新建用户资产
		beeUserAmountRouter.DELETE("deleteBeeUserAmount", beeUserAmountApi.DeleteBeeUserAmount)           // 删除用户资产
		beeUserAmountRouter.DELETE("deleteBeeUserAmountByIds", beeUserAmountApi.DeleteBeeUserAmountByIds) // 批量删除用户资产
		beeUserAmountRouter.PUT("updateBeeUserAmount", beeUserAmountApi.UpdateBeeUserAmount)              // 更新用户资产
	}
	{
		beeUserAmountRouterWithoutRecord.GET("findBeeUserAmount", beeUserAmountApi.FindBeeUserAmount)       // 根据ID获取用户资产
		beeUserAmountRouterWithoutRecord.GET("getBeeUserAmountList", beeUserAmountApi.GetBeeUserAmountList) // 获取用户资产列表
	}
	{
		beeUserAmountRouterWithoutAuth.GET("getBeeUserAmountPublic", beeUserAmountApi.GetBeeUserAmountPublic) // 获取用户资产列表
	}
}
