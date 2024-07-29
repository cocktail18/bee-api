package tuangoureceiptgetconsumed

/**
* 查询已验券信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const tuangou_receipt_getconsumed_url = "/ddzh/tuangou/receipt/getconsumed"

type TuangouReceiptGetconsumedRequest struct {
    /**
    *  团购券码 
    */
    ReceiptCode string `json:"receiptCode"`
}
type TuangouReceiptGetconsumedResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data RqueryConsumedResult `json:"data"`
    TraceId string `json:"traceId"`
}
type RadditionItemInfo struct {
    /**
    * 服务名
    */
    ServiceTitle string `json:"serviceTitle"`
    /**
    * 调价规则当前金额(当前金额，单位：元)
    */
    AdjustedAmount string `json:"adjustedAmount"`
    /**
    * 加项商品ID
    */
    ProductId string `json:"productId"`
    /**
    * 三方加项商品ID
    */
    AddItemThirdPartyId string `json:"addItemThirdPartyId"`
}
type RqueryConsumedResult struct {
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
    * 商品类型 1、泛商品如丽人派样活动商品等（若验证的券所对应的商品非团购时，该字段必返回）
    */
    ProductType int32 `json:"productType"`
    /**
    * 商品ID（若验证的券所对应的商品非团购时，该字段必返回，product_item_id含义参考商品管理API）
    */
    ProductItemId int64 `json:"productItemId"`
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
    * 所在订单商家营销金额
    */
    OrderShopPromoAmount float64 `json:"orderShopPromoAmount"`
    /**
    * 所在订单商家营销详情
    */
    OrderShopPromoDetails []ReceiptShopPromo `json:"orderShopPromoDetails"`
    /**
    * 验证帐号
    */
    VerifyAccount string `json:"verifyAccount"`
    /**
    * 验证方式：15-第三方验证;其他验证方式等
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
    * 附加项目信息列表
    */
    AdditionItemInfos []RadditionItemInfo `json:"additionItemInfos"`
}
type ReceiptShopPromo struct {
    /**
    * 优惠金额
    */
    PromoAmount float64 `json:"promoAmount"`
    /**
    * 优惠类型
    */
    PromoType int32 `json:"promoType"`
    /**
    * 说明信息
    */
    Desc string `json:"desc"`
}



func (req *TuangouReceiptGetconsumedRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TuangouReceiptGetconsumedResponse, error) {
    resp, err := client.InvokeApi(tuangou_receipt_getconsumed_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TuangouReceiptGetconsumedResponse
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

func (response *TuangouReceiptGetconsumedResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

