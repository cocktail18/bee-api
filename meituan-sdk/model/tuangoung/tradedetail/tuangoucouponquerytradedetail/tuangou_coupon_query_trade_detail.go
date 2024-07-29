package tuangoucouponquerytradedetail

/**
* 查询团购订单结算明细
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const tuangou_coupon_query_trade_detail_url = "/tuangou/coupon/queryTradeDetail"

type CouponTradeDetail struct {
    /**
    * 商家促销金额
    */
    BizCost float64 `json:"bizCost"`
    /**
    * 券进价，即美团从商家获取的团购券结算批发价，也是开店宝中的“顾客实付”-“服务费”。如果商家有另外的促销（指从开店宝发起的促销），则美团与商家的结算价格为结算批发价-商家促销金额；如无另外的促销，结算批发价即作为美团与商家的结算价格
    */
    BuyPrice float64 `json:"buyPrice"`
    /**
    * 券购买价，即用户在购买团购券时所需付的价格，这里指剔除商家促销金额后的价格，即实时价格
    */
    CouponBuyPrice float64 `json:"couponBuyPrice"`
    /**
    * 券码
    */
    CouponCode string `json:"couponCode"`
    /**
    * 项目id
    */
    DealId int32 `json:"dealId"`
    /**
    * 券面值，即人们常说的市场价
    */
    DealValue float64 `json:"dealValue"`
    /**
    * 商家预计应得金额
    */
    Due float64 `json:"due"`
    /**
    * 订单id
    */
    OrderId int64 `json:"orderId"`
    /**
    * 量贩项目的单张券原价（普通券单张券原价与项目总价相同）
    */
    SingleValue float64 `json:"singleValue"`
    /**
    * 验券时间
    */
    UseTime int64 `json:"useTime"`
    /**
    * 是否量贩：0：不是，1：是
    */
    Volume int32 `json:"volume"`
}
type TuangouCouponQueryTradeDetailResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []CouponTradeDetail `json:"data"`
    TraceId string `json:"traceId"`
}
type TuangouCouponQueryTradeDetailRequest struct {
    /**
    *  团购券码 
    */
    CouponCode string `json:"couponCode"`
}



func (req *TuangouCouponQueryTradeDetailRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TuangouCouponQueryTradeDetailResponse, error) {
    resp, err := client.InvokeApi(tuangou_coupon_query_trade_detail_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TuangouCouponQueryTradeDetailResponse
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

func (response *TuangouCouponQueryTradeDetailResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

