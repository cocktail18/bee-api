package orderrearrange

/**
* 模拟改派
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const order_rearrange_url = "/peisong/test/orderRearrange"

type OrderRearrangeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}
type OrderRearrangeRequest struct {
    /**
    * 即配送活动标识，由外部系统生成，不同order_id应对应不同的delivery_id，若因美团系统故障导致发单接口失败，可利用相同的delivery_id重新发单，系统视为同一次配送活动，若更换delivery_id，则系统视为两次独立配送活动。
    */
    DeliveryId int64 `json:"delivery_id"`
    /**
    * 美团配送内部订单id，最长不超过32个字符
    */
    MtPeisongId string `json:"mt_peisong_id"`
}



func (req *OrderRearrangeRequest) DoInvoke(client mtclient.MeituanClient) (*OrderRearrangeResponse, error) {
    resp, err := client.InvokeApi(order_rearrange_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response OrderRearrangeResponse
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

func (response *OrderRearrangeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

