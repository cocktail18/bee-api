package beeshop

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"

	"gitee.com/stuinfer/bee-api/cmd"
	"gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/db"
	"gitee.com/stuinfer/bee-api/logger"
	"gitee.com/stuinfer/bee-api/util"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/router"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BeeShopPlugin struct {
}

//go:embed menus.json
var menusJsonStr string

//go:embed apis.json
var apisJsonStr string

type SysBaseMenu struct {
	ID          uint   `json:"id"`
	MenuLevel   uint   `json:"menu_level"`
	ParentId    uint   `json:"parent_id" gorm:"comment:父菜单ID"`    // 父菜单ID
	Path        string `json:"path" gorm:"comment:路由path"`        // 路由path
	Name        string `json:"name" gorm:"comment:路由name"`        // 路由name
	Hidden      bool   `json:"hidden" gorm:"comment:是否在列表隐藏"`     // 是否在列表隐藏
	Component   string `json:"component" gorm:"comment:对应前端文件路径"` // 对应前端文件路径
	Sort        int    `json:"sort" gorm:"comment:排序标记"`          // 排序标记
	ActiveName  string `json:"active_name" gorm:"comment:高亮菜单"`
	KeepAlive   bool   `json:"keep_alive" gorm:"comment:是否缓存"`           // 是否缓存
	DefaultMenu bool   `json:"default_menu" gorm:"comment:是否是基础路由（开发中）"` // 是否是基础路由（开发中）
	Title       string `json:"title" gorm:"comment:菜单名"`                 // 菜单名
	Icon        string `json:"icon" gorm:"comment:菜单图标"`                 // 菜单图标
	CloseTab    bool   `json:"close_tab" gorm:"comment:自动关闭tab"`         // 自动关闭tab
	DeletedAt   string `json:"deleted_at"`
}

