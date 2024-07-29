package tagbind

/**
* 品牌门店打标签
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const tag_bind_url = "/foodmop/shop/tag/bind"

type BindFailResult struct {
    /**
    * SHOP_OFFLINE：门店不在线 TAG_ERR：标签错误
    */
    FailCode string `json:"failCode"`
    /**
    * 如果failCode = TAG_ERR必传： key(String)：标签code value(String)： TAG_AUDIT_REJECT：标签审核失败 TAG_IS_NULL：标签不存在 
    */
    TagFailMap TagFail `json:"tagFailMap"`
}
type ResultData struct {
    /**
    * key(String)：xbk门店id value(BindFailResult): 失败结果封装
    */
    Map DataMap `json:"map"`
}
type TagFail struct {
}
type TagBindRequest struct {
    /**
    *  Map&lt;String, List&lt;String&gt;&gt; key：厂商门店ID value：标签Code列表 {"1234" : [ "1", "2"]} 
    */
    Map DataMap `json:"map"`
}
type TagBindResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type DataMap struct {
    /**
    * 失败结果封装
    */
    BindFailResult BindFailResult `json:"BindFailResult"`
}



func (req *TagBindRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TagBindResponse, error) {
    resp, err := client.InvokeApi(tag_bind_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TagBindResponse
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

func (response *TagBindResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

