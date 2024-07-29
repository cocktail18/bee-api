package cpmbatchupdatelaunchpower

/**
* cpm批量修改投放状态
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const cpm_batch_update_launch_power_url = "/ad/launch/cpm/batchUpdate"

type BatchOperateRsp struct {
}
type CpmBatchUpdateLaunchPowerRequest struct {
    /**
    *  一批不超过50 
    */
    BatchOperateRequests []BatchOperateRequest `json:"batchOperateRequests"`
}
type BatchOperateRequest struct {
    /**
    *  投放ID 
    */
    LaunchId int32 `json:"launchId"`
    /**
    *  用户操作态，1-有效，5-暂停，9-删除 
    */
    Power int32 `json:"power"`
}
type CpmBatchUpdateLaunchPowerResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * private List successList; private Map failReasonMap; private Map successMap;
    */
    Data BatchOperateRsp `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *CpmBatchUpdateLaunchPowerRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CpmBatchUpdateLaunchPowerResponse, error) {
    resp, err := client.InvokeApi(cpm_batch_update_launch_power_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CpmBatchUpdateLaunchPowerResponse
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

func (response *CpmBatchUpdateLaunchPowerResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

