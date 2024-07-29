package check

/**
* 配送能力校验
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const check_url = "/peisong/order/check"

type CheckResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}
type CheckRequest struct {
    /**
    * 取货门店 id，即合作方向美团提供的门店 id。 注：测试门店的 shop_id 固定为 test_0001，仅用于对接时联调测试。
    */
    ShopId string `json:"shop_id"`
    /**
    * 配送服务代码，详情见合同 1）服务包 飞速达:4002 快速达:4011 及时达:4012 集中送:4013 跑腿B帮送:4031 2）新服务产品 具体可参考 新服务产品列表
    */
    DeliveryServiceCode int32 `json:"delivery_service_code"`
    /**
    * 收件人地址，最长不超过 512 个字符
    */
    ReceiverAddress string `json:"receiver_address"`
    /**
    * 收件人经度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 116398419
    */
    ReceiverLng int32 `json:"receiver_lng"`
    /**
    * 收件人纬度（火星坐标或百度坐标，和 coordinate_type 字段配合使用），坐标 * （10的六次方），如 39985005
    */
    ReceiverLat int32 `json:"receiver_lat"`
    /**
    * 坐标类型，0：火星坐标（高德，腾讯地图均采用火星坐标） 1：百度坐标 （默认值为0）
    */
    CoordinateType int32 `json:"coordinate_type"`
    /**
    * 预留字段，方便以后扩充校验规则，check_type = 1
    */
    CheckType int32 `json:"check_type"`
    /**
    * 模拟发单时间，时区为 GMT+8，当前距离 Epoch（1970年1月1日) 以秒计算的时间，即 unix-timestamp。
    */
    MockOrderTime int64 `json:"mock_order_time"`
}



func (req *CheckRequest) DoInvoke(client mtclient.MeituanClient) (*CheckResponse, error) {
    resp, err := client.InvokeApi(check_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response CheckResponse
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

func (response *CheckResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

