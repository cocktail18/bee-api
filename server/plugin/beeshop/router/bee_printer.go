package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/api"
	"github.com/gin-gonic/gin"
)

type BeePrinterRouter struct{}

// InitBeePrinterRouter 初始化 beePrinter表 路由信息
func (s *BeePrinterRouter) InitBeePrinterRouter(Router *gin.RouterGroup, PublicRouter *gin.RouterGroup) {
	beePrinterRouter := Router.Group("beePrinter").Use(middleware.OperationRecord())
	beePrinterRouterWithoutRecord := Router.Group("beePrinter")
	beePrinterRouterWithoutAuth := PublicRouter.Group("beePrinter")

	var beePrinterApi = api.BeePrinterApi{}
	{
		beePrinterRouter.POST("createBeePrinter", beePrinterApi.CreateBeePrinter)             // 新建beePrinter表
		beePrinterRouter.DELETE("deleteBeePrinter", beePrinterApi.DeleteBeePrinter)           // 删除beePrinter表
		beePrinterRouter.DELETE("deleteBeePrinterByIds", beePrinterApi.DeleteBeePrinterByIds) // 批量删除beePrinter表
		beePrinterRouter.PUT("updateBeePrinter", beePrinterApi.UpdateBeePrinter)              // 更新beePrinter表
	}
	{
		beePrinterRouterWithoutRecord.GET("findBeePrinter", beePrinterApi.FindBeePrinter)       // 根据ID获取beePrinter表
		beePrinterRouterWithoutRecord.GET("getBeePrinterList", beePrinterApi.GetBeePrinterList) // 获取beePrinter表列表
	}
	{
		beePrinterRouterWithoutAuth.GET("getBeePrinterPublic", beePrinterApi.GetBeePrinterPublic) // 获取beePrinter表列表
	}
}
