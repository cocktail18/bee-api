package search

/**
* 获取im消息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const search_url = "/mlive/immsg/search"

type SearchRequest struct {
    Biz string `json:"biz"`
}
type SearchResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *SearchRequest) DoInvoke(client mtclient.MeituanClient) (*SearchResponse, error) {
    resp, err := client.InvokeApi(search_url, 50, "", req)

    if err != nil {
        return nil, err
    }
    var response SearchResponse
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

func (response *SearchResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

