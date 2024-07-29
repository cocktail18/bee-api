package dishupdateprice

/**
* 更新菜品价格【sku的价格】
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const dish_update_price_url = "/waimai/dish/updatePrice"

type DishUpdatePriceRequest struct {
    /**
    *  菜品sku价格 
    */
    DishSkuPrices string `json:"dishSkuPrices"`
    /**
    *  ERP方门店id 最大长度100 
    */
    EPoiId string `json:"ePoiId"`
}
type DishUpdatePriceResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *DishUpdatePriceRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishUpdatePriceResponse, error) {
    resp, err := client.InvokeApi(dish_update_price_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishUpdatePriceResponse
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

func (response *DishUpdatePriceResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

