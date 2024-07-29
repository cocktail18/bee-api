package getreversionrateandrightslist

/**
* 回复率及权益查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_reversion_rate_and_rights_list_url = "/waimai/ng/im/getReversionRateAndRightsList"

type GetReversionRateAndRightsListData struct {
    /**
    * 是否获得5分钟回复率达标权益， 0-未达标 1-达标
    */
    BenefitStatus int32 `json:"benefitStatus"`
    /**
    * IM5分钟回复率
    */
    ReversionRate string `json:"reversionRate"`
    /**
    * IM5分钟目标回复率
    */
    TargetReversionRate int32 `json:"targetReversionRate"`
    /**
    * 紧急消息1分钟回复率， 0-未达标 1-达标
    */
    UrgentBenefitStatus int32 `json:"urgentBenefitStatus"`
    /**
    * 是否获得紧急消息1分钟回复率达标权益
    */
    UrgentReversionRate string `json:"urgentReversionRate"`
    /**
    * IM紧急消息目标回复率
    */
    UrgentTargetReversionRate string `json:"urgentTargetReversionRate"`
    /**
    * 门店权益列表 此次只有1，3，5，6 RIGHTS_UNKNOWN(-1, "未知权益", 1), RIGHTS_COMMENT_CONTACT(1, "评价联系", 1), RIGHTS_HONOR_TAG(2, "荣誉标签", 1), RIGHTS_COMMENT_UPDATE(3, "评价实时更新", 1), RIGHTS_IMPROVE_SCORE(4, "提升店铺评分", 2), RIGHTS_COMMENT_INVITE(5, "邀请顾客评价", 2), RIGHTS_BLOCK_LIST(6, "屏蔽恶意顾客", 2), RIGHTS_POSTER(7, "免费门店海报", 2), RIGHTS_PHONE_RETURN_VISIT(8, "电话回访顾客", 3), RIGHTS_SEND_COUPON(9, "发送优惠券", 3);
    */
    RightsList []Rights `json:"rightsList"`
    /**
    * 门店临时权益列表，同rightsList
    */
    TempRightsList []Rights `json:"tempRightsList"`
}
type Rights struct {
    Code string `json:"code"`
    Name string `json:"name"`
}
type GetReversionRateAndRightsListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data GetReversionRateAndRightsListData `json:"data"`
    TraceId string `json:"traceId"`
}
type GetReversionRateAndRightsListRequest struct {
}



func (req *GetReversionRateAndRightsListRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetReversionRateAndRightsListResponse, error) {
    resp, err := client.InvokeApi(get_reversion_rate_and_rights_list_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetReversionRateAndRightsListResponse
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

func (response *GetReversionRateAndRightsListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

