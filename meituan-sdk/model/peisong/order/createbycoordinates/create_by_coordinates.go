package createbycoordinates

/**
* 订单创建(送货分拣方式)
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const create_by_coordinates_url = "/peisong/order/createByCoordinates"

type CreateByCoordinatesData struct {
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
    * 路区信息
    */
    RoadArea string `json:"road_area"`
    /**
    * 目的地信息
    */
    DestinationId string `json:"destination_id"`
    /**
    * 订单配送距离，单位为米，当前订单计算的配送距离
    */
    DeliveryDistance int32 `json:"delivery_distance"`
    /**
    * 订单配送价格，单位为元，价格受距离、发单时间、配送高峰、特殊天气等不同计费因素影响
    */
    DeliveryFee float64 `json:"delivery_fee"`
}
type CreateByCoordinatesResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CreateByCoordinatesData `json:"data"`
    TraceId string `json:"traceId"`
}
type CreateByCoordinatesRequest struct {
    /**
    * 即配送活动标识，由外部系统生成，不同order_id应对应 不同的delivery_id，若因美团系统故障导致发单接口失 败，可利用相同的delivery_id重新发单，系统视为同一次 配送活动，若更换delivery_id，则系统视为两次独立配送 活动。
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
    * 配送服务代码，详情见合同 飞速达:4002 快速达:4011 及时达:4012 集中送:4013
    */
    DeliveryServiceCode int32 `json:"delivery_service_code"`
    /**
    * 货物取货类型，目前只支持1 1：客户配送至站点 2：美团自取
    */
    PickUpType int32 `json:"pick_up_type"`
    /**
    * 收件人名称，最长不超过256个字符
    */
    ReceiverName string `json:"receiver_name"`
    /**
    * 收件人电话，最长不超过64个字符
    */
    ReceiverPhone string `json:"receiver_phone"`
    /**
    * 收件人地址省，如上海市，江苏省，最长不超过10个字符
    */
    ReceiverProvince string `json:"receiver_province"`
    /**
    * 收件人地址市，如上海市，南京市，最长不超过10个字符
    */
    ReceiverCity string `json:"receiver_city"`
    /**
    * 收件人地址区县，如静安区，最长不超过10个字符
    */
    ReceiverCountry string `json:"receiver_country"`
    /**
    * 收件人地址街道，如宝山路街道，最长不超过30个字符
    */
    ReceiverTown string `json:"receiver_town"`
    /**
    * 收件人地址详情，如永兴路365号，最长不超过100个字符
    */
    ReceiverDetailAddress string `json:"receiver_detail_address"`
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
    * 货物价格，单位为元，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为0-5000
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
    * 货物重量，单位为kg，精确到小数点后两位（如果小数点后位数多于两位，则四舍五入保留两位小数），范围为0-50
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
    * 期望取货时间，时区为GMT+8，当前距离Epoch（1970年1月1日) 以秒计算的时间，即unix-timestamp。
    */
    ExpectedPickupTime int64 `json:"expected_pickup_time"`
    /**
    * 期望送达时间，时区为GMT+8，当前距离Epoch（1970年1月1日) 以秒计算的时间，即unix-timestamp。
    */
    ExpectedDeliveryTime int64 `json:"expected_delivery_time"`
    /**
    * 订单类型，目前只支持预约单 0: 及时单(尽快送达，限当日订单) 1: 预约单
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
}



func (req *CreateByCoordinatesRequest) DoInvoke(client mtclient.MeituanClient) (*CreateByCoordinatesResponse, error) {
    resp, err := client.InvokeApi(create_by_coordinates_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response CreateByCoordinatesResponse
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

func (response *CreateByCoordinatesResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

