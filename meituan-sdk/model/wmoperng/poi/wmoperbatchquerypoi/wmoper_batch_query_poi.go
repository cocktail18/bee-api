package wmoperbatchquerypoi

/**
* 批量获取门店详情信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoper_batch_query_poi_url = "/wmoper/ng/poi/mget"

type WmoperBatchQueryPoiResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []WmoperBatchQueryPoiData `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperBatchQueryPoiRequest struct {
    /**
    *  epoiId集合(半角逗号分隔) 
    */
    EpoiIds string `json:"epoiIds"`
}
type WmoperBatchQueryPoiData struct {
    AppId int64 `json:"app_id"`
    /**
    * epoiId
    */
    EpoiId string `json:"epoiId"`
    /**
    * 门店名称
    */
    Name string `json:"name"`
    /**
    * 门店地址
    */
    Address string `json:"address"`
    /**
    * 门店纬度
    */
    Latitude float64 `json:"latitude"`
    /**
    * 门店经度
    */
    Longitude float64 `json:"longitude"`
    /**
    * 门店图片地址
    */
    PicUrl string `json:"pic_url"`
    /**
    * 客服电话号码
    */
    Phone string `json:"phone"`
    /**
    * 门店电话号码(已废弃)
    */
    StandbyTel string `json:"standby_tel"`
    /**
    * 每个订单的配送费
    */
    ShippingFee float32 `json:"shipping_fee"`
    /**
    * 门店营业时间
    */
    ShippingTime string `json:"shipping_time"`
    /**
    * 门店推广信息
    */
    PromotionInfo string `json:"promotion_info"`
    Remark string `json:"remark"`
    /**
    * 门店的营业状态：1为可配送，3为休息中
    */
    OpenLevel int32 `json:"open_level"`
    /**
    * 门店上下线状态：0为下线，1为上线，2为上单中，3为审核通过可上线
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
    * 发票相关说明（invoice_suport为1时可用）
    */
    InvoiceDescription string `json:"invoice_description"`
    /**
    * 此字段已不再维护，请避免使用
    */
    CityId int64 `json:"city_id"`
    LocationId int64 `json:"location_id"`
    /**
    * 创建时间（当前距离Epoch（1970年1月1日） 以秒计算的时间，即unix-timestamp）
    */
    Ctime int32 `json:"ctime"`
    /**
    * 更新时间（当前距离Epoch（1970年1月1日） 以秒计算的时间，即unix-timestamp）
    */
    Utime int32 `json:"utime"`
    TagName string `json:"tag_name"`
    /**
    * 是否支持营业时间范围外预下单，1表支持，0表不支持
    */
    PreBook int32 `json:"pre_book"`
    /**
    * 是否支持营业时间范围内预下单，1表支持，0表不支持
    */
    TimeSelect int32 `json:"time_select"`
    PreBookMinDays int64 `json:"pre_book_min_days"`
    PreBookMaxDays int64 `json:"pre_book_max_days"`
    /**
    * 门店的配送方式，参考值：1003-美团跑腿；1001-专送（加盟）；1002-专送（自建）；1004-城市代理；2002-快送；3001-混合送（专送+快送)；2010-全城送；0000-商家自配；5001-聚合配。
    */
    LogisticsCodes string `json:"logistics_codes"`
    /**
    * 门店品类
    */
    ThirdTagName string `json:"third_tag_name"`
    /**
    * 门店的配送方式,参考值：1003-美团跑腿；1001-专送（加盟）；1002-专送（自建）；1004-城市代理；2002-快送；3001-混合送（专送+快送)；2010-全城送；0000-商家自配；5001-聚合配;4015-企客远距离配送。
    */
    LogisticsCode string `json:"logistics_code"`
}



func (req *WmoperBatchQueryPoiRequest) DoInvoke(client mtclient.MeituanClient) (*WmoperBatchQueryPoiResponse, error) {
    resp, err := client.InvokeApi(wmoper_batch_query_poi_url, 16, "", req)

    if err != nil {
        return nil, err
    }
    var response WmoperBatchQueryPoiResponse
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

func (response *WmoperBatchQueryPoiResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

