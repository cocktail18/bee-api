package riderlocation

/**
* 获取骑手当前位置
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const rider_location_url = "/peisong/order/rider/location"

type RiderLocationResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data RiderLocationData `json:"data"`
    TraceId string `json:"traceId"`
}
type RiderLocationRequest struct {
    /**
    * 配送活动标识
    */
    DeliveryId int64 `json:"delivery_id"`
    /**
    * 美团配送内部订单 id，最长不超过 32 个字符
    */
    MtPeisongId string `json:"mt_peisong_id"`
}
type RiderLocationData struct {
    /**
    * 纬度
    */
    Lat int32 `json:"lat"`
    /**
    * 经度
    */
    Lng string `json:"lng"`
}



func (req *RiderLocationRequest) DoInvoke(client mtclient.MeituanClient) (*RiderLocationResponse, error) {
    resp, err := client.InvokeApi(rider_location_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response RiderLocationResponse
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

func (response *RiderLocationResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

