package queryreceiptsbyreceiptids

/**
* 根据id批量查询团购券
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_receipts_by_receipt_ids_url = "/tuangou/coupon/queryReceiptsByReceiptIds"

type QueryReceiptsByReceiptIdsRequest struct {
    DealId int32 `json:"dealId"`
    /**
    *  逗号分割receiptId列表字符串 
    */
    ReceiptIds string `json:"receiptIds"`
}
type ReceiptPromotionDTO struct {
    PromotionId string `json:"promotionId"`
    ActivityProductId int32 `json:"activityProductId"`
    PromotionType int32 `json:"promotionType"`
    BizReduce string `json:"bizReduce"`
}
type QueryReceiptsByIdResp struct {
    /**
    * 团购券查询状态
    */
    Status ReceiptQueryStatusDTO `json:"status"`
    /**
    * 团购券id
    */
    ReceiptId int64 `json:"receiptId"`
    /**
    * 团购券状态
    */
    CouponStatus int32 `json:"couponStatus"`
    /**
    * 使用时间
    */
    CouponUseTime int64 `json:"couponUseTime"`
    /**
    * 验证类型
    */
    VerifyType string `json:"verifyType"`
    /**
    * 项目id
    */
    DealId int32 `json:"dealId"`
    /**
    * 项目名称
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 市场价
    */
    DealValue float64 `json:"dealValue"`
    /**
    * 项目价格
    */
    DealPrice float64 `json:"dealPrice"`
    /**
    * 是否代金券
    */
    IsVoucher bool `json:"isVoucher"`
    /**
    * 是否量贩
    */
    Volume int32 `json:"volume"`
    /**
    * 量贩项目的单张券原价
    */
    SingleValue BigDecimal `json:"singleValue"`
    /**
    * 团购券购买价格
    */
    CouponBuyPrice float64 `json:"couponBuyPrice"`
    TotalBizReduce float64 `json:"totalBizReduce"`
    OldBiz bool `json:"oldBiz"`
    /**
    * 团购券优惠
    */
    ReceiptPromotions []ReceiptPromotionDTO `json:"receiptPromotions"`
}
type QueryReceiptsByReceiptIdsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data QueryReceiptsByIdResp `json:"data"`
    TraceId string `json:"traceId"`
}
type BigDecimal struct {
}
type ReceiptQueryStatusDTO struct {
    /**
    * 状态码
    */
    Code int32 `json:"code"`
    /**
    * 状态描述
    */
    Msg string `json:"msg"`
}



func (req *QueryReceiptsByReceiptIdsRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryReceiptsByReceiptIdsResponse, error) {
    resp, err := client.InvokeApi(query_receipts_by_receipt_ids_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryReceiptsByReceiptIdsResponse
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

func (response *QueryReceiptsByReceiptIdsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

