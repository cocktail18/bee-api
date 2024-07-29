package orderchangenotice

/**
* 订单修改审核结果通知
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_change_notice_url = "/foodmop/order/orderChangeNotice"

type OrderChangeNoticeRequest struct {
    /**
    *  MT订单id 
    */
    OrderId int64 `json:"orderId"`
    /**
    *  请求时间戳 
    */
    Timestamp int64 `json:"timestamp"`
    /**
    *  商户订单号 
    */
    VendorOrderId string `json:"vendorOrderId"`
    /**
    *  订单修改类型： TAKE_TIME(1, "取餐时间") ORDER_TYPE_FROM_BOOKING_TO_REAL(2, "预约单改成实时单") 
    */
    OrderModifyType int32 `json:"orderModifyType"`
    /**
    *  订单修改结果： SUCCESS(0, "成功") FAIL(1, "失败") 
    */
    OrderModifyStatus int32 `json:"orderModifyStatus"`
    /**
    *  预计取餐时间，普通单必传 
    */
    TakeTime int64 `json:"takeTime"`
    /**
    *  操作者Id 
    */
    OperatorId string `json:"operatorId"`
    /**
    *  操作者名称 
    */
    OperatorName string `json:"operatorName"`
    /**
    *  预期制作时间，预约单必传 
    */
    MakeTime int64 `json:"makeTime"`
    /**
    *  修改失败原因： ORDER_NOT_FUND(1, "未查询到订单"), NOT_BOOKING_ORDER(2, "非预约单"), REPEAT_MODIFY(3, "不能重复修改") 
    */
    RejectType int32 `json:"rejectType"`
    /**
    *  拒绝修改理由 
    */
    RejectMessage string `json:"rejectMessage"`
    /**
    *  处理时间 
    */
    HandleTime int64 `json:"handleTime"`
    /**
    *  业务类型 10-MOP 
    */
    BizType int32 `json:"bizType"`
}
type OrderChangeNoticeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *OrderChangeNoticeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OrderChangeNoticeResponse, error) {
    resp, err := client.InvokeApi(order_change_notice_url, 51, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OrderChangeNoticeResponse
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

func (response *OrderChangeNoticeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

