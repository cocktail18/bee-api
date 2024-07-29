package getrecipientinfo

/**
* 配送类服务商提供收货人的信息
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const get_recipient_info_url = "/wmoper/ng/delivery/getRecipientInfo"

type GetRecipientInfoRequest struct {
    /**
    *  订单号 
    */
    OrderId string `json:"orderId"`
}
type GetRecipientInfoResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data GetRecipientInfoData `json:"data"`
    TraceId string `json:"traceId"`
}
type GetRecipientInfoData struct {
    /**
    * 收货人姓名
    */
    RecipientName string `json:"recipientName"`
    /**
    * 收货人手机号，C端启用号码保护时该字段展示隐私号，如15696424612_7472。C端未启用号码保护或隐私号降级时，该字段为真实手机号，如12345678901
    */
    RecipientPhone string `json:"recipientPhone"`
    /**
    * 收货人地址。订单完成3小时后展示“为保护顾客隐私，地址已隐藏”。实际的地址@#后是经纬度反查的地址，是用户订餐时定位的地址。
    */
    RecipientAddress string `json:"recipientAddress"`
    /**
    * 脱敏收货人地址。订单完成3小时后展示“为保护顾客隐私，地址已隐藏”
    */
    RecipientAddressDesensitization string `json:"recipientAddressDesensitization"`
    /**
    * 实际送餐地址纬度，美团使用的是高德坐标系
    */
    Latitude float64 `json:"latitude"`
    /**
    * 实际送餐地址经度，美团使用的是高德坐标系
    */
    Longitude float64 `json:"longitude"`
    /**
    * 收货人备份隐私号，具体为json格式的字符串数组 1、当recipientPhone字段为隐私号时，该字段作为其备份，目前为1个号码，将来为可能为多个，用英文逗号隔开“,”。 2、有至少2个可用的隐私号时，该字段返回备用的隐私号；如果只有1个隐私号可用或所有隐私号均不可用时，该字段返回内容为空。 3、recipientPhone和backupRecipientPhone的隐私号间没有优劣之分，任何一个隐私号故障时，请尝试用其他隐私号联系用户。当所有隐私号都故障且美团触发降级时，重新获取的recipientPhone字段才会展示真实号。
    */
    BackupRecipientPhone string `json:"backupRecipientPhone"`
    /**
    * 订单状态，订单状态。2新订单，4已接单，8订单完成，9订单取消
    */
    Status int32 `json:"status"`
}



func (req *GetRecipientInfoRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*GetRecipientInfoResponse, error) {
    resp, err := client.InvokeApi(get_recipient_info_url, 16, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response GetRecipientInfoResponse
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

func (response *GetRecipientInfoResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

