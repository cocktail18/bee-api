package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeCmsInfoRouter struct{}

// InitBeeCmsInfoRouter 初始化 cms信息 路由信息
func (s *BeeCmsInfoRouter) InitBeeCmsInfoRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeCmsInfoRouter := Router.Group("beeCmsInfo").Use(middleware.OperationRecord())
	beeCmsInfoRouterWithoutRecord := Router.Group("beeCmsInfo")
	beeCmsInfoRouterWithoutAuth := PublicRouter.Group("beeCmsInfo")

	var beeCmsInfoApi = api.BeeCmsInfoApi{}
	{
		beeCmsInfoRouter.POST("createBeeCmsInfo", beeCmsInfoApi.CreateBeeCmsInfo)             // 新建cms信息
		beeCmsInfoRouter.DELETE("deleteBeeCmsInfo", beeCmsInfoApi.DeleteBeeCmsInfo)           // 删除cms信息
		beeCmsInfoRouter.DELETE("deleteBeeCmsInfoByIds", beeCmsInfoApi.DeleteBeeCmsInfoByIds) // 批量删除cms信息
		beeCmsInfoRouter.PUT("updateBeeCmsInfo", beeCmsInfoApi.UpdateBeeCmsInfo)              // 更新cms信息
	}
	{
		beeCmsInfoRouterWithoutRecord.GET("findBeeCmsInfo", beeCmsInfoApi.FindBeeCmsInfo)       // 根据ID获取cms信息
		beeCmsInfoRouterWithoutRecord.GET("getBeeCmsInfoList", beeCmsInfoApi.GetBeeCmsInfoList) // 获取cms信息列表
	}
	{
		beeCmsInfoRouterWithoutAuth.GET("getBeeCmsInfoPublic", beeCmsInfoApi.GetBeeCmsInfoPublic) // 获取cms信息列表
	}
}
