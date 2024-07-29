package foodskusellstatus

/**
* 批量更新售卖状态
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const food_sku_sell_status_url = "/wmoper/ng/foodop/food/sku/sellStatus"

type Skus struct {
    /**
    *  sku码 
    */
    SkuId string `json:"sku_id"`
}
type FoodSkuSellStatusRequest struct {
    /**
    *  售卖状态，1表下架，0表上架 
    */
    SellStatus int32 `json:"sell_status"`
    /**
    *  菜品sku集合 
    */
    FoodData []FoodData `json:"food_data"`
}
type FoodData struct {
    /**
    *  服务商方的菜品id，最大长度128，不同门店可以重复，同一门店内不能重复 
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    *  skus 
    */
    Skus []Skus `json:"skus"`
}
type FoodSkuSellStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *FoodSkuSellStatusRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FoodSkuSellStatusResponse, error) {
    resp, err := client.InvokeApi(food_sku_sell_status_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response FoodSkuSellStatusResponse
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

func (response *FoodSkuSellStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

