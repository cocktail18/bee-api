package pagequerytokenpoilist

/**
* 适用门店查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const page_query_token_poi_list_url = "/ddzhkh/auth/token/pageQueryPoiList"

type CustomerPoiInfoTO struct {
    /**
    * 门店id
    */
    OpPoiId string `json:"opPoiId"`
    /**
    * 门店名称
    */
    Name string `json:"name"`
    /**
    * 地址
    */
    Address string `json:"address"`
    /**
    * 城市
    */
    CityName string `json:"cityName"`
}
type PageQueryTokenPoiListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Resp `json:"data"`
    TraceId string `json:"traceId"`
}
type Resp struct {
    /**
    * 总数
    */
    Total int32 `json:"total"`
    PoiInfoList []CustomerPoiInfoTO `json:"poiInfoList"`
}
type PageQueryTokenPoiListRequest struct {
    Limit int32 `json:"limit"`
    Offset int32 `json:"offset"`
}



func (req *PageQueryTokenPoiListRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*PageQueryTokenPoiListResponse, error) {
    resp, err := client.InvokeApi(page_query_token_poi_list_url, 59, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response PageQueryTokenPoiListResponse
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

func (response *PageQueryTokenPoiListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

