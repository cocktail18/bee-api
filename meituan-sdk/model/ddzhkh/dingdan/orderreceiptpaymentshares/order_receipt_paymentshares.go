package orderreceiptpaymentshares

/**
* 订单券码分摊金额查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_receipt_paymentshares_url = "/ddzhkh/dingdan/receipt/paymentshares"

type OrderReceiptPaymentsharesResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    Result []Result `json:"result"`
}
type OrderReceiptPaymentsharesRequest struct {
    /**
    *  订单id，24位数字 
    */
    OrderId string `json:"orderId"`
    /**
    *  订单类型，0-团购，1-次卡 
    */
    Type int32 `json:"type"`
}
type Result struct {
    /**
    * 券码
    */
    ReceiptCode string `json:"receiptCode"`
    /**
    * 分摊金额
    */
    Amount float64 `json:"amount"`
    /**
    * 分摊金额类型，-1-验证金额（用户实付+平台优惠），2-抵用券，5-积分，6-立减，8-商户抵用券，10-C端美团支付，12-优惠代码，15-美团立减，17-商家立减，18-美团商家立减，21-次卡，22-打折卡，23-B端美团支付，24-全渠道会员券，25-pos支付，26-线下认款平台，28-商家折上折
    */
    AmountType int32 `json:"amountType"`
}



func (req *OrderReceiptPaymentsharesRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderReceiptPaymentsharesResponse, error) {
    resp, err := client.InvokeApi(order_receipt_paymentshares_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderReceiptPaymentsharesResponse
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

func (response *OrderReceiptPaymentsharesResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

