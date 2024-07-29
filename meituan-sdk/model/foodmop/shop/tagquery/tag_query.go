package tagquery

/**
* 品牌门店标签查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const tag_query_url = "/foodmop/shop/tag/query"

type TagQueryRequest struct {
    /**
    *  品牌门店 id 列表 最多 100 个 
    */
    VendorShopIdList []string `json:"vendorShopIdList"`
}
type ResultData struct {
}
type TagQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * key：厂商门店ID value：标签Code列表 如{ "1234" : [ "1", "2"] }
    */
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *TagQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TagQueryResponse, error) {
    resp, err := client.InvokeApi(tag_query_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TagQueryResponse
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

func (response *TagQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

