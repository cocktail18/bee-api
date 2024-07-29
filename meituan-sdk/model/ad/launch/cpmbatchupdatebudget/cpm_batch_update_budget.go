package cpmbatchupdatebudget

/**
* cpm批量修改预算
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const cpm_batch_update_budget_url = "/ad/launch/cpm/batchUpdateBudget"

type BatchOperateRsp struct {
}
type CpmBatchUpdateBudgetResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * private List successList; private Map failReasonMap; private Map successMap;
    */
    Data BatchOperateRsp `json:"data"`
    TraceId string `json:"traceId"`
}
type BatchOperateRequest struct {
    /**
    *  计划ID 
    */
    PlanId int32 `json:"planId"`
    /**
    *  基础日预算 
    */
    DailyBudget int32 `json:"dailyBudget"`
}
type CpmBatchUpdateBudgetRequest struct {
    /**
    *  大小不超过50 
    */
    BatchOperateRequests []BatchOperateRequest `json:"batchOperateRequests"`
}



func (req *CpmBatchUpdateBudgetRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CpmBatchUpdateBudgetResponse, error) {
    resp, err := client.InvokeApi(cpm_batch_update_budget_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CpmBatchUpdateBudgetResponse
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

func (response *CpmBatchUpdateBudgetResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

