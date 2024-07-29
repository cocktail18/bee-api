package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeShopGoodsRouter struct{}

// InitBeeShopGoodsRouter 初始化 商品表 路由信息
func (s *BeeShopGoodsRouter) InitBeeShopGoodsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeShopGoodsRouter := Router.Group("beeShopGoods").Use(middleware.OperationRecord())
	beeShopGoodsRouterWithoutRecord := Router.Group("beeShopGoods")
	beeShopGoodsRouterWithoutAuth := PublicRouter.Group("beeShopGoods")

	var beeShopGoodsApi = api.BeeShopGoodsApi{}
	{
		beeShopGoodsRouter.POST("createBeeShopGoods", beeShopGoodsApi.CreateBeeShopGoods)             // 新建商品表
		beeShopGoodsRouter.DELETE("deleteBeeShopGoods", beeShopGoodsApi.DeleteBeeShopGoods)           // 删除商品表
		beeShopGoodsRouter.DELETE("deleteBeeShopGoodsByIds", beeShopGoodsApi.DeleteBeeShopGoodsByIds) // 批量删除商品表
		beeShopGoodsRouter.PUT("updateBeeShopGoods", beeShopGoodsApi.UpdateBeeShopGoods)              // 更新商品表
	}
	{
		beeShopGoodsRouterWithoutRecord.GET("findBeeShopGoods", beeShopGoodsApi.FindBeeShopGoods)       // 根据ID获取商品表
		beeShopGoodsRouterWithoutRecord.GET("getBeeShopGoodsList", beeShopGoodsApi.GetBeeShopGoodsList) // 获取商品表列表
	}
	{
		beeShopGoodsRouterWithoutAuth.GET("getBeeShopGoodsPublic", beeShopGoodsApi.GetBeeShopGoodsPublic) // 获取商品表列表
	}
}
