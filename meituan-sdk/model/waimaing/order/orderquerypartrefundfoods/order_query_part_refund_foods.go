package orderquerypartrefundfoods

/**
* 查询部分退款商品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_query_part_refund_foods_url = "/waimai/order/queryPartRefundFoods"

type OrderQueryPartRefundFoodsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderQueryPartRefundFoodsRequest struct {
    /**
    *  订单号 
    */
    OrderId int64 `json:"orderId"`
}



func (req *OrderQueryPartRefundFoodsRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderQueryPartRefundFoodsResponse, error) {
    resp, err := client.InvokeApi(order_query_part_refund_foods_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderQueryPartRefundFoodsResponse
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

func (response *OrderQueryPartRefundFoodsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

