package tagpush

/**
* 品牌门店标签推送
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const tag_push_url = "/foodmop/shop/tag/push"

type TagPushResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 审核拒绝标签code列表/参数不合法
    */
    Data []string `json:"data"`
    TraceId string `json:"traceId"`
}
type TagPushRequest struct {
    /**
    *  请求流水号，唯一 
    */
    RequestId string `json:"requestId"`
    /**
    *  标签code 【size &lt;= 20】 
    */
    TagList []Tag `json:"tagList"`
}
type Tag struct {
    /**
    *  标签Code 
    */
    TagCode string `json:"tagCode"`
    /**
    *  标签名称 
    */
    TagName string `json:"tagName"`
    /**
    *  1：新增 2：修改 3：删除 
    */
    Action int32 `json:"action"`
}



func (req *TagPushRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TagPushResponse, error) {
    resp, err := client.InvokeApi(tag_push_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TagPushResponse
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

func (response *TagPushResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

