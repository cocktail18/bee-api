package getconnectiontoken

/**
* 获取长连接的token
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_connection_token_url = "/wmoper/ng/im/getConnectionToken"

type GetConnectionTokenRequest struct {
}
type GetConnectionTokenData struct {
    /**
    * 建立长连接的token
    */
    ConnectionToken string `json:"connectionToken"`
    /**
    * 30分钟内，消息发送失败的用户数
    */
    UserCount int64 `json:"userCount"`
    /**
    * 长连接key，用于拼装websocket url
    */
    ConnectionId string `json:"connectionId"`
    /**
    * im消息加解秘钥
    */
    ImSecret string `json:"imSecret"`
}
type GetConnectionTokenResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data GetConnectionTokenData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *GetConnectionTokenRequest) DoInvoke(client mtclient.MeituanClient) (*GetConnectionTokenResponse, error) {
    resp, err := client.InvokeApi(get_connection_token_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response GetConnectionTokenResponse
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

func (response *GetConnectionTokenResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

