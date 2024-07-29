package bizordlist

/**
* 经营分析-订单列表
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const biz_ord_list_url = "/wmoper/ng/waimaiad/biz/ordlist"

type BizOrdListRequest struct {
    /**
    * 日期，yyyyMMdd格式
    */
    Dt string `json:"dt"`
    /**
    * 页数，从1开始
    */
    PageNum int32 `json:"pageNum"`
    /**
    * 一页数量，不能超过200
    */
    PageSize int32 `json:"pageSize"`
}
type OrdList struct {
    WmOrderViewId int64 `json:"wmOrderViewId"`
    OriginalPrice int64 `json:"originalPrice"`
    ActualPrice int64 `json:"actualPrice"`
    ShippingFee int64 `json:"shippingFee"`
    OrderTime string `json:"orderTime"`
    /**
    * json字符串，foodId&food名称
    */
    FoodSet string `json:"foodSet"`
    /**
    * json字符串
    */
    ActSet string `json:"actSet"`
    Dt int64 `json:"dt"`
}
type Data struct {
    Result Result `json:"result"`
    Count int64 `json:"count"`
    OrdList []OrdList `json:"ordList"`
}
type BizOrdListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Result struct {
    Success bool `json:"success"`
    Code string `json:"code"`
    Message string `json:"message"`
}



func (req *BizOrdListRequest) DoInvoke(client mtclient.MeituanClient) (*BizOrdListResponse, error) {
    resp, err := client.InvokeApi(biz_ord_list_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response BizOrdListResponse
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

func (response *BizOrdListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

