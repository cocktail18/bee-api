package updateorderfulfillinfo

/**
* 更新订单履约信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const updateorderfulfillinfo_url = "/ddzh/yuding/updateorderfulfillinfo"

type UpdateorderfulfillinfoRequest struct {
    /**
    *  订单id 
    */
    OrderId string `json:"orderId"`
    /**
    *  扩展信息 
    */
    AttrKeyValue string `json:"attrKeyValue"`
    /**
    *  生活服务行业新增字段，其他行业不传。交易类型，枚举值：1-预订，2-团单（团购后预约），商家根据此属性字段识别交易类型，预订业务该字段非必传，团购后预约业务必传 
    */
    Type int32 `json:"type"`
    /**
    *  履约状态： 已分配（通用）：40001 已取件（洗衣）：40101 已到达（通用）：40002 清洗中（洗衣）：41101 返回中（洗衣）：41102 履约完成（通用）：41001 
    */
    OrderFulfillStatus int32 `json:"orderFulfillStatus"`
}
type UpdateorderfulfillinfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *UpdateorderfulfillinfoRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*UpdateorderfulfillinfoResponse, error) {
    resp, err := client.InvokeApi(updateorderfulfillinfo_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response UpdateorderfulfillinfoResponse
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

func (response *UpdateorderfulfillinfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

