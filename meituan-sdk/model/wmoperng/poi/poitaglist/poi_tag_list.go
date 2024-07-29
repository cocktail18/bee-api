package poitaglist

/**
* 获取门店品类列表
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const poi_tag_list_url = "/wmoper/ng/poi/poiTag/list"

type PoiTagListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    /**
    * 品类ID
    */
    Id int32 `json:"id"`
    /**
    * 品类名称
    */
    Name string `json:"name"`
}
type PoiTagListRequest struct {
}



func (req *PoiTagListRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*PoiTagListResponse, error) {
    resp, err := client.InvokeApi(poi_tag_list_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response PoiTagListResponse
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

func (response *PoiTagListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

