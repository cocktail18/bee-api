package governwarninglistquery

/**
* 预警列表查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const govern_warning_list_query_url = "/wmoper/ng/govern/warning/list/query"

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
type GovernWarningListQueryRequest struct {
    /**
    *  页数，从1开始 
    */
    PageNum int32 `json:"pageNum"`
    /**
    *  页大小，最多20 
    */
    PageSize int32 `json:"pageSize"`
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
    EpoiId int32 `json:"epoiId"`
    /**
    * 预警时间，unix时间戳
    */
    WarningTime int32 `json:"warningTime"`
    /**
    * 预警规则id
    */
    WarningId int32 `json:"warningId"`
    /**
    * 预警类型
    */
    WarningTitle string `json:"warningTitle"`
    /**
    * 平台规则描述
    */
    WarningDesc string `json:"warningDesc"`
    /**
    * 违规原因
    */
    WarningReason string `json:"warningReason"`
    /**
    * 预警的处罚描述
    */
    PunishDesc string `json:"punishDesc"`
    /**
    * 预警计算开始时间，unix时间戳
    */
    CalStartTime int64 `json:"calStartTime"`
    /**
    * 预警计算束时间，unix时间戳
    */
    CalEndTime int64 `json:"calEndTime"`
    /**
    * 预警事实依据
    */
    WarningContent []ContentViewDTO `json:"warningContent"`
    /**
    * 规则跳转链接
    */
    CollegeInfo CollegeInfoViewDTO `json:"collegeInfo"`
    /**
    * 阅读状态
    */
    IsRead int32 `json:"isRead"`
    /**
    * 门店名称
    */
    WmPoiName string `json:"wmPoiName"`
}
type GovernWarningListQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type ContentViewDTO struct {
    /**
    * 事实依据类型
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



func (req *GovernWarningListQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GovernWarningListQueryResponse, error) {
    resp, err := client.InvokeApi(govern_warning_list_query_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GovernWarningListQueryResponse
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

func (response *GovernWarningListQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

