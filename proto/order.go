package proto

import (
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/model"
	"github.com/shopspring/decimal"
	"github.com/spf13/cast"
	"math"
	"time"
)

type BeeOrderGoods struct {
	AfterSale        string                 `json:"afterSale"`
	Amount           decimal.Decimal        `json:"amount"`
	AmountCoupon     decimal.Decimal        `json:"amountCoupon"`
	AmountSingle     decimal.Decimal        `json:"amountSingle"`
	AmountSingleBase decimal.Decimal        `json:"amountSingleBase"`
	BuyRewardEnd     bool                   `json:"buyRewardEnd"`
	CategoryID       int64                  `json:"categoryId"`
	CyTableStatus    enum.CyTableStatus     `json:"cyTableStatus"`
	DateAdd          string                 `json:"dateAdd"`
	FxType           enum.FxType            `json:"fxType"`
	GoodsID          int64                  `json:"goodsId"`
	GoodsIDStr       string                 `json:"goodsIdStr"`
	GoodsName        string                 `json:"goodsName"`
	ID               int64                  `json:"id"`
	InviterId        int64                  `json:"inviter_id"`
	IsProcessHis     bool                   `json:"isProcessHis"`
	IsScoreOrder     bool                   `json:"isScoreOrder"`
	LogisticsType    int                    `json:"logisticsType"`
	Number           int64                  `json:"number"`
	NumberNoFahuo    int64                  `json:"numberNoFahuo"`
	OrderID          string                 `json:"orderId"`
	Persion          int64                  `json:"persion"`
	Pic              string                 `json:"pic"`
	PriceID          int64                  `json:"priceId"`
	Property         string                 `json:"property"`
	PropertyChildIds string                 `json:"propertyChildIds"`
	Purchase         bool                   `json:"purchase"`
	RefundStatus     enum.OrderRefundStatus `json:"refundStatus"`
	Score            decimal.Decimal        `json:"score"`
	ShopID           int64                  `json:"shopId"`
	Status           enum.OrderGoodsStatus  `json:"status"`
	Type             int64                  `json:"type"`
	UID              int64                  `json:"uid"`
	Unit             string                 `json:"unit"`
	UserID           int64                  `json:"userId"`
}

func BeeOrderGoodsToDto(goods *model.BeeOrderGoods) *BeeOrderGoods {
	return &BeeOrderGoods{
		AfterSale:        goods.AfterSale,
		Amount:           goods.Amount,
		AmountCoupon:     goods.AmountCoupon,
		AmountSingle:     goods.AmountSingle,
		AmountSingleBase: goods.AmountSingleBase,
		BuyRewardEnd:     goods.BuyRewardEnd,
		CategoryID:       goods.CategoryId,
		CyTableStatus:    goods.CyTableStatus,
		DateAdd:          goods.DateAdd.ToString(),
		FxType:           goods.FxType,
		GoodsID:          goods.GoodsId,
		GoodsIDStr:       cast.ToString(goods.GoodsId),
		GoodsName:        goods.GoodsName,
		ID:               goods.Id,
		IsScoreOrder:     goods.IsScoreOrder,
		Number:           goods.Number,
		NumberNoFahuo:    goods.NumberNoFahuo,
		OrderID:          cast.ToString(goods.OrderId),
		Persion:          goods.Persion,
		Pic:              goods.Pic,
		PriceID:          goods.PriceId,
		Property:         goods.Property,
		PropertyChildIds: goods.PropertyChildIds,
		Purchase:         goods.Purchase,
		RefundStatus:     goods.RefundStatus,
		Score:            goods.Score,
		ShopID:           goods.ShopId,
		Status:           goods.Status,
		Type:             goods.Type,
		UID:              goods.Uid,
		UserID:           goods.UserId,
		Unit:             goods.Unit,
	}
}

type LogisticsMap struct {
}

