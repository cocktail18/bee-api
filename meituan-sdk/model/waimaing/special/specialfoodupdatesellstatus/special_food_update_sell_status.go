package specialfoodupdatesellstatus

/**
* 修改上下架状态
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const special_food_update_sell_status_url = "/waimai/ng/special/food/updateSellStatus"

type SpecialFoodUpdateSellStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}
type SpecialFoodUpdateSellStatusRequest struct {
    /**
    * 门店id 
    */
    EpoiId string `json:"epoiId"`
    /**
    * 业务标识，1-特价版
    */
    BusinessIdentify int32 `json:"businessIdentify"`
    /**
    * 菜品id
    */
    EDishCode string `json:"eDishCode"`
    /**
    * 菜品上下架状态，0表上架，1表下架
    */
    SellStatus int32 `json:"sellStatus"`
}



func (req *SpecialFoodUpdateSellStatusRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SpecialFoodUpdateSellStatusResponse, error) {
    resp, err := client.InvokeApi(special_food_update_sell_status_url, 2, appAuthToken, req)

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

