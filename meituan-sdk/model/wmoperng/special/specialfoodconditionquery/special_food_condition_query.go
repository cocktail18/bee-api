package specialfoodconditionquery

/**
* 查询单个商品(拼好饭)
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const special_food_condition_query_url = "/wmoper/ng/special/food/conditionQuery"

type SpecialFoodConditionQueryRequest struct {
    /**
    *  业务标识，1-拼好饭 
    */
    BusinessIdentify int32 `json:"businessIdentify"`
    /**
    *  三方菜品id，对应外卖侧的app_food_code 
    */
    EDishCode string `json:"eDishCode"`
}
type SpecialFoodConditionQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *SpecialFoodConditionQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SpecialFoodConditionQueryResponse, error) {
    resp, err := client.InvokeApi(special_food_condition_query_url, 16, appAuthToken, req)

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

