package specialfoodconditionquery

/**
* 单个查询商品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const special_food_condition_query_url = "/waimai/ng/special/food/conditionQuery"

type SpecialFoodConditionQueryRequest struct {
    /**
    * 门店id 
    */
    EpoiId string `json:"epoiId"`
    /**
    * 业务标识，1-特价版
    */
    BusinessIdentify int32 `json:"businessIdentify"`
    /**
    * 菜品id
    */
    EDishCode string `json:"eDishCode"`
}
type SpecialFoodConditionQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *SpecialFoodConditionQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SpecialFoodConditionQueryResponse, error) {
    resp, err := client.InvokeApi(special_food_condition_query_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SpecialFoodConditionQueryResponse
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

func (response *SpecialFoodConditionQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

