package governviolationquery

/**
* 违规列表查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const govern_violation_query_url = "/wmoper/ng/govern/violation/list/query"

type CollegeInfoViewDTO struct {
    /**
    * 规则标题
    */
    Title string `json:"title"`
    /**
    * 规则跳转链接
    */
    Url string `json:"url"`
}
type ResultData struct {
    /**
    * 违规列表总数
    */
    Total int32 `json:"total"`
    /**
    * 总页数
    */
    PageTotal int32 `json:"pageTotal"`
    /**
    * 当前页
    */
    PageNum int32 `json:"pageNum"`
    /**
    * 数据列表
    */
    Data []OpenViolationDetailViewDTO `json:"data"`
}
type OpenViolationDetailViewDTO struct {
    /**
    * 违规单id
    */
    Id int64 `json:"id"`
    /**
    * 门店id
    */
    WmPoiId int64 `json:"wmPoiId"`
    /**
    * 处罚时间，unix时间戳
    */
    PunishTime int32 `json:"punishTime"`
    /**
    * 违规规则id
    */
    ViolationId int32 `json:"violationId"`
    /**
    * 规则标题
    */
    ViolationTitle string `json:"violationTitle"`
    /**
    * 平台规则描述
    */
    ViolationDesc string `json:"violationDesc"`
    /**
    * 违规原因
    */
    ViolationReason string `json:"violationReason"`
    /**
    * 违规开始时间，unix时间戳
    */
    ViolationStartTime int64 `json:"violationStartTime"`
    /**
    * 违规结束时间，unix时间戳
    */
    ViolationEndTime int64 `json:"violationEndTime"`
    /**
    * 违规事实依据 返回的数量有限，50以上会被截断，需要查询更多通过分页查询违规事实依据接口
    */
    ViolationContent []ContentViewDTO `json:"violationContent"`
    /**
    * 处罚结果
    */
    PunishInfo []ViolationPunishInfoViewDTO `json:"punishInfo"`
    /**
    * 规则跳转链接
    */
    CollegeInfo CollegeInfoViewDTO `json:"collegeInfo"`
    /**
    * 申诉开始时间
    */
    AppealStartTime int32 `json:"appealStartTime"`
    /**
    * 申诉开始时间
    */
    AppealEndTime int32 `json:"appealEndTime"`
    /**
    * 阅读状态(是否在商家端查看)
    */
    IsRead int32 `json:"isRead"`
    /**
    * 申诉状态0,"不支持申诉",10,"可申诉",20,"申诉中",30,"申诉成功",40,"申诉失败",50,"逾期未申诉",60,"已确认违规" 20,"申诉中",30,"申诉成功",40,"申诉失败" 
    */
    AppealStatus int32 `json:"appealStatus"`
    /**
    * 门店名称
    */
    WmPoiName string `json:"wmPoiName"`
    /**
    * 违规分描述
    */
    ScorePunishDesc string `json:"scorePunishDesc"`
    /**
    * 规则二级类型0: "无法识别" 1: "一般违规" 51: "严重违规" 101: "积分节点一般违规" 102: "积分节点严重违规"
    */
    RuleSubType int32 `json:"ruleSubType"`
    /**
    * 违规单状态10:"待处罚,初始状态" 30:"处罚完成" 40:"处罚失败" 50:"罚单撤回"
    */
    Status int32 `json:"status"`
}
type GovernViolationQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type ViolationPunishInfoViewDTO struct {
    /**
    * 处罚描述
    */
    Title string `json:"title"`
    /**
    * 处罚内容：处罚时间格式化
    */
    Content string `json:"content"`
    /**
    * 处罚类型1:"门店置休" 2:"门店下线" 3:"删除违规商品" 4:"予以警告" 5:"门店降权" 6:"商品停售") 7:"下线所有活动" 8:"关闭到店自取" 9:"列入美团外卖商家黑名单") 10:"删除违规订单销量与评价" 11:"风控规则学习") 12:"扣除违约金" 13:"缩小配送范围" 14:"退还刷单不当得利" 15:"删除虚假销量" 16:"删除虚假评价" 99:"增加违规积分" 100:"预画管控配送范围"
    */
    Type int32 `json:"type"`
    /**
    * 处罚开始时间
    */
    StartTime int64 `json:"startTime"`
    /**
    * 处罚结束时间
    */
    EndTime int64 `json:"endTime"`
}
type ContentViewDTO struct {
    /**
    * 事实依据类型1 订单，2 商品，3 太平洋工单， 4 售后工单
    */
    Type int32 `json:"type"`
    /**
    * 总数
    */
    Total int32 `json:"total"`
    /**
    * 事实依据
    */
    ContentList []ContentDetailViewDTO `json:"contentList"`
}
type ContentDetailViewDTO struct {
    /**
    * 表明事实依据id
    */
    Id string `json:"id"`
}
type GovernViolationQueryRequest struct {
    /**
    *  页数，从1开始 
    */
    PageNum int32 `json:"pageNum"`
    /**
    *  页大小，最多20 
    */
    PageSize int32 `json:"pageSize"`
}



func (req *GovernViolationQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GovernViolationQueryResponse, error) {
    resp, err := client.InvokeApi(govern_violation_query_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GovernViolationQueryResponse
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

func (response *GovernViolationQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

