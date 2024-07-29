package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeShopGoodsSkuRouter struct{}

// InitBeeShopGoodsSkuRouter 初始化 商品sku 路由信息
func (s *BeeShopGoodsSkuRouter) InitBeeShopGoodsSkuRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeShopGoodsSkuRouter := Router.Group("beeShopGoodsSku").Use(middleware.OperationRecord())
	beeShopGoodsSkuRouterWithoutRecord := Router.Group("beeShopGoodsSku")
	beeShopGoodsSkuRouterWithoutAuth := PublicRouter.Group("beeShopGoodsSku")

	var beeShopGoodsSkuApi = api.BeeShopGoodsSkuApi{}
	{
		beeShopGoodsSkuRouter.POST("createBeeShopGoodsSku", beeShopGoodsSkuApi.CreateBeeShopGoodsSku)             // 新建商品sku
		beeShopGoodsSkuRouter.DELETE("deleteBeeShopGoodsSku", beeShopGoodsSkuApi.DeleteBeeShopGoodsSku)           // 删除商品sku
		beeShopGoodsSkuRouter.DELETE("deleteBeeShopGoodsSkuByIds", beeShopGoodsSkuApi.DeleteBeeShopGoodsSkuByIds) // 批量删除商品sku
		beeShopGoodsSkuRouter.PUT("updateBeeShopGoodsSku", beeShopGoodsSkuApi.UpdateBeeShopGoodsSku)              // 更新商品sku
	}
	{
		beeShopGoodsSkuRouterWithoutRecord.GET("findBeeShopGoodsSku", beeShopGoodsSkuApi.FindBeeShopGoodsSku)       // 根据ID获取商品sku
		beeShopGoodsSkuRouterWithoutRecord.GET("getBeeShopGoodsSkuList", beeShopGoodsSkuApi.GetBeeShopGoodsSkuList) // 获取商品sku列表
	}
	{
		beeShopGoodsSkuRouterWithoutAuth.GET("getBeeShopGoodsSkuPublic", beeShopGoodsSkuApi.GetBeeShopGoodsSkuPublic) // 获取商品sku列表
	}
}
