package querysuggestbid

/**
* 批量查询广告位次出价
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_suggest_bid_url = "/ad/launch/querySuggestBid"

type QuerySuggestBidResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data SuggestPriceDTO `json:"data"`
    TraceId string `json:"traceId"`
}
type SuggestPriceDTO struct {
}
type QuerySuggestBidRequest struct {
    /**
    *  门店以及对应的投放时段 
    */
    SuggestBid string `json:"suggestBid"`
}



func (req *QuerySuggestBidRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QuerySuggestBidResponse, error) {
    resp, err := client.InvokeApi(query_suggest_bid_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QuerySuggestBidResponse
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

func (response *QuerySuggestBidResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

