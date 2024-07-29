package lifereservemerchantmodify

/**
* 商家改约
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const lifereserve_merchantmodify_url = "/ddzh/yuding/lifereserve/merchantmodify"

type LifereserveMerchantmodifyResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type Data struct {
    /**
    * 改约结果
    */
    Result bool `json:"result"`
}
type LifereserveMerchantmodifyRequest struct {
    /**
    *  预约单号 
    */
    ReserveOrderId int64 `json:"reserveOrderId"`
    /**
    *  预约人姓名 
    */
    ReserveUserName string `json:"reserveUserName"`
    /**
    *  预约人电话 
    */
    ReservePhone string `json:"reservePhone"`
    /**
    *  预约地址 
    */
    ReserveAddress string `json:"reserveAddress"`
    /**
    *  预约起始时间（毫秒数） 
    */
    BookStartTime int64 `json:"bookStartTime"`
    /**
    *  预约结束时间 
    */
    BookEndTime int64 `json:"bookEndTime"`
    /**
    *  送件姓名 
    */
    DeliveryUserName string `json:"deliveryUserName"`
    /**
    *  送件地址 
    */
    DeliveryAddress string `json:"deliveryAddress"`
    /**
    *  送件电话 
    */
    DeliveryPhone string `json:"deliveryPhone"`
}



func (req *LifereserveMerchantmodifyRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*LifereserveMerchantmodifyResponse, error) {
    resp, err := client.InvokeApi(lifereserve_merchantmodify_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response LifereserveMerchantmodifyResponse
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

func (response *LifereserveMerchantmodifyResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

