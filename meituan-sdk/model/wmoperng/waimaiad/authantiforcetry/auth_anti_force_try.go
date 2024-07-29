package authantiforcetry

/**
* 增值平台接口校验
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const auth_anti_force_try_url = "/wmoper/ng/waimaiad/marketplace/app/auth"

type Data struct {
    Result Result `json:"result"`
    Status int64 `json:"status"`
}
type AuthAntiForceTryRequest struct {
    DeveloperId string `json:"developerId"`
    OpBizCode string `json:"opBizCode"`
    BizToken string `json:"bizToken"`
}
type AuthAntiForceTryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *AuthAntiForceTryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*AuthAntiForceTryResponse, error) {
    resp, err := client.InvokeApi(auth_anti_force_try_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response AuthAntiForceTryResponse
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

func (response *AuthAntiForceTryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

