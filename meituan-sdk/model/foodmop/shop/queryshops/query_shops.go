package queryshops

/**
* 门店配置查询（非必接）
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const query_shops_url = "/foodmop/shop/queryShops"

type ResultData struct {
    /**
    * 厂商门店ID
    */
    VendorShopId string `json:"vendorShopId"`
    /**
    * 门店详细配置
    */
    ShopConfig VendorShopConfigDTO `json:"shopConfig"`
    /**
    * 门店查询失败，返回错误原因
    */
    FailMsg string `json:"failMsg"`
}
type VendorShopConfigDTO struct {
    /**
    * 预约配置
    */
    ReserveConfig VendorReserveConfigDTO `json:"reserveConfig"`
    /**
    * 服务配置
    */
    ServiceConfig VendorServiceConfigDTO `json:"serviceConfig"`
    /**
    * 业务配置
    */
    BizConfig VendorBizConfigDTO `json:"bizConfig"`
}
type VendorReserveConfigDTO struct {
    /**
    * 是否支持预约，参数值： TRUE- 支持预约 FALSE- 不支持预约
    */
    MopReserveOpen string `json:"mopReserveOpen"`
    /**
    * MOP可预约时间，参数值范例： {\"1\":[{\"startTime\":\"08:00\",\"endTime\":\"20:00\"}]
    */
    MopReserveHour string `json:"mopReserveHour"`
    /**
    * 预约周期（单位分钟），参数值范例： 5 （ 配置5min，则每小时的0,5,10, ... ）
    */
    MopReservePeriod string `json:"mopReservePeriod"`
    /**
    * 距离现在的最早可预约时间分钟数（单位分钟）,参数值范例 10
    */
    EarliesReserveTimeFromNowByMinutes string `json:"earliesReserveTimeFromNowByMinutes"`
    /**
    * 距离现在的最晚可预约时间分钟数（单位分钟），参数值范例 100
    */
    LatestReserveTimeFromNowByMinutes string `json:"latestReserveTimeFromNowByMinutes"`
}
type QueryShopsResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 厂商门店配置信息列表，特殊说明：假设查询10个ID，9个能查到，List.size()==10, 返回9条VendorShopConfigDTO，1条返回failMsg
    */
    Data []ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type VendorBizConfigDTO struct {
    /**
    * 门店业务开通状态，参数值范例： {\"timestamp\":1675241667962,\"turnOn\":true}
    */
    VendorBizOpenStatus string `json:"vendorBizOpenStatus"`
    /**
    * 门店类型Code值
    */
    ShopType string `json:"shopType"`
    /**
    * 是否交通枢纽，默认false，暂无意义。
    */
    TransportationHub string `json:"transportationHub"`
}
type VendorServiceConfigDTO struct {
    /**
    * 门店营业状态，参数值： 10-营业 20-歇业
    */
    MopBizServiceStatus string `json:"mopBizServiceStatus"`
    /**
    * 营业时间配置 （json格式的字符串），参数值范例： {\"2022-03-30\":[{\"startTime\":\"00:00\",\"endTime\":\"00:06\"}]
    */
    MopBusinessHour string `json:"mopBusinessHour"`
}
type QueryShopsRequest struct {
    /**
    *  厂商门店ID(批量限制20) 
    */
    VendorShopIds []string `json:"vendorShopIds"`
}



func (req *QueryShopsRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*QueryShopsResponse, error) {
    resp, err := client.InvokeApi(query_shops_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response QueryShopsResponse
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

func (response *QueryShopsResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

