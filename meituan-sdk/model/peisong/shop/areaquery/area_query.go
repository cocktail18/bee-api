package areaquery

/**
* 查询合作方配送范围
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const area_query_url = "/peisong/shop/area/query"

type AreaQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data AreaQueryData `json:"data"`
    TraceId string `json:"traceId"`
}
type AreaQueryData struct {
    /**
    * 门店配送范围
    */
    Scope string `json:"scope"`
}
type AreaQueryRequest struct {
    /**
    * 配送服务代码
    */
    DeliveryServiceCode int32 `json:"delivery_service_code"`
    /**
    * 取货门店id
    */
    ShopId string `json:"shop_id"`
}



func (req *AreaQueryRequest) DoInvoke(client mtclient.MeituanClient) (*AreaQueryResponse, error) {
    resp, err := client.InvokeApi(area_query_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response AreaQueryResponse
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

func (response *AreaQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

