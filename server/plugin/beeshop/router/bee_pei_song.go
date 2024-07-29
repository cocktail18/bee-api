package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeePeiSongRouter struct{}

// InitBeePeiSongRouter 初始化 配送信息 路由信息
func (s *BeePeiSongRouter) InitBeePeiSongRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beePeiSongRouter := Router.Group("beePeiSong").Use(middleware.OperationRecord())
	beePeiSongRouterWithoutRecord := Router.Group("beePeiSong")
	beePeiSongRouterWithoutAuth := PublicRouter.Group("beePeiSong")

	var beePeiSongApi = api.BeePeiSongApi{}
	{
		beePeiSongRouter.POST("createBeePeiSong", beePeiSongApi.CreateBeePeiSong)             // 新建配送信息
		beePeiSongRouter.DELETE("deleteBeePeiSong", beePeiSongApi.DeleteBeePeiSong)           // 删除配送信息
		beePeiSongRouter.DELETE("deleteBeePeiSongByIds", beePeiSongApi.DeleteBeePeiSongByIds) // 批量删除配送信息
		beePeiSongRouter.PUT("updateBeePeiSong", beePeiSongApi.UpdateBeePeiSong)              // 更新配送信息
	}
	{
		beePeiSongRouterWithoutRecord.GET("findBeePeiSong", beePeiSongApi.FindBeePeiSong)       // 根据ID获取配送信息
		beePeiSongRouterWithoutRecord.GET("getBeePeiSongList", beePeiSongApi.GetBeePeiSongList) // 获取配送信息列表
	}
	{
		beePeiSongRouterWithoutAuth.GET("getBeePeiSongPublic", beePeiSongApi.GetBeePeiSongPublic) // 获取配送信息列表
	}
}
