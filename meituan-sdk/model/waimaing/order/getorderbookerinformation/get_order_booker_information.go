package getorderbookerinformation

/**
* 查询订单预订人隐私信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_order_booker_information_url = "/waimai/ng/order/getOrderBookerInformation"

type ResultData struct {
    /**
    * 订单Id
    */
    WmOrderId int32 `json:"wmOrderId"`
    /**
    * 预订人手机号
    */
    ReservationPhone string `json:"reservationPhone"`
}
type GetOrderBookerInformationResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ResultData `json:"data"`
    TraceId string `json:"traceId"`
}
type GetOrderBookerInformationRequest struct {
    /**
    *  订单id 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  查询类型。1蛋糕品类，2其他。 
    */
    Type int32 `json:"type"`
    /**
    *  补充原因。当type=2时为必填项 
    */
    SupplementaryReasons string `json:"supplementaryReasons"`
}



func (req *GetOrderBookerInformationRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetOrderBookerInformationResponse, error) {
    resp, err := client.InvokeApi(get_order_booker_information_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetOrderBookerInformationResponse
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

func (response *GetOrderBookerInformationResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

