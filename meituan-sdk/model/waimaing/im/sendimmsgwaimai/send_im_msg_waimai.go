package sendimmsgwaimai

/**
* 发送IM消息（接单）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const send_im_msg_waimai_url = "/waimai/ng/im/msg/send"

type SendImMsgWaimaiRequest struct {
    /**
    *  业务类型字段，1两方会话，4评价匿名群聊，5粉丝群聊， 6三方群聊 
    */
    BizType int32 `json:"biz_type"`
    /**
    *  json格式的消息体，不同bizType时，字段不同，解析方式也不同。 
    */
    PushContent PushContent `json:"push_content"`
    /**
    *  扩展字段 
    */
    ExtendedField string `json:"extended_field"`
}
type SendImMsgWaimaiResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}
type PushContent struct {
    /**
    *  消息id，确保消息唯一性，发送消息时，为三方的消息id，接收消息时，为美团的消息id。 
    */
    MsgId int64 `json:"msg_id"`
    /**
    *  加密后的消息内容。 
    */
    MsgContent string `json:"msg_content"`
    /**
    *  消息类型。 bizType为1两方会话时，支持：1文本，2图片，3其他； bizType为4评价匿名群聊时，支持： 0文本，2图片； bizType为5粉丝群聊时，支持：0文本，2图片； bizType为6三方群聊时，支持：0文本，2图片。 
    */
    MsgType int32 `json:"msg_type"`
    /**
    *  消息发送方 商家：1，用户：2。 
    */
    MsgSource int32 `json:"msg_source"`
    /**
    *  消息发送时间,10位时间戳。 
    */
    Cts int64 `json:"cts"`
    /**
    *  订单id。bizType为1两方会话时，此项需要。 
    */
    OrderId int64 `json:"order_id"`
    /**
    *  用户id。bizType为1两方会话、6三方群聊时，此项需要。 
    */
    OpenUserId int64 `json:"open_user_id"`
    /**
    *  场景tag，-1时为普通消息，&gt;0时为场景消息： 9003："用户投诉商家，紧急消息" 3002："餐品质量投诉" 3001："餐品错送/少送" 9001："用户已尝试联系客服的消息" 2203："补充备注" bizType为1两方会话时，此项需要。 
    */
    SceneTag int32 `json:"scene_tag"`
    /**
    *  场景名称， 为空时时普通消息，不为空时是场景消息，取值同sceneTag 。bizType为1两方会话时，此项需要。 
    */
    SceneName string `json:"scene_name"`
    /**
    *  0默认值，1未回复，2回复。bizType为1两方会话时，此项需要。 
    */
    HasReply int32 `json:"has_reply"`
    /**
    *  评价id。bizType为4评价联系时，此项需要。 
    */
    CommentId int64 `json:"comment_id"`
    /**
    *  群id。bizType为4评价匿名群聊、5粉丝群聊、6三方群聊时，此项需要。 
    */
    GroupId int64 `json:"group_id"`
    /**
    *  群名称。bizType为5粉丝群聊时，此项需要。 
    */
    GroupName string `json:"group_name"`
    /**
    *  骑手id。bizType为6三方群聊时，此项需要。 
    */
    RiderId int64 `json:"rider_id"`
}



func (req *SendImMsgWaimaiRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SendImMsgWaimaiResponse, error) {
    resp, err := client.InvokeApi(send_im_msg_waimai_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SendImMsgWaimaiResponse
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

func (response *SendImMsgWaimaiResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

