package waimaiimgetreadtime

/**
* 查询会话最新已读时间戳
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const waimai_im_get_read_time_url = "/waimai/ng/im/userReadTime"

type WaimaiImGetReadTimeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data WaimaiImGetReadTimeData `json:"data"`
    TraceId string `json:"traceId"`
}
type WaimaiImGetReadTimeRequest struct {
    Biz string `json:"biz"`
}
type WaimaiImGetReadTimeData struct {
    /**
    * 开发者应用标识
    */
    AppId int64 `json:"app_id"`
    /**
    * ERP 商户门店ID
    */
    EpoiId string `json:"epoiId"`
    /**
    * 美团C端用户ID
    */
    OpenUserId string `json:"open_user_id"`
    /**
    * 毫秒时间戳
    */
    LastReadTime int64 `json:"last_read_time"`
}



func (req *WaimaiImGetReadTimeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*WaimaiImGetReadTimeResponse, error) {
    resp, err := client.InvokeApi(waimai_im_get_read_time_url, 2, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response WaimaiImGetReadTimeResponse
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

func (response *WaimaiImGetReadTimeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

