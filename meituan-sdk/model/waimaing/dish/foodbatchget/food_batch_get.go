package foodbatchget

/**
* 批量查询门店菜品（包括查询套餐商品和普通商品）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const food_batch_get_url = "/waimai/ng/dish/food/batchGet"

type FoodBatchGetResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data RespData `json:"data"`
    TraceId string `json:"traceId"`
}
type RespData struct {
}
type FoodBatchGetRequest struct {
    /**
    *  商品第三方code列表, 不超过200个 
    */
    AppFoodCodes string `json:"app_food_codes"`
}



func (req *FoodBatchGetRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FoodBatchGetResponse, error) {
    resp, err := client.InvokeApi(food_batch_get_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response FoodBatchGetResponse
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

func (response *FoodBatchGetResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

