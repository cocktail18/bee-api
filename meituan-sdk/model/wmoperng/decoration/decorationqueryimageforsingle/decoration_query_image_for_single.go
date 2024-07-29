package decorationqueryimageforsingle

/**
* 商家开放平台查询招牌
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const decoration_query_image_for_single_url = "/wmoper/ng/decoration/queryImageForSingle"

type DecorationQueryImageForSingleData struct {
    /**
    * 图片地址
    */
    Url string `json:"url"`
    /**
    * 审核状态： 0：免审 1：审核中 2：审核通过 3：审核驳回
    */
    Status int32 `json:"status"`
    /**
    * 驳回code
    */
    StatusCode int32 `json:"statusCode"`
}
type DecorationQueryImageForSingleRequest struct {
}
type DecorationQueryImageForSingleResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []DecorationQueryImageForSingleData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *DecorationQueryImageForSingleRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DecorationQueryImageForSingleResponse, error) {
    resp, err := client.InvokeApi(decoration_query_image_for_single_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DecorationQueryImageForSingleResponse
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

func (response *DecorationQueryImageForSingleResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

