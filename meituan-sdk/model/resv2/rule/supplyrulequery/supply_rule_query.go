package supplyrulequery

/**
* 查询规则信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const supply_rule_query_url = "/resv2/rule/supply/query"

type VendorSpecialUnbookableRuleDTO struct {
    /**
    * 开始日期
    */
    StartDate int32 `json:"startDate"`
    /**
    * 结束日期
    */
    EndDate int32 `json:"endDate"`
}
type VendorGroupStockRuleDTO struct {
    /**
    * 桌位信息
    */
    TableGroupDTO VendorTableGroupDTO `json:"tableGroupDTO"`
    /**
    * 库存信息
    */
    VendorDetailStockRuleDTOS []VendorDetailStockRuleDTO `json:"vendorDetailStockRuleDTOS"`
}
type VendorDetailStockRuleDTO struct {
    /**
    * 星期
    */
    Weeks []int32 `json:"weeks"`
    /**
    * 餐段
    */
    TimePeriod []int32 `json:"timePeriod"`
    /**
    * 库存数量
    */
    StockCount int32 `json:"stockCount"`
}
type VendorTimeRuleDTO struct {
    /**
    * 星期
    */
    WeeDayList []int32 `json:"weeDayList"`
    /**
    * 规则时间信息
    */
    VendorTimePeriodDTOList []VendorTimePeriodDTO `json:"vendorTimePeriodDTOList"`
    /**
    * 提前可订规则
    */
    VendorAdvanceRuleDTO VendorAdvanceRuleDTO `json:"vendorAdvanceRuleDTO"`
}
type VendorTableRuleDTO struct {
    /**
    * 桌型id
    */
    TableTypeId int64 `json:"tableTypeId"`
    /**
    * 桌型类型 0-大厅 1-包间 3-特殊桌型
    */
    TableType int32 `json:"tableType"`
    /**
    * 桌型名称
    */
    TableTypeName string `json:"tableTypeName"`
    /**
    * 桌型可预订餐段的配置
    */
    VendorTimeRuleDTOS []VendorTimeRuleDTO `json:"vendorTimeRuleDTOS"`
    /**
    * 桌型可预订餐段的特殊配置
    */
    VendorSpecialRuleDTO VendorSpecialRuleDTO `json:"vendorSpecialRuleDTO"`
    /**
    * 库存规则
    */
    GroupVendorGroupStockRuleDTOS []VendorGroupStockRuleDTO `json:"groupVendorGroupStockRuleDTOS"`
}
type VendorSpecialRuleDTO struct {
    /**
    * 特殊可订规则
    */
    VendorSpecialBookableRuleDTOS []VendorSpecialBookableRuleDTO `json:"vendorSpecialBookableRuleDTOS"`
    /**
    * 特殊不可订规则
    */
    VendorSpecialUnbookableRuleDTOS []VendorSpecialUnbookableRuleDTO `json:"vendorSpecialUnbookableRuleDTOS"`
}
type VendorAdvanceRuleDTO struct {
    /**
    * 提前多少小时可订，单位秒
    */
    EarlyBookableHour int32 `json:"earlyBookableHour"`
    /**
    * 提前多少天可订,范围为[0,30]，XX天以后的座位才可订
    */
    EarlyBookableDay int32 `json:"earlyBookableDay"`
    /**
    * 最大可预订天数, 天数,范围为[0,30]
    */
    MaxBookableDay int32 `json:"maxBookableDay"`
}
type SupplyRuleQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []VendorTableRuleDTO `json:"data"`
    TraceId string `json:"traceId"`
}
type VendorSpecialBookableRuleDTO struct {
    /**
    * 规则时间信息
    */
    VendorTimePeriodDTOS []VendorTimePeriodDTO `json:"vendorTimePeriodDTOS"`
    /**
    * 提前可订规则
    */
    VendorAdvanceRuleDTO VendorAdvanceRuleDTO `json:"vendorAdvanceRuleDTO"`
    /**
    * 开始日期，秒级时间戳
    */
    StartDate int32 `json:"startDate"`
    /**
    * 结束日期，秒级时间戳
    */
    EndDate int32 `json:"endDate"`
}
type SupplyRuleQueryRequest struct {
}
type VendorTimePeriodDTO struct {
    /**
    * 餐段
    */
    TimePeriodType int32 `json:"timePeriodType"`
    /**
    * 开始小时
    */
    StartHour int32 `json:"startHour"`
    /**
    * 开始分钟
    */
    StartMinute int32 `json:"startMinute"`
    /**
    * 结束小时
    */
    EndHour int32 `json:"endHour"`
    /**
    * 结束分钟
    */
    EndMinute int32 `json:"endMinute"`
}
type VendorTableGroupDTO struct {
    /**
    * 该桌位最多可坐的人数
    */
    MaxPeople int32 `json:"maxPeople"`
    /**
    * 该桌位最少可坐的人数
    */
    MinPeople int32 `json:"minPeople"`
    /**
    * 桌位id
    */
    TableGroupId int64 `json:"tableGroupId"`
}



func (req *SupplyRuleQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SupplyRuleQueryResponse, error) {
    resp, err := client.InvokeApi(supply_rule_query_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SupplyRuleQueryResponse
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

func (response *SupplyRuleQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

