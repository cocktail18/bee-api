package creategroup

/**
* 创建群接口
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const create_group_url = "/waimai/ng/im/createGroup"

type CreateGroupRequest struct {
    /**
    *  群聊类型，评价联系：1 
    */
    GroupType int32 `json:"group_type"`
    /**
    *  评价联系需要的参数，group_type为1时的必填参数， 
    */
    CommentId int64 `json:"comment_id"`
}
type CreateGroupResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CreateGroupData `json:"data"`
    TraceId string `json:"traceId"`
}
type CreateGroupData struct {
    /**
    * 群id
    */
    GroupId int64 `json:"group_id"`
    /**
    * 创建时间
    */
    Ctime int64 `json:"ctime"`
    /**
    * 群聊状态
    */
    State int32 `json:"state"`
    /**
    * 关闭类型
    */
    CloseType string `json:"close_type"`
    /**
    * 是否是新群
    */
    NewGroup string `json:"new_group"`
}



func (req *CreateGroupRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CreateGroupResponse, error) {
    resp, err := client.InvokeApi(create_group_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CreateGroupResponse
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

func (response *CreateGroupResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

