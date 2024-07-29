package updatetablestatus

/**
* 变更桌位状态
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const update_table_status_url = "/resv2/table/status/update"

type UpdateTableStatusRequest struct {
    Data []Data `json:"data"`
}
type Data struct {
    /**
    *  桌位id 
    */
    TableId string `json:"tableId"`
    /**
    *  就餐时间 unix时间戳 
    */
    OrderTime int32 `json:"orderTime"`
    /**
    *  锁台开始时间 unix时间戳 
    */
    StartTime int32 `json:"startTime"`
    /**
    *  锁台结束时间 unix时间戳 
    */
    EndTime int32 `json:"endTime"`
    /**
    *  桌位状态 0：锁台 1：解锁 
    */
    Status int32 `json:"status"`
    /**
    *  餐段，与门店配置中餐段一致 
    */
    TimePeriod int32 `json:"timePeriod"`
    /**
    *  推送时间 当前时间戳 
    */
    ChangeTime int32 `json:"changeTime"`
}
type UpdateTableStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 返回的数据
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *UpdateTableStatusRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*UpdateTableStatusResponse, error) {
    resp, err := client.InvokeApi(update_table_status_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response UpdateTableStatusResponse
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

func (response *UpdateTableStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

