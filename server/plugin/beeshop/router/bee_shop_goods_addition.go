package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeShopGoodsAdditionRouter struct{}

// InitBeeShopGoodsAdditionRouter 初始化 商品额外信息 路由信息
func (s *BeeShopGoodsAdditionRouter) InitBeeShopGoodsAdditionRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeShopGoodsAdditionRouter := Router.Group("beeShopGoodsAddition").Use(middleware.OperationRecord())
	beeShopGoodsAdditionRouterWithoutRecord := Router.Group("beeShopGoodsAddition")
	beeShopGoodsAdditionRouterWithoutAuth := PublicRouter.Group("beeShopGoodsAddition")

	var beeShopGoodsAdditionApi = api.BeeShopGoodsAdditionApi{}
	{
		beeShopGoodsAdditionRouter.POST("createBeeShopGoodsAddition", beeShopGoodsAdditionApi.CreateBeeShopGoodsAddition)             // 新建商品额外信息
		beeShopGoodsAdditionRouter.DELETE("deleteBeeShopGoodsAddition", beeShopGoodsAdditionApi.DeleteBeeShopGoodsAddition)           // 删除商品额外信息
		beeShopGoodsAdditionRouter.DELETE("deleteBeeShopGoodsAdditionByIds", beeShopGoodsAdditionApi.DeleteBeeShopGoodsAdditionByIds) // 批量删除商品额外信息
		beeShopGoodsAdditionRouter.PUT("updateBeeShopGoodsAddition", beeShopGoodsAdditionApi.UpdateBeeShopGoodsAddition)              // 更新商品额外信息
	}
	{
		beeShopGoodsAdditionRouterWithoutRecord.GET("findBeeShopGoodsAddition", beeShopGoodsAdditionApi.FindBeeShopGoodsAddition)       // 根据ID获取商品额外信息
		beeShopGoodsAdditionRouterWithoutRecord.GET("getBeeShopGoodsAdditionList", beeShopGoodsAdditionApi.GetBeeShopGoodsAdditionList) // 获取商品额外信息列表
	}
	{
		beeShopGoodsAdditionRouterWithoutAuth.GET("getBeeShopGoodsAdditionPublic", beeShopGoodsAdditionApi.GetBeeShopGoodsAdditionPublic) // 获取商品额外信息列表
	}
}
