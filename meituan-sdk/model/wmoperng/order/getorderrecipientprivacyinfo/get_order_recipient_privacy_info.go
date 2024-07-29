package getorderrecipientprivacyinfo

/**
* 查询订单收餐人隐私信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_order_recipient_privacy_info_url = "/wmoper/ng/order/getOrderRecipientPrivacyInfo"

type GetOrderRecipientPrivacyInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type ResultData struct {
    /**
    * 订单Id
    */
    WmOrderId int64 `json:"wmOrderId"`
    /**
    * 收餐人手机号。当下单用户关闭隐私号保护，为收餐人真实手机号码，示例：13000220022。 当下单用户开启隐私号保护，为收餐人手机号码隐私号格式，示例：13101230000_1234。
    */
    RecipientPhone string `json:"recipientPhone"`
    /**
    * 收餐人脱敏真实号
    */
    RecipientPhoneDesensitization string `json:"recipientPhoneDesensitization"`
    /**
    * 收餐人脱敏地址
    */
    RecipientAddressDesensitization string `json:"recipientAddressDesensitization"`
}
type GetOrderRecipientPrivacyInfoRequest struct {
    /**
    *  订单id 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  查询类型。 1仅查询收餐人电话（含脱敏真实号）， 2仅查询收餐人脱敏地址， 3同时查询。 
    */
    Type int32 `json:"type"`
    /**
    *  原因类型。 1接单前联系用户-核对订单信息， 2接单前联系用户-商品缺货协商更换或退款， 3接单前联系用户-协商发货时间， 4接单后联系用户-商品缺货协商更换或退款， 5接单后联系用户-协商发货时间， 6接单后联系用户-告知出餐或发货进度， 7订单完成后联系用户-补发商品， 8订单完成后联系用户-售后回访， 9配送过程联系用户。 
    */
    ReasonType int32 `json:"reasonType"`
}



func (req *GetOrderRecipientPrivacyInfoRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetOrderRecipientPrivacyInfoResponse, error) {
    resp, err := client.InvokeApi(get_order_recipient_privacy_info_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetOrderRecipientPrivacyInfoResponse
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

func (response *GetOrderRecipientPrivacyInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

