package orderqueryzbshippingfee

/**
* 众包配送场景－查询配送费
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_query_zb_shipping_fee_url = "/waimai/order/queryZbShippingFee"

type OrderQueryZbShippingFeeRequest struct {
    /**
    *  订单号,多个订单号逗号隔开 
    */
    OrderIds string `json:"orderIds"`
}
type OrderQueryZbShippingFeeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []ZbShippingFeeInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type ZbShippingFeeInfo struct {
    /**
    * 订单号
    */
    OrderId string `json:"orderId"`
    /**
    * 用户可见订单号
    */
    OrderViewId string `json:"orderViewId"`
    /**
    * 配送费
    */
    ShippingFee string `json:"shippingFee"`
    /**
    * 配送费说明 
    */
    ShippingTip string `json:"shippingTip"`
}



func (req *OrderQueryZbShippingFeeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderQueryZbShippingFeeResponse, error) {
    resp, err := client.InvokeApi(order_query_zb_shipping_fee_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderQueryZbShippingFeeResponse
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

func (response *OrderQueryZbShippingFeeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

