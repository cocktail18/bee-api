package orderdelivered

/**
* 自配送场景－订单已送达
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_delivered_url = "/waimai/order/delivered"

type OrderDeliveredResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderDeliveredRequest struct {
    /**
    *  订单号 
    */
    OrderId int64 `json:"orderId"`
}



func (req *OrderDeliveredRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderDeliveredResponse, error) {
    resp, err := client.InvokeApi(order_delivered_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderDeliveredResponse
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

func (response *OrderDeliveredResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

