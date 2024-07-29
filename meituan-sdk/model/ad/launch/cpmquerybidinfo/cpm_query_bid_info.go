package cpmquerybidinfo

/**
* cpm查询出价信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const cpm_query_bid_info_url = "/ad/launch/cpm/queryBidInfo"

type BidInfo struct {
}
type CpmQueryBidInfoRequest struct {
    /**
    *  推广目的 
    */
    LaunchAim int32 `json:"launchAim"`
    /**
    *  实体类型 
    */
    EntityType int32 `json:"entityType"`
    /**
    *  门店ID，所有门店须在同一个一级前台类目下 
    */
    ShopIdList []int64 `json:"shopIdList"`
    /**
    *  推广时间，7*24h，由0-167间的数字组成，不能为空 
    */
    TimePeriod []int32 `json:"timePeriod"`
    /**
    *  展位ID列表 
    */
    TagIds []int32 `json:"tagIds"`
    /**
    *  投放区域类型，1-lbs，2-城市商圈 
    */
    LaunchType int32 `json:"launchType"`
    /**
    *  点评城市ID，launchType=2时必填 
    */
    DpCityIds []int32 `json:"dpCityIds"`
    /**
    *  lbs距离，launchType=1时必填 
    */
    LbsDistance int32 `json:"lbsDistance"`
    /**
    *  推广人群。0-全部人群，1-自定义人群，2-智选人群 
    */
    LaunchScope int32 `json:"launchScope"`
    /**
    *  人群包ID，选择自定义人群时必填。 
    */
    AudienceIds []int32 `json:"audienceIds"`
    /**
    *  日预算，单位元 
    */
    DailyBudget int32 `json:"dailyBudget"`
    /**
    *  批量参数 
    */
    Params []BatchCreateCpmLaunchParam `json:"params"`
}
type BatchCreateCpmLaunchParam struct {
    /**
    *  出价模式 
    */
    PriceMode int32 `json:"priceMode"`
}
type CpmQueryBidInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * // 出价模式 private int priceMode; // 默认出价，元 private BigDecimal defaultBid; // 保底价，元 private BigDecimal bottomPrice;
    */
    Data []BidInfo `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *CpmQueryBidInfoRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CpmQueryBidInfoResponse, error) {
    resp, err := client.InvokeApi(cpm_query_bid_info_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CpmQueryBidInfoResponse
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

func (response *CpmQueryBidInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

