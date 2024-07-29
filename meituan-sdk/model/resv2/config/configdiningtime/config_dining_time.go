package configdiningtime

/**
* 新增或修改门店的营业时段信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const config_dining_time_url = "/resv2/config/update"

type ConfigDiningTimeResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 返回的数据
    */
    Data string `json:"data"`
    TraceId string `json:"traceId"`
}
type ConfigDiningTimeRequest struct {
    Data []Data `json:"data"`
}
type Data struct {
    /**
    *  星期 
    */
    Weeks []int32 `json:"weeks"`
    /**
    *  具体配置 
    */
    DiningTime []DiningTime `json:"diningTime"`
}
type DiningTime struct {
    /**
    *  时段类型 
    */
    Type int32 `json:"type"`
    /**
    *  开始小时 
    */
    StartHour int32 `json:"startHour"`
    /**
    *  开始分钟 
    */
    StartMinute int32 `json:"startMinute"`
    /**
    *  结束小时 
    */
    EndHour int32 `json:"endHour"`
    /**
    *  结束分钟 
    */
    EndMinute int32 `json:"endMinute"`
}



func (req *ConfigDiningTimeRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*ConfigDiningTimeResponse, error) {
    resp, err := client.InvokeApi(config_dining_time_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response ConfigDiningTimeResponse
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

func (response *ConfigDiningTimeResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

