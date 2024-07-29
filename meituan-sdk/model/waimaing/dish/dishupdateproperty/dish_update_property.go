package dishupdateproperty

/**
* 批量创建/更新菜品属性
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const dish_update_property_url = "/waimai/dish/updateProperty"

type DishUpdatePropertyResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type DishUpdatePropertyRequest struct {
    DishProperty string `json:"dishProperty"`
}



func (req *DishUpdatePropertyRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishUpdatePropertyResponse, error) {
    resp, err := client.InvokeApi(dish_update_property_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishUpdatePropertyResponse
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

func (response *DishUpdatePropertyResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

