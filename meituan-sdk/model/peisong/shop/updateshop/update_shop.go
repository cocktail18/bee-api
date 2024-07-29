package updateshop

/**
* 修改门店
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const update_shop_url = "/peisong/shop/update"

type UpdateShopData struct {
    /**
    * 门店id
    */
    ShopId string `json:"shop_id"`
    /**
    * 状态码
    */
    Status int32 `json:"status"`
}
type UpdateShopResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data UpdateShopData `json:"data"`
    TraceId string `json:"traceId"`
}
type UpdateShopRequest struct {
    /**
    *  取货门店id，即合作方向美团提供的门店id 
    */
    ShopId string `json:"shop_id"`
    /**
    *  门店名称 说明：①门店名称格式请按照 【XX品牌-XX店 or XX品牌（XX店）】填写，例：百果园（望京店），括号内信息必须以【店】结尾；②名称最大不能超过50个字符。 注：该名称需与实体门店门牌保持一致，保证骑手取货可确认门店。 
    */
    ShopName string `json:"shop_name"`
    /**
    *  门店联系人姓名 
    */
    ContactName string `json:"contact_name"`
    /**
    *  联系电话 说明：支持手机号、支持400电话、支持固话（格式为【区号】【-】【座机号】【-】【分机号】，分机号如有则添加）。 
    */
    ContactPhone string `json:"contact_phone"`
    /**
    *  联系邮箱 
    */
    ContactEmail string `json:"contact_email"`
    /**
    *  门店地址 说明：①地址长度不得小于5个字符大于60个字符；②地址信息不得含有折扣/满减信息。 注：请提供真实准确的门店地址，便于骑手取货。 
    */
    ShopAddress string `json:"shop_address"`
    /**
    *  门牌号 
    */
    ShopAddressDetail string `json:"shop_address_detail"`
    /**
    *  门店经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 116398419 说明：门店坐标与地址所在城市、行政区等需保持一致。 注：请提供准确坐标，便于骑手取货。 
    */
    ShopLng int32 `json:"shop_lng"`
    /**
    *  门店纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 39985005 说明：门店坐标与地址所在城市、行政区等需保持一致。 注：请提供准确坐标，便于骑手取货。 
    */
    ShopLat int32 `json:"shop_lat"`
    /**
    *  坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标 （默认值为0） 
    */
    CoordinateType int32 `json:"coordinate_type"`
    /**
    *  配送服务代码，详情见合同 1）服务包 飞速达:4002 快速达:4011 及时达:4012 集中送:4013 跑腿B帮送:4031 例如：4011,4012(多个英文逗号隔开) 2）新服务产品 具体可参考新服务产品列表 
    */
    DeliveryServiceCodes string `json:"delivery_service_codes"`
    /**
    *  营业时间 例：[{"beginTime":"00:00","endTime":"24:00"}] 注：传入后美团根据区域可配送时间取交集时间作为门店配送时间 
    */
    BusinessHours string `json:"business_hours"`
}



func (req *UpdateShopRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*UpdateShopResponse, error) {
    resp, err := client.InvokeApi(update_shop_url, 19, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response UpdateShopResponse
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

func (response *UpdateShopResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