func CreateBeeShopPlug() *BeeShopPlugin {
	ins := &BeeShopPlugin{}
	if global.GVA_CONFIG.BeeShop.Disable || global.GVA_DB == nil { //未初始化
		return &BeeShopPlugin{}
	}
	menus := make([]SysBaseMenu, 0)
	_ = json.Unmarshal([]byte(menusJsonStr), &menus)
	sort.Slice(menus, func(i, j int) bool {
		return menus[i].ParentId < menus[j].ParentId
	})
	menuIdMapper := make(map[uint]uint) // 旧的到新的映射
	for _, menu := range menus {
		if menu.DeletedAt != "" {
			continue
		}
		cnt := int64(0)
		global.GVA_DB.Find(&[]system.SysBaseMenu{}, "name in (?)", menu.Name).Count(&cnt)
		if cnt > 0 {
			global.GVA_LOG.Info("存在同名menu，跳过")
			continue
		}
		menuPath := menu.Path
		component := menu.Component
		if component != "view/routerHolder.vue" {
			component = strings.ReplaceAll(component, "view/bee/", "plugin/beeshop/view/")
		}
		parentId := uint(0)
		if menu.ParentId != 0 {
			parentId = menuIdMapper[menu.ParentId]
		}
		item := system.SysBaseMenu{
			MenuLevel: menu.MenuLevel,
			ParentId:  parentId,
			Path:      menuPath,
			Name:      menu.Name,
			Hidden:    menu.Hidden,
			Component: component,
			Sort:      menu.Sort,
			Meta: system.Meta{
				ActiveName:  menu.ActiveName,
				KeepAlive:   menu.KeepAlive,
				DefaultMenu: menu.DefaultMenu,
				Title:       menu.Title,
				Icon:        menu.Icon,
				CloseTab:    menu.CloseTab,
			},
		}
		err := global.GVA_DB.Create(&item).Error
		if err != nil {
			fmt.Println(err)
		}
		menuIdMapper[menu.ID] = item.ID
	}

	// 下方会自动注册菜单 第一个参数为菜单一级路由信息一般为定义好的组名 第二个参数为真实使用的web页面路由信息
	// 具体值请根据实际情况修改
	//utils.RegisterMenus()

	sysApis := make([]system.SysApi, 0)
	_ = json.Unmarshal([]byte(apisJsonStr), &sysApis)
	apiAdd := make([]system.SysApi, 0)
	for _, api := range sysApis {
		if api.DeletedAt.Valid {
			continue
		}
		apiAdd = append(apiAdd, system.SysApi{
			Path:        "/bee-shop" + api.Path,
			Description: api.Description,
			ApiGroup:    "bee-shop",
			Method:      api.Method,
		})
	}
	// 下方会自动注册api 以下格式为示例格式，请按照实际情况修改
	utils.RegisterApis(
		apiAdd...,
	)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/bee-shop/beePayLog/getBeePayTotal",
			Description: "支付总额",
			ApiGroup:    "bee-shop",
			Method:      "GET",
		},
	)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/bee-shop/beeOrder/updateBeeOrderExtJsonStr",
			Description: "更新订单额外json信息",
			ApiGroup:    "bee-shop",
			Method:      "PUT",
		},
	)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/bee-shop/beeOrder/updateBeeOrderStatus",
			Description: "更新订单status字段",
			ApiGroup:    "bee-shop",
			Method:      "PUT",
		},
	)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/bee-shop/beeOrder/markBeeOrderPaid",
			Description: "标识为已支付",
			ApiGroup:    "bee-shop",
			Method:      "PUT",
		},
	)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/bee-shop/beeCyTable/getBeeCyTableQrCode",
			Description: "获取桌子二维码",
			ApiGroup:    "bee-shop",
			Method:      "GET",
		},
	)
	utils.RegisterApis(ins.genPluginApi("beeRechargeSendRule", "充值优惠规则")...)
	utils.RegisterApis(ins.genPluginApi("beePrinter", "打印机")...)
	utils.RegisterApis(ins.genPluginApi("beeQueue", "排队叫号")...)
	utils.RegisterApis(ins.genPluginApi("beeUserQueue", "取号列表")...)
	utils.RegisterApis(ins.genPluginApi("beeOrderPeisong", "订单配送")...)
	utils.RegisterApis(ins.genPluginApi("beeOrderPeisongLog", "订单配送日志")...)
	utils.RegisterApis(ins.genPluginApi("beeDelivery", "配送供应商")...)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/bee-shop/beeOrderPeisong/cancelBeeOrderPeisong",
			Description: "取消配送",
			ApiGroup:    "bee-shop",
			Method:      "POST",
		},
	)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/bee-shop/beeOrderPeisong/notifyBeeOrderPeisong",
			Description: "通知配送",
			ApiGroup:    "bee-shop",
			Method:      "POST",
		},
	)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/bee-shop/beeOrderPeisong/getBeeOrderPeisongDetail",
			Description: "获取配送信息详情",
			ApiGroup:    "bee-shop",
			Method:      "GET",
		},
	)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/bee-shop/beeDelivery/bindYunlabaShop",
			Description: "绑定云喇叭",
			ApiGroup:    "bee-shop",
			Method:      "PUT",
		},
	)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/bee-shop/beeOrder/shippedBeeOrder",
			Description: "确认已发货",
			ApiGroup:    "bee-shop",
			Method:      "PUT",
		},
	)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/bee-shop/beePrinter/testBeePrinter",
			Description: "打印机测试",
			ApiGroup:    "bee-shop",
			Method:      "POST",
		},
	)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/bee-shop/beeQueue/passBeeQueue",
			Description: "过号",
			ApiGroup:    "bee-shop",
			Method:      "PUT",
		},
	)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/bee-shop/beeQueue/callBeeQueue",
			Description: "叫号",
			ApiGroup:    "bee-shop",
			Method:      "PUT",
		},
	)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/bee-shop/beeQueue/reCallBeeQueue",
			Description: "重新叫号",
			ApiGroup:    "bee-shop",
			Method:      "PUT",
		},
	)
	utils.RegisterApis(
		system.SysApi{
			Path:        "/bee-shop/beeQueue/nextBeeQueue",
			Description: "叫号下一位",
			ApiGroup:    "bee-shop",
			Method:      "PUT",
		},
	)
	ins.registerBaseMenu("shop-order-admin", system.SysBaseMenu{
		Path: "beeOrderPeisong",
		Name: "beeOrderPeisong",
		Meta: system.Meta{
			Title: "订单配送信息",
		},
	})

	ins.registerBaseMenu("beeFinancialManager", system.SysBaseMenu{
		Path: "beeRechargeSendRule",
		Name: "beeRechargeSendRule",
		Meta: system.Meta{
			Title: "充值优惠规则配置",
		},
	})

	ins.registerBaseMenu("shop-order-admin", system.SysBaseMenu{
		Path: "beeOrderPeisongLog",
		Name: "beeOrderPeisongLog",
		Meta: system.Meta{
			Title: "订单配送信息日志",
		},
	})
	ins.registerBaseMenu("shop-base-info", system.SysBaseMenu{
		Path: "beePrinter",
		Name: "beePrinter",
		Meta: system.Meta{
			Title: "打印机配置",
		},
	})
	ins.registerBaseMenu("bee_logistics_admin", system.SysBaseMenu{
		Path: "beeDelivery",
		Name: "beeDelivery",
		Meta: system.Meta{
			Title: "配送供应商配置",
		},
	})
	ins.registerBaseMenu("bee_logistics_admin", system.SysBaseMenu{
		Path:      "beeDeliveryBindYunlaba",
		Name:      "beeDeliveryBindYunlaba",
		Component: "plugin/beeshop/view/beeDelivery/beeDeliveryYunlabaBind.vue",
		Meta: system.Meta{
			Title: "绑定云喇叭账户",
		},
	})
	ins.registerBaseMenu("bee_index", system.SysBaseMenu{
		Path: "beeDashboard",
		Name: "beeDashboard",
		Meta: system.Meta{
			Title: "商城大盘",
			Icon:  "data-analysis",
		},
	})
	ins.registerBaseMenu("bee_index", system.SysBaseMenu{
		Path:      "beeQueueAdmin",
		Name:      "beeQueueAdmin",
		Component: "view/routerHolder.vue",
		Meta: system.Meta{
			Title: "排队叫号",
			Icon:  "goblet-square",
		},
	})
	ins.registerBaseMenu("beeQueueAdmin", system.SysBaseMenu{
		Path: "beeQueue",
		Name: "beeQueue",
		Meta: system.Meta{
			Title: "队列维护",
		},
	})
	ins.registerBaseMenu("beeQueueAdmin", system.SysBaseMenu{
		Path: "beeUserQueue",
		Name: "beeUserQueue",
		Meta: system.Meta{
			Title: "取号列表",
		},
	})
	ins.registerBaseMenu("bee_index", system.SysBaseMenu{
		Path:      "beeOrderTodo",
		Name:      "beeOrderTodo",
		Component: "plugin/beeshop/view/beeOrder/beeOrderTodo.vue",
		Meta: system.Meta{
			Title: "待发货订单",
			Icon:  "histogram",
		},
	})
	// 下方会自动注册api 以下格式为示例格式，请按照实际情况修改
	//utils.RegisterApis(
	//	system.SysApi{
	//		Path:        "/bee-shop/routerName",
	//		Description: "BeeShop接口",
	//		ApiGroup:    "bee-shop",
	//		Method:      "POST",
	//	},
	//)
	initBeeDict()

	go func() {
		ins.startBeeApi()
	}()
	return ins
}

