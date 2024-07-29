package addatahistory

/**
* 广告历史数据，区分到天
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const ad_data_history_url = "/wmoper/ng/waimaiad/data/history"

type AdDataInfoList struct {
    Time string `json:"time"`
    Show string `json:"show"`
    Click string `json:"click"`
    Debit string `json:"debit"`
}
type Data struct {
    Result Result `json:"result"`
    AdDataInfoList []AdDataInfoList `json:"adDataInfoList"`
    Show string `json:"show"`
    Click string `json:"click"`
    Debit string `json:"debit"`
}
type AdDataHistoryRequest struct {
    /**
    * 外卖广告产品类型，1表示点金推广
    */
    Channel int32 `json:"channel"`
    /**
    * 查询的历史天数，不能大于30天
    */
    Days int32 `json:"days"`
    /**
    * 查询的数据类型，1：全部数据，2：仅KA系统数据，3：仅单门店数据
    */
    Type int32 `json:"type"`
}
type AdDataHistoryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *AdDataHistoryRequest) DoInvoke(client mtclient.MeituanClient) (*AdDataHistoryResponse, error) {
    resp, err := client.InvokeApi(ad_data_history_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response AdDataHistoryResponse
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

func (response *AdDataHistoryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

