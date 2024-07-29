package updateprinterstatus

/**
* 4.1.1.三方云打印机状态（绑定解绑）回传
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const update_printer_status_url = "/dcps/print/printer/status/update"

type TNoDataResponse struct {
    Resp ResponseResult `json:"resp"`
}
type UpdatePrinterStatusResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data TNoDataResponse `json:"data"`
    TraceId string `json:"traceId"`
}
type ResponseResult struct {
    /**
    * 200
    */
    Code int32 `json:"code"`
    /**
    * 成功
    */
    Msg string `json:"msg"`
}
type UpdatePrinterStatusRequest struct {
    /**
    *  加密门店id 
    */
    OpBizCode string `json:"opBizCode"`
    /**
    *  设备号（sn码） 
    */
    DeviceNumber string `json:"deviceNumber"`
    /**
    *  绑定状态。0解绑；1绑定 
    */
    BindStatus int32 `json:"bindStatus"`
    /**
    *  打印机备注名称 
    */
    Note string `json:"note"`
    /**
    *  打印机服务商类型 
    */
    PrinterMerchantType int32 `json:"printerMerchantType"`
    /**
    *  打印机型号 
    */
    Model string `json:"model"`
    /**
    *  更新时间戳 
    */
    UpdateTime int64 `json:"updateTime"`
}



func (req *UpdatePrinterStatusRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*UpdatePrinterStatusResponse, error) {
    resp, err := client.InvokeApi(update_printer_status_url, 46, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response UpdatePrinterStatusResponse
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

func (response *UpdatePrinterStatusResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

