package adupdateplanbudget

/**
* 修改广告投放计划预算
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const ad_update_plan_budget_url = "/wmoper/ng/waimaiad/plan/update/budget"

type AdUpdatePlanBudgetRequest struct {
    /**
    * 外卖广告的产品类型，1表示点金推广
    */
    Channel int32 `json:"channel"`
    /**
    * 预算，单位分
    */
    Budget int32 `json:"budget"`
}
type Data struct {
    Success bool `json:"success"`
    Code string `json:"code"`
}
type AdUpdatePlanBudgetResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *AdUpdatePlanBudgetRequest) DoInvoke(client mtclient.MeituanClient) (*AdUpdatePlanBudgetResponse, error) {
    resp, err := client.InvokeApi(ad_update_plan_budget_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response AdUpdatePlanBudgetResponse
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

func (response *AdUpdatePlanBudgetResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

