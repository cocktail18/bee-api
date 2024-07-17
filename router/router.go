package router

import (
	"gitee.com/stuinfer/bee-api/api"
	config2 "gitee.com/stuinfer/bee-api/config"
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/service"
	"gitee.com/stuinfer/bee-api/sys"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strings"
)

var router = gin.New()

var whiteUrlList = []string{
	"/user/wxapp/authorize",
	"/user/wxapp/login",
	"/user/level/list",
}

// 注册主账号id到context
func regSysUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		uri := c.Request.RequestURI
		idx := strings.Index(uri[1:], "/")
		domain := uri[1 : idx+1]

		sysUserInfo, err := sys.GetUserSrv().GetByDomain(domain)
		if err != nil {
			c.Abort()
			api.BaseApi{}.Fail(c, enum.ResCodeUnauthorized, "domain错误，请检查")
			return
		}
		c.Request.URL.Path = uri[idx+1:]
		c.Request.RequestURI = uri[idx+1:]
		c.Set(string(enum.CtxKeySysUser), sysUserInfo)
		c.Next()
	}
}

func CheckToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		if lo.Contains(whiteUrlList, c.Request.RequestURI) {
			c.Next()
			return
		}

		token := c.Query("token")
		if token == "" {
			token = c.PostForm("token")
		}

		if token == "" {
			c.Abort()
			api.BaseApi{}.Fail(c, enum.ResCodeTokenInvalid, "当前登录token无效，请重新登录")
			return
		}
		userInfo, err := service.GetUserSrv().GetUserInfo(c, token)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.Abort()
				api.BaseApi{}.Fail(c, enum.ResCodeTokenInvalid, "当前登录token无效，请重新登录")
				return
			}
			c.Abort()
			api.BaseApi{}.Fail(c, enum.ResCodeFail, err.Error())
			return
		}
		c.Set(enum.UserInfoKey, userInfo)
		c.Next()
	}
}

type justFilesFilesystem struct {
	fs http.FileSystem
}

func (fs justFilesFilesystem) Open(name string) (http.File, error) {
	f, err := fs.fs.Open(name)
	if err != nil {
		return nil, err
	}

	stat, err := f.Stat()
	if stat.IsDir() {
		return nil, os.ErrPermission
	}

	return f, nil
}

