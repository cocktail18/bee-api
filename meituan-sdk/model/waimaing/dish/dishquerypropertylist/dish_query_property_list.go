package dishquerypropertylist

/**
* 查询菜品属性
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const dish_query_property_list_url = "/waimai/dish/queryPropertyList"

type DishQueryPropertyListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 菜品属性数组
    */
    Data []DishPropertyInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type DishPropertyInfo struct {
    /**
    * 菜品属性名
    */
    PropertyName string `json:"propertyName"`
    /**
    * 菜品具体属性
    */
    Values []int64 `json:"values"`
}
type DishQueryPropertyListRequest struct {
    /**
    *  ERP方菜品id 
    */
    EDishCode string `json:"eDishCode"`
}



func (req *DishQueryPropertyListRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DishQueryPropertyListResponse, error) {
    resp, err := client.InvokeApi(dish_query_property_list_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DishQueryPropertyListResponse
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

func (response *DishQueryPropertyListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

