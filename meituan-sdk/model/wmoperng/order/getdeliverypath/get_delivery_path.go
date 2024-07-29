package getdeliverypath

/**
* 查询众包骑手坐标
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_delivery_path_url = "/wmoper/ng/order/getDeliveryPath"

type GetDeliveryPathResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data []GetDeliveryPathData `json:"data"`
    TraceId string `json:"traceId"`
}
type GetDeliveryPathRequest struct {
    /**
    *  订单ID 
    */
    OrderId int64 `json:"orderId"`
}
type GetDeliveryPathData struct {
    /**
    * 纬度
    */
    Latitude int64 `json:"latitude"`
    /**
    * 经度
    */
    Longitude int64 `json:"longitude"`
    /**
    * 信息采集时间戳
    */
    Time int32 `json:"time"`
}



func (req *GetDeliveryPathRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetDeliveryPathResponse, error) {
    resp, err := client.InvokeApi(get_delivery_path_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetDeliveryPathResponse
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

func (response *GetDeliveryPathResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

