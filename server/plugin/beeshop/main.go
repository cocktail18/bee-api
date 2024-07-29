package beeshop

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"gitee.com/stuinfer/bee-api/cmd"
	"gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/logger"
	"gitee.com/stuinfer/bee-api/util"
	config2 "github.com/flipped-aurora/gin-vue-admin/server/config"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/beeshop/router"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/plugin-tool/utils"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"io/ioutil"
	"sort"
	"strings"
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
	if !global.GVA_CONFIG.BeeShop.Enable {
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
		menuPath := "_" + menu.Path
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
			ApiGroup:    api.ApiGroup,
			Method:      api.Method,
		})
	}
	// 下方会自动注册api 以下格式为示例格式，请按照实际情况修改
	utils.RegisterApis(
		apiAdd...,
	)
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
	ins := &BeeShopPlugin{}
	go func() {
		ins.startBeeApi()
	}()
	return ins
}

func (*BeeShopPlugin) startBeeApi() {
	var dbCfg config2.SpecializedDB
	for _, info := range global.GVA_CONFIG.DBList {
		if info.Disable { // 禁用数据库
			continue
		}
		if info.AliasName == "bee" {
			dbCfg = info
		}
	}
	if dbCfg.Path == "" {
		dbCfg.Path = "127.0.0.1"
	}
	lockFile := "./bee-init.lock"
	b, err := util.FileExists(lockFile)
	if err != nil {
		panic(err)
	}
	listen := global.GVA_CONFIG.BeeShop.Listen
	if listen == "" {
		listen = "127.0.0.1:18083"
	}
	appCfg := &config.AppConfig{
		App: &config.App{
			Listen:  listen,
			DfsHost: "",
		},
		DB: &config.AppDBConfig{
			Host:     dbCfg.Path,
			Port:     cast.ToInt(dbCfg.Port),
			User:     dbCfg.Username,
			Password: dbCfg.Password,
			Database: dbCfg.Dbname,
			Drop:     false,
			NeedInit: !b,
		},
		Default: &config.DefaultConfig{
			Wx: &config.WxConfig{
				AppId:  "test", //改成你的
				Secret: "test", //改成你的
			},
			SysUser: &config.SysUser{
				Domain:   "cocktailBeeOrder",
				Username: "bee",
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
	_ = ioutil.WriteFile(lockFile, []byte("初始化完成后请不要删除该文件"), 0644)
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

	beeRouter.InitBeeOrderLogRouter(privateGroup, publicGroup)
}

func (*BeeShopPlugin) RouterPath() string {
	return "bee-shop"
}
