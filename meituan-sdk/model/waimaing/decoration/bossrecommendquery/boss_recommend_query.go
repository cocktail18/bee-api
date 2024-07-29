package bossrecommendquery

/**
* 商家开放平台查询老板推荐
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const boss_recommend_query_url = "/waimai/ng/decoration/bossRecommendQuery"

type BossRecommendQueryRequest struct {
}
type BossRecommendQueryData struct {
    /**
    * 老板推荐是否开启 0:关闭 1:自定义 2:智能推荐
    */
    State int32 `json:"state"`
}
type BossRecommendQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data BossRecommendQueryData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *BossRecommendQueryRequest) DoInvoke(client mtclient.MeituanClient) (*BossRecommendQueryResponse, error) {
    resp, err := client.InvokeApi(boss_recommend_query_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response BossRecommendQueryResponse
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

func (response *BossRecommendQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

