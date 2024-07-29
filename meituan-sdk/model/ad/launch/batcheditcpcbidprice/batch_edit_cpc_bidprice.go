package batcheditcpcbidprice

/**
* 修改推广出价
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const batch_edit_cpc_bidprice_url = "/ad/launch/batchEditCpcBidPrice"

type BatchEditCpcBidpriceResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    Data BatchEditCpcBidpriceData `json:"data"`
    TraceId string `json:"traceId"`
}
type BatchEditCpcBidpriceData struct {
    /**
    * 成功的推广列表
    */
    SuccessLaunchs []int64 `json:"successLaunchs"`
    /**
    * 创建失败的推广map，key为原因，value为推广列表
    */
    FailLaunchs FailLaunch `json:"failLaunchs"`
}
type BatchEditCpcBidpriceRequest struct {
    /**
    * 批量修改出价,单次最对支持50个
    */
    Launchs string `json:"launchs"`
    /**
    *  子账号 
    */
    SonAdaccountId int32 `json:"sonAdaccountId"`
}
type FailLaunch struct {
}



func (req *BatchEditCpcBidpriceRequest) DoInvoke(client mtclient.MeituanClient) (*BatchEditCpcBidpriceResponse, error) {
    resp, err := client.InvokeApi(batch_edit_cpc_bidprice_url, 22, "", req)

    if err != nil {
        return nil, err
    }
    var response BatchEditCpcBidpriceResponse
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

func (response *BatchEditCpcBidpriceResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

