package decorationbossrecommendcreate

/**
* 商家开放平台创建、修改、删除老板推荐
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const decoration_boss_recommend_create_url = "/wmoper/ng/decorationop/bossRecommendCreate"

type DecorationBossRecommendCreateRequest struct {
    /**
    *  ERP自定义商品ID列表，英文逗号分隔 
    */
    AppFoodCodes string `json:"appFoodCodes"`
    /**
    *  老板推荐是否开启 0：关闭 1：自定义 2：智能推荐 
    */
    State int32 `json:"state"`
}
type DecorationBossRecommendCreateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *DecorationBossRecommendCreateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DecorationBossRecommendCreateResponse, error) {
    resp, err := client.InvokeApi(decoration_boss_recommend_create_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DecorationBossRecommendCreateResponse
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

func (response *DecorationBossRecommendCreateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

