package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeShopGoodsImagesRouter struct{}

// InitBeeShopGoodsImagesRouter 初始化 商品图 路由信息
func (s *BeeShopGoodsImagesRouter) InitBeeShopGoodsImagesRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeShopGoodsImagesRouter := Router.Group("beeShopGoodsImages").Use(middleware.OperationRecord())
	beeShopGoodsImagesRouterWithoutRecord := Router.Group("beeShopGoodsImages")
	beeShopGoodsImagesRouterWithoutAuth := PublicRouter.Group("beeShopGoodsImages")

	var beeShopGoodsImagesApi = api.BeeShopGoodsImagesApi{}
	{
		beeShopGoodsImagesRouter.POST("createBeeShopGoodsImages", beeShopGoodsImagesApi.CreateBeeShopGoodsImages)             // 新建商品图
		beeShopGoodsImagesRouter.DELETE("deleteBeeShopGoodsImages", beeShopGoodsImagesApi.DeleteBeeShopGoodsImages)           // 删除商品图
		beeShopGoodsImagesRouter.DELETE("deleteBeeShopGoodsImagesByIds", beeShopGoodsImagesApi.DeleteBeeShopGoodsImagesByIds) // 批量删除商品图
		beeShopGoodsImagesRouter.PUT("updateBeeShopGoodsImages", beeShopGoodsImagesApi.UpdateBeeShopGoodsImages)              // 更新商品图
	}
	{
		beeShopGoodsImagesRouterWithoutRecord.GET("findBeeShopGoodsImages", beeShopGoodsImagesApi.FindBeeShopGoodsImages)       // 根据ID获取商品图
		beeShopGoodsImagesRouterWithoutRecord.GET("getBeeShopGoodsImagesList", beeShopGoodsImagesApi.GetBeeShopGoodsImagesList) // 获取商品图列表
	}
	{
		beeShopGoodsImagesRouterWithoutAuth.GET("getBeeShopGoodsImagesPublic", beeShopGoodsImagesApi.GetBeeShopGoodsImagesPublic) // 获取商品图列表
	}
}