func (*BeeShopPlugin) registerBaseMenu(parentMenu string, menu system.SysBaseMenu) {
	var cnt int64
	global.GVA_DB.Find(&[]system.SysBaseMenu{}, "name in (?)", menu.Name).Count(&cnt)
	if cnt > 0 {
		global.GVA_LOG.Info("存在同名menu，跳过", zap.String("name", menu.Name))
		return
	}
	var parent system.SysBaseMenu
	if err := global.GVA_DB.Where("name = ?", parentMenu).First(&parent).Error; err != nil {
		global.GVA_LOG.Error("获取父菜单失败", zap.Error(err), zap.String("name", parentMenu))
		return
	}
	menu.CreatedAt = time.Now()
	menu.ParentId = parent.ID
	menu.Hidden = false
	if menu.Component == "" {
		menu.Component = fmt.Sprintf("plugin/beeshop/view/%s/%s.vue", menu.Path, menu.Path)
	}
	if err := global.GVA_DB.Create(&menu).Error; err != nil {
		global.GVA_LOG.Error("生成菜单失败", zap.Error(err), zap.String("name", menu.Name))
		return
	}
}

func (*BeeShopPlugin) genPluginApi(tableName string, remark string) []system.SysApi {
	tpls := []system.SysApi{
		{ApiGroup: "bee-shop", Path: "/bee-shop/{{tbName}}/create{{tbNameFirstLetterUpper}}", Method: "POST", Description: "新增{{remark}}"},
		{ApiGroup: "bee-shop", Path: "/bee-shop/{{tbName}}/delete{{tbNameFirstLetterUpper}}", Method: "DELETE", Description: "删除{{remark}}"},
		{ApiGroup: "bee-shop", Path: "/bee-shop/{{tbName}}/delete{{tbNameFirstLetterUpper}}ByIds", Method: "DELETE", Description: "批量删除{{remark}}"},
		{ApiGroup: "bee-shop", Path: "/bee-shop/{{tbName}}/update{{tbNameFirstLetterUpper}}", Method: "PUT", Description: "更新{{remark}}"},
		{ApiGroup: "bee-shop", Path: "/bee-shop/{{tbName}}/find{{tbNameFirstLetterUpper}}", Method: "GET", Description: "根据ID获取{{remark}}"},
		{ApiGroup: "bee-shop", Path: "/bee-shop/{{tbName}}/get{{tbNameFirstLetterUpper}}List", Method: "GET", Description: "获取{{remark}}列表"},
	}
	for i, tpl := range tpls {
		tpls[i].Description = strings.ReplaceAll(tpl.Description, "{{remark}}", remark)
		tpls[i].Path = strings.ReplaceAll(tpls[i].Path, "{{tbName}}", tableName)
		tpls[i].Path = strings.ReplaceAll(tpls[i].Path, "{{tbNameFirstLetterUpper}}", strings.ToUpper(string(tableName[0]))+tableName[1:])
	}
	return tpls
}

