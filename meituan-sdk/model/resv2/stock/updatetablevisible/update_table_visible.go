package updatetablevisible

/**
* 清台/锁台
* This file was automatically generated.
*/
import (
    "meituan/gosdk/constants"
    "meituan/gosdk/api_error"
    "meituan/gosdk/mtclient"
    "encoding/json"
)

const update_table_visible_url = "/resv2/stock/supply/updateTableVisible"

type VendorSpuDTO struct {
    /**
    *  日期 秒级时间戳 
    */
    Day int32 `json:"day"`
    /**
    *  美团桌型ID，推荐该字段必传，或者 「tableType、tableTypeName」两个字段必传 
    */
    TableTypeId int64 `json:"tableTypeId"`
    /**
    *  桌型类型，0-大厅 1-包间 3-自定义 
    */
    TableType int32 `json:"tableType"`
    /**
    *  桌型名称 
    */
    TableTypeName string `json:"tableTypeName"`
    /**
    *  餐段, 1234早中晚夜 
    */
    Period int32 `json:"period"`
    /**
    *  1:锁台 2:清台 
    */
    Type int32 `json:"type"`
}
type UpdateTableVisibleRequest struct {
    VendorSpuDTO VendorSpuDTO `json:"vendorSpuDTO"`
    Uuid string `json:"uuid"`
}
type UpdateTableVisibleResponse struct {
    Code string `json:"code"`
    Msg string `json:"msg"`
    /**
    * 操作结果
    */
    Data bool `json:"data"`
    TraceId string `json:"traceId"`
}



func (req *UpdateTableVisibleRequest) DoInvoke(client mtclient.MeituanClient, appAuthToken string) (*UpdateTableVisibleResponse, error) {
    resp, err := client.InvokeApi(update_table_visible_url, 7, appAuthToken, req)

    if err != nil {
        return nil, err
    }
    var response UpdateTableVisibleResponse
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

func (response *UpdateTableVisibleResponse) IsSuccess() bool {
    return response.Code == constants.SuccessCode
}

