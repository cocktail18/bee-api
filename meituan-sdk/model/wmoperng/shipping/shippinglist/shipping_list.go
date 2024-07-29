package shippinglist

/**
* 查询门店配送范围
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const shipping_list_url = "/wmoper/ng/shipping/list"

type Data struct {
    /**
    * 服务商方提供的配送范围id
    */
    AppShippingCode string `json:"app_shipping_code"`
    /**
    * 配送范围type为1时: [{"x":39941199,"y":116385384},{"x":39926983,"y":116361694},{"x ":39921586,"y":116398430}]， type为2时: {"r":1000,"x":40029380,"y":116418390}， 需要对其urlEncode，x代表纬度，y代表经度（美团使用的是高德坐标系，也就是火星坐标系，如果是百度坐标系需要转换，配送范围坐标需要乘以一百万) app_shipping_code string 125 
    */
    Area string `json:"area"`
    LogisticsCode string `json:"logistics_code"`
    LogisticsType string `json:"logistics_type"`
    /**
    * 该配送区域的起送价
    */
    MinPrice float64 `json:"min_price"`
    /**
    * 该配送区域的配送费(建议填写这个字段设定配送费,如果此字段为空,则以门店保存的配送费为准)
    */
    ShippingFee float64 `json:"shipping_fee"`
    TimeRange string `json:"time_range"`
    /**
    * 配置范围类型，（1表示多个配送范围由多个多边形组成）只支持多边形
    */
    Type string `json:"type"`
}
type ShippingListResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []Data `json:"data"`
    TraceId string `json:"traceId"`
}
type ShippingListRequest struct {
}



func (req *ShippingListRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ShippingListResponse, error) {
    resp, err := client.InvokeApi(shipping_list_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ShippingListResponse
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

func (response *ShippingListResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

