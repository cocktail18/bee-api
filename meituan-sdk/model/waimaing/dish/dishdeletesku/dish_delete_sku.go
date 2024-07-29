package dishdeletesku

/**
* 删除菜品sku
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const dish_delete_sku_url = "/waimai/dish/deleteSku"

type DishDeleteSkuRequest struct {
    /**
    *  ERP方skuId 
    */
    EDishSkuCode string `json:"eDishSkuCode"`
    /**
    *  ERP方菜品id 
    */
    EDishCode string `json:"eDishCode"`
}
type DishDeleteSkuResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *DishDeleteSkuRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishDeleteSkuResponse, error) {
    resp, err := client.InvokeApi(dish_delete_sku_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishDeleteSkuResponse
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

func (response *DishDeleteSkuResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

