package orderetamodifyagree

/**
* 同意ETA修改请求
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_eta_modify_agree_url = "/waimai/ng/order/eta/modify/agree"

type OrderEtaModifyAgreeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 同意订单eta修改是否成功
    */
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderEtaModifyAgreeRequest struct {
    OrderId string `json:"orderId"`
}



func (req *OrderEtaModifyAgreeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderEtaModifyAgreeResponse, error) {
    resp, err := client.InvokeApi(order_eta_modify_agree_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderEtaModifyAgreeResponse
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

func (response *OrderEtaModifyAgreeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

