package commentreply

/**
* 商家回复评论
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const comment_reply_url = "/review/reply"

type CommentReplyData struct {
    /**
    * 商家回复内容
    */
    BizReply string `json:"bizReply"`
    /**
    * 回复时间
    */
    ReplyTime string `json:"replyTime"`
}
type CommentReplyResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CommentReplyData `json:"data"`
    TraceId string `json:"traceId"`
}
type CommentReplyRequest struct {
    /**
    *  评论id(由查询评价接口返回) 
    */
    FeedbackId int64 `json:"feedbackId"`
    /**
    *  评论内容，只支持文字，不支持图片等其他格式 
    */
    Comment string `json:"comment"`
}



func (req *CommentReplyRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CommentReplyResponse, error) {
    resp, err := client.InvokeApi(comment_reply_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CommentReplyResponse
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

func (response *CommentReplyResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

