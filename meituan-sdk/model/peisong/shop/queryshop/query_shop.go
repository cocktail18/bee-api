package queryshop

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

const query_shop_url = "/peisong/shop/query"

type QueryShopData struct {
    /**
    * 取货门店id，即合作方向美团提供的门店id
    */
    ShopId string `json:"shop_id"`
    /**
    * 城市ID，见城市代码表
    */
    City int32 `json:"city"`
    /**
    * 一级品类，见附件品类代码表
    */
    Category int32 `json:"category"`
    /**
    * 二级品类，见附件品类代码表
    */
    SecondCategory int32 `json:"second_category"`
    /**
    * 门店联系人
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
    * 门店经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 116398419
    */
    ShopLng int32 `json:"shop_lng"`
    /**
    * 门店纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 39985005
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
    * 配送时间 例：[{"beginTime":"00:00","endTime":"24:00"}]
    */
    DeliveryHours string `json:"delivery_hours"`
    /**
    * 是否支持预约单，0：不支持，1：支持
    */
    Prebook int32 `json:"prebook"`
    /**
    * 是否支持营业时间外预约单，0：不支持，1：支持
    */
    PrebookOutOfBizTime int32 `json:"prebook_out_of_biz_time"`
    /**
    * 预约单时间段，格式为{"start": "0", "end": "2"}，单位为天
    */
    PrebookPeriod string `json:"prebook_period"`
    /**
    * 门店当前可支持的结算方式下的支付方式，支付方式，0、账期支付，1、余额支付； 账期支付请关注月度账单；余额支付请联系运营协助操作账户开通、充值等，开通余额支付时可参与相应营销活动、使用优惠券等。
    */
    PayTypeCodes []int32 `json:"pay_type_codes"`
}
type QueryShopRequest struct {
    /**
    * 取货门店id，即合作方向美团提供的门店id
    */
    ShopId string `json:"shop_id"`
}
type QueryShopResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data QueryShopData `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *QueryShopRequest) DoInvoke(client mtclient.MeituanClient) (*QueryShopResponse, error) {
    resp, err := client.InvokeApi(query_shop_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response QueryShopResponse
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

func (response *QueryShopResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

