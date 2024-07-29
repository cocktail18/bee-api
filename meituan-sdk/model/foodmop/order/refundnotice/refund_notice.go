package refundnotice

/**
* 品牌方对美团发起的退款审核确认
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const refund_notice_url = "/foodmop/order/refundNotice"

type RefundNoticeRequest struct {
    /**
    *  MT订单id 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  退款流水，申请退款时会传给品牌 
    */
    RefundNumber string `json:"refundNumber"`
    /**
    *  请求时间戳 
    */
    Timestap int64 `json:"timestap"`
    /**
    *  商户订单号， 商户查不到订单时为空 
    */
    VendorOrderId string `json:"vendorOrderId"`
    /**
    *  退款申请确认结果： AGREE(0, "同意") REJECT(1, "拒绝") 
    */
    OrderRefundStatus int32 `json:"orderRefundStatus"`
    /**
    *  拒绝退款原因： BEGIN_MAKE(1, "已开始制作")， ORDER_NOT_FUND(2, "未查询到订单")， 
    */
    VendorRejectRefundType int32 `json:"vendorRejectRefundType"`
    /**
    *  操作者Id 
    */
    OperatorId string `json:"operatorId"`
    /**
    *  操作者名称 
    */
    OperatorName string `json:"operatorName"`
    /**
    *  失败的详细信息 
    */
    Message string `json:"message"`
    /**
    *  处理时间 
    */
    HandleTime int64 `json:"handleTime"`
    /**
    *  业务类型 10-MOP 
    */
    BizType int32 `json:"bizType"`
}
type ResultData struct {
}
type RefundNoticeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *RefundNoticeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*RefundNoticeResponse, error) {
    resp, err := client.InvokeApi(refund_notice_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response RefundNoticeResponse
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

func (response *RefundNoticeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

