package commentcomplainreport

/**
* 评价申诉接口
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const comment_complain_report_url = "/waimai/ng/comment/complain/report"

type CommentComplainReportResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type CommentComplainReportRequest struct {
    /**
    *  评价Id 
    */
    CommentId string `json:"commentId"`
    /**
    *  申述类型 ID code name 类别 31 WM_RIDER_TIMEOUT 骑手配送超时 配送原因导致差评 32 WM_RIDER_SERVICE 骑手服务问题 34 WM_CHOOSES_WRONG_SCORE 用户选错评分 用户操作错误导致差评 35 WM_ENVIRONMENTAL 用户选择环保单 36 WM_NOT_THE_ORDER 评价与订单不符 37 WM_NOT_RECEIVED 用户未收餐 39 WM_NO_REMARKS 用户未备注产生差评 用户提出额外要求导致差评 40 WM_UNREASONABLE_REQUEST 用户提出不合理要求 41 WM_USER_THREATEN 用户以差评威胁商家 43 WM_ACCOMPANY_BAD_REVIEW 同行或特殊身份给差评 同行或特殊身份给差评 44 WM_MULTIPLE_BAD_REVIEW 多次差评 46 WM_ADVERT 不文明用语，泄露隐私，广告信息 广告或无意义评价 47 WM_NO_MEAN 内容无意义评价 48 WM_OTHER_CONTENT 其他 其他 
    */
    ReportTypeCode int32 `json:"reportTypeCode"`
    /**
    *  申诉内容 
    */
    ReportContent string `json:"reportContent"`
    /**
    *  差评申述图片 逗号分隔图片id 
    */
    PicIds string `json:"picIds"`
}



func (req *CommentComplainReportRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CommentComplainReportResponse, error) {
    resp, err := client.InvokeApi(comment_complain_report_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CommentComplainReportResponse
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

func (response *CommentComplainReportResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

