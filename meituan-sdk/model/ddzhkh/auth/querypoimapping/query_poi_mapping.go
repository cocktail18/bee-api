package querypoimapping

/**
* 对应客户门店ID映射关系
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_poi_mapping_url = "/ddzhkh/auth/token/queryPoiMapping"

type QueryPoiMappingResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type PoiInfoTO struct {
    /**
    * 美团门店id
    */
    PoiId int64 `json:"poiId"`
    /**
    * 门店混淆id
    */
    OpPoiId string `json:"opPoiId"`
    /**
    * 错误信息
    */
    ErrMsg string `json:"errMsg"`
}
type QueryPoiMappingRequest struct {
    /**
    *  列表长度大于0，小于等于100 
    */
    PoiIds []int64 `json:"poiIds"`
}
type Data struct {
    Result []PoiInfoTO `json:"result"`
}



func (req *QueryPoiMappingRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryPoiMappingResponse, error) {
    resp, err := client.InvokeApi(query_poi_mapping_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryPoiMappingResponse
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

func (response *QueryPoiMappingResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

