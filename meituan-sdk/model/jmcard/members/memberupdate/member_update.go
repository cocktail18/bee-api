package memberupdate

/**
* 更新会员信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const member_update_url = "/jmcard/members/update"

type MemberUpdateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Resp `json:"data"`
    TraceId string `json:"traceId"`
}
type Resp struct {
    /**
    * 业务状态
    */
    Status string `json:"status"`
    /**
    * 业务异常时的报错信息
    */
    ErrMsg string `json:"errMsg"`
}
type MemberUpdateRequest struct {
    /**
    *  商户唯一标识，领卡事件中的userOpenId 
    */
    UserOpenId string `json:"userOpenId"`
    /**
    *  该用户使用的卡模板key，应该能对应到一张唯一的卡模板，否则报错 
    */
    DefaultTemplateKey string `json:"defaultTemplateKey"`
    /**
    *  会员卡更新内容 
    */
    ExternalData ExternalData `json:"externalData"`
}
type ExternalData struct {
    /**
    *  商户外部会员卡卡号 
    */
    ExternalCardNo string `json:"externalCardNo"`
    /**
    *  会员卡积分，积分必须为数字型（可为浮点型，带2位小数点） 
    */
    Point string `json:"point"`
    /**
    *  资金卡余额，单位：元，精确到小数点后两位 
    */
    Balance string `json:"balance"`
    /**
    *  优惠券建议展示X张 
    */
    Coupon string `json:"coupon"`
}



func (req *MemberUpdateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*MemberUpdateResponse, error) {
    resp, err := client.InvokeApi(member_update_url, 15, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response MemberUpdateResponse
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

func (response *MemberUpdateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

