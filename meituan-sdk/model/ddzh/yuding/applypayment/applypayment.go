package applypayment

/**
* 预支付
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const applypayment_url = "/ddzh/yuding/applypayment"

type Data struct {
    /**
    * 统一订单号
    */
    UnifiedOrderId string `json:"unifiedOrderId"`
    /**
    * 支付参数
    */
    PayParams PayParam `json:"payParams"`
}
type ApplypaymentRequest struct {
    /**
    *  美团统一订单号 
    */
    UnifiedOrderId string `json:"unifiedOrderId"`
    /**
    *  支付金额类型，1-订金，2-尾款 
    */
    PayAmountType int32 `json:"payAmountType"`
    /**
    *  支付请求类型，1-首次请求支付参数（仅支持尾款。订金的首次支付请求，需要调book.createandprepayorder接口，完成下单&amp;请求支付参数）， 2-二次支付请求（支持订金和尾款），3-二次支付且要求修改支付金额（仅支持尾款） 
    */
    PayActionType int32 `json:"payActionType"`
    /**
    *  尾款金额(尾款首次请求支付参数时需传递； 尾款改价时，需传递最新金额。) 
    */
    FinalPaymentAmount string `json:"finalPaymentAmount"`
    /**
    *  商家透传自定义支付参数，美团落地到支付单，支付结果通知商家时回传 
    */
    ThirdPartyPaymentInfo string `json:"thirdPartyPaymentInfo"`
    /**
    *  美团skuId 
    */
    ProductItemId string `json:"productItemId"`
    /**
    *  通用业务参数加密串 
    */
    GeneralBizData string `json:"generalBizData"`
}
type DegradeInfo struct {
    Key string `json:"key"`
    Value string `json:"value"`
}
type ApplypaymentResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type PayParam struct {
    /**
    * 美团支付流水号
    */
    TradeNO string `json:"tradeNO"`
    /**
    * 美团支付token
    */
    PayToken string `json:"payToken"`
    /**
    * 订单号
    */
    UnifiedOrderId string `json:"unifiedOrderId"`
    /**
    * 是否需要重定向
    */
    NeedRedirect bool `json:"needRedirect"`
    /**
    * 微信支付参数
    */
    WxPayUrl string `json:"wxPayUrl"`
    /**
    * 降级信息
    */
    DegradeInfo DegradeInfo `json:"degradeInfo"`
    /**
    * 降级状态
    */
    DegradeStatus bool `json:"degradeStatus"`
}



func (req *ApplypaymentRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ApplypaymentResponse, error) {
    resp, err := client.InvokeApi(applypayment_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ApplypaymentResponse
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

func (response *ApplypaymentResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

