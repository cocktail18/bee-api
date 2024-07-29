package cpmbatchupdatebid

/**
* cpm批量更新出价
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const cpm_batch_update_bid_url = "/ad/launch/cpm/batchUpdateBid"

type CpmBatchUpdateBidResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * private List successList; private Map failReasonMap; private Map successMap;
    */
    Data BatchOperateRsp `json:"data"`
    TraceId string `json:"traceId"`
}
type BatchOperateRsp struct {
}
type BatchOperateRequest struct {
    /**
    *  投放ID 
    */
    LaunchId int32 `json:"launchId"`
    /**
    *  出价模式，当前只支持2-千次曝光出价，4-ocpx点击出价 
    */
    PriceMode int32 `json:"priceMode"`
    /**
    *  出价，千次曝光出价时为整数，点击出价时保留两位小数 
    */
    Bid float64 `json:"bid"`
    /**
    *  点击出价时，须在(0,100)范围内 
    */
    FloatRatio int32 `json:"floatRatio"`
}
type CpmBatchUpdateBidRequest struct {
    /**
    *  一批不超过50 
    */
    BatchOperateRequests []BatchOperateRequest `json:"batchOperateRequests"`
}



func (req *CpmBatchUpdateBidRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CpmBatchUpdateBidResponse, error) {
    resp, err := client.InvokeApi(cpm_batch_update_bid_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CpmBatchUpdateBidResponse
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

func (response *CpmBatchUpdateBidResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

