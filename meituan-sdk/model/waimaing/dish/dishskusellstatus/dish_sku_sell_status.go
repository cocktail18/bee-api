package dishskusellstatus

/**
* 批量更新菜品售卖状态
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const dish_sku_sell_status_url = "/waimai/ng/dish/sku/sellStatus"

type Skus struct {
    /**
    *  sku码 
    */
    SkuId string `json:"sku_id"`
}
type DishSkuSellStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
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
type DishSkuSellStatusRequest struct {
    /**
    *  售卖状态，1表下架，0表上架 
    */
    SellStatus int32 `json:"sell_status"`
    /**
    *  菜品sku集合 
    */
    FoodData []FoodData `json:"food_data"`
}



func (req *DishSkuSellStatusRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishSkuSellStatusResponse, error) {
    resp, err := client.InvokeApi(dish_sku_sell_status_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishSkuSellStatusResponse
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

func (response *DishSkuSellStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

