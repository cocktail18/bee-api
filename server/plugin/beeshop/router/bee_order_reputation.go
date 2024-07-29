package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeOrderReputationRouter struct{}

// InitBeeOrderReputationRouter 初始化 商品评价 路由信息
func (s *BeeOrderReputationRouter) InitBeeOrderReputationRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeOrderReputationRouter := Router.Group("beeOrderReputation").Use(middleware.OperationRecord())
	beeOrderReputationRouterWithoutRecord := Router.Group("beeOrderReputation")
	beeOrderReputationRouterWithoutAuth := PublicRouter.Group("beeOrderReputation")

	var beeOrderReputationApi = api.BeeOrderReputationApi{}
	{
		beeOrderReputationRouter.POST("createBeeOrderReputation", beeOrderReputationApi.CreateBeeOrderReputation)             // 新建商品评价
		beeOrderReputationRouter.DELETE("deleteBeeOrderReputation", beeOrderReputationApi.DeleteBeeOrderReputation)           // 删除商品评价
		beeOrderReputationRouter.DELETE("deleteBeeOrderReputationByIds", beeOrderReputationApi.DeleteBeeOrderReputationByIds) // 批量删除商品评价
		beeOrderReputationRouter.PUT("updateBeeOrderReputation", beeOrderReputationApi.UpdateBeeOrderReputation)              // 更新商品评价
	}
	{
		beeOrderReputationRouterWithoutRecord.GET("findBeeOrderReputation", beeOrderReputationApi.FindBeeOrderReputation)       // 根据ID获取商品评价
		beeOrderReputationRouterWithoutRecord.GET("getBeeOrderReputationList", beeOrderReputationApi.GetBeeOrderReputationList) // 获取商品评价列表
	}
	{
		beeOrderReputationRouterWithoutAuth.GET("getBeeOrderReputationPublic", beeOrderReputationApi.GetBeeOrderReputationPublic) // 获取商品评价列表
	}
}
