package blockuser

/**
* 商家屏蔽顾客
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const block_user_url = "/wmoper/ng/im/blockUser"

type BlockUserResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}
type BlockUserRequest struct {
    /**
    *  用户id 
    */
    OpenUserId string `json:"open_user_id"`
    /**
    *  屏蔽类型，1：骚扰、辱骂我，言辞恶劣，2：让我加微信又不下单，疑似诈骗，3：其他 
    */
    BlockCode int64 `json:"block_code"`
    /**
    *  屏蔽原因，当屏蔽类型为「其他」时该值必填 
    */
    BlockReasonDesc string `json:"block_reason_desc"`
    /**
    *  照片id，以，分割，最少1张，最多9张，通过图片上传接口获取图片id 
    */
    VoucherPicIds string `json:"voucher_pic_ids"`
}



func (req *BlockUserRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*BlockUserResponse, error) {
    resp, err := client.InvokeApi(block_user_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response BlockUserResponse
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

func (response *BlockUserResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

