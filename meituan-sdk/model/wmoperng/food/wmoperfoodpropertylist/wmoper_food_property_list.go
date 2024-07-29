package wmoperfoodpropertylist

/**
* 获取菜品属性
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_food_property_list_url = "/wmoper/ng/food/queryFoodPropertyList"

type WmoperFoodPropertyListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 菜品属性数组
    */
    Data []Data `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperFoodPropertyListRequest struct {
    /**
    *  菜品id 
    */
    AppFoodCode string `json:"app_food_code"`
}
type Data struct {
    /**
    * 菜品属性名
    */
    PropertyName string `json:"property_name"`
    /**
    * 菜品具体属性
    */
    Values []string `json:"values"`
}



func (req *WmoperFoodPropertyListRequest) DoInvoke(client mtclient.MeituanClient) (*WmoperFoodPropertyListResponse, error) {
    resp, err := client.InvokeApi(wmoper_food_property_list_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response WmoperFoodPropertyListResponse
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

func (response *WmoperFoodPropertyListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

