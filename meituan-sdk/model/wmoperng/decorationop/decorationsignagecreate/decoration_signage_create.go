package decorationsignagecreate

/**
* 商家开放平台创建招牌
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const decoration_signage_create_url = "/wmoper/ng/decorationop/signageCreate"

type DecorationSignageCreateRequest struct {
    /**
    *  epoiId门店列表，通过英文逗号分隔 
    */
    PoiIds string `json:"poiIds"`
    /**
    *  图片code，服务商可通过接口【上传无水印照片】获取 
    */
    PicCode string `json:"pic_code"`
}
type DecorationSignageCreateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *DecorationSignageCreateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DecorationSignageCreateResponse, error) {
    resp, err := client.InvokeApi(decoration_signage_create_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DecorationSignageCreateResponse
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

func (response *DecorationSignageCreateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

