package updateposterstatus

/**
* 商家开放平台使用海报
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const update_poster_status_url = "/waimai/ng/decoration/updatePosterStatus"

type UpdatePosterStatusRequest struct {
    /**
    *  海报id 
    */
    PosterId int64 `json:"posterId"`
    /**
    *  展示状态 0：不展示 1：展示 
    */
    Valid int32 `json:"valid"`
}
type UpdatePosterStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *UpdatePosterStatusRequest) DoInvoke(client mtclient.MeituanClient) (*UpdatePosterStatusResponse, error) {
    resp, err := client.InvokeApi(update_poster_status_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response UpdatePosterStatusResponse
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

func (response *UpdatePosterStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

