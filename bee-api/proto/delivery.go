package proto

import "github.com/shopspring/decimal"

// QueryDeliveryResult 查询配送订单信息
type QueryDeliveryResult struct {
	OrderId          string          `json:"orderId"`          //	第三方订单编号
	StatusCode       int             `json:"statusCode"`       //	订单状态(待接单＝1,待取货＝2,配送中＝3,已完成＝4,已取消＝5, 指派单=8,妥投异常之物品返回中=9, 妥投异常之物品返回完成=10, 骑士到店=100,创建达达运单失败=1000 可参考文末的状态说明）
	StatusMsg        string          `json:"statusMsg"`        //	订单状态
	TransporterName  string          `json:"transporterName"`  //	骑手姓名
	TransporterPhone string          `json:"transporterPhone"` //	骑手电话
	TransporterLng   string          `json:"transporterLng"`   //	骑手经度
	TransporterLat   string          `json:"transporterLat"`   //	骑手纬度
	DeliveryFee      decimal.Decimal `json:"deliveryFee"`      //	订单总费用，包含运费、小费、保价费
	Tips             decimal.Decimal `json:"tips"`             //	小费,单位为元
	CouponFee        decimal.Decimal `json:"couponFee"`        //	优惠券费用,单位为元
	InsuranceFee     decimal.Decimal `json:"insuranceFee"`     //	保价费,单位为元
	ActualFee        decimal.Decimal `json:"actualFee"`        //	订单实际费用消耗，订单总费用减优惠券
	Distance         float64         `json:"distance"`         //	配送距离,单位为米
	CreateTime       string          `json:"createTime"`       //	发单时间
	AcceptTime       string          `json:"acceptTime"`       //	接单时间,若未接单,则为空
	FetchTime        string          `json:"fetchTime"`        //	取货时间,若未取货,则为空
	FinishTime       string          `json:"finishTime"`       //	送达时间,若未送达,则为空
	CancelTime       string          `json:"cancelTime"`       //	取消时间,若未取消,则为空
	OrderFinishCode  string          `json:"orderFinishCode"`  //	收货码
	DeductFee        decimal.Decimal `json:"deductFee"`        //	违约金
}

// QueryDeliverFeeResult 查询配送订单费用res
type QueryDeliverFeeResult struct {
	Distance     float64         `json:"distance"`     //		是	配送距离(单位：米)
	DeliveryNo   string          `json:"deliveryNo"`   //		是	平台订单号
	Fee          decimal.Decimal `json:"fee"`          //	是	实际运费(单位：元)，运费减去优惠券费用
	DeliverFee   decimal.Decimal `json:"deliver_fee"`  //		是	运费(单位：元)
	CouponFee    decimal.Decimal `json:"couponFee"`    //	否	优惠券费用(单位：元)
	Tips         float64         `json:"tips"`         //	否	小费（单位：元，精确小数点后一位，小费金额不能高于订单金额。）
	InsuranceFee decimal.Decimal `json:"insuranceFee"` //		否	保价费(单位：元)
}

// QueryDeliverFeeReq 查询配送订单费用req
type QueryDeliverFeeReq struct {
	ShopNo          string          `json:"shop_no"`                                   //	是 	门店编号，门店创建后可在门店列表查看
	OriginId        string          `json:"origin_id"`                                 //	是	第三方订单ID
	CargoPrice      decimal.Decimal `json:"cargo_price"`                               //	是	订单金额（单位：元）
	IsPrepay        int8            `json:"is_prepay"`                                 //	是	是否需要垫付 1:是 0:否 (垫付订单金额，非运费)
	ReceiverName    string          `json:"receiver_name"`                             //	是	收货人姓名
	ReceiverAddress string          `json:"receiver_address"`                          //	是	收货人地址
	Callback        string          `json:"callback"`                                  //	是	回调URL（查看回调说明）
	CargoWeight     float64         `json:"cargo_weight"`                              //	是	订单重量（单位：Kg）
	ReceiverLat     float64         `json:"proto.QueryDeliverFeeResult{}receiver_lat"` //	是	收货人地址纬度（支持火星坐标系(高德、腾讯地图))
	ReceiverLng     float64         `json:"receiver_lng"`                              // 	是	收货人地址经度（支持火星坐标系(高德、腾讯地图))
	ReceiverPhone   string          `json:"receiver_phone,omitempty"`                  //	否	收货人手机号（手机号和座机号必填一项），如传入隐私号，请使用英文逗号","分隔格式，并确保支持接收短信
	ReceiverTel     string          `json:"receiver_tel,omitempty"`                    //	否	收货人座机号（手机号和座机号必填一项）
	Tips            decimal.Decimal `json:"tips,omitempty"`                            //	否	小费（单位：元，精确小数点后一位，小费金额不能高于订单金额。）
	Info            string          `json:"info,omitempty"`                            //	否	订单备注
	CargoType       int             `json:"cargo_type,omitempty"`                      //	否	支持配送的物品品类见链接。请选择门店真实配送品类，如无法判断可咨询您的销售经理。
	CargoNum        int             `json:"cargo_num,omitempty"`                       //	否	订单商品数量
}

// AddOrderDirectReq 直接下单req
type AddOrderDirectReq struct {
	QueryDeliverFeeReq
	DeliveryNo  string `json:"deliveryNo"` //	是	平台订单编号
	DaySeq      int64  `json:"daySeq"`     //流水号
	Source      string `json:"source"`
	ShopAddress string `json:"shop_address"`
	ShopName    string `json:"shop_name"`
	ShopPhone   string `json:"shop_phone"`
}

// AddOrderDirectResp 直接下单resp
type AddOrderDirectResp struct {
	QueryDeliverFeeResult
}

type AddDeliverOrderReq struct {
	DeliveryNo  string `json:"deliveryNo"`     //	是	平台订单编号
	EnableReset bool   `json:"enableReset"`    //	否	是否重置订单备注等字段。true：重置，false或缺失：不重置
	Info        string `json:"info,omitempty"` //	否	订单备注，长度限制128
}

type CancelDeliverOrderReq struct {
	OrderId        string `json:"order_id"`
	CancelReasonId int64  `json:"cancel_reason_id"`
	CancelReason   string `json:"cancel_reason"`
	ShopId         string `json:"shop_id"`
	Source         string `json:"source"` //订单来源
}

type CancelDeliverOrderRes struct {
	DeductFee decimal.Decimal `json:"deduct_fee"`
}
