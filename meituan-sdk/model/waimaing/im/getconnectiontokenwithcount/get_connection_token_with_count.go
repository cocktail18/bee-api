package getconnectiontokenwithcount

/**
* 获取多个长连接
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_connection_token_with_count_url = "/waimai/ng/im/getConnectionTokenWithCount"

type GetConnectionTokenWithCountRequest struct {
    /**
    *  token数量，推荐3，最多为5 
    */
    RequiredTokenCount int32 `json:"required_token_count"`
}
type GetConnectionTokenWithCountResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data GetConnectionTokenWithCountData `json:"data"`
    TraceId string `json:"traceId"`
}
type GetConnectionTokenWithCountData struct {
    /**
    * 建立长连接的token
    */
    TokenInfos []TokenInfos `json:"token_infos"`
    Appkey int64 `json:"appkey"`
    /**
    * 消息加解密密钥
    */
    ImSecret string `json:"imSecret"`
}
type TokenInfos struct {
    /**
    * 建立长连接的token
    */
    ConnectionToken string `json:"connectionToken"`
    /**
    * X分钟内，IM两方会话消息发送失败的用户数
    */
    UserSendFailCount int64 `json:"userSendFailCount"`
    /**
    * X分钟内，评价联系消息发送失败的评价数
    */
    CommentSendFailCount string `json:"commentSendFailCount"`
}



func (req *GetConnectionTokenWithCountRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetConnectionTokenWithCountResponse, error) {
    resp, err := client.InvokeApi(get_connection_token_with_count_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetConnectionTokenWithCountResponse
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

func (response *GetConnectionTokenWithCountResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

