package batchquerycpclaunchinfos

/**
* 获取推广信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const batch_query_cpc_launchinfos_url = "/ad/launch/batchQueryCpcLaunchInfos"

type TargetInfo struct {
    /**
    * 定向id
    */
    TargetId int32 `json:"targetId"`
    /**
    * 出价
    */
    BidPrice int32 `json:"bidPrice"`
    /**
    * 1点评，2美团
    */
    Platform int32 `json:"platform"`
    /**
    * 词包
    */
    Keyword string `json:"keyword"`
}
type BatchQueryCpcLaunchinfosRequest struct {
    /**
    * 需要查询的推广id,单次最多支持50个
    */
    LaunchIds []int64 `json:"launchIds"`
    SonAdaccountId int32 `json:"sonAdaccountId"`
}
type BatchQueryCpcLaunchinfosResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []LaunchInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type LaunchInfo struct {
    /**
    * 推广id
    */
    Launchid int32 `json:"launchid"`
    /**
    * 推广名称
    */
    LaunchName string `json:"launchName"`
    /**
    * 门店名称
    */
    ShopName string `json:"shopName"`
    /**
    * 日预算
    */
    BudgetBasic int32 `json:"budgetBasic"`
    /**
    * 高峰日上浮比例
    */
    BudgetFloatRatio int32 `json:"budgetFloatRatio"`
    /**
    * 是否支持匀速投放 0 不支持 1 支持
    */
    DeliverySpeed int32 `json:"deliverySpeed"`
    /**
    * 推广开始时间
    */
    StartTime string `json:"startTime"`
    /**
    * 推广结束时间
    */
    EndTime string `json:"endTime"`
    /**
    * 1有效，2过期，3暂停，4审核中，5审核驳回 6 删除 8 未投放（推广计划暂停） 9待推广，10 状态异常
    */
    LaunchStatus int32 `json:"launchStatus"`
    /**
    * 时间定向 7*24小时
    */
    TimeSlotPeriod []int64 `json:"timeSlotPeriod"`
    TargetBean []TargetInfo `json:"TargetBean"`
}



func (req *BatchQueryCpcLaunchinfosRequest) DoInvoke(client mtclient.MeituanClient) (*BatchQueryCpcLaunchinfosResponse, error) {
    resp, err := client.InvokeApi(batch_query_cpc_launchinfos_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response BatchQueryCpcLaunchinfosResponse
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

func (response *BatchQueryCpcLaunchinfosResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

