package tuangoureceiptquerybymobile

/**
* 手机号查询可用团购券
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const tuangou_receipt_querybymobile_url = "/ddzh/tuangou/receipt/querybymobile"

type TuangouReceiptQuerybymobileRequest struct {
    /**
    *  点评用户绑定的手机号 国内手机即11位手机号，国外手机格式为【地区码_手机号】 例如国内手机：13866666666，国外手机：65_666666 
    */
    Mobile string `json:"mobile"`
    /**
    *  团购ID 
    */
    DealGroupId int64 `json:"dealGroupId"`
    /**
    *  套餐ID 
    */
    DealId int64 `json:"dealId"`
    /**
    *  起始位置，从0开始 
    */
    Offset int32 `json:"offset"`
    /**
    *  查询数量，最大不超过30 
    */
    Limit int32 `json:"limit"`
    /**
    *  平台：1-大众点评，2-美团 
    */
    Platform int32 `json:"platform"`
}
type TuangouReceiptQuerybymobileResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data Data `json:"data"`
    TraceId string `json:"traceId"`
}
type UserReceiptDTO struct {
    /**
    * 团购券码
    */
    SerialNumber string `json:"serialNumber"`
}
type Data struct {
    Result []UserReceiptDTO `json:"result"`
}



func (req *TuangouReceiptQuerybymobileRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*TuangouReceiptQuerybymobileResponse, error) {
    resp, err := client.InvokeApi(tuangou_receipt_querybymobile_url, 58, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response TuangouReceiptQuerybymobileResponse
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

func (response *TuangouReceiptQuerybymobileResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

