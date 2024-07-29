package adpricerange

/**
* 获取广告不同类型价格范围
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const ad_price_range_url = "/wmoper/ng/waimaiad/pricerange"

type AdPriceRangeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    Result Result `json:"result"`
    Min int64 `json:"min"`
    Max int64 `json:"max"`
}
type AdPriceRangeRequest struct {
    /**
    * 外卖广告产品类型，1是点金推广
    */
    Channel int32 `json:"channel"`
    /**
    * 1：预算，2：出价，3：定向出价
    */
    Type int32 `json:"type"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *AdPriceRangeRequest) DoInvoke(client mtclient.MeituanClient) (*AdPriceRangeResponse, error) {
    resp, err := client.InvokeApi(ad_price_range_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response AdPriceRangeResponse
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

func (response *AdPriceRangeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

