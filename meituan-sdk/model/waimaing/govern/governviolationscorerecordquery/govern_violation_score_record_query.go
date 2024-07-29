package governviolationscorerecordquery

/**
* 积分记录查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const govern_violation_score_record_query_url = "/waimai/ng/govern/violation/score/record/query"

type GovernViolationScoreRecordQueryRequest struct {
    PageNum int32 `json:"pageNum"`
    PageSize int32 `json:"pageSize"`
    /**
    *  1-一般违规 51-严重违规 0-同时包含一般和严重违规 
    */
    ViolationType int32 `json:"violationType"`
}
type ResultData struct {
    /**
    * 总数
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
    Data []ViolationScoreRecordDTO `json:"data"`
}
type GovernViolationScoreRecordQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type ViolationScoreRecordDTO struct {
    /**
    * 变更原因
    */
    ReasonName string `json:"reasonName"`
    /**
    * 违规类型
    */
    ViolationType int32 `json:"violationType"`
    /**
    * 变更分数
    */
    DeltaScore int32 `json:"deltaScore"`
    /**
    * 变更前分数
    */
    FromScore int32 `json:"fromScore"`
    /**
    * 变更后分数
    */
    ToScore string `json:"toScore"`
    /**
    * 变更时间
    */
    DateTime int64 `json:"dateTime"`
    /**
    * 对应违规单id
    */
    ViolationId int64 `json:"violationId"`
}



func (req *GovernViolationScoreRecordQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GovernViolationScoreRecordQueryResponse, error) {
    resp, err := client.InvokeApi(govern_violation_score_record_query_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GovernViolationScoreRecordQueryResponse
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

func (response *GovernViolationScoreRecordQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

