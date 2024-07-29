package batcheditcpcplanbudget

/**
* 修改预算
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const batch_edit_cpc_planbudget_url = "/ad/launch/batchEditCpcPlanBudget"

type BatchEditCpcPlanbudgetData struct {
    /**
    * 成功的推广列表
    */
    SuccessLaunchs []int64 `json:"successLaunchs"`
    /**
    * 创建失败的推广map，key为原因，value为推广列表
    */
    FailLaunchs FailLaunch `json:"failLaunchs"`
}
type BatchEditCpcPlanbudgetResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data BatchEditCpcPlanbudgetData `json:"data"`
    TraceId string `json:"traceId"`
}
type FailLaunch struct {
}
type BatchEditCpcPlanbudgetRequest struct {
    /**
    * 批量修改预算,单次最多支持50个
    */
    Launchs string `json:"launchs"`
    SonAdaccountId int32 `json:"sonAdaccountId"`
}



func (req *BatchEditCpcPlanbudgetRequest) DoInvoke(client mtclient.MeituanClient) (*BatchEditCpcPlanbudgetResponse, error) {
    resp, err := client.InvokeApi(batch_edit_cpc_planbudget_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response BatchEditCpcPlanbudgetResponse
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

func (response *BatchEditCpcPlanbudgetResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

