package sendliveroomcomment

/**
* 发送直播间评论
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const send_live_room_comment_url = "/mlive/comment/send"

type SendLiveRoomCommentRequest struct {
    /**
    *  数字人形象id 
    */
    DigitalImageId string `json:"digitalImageId"`
    /**
    *  直播评论DTO 
    */
    LiveRoomCommentDTO LiveRoomCommentDTO `json:"liveRoomCommentDTO"`
    /**
    *  回复评论的id，非必传 
    */
    FollowCommentId string `json:"followCommentId"`
    /**
    *  直播id 
    */
    LiveId int64 `json:"liveId"`
    /**
    *  幂等id 
    */
    UniqueId string `json:"uniqueId"`
    /**
    *  链接别名，若无，使用uuid即可 
    */
    Alias string `json:"alias"`
}
type LiveRoomCommentDTO struct {
    /**
    *  消息内容 
    */
    Message string `json:"message"`
}
type SendLiveRoomCommentResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 布尔值，code=0时，返回true
    */
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *SendLiveRoomCommentRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SendLiveRoomCommentResponse, error) {
    resp, err := client.InvokeApi(send_live_room_comment_url, 50, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SendLiveRoomCommentResponse
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

func (response *SendLiveRoomCommentResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

