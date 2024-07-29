package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeRegionRouter struct{}

// InitBeeRegionRouter 初始化 地址库 路由信息
func (s *BeeRegionRouter) InitBeeRegionRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeRegionRouter := Router.Group("beeRegion").Use(middleware.OperationRecord())
	beeRegionRouterWithoutRecord := Router.Group("beeRegion")
	beeRegionRouterWithoutAuth := PublicRouter.Group("beeRegion")

	var beeRegionApi = api.BeeRegionApi{}
	{
		beeRegionRouter.POST("createBeeRegion", beeRegionApi.CreateBeeRegion)             // 新建地址库
		beeRegionRouter.DELETE("deleteBeeRegion", beeRegionApi.DeleteBeeRegion)           // 删除地址库
		beeRegionRouter.DELETE("deleteBeeRegionByIds", beeRegionApi.DeleteBeeRegionByIds) // 批量删除地址库
		beeRegionRouter.PUT("updateBeeRegion", beeRegionApi.UpdateBeeRegion)              // 更新地址库
	}
	{
		beeRegionRouterWithoutRecord.GET("findBeeRegion", beeRegionApi.FindBeeRegion)               // 根据ID获取地址库
		beeRegionRouterWithoutRecord.GET("getBeeRegionList", beeRegionApi.GetBeeRegionList)         // 获取地址库列表
		beeRegionRouterWithoutRecord.GET("getBeeRegionCityTree", beeRegionApi.GetBeeRegionCityTree) // 获取地址库省市区树
	}
	{
		beeRegionRouterWithoutAuth.GET("getBeeRegionPublic", beeRegionApi.GetBeeRegionPublic) // 获取地址库列表
	}
}
