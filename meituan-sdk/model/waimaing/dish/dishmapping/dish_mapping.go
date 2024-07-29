package dishmapping

/**
* 建立菜品映射
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const dish_mapping_url = "/waimai/dish/mapping"

type DishMappingRequest struct {
    DishMappings string `json:"dishMappings"`
    OrderEntranceType int32 `json:"orderEntranceType"`
}
type DishMappingResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *DishMappingRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishMappingResponse, error) {
    resp, err := client.InvokeApi(dish_mapping_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishMappingResponse
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

func (response *DishMappingResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

