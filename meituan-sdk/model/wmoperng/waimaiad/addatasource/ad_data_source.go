package addatasource

/**
* 获取不同位置的效果数据
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const ad_data_source_url = "/wmoper/ng/waimaiad/data/source"

type SearchPageList struct {
    Pos int64 `json:"pos"`
    Click string `json:"click"`
    Debit string `json:"debit"`
    AvgPrice int64 `json:"avgPrice"`
    DebitPercent float64 `json:"debitPercent"`
}
type Data struct {
    Result Result `json:"result"`
    HomePageList []HomePageList `json:"homePageList"`
    ChannelPageList []ChannelPageList `json:"channelPageList"`
    SearchPageList []SearchPageList `json:"searchPageList"`
}
type HomePageList struct {
    Pos int64 `json:"pos"`
    Click string `json:"click"`
    Debit string `json:"debit"`
    AvgPrice int64 `json:"avgPrice"`
    DebitPercent float64 `json:"debitPercent"`
}
type AdDataSourceRequest struct {
    /**
    * 外卖广告产品类型，1表示点金推广
    */
    Channel int32 `json:"channel"`
    /**
    * 查询的历史天数，不能大于30天
    */
    Days int32 `json:"days"`
    /**
    * 查询的日期，不能早于T-31
    */
    Date int32 `json:"date"`
}
type AdDataSourceResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type ChannelPageList struct {
    Pos int64 `json:"pos"`
    Click string `json:"click"`
    Debit string `json:"debit"`
    AvgPrice int64 `json:"avgPrice"`
    DebitPercent float64 `json:"debitPercent"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *AdDataSourceRequest) DoInvoke(client mtclient.MeituanClient) (*AdDataSourceResponse, error) {
    resp, err := client.InvokeApi(ad_data_source_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response AdDataSourceResponse
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

func (response *AdDataSourceResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

