package foodupdateappfoodcodebyorigin

/**
* 根据原商品编码更换新商品编码
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const food_update_app_food_code_by_origin_url = "/waimai/ng/dish/food/updateAppFoodCodeByOrigin"

type FoodUpdateAppFoodCodeByOriginResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type FoodUpdateAppFoodCodeByOriginRequest struct {
    Biz string `json:"biz"`
}



func (req *FoodUpdateAppFoodCodeByOriginRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FoodUpdateAppFoodCodeByOriginResponse, error) {
    resp, err := client.InvokeApi(food_update_app_food_code_by_origin_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response FoodUpdateAppFoodCodeByOriginResponse
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

func (response *FoodUpdateAppFoodCodeByOriginResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

