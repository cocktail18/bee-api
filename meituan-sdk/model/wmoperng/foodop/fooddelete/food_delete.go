package fooddelete

/**
* 删除菜品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const food_delete_url = "/wmoper/ng/foodop/food/delete"

type FoodDeleteResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type FoodDeleteRequest struct {
    /**
    *  服务商方的菜品id， 最大长度128，不同门店可以重复，同一门店内不能重复 
    */
    AppFoodCode string `json:"app_food_code"`
}



func (req *FoodDeleteRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FoodDeleteResponse, error) {
    resp, err := client.InvokeApi(food_delete_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response FoodDeleteResponse
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

func (response *FoodDeleteResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

