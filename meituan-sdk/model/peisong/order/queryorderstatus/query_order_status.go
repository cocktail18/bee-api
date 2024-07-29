package queryorderstatus

/**
* 订单状态查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_order_status_url = "/peisong/order/queryStatus"

type QueryOrderStatusData struct {
    /**
    * 配送活动标识
    */
    DeliveryId int64 `json:"delivery_id"`
    /**
    * 美团配送内部订单id，最长不超过32个字符
    */
    MtPeisongId string `json:"mt_peisong_id"`
    /**
    * 外部订单号，最长不超过32个字符
    */
    OrderId string `json:"order_id"`
    /**
    * 状态代码，可选值为 0：待调度 20：已接单 30：已取货 50：已送达 99：已取消
    */
    Status int32 `json:"status"`
    /**
    * 订单状态变更时间，格式为unix-timestamp
    */
    OperateTime int32 `json:"operate_time"`
    /**
    * 配送员姓名（订单已被骑手接单后会返回骑手信息）
    */
    CourierName string `json:"courier_name"`
    /**
    * 配送员电话（订单已被骑手接单后会返回骑手信息）
    */
    CourierPhone string `json:"courier_phone"`
    /**
    * 取消原因id，详情请参考 美团配送开放平台接口文档--门户页面-4.3，订单取消原因列表
    */
    CancelReasonId int32 `json:"cancel_reason_id"`
    /**
    * 取消原因详情，最长不超过256个字符
    */
    CancelReason string `json:"cancel_reason"`
    /**
    * 非必须，订单预计送达时间
    */
    PredictDeliveryTime string `json:"predict_delivery_time"`
    /**
    * 订单配送距离，单位为米
    */
    DeliveryDistance int32 `json:"delivery_distance"`
    /**
    * 订单配送价格，单位为元
    */
    DeliveryFee float64 `json:"delivery_fee"`
    /**
    * 实际支付价格，单位为元；使用优惠券等扣减后价格；支付方式为账期时支付价格以账单为准
    */
    PayAmount float64 `json:"pay_amount"`
    /**
    * 结算方式，1、实时结算，2、账期结算；账期结算下包含的支付方式为账期；实时结算下包含的支付方式为余额；后续会扩展多种支付方式
    */
    SettlementModeCode int32 `json:"settlement_mode_code"`
    /**
    * 优惠券金额，单位为元； 当前订单存在使用符合条件的优惠券则有值，反之无值或为0
    */
    CouponsAmount float64 `json:"coupons_amount"`
}
type QueryOrderStatusRequest struct {
    /**
    * 配送活动标识
    */
    DeliveryId int64 `json:"delivery_id"`
    /**
    * 美团配送内部订单id，最长不超过32个字符
    */
    MtPeisongId string `json:"mt_peisong_id"`
}
type QueryOrderStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data QueryOrderStatusData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryOrderStatusRequest) DoInvoke(client mtclient.MeituanClient) (*QueryOrderStatusResponse, error) {
    resp, err := client.InvokeApi(query_order_status_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response QueryOrderStatusResponse
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

func (response *QueryOrderStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

