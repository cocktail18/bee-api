package decorationbossrecommendquery

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

const decoration_boss_recommend_query_url = "/wmoper/ng/decoration/bossRecommendQuery"

type DecorationBossRecommendQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data DecorationBossRecommendQueryData `json:"data"`
    TraceId string `json:"traceId"`
}
type DecorationBossRecommendQueryData struct {
    /**
    * ERP 商品ID 列表
    */
    AppFoodCodes []string `json:"app_food_codes"`
    /**
    * 老板推荐是否开启 0:关闭 1:自定义 2:智能推荐
    */
    State int32 `json:"state"`
}
type DecorationBossRecommendQueryRequest struct {
}



func (req *DecorationBossRecommendQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DecorationBossRecommendQueryResponse, error) {
    resp, err := client.InvokeApi(decoration_boss_recommend_query_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DecorationBossRecommendQueryResponse
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

func (response *DecorationBossRecommendQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

