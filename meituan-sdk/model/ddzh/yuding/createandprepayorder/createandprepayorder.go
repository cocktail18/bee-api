package createandprepayorder

/**
* 下单并预支付
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const createandprepayorder_url = "/ddzh/yuding/createandprepayorder"

type CreateandprepayorderRequest struct {
    /**
    *  美团侧skuId 
    */
    ProductItemId string `json:"productItemId"`
    /**
    *  商家透传信息，美团落地到订单上(Json字符串，商家侧传用户填写的搬家信息) 
    */
    ThirdPartyFulfillInfo string `json:"thirdPartyFulfillInfo"`
    /**
    *  商家透传自定义支付参数，美团落地到支付单，支付结果通知商家时回传(Json字符串，商家可以通过回传的该信息，来识别是哪一次支付。 也可以在其中设置微信代金券标识goodsTag等信息，会透传到美团支付侧 示例：thirdPartyPaymentInfo="{\"goodsTag\":\"1111\",\"otherInfo\":\"xxx\"}") 
    */
    ThirdPartyPaymentInfo string `json:"thirdPartyPaymentInfo"`
    /**
    *  通用业务参数加密串 
    */
    GeneralBizData string `json:"generalBizData"`
    /**
    *  服务开始时间(单位：时间戳) 
    */
    ArriveTime int64 `json:"arriveTime"`
    /**
    *  服务结束时间（单位：时间戳） 
    */
    LeaveTime int64 `json:"leaveTime"`
    /**
    *  订单价格 
    */
    TotalAmount string `json:"totalAmount"`
    /**
    *  购买数量 
    */
    Quantity int32 `json:"quantity"`
    /**
    *  收货人手机号 
    */
    Mobile string `json:"mobile"`
    /**
    *  用户名称 
    */
    BookName string `json:"bookName"`
    /**
    *  用户备注 
    */
    OrderRemark string `json:"orderRemark"`
}
type CreateAndPrepayOrderResponse struct {
    /**
    * 统一订单号，外界主要感知该订单号，预支付时传该订单号
    */
    UnifiedOrderId string `json:"unifiedOrderId"`
    /**
    * 支付参数
    */
    PayParams PayParams `json:"payParams"`
}
type PayParams struct {
    /**
    * 订单号 
    */
    UnifiedOrderId string `json:"unifiedOrderId"`
    /**
    * 降级状态
    */
    DegradeStatus bool `json:"degradeStatus"`
    /**
    * 微信支付参数 
    */
    WxPayUrl string `json:"wxPayUrl"`
    /**
    * 美团支付流水号
    */
    TradeNO string `json:"tradeNO"`
    /**
    * 是否需要重定向
    */
    NeedRedirect bool `json:"needRedirect"`
    /**
    * 美团支付token 
    */
    PayToken string `json:"payToken"`
    /**
    * 降级信息
    */
    DegradeInfo DegradeInfo `json:"degradeInfo"`
}
type DegradeInfo struct {
    Key string `json:"key"`
    Value string `json:"value"`
}
type CreateandprepayorderResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CreateAndPrepayOrderResponse `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *CreateandprepayorderRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CreateandprepayorderResponse, error) {
    resp, err := client.InvokeApi(createandprepayorder_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CreateandprepayorderResponse
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

func (response *CreateandprepayorderResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

