package dada_sdk

import (
	"context"
	"github.com/shopspring/decimal"
)

type QueryDeliverFeeReq struct {
	ShopNo          string          `json:"shop_no"`                  //	门店编号，门店创建后可在门店列表查看
	OriginId        string          `json:"origin_id"`                //是	第三方订单ID
	CargoPrice      decimal.Decimal `json:"cargo_price"`              //	是	订单金额（单位：元）
	IsPrepay        int8            `json:"is_prepay"`                //	是	是否需要垫付 1:是 0:否 (垫付订单金额，非运费)
	ReceiverName    string          `json:"receiver_name"`            //	是	收货人姓名
	ReceiverAddress string          `json:"receiver_address"`         //	是	收货人地址
	Callback        string          `json:"callback"`                 //	是	回调URL（查看回调说明）
	CargoWeight     float64         `json:"cargo_weight"`             //	是	订单重量（单位：Kg）
	ReceiverLat     float64         `json:"receiver_lat"`             //	是	收货人地址纬度（支持火星坐标系(高德、腾讯地图))
	ReceiverLng     float64         `json:"receiver_lng"`             // 是	收货人地址经度（支持火星坐标系(高德、腾讯地图))
	ReceiverPhone   string          `json:"receiver_phone,omitempty"` //	否	收货人手机号（手机号和座机号必填一项），如传入隐私号，请使用英文逗号","分隔格式，并确保支持接收短信
	ReceiverTel     string          `json:"receiver_tel,omitempty"`   //	否	收货人座机号（手机号和座机号必填一项）
	Tips            decimal.Decimal `json:"tips,omitempty"`           //	否	小费（单位：元，精确小数点后一位，小费金额不能高于订单金额。）
	Info            string          `json:"info,omitempty"`           //	否	订单备注
	CargoType       int             `json:"cargo_type,omitempty"`     //	否	支持配送的物品品类见链接。请选择门店真实配送品类，如无法判断可咨询您的销售经理。
	CargoNum        int             `json:"cargo_num,omitempty"`      //	否	订单商品数量
	InvoiceTitle    string          `json:"invoice_title,omitempty"`  //	否	发票抬头
	OriginMark      string          `json:"origin_mark,omitempty"`    //	否	订单来源标识（支持传汉字、字母，最大长度为10），请传入和小票相同的订单来源名称，避免骑士取错货；除以下已定义枚举值，如需自定义请联系业务经理开通。
	//传参：在骑士APP展示内容
	//elm：饿了么
	//mt：美团外卖
	//jdwl：京东物流
	OriginMarkNo   string `json:"origin_mark_no,omitempty"`   //	否	订单来源编号（支持数字、字母，最大长度为30），该字段会和origin_mark组合展示给骑士；示例：origin_mark传"mt"，origin_mark_no传"001"，则给骑士展示"美团外卖#001"
	IsUseInsurance int    `json:"is_use_insurance,omitempty"` //	否	是否使用保价费（0：不使用保价，1：使用保价； 同时，请确保填写了订单金额（cargo_price）
	//商品保价费(当商品出现损坏，可获取一定金额的赔付)
	//保费=配送物品实际价值*费率（5‰），配送物品价值及最高赔付不超过10000元， 最高保费为50元（物品价格最小单位为100元，不足100元部分按100元认定，保价费向上取整数， 如：物品声明价值为201元，保价费为300元*5‰=1.5元，取整数为2元。）
	//若您选择不保价，若物品出现丢失或损毁，最高可获得平台30元优惠券。 （优惠券直接存入用户账户中）。
	IsFinishCodeNeeded    int           `json:"is_finish_code_needed,omitempty"`    //	否	收货码（0：不需要；1：需要。收货码的作用是：骑手必须输入收货码才能完成订单妥投）
	DelayPublishTime      int           `json:"delay_publish_time,omitempty"`       //	否	预约发单时间（unix时间戳(10位)，精确到分），可以支持未来3天内的订单发预约单，如预约发单时间晚于当前时间会作为即时单下发。 查看预约单下单说明
	IsExpectFinishOrder   int           `json:"is_expect_finish_order,omitempty"`   //	否	是否根据期望送达时间预约发单（0-否，即时发单；1-是，预约发单），如传1则期望送达时间必传。 查看预约单下单说明
	ExpectFinishTimeLimit int64         `json:"expect_finish_time_limit,omitempty"` //	否	期望送达时间（单位秒，不早于当前时间） 查看预约单下单说明
	IsDirectDelivery      int           `json:"is_direct_delivery,omitempty"`       //	否	是否选择直拿直送（0：不需要；1：需要。选择直拿直送后，同一时间骑士只能配送此订单至完成，同时，也会相应的增加配送费用）
	ProductList           []ProductInfo `json:"product_list,omitempty"`             //	否	订单商品明细
	PickUpPos             string        `json:"pick_up_pos,omitempty"`              //	否	货架信息,该字段可在骑士APP订单备注中展示
	FetchCode             string        `json:"fetch_code,omitempty"`               //	否	取货码，在骑士取货时展示给骑士，门店通过取货码在骑士取货阶段核实骑士
	BusinessType          int           `json:"business_type,omitempty"`            //	否	物流配送方向(0-正向送货单、2-售后退货取件单)。如为售后单需传入发货人相关信息(售后场景下发货人一般为顾客，收货人为门店)，并请联系业务经理开通服务
	SupplierName          string        `json:"supplier_name,omitempty"`            //	否	发货人姓名，订单发货信息默认为门店信息，如需自定义则需传发货人5项信息(发货人姓名、地址、电话、纬度、经度)，并请联系业务经理
	SupplierAddress       string        `json:"supplier_address,omitempty"`         //		否	发货人地址，如不传订单发货信息默认取门店信息
	SupplierPhone         string        `json:"supplier_phone,omitempty"`           //		否	发货人电话，如不传订单发货信息默认取门店信息
	SupplierLat           float64       `json:"supplier_lat,omitempty"`             //否	发货人纬度，如不传订单发货信息默认取门店信息
	SupplierLng           float64       `json:"supplier_lng,omitempty"`             //	否	发货人经度，如不传订单发货信息默认取门店信息
}

