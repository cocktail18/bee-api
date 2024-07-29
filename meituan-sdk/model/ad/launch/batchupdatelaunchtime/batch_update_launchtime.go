package batchupdatelaunchtime

/**
* 批量修改投放时间段
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const batch_update_launchtime_url = "/ad/launch/batchUpdateLaunchTime"

type BatchOptResult struct {
}
type BatchUpdateLaunchtimeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * private List successObjects; private Map failObjects;
    */
    Data BatchOptResult `json:"data"`
    TraceId string `json:"traceId"`
}
type BatchUpdateLaunchtimeRequest struct {
    /**
    *  
    */
    LaunchDateList string `json:"launchDateList"`
}



func (req *BatchUpdateLaunchtimeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*BatchUpdateLaunchtimeResponse, error) {
    resp, err := client.InvokeApi(batch_update_launchtime_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response BatchUpdateLaunchtimeResponse
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

func (response *BatchUpdateLaunchtimeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

