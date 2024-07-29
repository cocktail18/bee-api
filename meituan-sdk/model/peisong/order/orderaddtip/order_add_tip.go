package orderaddtip

/**
* 增加小费接口
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_add_tip_url = "/peisong/order/addTip"

type OrderAddTipRequest struct {
    /**
    * 配送活动标识
    */
    DeliveryId int64 `json:"delivery_id"`
    /**
    * 美团配送内部订单id，最长不超过32个字符
    */
    MtPeisongId string `json:"mt_peisong_id"`
    /**
    * 小费金额，精确到元，需要为1或1的倍数，上限20，允许每个运单最多加5次；
    */
    TipAmount int32 `json:"tip_amount"`
    /**
    * 加小费传入的唯一标识，用来幂等处理，最长不超过128个字符
    */
    SerialNumber string `json:"serial_number"`
}
type OrderAddTipResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *OrderAddTipRequest) DoInvoke(client mtclient.MeituanClient) (*OrderAddTipResponse, error) {
    resp, err := client.InvokeApi(order_add_tip_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response OrderAddTipResponse
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

func (response *OrderAddTipResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

