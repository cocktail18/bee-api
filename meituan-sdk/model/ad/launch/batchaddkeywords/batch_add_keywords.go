package batchaddkeywords

/**
* 批量新增关键词定向
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const batch_add_keywords_url = "/ad/launch/batchAddKeywords"

type BatchAddKeywordsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * private List successObjects; private Map failObjects;
    */
    Data BatchOptResult `json:"data"`
    TraceId string `json:"traceId"`
}
type BatchOptResult struct {
}
type BatchAddKeywordsRequest struct {
    /**
    *  premiumName：溢价关键词定向名 keywords：词包列表，用于新增品类词包，地址词，多个词逗号隔开 words:用于新增推荐词，多个词逗号隔开 itemTypes：只在选择门店热搜词的情况下传固定值3 dpBidPrice：点评出价（单位分） mtBidPrice：美团出价（单位分） 
    */
    KeywordPremiumList string `json:"keywordPremiumList"`
}



func (req *BatchAddKeywordsRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*BatchAddKeywordsResponse, error) {
    resp, err := client.InvokeApi(batch_add_keywords_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response BatchAddKeywordsResponse
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

func (response *BatchAddKeywordsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

