package batchupdatekeywords

/**
* 批量更新关键词定向
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const batch_update_keywords_url = "/ad/launch/batchUpdateKeywords"

type BatchOptResult struct {
}
type BatchUpdateKeywordsRequest struct {
    /**
    *  premiumName：溢价关键词定向名 keywords：词包列表，用于新增品类词包，地址词，多个词逗号隔开 words:用于新增推荐词，多个词逗号隔开 itemTypes：只在选择门店热搜词的情况下传固定值3 dpBidPrice：点评出价（单位分） mtBidPrice：美团出价（单位分） 
    */
    KeywordPremiumList string `json:"keywordPremiumList"`
}
type BatchUpdateKeywordsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * private List successObjects; private Map failObjects;
    */
    Data BatchOptResult `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *BatchUpdateKeywordsRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*BatchUpdateKeywordsResponse, error) {
    resp, err := client.InvokeApi(batch_update_keywords_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response BatchUpdateKeywordsResponse
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

func (response *BatchUpdateKeywordsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

