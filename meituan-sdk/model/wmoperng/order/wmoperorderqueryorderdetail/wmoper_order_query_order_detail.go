package wmoperorderqueryorderdetail

/**
* 查询订单详情
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_order_query_order_detail_url = "/wmoper/order/queryOrderDetail"

type WmoperOrderQueryOrderDetailData struct {
    /**
    * 活动信息
    */
    Activity []Activity `json:"activity"`
    /**
    * 备注
    */
    Caution string `json:"caution"`
    /**
    * 流水单号
    */
    DaySeq int32 `json:"daySeq"`
    /**
    * 用户预计送达时间
    */
    DeliveryTime int64 `json:"deliveryTime"`
    /**
    * 菜品信息
    */
    Detail []Detail `json:"detail"`
    /**
    * erp门店id
    */
    EpoiId string `json:"epoiId"`
    /**
    * 实际送餐地址纬度
    */
    Latitude float64 `json:"latitude"`
    /**
    * 配送单取消时间
    */
    LogisticsCancelTime int64 `json:"logisticsCancelTime"`
    /**
    * 配送方式
    */
    LogisticsCode string `json:"logisticsCode"`
    /**
    * 配送完成时间
    */
    LogisticsCompletedTime int64 `json:"logisticsCompletedTime"`
    /**
    * 配送单确认时间
    */
    LogisticsConfirmTime int64 `json:"logisticsConfirmTime"`
    /**
    * 骑手取单时间
    */
    LogisticsFetchTime int64 `json:"logisticsFetchTime"`
    /**
    * 配送单下单时间
    */
    LogisticsSendTime int64 `json:"logisticsSendTime"`
    /**
    * 实际送餐地址经度
    */
    Longitude float64 `json:"longitude"`
    /**
    * 订单取消时间
    */
    OrderCancelTime int64 `json:"orderCancelTime"`
    /**
    * 订单完成时间
    */
    OrderCompletedTime int64 `json:"orderCompletedTime"`
    /**
    * 商户确认时间
    */
    OrderConfirmTime int64 `json:"orderConfirmTime"`
    /**
    * 订单ID（订单展示Id。用户下单时看到的订单号）
    */
    OrderId int64 `json:"orderId"`
    /**
    * 商户收到时间
    */
    OrderReceiveTime int64 `json:"orderReceiveTime"`
    /**
    * 用户下单时间
    */
    OrderSendTime int64 `json:"orderSendTime"`
    /**
    * 商品原价
    */
    OriginalPrice float64 `json:"originalPrice"`
    /**
    * 支付类型，1表货到付款，2表在线支付
    */
    PayType int32 `json:"payType"`
    /**
    * 取餐类型（0：普通取餐；1：到店取餐）
    */
    PickType int32 `json:"pickType"`
    /**
    * 门店ID
    */
    PoiId int64 `json:"poiId"`
    /**
    * 商家对账信息，此字段为JSON字符串
    */
    PoiReceiveDetail string `json:"poi_receive_detail"`
    /**
    * 收货地址
    */
    RecipientAddress string `json:"recipientAddress"`
    /**
    * 收货人名字
    */
    RecipientName string `json:"recipientName"`
    /**
    * 收货人手机号
    */
    RecipientPhone string `json:"recipientPhone"`
    /**
    * 配送员电话
    */
    ShipperPhone string `json:"shipperPhone"`
    /**
    * 门店配送费
    */
    ShippingFee float64 `json:"shippingFee"`
    /**
    * 订单状态 2 新订单，4 已接单，8 订单完成，9 订单取消
    */
    Status int32 `json:"status"`
    /**
    * 订单金额
    */
    Total float64 `json:"total"`
    /**
    * 透传给第三方的偏移后的user_id
    */
    UserIdView int64 `json:"userIdView"`
}
type WmoperOrderQueryOrderDetailRequest struct {
    /**
    *  订单ID（订单展示Id。用户下单时看到的订单号） 
    */
    OrderId int64 `json:"orderId"`
}
type Activity struct {
    /**
    * 活动ID
    */
    ActDetailId int32 `json:"act_detail_id"`
    /**
    * 优惠金额中美团承担的部分
    */
    MtCharge float64 `json:"mt_charge"`
    /**
    * 优惠金额中商家承担的部分
    */
    PoiCharge float64 `json:"poi_charge"`
    /**
    * 该活动用户实际享受优惠金额
    */
    ReduceFee float64 `json:"reduce_fee"`
    /**
    * 优惠说明
    */
    Remark string `json:"remark"`
    /**
    * 活动类型
    */
    Type int32 `json:"type"`
}
type WmoperOrderQueryOrderDetailResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data WmoperOrderQueryOrderDetailData `json:"data"`
    TraceId string `json:"traceId"`
}
type Detail struct {
    /**
    * 菜品code
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    * 餐盒数量
    */
    BoxNum float64 `json:"box_num"`
    /**
    * 餐盒价格
    */
    BoxPrice float64 `json:"box_price"`
    /**
    * 菜品名称
    */
    FoodName string `json:"food_name"`
    /**
    * 菜品属性
    */
    FoodProperty string `json:"food_property"`
    /**
    * 透传给第三方的偏移后的sku的id
    */
    MtSkuId int64 `json:"mt_sku_id"`
    /**
    * 透传给第三方的偏移后的spu的id
    */
    MtSpuId int64 `json:"mt_spu_id"`
    /**
    * 透传给第三方的偏移后的tag_id
    */
    MtTagId int64 `json:"mt_tag_id"`
    /**
    * 商品单价
    */
    Price float64 `json:"price"`
    /**
    * 商品数量
    */
    Quantity float64 `json:"quantity"`
    /**
    * sku编码
    */
    SkuId string `json:"sku_id"`
    /**
    * 菜品规格名称
    */
    Spec string `json:"spec"`
    /**
    * 透传给第三方的tag_name
    */
    TagName string `json:"tag_name"`
    /**
    * 单位
    */
    Unit string `json:"unit"`
}



func (req *WmoperOrderQueryOrderDetailRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperOrderQueryOrderDetailResponse, error) {
    resp, err := client.InvokeApi(wmoper_order_query_order_detail_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperOrderQueryOrderDetailResponse
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

func (response *WmoperOrderQueryOrderDetailResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

