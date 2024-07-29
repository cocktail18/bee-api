package queryliveroomcomment

/**
* 查询直播间用户评论
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_live_room_comment_url = "/mlive/comment/query"

type QueryLiveRoomCommentResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []LiveRoomCommentDTO `json:"data"`
    TraceId string `json:"traceId"`
}
type LiveRoomCommentDTO struct {
    /**
    * 游标，防止数据重复查询
    */
    Cursor int64 `json:"cursor"`
    /**
    * 昵称
    */
    Nickname string `json:"nickname"`
    /**
    * 消息内容
    */
    Message string `json:"message"`
    /**
    * 消息创建时间戳
    */
    CreateTime int64 `json:"createTime"`
}
type QueryLiveRoomCommentRequest struct {
    /**
    *  查询游标，偏移量（该字段已废弃，请通过pageNum字段进行分页查询） 
    */
    Cursor int64 `json:"cursor"`
    /**
    *  消息条数限制，最大200 
    */
    Limit int64 `json:"limit"`
    /**
    *  请求开始时间，yyyy-MM-dd HH:mm:ss 格式 
    */
    StartTime string `json:"startTime"`
    /**
    *  请求结束时间，yyyy-MM-dd HH:mm:ss 格式 
    */
    EndTime string `json:"endTime"`
    /**
    *  直播场次id 
    */
    LiveId int64 `json:"liveId"`
    /**
    *  页码 ，时间区间内的页数，从1开始， 最大限制30页 
    */
    PageNum int32 `json:"pageNum"`
}



func (req *QueryLiveRoomCommentRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryLiveRoomCommentResponse, error) {
    resp, err := client.InvokeApi(query_live_room_comment_url, 50, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryLiveRoomCommentResponse
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

func (response *QueryLiveRoomCommentResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

