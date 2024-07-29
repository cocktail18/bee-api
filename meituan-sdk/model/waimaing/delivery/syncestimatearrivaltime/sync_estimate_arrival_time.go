package syncestimatearrivaltime

/**
* 自配订单同步预计送达时间信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const sync_estimate_arrival_time_url = "/waimai/ng/delivery/medicine/syncEstimateArrivalTime"

type SyncEstimateArrivalTimeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 返回成功或失败。 如果非自配送订单，且非医药类订单，则直接返回失败。
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type SyncEstimateArrivalTimeRequest struct {
    /**
    *  订单号（同订单展示ID） 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  订单预计送达时间，为10位秒级时间戳 
    */
    EstimateArrivalTime string `json:"estimateArrivalTime"`
}



func (req *SyncEstimateArrivalTimeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SyncEstimateArrivalTimeResponse, error) {
    resp, err := client.InvokeApi(sync_estimate_arrival_time_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SyncEstimateArrivalTimeResponse
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

func (response *SyncEstimateArrivalTimeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

