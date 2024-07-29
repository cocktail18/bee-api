package specialfoodupdatestock

/**
* 修改库存
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const special_food_update_stock_url = "/waimai/ng/special/food/updateStock"

type SpecialFoodUpdateStockResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}
type SpecialFoodUpdateStockRequest struct {
    /**
    * 门店id 
    */
    EpoiId string `json:"epoiId"`
    /**
    * 业务标识，1-特价版
    */
    BusinessIdentify int32 `json:"businessIdentify"`
    /**
    * 菜品规格id
    */
    SkuId string `json:"skuId"`
    /**
    * 菜品id
    */
    EDishCode string `json:"eDishCode"`
    /**
    * 当前库存值
    */
    Stock string `json:"stock"`
}



func (req *SpecialFoodUpdateStockRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SpecialFoodUpdateStockResponse, error) {
    resp, err := client.InvokeApi(special_food_update_stock_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SpecialFoodUpdateStockResponse
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

func (response *SpecialFoodUpdateStockResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

