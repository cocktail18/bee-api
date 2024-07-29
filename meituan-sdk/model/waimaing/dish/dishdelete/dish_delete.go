package dishdelete

/**
* 删除菜品
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const dish_delete_url = "/waimai/dish/delete"

type DishDeleteResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type DishDeleteRequest struct {
    /**
    *  ERP方菜品id 
    */
    EDishCode string `json:"eDishCode"`
}



func (req *DishDeleteRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishDeleteResponse, error) {
    resp, err := client.InvokeApi(dish_delete_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishDeleteResponse
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

func (response *DishDeleteResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

