package foodskuincstock

/**
* 增加SKU库存
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const food_sku_inc_stock_url = "/wmoper/ng/foodop/food/sku/inc_stock"

type Skus struct {
    /**
    *  sku码 
    */
    SkuId string `json:"sku_id"`
    /**
    *  此处为sku库存的增加值，不能为负数或小数 
    */
    Stock string `json:"stock"`
}
type FoodSkuIncStockResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type FoodSkuIncStockRequest struct {
    /**
    *  菜品sku价格集合 
    */
    FoodData []FoodData `json:"food_data"`
}
type FoodData struct {
    /**
    *  服务商方的菜品id 最大长度128，不同门店可以重复，同一门店内不能重复 
    */
    AppFoodCode string `json:"app_food_code"`
    /**
    *  skus 
    */
    Skus []Skus `json:"skus"`
}



func (req *FoodSkuIncStockRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FoodSkuIncStockResponse, error) {
    resp, err := client.InvokeApi(food_sku_inc_stock_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response FoodSkuIncStockResponse
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

func (response *FoodSkuIncStockResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

