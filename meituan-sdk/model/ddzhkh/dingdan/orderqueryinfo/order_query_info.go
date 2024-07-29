package orderqueryinfo

/**
* 订单及券码状态查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_query_info_url = "/ddzhkh/dingdan/query/info"

type Order struct {
    /**
    * 统一订单Id
    */
    UnifiedOrderId string `json:"unifiedOrderId"`
    /**
    * 订单Id，用户订单详情页实际看到的订单Id
    */
    OrderId int64 `json:"orderId"`
    /**
    * 商品快照结构化信息
    */
    ProductSnapshotContentStruct string `json:"productSnapshotContentStruct"`
    /**
    * 订单状态，1-初始状态，处于下单流程中，2-创建成功，订单创建成功，等待用户支付，3-订单取消，订单超时未支付，4-订单终止，订单发券失败，支付金额会退款给用户，5-订单完成，订单发券成功
    */
    Status int32 `json:"status"`
}
type Receipt struct {
    /**
    * 券码号
    */
    ReceiptCode string `json:"receiptCode"`
    /**
    * 券状态， 1-未使用、未过期，2-已使用，3-未使用、已过期，4-已退款 （未消费退款），6-已退款 （已消费退）
    */
    Status int32 `json:"status"`
    /**
    * 团购券可用开始时间，时间戳
    */
    ReceiptBeginTime int64 `json:"receiptBeginTime"`
    /**
    * 团购券可用截止时间，时间戳
    */
    ReceiptEndTime int64 `json:"receiptEndTime"`
    /**
    * 金额分摊信息 注：类型为对象类型，详见扩展参数
    */
    AmountShareList []AmountShare `json:"amountShareList"`
}
type AmountShare struct {
    /**
    * 资金类型，2-抵用券，5-积分 6-立减，8-商户抵用券，10-C端美团支付，12-优惠代码，15-美团立减，17-商家立减，18-美团商家立减，21-次卡，22-打折卡，23-B端美团支付，24-全渠道会员券，25-pos支付，26-线下认款平台，28-商家折上折，29-美团分销支付
    */
    AmountType int64 `json:"amountType"`
    /**
    * 优惠名称
    */
    Title string `json:"title"`
    /**
    * 优惠金额(单位分)
    */
    Value int64 `json:"value"`
}
type OrderInfoResponseDTO struct {
    Order Order `json:"order"`
    ReceiptList []Receipt `json:"receiptList"`
}
type OrderQueryInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data OrderInfoResponseDTO `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderQueryInfoRequest struct {
    /**
    *  订单ID（用户订单详情页实际看到的订单号） orderId 与 unifiedOrderId 必须二选一，同时传参时以unifiedOrderId 为准，以 orderId 查询订单信息 不支持持非团购 业务 
    */
    OrderId string `json:"orderId"`
    /**
    *  需要查询的信息，1-订单，2-券码 
    */
    Type []int32 `json:"type"`
    /**
    *  统一订单ID（开放平台对外透出的订单号） orderId 与 unifiedOrderId 必须二选一，同时传参时以unifiedOrderId 为准，以 unifiedOrderId 查询订单信息 支持非团购 业务 
    */
    UnifiedOrderId string `json:"unifiedOrderId"`
}



func (req *OrderQueryInfoRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderQueryInfoResponse, error) {
    resp, err := client.InvokeApi(order_query_info_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderQueryInfoResponse
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

func (response *OrderQueryInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

