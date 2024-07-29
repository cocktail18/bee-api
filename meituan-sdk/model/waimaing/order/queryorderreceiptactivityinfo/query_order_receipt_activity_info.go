package queryorderreceiptactivityinfo

/**
* 查询活动分摊接口
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_order_receipt_activity_info_url = "/waimai/ng/order/getOrderReceiptActivityInfo"

type Activities struct {
    /**
    * 活动名称
    */
    Name string `json:"name"`
    /**
    * 活动类型id
    */
    Type int32 `json:"type"`
    /**
    * 该活动上非商家承担
    */
    MtCharge float64 `json:"mt_charge"`
    /**
    * 该活动上商家承担
    */
    PoiCharge float64 `json:"poi_charge"`
}
type ShippingFeeActivities struct {
    /**
    * 活动名称
    */
    Name string `json:"name"`
    /**
    * 活动类型id
    */
    Type int32 `json:"type"`
    /**
    * 该活动上非商家承担
    */
    MtCharge float64 `json:"mt_charge"`
    /**
    * 该活动上商家承担
    */
    PoiCharge float64 `json:"poi_charge"`
}
type QueryOrderReceiptActivityInfoData struct {
    /**
    * 商品优惠分摊信息
    */
    OrderDetailActInfoList []OrderDetailActInfoList `json:"orderDetailActInfoList"`
    /**
    * 原始配送费
    */
    OriginalShippingFee float64 `json:"original_shipping_fee"`
    /**
    * 配送费活动详情
    */
    ShippingFeeActivities []ShippingFeeActivities `json:"shipping_fee_activities"`
}
type QueryOrderReceiptActivityInfoRequest struct {
    Biz string `json:"biz"`
}
type QueryOrderReceiptActivityInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data QueryOrderReceiptActivityInfoData `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderDetailActInfoList struct {
    /**
    * APP方菜品id
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    * sku编码
    */
    SkuId string `json:"sku_id"`
    /**
    * sku总数量
    */
    SkuNum int32 `json:"sku_num"`
    /**
    * 商品原价
    */
    OriginPrice float64 `json:"origin_price"`
    /**
    * sku优惠前总金额
    */
    TotalOriginPrice float64 `json:"total_origin_price"`
    /**
    * sku优惠后总金额
    */
    TotalActivityPrice float64 `json:"total_activity_price"`
    /**
    * sku总优惠金额
    */
    TotalReducePrice float64 `json:"total_reduce_price"`
    /**
    * sku优惠后平均金额
    */
    ActivityPrice float64 `json:"activity_price"`
    /**
    * 该sku非商家承担活动总金额
    */
    TotalMtCharge float64 `json:"total_mt_charge"`
    /**
    * 该sku商家承担活动总金额
    */
    TotalPoiCharge float64 `json:"total_poi_charge"`
    /**
    * 该sku总餐盒费
    */
    TotalBoxPrice float64 `json:"total_box_price"`
    /**
    * 该sku单商品餐盒费
    */
    BoxPrice float64 `json:"box_price"`
    /**
    * 该sku单商品的餐盒数
    */
    BoxNum float64 `json:"box_num"`
    /**
    * 该sku参与的所有活动情况
    */
    Activities []Activities `json:"activities"`
}



func (req *QueryOrderReceiptActivityInfoRequest) DoInvoke(client mtclient.MeituanClient) (*QueryOrderReceiptActivityInfoResponse, error) {
    resp, err := client.InvokeApi(query_order_receipt_activity_info_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response QueryOrderReceiptActivityInfoResponse
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

func (response *QueryOrderReceiptActivityInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

