package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeShopGoodsCategoryRouter struct{}

// InitBeeShopGoodsCategoryRouter 初始化 商品分类 路由信息
func (s *BeeShopGoodsCategoryRouter) InitBeeShopGoodsCategoryRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeShopGoodsCategoryRouter := Router.Group("beeShopGoodsCategory").Use(middleware.OperationRecord())
	beeShopGoodsCategoryRouterWithoutRecord := Router.Group("beeShopGoodsCategory")
	beeShopGoodsCategoryRouterWithoutAuth := PublicRouter.Group("beeShopGoodsCategory")

	var beeShopGoodsCategoryApi = api.BeeShopGoodsCategoryApi{}
	{
		beeShopGoodsCategoryRouter.POST("createBeeShopGoodsCategory", beeShopGoodsCategoryApi.CreateBeeShopGoodsCategory)             // 新建商品分类
		beeShopGoodsCategoryRouter.DELETE("deleteBeeShopGoodsCategory", beeShopGoodsCategoryApi.DeleteBeeShopGoodsCategory)           // 删除商品分类
		beeShopGoodsCategoryRouter.DELETE("deleteBeeShopGoodsCategoryByIds", beeShopGoodsCategoryApi.DeleteBeeShopGoodsCategoryByIds) // 批量删除商品分类
		beeShopGoodsCategoryRouter.PUT("updateBeeShopGoodsCategory", beeShopGoodsCategoryApi.UpdateBeeShopGoodsCategory)              // 更新商品分类
	}
	{
		beeShopGoodsCategoryRouterWithoutRecord.GET("findBeeShopGoodsCategory", beeShopGoodsCategoryApi.FindBeeShopGoodsCategory)       // 根据ID获取商品分类
		beeShopGoodsCategoryRouterWithoutRecord.GET("getBeeShopGoodsCategoryList", beeShopGoodsCategoryApi.GetBeeShopGoodsCategoryList) // 获取商品分类列表
	}
	{
		beeShopGoodsCategoryRouterWithoutAuth.GET("getBeeShopGoodsCategoryPublic", beeShopGoodsCategoryApi.GetBeeShopGoodsCategoryPublic) // 获取商品分类列表
	}
}
