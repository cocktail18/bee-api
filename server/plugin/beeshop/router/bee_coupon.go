package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeCouponRouter struct{}

// InitBeeCouponRouter 初始化 优惠券 路由信息
func (s *BeeCouponRouter) InitBeeCouponRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeCouponRouter := Router.Group("beeCoupon").Use(middleware.OperationRecord())
	beeCouponRouterWithoutRecord := Router.Group("beeCoupon")
	beeCouponRouterWithoutAuth := PublicRouter.Group("beeCoupon")

	var beeCouponApi = api.BeeCouponApi{}
	{
		beeCouponRouter.POST("createBeeCoupon", beeCouponApi.CreateBeeCoupon)             // 新建优惠券
		beeCouponRouter.DELETE("deleteBeeCoupon", beeCouponApi.DeleteBeeCoupon)           // 删除优惠券
		beeCouponRouter.DELETE("deleteBeeCouponByIds", beeCouponApi.DeleteBeeCouponByIds) // 批量删除优惠券
		beeCouponRouter.PUT("updateBeeCoupon", beeCouponApi.UpdateBeeCoupon)              // 更新优惠券
	}
	{
		beeCouponRouterWithoutRecord.GET("findBeeCoupon", beeCouponApi.FindBeeCoupon)       // 根据ID获取优惠券
		beeCouponRouterWithoutRecord.GET("getBeeCouponList", beeCouponApi.GetBeeCouponList) // 获取优惠券列表
	}
	{
		beeCouponRouterWithoutAuth.GET("getBeeCouponPublic", beeCouponApi.GetBeeCouponPublic) // 获取优惠券列表
	}
}
