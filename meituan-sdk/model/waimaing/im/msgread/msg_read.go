package msgread

/**
* 设置设置消息已读状态
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const msg_read_url = "/waimai/ng/im/msgRead"

type MsgReadResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type MsgReadRequest struct {
    Biz string `json:"biz"`
}



func (req *MsgReadRequest) DoInvoke(client mtclient.MeituanClient) (*MsgReadResponse, error) {
    resp, err := client.InvokeApi(msg_read_url, 2, "", req)

    if err != nil {
        return nil, err
    }
    var response MsgReadResponse
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

func (response *MsgReadResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

