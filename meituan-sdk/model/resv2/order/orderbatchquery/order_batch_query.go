package orderbatchquery

/**
* 订单批量查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_batch_query_url = "/resv2/order/batchQuery"

type OrderBatchQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderBatchQueryRequest struct {
    PageNo int32 `json:"pageNo"`
    PageSize int32 `json:"pageSize"`
    BookingTimeGe int64 `json:"bookingTimeGe"`
    BookingTimeLe int64 `json:"bookingTimeLe"`
    UpdateTimeGe int64 `json:"updateTimeGe"`
    UpdateTimeLe int64 `json:"updateTimeLe"`
}



func (req *OrderBatchQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderBatchQueryResponse, error) {
    resp, err := client.InvokeApi(order_batch_query_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderBatchQueryResponse
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

func (response *OrderBatchQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

