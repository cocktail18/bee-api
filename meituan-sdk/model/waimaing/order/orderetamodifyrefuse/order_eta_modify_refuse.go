package orderetamodifyrefuse

/**
* 拒绝ETA修改请求
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_eta_modify_refuse_url = "/waimai/ng/order/eta/modify/refuse"

type OrderEtaModifyRefuseResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 拒绝订单eta修改是否成功
    */
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}
type OrderEtaModifyRefuseRequest struct {
    /**
    *  订单Id 
    */
    OrderId string `json:"orderId"`
}



func (req *OrderEtaModifyRefuseRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderEtaModifyRefuseResponse, error) {
    resp, err := client.InvokeApi(order_eta_modify_refuse_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderEtaModifyRefuseResponse
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

func (response *OrderEtaModifyRefuseResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

