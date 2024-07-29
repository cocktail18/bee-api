package cpmbatchquerylaunches

/**
* cpm批量查询投放
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const cpm_batch_query_launches_url = "/ad/launch/cpm/batchQueryLaunchInfo"

type CpmBatchQueryLaunchesResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data PageResult `json:"data"`
    TraceId string `json:"traceId"`
}
type T struct {
}
type CpmBatchQueryLaunchesRequest struct {
    /**
    *  推广ID列表 
    */
    LaunchIds []int32 `json:"launchIds"`
    /**
    *  推广状态列表，此参数和launchIds互斥 
    */
    LaunchStatusList []int32 `json:"launchStatusList"`
    /**
    *  分页请求页码，传launchStatusList时可传。默认查第一页，每页100个 
    */
    PageNum int32 `json:"pageNum"`
}
type PageResult struct {
    /**
    * 当前页码
    */
    PageNum int32 `json:"pageNum"`
    /**
    * 是否有下一页
    */
    HaveNextPage bool `json:"haveNextPage"`
    /**
    * 查询结果，T为泛型
    */
    Data []T `json:"data"`
}



func (req *CpmBatchQueryLaunchesRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CpmBatchQueryLaunchesResponse, error) {
    resp, err := client.InvokeApi(cpm_batch_query_launches_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CpmBatchQueryLaunchesResponse
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

func (response *CpmBatchQueryLaunchesResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

