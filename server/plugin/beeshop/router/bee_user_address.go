package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeUserAddressRouter struct{}

// InitBeeUserAddressRouter 初始化 用户地址 路由信息
func (s *BeeUserAddressRouter) InitBeeUserAddressRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeUserAddressRouter := Router.Group("beeUserAddress").Use(middleware.OperationRecord())
	beeUserAddressRouterWithoutRecord := Router.Group("beeUserAddress")
	beeUserAddressRouterWithoutAuth := PublicRouter.Group("beeUserAddress")

	var beeUserAddressApi = api.BeeUserAddressApi{}
	{
		beeUserAddressRouter.POST("createBeeUserAddress", beeUserAddressApi.CreateBeeUserAddress)             // 新建用户地址
		beeUserAddressRouter.DELETE("deleteBeeUserAddress", beeUserAddressApi.DeleteBeeUserAddress)           // 删除用户地址
		beeUserAddressRouter.DELETE("deleteBeeUserAddressByIds", beeUserAddressApi.DeleteBeeUserAddressByIds) // 批量删除用户地址
		beeUserAddressRouter.PUT("updateBeeUserAddress", beeUserAddressApi.UpdateBeeUserAddress)              // 更新用户地址
	}
	{
		beeUserAddressRouterWithoutRecord.GET("findBeeUserAddress", beeUserAddressApi.FindBeeUserAddress)       // 根据ID获取用户地址
		beeUserAddressRouterWithoutRecord.GET("getBeeUserAddressList", beeUserAddressApi.GetBeeUserAddressList) // 获取用户地址列表
	}
	{
		beeUserAddressRouterWithoutAuth.GET("getBeeUserAddressPublic", beeUserAddressApi.GetBeeUserAddressPublic) // 获取用户地址列表
	}
}
