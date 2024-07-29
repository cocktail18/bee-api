package signagecreate

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

const signage_create_url = "/waimai/ng/decoration/signageCreate"

type SignageCreateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type SignageCreateRequest struct {
    /**
    *  epoiId门店列表，通过英文逗号分隔 
    */
    PoiIds string `json:"poiIds"`
    /**
    *  图片code，服务商可通过接口【上传无水印照片】获取 
    */
    PicCode string `json:"picCode"`
}



func (req *SignageCreateRequest) DoInvoke(client mtclient.MeituanClient) (*SignageCreateResponse, error) {
    resp, err := client.InvokeApi(signage_create_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response SignageCreateResponse
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

func (response *SignageCreateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

