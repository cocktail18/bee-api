package supplyrulesave

/**
* 设置可订规则
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const supply_rule_save_url = "/resv2/rule/supply/save"

type GroupVendorGroupStockRuleDTO struct {
    TableGroupDTO TableGroupDTO `json:"tableGroupDTO"`
    VendorDetailStockRuleDTOS []VendorDetailStockRuleDTO `json:"vendorDetailStockRuleDTOS"`
}
type TableGroupDTO struct {
    MinPeople int32 `json:"minPeople"`
    MaxPeople int32 `json:"maxPeople"`
}
type VendorSpecialUnbookableRuleDTO struct {
    /**
    *  特殊不可订规则的开始日期，当天0点的时间戳，单位秒，闭区间 
    */
    StartDate int32 `json:"startDate"`
    /**
    *  特殊不可订规则的结束日期，当天0点的时间戳，单位秒，闭区间 
    */
    EndDate int32 `json:"endDate"`
}
type VendorDetailStockRuleDTO struct {
    /**
    *  星期, 1-7分别表示星期一到星期日 
    */
    Weeks []int32 `json:"weeks"`
    /**
    *  餐段，1234早中晚夜 
    */
    TimePeriod int32 `json:"timePeriod"`
    /**
    *  库存数 
    */
    StockCount int32 `json:"stockCount"`
}
type VendorTimeRuleDTO struct {
    /**
    *  星期 
    */
    WeeDayList []int32 `json:"weeDayList"`
    /**
    *  餐段开始结束时间 
    */
    VendorTimePeriodDTOList []VendorTimePeriodDTO `json:"vendorTimePeriodDTOList"`
    /**
    *  提前可订规则 
    */
    VendorAdvanceRuleDTO VendorAdvanceRuleDTO `json:"vendorAdvanceRuleDTO"`
}
type VendorTimePeriodDTOSSub struct {
    /**
    *  餐段 
    */
    TimePeriodType int32 `json:"timePeriodType"`
    /**
    *  餐段起始小时 限制：0-23 
    */
    StartHour int32 `json:"startHour"`
    /**
    *  餐段起始分钟，限制0-59 
    */
    StartMinute int32 `json:"startMinute"`
    /**
    *  餐段结束小时 限制：0-23 
    */
    EndHour int32 `json:"endHour"`
    /**
    *  餐段结束分钟，限制0-59 
    */
    EndMinute int32 `json:"endMinute"`
}
type VendorTableRuleDTO struct {
    TableType int32 `json:"tableType"`
    TableTypeName string `json:"tableTypeName"`
    VendorTimeRuleDTOS []VendorTimeRuleDTO `json:"vendorTimeRuleDTOS"`
    VendorSpecialRuleDTO VendorSpecialRuleDTO `json:"vendorSpecialRuleDTO"`
    GroupVendorGroupStockRuleDTOS []GroupVendorGroupStockRuleDTO `json:"groupVendorGroupStockRuleDTOS"`
}
type VendorSpecialRuleDTO struct {
    VendorSpecialBookableRuleDTOS []VendorSpecialBookableRuleDTO `json:"vendorSpecialBookableRuleDTOS"`
    /**
    *  特殊不可订规则 
    */
    VendorSpecialUnbookableRuleDTOS []VendorSpecialUnbookableRuleDTO `json:"vendorSpecialUnbookableRuleDTOS"`
}
type SupplyRuleSaveResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type VendorAdvanceRuleDTO struct {
    /**
    *  提前多少秒可订 
    */
    EarlyBookableHour int32 `json:"earlyBookableHour"`
    /**
    *  提前多少天可订,范围为[0,30] 
    */
    EarlyBookableDay int32 `json:"earlyBookableDay"`
    /**
    *  最大可预订天数 
    */
    MaxBookableDay int32 `json:"maxBookableDay"`
}
type SupplyRuleSaveRequest struct {
    VendorTableRuleDTOS []VendorTableRuleDTO `json:"vendorTableRuleDTOS"`
    /**
    *  请求唯一id 
    */
    Uuid string `json:"uuid"`
}
type VendorSpecialBookableRuleDTO struct {
    VendorTimePeriodDTOS []VendorTimePeriodDTOSSub `json:"vendorTimePeriodDTOS"`
    VendorAdvanceRuleDTO VendorAdvanceRuleDTO `json:"vendorAdvanceRuleDTO"`
    /**
    *  开始日期 
    */
    StartDate int32 `json:"startDate"`
    /**
    *  结束日期 
    */
    EndDate int32 `json:"endDate"`
}
type VendorTimePeriodDTO struct {
    /**
    *  餐段 
    */
    TimePeriodType int32 `json:"timePeriodType"`
    /**
    *  餐段起始小时 限制：0-23 
    */
    StartHour int32 `json:"startHour"`
    /**
    *  餐段起始分钟：限制0-59 
    */
    StartMinute int32 `json:"startMinute"`
    /**
    *  餐段结束小时 限制：0-23 
    */
    EndHour int32 `json:"endHour"`
    /**
    *  餐段结束分钟：限制0-59 
    */
    EndMinute int32 `json:"endMinute"`
}



func (req *SupplyRuleSaveRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*SupplyRuleSaveResponse, error) {
    resp, err := client.InvokeApi(supply_rule_save_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response SupplyRuleSaveResponse
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

func (response *SupplyRuleSaveResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

