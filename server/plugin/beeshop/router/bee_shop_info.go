package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeShopInfoRouter struct{}

// InitBeeShopInfoRouter 初始化 商店信息 路由信息
func (s *BeeShopInfoRouter) InitBeeShopInfoRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeShopInfoRouter := Router.Group("beeShopInfo").Use(middleware.OperationRecord())
	beeShopInfoRouterWithoutRecord := Router.Group("beeShopInfo")
	beeShopInfoRouterWithoutAuth := PublicRouter.Group("beeShopInfo")

	var beeShopInfoApi = api.BeeShopInfoApi{}
	{
		beeShopInfoRouter.POST("createBeeShopInfo", beeShopInfoApi.CreateBeeShopInfo)             // 新建商店信息
		beeShopInfoRouter.DELETE("deleteBeeShopInfo", beeShopInfoApi.DeleteBeeShopInfo)           // 删除商店信息
		beeShopInfoRouter.DELETE("deleteBeeShopInfoByIds", beeShopInfoApi.DeleteBeeShopInfoByIds) // 批量删除商店信息
		beeShopInfoRouter.PUT("updateBeeShopInfo", beeShopInfoApi.UpdateBeeShopInfo)              // 更新商店信息
	}
	{
		beeShopInfoRouterWithoutRecord.GET("findBeeShopInfo", beeShopInfoApi.FindBeeShopInfo)       // 根据ID获取商店信息
		beeShopInfoRouterWithoutRecord.GET("getBeeShopInfoList", beeShopInfoApi.GetBeeShopInfoList) // 获取商店信息列表
	}
	{
		beeShopInfoRouterWithoutAuth.GET("getBeeShopInfoPublic", beeShopInfoApi.GetBeeShopInfoPublic) // 获取商店信息列表
	}
}
