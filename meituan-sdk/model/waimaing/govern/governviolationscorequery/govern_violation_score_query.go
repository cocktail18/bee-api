package governviolationscorequery

/**
* 积分信息查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const govern_violation_score_query_url = "/waimai/ng/govern/violation/score/query"

type GovernViolationScoreQueryRequest struct {
}
type GovernViolationScoreQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ReaultData `json:"data"`
    TraceId string `json:"traceId"`
}
type ReaultData struct {
    /**
    * 门店Id
    */
    AppPoiCode string `json:"appPoiCode"`
    /**
    * 一般违规积分
    */
    NormalScore int32 `json:"normalScore"`
    /**
    * 严重违规积分
    */
    CriticalScore int32 `json:"criticalScore"`
}



func (req *GovernViolationScoreQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GovernViolationScoreQueryResponse, error) {
    resp, err := client.InvokeApi(govern_violation_score_query_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GovernViolationScoreQueryResponse
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

func (response *GovernViolationScoreQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

