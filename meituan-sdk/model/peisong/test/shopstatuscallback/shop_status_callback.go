package shopstatuscallback

/**
* 模拟门店状态回调测试
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const shop_status_callback_url = "/peisong/test/shopStatusCallback"

type ShopStatusCallbackRequest struct {
    /**
    * 取货门店id，即合作方向美团提供的门店id
    */
    ShopId string `json:"shop_id"`
    /**
    * 10-审核驳回 20-审核通过 30-创建成功 40-上线可发单
    */
    Status int32 `json:"status"`
}
type ShopStatusCallbackResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    TraceId string `json:"traceId"`
}



func (req *ShopStatusCallbackRequest) DoInvoke(client mtclient.MeituanClient) (*ShopStatusCallbackResponse, error) {
    resp, err := client.InvokeApi(shop_status_callback_url, 19, "", req)

    if err != nil {
        return nil, err
    }
    var response ShopStatusCallbackResponse
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

func (response *ShopStatusCallbackResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

