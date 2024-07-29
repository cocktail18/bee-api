package ordergetrealrecipientaddress

/**
* 查询真实地址接口
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_get_real_recipient_address_url = "/waimai/ng/order/getRealRecipientAddress"

type OrderGetRealRecipientAddressRequest struct {
    /**
    *  订单ID 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  查询原因类型枚举值： 1：商家自有运力。 2：美团运力转商家自有运力。 0：其他。 
    */
    QueryReasonType int32 `json:"queryReasonType"`
    /**
    *  查询原因补充说明： 当【查询原因类型=其他】时，必须填写原因， 即query_reason_type=0时，必填。 
    */
    AddOtherReason string `json:"addOtherReason"`
}
type OrderGetRealRecipientAddressResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data OrderGetRealRecipientAddressData `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderGetRealRecipientAddressData struct {
    /**
    * 订单Id
    */
    OrderId string `json:"order_id"`
    /**
    * 收货人地址。 订单完成3小时后展示“为保护顾客隐私，地址已隐藏”
    */
    RecipientAddress string `json:"recipient_address"`
}



func (req *OrderGetRealRecipientAddressRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderGetRealRecipientAddressResponse, error) {
    resp, err := client.InvokeApi(order_get_real_recipient_address_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderGetRealRecipientAddressResponse
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

func (response *OrderGetRealRecipientAddressResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

