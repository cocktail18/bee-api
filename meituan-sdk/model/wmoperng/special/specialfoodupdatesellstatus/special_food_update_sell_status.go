package specialfoodupdatesellstatus

/**
* 修改商品上下架状态(拼好饭)
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const special_food_update_sell_status_url = "/wmoper/ng/special/food/updateSellStatus"

type SpecialFoodUpdateSellStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}
type SpecialFoodUpdateSellStatusRequest struct {
    /**
    *  业务标识，1-拼好饭 
    */
    BusinessIdentify int32 `json:"businessIdentify"`
    /**
    *  三方菜品id，当eDishCode为default时，表示清空该值，当eDishCode为空字符串时，表示不处理该字段，对应外卖侧的app_food_code 
    */
    EDishCode string `json:"eDishCode"`
    /**
    *  菜品上下架状态，0表上架，1表下架 
    */
    SellStatus int32 `json:"sellStatus"`
}



func (req *SpecialFoodUpdateSellStatusRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SpecialFoodUpdateSellStatusResponse, error) {
    resp, err := client.InvokeApi(special_food_update_sell_status_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SpecialFoodUpdateSellStatusResponse
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

func (response *SpecialFoodUpdateSellStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

