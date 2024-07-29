package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeShopGoodsContentRouter struct{}

// InitBeeShopGoodsContentRouter 初始化 商品详情描述 路由信息
func (s *BeeShopGoodsContentRouter) InitBeeShopGoodsContentRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeShopGoodsContentRouter := Router.Group("beeShopGoodsContent").Use(middleware.OperationRecord())
	beeShopGoodsContentRouterWithoutRecord := Router.Group("beeShopGoodsContent")
	beeShopGoodsContentRouterWithoutAuth := PublicRouter.Group("beeShopGoodsContent")

	var beeShopGoodsContentApi = api.BeeShopGoodsContentApi{}
	{
		beeShopGoodsContentRouter.POST("createBeeShopGoodsContent", beeShopGoodsContentApi.CreateBeeShopGoodsContent)             // 新建商品详情描述
		beeShopGoodsContentRouter.DELETE("deleteBeeShopGoodsContent", beeShopGoodsContentApi.DeleteBeeShopGoodsContent)           // 删除商品详情描述
		beeShopGoodsContentRouter.DELETE("deleteBeeShopGoodsContentByIds", beeShopGoodsContentApi.DeleteBeeShopGoodsContentByIds) // 批量删除商品详情描述
		beeShopGoodsContentRouter.PUT("updateBeeShopGoodsContent", beeShopGoodsContentApi.UpdateBeeShopGoodsContent)              // 更新商品详情描述
	}
	{
		beeShopGoodsContentRouterWithoutRecord.GET("findBeeShopGoodsContent", beeShopGoodsContentApi.FindBeeShopGoodsContent)       // 根据ID获取商品详情描述
		beeShopGoodsContentRouterWithoutRecord.GET("getBeeShopGoodsContentList", beeShopGoodsContentApi.GetBeeShopGoodsContentList) // 获取商品详情描述列表
	}
	{
		beeShopGoodsContentRouterWithoutAuth.GET("getBeeShopGoodsContentPublic", beeShopGoodsContentApi.GetBeeShopGoodsContentPublic) // 获取商品详情描述列表
	}
}
