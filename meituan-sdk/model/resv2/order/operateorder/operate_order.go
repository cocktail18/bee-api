package operateorder

/**
* 向开放平台发起订单接单/拒单/确认到店/未到店/分配订单桌位操作
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const operate_order_url = "/resv2/order/operation"

type Data struct {
    /**
    *  订单ID 
    */
    OrderSerializedId string `json:"orderSerializedId"`
    /**
    *  操作类型 1-接单 2-拒单 3-用户到店 4-用户未到店 
    */
    OperationType int32 `json:"operationType"`
    /**
    *  桌位id operationType=1时必填 
    */
    TableId string `json:"tableId"`
    /**
    *  桌位名称 operationType=1时必填 
    */
    TableName string `json:"tableName"`
    /**
    *  时段开始时间 operationType=1时必填 
    */
    StartTime int32 `json:"startTime"`
    /**
    *  时段结束时间 operationType=1时必填 
    */
    EndTime int32 `json:"endTime"`
    /**
    *  实收金额 operationType=3时必填 
    */
    Consume float64 `json:"consume"`
    /**
    *  应收金额 operationType=3时必填 
    */
    OriginPrice float64 `json:"originPrice"`
    /**
    *  拒单原因 1 - 没有匹配的桌型 2 - 该桌型已订满 3 - 不在可预订时间内 4 - 备注无法满足 5 - 其他原因 
    */
    RejectReason int32 `json:"rejectReason"`
    /**
    *  拒单其他原因说明 
    */
    OtherReason string `json:"otherReason"`
}
type OperateOrderResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 返回的数据
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type OperateOrderRequest struct {
    Data Data `json:"data"`
}



func (req *OperateOrderRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*OperateOrderResponse, error) {
    resp, err := client.InvokeApi(operate_order_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response OperateOrderResponse
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

func (response *OperateOrderResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

