package queryorderpaymentinfo

/**
* 查询支付结果
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const queryorderpaymentinfo_url = "/ddzh/yuding/queryorderpaymentinfo"

type Order struct {
    /**
    * 统一订单号
    */
    UnifiedOrderId string `json:"unifiedOrderId"`
}
type OrderInfoResponse struct {
    /**
    * 订单信息
    */
    Order Order `json:"order"`
    /**
    * 支付单信息
    */
    OrderPaymentList []OrderPayment `json:"orderPaymentList"`
}
type QueryorderpaymentinfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data OrderInfoResponse `json:"data"`
    TraceId string `json:"traceId"`
}
type QueryorderpaymentinfoRequest struct {
    /**
    *  美团统一订单号 
    */
    UnifiedOrderId string `json:"unifiedOrderId"`
}
type OrderPayment struct {
    /**
    * 支付单总金额
    */
    TotalAmount string `json:"totalAmount"`
    /**
    * 支付批次
    */
    PayBatchNo int32 `json:"payBatchNo"`
    /**
    * 支付成功时间（取不到则返回null）
    */
    SuccessTime int64 `json:"successTime"`
    /**
    * 预支付时，三方传递的自定义支付参数，查询时回传
    */
    ThirdPartyPaymentInfo string `json:"thirdPartyPaymentInfo"`
    /**
    * 支付单状态，1-待支付，2-已取消，3-支付成功，4支付失败
    */
    Status int32 `json:"status"`
}



func (req *QueryorderpaymentinfoRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryorderpaymentinfoResponse, error) {
    resp, err := client.InvokeApi(queryorderpaymentinfo_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryorderpaymentinfoResponse
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

func (response *QueryorderpaymentinfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