type OrderInfo struct {
	Amount            decimal.Decimal `json:"amount"`
	AmountCard        decimal.Decimal `json:"amountCard"`
	AmountCoupons     decimal.Decimal `json:"amountCoupons"`
	AmountLogistics   decimal.Decimal `json:"amountLogistics"`
	AmountReal        decimal.Decimal `json:"amountReal"`
	AmountRefundTotal decimal.Decimal `json:"amountRefundTotal"`
	AmountTax         decimal.Decimal `json:"amountTax"`
	AmountTaxGst      decimal.Decimal `json:"amountTaxGst"`
	AmountTaxService  decimal.Decimal `json:"amountTaxService"`
	AutoDeliverStatus int             `json:"autoDeliverStatus"`
	DateAdd           string          `json:"dateAdd"`
	DateClose         string          `json:"dateClose"`
	DatePay           string          `json:"datePay"`
	DateUpdate        string          `json:"dateUpdate"`
	DifferHours       int             `json:"differHours"`
	GoodsNumber       int             `json:"goodsNumber"`
	HasRefund         bool            `json:"hasRefund"`
	HxNumber          string          `json:"hxNumber"`
	ID                int             `json:"id"`
	IP                string          `json:"ip"`
	IsCanHx           bool            `json:"isCanHx"`
	IsDelUser         bool            `json:"isDelUser"`
	IsEnd             bool            `json:"isEnd"`
	IsHasBenefit      bool            `json:"isHasBenefit"`
	IsNeedLogistics   bool            `json:"isNeedLogistics"`
	IsPay             bool            `json:"isPay"`
	IsScoreOrder      bool            `json:"isScoreOrder"`
	IsSuccessPingtuan bool            `json:"isSuccessPingtuan"`
	Jd8Status         int             `json:"jd8Status"`
	NearbyCloseMillis int             `json:"nearbyCloseMillis"`
	OrderNumber       string          `json:"orderNumber"`
	OrderType         int             `json:"orderType"`
	PeriodAutoPay     bool            `json:"periodAutoPay"`
	Pid               int             `json:"pid"`
	Qudanhao          string          `json:"qudanhao"`
	RefundStatus      int             `json:"refundStatus"`
	Remark            string          `json:"remark"`
	Score             int             `json:"score"`
	ScoreDeduction    int             `json:"scoreDeduction"`
	ShopID            int             `json:"shopId"`
	ShopIDZt          int             `json:"shopIdZt"`
	ShopNameZt        string          `json:"shopNameZt"`
	Status            int             `json:"status"`
	StatusStr         string          `json:"statusStr"`
	Trips             decimal.Decimal `json:"trips"`
	Type              int             `json:"type"`
	UID               int             `json:"uid"`
	UserID            int             `json:"userId"`
}

type GetOrderListResp struct {
	GoodsMap     map[string][]*BeeOrderGoods `json:"goodsMap"`
	LogisticsMap *LogisticsMap               `json:"logisticsMap"`
	OrderList    []*OrderInfo                `json:"orderList"`
	TotalPage    int                         `json:"totalPage"`
	TotalRow     int                         `json:"totalRow"`
}

type CreateOrderResp struct {
	Id                     int64             `json:"id"`
	Amount                 decimal.Decimal   `json:"amount"`     //货款
	AmountReal             decimal.Decimal   `json:"amountReal"` //实际需要支付金额
	AmountTax              int               `json:"amountTax"`
	AmountTaxGst           int               `json:"amountTaxGst"`
	AmountTaxService       int               `json:"amountTaxService"`
	AmountTotle            decimal.Decimal   `json:"amountTotle"`            //扣除积分、余额前的金额
	AmountTotleOriginal    decimal.Decimal   `json:"amountTotleOriginal"`    //优惠前商品金额
	AmountCoupons          decimal.Decimal   `json:"amountCoupons"`          //优惠券抵扣金额
	AmountLogistics        decimal.Decimal   `json:"amountLogistics"`        //运费原价
	AmountLogisticsReal    decimal.Decimal   `json:"amountLogisticsReal"`    //真实运费金额
	AmountLogisticsCoupons decimal.Decimal   `json:"amountLogisticsCoupons"` //优惠券抵扣运费金额
	CouponId               []int64           `json:"couponId"`
	CouponUserList         []*UserCouponResp `json:"couponUserList"`
	DeductionMoney         int               `json:"deductionMoney"`
	DeductionScore         int               `json:"deductionScore"`
	FreightScore           int               `json:"freightScore"`
	GoodsNumber            int64             `json:"goodsNumber"`
	IsNeedLogistics        bool              `json:"isNeedLogistics"`
	Overseas               bool              `json:"overseas"`
	Score                  int               `json:"score"`
	HxNumber               string            `json:"hxNumber"`
	IsPay                  bool              `json:"isPay"`
	OrderNumber            string            `json:"orderNumber"`
	NearbyCloseMillis      int64             `json:"nearbyCloseMillis"` // 订单关闭时间
	Status                 enum.OrderStatus  `json:"status"`
}

