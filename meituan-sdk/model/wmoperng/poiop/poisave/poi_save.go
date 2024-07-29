package poisave

/**
* 创建或更新门店信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const poi_save_url = "/wmoper/ng/poiop/poi/save"

type PoiSaveRequest struct {
    /**
    * 门店名称
    */
    Name string `json:"name"`
    /**
    * 门店地址
    */
    Address string `json:"address"`
    /**
    * 门店纬度，（美团使用的是高德坐标系，也就是火星坐标系，如果是百度坐标系需要转换），(门店坐标不需要乘以一百万)
    */
    Latitude float32 `json:"latitude"`
    /**
    * 门店经度，（美团使用的是高德坐标系，也就是火星坐标系，如果是百度坐标系需要转换），(门店坐标不需要乘以一百万)
    */
    Longitude float32 `json:"longitude"`
    /**
    * 门店头图图片地址，图片比例1:1，需要为JPG/JPEG格式，大于256x256
    */
    PicUrl string `json:"pic_url"`
    /**
    * 门店头图图片地址，图片比例4:3，需要为JPG/JPEG格式，大于400x300
    */
    PicUrlLarge string `json:"pic_url_large"`
    /**
    * 客服电话号码 （注意：此号码留客服号码）
    */
    Phone string `json:"phone"`
    /**
    * 门店电话号码 （注意：此号码留商家电话）(已废弃)
    */
    StandbyTel string `json:"standby_tel"`
    /**
    * 每个订单的配送费
    */
    ShippingFee float32 `json:"shipping_fee"`
    /**
    * 门店营业时间 （注意格式，且保证不同时间段之间不存在交集）
    */
    ShippingTime string `json:"shipping_time"`
    /**
    * 门店推广信息
    */
    PromotionInfo string `json:"promotion_info"`
    /**
    * 门店的营业状态：1为可配送，3为休息中
    */
    OpenLevel int32 `json:"open_level"`
    /**
    * 门店上下线状态：1为上线，0为下线，门店必须在菜品、配送范围和门店信息都齐全后，才能提交上线（注意：此字段不为1时门店不会提交审核）
    */
    IsOnline int32 `json:"is_online"`
    /**
    * 门店是否支持发票
    */
    InvoiceSupport int32 `json:"invoice_support"`
    /**
    * 门店支持开发票的最小订单价（invoice_suport为1时可用）
    */
    InvoiceMinPrice int32 `json:"invoice_min_price"`
    /**
    * 发票相关说明（invoice_support为1时可用）
    */
    InvoiceDescription string `json:"invoice_description"`
    /**
    * 目前最多支持上传两个门店品类：一个主营品类（上传的第一个品类为主营品类），一个辅营品类；部分门店品类只支持上传一个品类（例如：火锅，特色菜，地方菜，东南亚菜，日韩料理，生活超市）
    */
    ThirdTagName string `json:"third_tag_name"`
    /**
    * 是否支持营业时间范围外预下单。1表支持，0表不支持
    */
    PreBook int32 `json:"pre_book"`
    /**
    * 是否支持营业时间范围内预下单。1表支持，0表不支持。此字段开启后不可关闭
    */
    TimeSelect string `json:"time_select"`
}
type PoiSaveResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *PoiSaveRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*PoiSaveResponse, error) {
    resp, err := client.InvokeApi(poi_save_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response PoiSaveResponse
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

func (response *PoiSaveResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

