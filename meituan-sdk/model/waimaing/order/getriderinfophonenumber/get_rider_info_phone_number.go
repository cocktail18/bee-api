package getriderinfophonenumber

/**
* 隐私号-批量拉取骑手手机号
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_rider_info_phone_number_url = "/waimai/order/getRiderInfoPhoneNumber"

type GetRiderInfoPhoneNumberRequest struct {
    DeveloperId int64 `json:"developerId"`
    /**
    *  分页查询的偏移量 
    */
    DegradOffset int32 `json:"degradOffset"`
    /**
    *  每页条数，需小于等于1000 
    */
    DegradLimit int32 `json:"degradLimit"`
}
type GetRiderInfoPhoneNumberResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []RiderPhoneNumberInfo `json:"data"`
    TraceId string `json:"traceId"`
}
type RiderPhoneNumberInfo struct {
    /**
    * ERP门店ID
    */
    EPoiId string `json:"ePoiId"`
    /**
    * 订单Id
    */
    OrderId int64 `json:"orderId"`
    /**
    * 订单展示Id
    */
    OrderIdView int64 `json:"orderIdView"`
    /**
    * 骑手名字
    */
    RiderName string `json:"riderName"`
    /**
    * 骑手真实手机号
    */
    RiderRealPhoneNumber string `json:"riderRealPhoneNumber"`
}



func (req *GetRiderInfoPhoneNumberRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetRiderInfoPhoneNumberResponse, error) {
    resp, err := client.InvokeApi(get_rider_info_phone_number_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetRiderInfoPhoneNumberResponse
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

func (response *GetRiderInfoPhoneNumberResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

