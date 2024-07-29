package orderconfirm

/**
* 商家确认接单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_confirm_url = "/waimai/order/confirm"

type OrderConfirmRequest struct {
    /**
    *  订单号 
    */
    OrderId int64 `json:"orderId"`
}
type OrderConfirmResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 确认信息
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *OrderConfirmRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderConfirmResponse, error) {
    resp, err := client.InvokeApi(order_confirm_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderConfirmResponse
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

func (response *OrderConfirmResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