func (*BeeShopPlugin) startBeeApi() {
	beeDb := global.GetGlobalDBByDBName("bee")
	if beeDb == nil { //直接用同一个库
		db.SetDB(global.GVA_DB)
	} else {
		db.SetDB(beeDb)
	}

	lockFile := "./bee-init.lock"
	b, err := util.FileExists(lockFile)
	if err != nil {
		panic(err)
	}
	if !b {
		if err = ioutil.WriteFile(lockFile, []byte("初始化完成后请不要删除该文件"), 0644); err != nil {
			logger.GetLogger().Error("写入初始化文件失败", zap.Error(err))
			panic(err.Error())
		}
	}
	listen := global.GVA_CONFIG.BeeShop.Listen
	if listen == "" {
		listen = "127.0.0.1:18083"
	}
	appCfg := &config.AppConfig{
		App: &config.App{
			Listen:  listen,
			DfsHost: util.IF(global.GVA_CONFIG.BeeShop.Host == "", "http://127.0.0.1:18083", global.GVA_CONFIG.BeeShop.Host),
			Host:    util.IF(global.GVA_CONFIG.BeeShop.Host == "", "http://127.0.0.1:18083", global.GVA_CONFIG.BeeShop.Host),
		},
		DB: &config.AppDBConfig{
			Drop:     false,
			NeedInit: false,
		},
		Default: &config.DefaultConfig{
			Wx: &config.WxConfig{
				AppId:  "wxa4bc935c9ba87528",               //改成你的
				Secret: "368ce6cfeeee84764a694ff1cd2ef9ea", //改成你的
			},
			SysUser: &config.SysUser{
				Domain:   "9kuai8coffee.asia",
				Username: "admin",
				Password: "123456",
			},
		},
		Upload: &config.UploadConfig{
			OssType: global.GVA_CONFIG.System.OssType,
			Local: config.Local{
				Path:      global.GVA_CONFIG.Local.Path,
				StorePath: global.GVA_CONFIG.Local.StorePath,
			},
			Qiniu: config.Qiniu{
				Zone:          global.GVA_CONFIG.Qiniu.Zone,
				AccessKey:     global.GVA_CONFIG.Qiniu.AccessKey,
				SecretKey:     global.GVA_CONFIG.Qiniu.SecretKey,
				Bucket:        global.GVA_CONFIG.Qiniu.Bucket,
				ImgPath:       global.GVA_CONFIG.Qiniu.ImgPath,
				UseHTTPS:      global.GVA_CONFIG.Qiniu.UseHTTPS,
				UseCdnDomains: global.GVA_CONFIG.Qiniu.UseCdnDomains,
			},
			AliyunOSS: config.AliyunOSS{
				Endpoint:        global.GVA_CONFIG.AliyunOSS.Endpoint,
				AccessKeyId:     global.GVA_CONFIG.AliyunOSS.AccessKeyId,
				AccessKeySecret: global.GVA_CONFIG.AliyunOSS.AccessKeySecret,
				BucketName:      global.GVA_CONFIG.AliyunOSS.BucketName,
				BucketUrl:       global.GVA_CONFIG.AliyunOSS.BucketUrl,
				BasePath:        global.GVA_CONFIG.AliyunOSS.BasePath,
			},
			HuaWeiObs: config.HuaWeiObs{
				Path:      global.GVA_CONFIG.HuaWeiObs.Path,
				Bucket:    global.GVA_CONFIG.HuaWeiObs.Bucket,
				Endpoint:  global.GVA_CONFIG.HuaWeiObs.Endpoint,
				AccessKey: global.GVA_CONFIG.HuaWeiObs.AccessKey,
				SecretKey: global.GVA_CONFIG.HuaWeiObs.SecretKey,
			},
			TencentCOS: config.TencentCOS{
				Bucket:     global.GVA_CONFIG.TencentCOS.Bucket,
				Region:     global.GVA_CONFIG.TencentCOS.Region,
				SecretID:   global.GVA_CONFIG.TencentCOS.SecretID,
				SecretKey:  global.GVA_CONFIG.TencentCOS.SecretKey,
				BaseURL:    global.GVA_CONFIG.TencentCOS.BaseURL,
				PathPrefix: global.GVA_CONFIG.TencentCOS.PathPrefix,
			},
			AwsS3: config.AwsS3{
				Bucket:           global.GVA_CONFIG.AwsS3.Bucket,
				Region:           global.GVA_CONFIG.AwsS3.Region,
				Endpoint:         global.GVA_CONFIG.AwsS3.Endpoint,
				SecretID:         global.GVA_CONFIG.AwsS3.SecretID,
				SecretKey:        global.GVA_CONFIG.AwsS3.SecretKey,
				BaseURL:          global.GVA_CONFIG.AwsS3.BaseURL,
				PathPrefix:       global.GVA_CONFIG.AwsS3.PathPrefix,
				S3ForcePathStyle: global.GVA_CONFIG.AwsS3.S3ForcePathStyle,
				DisableSSL:       global.GVA_CONFIG.AwsS3.DisableSSL,
			},
		},
	}
	logger.SetLogger(global.GVA_LOG)
	cmd.Run(appCfg)

}

