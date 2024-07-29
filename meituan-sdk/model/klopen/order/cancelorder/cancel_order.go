package cancelorder

/**
* 申请取消订单
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const cancel_order_url = "/kl/open/order/cancel"

type CancelOrderRequest struct {
    /**
    *  快驴订单编号 
    */
    OrderId string `json:"orderId"`
    /**
    *  取消原因 
    */
    CancelReason string `json:"cancelReason"`
    /**
    *  取消备注 
    */
    CancelRemark string `json:"cancelRemark"`
    /**
    *  取消人 
    */
    CancelUser string `json:"cancelUser"`
    /**
    *  品牌标识，除华住外其他服务商必传。 
    */
    BrandIdentify string `json:"brandIdentify"`
}
type CancelOrderResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *CancelOrderRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*CancelOrderResponse, error) {
    resp, err := client.InvokeApi(cancel_order_url, 27, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response CancelOrderResponse
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

func (response *CancelOrderResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

