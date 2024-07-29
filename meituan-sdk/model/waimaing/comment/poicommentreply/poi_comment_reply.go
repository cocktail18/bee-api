package poicommentreply

/**
* 外卖评价回复
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const poi_comment_reply_url = "/waimai/poi/addReply"

type PoiCommentReplyRequest struct {
    /**
    *  外卖评价ID 
    */
    CommentId int64 `json:"commentId"`
    /**
    *  商家回复内容 
    */
    Reply string `json:"reply"`
}
type PoiCommentReplyResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *PoiCommentReplyRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*PoiCommentReplyResponse, error) {
    resp, err := client.InvokeApi(poi_comment_reply_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response PoiCommentReplyResponse
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

func (response *PoiCommentReplyResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

