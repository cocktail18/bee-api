package createfansgroup

/**
* 创建粉丝群
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const create_fans_group_url = "/wmoper/ng/im/createFansGroup"

type CreateFansGroupResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CreateFansGroup `json:"data"`
    TraceId string `json:"traceId"`
}
type CreateFansGroupRequest struct {
    /**
    *  入群引导语 
    */
    GroupIntro string `json:"groupIntro"`
    /**
    *  用户入群欢迎语 
    */
    GroupWelcomeMsg string `json:"groupWelcomeMsg"`
    GroupEnterCoupon GroupEnterCoupon `json:"groupEnterCoupon"`
}
type GroupEnterCoupon struct {
    /**
    *  优惠券金额 
    */
    Price int32 `json:"price"`
    /**
    *  使用门槛 
    */
    LimitPrice int32 `json:"limitPrice"`
    /**
    *  有效期 
    */
    Validity int32 `json:"validity"`
}
type CreateFansGroup struct {
    /**
    * 大象群ID 偏移
    */
    GroupId int64 `json:"groupId"`
    /**
    * 群名称
    */
    GroupName string `json:"groupName"`
}



func (req *CreateFansGroupRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CreateFansGroupResponse, error) {
    resp, err := client.InvokeApi(create_fans_group_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CreateFansGroupResponse
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

func (response *CreateFansGroupResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

