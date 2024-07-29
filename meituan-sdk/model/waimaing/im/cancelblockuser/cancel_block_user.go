package cancelblockuser

/**
* 商家解除屏蔽顾客
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const cancel_block_user_url = "/waimai/ng/im/cancelBlockUser"

type CancelBlockUserRequest struct {
    /**
    *  多个用户id，以,分割， eg: 123456,654321 
    */
    OpenUserId string `json:"open_user_id"`
}
type CancelBlockUserResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *CancelBlockUserRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CancelBlockUserResponse, error) {
    resp, err := client.InvokeApi(cancel_block_user_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CancelBlockUserResponse
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

func (response *CancelBlockUserResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

