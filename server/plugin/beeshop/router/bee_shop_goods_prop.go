package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeShopGoodsPropRouter struct{}

// InitBeeShopGoodsPropRouter 初始化 商品属性 路由信息
func (s *BeeShopGoodsPropRouter) InitBeeShopGoodsPropRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeShopGoodsPropRouter := Router.Group("beeShopGoodsProp").Use(middleware.OperationRecord())
	beeShopGoodsPropRouterWithoutRecord := Router.Group("beeShopGoodsProp")
	beeShopGoodsPropRouterWithoutAuth := PublicRouter.Group("beeShopGoodsProp")

	var beeShopGoodsPropApi = api.BeeShopGoodsPropApi{}
	{
		beeShopGoodsPropRouter.POST("createBeeShopGoodsProp", beeShopGoodsPropApi.CreateBeeShopGoodsProp)             // 新建商品属性
		beeShopGoodsPropRouter.DELETE("deleteBeeShopGoodsProp", beeShopGoodsPropApi.DeleteBeeShopGoodsProp)           // 删除商品属性
		beeShopGoodsPropRouter.DELETE("deleteBeeShopGoodsPropByIds", beeShopGoodsPropApi.DeleteBeeShopGoodsPropByIds) // 批量删除商品属性
		beeShopGoodsPropRouter.PUT("updateBeeShopGoodsProp", beeShopGoodsPropApi.UpdateBeeShopGoodsProp)              // 更新商品属性
	}
	{
		beeShopGoodsPropRouterWithoutRecord.GET("findBeeShopGoodsProp", beeShopGoodsPropApi.FindBeeShopGoodsProp)       // 根据ID获取商品属性
		beeShopGoodsPropRouterWithoutRecord.GET("getBeeShopGoodsPropList", beeShopGoodsPropApi.GetBeeShopGoodsPropList) // 获取商品属性列表
	}
	{
		beeShopGoodsPropRouterWithoutAuth.GET("getBeeShopGoodsPropPublic", beeShopGoodsPropApi.GetBeeShopGoodsPropPublic) // 获取商品属性列表
	}
}
