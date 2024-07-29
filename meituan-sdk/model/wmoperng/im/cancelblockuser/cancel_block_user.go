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

const cancel_block_user_url = "/wmoper/ng/im/cancelBlockUser"

type CancelBlockUserRequest struct {
    /**
    *  用户id 
    */
    OpenUserId string `json:"open_user_id"`
}
type CancelBlockUserResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *CancelBlockUserRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CancelBlockUserResponse, error) {
    resp, err := client.InvokeApi(cancel_block_user_url, 16, appAuthToken, req)

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

