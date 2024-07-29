package createshop

/**
* 创建配送门店
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const create_shop_url = "/peisong/shop/create"

type CreateShopRequest struct {
    /**
    * 取货门店id，即合作方向美团提供的门店id
    */
    ShopId string `json:"shop_id"`
    /**
    * 门店名称 说明：门店名称格式请按照 【XX品牌-XX店】填写，例：百果园-望京店，注：该名称需与实体门店门牌保持一致，保证骑手取货可确认门店。
    */
    ShopName string `json:"shop_name"`
    /**
    * 一级品类，见附件品类代码表 说明：品类需按门店真实配送品类选择，如无法判断可咨询您的销售经理。
    */
    Category int32 `json:"category"`
    /**
    * 二级品类，见附件品类代码表 说明：品类需按门店真实配送品类选择，如无法判断可咨询您的销售经理。
    */
    SecondCategory int32 `json:"second_category"`
    /**
    * 门店联系人姓名
    */
    ContactName string `json:"contact_name"`
    /**
    * 联系电话
    */
    ContactPhone string `json:"contact_phone"`
    /**
    * 联系邮箱
    */
    ContactEmail string `json:"contact_email"`
    /**
    * 门店地址
    */
    ShopAddress string `json:"shop_address"`
    /**
    * 门牌号
    */
    ShopAddressDetail string `json:"shop_address_detail"`
    /**
    * 门店经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 116398419 说明：请提供准确坐标，便于骑手取货
    */
    ShopLng int32 `json:"shop_lng"`
    /**
    * 门店纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 39985005 说明：请提供准确坐标，便于骑手取货
    */
    ShopLat int32 `json:"shop_lat"`
    /**
    * 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标 （默认值为0）
    */
    CoordinateType int32 `json:"coordinate_type"`
    /**
    * 配送服务代码，详情见合同 1）服务包 飞速达:4002 快速达:4011 及时达:4012 集中送:4013 跑腿B帮送:4031 例如：4011,4012(多个英文逗号隔开) 2）新服务产品 具体可参考新服务产品列表
    */
    DeliveryServiceCodes string `json:"delivery_service_codes"`
    /**
    * 营业时间 例：[{"beginTime":"00:00","endTime":"24:00"}] 注：传入后美团根据区域可配送时间取交集时间作为门店配送时间
    */
    BusinessHours string `json:"business_hours"`
}
type CreateShopData struct {
    /**
    * 取货门店id
    */
    ShopId string `json:"shop_id"`
    /**
    * 状态码
    */
    Status int32 `json:"status"`
}
type CreateShopResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data CreateShopData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *CreateShopRequest) DoInvoke(client mtclient.MeituanClient) (*CreateShopResponse, error) {
    resp, err := client.InvokeApi(create_shop_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response CreateShopResponse
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

func (response *CreateShopResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

