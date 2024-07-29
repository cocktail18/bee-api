package queryimageforsingle

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

const query_image_for_single_url = "/waimai/ng/decoration/queryImageForSingle"

type QueryImageForSingleResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []ImageForSingleInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type ImageForSingleInfo struct {
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
type QueryImageForSingleRequest struct {
}



func (req *QueryImageForSingleRequest) DoInvoke(client mtclient.MeituanClient) (*QueryImageForSingleResponse, error) {
    resp, err := client.InvokeApi(query_image_for_single_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response QueryImageForSingleResponse
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

func (response *QueryImageForSingleResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

