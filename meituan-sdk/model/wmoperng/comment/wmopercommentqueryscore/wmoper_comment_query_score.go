package wmopercommentqueryscore

/**
* 获取门店评分
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_comment_query_score_url = "/wmoper/ng/comment/queryScore"

type WmoperCommentQueryScoreResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data WmoperCommentQueryScoreData `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperCommentQueryScoreRequest struct {
    /**
    *  业务标识 传“B_API”只返回外卖评价分数，不传或者传其他返回所有 
    */
    BizKey string `json:"bizKey"`
}
type WmoperCommentQueryScoreData struct {
    /**
    * 服务商门店id
    */
    EpoiId string `json:"epoiId"`
    /**
    * 餐品评分
    */
    AvgPoiScore float64 `json:"avgPoiScore"`
    /**
    * 口味评分
    */
    AvgTasteScore float64 `json:"avgTasteScore"`
    /**
    * 包装评分
    */
    AvgPackingScore float64 `json:"avgPackingScore"`
    /**
    * 配送评分
    */
    AvgDeliveryScore float64 `json:"avgDeliveryScore"`
    /**
    * 配送满意度。在0~100之间，返回参数中没有%
    */
    DeliverySatisfaction int32 `json:"deliverySatisfaction"`
}



func (req *WmoperCommentQueryScoreRequest) DoInvoke(client mtclient.MeituanClient) (*WmoperCommentQueryScoreResponse, error) {
    resp, err := client.InvokeApi(wmoper_comment_query_score_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response WmoperCommentQueryScoreResponse
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

func (response *WmoperCommentQueryScoreResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