type ListOrderResp struct {
	TotalRow      int64                               `json:"totalRow"`
	TotalPage     int64                               `json:"totalPage"`
	OrderList     []*OrderDto                         `json:"orderList"`
	LogisticsMap  map[string]*model.BeeOrderLogistics `json:"logisticsMap"`
	OrderGoodsMap map[string][]*model.BeeOrderGoods   `json:"goodsMap"`
}

type ListOrderReq struct {
	DateAddBegin              string               `json:"dateAddBegin" form:"DateAddBegin"`
	DateAddEnd                string               `json:"dateAddEnd" form:"DateAddEnd"`
	DateRange                 string               `json:"dateRange" form:"DateRange"`           //	下单时间范围:今天 today;昨天 yestoday; 本周 week; 本月 month;上月 lastmonth;今年 year;去年 lastyear;近30天 30day	query	false
	GoodReputation            enum.OrderReputation `json:"goodReputation" form:"goodReputation"` // 评价状态：0 差评 1 中评 2 好评		false
	GoodsId                   string               `json:"goodsId" form:"GoodReputation"`
	HasRefund                 *bool                `json:"hasRefund" form:"hasRefund"`                       //	true / false ，读取是否有退款的订单		false
	IsNeedLogistics           *bool                `json:"isNeedLogistics" form:"isNeedLogistics"`           //订单是否需要物流		false
	LogisticsMobile           string               `json:"logisticsMobile" form:"GoodsId"`                   //收货人手机号码
	LogisticsNameLike         string               `json:"logisticsNameLike" form:"HasRefund"`               //收货人姓名	query	false
	LogisticsNameOrmobileLike string               `json:"logisticsNameOrmobileLike" form:"IsNeedLogistics"` //收货人姓名/手机号码	query	false
	Mod                       int                  `json:"mod" form:"mod"`                                   //	0 显示自己的订单(不传默认0)； 1 为团长查看所有的订单; 2 查看收件人手机号码是自己的订单; 3 团队订单（下级分销商订单）; 4 查看我邀请的订单（推广人）; 5 查看我的拼团用户的订单; 6 接龙团长查看订单; 7 技师查看自己服务的订单	query	false

	OrderNumber        string            `json:"orderNumber" form:"LogisticsMobile"`              //	订单编号，如：OD1703281618955938		false
	Page               int               `json:"page" form:"page"`                                //	获取第几页数据,示例值(1)	query	false
	PageSize           int               `json:"pageSize" form:"pageSize"`                        //每页显示多少数据,示例值(50)	query	false
	PeisongStatusBatch string            `json:"peisongStatusBatch" form:"LogisticsNameLike"`     //批量拉取多种配送状态订单，多个状态之间用逗号隔开，1 待接单； 2 待取货（餐）； 3 配送中； 4 已完成； 5 已取消； 6 预留； 7 预留； 8 预留； 9 退回中； 10 已退回；	query	false
	PingtuanOpenId     string            `json:"pingtuanOpenId" form:"LogisticsNameOrmobileLike"` //	拼团购买的团号		false
	RemarkLike         string            `json:"remarkLike" form:"Mod"`                           //根据下单备注搜索	query	false
	ShopId             string            `json:"shopId" form:"OrderNumber"`                       //	所属门店ID，用于查询该门店下的订单		false
	ShowAttendant      *bool             `json:"showAttendant" form:"showAttendant"`              //显示技师信息	query	false
	ShowExtJson        *bool             `json:"showExtJson" form:"showExtJson"`                  //显示扩展属性	query	false
	ShowShopInfo       *bool             `json:"showShopInfo" form:"showShopInfo"`                //获取门店数据	query	false
	Status             *enum.OrderStatus `json:"status" form:"status"`                            //-1 已关闭 0 待支付 1 已支付待发货 2 已发货待确认 3 确认收货待评价 4 已评价		false
	StatusBatch        string            `json:"statusBatch" form:"Page"`                         //，多个状态之间用逗号隔开，-1 已关闭 0 待支付 1 已支付待发货 2 已发货待确认 3 确认收货待评价 4 已评价	query	false
	StatusOuterBatch   string            `json:"statusOuterBatch" form:"PageSize"`                //批量拉取 status_outer 状态，多个状态之间用逗号隔开	query	false
	Type2              *bool             `json:"type2" form:"type2"`                              //	自定义的订单类型		false
}

