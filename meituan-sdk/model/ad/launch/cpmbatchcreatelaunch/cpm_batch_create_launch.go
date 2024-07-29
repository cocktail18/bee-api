package cpmbatchcreatelaunch

/**
* cpm批量创建投放
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const cpm_batch_create_launch_url = "/ad/launch/cpm/batchCreateLaunch"

type CreativeElementDTO struct {
    /**
    *  模板ID 
    */
    TemplateId int32 `json:"templateId"`
    /**
    *  创意主图 
    */
    PicUrl string `json:"picUrl"`
    /**
    *  创意标题 
    */
    Title string `json:"title"`
}
type CpmBatchCreateLaunchResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * // 成功数据 private List successList; // 失败数据和原因 private Map failReasonMap;
    */
    Data BatchOperateRsp `json:"data"`
    TraceId string `json:"traceId"`
}
type BatchOperateRsp struct {
}
type CpmBatchCreateLaunchRequest struct {
    /**
    *  推广目的，当前仅开放自定义：3 
    */
    LaunchAim int32 `json:"launchAim"`
    /**
    *  实体类型，当前仅开放自定义落地页链接：11 
    */
    EntityType int32 `json:"entityType"`
    /**
    *  门店ID，所有门店须在同一个一级前台类目下，不超过50个 
    */
    ShopIdList []int64 `json:"shopIdList"`
    /**
    *  投放日期，只填一个时段，暂不开放多时段 
    */
    Duration []CpmTimeIntervalDTO `json:"duration"`
    /**
    *  推广时间，7*24h，由0-167间的数字组成，不能为空 
    */
    TimePeriod []int32 `json:"timePeriod"`
    /**
    *  展位ID列表，仅开放410-点评首页信息流、552-美团首页猜你喜欢 
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
    *  lbs距离，单位米，1000的整数倍，launchType=1时必填 
    */
    LbsDistance int32 `json:"lbsDistance"`
    /**
    *  推广人群。0-全部人群，1-自定义人群，2-智选人群。默认智选人群 
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
type CpmTimeIntervalDTO struct {
    /**
    *  投放开始日期，格式"YYYY-MM-DD" 
    */
    BeginDate string `json:"beginDate"`
    /**
    *  投放结束日期，格式"YYYY-MM-DD"，不填默认为"2099-12-31" 
    */
    EndDate string `json:"endDate"`
}
type BatchCreateCpmLaunchParam struct {
    /**
    *  出价 
    */
    BidPrice float64 `json:"bidPrice"`
    /**
    *  出价模式 
    */
    PriceMode int32 `json:"priceMode"`
    /**
    *  点击出价上浮比例 
    */
    FloatRatio int32 `json:"floatRatio"`
    /**
    *  落地页链接 
    */
    LandingUrl string `json:"landingUrl"`
    /**
    *  创意头像 
    */
    AvatarUrl string `json:"avatarUrl"`
    /**
    *  创意昵称 
    */
    BrandName string `json:"brandName"`
    /**
    *  其余创意元素 
    */
    Elements []CreativeElementDTO `json:"elements"`
}



func (req *CpmBatchCreateLaunchRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CpmBatchCreateLaunchResponse, error) {
    resp, err := client.InvokeApi(cpm_batch_create_launch_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CpmBatchCreateLaunchResponse
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

func (response *CpmBatchCreateLaunchResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