func NewRouter() *gin.Engine {
	router.Use(gin.Logger(), gin.Recovery())
	router.StaticFS(config2.GetStorePath(), justFilesFilesystem{http.Dir(config2.GetStorePath())})
	domainGroup := router.Group("/:domain", regSysUser())
	configGroup := domainGroup.Group("/config")
	{
		configGroup.GET("/values", (api.ConfigApi{}).Values)
	}
	qrcodeGroup := domainGroup.Group("/qrcode")
	{
		qrcodeGroup.POST("/wxa/unlimit", (api.QrcodeApi{}).WxaUnLimit)
	}
	shopGroup := domainGroup.Group("/shop")
	{
		{ // 门店
			shopGroup.GET("/subshop/detail/v2", (api.ShopApi{}).SubShopDetail)
			shopGroup.POST("/subshop/list", (api.ShopApi{}).SubShopList)
		}
		{ //商品
			shopGroup.GET("/goods/category/all", (api.GoodsApi{}).CategoryAll)
			shopGroup.POST("/goods/list", (api.GoodsApi{}).List)
			shopGroup.GET("/goods/list/v2", (api.GoodsApi{}).List)
			shopGroup.POST("/goods/price", (api.GoodsApi{}).Price)
			shopGroup.GET("/goods/detail", (api.GoodsApi{}).Detail)
			shopGroup.GET("/goods/goodsAddition", (api.GoodsApi{}).GoodsAddition)
			shopGroup.Any("/goods/times/schedule", (api.GoodsApi{}).TimesSchedule)
		}
		{
			//砍价
			shopGroup.GET("/goods/kanjia/info", (api.KanjiaApi{}).Info)
			shopGroup.GET("/goods/kanjia/help", (api.KanjiaApi{}).Help)
			shopGroup.GET("/goods/kanjia/clear", (api.KanjiaApi{}).Clear)
			shopGroup.GET("/goods/kanjia/my", (api.KanjiaApi{}).My)
			shopGroup.GET("/goods/kanjia/myHelp", (api.KanjiaApi{}).MyHelp)
			shopGroup.GET("/goods/kanjia/set/v2", (api.KanjiaApi{}).Set)
			shopGroup.POST("/goods/kanjia/join", (api.KanjiaApi{}).Join)
		}
	}
	cmsGroup := domainGroup.Group("/cms")
	{
		cmsGroup.GET("/page/info/v2", (api.CmsApi{}).Info)
	}
	commentGroup := domainGroup.Group("/comment", CheckToken())
	{
		//意见反馈
		commentGroup.POST("/add", (api.CommentApi{}).Add)
	}
	noticeGroup := domainGroup.Group("/notice")
	{
		noticeGroup.GET("/last-one", (api.NoticeApi{}).LastOne)
		noticeGroup.GET("/detail", (api.NoticeApi{}).Detail)
	}
	bannerGroup := domainGroup.Group("/banner")
	{
		bannerGroup.GET("/list", (api.BannerApi{}).List)
	}
	commonGroup := router.Group("/common")
	{
		commonGroup.GET("/region/v2/province", (api.CommonApi{}).Province)
		commonGroup.GET("/region/v2/child", (api.CommonApi{}).Child)
		commonGroup.POST("/map/qq/distance", (api.CommonApi{}).MapQQDistance)
	}
	yuyueGroup := domainGroup.Group("/yuyue", CheckToken())
	{
		yuyueGroup.POST("/join", (api.YuyueAPi{}).Join)
	}
	discountsGroup := domainGroup.Group("/discounts", CheckToken())
	{
		discountsGroup.GET("/my", (api.DiscountsApi{}).My)
		discountsGroup.GET("/detail", (api.DiscountsApi{}).CouponDetail)
		discountsGroup.GET("/statistics", (api.DiscountsApi{}).Statistics)
		discountsGroup.GET("/coupons", (api.DiscountsApi{}).Coupons)
		discountsGroup.POST("/fetch", (api.DiscountsApi{}).Fetch)
		discountsGroup.POST("/exchange", (api.DiscountsApi{}).Exchange)
	}
	payBillGroup := domainGroup.Group("/payBill", CheckToken())
	{ //满减
		payBillGroup.GET("/discounts", (api.PayBillApi{}).Discounts)
		payBillGroup.POST("/pay", (api.PayBillApi{}).Pay)
	}
	userGroup := domainGroup.Group("/user", CheckToken())
	{
		userGroup.GET("/bindSeller", (api.UserApi{}).BindSeller)
		userGroup.GET("/my", (api.UserApi{}).My)
		userGroup.POST("/level/list", (api.UserApi{}).LevelList)
		userGroup.GET("/detail", (api.UserApi{}).Detail)
		userGroup.POST("/modify", (api.UserApi{}).Modify)
		userGroup.GET("/amount", (api.UserApi{}).Amount)
		userGroup.GET("/check-token", (api.UserApi{}).CheckToken)
		userGroup.POST("/wxapp/authorize", (api.UserApi{}).WxAppAuthorize)
		userGroup.POST("/wxapp/login", (api.UserApi{}).WxAppAuthorize)
		userGroup.POST("/wxapp/bindMobile", (api.UserApi{}).BindMobile)
		userGroup.POST("/cashLog", (api.UserApi{}).CashLog)
		userGroup.POST("/cashLog/v2", (api.UserApi{}).CashLog)
		userGroup.GET("/dynamicUserCode", (api.UserApi{}).GetDynamicUserCode)

		userGroup.POST("/payLogs", (api.UserApi{}).PayLogs)
		userGroup.GET("/recharge/send/rule", (api.UserApi{}).RechargeSendRule)

		{ // 收获地址
			userGroup.GET("/shipping-address/list", (api.ShippingAddressApi{}).List)
			userGroup.GET("/shipping-address/default/v2", (api.ShippingAddressApi{}).Default)
			userGroup.GET("/shipping-address/detail/v2", (api.ShippingAddressApi{}).Detail)
			userGroup.POST("/shipping-address/add", (api.ShippingAddressApi{}).Add)
			userGroup.POST("/shipping-address/update", (api.ShippingAddressApi{}).Update)
			userGroup.POST("/shipping-address/delete", (api.ShippingAddressApi{}).Delete)
		}
	}
	shoppingCartGroup := domainGroup.Group("/shopping-cart", CheckToken())
	{
		shoppingCartGroup.GET("/info", (api.ShoppingCartAPi{}).Info)
		shoppingCartGroup.POST("/add", (api.ShoppingCartAPi{}).Add)
		shoppingCartGroup.POST("/remove", (api.ShoppingCartAPi{}).Remove)
		shoppingCartGroup.POST("/empty", (api.ShoppingCartAPi{}).Empty)
		shoppingCartGroup.POST("/modifyNumber", (api.ShoppingCartAPi{}).ModifyNumber)
	}
	orderGroup := domainGroup.Group("/order", CheckToken())
	{
		orderGroup.GET("/detail", (api.OrderApi{}).Detail)
		orderGroup.GET("/list", (api.OrderApi{}).List)
		orderGroup.POST("/list", (api.OrderApi{}).List)
		orderGroup.POST("/pay", (api.OrderApi{}).Pay)
		orderGroup.POST("/create", (api.OrderApi{}).Create)
		orderGroup.POST("/close", (api.OrderApi{}).Close)
		orderGroup.POST("/delete", (api.OrderApi{}).Delete)
		orderGroup.POST("/delivery", (api.OrderApi{}).Delivery)
		orderGroup.POST("/reputation", (api.OrderApi{}).Reputation) //评价
		orderGroup.POST("/hx", (api.OrderApi{}).Hx)                 //核销
	}
	scoreGroup := domainGroup.Group("/score", CheckToken())
	{
		scoreGroup.GET("/sign/logs", (api.ScoreApi{}).SignLogs)
		scoreGroup.POST("/sign/logs", (api.ScoreApi{}).SignLogs)
		scoreGroup.POST("/sign", (api.ScoreApi{}).Sign)
		scoreGroup.POST("/logs", (api.ScoreApi{}).Logs)
		scoreGroup.POST("/share/wxa/group", (api.ScoreApi{}).WxaGroup)
	}
	payGroup := domainGroup.Group("/pay", CheckToken())
	{
		payGroup.POST("/wx/wxapp", (api.PayApi{}).WxApp)
	}

	notifyGroup := domainGroup.Group("/notify")
	{
		notifyGroup.POST("/wx/pay", (api.PayApi{}).WxPayCallBack)
	}

	//扫码点餐
	cyTableGroup := domainGroup.Group("/cyTable", CheckToken())
	{
		cyTableGroup.POST("/add-order", (api.CyTableAPi{}).AddOrder)
		cyTableGroup.POST("/token", (api.CyTableAPi{}).Token)
	}
	//配送费
	feeGroup := domainGroup.Group("/fee")
	{
		feeGroup.GET("/peisong/list", (api.FeeAPi{}).ListPeiSong)
	}

	//文件上传
	dfsGroup := domainGroup.Group("/dfs", CheckToken())
	{
		dfsGroup.POST("/upload/file", (api.DfsApi{}).UploadFile)
	}
	//排队
	queuingGroup := domainGroup.Group("/queuing")
	{
		queuingGroup.POST("/get", CheckToken(), (api.QueuingApi{}).Get)
		queuingGroup.GET("/types", (api.QueuingApi{}).Types)
		queuingGroup.GET("/my", CheckToken(), (api.QueuingApi{}).My)
	}

	return router
}
