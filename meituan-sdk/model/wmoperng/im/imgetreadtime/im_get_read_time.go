package imgetreadtime

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

const im_get_read_time_url = "/wmoper/ng/im/userReadTime"

type ImGetReadTimeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data ImGetReadTimeData `json:"data"`
    TraceId string `json:"traceId"`
}
type ImGetReadTimeData struct {
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
type ImGetReadTimeRequest struct {
    /**
    *  用户ID 
    */
    OpenUserId int64 `json:"openUserId"`
}



func (req *ImGetReadTimeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ImGetReadTimeResponse, error) {
    resp, err := client.InvokeApi(im_get_read_time_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ImGetReadTimeResponse
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

func (response *ImGetReadTimeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

