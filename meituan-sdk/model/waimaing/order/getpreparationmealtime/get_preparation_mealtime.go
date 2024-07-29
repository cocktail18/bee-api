package getpreparationmealtime

/**
* 商家获取备餐时间
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_preparation_mealtime_url = "/waimai/ng/order/getPreparationMealTime"

type ResultData struct {
    /**
    * 备餐时长 （单位秒）
    */
    PreparationMealTime int32 `json:"preparationMealTime"`
}
type GetPreparationMealtimeRequest struct {
    /**
    *  订单id 
    */
    OrderId int64 `json:"orderId"`
}
type GetPreparationMealtimeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *GetPreparationMealtimeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetPreparationMealtimeResponse, error) {
    resp, err := client.InvokeApi(get_preparation_mealtime_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetPreparationMealtimeResponse
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

func (response *GetPreparationMealtimeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

