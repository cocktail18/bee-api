package orderdelivering

/**
* 自配送－配送状态
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_delivering_url = "/waimai/order/delivering"

type OrderDeliveringResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderDeliveringRequest struct {
    /**
    *  配送人名称 
    */
    CourierName string `json:"courierName"`
    /**
    *  订单号 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  配送人电话 
    */
    CourierPhone string `json:"courierPhone"`
}



func (req *OrderDeliveringRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderDeliveringResponse, error) {
    resp, err := client.InvokeApi(order_delivering_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderDeliveringResponse
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

func (response *OrderDeliveringResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

