package wmoperngquerypoidetail

/**
* 获取门店详情信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const wmoperng_query_poi_detail_url = "/wmoper/ng/poi/detail"

type WmoperngQueryPoiDetailResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []Data `json:"data"`
    TraceId string `json:"traceId"`
}
type WmoperngQueryPoiDetailRequest struct {
    Biz string `json:"biz"`
}
type Data struct {
    AppId int64 `json:"app_id"`
    EpoiId string `json:"epoiId"`
    Name string `json:"name"`
    Address string `json:"address"`
    Latitude float64 `json:"latitude"`
    Longitude float64 `json:"longitude"`
    PicUrl string `json:"pic_url"`
    Phone string `json:"phone"`
    StandbyTel string `json:"standby_tel"`
    ShippingFee float64 `json:"shipping_fee"`
    ShippingTime string `json:"shipping_time"`
    PromotionInfo string `json:"promotion_info"`
    Remark string `json:"remark"`
    OpenLevel int64 `json:"open_level"`
    IsOnline int64 `json:"is_online"`
    InvoiceSupport int64 `json:"invoice_support"`
    InvoiceMinPrice int64 `json:"invoice_min_price"`
    InvoiceDescription string `json:"invoice_description"`
    CityId int64 `json:"city_id"`
    LocationId int64 `json:"location_id"`
    Ctime int64 `json:"ctime"`
    Utime int64 `json:"utime"`
    TagName string `json:"tag_name"`
    PreBook int64 `json:"pre_book"`
    TimeSelect int64 `json:"time_select"`
    PreBookMinDays int64 `json:"pre_book_min_days"`
    PreBookMaxDays int64 `json:"pre_book_max_days"`
    LogisticsCodes string `json:"logistics_codes"`
}



func (req *WmoperngQueryPoiDetailRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WmoperngQueryPoiDetailResponse, error) {
    resp, err := client.InvokeApi(wmoperng_query_poi_detail_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WmoperngQueryPoiDetailResponse
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

func (response *WmoperngQueryPoiDetailResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

