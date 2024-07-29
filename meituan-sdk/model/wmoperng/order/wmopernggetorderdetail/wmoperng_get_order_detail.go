package wmopernggetorderdetail

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

const wmoperng_get_order_detail_url = "/wmoper/ng/order/detail"

type Activity struct {
    ActDetailId int64 `json:"act_detail_id"`
    ReduceFee float64 `json:"reduce_fee"`
    MtCharge float64 `json:"mt_charge"`
    PoiCharge float64 `json:"poi_charge"`
    Remark string `json:"remark"`
    Type int64 `json:"type"`
}
type Data struct {
    OrderId int64 `json:"order_id"`
    RecipientAddress string `json:"recipient_address"`
    RecipientPhone string `json:"recipient_phone"`
    RecipientName string `json:"recipient_name"`
    ShippingFee float64 `json:"shipping_fee"`
    Total float64 `json:"total"`
    OriginalPrice float64 `json:"original_price"`
    Caution string `json:"caution"`
    ShipperPhone string `json:"shipper_phone"`
    Status int64 `json:"status"`
    DeliveryTime int64 `json:"delivery_time"`
    PayType int64 `json:"pay_type"`
    PickType int64 `json:"pick_type"`
    DaySeq int64 `json:"day_seq"`
    Detail []Detail `json:"detail"`
    Activity []Activity `json:"activity"`
    OrderSendTime int64 `json:"orderSendTime"`
    OrderConfirmTime int64 `json:"orderConfirmTime"`
    OrderReceiveTime int64 `json:"orderReceiveTime"`
    OrderCancelTime int64 `json:"orderCancelTime"`
    OrderCompletedTime int64 `json:"orderCompletedTime"`
    LogisticsSendTime int64 `json:"logisticsSendTime"`
    LogisticsConfirmTime int64 `json:"logisticsConfirmTime"`
    LogisticsCancelTime int64 `json:"logisticsCancelTime"`
    LogisticsFetchTime int64 `json:"logisticsFetchTime"`
    LogisticsCompletedTime int64 `json:"logisticsCompletedTime"`
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    PoiReceiveDetail string `json:"poiReceiveDetail"`
    UserIdView int64 `json:"userIdView"`
}
type WmoperngGetOrderDetailRequest struct {
    /**
    * 订单id
    */
    OrderId int64 `json:"orderId"`
}
type WmoperngGetOrderDetailResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Detail struct {
    AppFoodCode string `json:"app_food_code"`
    FoodName string `json:"food_name"`
    SkuId string `json:"sku_id"`
    Quantity float64 `json:"quantity"`
    Price float64 `json:"price"`
    BoxNum float64 `json:"box_num"`
    BoxPrice float64 `json:"box_price"`
    Unit string `json:"unit"`
    FoodProperty string `json:"food_property"`
    Spec string `json:"spec"`
    MtSkuId int64 `json:"mt_sku_id"`
    MtSpuId int64 `json:"mt_spu_id"`
    MtTagId int64 `json:"mt_tag_id"`
    TagName string `json:"tag_name"`
}



func (req *WmoperngGetOrderDetailRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperngGetOrderDetailResponse, error) {
    resp, err := client.InvokeApi(wmoperng_get_order_detail_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperngGetOrderDetailResponse
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

func (response *WmoperngGetOrderDetailResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