func (*BeeShopPlugin) Register(group *gin.RouterGroup) {
	beeRouter := router.RouterGroupApp

	privateGroup := group.Group("")
	publicGroup := group.Group("")

	privateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())

	beeRouter.InitBeeUserMapperRouter(privateGroup, publicGroup)
	beeRouter.InitBeeShopGoodsCategoryRouter(privateGroup, publicGroup)
	beeRouter.InitBeeShopInfoRouter(privateGroup, publicGroup)
	beeRouter.InitBeeUserBalanceLogRouter(privateGroup, publicGroup)
	beeRouter.InitBeeCouponRouter(privateGroup, publicGroup)
	beeRouter.InitBeeUserCouponRouter(privateGroup, publicGroup)
	beeRouter.InitBeeCashLogRouter(privateGroup, publicGroup)
	beeRouter.InitBeeConfigRouter(privateGroup, publicGroup)
	beeRouter.InitBeeShopGoodsPropRouter(privateGroup, publicGroup)
	beeRouter.InitBeeShopGoodsAdditionRouter(privateGroup, publicGroup)
	beeRouter.InitBeeShopGoodsRouter(privateGroup, publicGroup)
	beeRouter.InitBeeShopGoodsImagesRouter(privateGroup, publicGroup)

	beeRouter.InitBeeShopGoodsSkuRouter(privateGroup, publicGroup)

	beeRouter.InitBeeShopGoodsContentRouter(privateGroup, publicGroup)
	beeRouter.InitBeeWxConfigRouter(privateGroup, publicGroup)

	beeRouter.InitBeeCommentRouter(privateGroup, publicGroup)
	beeRouter.InitBeeOrderRouter(privateGroup, publicGroup)
	beeRouter.InitBeeOrderReputationRouter(privateGroup, publicGroup)
	beeRouter.InitBeeBannerRouter(privateGroup, publicGroup)
	beeRouter.InitBeeUserAmountRouter(privateGroup, publicGroup)
	beeRouter.InitBeeCyTableRouter(privateGroup, publicGroup)

	beeRouter.InitBeeRegionRouter(privateGroup, publicGroup)
	beeRouter.InitBeeCmsInfoRouter(privateGroup, publicGroup)
	beeRouter.InitBeeUserAddressRouter(privateGroup, publicGroup)

	beeRouter.InitBeePeiSongRouter(privateGroup, publicGroup)
	beeRouter.InitBeeUserRouter(privateGroup, publicGroup)
	beeRouter.InitBeeOrderGoodsRouter(privateGroup, publicGroup)
	beeRouter.InitBeeOrderLogisticsRouter(privateGroup, publicGroup)
	beeRouter.InitBeeUploadFileRouter(privateGroup, publicGroup)
	beeRouter.InitBeePayLogRouter(privateGroup, publicGroup)
	beeRouter.InitBeeWxPayConfigRouter(privateGroup, publicGroup)
	beeRouter.InitBeeNoticeRouter(privateGroup, publicGroup)
	beeRouter.InitBeeLogisticsRouter(privateGroup, publicGroup)
	beeRouter.InitBeeOrderPeisongRouter(privateGroup, publicGroup)

	beeRouter.InitBeeOrderLogRouter(privateGroup, publicGroup)
	beeRouter.InitBeePrinterRouter(privateGroup, publicGroup)
	beeRouter.InitBeeDeliveryRouter(privateGroup, publicGroup)
	beeRouter.InitBeeOrderPeisongLogRouter(privateGroup, publicGroup)
	beeRouter.InitBeeQueueRouter(privateGroup, publicGroup)
	beeRouter.InitBeeUserQueueRouter(privateGroup, publicGroup)
	beeRouter.InitBeeRechargeSendRuleRouter(privateGroup, publicGroup)
}

func (*BeeShopPlugin) RouterPath() string {
	return "bee-shop"
}
