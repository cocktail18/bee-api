package ordermodify

/**
* 订单修改
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_modify_url = "/peisong/order/modify"

type OrderModifyRequest struct {
    /**
    * 美团配送内部订单id，最长不超过32个字符
    */
    MtPeisongId string `json:"mt_peisong_id"`
    /**
    * 收件人电话，最长不超过64个字符
    */
    ReceiverPhone string `json:"receiver_phone"`
    /**
    * 收件人地址详情，如永兴路365号，最长不超过100个字符
    */
    ReceiverDetailAddress string `json:"receiver_detail_address"`
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
    * 备注，最长不超过200个字符
    */
    Note string `json:"note"`
}
type OrderModifyResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *OrderModifyRequest) DoInvoke(client mtclient.MeituanClient) (*OrderModifyResponse, error) {
    resp, err := client.InvokeApi(order_modify_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response OrderModifyResponse
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

func (response *OrderModifyResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

