package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeUserRouter struct{}

// InitBeeUserRouter 初始化 beeUser表 路由信息
func (s *BeeUserRouter) InitBeeUserRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeUserRouter := Router.Group("beeUser").Use(middleware.OperationRecord())
	beeUserRouterWithoutRecord := Router.Group("beeUser")
	beeUserRouterWithoutAuth := PublicRouter.Group("beeUser")

	var beeUserApi = api.BeeUserApi{}
	{
		beeUserRouter.POST("createBeeUser", beeUserApi.CreateBeeUser)             // 新建beeUser表
		beeUserRouter.DELETE("deleteBeeUser", beeUserApi.DeleteBeeUser)           // 删除beeUser表
		beeUserRouter.DELETE("deleteBeeUserByIds", beeUserApi.DeleteBeeUserByIds) // 批量删除beeUser表
		beeUserRouter.PUT("updateBeeUser", beeUserApi.UpdateBeeUser)              // 更新beeUser表
	}
	{
		beeUserRouterWithoutRecord.GET("findBeeUser", beeUserApi.FindBeeUser)       // 根据ID获取beeUser表
		beeUserRouterWithoutRecord.GET("getBeeUserList", beeUserApi.GetBeeUserList) // 获取beeUser表列表
	}
	{
		beeUserRouterWithoutAuth.GET("getBeeUserPublic", beeUserApi.GetBeeUserPublic) // 获取beeUser表列表
	}
}
