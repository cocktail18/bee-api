package memberquery

/**
* 查询用户会员卡信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const member_query_url = "/jmcard/members/query"

type CardResult struct {
    /**
    * 表单内容列表
    */
    FormItemList []FormItem `json:"formItemList"`
    /**
    * 用户领卡时间
    */
    ReceiveTime int64 `json:"receiveTime"`
    /**
    * 用户卡过期时间
    */
    ExpireTime int64 `json:"expireTime"`
    /**
    * 卡号
    */
    CardNumber string `json:"cardNumber"`
}
type MemberQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Result `json:"data"`
    TraceId string `json:"traceId"`
}
type FormItem struct {
    /**
    * 注册信息类别，可选值为NAME,SEX,PHONE,BIRTHDAY,CITY
    */
    ItemEnum string `json:"itemEnum"`
    /**
    * 值
    */
    Value string `json:"value"`
}
type MemberQueryRequest struct {
    /**
    *  商户唯一标识，领卡事件中的userOpenId 
    */
    UserOpenId string `json:"userOpenId"`
}
type Result struct {
    /**
    * 业务状态
    */
    Status string `json:"status"`
    Result CardResult `json:"result"`
    /**
    * 业务异常时的描述消息
    */
    ErrMsg string `json:"errMsg"`
}



func (req *MemberQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*MemberQueryResponse, error) {
    resp, err := client.InvokeApi(member_query_url, 15, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response MemberQueryResponse
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

func (response *MemberQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

