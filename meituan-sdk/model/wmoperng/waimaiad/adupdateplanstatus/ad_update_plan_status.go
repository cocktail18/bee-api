package adupdateplanstatus

/**
* 修改广告投放计划状态（不能新建）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const ad_update_plan_status_url = "/wmoper/ng/waimaiad/plan/update/status"

type AdUpdatePlanStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type AdUpdatePlanStatusRequest struct {
    /**
    * 外卖广告的产品类型，1表示点金推广
    */
    Channel int32 `json:"channel"`
    /**
    * 计划状态。0：开启，1：关闭
    */
    Status int32 `json:"status"`
}
type Data struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *AdUpdatePlanStatusRequest) DoInvoke(client mtclient.MeituanClient) (*AdUpdatePlanStatusResponse, error) {
    resp, err := client.InvokeApi(ad_update_plan_status_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response AdUpdatePlanStatusResponse
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

func (response *AdUpdatePlanStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

