package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeUserMapperRouter struct{}

// InitBeeUserMapperRouter 初始化 beeUserMapper表 路由信息
func (s *BeeUserMapperRouter) InitBeeUserMapperRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeUserMapperRouter := Router.Group("beeUserMapper").Use(middleware.OperationRecord())
	beeUserMapperRouterWithoutRecord := Router.Group("beeUserMapper")
	beeUserMapperRouterWithoutAuth := PublicRouter.Group("beeUserMapper")

	var beeUserMapperApi = api.BeeUserMapperApi{}
	{
		beeUserMapperRouter.POST("createBeeUserMapper", beeUserMapperApi.CreateBeeUserMapper)             // 新建beeUserMapper表
		beeUserMapperRouter.DELETE("deleteBeeUserMapper", beeUserMapperApi.DeleteBeeUserMapper)           // 删除beeUserMapper表
		beeUserMapperRouter.DELETE("deleteBeeUserMapperByIds", beeUserMapperApi.DeleteBeeUserMapperByIds) // 批量删除beeUserMapper表
		beeUserMapperRouter.PUT("updateBeeUserMapper", beeUserMapperApi.UpdateBeeUserMapper)              // 更新beeUserMapper表
	}
	{
		beeUserMapperRouterWithoutRecord.GET("findBeeUserMapper", beeUserMapperApi.FindBeeUserMapper)       // 根据ID获取beeUserMapper表
		beeUserMapperRouterWithoutRecord.GET("getBeeUserMapperList", beeUserMapperApi.GetBeeUserMapperList) // 获取beeUserMapper表列表
	}
	{
		beeUserMapperRouterWithoutAuth.GET("getBeeUserMapperPublic", beeUserMapperApi.GetBeeUserMapperPublic) // 获取beeUserMapper表列表
	}
}
