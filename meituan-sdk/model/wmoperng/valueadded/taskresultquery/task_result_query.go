package taskresultquery

/**
* 查询任务结果
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const task_result_query_url = "/wmoper/ng/valueadded/async/taskResultQuery"

type TaskResultQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []TaskResultQueryData `json:"data"`
    TraceId string `json:"traceId"`
}
type TaskResultQueryData struct {
    SendNum int64 `json:"sendNum"`
    SendTime int64 `json:"sendTime"`
    SuccessUser []int64 `json:"successUser"`
    FailUser []FailUser `json:"failUser"`
}
type TaskResultQueryRequest struct {
    /**
    *  枚举 2 根据系统标签获取用户列表 3 根据自定以标签获取用户列表 4 建券并发券 5 手机号发券 
    */
    TaskType int32 `json:"taskType"`
    /**
    *  任务请求id。 支持手机号发券生成的任务id 
    */
    QueryId string `json:"queryId"`
}
type FailUser struct {
    UserId int64 `json:"userId"`
    Reason string `json:"reason"`
}



func (req *TaskResultQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TaskResultQueryResponse, error) {
    resp, err := client.InvokeApi(task_result_query_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TaskResultQueryResponse
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

func (response *TaskResultQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

