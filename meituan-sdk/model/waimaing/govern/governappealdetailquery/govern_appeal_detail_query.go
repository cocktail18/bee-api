package governappealdetailquery

/**
* 申诉详情查询
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const govern_appeal_detail_query_url = "/waimai/ng/govern/appeal/detail/query"

type GovernAppealDetailQueryResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data OpenViolationAppealDTO `json:"data"`
    TraceId string `json:"traceId"`
}
type GovernAppealDetailQueryRequest struct {
    /**
    *  违规单Id 
    */
    ViolationDetailId int64 `json:"violationDetailId"`
}
type AppealMediaDTO struct {
    /**
    * 1：图片 ，2：视频
    */
    Type int32 `json:"type"`
    /**
    * 图片url
    */
    PicUrl string `json:"picUrl"`
    /**
    * 视频url
    */
    VideoUrl string `json:"videoUrl"`
}
type OpenViolationAppealDTO struct {
    /**
    * 申诉单id
    */
    Id int32 `json:"id"`
    /**
    * 门店id
    */
    AppPoiCode string `json:"appPoiCode"`
    /**
    * 0,"不支持申诉",10,"可申诉",20,"申诉中",30,"申诉成功",40,"申诉失败",50,"逾期未申诉",60,"已确认违规"
    */
    Status int32 `json:"status"`
    /**
    * 处理结果文案
    */
    Description string `json:"description"`
    /**
    * 商家申请文字材料
    */
    Content string `json:"content"`
    /**
    * 申诉媒体材料
    */
    MediaList []AppealMediaDTO `json:"mediaList"`
    /**
    * 违规单id
    */
    ViolationDetailId int64 `json:"violationDetailId"`
    /**
    * 申诉时间
    */
    Ctime int64 `json:"ctime"`
    /**
    * 最后处理时间
    */
    Utime int64 `json:"utime"`
}



func (req *GovernAppealDetailQueryRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GovernAppealDetailQueryResponse, error) {
    resp, err := client.InvokeApi(govern_appeal_detail_query_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GovernAppealDetailQueryResponse
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

func (response *GovernAppealDetailQueryResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