type OrderDto struct {
	model.BeeOrder
	DifferHours int64 `json:"differ_hours,omitempty"` // 下单离现在相差多少个小时
}

func BeeOrder2Dto(item *model.BeeOrder) *OrderDto {
	return &OrderDto{
		BeeOrder:    *item,
		DifferHours: int64(math.Floor(time.Now().Sub(time.Time(item.DateAdd)).Hours())),
	}
}

type BeeOrderReputationReq struct {
	Token       string                    `json:"token"`
	OrderId     string                    `json:"orderId"`
	Reputations []*BeeOrderReputationItem `json:"reputations"`
}

type BeeOrderReputationItem struct {
	Id         string               `json:"id"`
	Reputation enum.OrderReputation `json:"reputation"`
	Remark     string               `json:"remark"`
	Pics       []string             `json:"pics"`
}

type CreateOrderReq struct {
	GoodsJsonStr        string `json:"goodsJsonStr" form:"goodsJsonStr"`               //	购买的商品、规格尺寸、数量信息的数组，如：[{"goodsId":11,"number":2,"propertyChildIds":"","logisticsType":0, "days": ["2019-07-26", "2019-07-27"],"wholesaleId":0}, {"goodsId":8,"number":3,"propertyChildIds":"2:9","logisticsType":0, "inviter_id":邀请用户ID, "days": ["2019-07-26", "2019-07-27"]}]	query	true
	Address             string `json:"address" form:"address"`                         //	收货地址详细地址	query	false
	AdminPurchase       string `json:"adminPurchase" form:"adminPurchase"`             //	区管从总店进货	query	false
	AttendanDay         string `json:"attendanDay" form:"attendanDay"`                 //	服务日期， 格式 yyyy-MM-dd	query	false
	AttendanTime        string `json:"attendanTime" form:"attendanTime"`               //	服务时间	query	false
	AttendanTime2       string `json:"attendanTime2" form:"attendanTime2"`             //	服务时间2，预留字段	query	false
	AttendantId         string `json:"attendantId" form:"attendantId"`                 //	技师/服务人员ID	query	false
	AutoDeliver         string `json:"autoDeliver" form:"autoDeliver"`                 //	是否自动发货，true 启用，需要开通自动发货插件后方可使用	query	false
	AutoPeisong         string `json:"autoPeisong" form:"autoPeisong"`                 //	是否启用配送联动,激活并使用配送模块	query	false
	Calculate           string `json:"calculate" form:"calculate"`                     //	true 不实际下单，而是返回价格计算	query	false
	CardId              string `json:"cardId" form:"cardId"`                           //	用户名下会员卡记录ID，注意不是会员卡ID	query	false
	CityId              string `json:"cityId" form:"cityId"`                           //	收货地址城市编码	query	false
	Code                string `json:"code" form:"code"`                               //	收货地址邮政编码	query	false
	CouponId            string `json:"couponId" form:"couponId"`                       //   使用的优惠券编号,多张优惠券请用英文的逗号分隔	query	false
	CouponPriority      string `json:"couponPriority" form:"couponPriority"`           //	传true，自动使用一张最优惠的优惠券	query	false
	DadaShopNo          string `json:"dadaShopNo" form:"dadaShopNo"`                   //	达达商户ID	query	false
	DeductionScore      string `json:"deductionScore" form:"deductionScore"`           //	用多少积分来抵扣本次交易，请配合后台积分抵扣规则使用，小于等于0代表自动用抵扣，保留绝对值金额不允许抵扣	query	false
	DistrictId          string `json:"districtId" form:"districtId"`                   //	收货地址区县编码	query	false
	Email               string `json:"email" form:"email"`                             //	京东云交易订单接收邮箱地址	query	false
	ExtJsonStr          string `json:"extJsonStr" form:"extJsonStr"`                   //	订单扩展属性信息,JSON格式	query	false
	ExtType             string `json:"extType" form:"extType"`                         //	自定义订单类型，最多32字	query	false
	FilterShopId        string `json:"filterShopId" form:"filterShopId"`               //	只下单指定门店商品，其他商品将被忽略	query	false
	GoodsType           string `json:"goodsType" form:"goodsType"`                     //	0 自营商品； 1 京东vop商品; 2 京东权益商品	query	false
	Grabkey             string `json:"grabkey" form:"grabkey"`                         //	如果开启秒杀模式，必须先调用秒杀抢占名额接口获取grabkey	query	false
	Idcard              string `json:"idcard" form:"idcard"`                           //	身份证号码【国际件使用】	query	false
	IgnoreDiscountPrice string `json:"ignoreDiscountPrice" form:"ignoreDiscountPrice"` //	如果传，将忽略商品设置里面的优惠金额	query	false
	IsCanHx             string `json:"isCanHx" form:"isCanHx"`                         //	是否支持核销，true 支持，默认不支持	query	false
	Kjid                string `json:"kjid" form:"kjid"`                               //	砍价编号，可直接购买砍价商品	query	false
	Lat                 string `json:"lat" form:"lat"`                                 //	收货地址的纬度	query	false
	LinkMan             string `json:"linkMan" form:"linkMan"`                         //	收货地址联系人信息	query	false
	ListingId           string `json:"listingId" form:"listingId"`                     //	如果是接龙下单，需要传接龙ID	query	false
	Lng                 string `json:"lng" form:"lng"`                                 //	收货地址的经度	query	false
	Menpai              string `json:"menpai" form:"menpai"`                           //	门牌	query	false
	Mobile              string `json:"mobile" form:"mobile"`                           //	收货地址联系人手机号码	query	false
	OrderInviterId      string `json:"orderInviterId" form:"orderInviterId"`           //	推荐人用户id	query	false
	OrderPeriod         string `json:"orderPeriod" form:"orderPeriod"`                 //	周期订单参数: {"unit":1,"duration":1,"dateStart":"2020-05-05 11:47:19", "dateStart": false}, unit: 0 天 1 月 2 年	query	false
	PayOnDelivery       string `json:"payOnDelivery" form:"payOnDelivery"`             //	1 为货到付款，其他数字为先支付	query	false
	PeisongFeeId        string `json:"peisongFeeId" form:"peisongFeeId"`               //	配送费的id	query	false
	PeisongType         string `json:"peisongType" form:"peisongType"`                 //	配送类型，kd 代表快递；zq代表到店自取； keloop 快跑者; pszq 商家配送自取(还需收取快递费); meituan 美团配送; shansong 闪送	query	false
	PickPointId         string `json:"pickPointId" form:"pickPointId"`                 //	取货点ID	query	false
	PingtuanOpenId      string `json:"pingtuanOpenId" form:"pingtuanOpenId"`           //	拼团购买的团号	query	false
	ProvinceId          string `json:"provinceId" form:"provinceId"`                   //	收货地址省份编码	query	false
	Remark              string `json:"remark" form:"remark"`                           //	下单备注信息	query	false
	ScoreExchangeClose  string `json:"scoreExchangeClose" form:"scoreExchangeClose"`   //	如果不传或者传false，积分商品积分不足的时候，自动根据兑换比例折算成钱结算	query	false
	ScoreForceExchange  string `json:"scoreForceExchange" form:"scoreForceExchange"`   //	如果传true，所有的积分折算成金额进行结算	query	false
	ShipmentType        string `json:"shipmentType" form:"shipmentType"`               //	京东云交易配送方式。1：京东配送 2：京配转三方配送 3：第三方配送 4：普通快递配送 9：不支持配送	query	false
	ShopIdZt            int64  `json:"shopIdZt" form:"shopIdZt"`                       //	自提门店id	query	false
	ShopNameZt          string `json:"shopNameZt" form:"shopNameZt"`                   //	自提门店名称	query	false
	ShopPurchase        string `json:"shopPurchase" form:"shopPurchase"`               //	门店从区管进货	query	false
	StreetId            string `json:"streetId" form:"streetId"`                       //	收货地址街道/社区编码	query	false
	Trips               string `json:"trips" form:"trips"`                             //	小费金额，用户可自行输入小费金额	query	false
	Xiaoqu              string `json:"xiaoqu" form:"xiaoqu"`                           //	小区名称	query	false
	TableNum            string `json:"tableNum" form:"tableNum"`
}

type GetOrderDetailResp struct {
	OrderInfo *OrderDto                `json:"orderInfo"`
	Goods     []*BeeOrderGoods         `json:"goods"`
	Logistics *model.BeeOrderLogistics `json:"logistics"`
	Logs      []*model.BeeOrderLog     `json:"logs"`
}
