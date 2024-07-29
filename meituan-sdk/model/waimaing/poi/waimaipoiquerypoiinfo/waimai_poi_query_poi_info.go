package waimaipoiquerypoiinfo

/**
* 查询门店信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const waimai_poi_query_poi_info_url = "/waimai/poi/queryPoiInfo"

type WaimaiPoiQueryPoiInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []PoiInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type PoiInfo struct {
    /**
    * 门店地址
    */
    Address string `json:"address"`
    /**
    * ERP方门店id 最大长度100
    */
    EPoiId string `json:"ePoiId"`
    /**
    * 发票相关说明
    */
    InvoiceDescription string `json:"invoiceDescription"`
    /**
    * 门店支持开发票的最小订单价
    */
    InvoiceMinPrice int32 `json:"invoiceMinPrice"`
    /**
    * 门店是否支持发票 1-可开发票 0-不支持
    */
    InvoiceSupport int32 `json:"invoiceSupport"`
    /**
    * 是否在线 1-上线 0-下线
    */
    IsOnline int32 `json:"isOnline"`
    /**
    * 是否营业 1可配送 3休息中 0-未上线
    */
    IsOpen int32 `json:"isOpen"`
    /**
    * 门店经度 美团使用的是高德坐标系，也就是火星坐标系，如果是百度坐标系需要转换）(门店坐标不需要乘以一百万）
    */
    Latitude string `json:"latitude"`
    /**
    * 门店纬度 美团使用的是高德坐标系，也就是火星坐标系，如果是百度坐标系需要转换）(门店坐标不需要乘以一百万）
    */
    Longitude string `json:"longitude"`
    /**
    * 门店名称
    */
    Name string `json:"name"`
    /**
    * 门店公告信息 2000字以内
    */
    NoticeInfo string `json:"noticeInfo"`
    /**
    * 营业时间
    */
    OpenTime string `json:"openTime"`
    /**
    * 门店电话
    */
    Phone string `json:"phone"`
    /**
    * 门店图片url
    */
    PictureUrl string `json:"pictureUrl"`
    /**
    * 是否支持营业时间范围外预下单 0-不支持,1 表示支持
    */
    PreBook int32 `json:"preBook"`
    /**
    * 预下单最大日期 2-后天, 预下单支持从最小日期到最大日期的连续日期。例如：preBookMinDays =0，preBookMaxDays=2，表示支持当天，明天，后天的预下单
    */
    PreBookMaxDays int32 `json:"preBookMaxDays"`
    /**
    * 预下单最小日期 0-从当天算起
    */
    PreBookMinDays int32 `json:"preBookMinDays"`
    /**
    * 配送费
    */
    ShippingFee float32 `json:"shippingFee"`
    StandbyTel string `json:"standbyTel"`
    /**
    * 美团品类名称
    */
    TagName string `json:"tagName"`
    /**
    * 是否支持营业时间内预下单 0-不支持,1 表示支持
    */
    TimeSelect int32 `json:"timeSelect"`
    /**
    * 门店的配送方式,参考值：1003-美团跑腿；1001-专送（加盟）；1002-专送（自建）；1004-城市代理；2002-快送；3001-混合送（专送+快送)；2010-全城送；0000-商家自配；5001-聚合配;4015-企客远距离配送。
    */
    LogisticsCodes string `json:"logisticsCodes"`
}
type WaimaiPoiQueryPoiInfoRequest struct {
    /**
    *  ERP方门店id，英文逗号分隔 
    */
    EPoiIds string `json:"ePoiIds"`
}



func (req *WaimaiPoiQueryPoiInfoRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WaimaiPoiQueryPoiInfoResponse, error) {
    resp, err := client.InvokeApi(waimai_poi_query_poi_info_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WaimaiPoiQueryPoiInfoResponse
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

func (response *WaimaiPoiQueryPoiInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

