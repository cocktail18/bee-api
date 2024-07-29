package addatadetail

/**
* 广告明细数据，区分小时
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const ad_data_detail_url = "/wmoper/ng/waimaiad/data/detail"

type AdDataInfoList struct {
    Time string `json:"time"`
    Show string `json:"show"`
    Click string `json:"click"`
    Debit string `json:"debit"`
}
type AdDataDetailRequest struct {
    /**
    * 外卖广告产品类型，1表示点金推广
    */
    Channel int32 `json:"channel"`
    /**
    * 日期
    */
    Date string `json:"date"`
}
type Data struct {
    Result Result `json:"result"`
    AdDataInfoList []AdDataInfoList `json:"adDataInfoList"`
    Show string `json:"show"`
    Click string `json:"click"`
    Debit string `json:"debit"`
}
type AdDataDetailResponse struct {
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



func (req *AdDataDetailRequest) DoInvoke(client mtclient.MeituanClient) (*AdDataDetailResponse, error) {
    resp, err := client.InvokeApi(ad_data_detail_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response AdDataDetailResponse
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

func (response *AdDataDetailResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

