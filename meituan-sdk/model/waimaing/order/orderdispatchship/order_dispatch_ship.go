package orderdispatchship

/**
* 美团专送场景－发配送
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_dispatch_ship_url = "/waimai/order/dispatchShip"

type OrderDispatchShipResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderDispatchShipRequest struct {
    /**
    *  订单号 
    */
    OrderId int64 `json:"orderId"`
}



func (req *OrderDispatchShipRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderDispatchShipResponse, error) {
    resp, err := client.InvokeApi(order_dispatch_ship_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderDispatchShipResponse
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

func (response *OrderDispatchShipResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

