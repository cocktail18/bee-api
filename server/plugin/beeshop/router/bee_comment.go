package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeeCommentRouter struct{}

// InitBeeCommentRouter 初始化 评论表 路由信息
func (s *BeeCommentRouter) InitBeeCommentRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beeCommentRouter := Router.Group("beeComment").Use(middleware.OperationRecord())
	beeCommentRouterWithoutRecord := Router.Group("beeComment")
	beeCommentRouterWithoutAuth := PublicRouter.Group("beeComment")

	var beeCommentApi = api.BeeCommentApi{}
	{
		beeCommentRouter.POST("createBeeComment", beeCommentApi.CreateBeeComment)             // 新建评论表
		beeCommentRouter.DELETE("deleteBeeComment", beeCommentApi.DeleteBeeComment)           // 删除评论表
		beeCommentRouter.DELETE("deleteBeeCommentByIds", beeCommentApi.DeleteBeeCommentByIds) // 批量删除评论表
		beeCommentRouter.PUT("updateBeeComment", beeCommentApi.UpdateBeeComment)              // 更新评论表
	}
	{
		beeCommentRouterWithoutRecord.GET("findBeeComment", beeCommentApi.FindBeeComment)       // 根据ID获取评论表
		beeCommentRouterWithoutRecord.GET("getBeeCommentList", beeCommentApi.GetBeeCommentList) // 获取评论表列表
	}
	{
		beeCommentRouterWithoutAuth.GET("getBeeCommentPublic", beeCommentApi.GetBeeCommentPublic) // 获取评论表列表
	}
}
