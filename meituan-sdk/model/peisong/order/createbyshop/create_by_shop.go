package createbyshop

/**
* 订单创建（门店方式）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const create_by_shop_url = "/peisong/order/createByShop"

type CreateByShopRequest struct {
    /**
    *  即配送活动标识，由外部系统生成，不同order_id应对应不同的delivery_id，若因美团系统故障导致发单接口失败，可利用相同的delivery_id重新发单，系统视为同一次配送活动，若更换delivery_id，则系统视为两次独立配送活动。 请关注并不超过字段类型要求、长度以及最大值等。 
    */
    DeliveryId int64 `json:"delivery_id"`
    /**
    * 订单id，即该订单在合作方系统中的id，最长不超过32个字符 注：目前若某一订单正在配送中（状态不为取消），再次发送同一订单（order_id相同）将返回同一mt_peisong_id
    */
    OrderId string `json:"order_id"`
    /**
    * 订单来源： 101.美团（外卖&闪购） 102.饿了么 103.京东到家 201.商家web网站 202.商家小程序-微信 203.商家小程序-支付宝 204.商家APP 205.商家热线 其他，请直接填写中文字符串，最长不超过20个字符 非「其他」需传代码
    */
    OuterOrderSourceDesc string `json:"outer_order_source_desc"`
    /**
    * 原平台订单号，如订单来源为美团，该字段必填，且为美团平台生成的订单号，最长不超过32个字符
    */
    OuterOrderSourceNo string `json:"outer_order_source_no"`
    /**
    * 取货门店id，即合作方向美团提供的门店id 注：测试门店的shop_id固定为 test_0001 ，仅用于对接时联调测试。
    */
    ShopId string `json:"shop_id"`
    /**
    * 配送服务代码，详情见合同 1）服务包 飞速达:4002 快速达:4011 及时达:4012 集中送:4013 跑腿B帮送：4031 2）新服务产品 具体可参考 新服务产品列表
    */
    DeliveryServiceCode int32 `json:"delivery_service_code"`
    /**
    * 收件人名称，最长不超过256个字符
    */
    ReceiverName string `json:"receiver_name"`
    /**
    * 收件人地址，最长不超过512个字符
    */
    ReceiverAddress string `json:"receiver_address"`
    /**
    *  收件人电话，最长不超过64个字符；收件人首要联系方式，只支持传输一个真实手机号码或隐私号；传输隐私号的分机号时，请保持英文”,“隔开；若收件人手机号是真实号码，可联系运营开通隐私号服务；不支持带*号的手机号码 
    */
    ReceiverPhone string `json:"receiver_phone"`
    /**
    *  备用手机号，最长不超过64字符；收件人其他的真实手机号或有效隐私号，可支持至多2个备用手机号，需要通过英文“;”区分；传输隐私号格式请保持分机号英文”,“隔开（如：135XXXXXXX,XXXX） 
    */
    ReceiverPhoneSpare string `json:"receiver_phone_spare"`
    /**
    * 收件人经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 116398419
    */
    ReceiverLng int32 `json:"receiver_lng"`
    /**
    * 收件人纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 39985005
    */
    ReceiverLat int32 `json:"receiver_lat"`
    /**
    * 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标 （默认值为0）
    */
    CoordinateType int32 `json:"coordinate_type"`
    /**
    * 货物价格，单位为元，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为0-5000;跑腿B（4031）帮送服务包：范围0-3000
    */
    GoodsValue float64 `json:"goods_value"`
    /**
    * 货物高度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为0-45
    */
    GoodsHeight float64 `json:"goods_height"`
    /**
    * 货物宽度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为0-50
    */
    GoodsWidth float64 `json:"goods_width"`
    /**
    * 货物长度，单位为cm，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为0-65
    */
    GoodsLength float64 `json:"goods_length"`
    /**
    * 货物重量，单位为kg，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为0-50;跑腿B（4031）帮送服务包：范围0-20
    */
    GoodsWeight float64 `json:"goods_weight"`
    /**
    *  货物详情，最长不超过10240个字符，内容为JSON格式，示例如下： { "goods":[ { "goodCount": 1, "goodName": "货品名称", "goodPrice": 9.99, "goodUnit": "个", "goodUnitCode":"10008" } ] } goods对应货物列表； goodCount表示货物数量，int类型，必填且必须大于0； goodName表示货品名称，String类型，必填且不能为空； goodPrice表示货品单价，double类型，选填，数值不小于0，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数）； goodUnit表示货品单位，String类型，选填，最长不超过20个字符； goodUnitCode表示货品标准计量单位，String类型，有特殊计费模式商家请必传；若字段【goodUnit】与当前字段同时传值不一致时，则以当前货品标准计量单位为准；参照附录- 货品标准计量单位列表 。 强烈建议提供，方便骑手在取货时确认货品信息 
    */
    GoodsDetail string `json:"goods_detail"`
    /**
    * 货物取货信息，用于骑手到店取货，最长不超过100个字符
    */
    GoodsPickupInfo string `json:"goods_pickup_info"`
    /**
    * 货物交付信息，最长不超过100个字符
    */
    GoodsDeliveryInfo string `json:"goods_delivery_info"`
    /**
    * 期望取货时间，时区为GMT+8，当前距离Epoch（1970年1月1日) 以秒计算的时间，即unix-timestamp。跑腿B（4031）帮送服务包预约单：发单后60分钟到次日23:50，10的倍数 跑腿B（4031）帮送服务包即时单：order type=0时 expected_pickup_time字段不允许传任何值,返回错误 文案：立即取件 无需输入取件时间。非跑腿B帮送服务包：非必填；跑腿B帮送服务包-即时单：非必填；跑腿B帮送服务包-预约单：必填
    */
    ExpectedPickupTime int64 `json:"expected_pickup_time"`
    /**
    * 期望送达时间，时区为GMT+8，当前距离Epoch（1970年1月1日) 以秒计算的时间，即unix-timestamp 即时单：以发单时间 + 服务包时效作为期望送达时间（当天送服务包需客户指定期望送达时间） 预约单：以客户传参数据为准（预约时间必须大于当前下单时间+服务包时效+3分钟）。即时单期望送达时间规则:1.飞速达(4002) 发单时间+45mins 2.快速达(4011) 发单时间+1h 3.及时达(4012) 发单时间+2h 4.集中送(4013) 发单时间+2h。跑腿B（4031）帮送服务包传此值不生效。跑腿B帮送服务包时，此字段为非必填。
    */
    ExpectedDeliveryTime int64 `json:"expected_delivery_time"`
    /**
    * 订单类型，默认为0 0: 即时单(尽快送达，限当日订单) 1: 预约单
    */
    OrderType int32 `json:"order_type"`
    /**
    * 门店订单流水号，建议提供，方便骑手门店取货，最长不超过32个字符
    */
    PoiSeq string `json:"poi_seq"`
    /**
    * 备注，最长不超过200个字符。
    */
    Note string `json:"note"`
    /**
    * 骑手应付金额，单位为元，精确到分【预留字段】
    */
    CashOnDelivery float64 `json:"cash_on_delivery"`
    /**
    * 骑手应收金额，单位为元，精确到分【预留字段】
    */
    CashOnPickup float64 `json:"cash_on_pickup"`
    /**
    * 发票抬头，最长不超过256个字符【预留字段】
    */
    InvoiceTitle string `json:"invoice_title"`
    /**
    * 加小费字段，加小费，精确到元，需要为1或1的倍数，上限20，允许每个运单最多加10次； 非跑腿B帮送服务包不支持
    */
    TipAmount int32 `json:"tip_amount"`
    /**
    * 收货码开关，默认为0，0=off，1=on，当收件人手机号为隐私号时，不支持收货码； 非跑腿B帮送服务包不支持
    */
    GoodsCodeSwitch int32 `json:"goods_code_switch"`
    /**
    *  支付方式，0、账期支付，1、余额支付； 当前门店可支持的支付方式可以通过查询门店信息接口获取；选择余额支付时，建单成功后若账户余额扣款失败，则美团系统自动会发起订单取消；建单成功后会返回支付方式对应的付款方式。 
    */
    PayTypeCode int32 `json:"pay_type_code"`
    /**
    *  优惠券ID； 当前是否有可用优惠券可以通过配送能力校验接口获取；若传值，则会校验优惠券ID在当前时刻是否符合使用规则，若不传值，则不会使用优惠券。 
    */
    CouponsId string `json:"coupons_id"`
}
type CreateByShopData struct {
    /**
    * 美团配送内部订单id
    */
    MtPeisongId string `json:"mt_peisong_id"`
    /**
    * 配送活动标识
    */
    DeliveryId int64 `json:"delivery_id"`
    /**
    * 外部订单id
    */
    OrderId string `json:"order_id"`
    /**
    * 订单配送距离，单位为米，当前订单计算的配送距离
    */
    DeliveryDistance int32 `json:"delivery_distance"`
    /**
    * 订单配送价格，单位为元，价格受距离、发单时间、配送高峰、特殊天气等不同计费因素影响
    */
    DeliveryFee float64 `json:"delivery_fee"`
    /**
    * 优惠可用金额，单位为元；使用余额支付时有值；总配送费仲可使用优惠券抵扣的部分配送费用，金额受到不同计费因素以及优惠券使用规则影响
    */
    BaseDeliveryFee float64 `json:"base_delivery_fee"`
    /**
    * 实际支付价格，单位为元；使用优惠券等扣减后价格；支付方式为账期时支付价格以账单为准
    */
    PayAmount float64 `json:"pay_amount"`
    /**
    * 结算方式，1、实时结算，2、账期结算；账期结算下包含的支付方式为账期；实时结算下包含的支付方式为余额；后续会扩展多种支付方式
    */
    SettlementModeCode int32 `json:"settlement_mode_code"`
    /**
    * 优惠券金额，单位为元； 当前订单创建时刻传入的优惠券且符合使用条件时则有值，反之无值或为0。
    */
    CouponsAmount float64 `json:"coupons_amount"`
}
type CreateByShopResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CreateByShopData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *CreateByShopRequest) DoInvoke(client mtclient.MeituanClient) (*CreateByShopResponse, error) {
    resp, err := client.InvokeApi(create_by_shop_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response CreateByShopResponse
    err = json.Unmarshal(resp, &response)
    if err != nil {
        return nil, err
    }

    if response.IsSuccess() {
        return &response, nil
    } else {
        return &response, &api_error.ApiError{Code: response.Code, Msg: response.Msg, TraceId: response.TraceId}
    }
}

func (response *CreateByShopResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

