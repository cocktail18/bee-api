package querystorefansgroup

/**
* 查询门店的粉丝群
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_store_fans_group_url = "/waimai/ng/im/queryStoreFansGroup"

type QueryStoreFansGroup struct {
    /**
    * 大象群ID 偏移
    */
    GroupId int64 `json:"groupId"`
    /**
    * 群名称
    */
    GroupName string `json:"groupName"`
    /**
    * 粉丝群序号
    */
    GroupIndex int32 `json:"groupIndex"`
    /**
    * 群聊人数
    */
    MemberCount int32 `json:"memberCount"`
    /**
    * 群公告
    */
    GroupAnnounce string `json:"groupAnnounce"`
    /**
    * 入群导语
    */
    GroupIntro string `json:"groupIntro"`
    /**
    * 群入口开关 0-开启 1-关闭
    */
    GroupEntrance int32 `json:"groupEntrance"`
    /**
    * 群状态 1-正常 2-关闭
    */
    Status int32 `json:"status"`
}
type QueryStoreFansGroupRequest struct {
    /**
    *  页码 从 1 开始 
    */
    Page int32 `json:"page"`
    /**
    *  页数据条数 
    */
    PageSize int32 `json:"pageSize"`
}
type QueryStoreFansGroupResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data QueryStoreFansGroup `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryStoreFansGroupRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryStoreFansGroupResponse, error) {
    resp, err := client.InvokeApi(query_store_fans_group_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryStoreFansGroupResponse
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

func (response *QueryStoreFansGroupResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

