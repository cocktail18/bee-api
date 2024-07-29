package bossrecommendcreate

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

const boss_recommend_create_url = "/waimai/ng/decoration/bossRecommendCreate"

type BossRecommendCreateRequest struct {
    /**
    *  ERP自定义商品ID列表，英文逗号分隔 
    */
    AppFoodCodes string `json:"appFoodCodes"`
    /**
    *  老板推荐是否开启 0：关闭 1：自定义 2：智能推荐 
    */
    State int32 `json:"state"`
}
type BossRecommendCreateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *BossRecommendCreateRequest) DoInvoke(client mtclient.MeituanClient) (*BossRecommendCreateResponse, error) {
    resp, err := client.InvokeApi(boss_recommend_create_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response BossRecommendCreateResponse
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

func (response *BossRecommendCreateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

