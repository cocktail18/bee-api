package getsystemlabels

/**
* 获取系统标签
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_system_labels_url = "/waimai/ng/valueadded/getSystemLabels"

type GetSystemLabelsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data GetSystemLabelsData `json:"data"`
    TraceId string `json:"traceId"`
}
type LabelList struct {
    /**
    * 系统标签id
    */
    LabelId int64 `json:"labelId"`
    /**
    * 标签名称
    */
    Name string `json:"name"`
    /**
    * 描述
    */
    Desc string `json:"desc"`
}
type GetSystemLabelsData struct {
    /**
    * 系统标签列表
    */
    LabelList []LabelList `json:"labelList"`
}
type GetSystemLabelsRequest struct {
}



func (req *GetSystemLabelsRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetSystemLabelsResponse, error) {
    resp, err := client.InvokeApi(get_system_labels_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetSystemLabelsResponse
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

func (response *GetSystemLabelsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

