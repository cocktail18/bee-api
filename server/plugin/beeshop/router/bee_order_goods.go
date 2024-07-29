package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeOrderGoodsRouter struct{}

// InitBeeOrderGoodsRouter 初始化 订单商品信息 路由信息
func (s *BeeOrderGoodsRouter) InitBeeOrderGoodsRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeOrderGoodsRouter := Router.Group("beeOrderGoods").Use(middleware.OperationRecord())
	beeOrderGoodsRouterWithoutRecord := Router.Group("beeOrderGoods")
	beeOrderGoodsRouterWithoutAuth := PublicRouter.Group("beeOrderGoods")

	var beeOrderGoodsApi = api.BeeOrderGoodsApi{}
	{
		beeOrderGoodsRouter.POST("createBeeOrderGoods", beeOrderGoodsApi.CreateBeeOrderGoods)             // 新建订单商品信息
		beeOrderGoodsRouter.DELETE("deleteBeeOrderGoods", beeOrderGoodsApi.DeleteBeeOrderGoods)           // 删除订单商品信息
		beeOrderGoodsRouter.DELETE("deleteBeeOrderGoodsByIds", beeOrderGoodsApi.DeleteBeeOrderGoodsByIds) // 批量删除订单商品信息
		beeOrderGoodsRouter.PUT("updateBeeOrderGoods", beeOrderGoodsApi.UpdateBeeOrderGoods)              // 更新订单商品信息
	}
	{
		beeOrderGoodsRouterWithoutRecord.GET("findBeeOrderGoods", beeOrderGoodsApi.FindBeeOrderGoods)       // 根据ID获取订单商品信息
		beeOrderGoodsRouterWithoutRecord.GET("getBeeOrderGoodsList", beeOrderGoodsApi.GetBeeOrderGoodsList) // 获取订单商品信息列表
	}
	{
		beeOrderGoodsRouterWithoutAuth.GET("getBeeOrderGoodsPublic", beeOrderGoodsApi.GetBeeOrderGoodsPublic) // 获取订单商品信息列表
	}
}
