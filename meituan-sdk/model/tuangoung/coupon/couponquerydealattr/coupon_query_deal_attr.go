package couponquerydealattr

/**
* 查询团购项目限制条件接口
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const coupon_query_deal_attr_url = "/tuangou/coupon/queryDealAttr"

type CouponQueryDealAttrRequest struct {
    /**
    *  项目id列表，不超过50个，英文逗号隔开【,】。 
    */
    DealIds string `json:"dealIds"`
}
type AttrInfoList struct {
    /**
    * attrId
    */
    AttrId int32 `json:"attrId"`
    /**
    * 拓展属性的值（由团购侧接口返回，不作处理）
    */
    Value string `json:"value"`
}
type CouponQueryDealAttrData struct {
    DealInfoList []DealInfoList `json:"dealInfoList"`
}
type DealInfoList struct {
    /**
    * 当前dealId对应的attrId以及属性值
    */
    AttrInfoList []AttrInfoList `json:"attrInfoList"`
    /**
    * dealId
    */
    DealId int32 `json:"dealId"`
}
type CouponQueryDealAttrResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CouponQueryDealAttrData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *CouponQueryDealAttrRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CouponQueryDealAttrResponse, error) {
    resp, err := client.InvokeApi(coupon_query_deal_attr_url, 1, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CouponQueryDealAttrResponse
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

func (response *CouponQueryDealAttrResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

