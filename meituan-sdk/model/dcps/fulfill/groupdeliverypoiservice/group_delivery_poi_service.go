package groupdeliverypoiservice

/**
* 团购配送门店服务接口，用于获取门店信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const group_delivery_poi_service_url = "/dcps/fulfill/poi/getPoiInfo"

type GroupDeliveryPoiServiceResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type GroupDeliveryPoiServiceRequest struct {
}
type Data struct {
    /**
    * 授权实体标识
    */
    OpBizCode string `json:"opBizCode"`
    /**
    * 授权实体名称
    */
    Name string `json:"name"`
}



func (req *GroupDeliveryPoiServiceRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GroupDeliveryPoiServiceResponse, error) {
    resp, err := client.InvokeApi(group_delivery_poi_service_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GroupDeliveryPoiServiceResponse
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

func (response *GroupDeliveryPoiServiceResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

