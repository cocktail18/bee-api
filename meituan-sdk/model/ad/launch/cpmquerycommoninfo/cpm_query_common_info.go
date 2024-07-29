package cpmquerycommoninfo

/**
* cpm查询常用信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const cpm_query_common_info_url = "/ad/launch/cpm/queryCommonInfo"

type CpmQueryCommonInfoRequest struct {
    /**
    *  实体状态类型 
    */
    EntityType int32 `json:"entityType"`
    /**
    *  门店id列表，最多支持50个 
    */
    ShopIdList []int64 `json:"shopIdList"`
    /**
    *  展位ID列表 
    */
    TagIds []bool `json:"tagIds"`
}
type CpmQueryCommonInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * private BudgetInfo budgetInfo; private List audienceDTOs; private List templateInfoList; private DistanceInfo distanceInfo; 
    */
    Data CpmCommonInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type TemplateInfo struct {
}
type CpmCommonInfo struct {
    /**
    * 预算信息
    */
    BudgetInfo BudgetInfo `json:"budgetInfo"`
    /**
    * 人群包列表
    */
    AudienceDTOs []AudienceInfo `json:"audienceDTOs"`
    /**
    * 模板信息
    */
    TemplateInfoList []TemplateInfo `json:"templateInfoList"`
    /**
    * LBS距离信息
    */
    DistanceInfo DistanceInfo `json:"distanceInfo"`
}
type AudienceInfo struct {
}
type DistanceInfo struct {
}
type BudgetInfo struct {
}



func (req *CpmQueryCommonInfoRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CpmQueryCommonInfoResponse, error) {
    resp, err := client.InvokeApi(cpm_query_common_info_url, 22, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CpmQueryCommonInfoResponse
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

func (response *CpmQueryCommonInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

