package orderbatchfetchabnormalorder

/**
* 批量拉取异常订单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_batch_fetch_abnormal_order_url = "/waimai/ng/order/batchFetchAbnormalOrder"

type OrderBatchFetchAbnormalOrderResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 异常订单id列表
    */
    Data []int64 `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderBatchFetchAbnormalOrderRequest struct {
    Biz string `json:"biz"`
}



func (req *OrderBatchFetchAbnormalOrderRequest) DoInvoke(client mtclient.MeituanClient) (*OrderBatchFetchAbnormalOrderResponse, error) {
    resp, err := client.InvokeApi(order_batch_fetch_abnormal_order_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response OrderBatchFetchAbnormalOrderResponse
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

func (response *OrderBatchFetchAbnormalOrderResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

