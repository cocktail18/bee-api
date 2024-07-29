package batchupdatekeywordsbit

/**
* 更新关键词出价
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const batch_update_keywords_bit_url = "/ad/launch/batchUpdateKeywordBidPrice"

type BatchOptResult struct {
}
type BatchUpdateKeywordsBitResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * private List successObjects; private Map failObjects;
    */
    Data BatchOptResult `json:"data"`
    TraceId string `json:"traceId"`
}
type BatchUpdateKeywordsBitRequest struct {
    /**
    *  launchPremiumId：溢价关键词定向ID dpBidPrice：点评出价（单位分） mtBidPrice：美团出价（单位分） 
    */
    KeywordPremiumPrice string `json:"keywordPremiumPrice"`
}



func (req *BatchUpdateKeywordsBitRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*BatchUpdateKeywordsBitResponse, error) {
    resp, err := client.InvokeApi(batch_update_keywords_bit_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response BatchUpdateKeywordsBitResponse
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

func (response *BatchUpdateKeywordsBitResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

