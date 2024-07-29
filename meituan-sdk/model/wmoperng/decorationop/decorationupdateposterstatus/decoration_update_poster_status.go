package decorationupdateposterstatus

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

const decoration_update_poster_status_url = "/wmoper/ng/decorationop/updatePosterStatus"

type DecorationUpdatePosterStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}
type DecorationUpdatePosterStatusRequest struct {
    /**
    *  海报id 
    */
    PosterId int64 `json:"posterId"`
    /**
    *  展示状态 0：不展示 1：展示 
    */
    Valid int32 `json:"valid"`
}



func (req *DecorationUpdatePosterStatusRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*DecorationUpdatePosterStatusResponse, error) {
    resp, err := client.InvokeApi(decoration_update_poster_status_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response DecorationUpdatePosterStatusResponse
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

func (response *DecorationUpdatePosterStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

