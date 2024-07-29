package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeBannerRouter struct{}

// InitBeeBannerRouter 初始化 商店banner 路由信息
func (s *BeeBannerRouter) InitBeeBannerRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeBannerRouter := Router.Group("beeBanner").Use(middleware.OperationRecord())
	beeBannerRouterWithoutRecord := Router.Group("beeBanner")
	beeBannerRouterWithoutAuth := PublicRouter.Group("beeBanner")

	var beeBannerApi = api.BeeBannerApi{}
	{
		beeBannerRouter.POST("createBeeBanner", beeBannerApi.CreateBeeBanner)             // 新建商店banner
		beeBannerRouter.DELETE("deleteBeeBanner", beeBannerApi.DeleteBeeBanner)           // 删除商店banner
		beeBannerRouter.DELETE("deleteBeeBannerByIds", beeBannerApi.DeleteBeeBannerByIds) // 批量删除商店banner
		beeBannerRouter.PUT("updateBeeBanner", beeBannerApi.UpdateBeeBanner)              // 更新商店banner
	}
	{
		beeBannerRouterWithoutRecord.GET("findBeeBanner", beeBannerApi.FindBeeBanner)       // 根据ID获取商店banner
		beeBannerRouterWithoutRecord.GET("getBeeBannerList", beeBannerApi.GetBeeBannerList) // 获取商店banner列表
	}
	{
		beeBannerRouterWithoutAuth.GET("getBeeBannerPublic", beeBannerApi.GetBeeBannerPublic) // 获取商店banner列表
	}
}
