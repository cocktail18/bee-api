package tuangoureceiptquerylistbydate

/**
* 验券记录
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const tuangou_receipt_querylistbydate_url = "/ddzh/tuangou/receipt/querylistbydate"

type ReceiptQueryBaseResult struct {
    /**
    * 券码
    */
    ReceiptCode string `json:"receiptCode"`
    /**
    * 套餐ID
    */
    DealId int64 `json:"dealId"`
    /**
    * 团购ID
    */
    DealGroupId int64 `json:"dealGroupId"`
    /**
    * 商品名称
    */
    DealTitle string `json:"dealTitle"`
    /**
    * 商品售卖价格
    */
    DealPrice float64 `json:"dealPrice"`
    /**
    * 商品类型（仅支持消费码）
    */
    ProductType int32 `json:"productType"`
    /**
    * 商品消费码（仅支持消费码验券）
    */
    ProductItemId int64 `json:"productItemId"`
    /**
    * 商品市场价
    */
    DealMarketPrice float64 `json:"dealMarketPrice"`
    /**
    * 验证账号
    */
    VerifyAccount string `json:"verifyAccount"`
    /**
    * 验证方式
    */
    VerifyChannel string `json:"verifyChannel"`
    /**
    * 验券时间
    */
    VerifyTime string `json:"verifyTime"`
    /**
    * 业务类型
    */
    BizType int32 `json:"bizType"`
    /**
    * 流水号
    */
    FlowId string `json:"flowId"`
    /**
    * 退款状态 0未使用1已使用4未消费退6已消费退
    */
    RefundStatus int32 `json:"refundStatus"`
}
type RqueryListByDateResult struct {
    /**
    * 已验券数量
    */
    TotalCount int32 `json:"totalCount"`
    Records []ReceiptQueryBaseResult `json:"records"`
}
type TuangouReceiptQuerylistbydateRequest struct {
    /**
    *  日期（如2016-01-01) 
    */
    Date string `json:"date"`
    /**
    *  业务类型： 0-普通团购 ，201-次卡，203-拼团 
    */
    BizType int32 `json:"bizType"`
    /**
    *  查询起始位置，从0开始 
    */
    Offset int32 `json:"offset"`
    /**
    *  0标识团购，1非团购券查询，默认为0，团购券查询 
    */
    Type int32 `json:"type"`
    /**
    *  查询条数，最大300 
    */
    Limit int32 `json:"limit"`
}
type TuangouReceiptQuerylistbydateResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data RqueryListByDateResult `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *TuangouReceiptQuerylistbydateRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TuangouReceiptQuerylistbydateResponse, error) {
    resp, err := client.InvokeApi(tuangou_receipt_querylistbydate_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TuangouReceiptQuerylistbydateResponse
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

func (response *TuangouReceiptQuerylistbydateResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

