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

const food_update_app_food_code_by_origin_url = "/wmoper/ng/foodop/food/updateAppFoodCodeByOrigin"

type FoodUpdateAppFoodCodeByOriginResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type FoodUpdateAppFoodCodeByOriginRequest struct {
    /**
    *  原app_food_code， 为商品的服务商方的商品id 
    */
    AppFoodCodeOrigin string `json:"app_food_code_origin"`
    /**
    *  新app_food_code， 为商品的服务商方的商品id，不同门店可以重复，同一门店内不能重复，最大长度128 
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    *  原sku_id， 为商品sku的唯一标示,如变更则必须 
    */
    SkuIdOrigin string `json:"sku_id_origin"`
    /**
    *  新sku_id， 为商品sku的唯一标示，如变更则必须 
    */
    SkuId string `json:"sku_id"`
}



func (req *FoodUpdateAppFoodCodeByOriginRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FoodUpdateAppFoodCodeByOriginResponse, error) {
    resp, err := client.InvokeApi(food_update_app_food_code_by_origin_url, 16, appAuthToken, req)

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

