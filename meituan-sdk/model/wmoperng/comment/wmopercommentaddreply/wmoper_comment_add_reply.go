package wmopercommentaddreply

/**
* 根据评价id添加商家回复
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_comment_add_reply_url = "/wmoper/ng/comment/addReply"

type WmoperCommentAddReplyResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperCommentAddReplyRequest struct {
    /**
    *  外卖评价ID 
    */
    CommentId int64 `json:"commentId"`
    /**
    *  ERP方门店id 最大长度100 
    */
    EPoiId string `json:"ePoiId"`
    /**
    *  商家回复内容 
    */
    Reply string `json:"reply"`
}



func (req *WmoperCommentAddReplyRequest) DoInvoke(client mtclient.MeituanClient) (*WmoperCommentAddReplyResponse, error) {
    resp, err := client.InvokeApi(wmoper_comment_add_reply_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response WmoperCommentAddReplyResponse
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

func (response *WmoperCommentAddReplyResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

