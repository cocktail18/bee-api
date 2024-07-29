package evaluate

/**
* 评价骑手
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const evaluate_url = "/peisong/order/evaluate"

type EvaluateData struct {
    /**
    * 配送活动标识
    */
    DeliveryId int64 `json:"delivery_id"`
    /**
    * 美团配送内部订单id，最长不超过32个字符
    */
    MtPeisongId string `json:"mt_peisong_id"`
    /**
    * 外部订单号，最长不超过32个字符
    */
    OrderId string `json:"order_id"`
}
type EvaluateRequest struct {
    /**
    * 配送活动标识
    */
    DeliveryId int64 `json:"delivery_id"`
    /**
    * 美团配送内部订单id，最长不超过32个字符
    */
    MtPeisongId string `json:"mt_peisong_id"`
    /**
    * 评分（5分制） 预留参数，不作为骑手反馈参考 合作方需传入0-5之间分数或者不传，否则报错
    */
    Score int32 `json:"score"`
    /**
    * 评论内容（评论的字符长度需小于1024）
    */
    CommentContent string `json:"comment_content"`
}
type EvaluateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data EvaluateData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *EvaluateRequest) DoInvoke(client mtclient.MeituanClient) (*EvaluateResponse, error) {
    resp, err := client.InvokeApi(evaluate_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response EvaluateResponse
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

func (response *EvaluateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

