package dishupdatestock

/**
* 更新菜品库存【sku的库存】
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const dish_update_stock_url = "/waimai/dish/updateStock"

type DishUpdateStockResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type DishUpdateStockRequest struct {
    /**
    *  菜品库存列表 参考dishSkuStocks说明 使用json数组格式传递参数 
    */
    DishSkuStocks string `json:"dishSkuStocks"`
    /**
    *  ERP方门店id 最大长度100 
    */
    EPoiId string `json:"ePoiId"`
}



func (req *DishUpdateStockRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishUpdateStockResponse, error) {
    resp, err := client.InvokeApi(dish_update_stock_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishUpdateStockResponse
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

func (response *DishUpdateStockResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

