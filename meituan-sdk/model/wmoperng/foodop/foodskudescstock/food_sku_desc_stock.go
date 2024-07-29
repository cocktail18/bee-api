package foodskudescstock

/**
* 减少SKU库存
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const food_sku_desc_stock_url = "/wmoper/ng/foodop/food/sku/desc_stock"

type Skus struct {
    /**
    *  sku码 
    */
    SkuId string `json:"sku_id"`
    /**
    *  此处为sku库存的减少值，不能为负数或小数 
    */
    Stock string `json:"stock"`
}
type FoodSkuDescStockRequest struct {
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
type FoodSkuDescStockResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *FoodSkuDescStockRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*FoodSkuDescStockResponse, error) {
    resp, err := client.InvokeApi(food_sku_desc_stock_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response FoodSkuDescStockResponse
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

func (response *FoodSkuDescStockResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

