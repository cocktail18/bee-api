package foodskudelete

/**
* 删除SKU信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const food_sku_delete_url = "/wmoper/ng/foodop/food/sku/delete"

type FoodSkuDeleteResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type FoodSkuDeleteRequest struct {
    /**
    *  服务商方的菜品id， 最大长度128，不同门店可以重复，同一门店内不能重复 
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    *  服务商方的菜品sku_id 
    */
    SkuId string `json:"sku_id"`
}



func (req *FoodSkuDeleteRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FoodSkuDeleteResponse, error) {
    resp, err := client.InvokeApi(food_sku_delete_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response FoodSkuDeleteResponse
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

func (response *FoodSkuDeleteResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

