package bizordcount

/**
* 经营分析-订单数
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const biz_ord_count_url = "/wmoper/ng/waimaiad/biz/ordcount"

type BizOrdCountResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    Result Result `json:"result"`
    Count int64 `json:"count"`
}
type BizOrdCountRequest struct {
    /**
    * 日期，yyyyMMdd格式
    */
    Dt string `json:"dt"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *BizOrdCountRequest) DoInvoke(client mtclient.MeituanClient) (*BizOrdCountResponse, error) {
    resp, err := client.InvokeApi(biz_ord_count_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response BizOrdCountResponse
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

func (response *BizOrdCountResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

