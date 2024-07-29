package setautosendconfig

/**
* 自动回复设置
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const set_auto_send_config_url = "/waimai/ng/im/setAutoSendConfig"

type SetAutoSendConfigResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type SetAutoSendConfigRequest struct {
    /**
    *  本次只有18，邀请顾客评价 
    */
    AutoSendType int32 `json:"auto_send_type"`
    /**
    *  当auto_send_type为18时，传0 UPDATE_STATUS(0), UPDATE_CONTENT(1), UPDATE_CONTENT_ONLY(2); 
    */
    OpType int32 `json:"op_type"`
    /**
    *  权益开关状态 0 关闭 1 打开; op_type为0、1时必填; 
    */
    Status int32 `json:"status"`
    /**
    *  op_type为1、2时必填 
    */
    Content string `json:"content"`
}



func (req *SetAutoSendConfigRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SetAutoSendConfigResponse, error) {
    resp, err := client.InvokeApi(set_auto_send_config_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SetAutoSendConfigResponse
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

func (response *SetAutoSendConfigResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

