package preparationmealcomplete

/**
* 商家确认已完成出餐
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const preparation_meal_complete_url = "/wmoper/ng/order/preparationMealComplete"

type PreparationMealCompleteResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type PreparationMealCompleteRequest struct {
    /**
    *  订单id 
    */
    OrderId string `json:"orderId"`
}



func (req *PreparationMealCompleteRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*PreparationMealCompleteResponse, error) {
    resp, err := client.InvokeApi(preparation_meal_complete_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response PreparationMealCompleteResponse
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

func (response *PreparationMealCompleteResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

