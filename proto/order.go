package proto

import (
	"gitee.com/stuinfer/bee-api/enum"
	"gitee.com/stuinfer/bee-api/model"
	"github.com/shopspring/decimal"
	"math"
	"time"
)

type BeeOrderGoods struct {
	AfterSale        string             `json:"afterSale"`
	Amount           decimal.Decimal    `json:"amount"`
	AmountCoupon     decimal.Decimal    `json:"amountCoupon"`
	AmountSingle     decimal.Decimal    `json:"amountSingle"`
	AmountSingleBase decimal.Decimal    `json:"amountSingleBase"`
	BuyRewardEnd     bool               `json:"buyRewardEnd"`
	CategoryID       int64              `json:"categoryId"`
	CyTableStatus    enum.CyTableStatus `json:"cyTableStatus"`
	DateAdd          string             `json:"dateAdd"`
	FxType           enum.FxType        `json:"fxType"`
	GoodsID          int64              `json:"goodsId"`
	GoodsIDStr       string             `json:"goodsIdStr"`
	GoodsName        string             `json:"goodsName"`
	ID               int64              `json:"id"`
	InviterId        int64              `json:"inviter_id"`
	IsProcessHis     bool               `json:"isProcessHis"`
	IsScoreOrder     bool               `json:"isScoreOrder"`
	LogisticsType    int                `json:"logisticsType"`
	Number           int64              `json:"number"`
	NumberNoFahuo    int64              `json:"numberNoFahuo"`
	OrderID          string             `json:"orderId"`
	Persion          int64              `json:"persion"`
	Pic              string             `json:"pic"`
	PriceID          int64              `json:"priceId"`
	Property         string             `json:"property"`
	PropertyChildIds string             `json:"propertyChildIds"`
	Purchase         bool               `json:"purchase"`
	RefundStatus     int                `json:"refundStatus"`
	Score            decimal.Decimal    `json:"score"`
	ShopID           int64              `json:"shopId"`
	Status           int                `json:"status"`
	Type             int                `json:"type"`
	UID              int64              `json:"uid"`
	Unit             string             `json:"unit"`
	UserID           int64              `json:"userId"`
}

type LogisticsMap struct {
}

type OrderList struct {
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
	GoodsMap     map[string][]BeeOrderGoods `json:"goodsMap"`
	LogisticsMap LogisticsMap               `json:"logisticsMap"`
	OrderList    []OrderList                `json:"orderList"`
	TotalPage    int                        `json:"totalPage"`
	TotalRow     int                        `json:"totalRow"`
}

type CreateOrderResp struct {
	AmountLogistics     int               `json:"amountLogistics"`
	AmountReal          decimal.Decimal   `json:"amountReal"`
	AmountTax           int               `json:"amountTax"`
	AmountTaxGst        int               `json:"amountTaxGst"`
	AmountTaxService    int               `json:"amountTaxService"`
	AmountTotle         decimal.Decimal   `json:"amountTotle"`
	AmountTotleOriginal decimal.Decimal   `json:"amountTotleOriginal"`
	CouponAmount        int               `json:"couponAmount"`
	CouponId            []int64           `json:"couponId"`
	CouponUserList      []*UserCouponResp `json:"couponUserList"`
	DeductionMoney      int               `json:"deductionMoney"`
	DeductionScore      int               `json:"deductionScore"`
	FreightScore        int               `json:"freightScore"`
	GoodsNumber         int               `json:"goodsNumber"`
	IsNeedLogistics     bool              `json:"isNeedLogistics"`
	Overseas            bool              `json:"overseas"`
	Score               int               `json:"score"`
}

type ListOrderResp struct {
	TotalRow      int64                               `json:"totalRow"`
	TotalPage     int64                               `json:"totalPage"`
	OrderList     []*OrderDto                         `json:"orderList"`
	LogisticsMap  map[string]*model.BeeOrderLogistics `json:"logisticsMap"`
	OrderGoodsMap map[string]*model.BeeOrderGoods     `json:"goodsMap"`
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
