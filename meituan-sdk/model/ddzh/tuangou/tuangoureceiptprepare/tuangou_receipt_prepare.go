package tuangoureceiptprepare

/**
* 输码验券校验
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const tuangou_receipt_prepare_url = "/ddzh/tuangou/receipt/prepare"

type TuangouReceiptPrepareResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type PaymentDetailSub struct {
    /**
    * 支付详情id
    */
    PaymentDetailId string `json:"paymentDetailId"`
    /**
    * 支付金额
    */
    Amount string `json:"amount"`
    /**
    * 金额类型
    */
    AmountType int32 `json:"amountType"`
}
type TuangouReceiptPrepareRequest struct {
    /**
    *  团购券码，必须未验证 
    */
    ReceiptCode string `json:"receiptCode"`
}
type Data struct {
    /**
    * 可验证的张数
    */
    Count int32 `json:"count"`
    /**
    * 订单ID
    */
    OrderId string `json:"orderId"`
    /**
    * 验证券码
    */
    ReceiptCode string `json:"receiptCode"`
    /**
    * 套餐ID（若验证的券所对应的商品为团购时，该字段必返回）
    */
    DealId int64 `json:"dealId"`
    /**
    * 团购ID，团购id与套餐id是一对多的关系（若验证的券所对应的商品为团购时，该字段必返回）
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
    * 商品市场价
    */
    DealMarketPrice float64 `json:"dealMarketPrice"`
    /**
    * 用户手机号
    */
    Mobile string `json:"mobile"`
    /**
    * 业务类型 0:普通团购 203:拼团 205:次卡 217:通兑标品
    */
    BizType int32 `json:"bizType"`
    /**
    * 支付明细(订单维度，如果订单有多张券可以通过订单券码分摊金额查询接口查询)
    */
    PaymentDetail []PaymentDetailSub `json:"paymentDetail"`
    /**
    * 券过期时间
    */
    ReceiptEndDate int64 `json:"receiptEndDate"`
    /**
    * 是否是团购次卡
    */
    TgTimesCardFlag bool `json:"tgTimesCardFlag"`
    /**
    * 团购次卡可用总次数
    */
    PurchaseToConsumeRatio string `json:"purchaseToConsumeRatio"`
    /**
    * 团购次卡单次卡价格 
    */
    TimesCardUnitPrice string `json:"timesCardUnitPrice"`
    /**
    * 补贴金额（平台分摊）
    */
    PlatformAmount string `json:"platformAmount"`
    /**
    * 补贴金额（商户分摊）
    */
    MerchantAmount string `json:"merchantAmount"`
    /**
    * 体检加项信息
    */
    AdditionItemInfos []RadditionItemInfo `json:"additionItemInfos"`
    /**
    * 商品类型 1、泛商品如丽人派样活动商品等（若验证的券所对应的商品非团购时，该字段必返回）
    */
    ProductType int64 `json:"productType"`
    /**
    * 商品ID（若验证的券所对应的商品非团购时，该字段必返回，productItemId含义参考商品管理API）
    */
    ProductItemId int64 `json:"productItemId"`
    /**
    * 三方商品ID（仅限家居、珠宝行业）
    */
    ThirdPartyProductId string `json:"thirdPartyProductId"`
    /**
    * Map > JSON格式。多团单维度券信息，如果为null则为单团单,key为productItemId
    */
    ReceiptInfoMap string `json:"receiptInfoMap"`
}
type RadditionItemInfo struct {
    /**
    * 服务名称
    */
    ServiceTitle string `json:"serviceTitle"`
    /**
    * 加项价格
    */
    AdjustedAmount string `json:"adjustedAmount"`
    /**
    * 商品ID
    */
    ProductId string `json:"productId"`
    /**
    * 三方加项ID
    */
    AddItemThirdPartyId string `json:"addItemThirdPartyId"`
}



func (req *TuangouReceiptPrepareRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TuangouReceiptPrepareResponse, error) {
    resp, err := client.InvokeApi(tuangou_receipt_prepare_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TuangouReceiptPrepareResponse
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

func (response *TuangouReceiptPrepareResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

