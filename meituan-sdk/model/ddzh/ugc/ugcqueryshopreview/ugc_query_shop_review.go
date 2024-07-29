package ugcqueryshopreview

/**
* 单一门店查询评论数据
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const ugc_query_shop_review_url = "/ddzh/ugc/queryshopreview"

type UgcQueryShopReviewResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data RqueryShopReviewRes `json:"data"`
    TraceId string `json:"traceId"`
}
type ScoreDetail struct {
    /**
    * 评分名称，比如环境、服务、菜品等
    */
    Title string `json:"title"`
    /**
    * 具体评分值
    */
    Score int32 `json:"score"`
    /**
    * 精确评分值（用户打半星时才会出现）
    */
    AccurateScore int32 `json:"accurateScore"`
}
type TradeDetail struct {
    /**
    * 名称
    */
    Title string `json:"title"`
    /**
    * 具体内容
    */
    Content string `json:"content"`
}
type UgcQueryShopReviewRequest struct {
    Star int32 `json:"star"`
    Offset int32 `json:"offset"`
    Limit int32 `json:"limit"`
    /**
    *  评论添加日期开始，格式如 2016-05-24 10:10:10 
    */
    BeginTime string `json:"beginTime"`
    /**
    *  评论添加日期结束，格式如 2016-06-24 10:10:10，查询区间最多一年 
    */
    EndTime string `json:"endTime"`
    /**
    *  1-大众点评，2-美团 （美团的评价与大众点评的评价是分开的，需要单独查） 
    */
    Platform int32 `json:"platform"`
}
type ReviewInfoDTO struct {
    /**
    * 评价ID
    */
    ReviewId string `json:"reviewId"`
    /**
    * 评价星级,50-5星，40-4星，30-3星，20-2星，10-1星
    */
    Star int32 `json:"star"`
    /**
    * 精确星级，50-5星, 45-4.5星, 40-4星, 35-3.5星, 30-3星, 25-2.5星, 20-2星, 15-1.5星, 10-1星, 5-0.5星
    */
    AccurateStar int32 `json:"accurateStar"`
    /**
    * 评价时间，格式如 2016-05-24 12:00 ，数据精确到分钟
    */
    ReviewTime string `json:"reviewTime"`
    /**
    * 如果是消费评价，则存在订单编号(一期仅支持团购)
    */
    ConsumeOrderId string `json:"consumeOrderId"`
    /**
    * 如果是消费评价，则存在消费金额
    */
    ConsumeAmount string `json:"consumeAmount"`
    /**
    * 如果是消费评价，则存在消费时间 格式如 2016-05-24 12:00 ，数据精确到分钟
    */
    ConsumeTime string `json:"consumeTime"`
    /**
    * 如果是消费评价，则存在消费券号,券号用逗号","分隔(可能没有)
    */
    SerialNumbers string `json:"serialNumbers"`
    /**
    * 评论质量，1-优质评价，0-非优质评价
    */
    ReviewQuality int32 `json:"reviewQuality"`
    /**
    * 详细评分，比如菜品、服务、环境等
    */
    ScoreDetails []ScoreDetail `json:"scoreDetails"`
    /**
    * 交易基础信息
    */
    TradeDetails []TradeDetail `json:"tradeDetails"`
}
type RqueryShopReviewRes struct {
    ReviewInfoDTOList []ReviewInfoDTO `json:"reviewInfoDTOList"`
}



func (req *UgcQueryShopReviewRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*UgcQueryShopReviewResponse, error) {
    resp, err := client.InvokeApi(ugc_query_shop_review_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response UgcQueryShopReviewResponse
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

func (response *UgcQueryShopReviewResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

