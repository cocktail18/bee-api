package batchupdatecpclaunchstatus

/**
* 更新推广状态
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const batch_update_cpclaunch_status_url = "/ad/launch/batchUpdateCpcLaunchStatus"

type BatchUpdateCpclaunchStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}
type BatchUpdateCpclaunchStatusRequest struct {
    /**
    * 推广id列表，单次最多支持50个
    */
    LaunchIds []int64 `json:"launchIds"`
    /**
    * 需要修改的推广状态，1有效,5暂停,9删除
    */
    Power int32 `json:"power"`
    SonAdaccountId int32 `json:"sonAdaccountId"`
}



func (req *BatchUpdateCpclaunchStatusRequest) DoInvoke(client mtclient.MeituanClient) (*BatchUpdateCpclaunchStatusResponse, error) {
    resp, err := client.InvokeApi(batch_update_cpclaunch_status_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response BatchUpdateCpclaunchStatusResponse
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

func (response *BatchUpdateCpclaunchStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