type ProductInfo struct {
	SkuName      string  `json:"sku_name"`       //		是	商品名称，限制长度128
	SrcProductNo string  `json:"src_product_no"` //	是	商品编码，限制长度64
	Count        float64 `json:"count"`          //	是	商品数量，精确到小数点后两位
	Unit         string  `json:"unit"`           //	否	商品单位，默认：件
}

type QueryDeliverFeeRes struct {
	Distance     float64         `json:"distance"`     //		是	配送距离(单位：米)
	DeliveryNo   string          `json:"deliveryNo"`   //		是	平台订单号
	Fee          decimal.Decimal `json:"fee"`          //	是	实际运费(单位：元)，运费减去优惠券费用
	DeliverFee   decimal.Decimal `json:"deliver_fee"`  //		是	运费(单位：元)
	CouponFee    decimal.Decimal `json:"couponFee"`    //	否	优惠券费用(单位：元)
	Tips         float64         `json:"tips"`         //	否	小费（单位：元，精确小数点后一位，小费金额不能高于订单金额。）
	InsuranceFee decimal.Decimal `json:"insuranceFee"` //		否	保价费(单位：元)
}

// QueryDeliverFee 查询运费
func (sdk *DadaSdk) QueryDeliverFee(ctx context.Context, req *QueryDeliverFeeReq) (*QueryDeliverFeeRes, error) {
	var res = &QueryDeliverFeeRes{}
	err := sdk.post(ctx, "/api/order/queryDeliverFee", req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type AddAfterQueryReq struct {
	DeliveryNo  string `json:"deliveryNo"`     //	是	平台订单编号
	EnableReset bool   `json:"enableReset"`    //	否	是否重置订单备注等字段。true：重置，false或缺失：不重置
	Info        string `json:"info,omitempty"` //	否	订单备注，长度限制128
}

func (sdk *DadaSdk) AddAfterQuery(ctx context.Context, req *AddAfterQueryReq) error {
	//达达平台询价单号仅在3分钟内有效
	err := sdk.post(ctx, "/api/order/addAfterQuery", req, nil)
	if err != nil {
		return err
	}
	return nil
}

type QueryOrderReq struct {
	OrderId string `json:"order_id"`
}

type QueryOrderRes struct {
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

func (sdk *DadaSdk) QueryOrder(ctx context.Context, req *QueryOrderReq) (*QueryOrderRes, error) {
	var res = &QueryOrderRes{}
	err := sdk.post(ctx, "/api/order/status/query", req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type CancelOrderReq struct {
	OrderId        string `json:"order_id"` //第三方订单编号
	CancelReasonId int64  `json:"cancel_reason_id"`
	CancelReason   string `json:"cancel_reason"`
}

type CancelOrderRes struct {
	DeductFee decimal.Decimal `json:"deduct_fee"`
}

func (sdk *DadaSdk) CancelOrder(ctx context.Context, req *CancelOrderReq) (*CancelOrderRes, error) {
	var res = &CancelOrderRes{}
	err := sdk.post(ctx, "/api/order/formalCancel", req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
